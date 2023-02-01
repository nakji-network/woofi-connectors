package woofi

import (
	"context"
	"math/rand"
	"sync/atomic"
	"time"

	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/chain/ethereum"
	"github.com/nakji-network/connector/kafkautils"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const TimeoutDuration = 5 * time.Minute

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

// Start starts subscribing and doing backfill at the same time
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

// Backfill does backfill only
func (c *Connector) Backfill() {
	var (
		events    []proto.Message
		addresses []ethcommon.Address
	)

	for _, contract := range c.contracts {
		events = append(events, contract.Events()...)
		addresses = append(addresses, ethcommon.HexToAddress(contract.Address()))
	}

	c.Connector = ethereum.NewConnector(context.Background(), addresses, c.NetworkName, connector.BackfillOption())

	c.RegisterProtos(kafkautils.MsgTypeBf, events...)
	go c.backfill(addresses)
	c.handleBfChannel()
}

func (c *Connector) handleBfChannel() {
	for msg := range c.processBfLogs() {
		c.EventSink <- &kafkautils.Message{
			MsgType:  kafkautils.MsgTypeBf,
			ProtoMsg: msg,
		}
	}
	log.Info().Msg("backfill stopped")
}

func (c *Connector) handleFctChannel() {
	for msg := range c.processFctLogs() {
		c.EventSink <- &kafkautils.Message{
			MsgType:  kafkautils.MsgTypeFct,
			ProtoMsg: msg,
		}
	}
}

func (c *Connector) backfill(addresses []ethcommon.Address) {
	defer atomic.StoreInt32(&c.bfDone, 1)
	log.Info().Msg("backfill started")
	logs, err := ethereum.HistoricalEventsWithQueryParams(context.Background(), c.Client, addresses, c.FromBlock, c.NumBlocks)
	if err != nil {
		log.Error().Err(err).Msg("backfill failed")
		return
	}
	for vLog := range logs {
		c.bfChan <- vLog
	}
}

func (c *Connector) parse(contract ISmartContract, vLog types.Log) (protoreflect.ProtoMessage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), TimeoutDuration)
	defer cancel()

	ts, err := c.getBlockTime(ctx, vLog)
	if err != nil {
		log.Warn().Err(err).Str("blockHash", vLog.BlockHash.String()).Msg("failed to retrieve block timestamp, retry later...")
		return nil, err
	}

	tx, err := c.transactionByHash(ctx, vLog.TxHash)
	if err != nil {
		log.Warn().Err(err).Str("txHash", vLog.TxHash.String()).Msg("failed to retrieve transaction by hash, retry later...")
		return nil, err
	}

	sender, err := c.transactionSender(ctx, tx, vLog.BlockHash, vLog.TxIndex)
	if err != nil {
		log.Warn().Err(err).Str("txHash", vLog.TxHash.String()).Msg("failed to retrieve transaction sender, retry later...")
		return nil, err
	}

	msg := contract.Message(Log{
		Log:             vLog,
		BlockTime:       time.Unix(int64(ts), 0),
		SenderAddress:   &sender,
		ReceiverAddress: tx.To(),
	})

	return msg, nil
}

func (c *Connector) processBfLogs() <-chan protoreflect.ProtoMessage {
	out := make(chan protoreflect.ProtoMessage)
	go func() {
		for {
			select {
			case vLog := <-c.bfChan:
				contract := c.GetContract(vLog.Address.String())
				if contract == nil {
					log.Error().Str("address", vLog.Address.String()).Msg("event from unsupported address")
					continue
				}
				msg, err := c.parse(contract, vLog)
				if err != nil {
					c.bfChan <- vLog // Put the log back to the end of the queue to retry later
					continue
				}
				out <- msg
			case <-time.After(100 * time.Millisecond):
				// Exit when backfill is done and there is no message left to process
				if atomic.LoadInt32(&c.bfDone) == 1 {
					close(out)
					return
				}
			}
		}
	}()
	return out
}

func (c *Connector) processFctLogs() <-chan protoreflect.ProtoMessage {
	out := make(chan protoreflect.ProtoMessage)
	go func() {
		for vLog := range c.fctChan {
			contract := c.GetContract(vLog.Address.String())
			if contract == nil {
				log.Error().Str("address", vLog.Address.String()).Msg("event from unsupported address")
				continue
			}
			msg, err := c.parse(contract, vLog)
			if err != nil {
				c.fctChan <- vLog // Put the log back to the end of the queue to retry later
				continue
			}
			out <- msg
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
