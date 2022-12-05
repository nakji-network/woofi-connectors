package WooPPV4

import (
	"strings"

	"github.com/nakji-network/connector/common"
	"github.com/nakji-network/woofi-connectors/woofi"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SmartContract struct {
	abi       abi.ABI
	addr      string
	addrBytes []byte
}

func NewContract(addr string) *SmartContract {
	contractAbi, err := abi.JSON(strings.NewReader(WooPPV4ABI))
	if err != nil {
		log.Fatal().Err(err).Msg("error reading WooPPV4ABI")
	}
	return &SmartContract{abi: contractAbi, addr: addr, addrBytes: ethcommon.HexToAddress(addr).Bytes()}
}

func (sc *SmartContract) Address() string {
	return sc.addr
}

func (sc *SmartContract) Events() []proto.Message {
	return []proto.Message{
		&AdminUpdated{},
		&FeeAddrUpdated{},
		&Paused{},
		&Unpaused{},
		&WooracleUpdated{},
		&Deposit{},
		&OwnershipTransferred{},
		&WooSwap{},
		&Migrate{},
		&Withdraw{},
	}
}

func (sc *SmartContract) Message(vLog woofi.Log) proto.Message {
	ev, err := sc.abi.EventByID(vLog.Topics[0])
	if err != nil {
		log.Warn().Err(err).Msg("EventByID error, skipping")
		return nil
	}
	switch ev.Name {
	case "AdminUpdated":
		e := new(WooPPV4AdminUpdated)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		msg := &AdminUpdated{
			Ts:              timestamppb.New(vLog.BlockTime),
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			Addr:            e.Addr.Bytes(),
			Flag:            e.Flag,
			ContractAddress: sc.addrBytes,
		}
		if vLog.SenderAddress != nil {
			msg.SenderAddress = vLog.SenderAddress.Bytes()
		}
		if vLog.ReceiverAddress != nil {
			msg.ReceiverAddress = vLog.ReceiverAddress.Bytes()
		}
		return msg
	case "FeeAddrUpdated":
		e := new(WooPPV4FeeAddrUpdated)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		msg := &FeeAddrUpdated{
			Ts:              timestamppb.New(vLog.BlockTime),
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			NewFeeAddr:      e.NewFeeAddr.Bytes(),
			ContractAddress: sc.addrBytes,
		}
		if vLog.SenderAddress != nil {
			msg.SenderAddress = vLog.SenderAddress.Bytes()
		}
		if vLog.ReceiverAddress != nil {
			msg.ReceiverAddress = vLog.ReceiverAddress.Bytes()
		}
		return msg
	case "Paused":
		e := new(WooPPV4Paused)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		msg := &Paused{
			Ts:              timestamppb.New(vLog.BlockTime),
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			Account:         e.Account.Bytes(),
			ContractAddress: sc.addrBytes,
		}
		if vLog.SenderAddress != nil {
			msg.SenderAddress = vLog.SenderAddress.Bytes()
		}
		if vLog.ReceiverAddress != nil {
			msg.ReceiverAddress = vLog.ReceiverAddress.Bytes()
		}
		return msg
	case "Unpaused":
		e := new(WooPPV4Unpaused)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		msg := &Unpaused{
			Ts:              timestamppb.New(vLog.BlockTime),
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			Account:         e.Account.Bytes(),
			ContractAddress: sc.addrBytes,
		}
		if vLog.SenderAddress != nil {
			msg.SenderAddress = vLog.SenderAddress.Bytes()
		}
		if vLog.ReceiverAddress != nil {
			msg.ReceiverAddress = vLog.ReceiverAddress.Bytes()
		}
		return msg
	case "WooracleUpdated":
		e := new(WooPPV4WooracleUpdated)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		msg := &WooracleUpdated{
			Ts:              timestamppb.New(vLog.BlockTime),
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			NewWooracle:     e.NewWooracle.Bytes(),
			ContractAddress: sc.addrBytes,
		}
		if vLog.SenderAddress != nil {
			msg.SenderAddress = vLog.SenderAddress.Bytes()
		}
		if vLog.ReceiverAddress != nil {
			msg.ReceiverAddress = vLog.ReceiverAddress.Bytes()
		}
		return msg
	case "Deposit":
		e := new(WooPPV4Deposit)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		msg := &Deposit{
			Ts:              timestamppb.New(vLog.BlockTime),
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			Token:           e.Token.Bytes(),
			Sender:          e.Sender.Bytes(),
			Amount:          e.Amount.Bytes(),
			ContractAddress: sc.addrBytes,
		}
		if vLog.SenderAddress != nil {
			msg.SenderAddress = vLog.SenderAddress.Bytes()
		}
		if vLog.ReceiverAddress != nil {
			msg.ReceiverAddress = vLog.ReceiverAddress.Bytes()
		}
		return msg
	case "OwnershipTransferred":
		e := new(WooPPV4OwnershipTransferred)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		msg := &OwnershipTransferred{
			Ts:              timestamppb.New(vLog.BlockTime),
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			PreviousOwner:   e.PreviousOwner.Bytes(),
			NewOwner:        e.NewOwner.Bytes(),
			ContractAddress: sc.addrBytes,
		}
		if vLog.SenderAddress != nil {
			msg.SenderAddress = vLog.SenderAddress.Bytes()
		}
		if vLog.ReceiverAddress != nil {
			msg.ReceiverAddress = vLog.ReceiverAddress.Bytes()
		}
		return msg
	case "WooSwap":
		e := new(WooPPV4WooSwap)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		msg := &WooSwap{
			Ts:              timestamppb.New(vLog.BlockTime),
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			FromToken:       e.FromToken.Bytes(),
			ToToken:         e.ToToken.Bytes(),
			FromAmount:      e.FromAmount.Bytes(),
			ToAmount:        e.ToAmount.Bytes(),
			From:            e.From.Bytes(),
			To:              e.To.Bytes(),
			RebateTo:        e.RebateTo.Bytes(),
			SwapVol:         e.SwapVol.Bytes(),
			SwapFee:         e.SwapFee.Bytes(),
			ContractAddress: sc.addrBytes,
		}
		if vLog.SenderAddress != nil {
			msg.SenderAddress = vLog.SenderAddress.Bytes()
		}
		if vLog.ReceiverAddress != nil {
			msg.ReceiverAddress = vLog.ReceiverAddress.Bytes()
		}
		return msg
	case "Migrate":
		e := new(WooPPV4Migrate)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		msg := &Migrate{
			Ts:              timestamppb.New(vLog.BlockTime),
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			Token:           e.Token.Bytes(),
			Receiver:        e.Receiver.Bytes(),
			Amount:          e.Amount.Bytes(),
			ContractAddress: sc.addrBytes,
		}
		if vLog.SenderAddress != nil {
			msg.SenderAddress = vLog.SenderAddress.Bytes()
		}
		if vLog.ReceiverAddress != nil {
			msg.ReceiverAddress = vLog.ReceiverAddress.Bytes()
		}
		return msg
	case "Withdraw":
		e := new(WooPPV4Withdraw)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		msg := &Withdraw{
			Ts:              timestamppb.New(vLog.BlockTime),
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			Token:           e.Token.Bytes(),
			Receiver:        e.Receiver.Bytes(),
			Amount:          e.Amount.Bytes(),
			ContractAddress: sc.addrBytes,
		}
		if vLog.SenderAddress != nil {
			msg.SenderAddress = vLog.SenderAddress.Bytes()
		}
		if vLog.ReceiverAddress != nil {
			msg.ReceiverAddress = vLog.ReceiverAddress.Bytes()
		}
		return msg
	default:
		log.Error().Msgf("invalid event: %s", ev.Name)
		return nil
	}
}
