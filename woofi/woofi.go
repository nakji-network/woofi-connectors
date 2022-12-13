package woofi

import (
	"context"
	"math/rand"
	"sync/atomic"
	"time"

	"github.com/nakji-network/connector/chain/ethereum"
	"github.com/nakji-network/connector/kafkautils"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Config struct {
	NetworkName string
	FromBlock   uint64
	NumBlocks   uint64
}

type Connector struct {
	*Config
	*ethereum.Connector
	contracts map[string]ISmartContract
	bfDone    int32
	bfChan    chan types.Log
	fctChan   chan types.Log
}

func New(config *Config) *Connector {
	return &Connector{
		Config:    config,
		contracts: make(map[string]ISmartContract),
		bfChan:    make(chan types.Log, 1024),
		fctChan:   make(chan types.Log, 1024),
	}
}

func (c *Connector) AddContract(sc ISmartContract) {
	c.contracts[sc.Address()] = sc
}

func (c *Connector) GetContract(addr string) ISmartContract {
	return c.contracts[addr]
}

func (c *Connector) Start() {
	var (
		events    []proto.Message
		addresses []ethcommon.Address
	)

	for _, contract := range c.contracts {
		events = append(events, contract.Events()...)
		addresses = append(addresses, ethcommon.HexToAddress(contract.Address()))
	}

	c.Connector = ethereum.NewConnector(context.Background(), addresses, c.NetworkName)

	c.RegisterProtos(kafkautils.MsgTypeFct, events...)

	c.Sub.Subscribe(context.Background())

	if c.FromBlock > 0 || c.NumBlocks > 0 {
		c.RegisterProtos(kafkautils.MsgTypeBf, events...)
		go c.backfill(addresses)
		go c.handleBfChannel()
	}

	go c.handleFctChannel()

	for {
		select {
		case <-c.Sub.Done():
			log.Info().Msg("connector shutdown")
			return

		//	Listen to error channel
		case err := <-c.Sub.Err():
			log.Error().Err(err).Str("network", c.NetworkName).Msg("subscription failed")
			return

		//	Listen to event logs
		case vLog := <-c.Sub.Logs():
			c.fctChan <- vLog.Log
		}
	}
}

func (c *Connector) handleBfChannel() {
	msgs := c.processLogs(c.bfChan)
loop:
	for {
		select {
		case msg := <-msgs:
			c.EventSink <- &kafkautils.Message{
				MsgType:  kafkautils.MsgTypeBf,
				ProtoMsg: msg,
			}
		default:
			// Exit when backfill is done and there is no message left to process
			if atomic.LoadInt32(&c.bfDone) == 1 {
				log.Info().Msg("backfill stopped")
				break loop
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (c *Connector) handleFctChannel() {
	for msg := range c.processLogs(c.fctChan) {
		c.EventSink <- &kafkautils.Message{
			MsgType:  kafkautils.MsgTypeFct,
			ProtoMsg: msg,
		}
	}
}

func (c *Connector) backfill(addresses []ethcommon.Address) {
	log.Info().Msg("backfill started")
	logs, err := ethereum.HistoricalEventsWithQueryParams(context.Background(), c.Client, addresses, c.FromBlock, c.NumBlocks)
	if err != nil {
		log.Error().Err(err).Msg("backfill failed")
		return
	}
	for vLog := range logs {
		c.bfChan <- vLog
	}
	atomic.StoreInt32(&c.bfDone, 1)
}

func (c *Connector) processLogs(logs chan types.Log) <-chan protoreflect.ProtoMessage {
	out := make(chan protoreflect.ProtoMessage)
	go func() {
		for vLog := range logs {
			contract := c.GetContract(vLog.Address.String())
			if contract == nil {
				log.Error().Str("address", vLog.Address.String()).Msg("event from unsupported address")
				continue
			}

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)

			ts, err := c.getBlockTime(ctx, vLog)
			if err != nil {
				log.Warn().Err(err).Str("blockHash", vLog.BlockHash.String()).Msg("failed to retrieve block timestamp, retry later...")
				logs <- vLog
				cancel()
				continue
			}

			tx, err := c.transactionByHash(ctx, vLog.TxHash)
			if err != nil {
				log.Warn().Err(err).Str("txHash", vLog.TxHash.String()).Msg("failed to retrieve transaction by hash, retry later...")
				logs <- vLog
				cancel()
				continue
			}

			sender, err := c.transactionSender(ctx, tx, vLog.BlockHash, vLog.TxIndex)
			if err != nil {
				log.Warn().Err(err).Str("txHash", vLog.TxHash.String()).Msg("failed to retrieve transaction sender, retry later...")
				logs <- vLog
				cancel()
				continue
			}

			cancel()

			out <- contract.Message(Log{
				Log:             vLog,
				BlockTime:       time.Unix(int64(ts), 0),
				SenderAddress:   &sender,
				ReceiverAddress: tx.To(),
			})
		}
	}()
	return out
}

func (c *Connector) getBlockTime(ctx context.Context, vLog types.Log) (uint64, error) {
	for retry := 0; ; retry++ {
		ts, err := c.Sub.GetBlockTime(ctx, vLog)
		if err == nil {
			return ts, nil
		}
		if !isRetryableError(err) {
			return 0, err
		}
		// Do exponential backoff with 10% jitter
		backoff := float64(int(1) << retry)
		backoff += backoff * (0.1 * rand.Float64())
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		case <-time.After(time.Second * time.Duration(backoff)):
			continue
		}
	}
}

func (c *Connector) transactionByHash(ctx context.Context, hash ethcommon.Hash) (*types.Transaction, error) {
	for retry := 0; ; retry++ {
		tx, err := c.Sub.TransactionByHash(ctx, hash)
		if err == nil {
			return tx, nil
		}
		if !isRetryableError(err) {
			return nil, err
		}
		// Do exponential backoff with 10% jitter
		backoff := float64(int(1) << retry)
		backoff += backoff * (0.1 * rand.Float64())
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(time.Second * time.Duration(backoff)):
			continue
		}
	}
}

func (c *Connector) transactionSender(ctx context.Context, tx *types.Transaction, block ethcommon.Hash, index uint) (ethcommon.Address, error) {
	for retry := 0; ; retry++ {
		sender, err := c.Sub.TransactionSender(ctx, tx, block, index)
		if err == nil {
			return sender, nil
		}
		if !isRetryableError(err) {
			return ethcommon.Address{}, err
		}
		// Do exponential backoff with 10% jitter
		backoff := float64(int(1) << retry)
		backoff += backoff * (0.1 * rand.Float64())
		select {
		case <-ctx.Done():
			return ethcommon.Address{}, ctx.Err()
		case <-time.After(time.Second * time.Duration(backoff)):
			continue
		}
	}
}

func isRetryableError(err error) bool {
	s := err.Error()
	return s == "not found" || s == "websocket: close 1006 (abnormal closure): unexpected EOF"
}
