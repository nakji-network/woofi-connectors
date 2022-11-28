package WOOPP

import (
	"strings"

	"github.com/nakji-network/connector/common"
	"github.com/nakji-network/woofi-connectors/woofi"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SmartContract struct {
	abi  abi.ABI
	addr string
}

func NewContract(addr string) *SmartContract {
	contractAbi, err := abi.JSON(strings.NewReader(WOOPPABI))
	if err != nil {
		log.Fatal().Err(err).Msg("error reading WooPPV2ABI")
	}
	return &SmartContract{abi: contractAbi, addr: addr}
}

func (sc *SmartContract) Address() string {
	return sc.addr
}

func (sc *SmartContract) Events() []proto.Message {
	return []proto.Message{
		&Paused{},
		&OwnershipTransferred{},
		&ParametersUpdated{},
		&FeeManagerUpdated{},
		&RewardManagerUpdated{},
		&Unpaused{},
		&WooGuardianUpdated{},
		&WooSwap{},
		&OwnershipTransferPrepared{},
		&WooracleUpdated{},
		&StrategistUpdated{},
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
	case "Paused":
		e := new(WOOPPPaused)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Paused{
			Ts:          timestamppb.New(vLog.BlockTime),
			BlockNumber: vLog.BlockNumber,
			Index:       uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
			Account:     e.Account.Bytes(),
		}
	case "OwnershipTransferred":
		e := new(WOOPPOwnershipTransferred)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &OwnershipTransferred{
			Ts:            timestamppb.New(vLog.BlockTime),
			BlockNumber:   vLog.BlockNumber,
			Index:         uint64(vLog.Index),
			TxHash:        vLog.TxHash.Bytes(),
			PreviousOwner: e.PreviousOwner.Bytes(),
			NewOwner:      e.NewOwner.Bytes(),
		}
	case "ParametersUpdated":
		e := new(WOOPPParametersUpdated)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &ParametersUpdated{
			Ts:           timestamppb.New(vLog.BlockTime),
			BlockNumber:  vLog.BlockNumber,
			Index:        uint64(vLog.Index),
			TxHash:       vLog.TxHash.Bytes(),
			BaseToken:    e.BaseToken.Bytes(),
			NewThreshold: e.NewThreshold.Bytes(),
			NewR:         e.NewR.Bytes(),
		}
	case "FeeManagerUpdated":
		e := new(WOOPPFeeManagerUpdated)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &FeeManagerUpdated{
			Ts:            timestamppb.New(vLog.BlockTime),
			BlockNumber:   vLog.BlockNumber,
			Index:         uint64(vLog.Index),
			TxHash:        vLog.TxHash.Bytes(),
			NewFeeManager: e.NewFeeManager.Bytes(),
		}
	case "RewardManagerUpdated":
		e := new(WOOPPRewardManagerUpdated)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &RewardManagerUpdated{
			Ts:               timestamppb.New(vLog.BlockTime),
			BlockNumber:      vLog.BlockNumber,
			Index:            uint64(vLog.Index),
			TxHash:           vLog.TxHash.Bytes(),
			NewRewardManager: e.NewRewardManager.Bytes(),
		}
	case "Unpaused":
		e := new(WOOPPUnpaused)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Unpaused{
			Ts:          timestamppb.New(vLog.BlockTime),
			BlockNumber: vLog.BlockNumber,
			Index:       uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
			Account:     e.Account.Bytes(),
		}
	case "WooGuardianUpdated":
		e := new(WOOPPWooGuardianUpdated)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &WooGuardianUpdated{
			Ts:             timestamppb.New(vLog.BlockTime),
			BlockNumber:    vLog.BlockNumber,
			Index:          uint64(vLog.Index),
			TxHash:         vLog.TxHash.Bytes(),
			NewWooGuardian: e.NewWooGuardian.Bytes(),
		}
	case "WooSwap":
		e := new(WOOPPWooSwap)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &WooSwap{
			Ts:          timestamppb.New(vLog.BlockTime),
			BlockNumber: vLog.BlockNumber,
			Index:       uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
			FromToken:   e.FromToken.Bytes(),
			ToToken:     e.ToToken.Bytes(),
			FromAmount:  e.FromAmount.Bytes(),
			ToAmount:    e.ToAmount.Bytes(),
			From:        e.From.Bytes(),
			To:          e.To.Bytes(),
			RebateTo:    e.RebateTo.Bytes(),
		}
	case "OwnershipTransferPrepared":
		e := new(WOOPPOwnershipTransferPrepared)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &OwnershipTransferPrepared{
			Ts:            timestamppb.New(vLog.BlockTime),
			BlockNumber:   vLog.BlockNumber,
			Index:         uint64(vLog.Index),
			TxHash:        vLog.TxHash.Bytes(),
			PreviousOwner: e.PreviousOwner.Bytes(),
			NewOwner:      e.NewOwner.Bytes(),
		}
	case "WooracleUpdated":
		e := new(WOOPPWooracleUpdated)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &WooracleUpdated{
			Ts:          timestamppb.New(vLog.BlockTime),
			BlockNumber: vLog.BlockNumber,
			Index:       uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
			NewWooracle: e.NewWooracle.Bytes(),
		}
	case "StrategistUpdated":
		e := new(WOOPPStrategistUpdated)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &StrategistUpdated{
			Ts:          timestamppb.New(vLog.BlockTime),
			BlockNumber: vLog.BlockNumber,
			Index:       uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
			Strategist:  e.Strategist.Bytes(),
			Flag:        e.Flag,
		}
	case "Withdraw":
		e := new(WOOPPWithdraw)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Withdraw{
			Ts:          timestamppb.New(vLog.BlockTime),
			BlockNumber: vLog.BlockNumber,
			Index:       uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
			Token:       e.Token.Bytes(),
			To:          e.To.Bytes(),
			Amount:      e.Amount.Bytes(),
		}
	}
	return nil
}
