// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package WooPPV4

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// WooPPV4MetaData contains all meta data concerning the WooPPV4 contract.
var WooPPV4MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_quoteToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"AdminUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFeeAddr\",\"type\":\"address\"}],\"name\":\"FeeAddrUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Migrate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rebateTo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapVol\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"WooSwap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newWooracle\",\"type\":\"address\"}],\"name\":\"WooracleUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"depositAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wooracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeAddr\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newPool\",\"type\":\"address\"}],\"name\":\"migrateToNewPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"poolSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"}],\"name\":\"query\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIWooLendingManager\",\"name\":\"lendManager\",\"type\":\"address\"}],\"name\":\"repayWeeklyLending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeAddr\",\"type\":\"address\"}],\"name\":\"setFeeAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"name\":\"setFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wooracle\",\"type\":\"address\"}],\"name\":\"setWooracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"skimMulTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minToAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rebateTo\",\"type\":\"address\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"realToAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenInfos\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"reserve\",\"type\":\"uint192\"},{\"internalType\":\"uint16\",\"name\":\"feeRate\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"}],\"name\":\"tryQuery\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unclaimedFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wooracle\",\"outputs\":[{\"internalType\":\"contractIWooracleV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// WooPPV4ABI is the input ABI used to generate the binding from.
// Deprecated: Use WooPPV4MetaData.ABI instead.
var WooPPV4ABI = WooPPV4MetaData.ABI

// WooPPV4 is an auto generated Go binding around an Ethereum contract.
type WooPPV4 struct {
	WooPPV4Caller     // Read-only binding to the contract
	WooPPV4Transactor // Write-only binding to the contract
	WooPPV4Filterer   // Log filterer for contract events
}

// WooPPV4Caller is an auto generated read-only Go binding around an Ethereum contract.
type WooPPV4Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WooPPV4Transactor is an auto generated write-only Go binding around an Ethereum contract.
type WooPPV4Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WooPPV4Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WooPPV4Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WooPPV4Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WooPPV4Session struct {
	Contract     *WooPPV4          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WooPPV4CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WooPPV4CallerSession struct {
	Contract *WooPPV4Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// WooPPV4TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WooPPV4TransactorSession struct {
	Contract     *WooPPV4Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// WooPPV4Raw is an auto generated low-level Go binding around an Ethereum contract.
type WooPPV4Raw struct {
	Contract *WooPPV4 // Generic contract binding to access the raw methods on
}

// WooPPV4CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WooPPV4CallerRaw struct {
	Contract *WooPPV4Caller // Generic read-only contract binding to access the raw methods on
}

// WooPPV4TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WooPPV4TransactorRaw struct {
	Contract *WooPPV4Transactor // Generic write-only contract binding to access the raw methods on
}

// NewWooPPV4 creates a new instance of WooPPV4, bound to a specific deployed contract.
func NewWooPPV4(address common.Address, backend bind.ContractBackend) (*WooPPV4, error) {
	contract, err := bindWooPPV4(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WooPPV4{WooPPV4Caller: WooPPV4Caller{contract: contract}, WooPPV4Transactor: WooPPV4Transactor{contract: contract}, WooPPV4Filterer: WooPPV4Filterer{contract: contract}}, nil
}

// NewWooPPV4Caller creates a new read-only instance of WooPPV4, bound to a specific deployed contract.
func NewWooPPV4Caller(address common.Address, caller bind.ContractCaller) (*WooPPV4Caller, error) {
	contract, err := bindWooPPV4(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WooPPV4Caller{contract: contract}, nil
}

// NewWooPPV4Transactor creates a new write-only instance of WooPPV4, bound to a specific deployed contract.
func NewWooPPV4Transactor(address common.Address, transactor bind.ContractTransactor) (*WooPPV4Transactor, error) {
	contract, err := bindWooPPV4(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WooPPV4Transactor{contract: contract}, nil
}

// NewWooPPV4Filterer creates a new log filterer instance of WooPPV4, bound to a specific deployed contract.
func NewWooPPV4Filterer(address common.Address, filterer bind.ContractFilterer) (*WooPPV4Filterer, error) {
	contract, err := bindWooPPV4(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WooPPV4Filterer{contract: contract}, nil
}

// bindWooPPV4 binds a generic wrapper to an already deployed contract.
func bindWooPPV4(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WooPPV4ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WooPPV4 *WooPPV4Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WooPPV4.Contract.WooPPV4Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WooPPV4 *WooPPV4Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WooPPV4.Contract.WooPPV4Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WooPPV4 *WooPPV4Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WooPPV4.Contract.WooPPV4Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WooPPV4 *WooPPV4CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WooPPV4.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WooPPV4 *WooPPV4TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WooPPV4.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WooPPV4 *WooPPV4TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WooPPV4.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address token) view returns(uint256)
func (_WooPPV4 *WooPPV4Caller) Balance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WooPPV4.contract.Call(opts, &out, "balance", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address token) view returns(uint256)
func (_WooPPV4 *WooPPV4Session) Balance(token common.Address) (*big.Int, error) {
	return _WooPPV4.Contract.Balance(&_WooPPV4.CallOpts, token)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address token) view returns(uint256)
func (_WooPPV4 *WooPPV4CallerSession) Balance(token common.Address) (*big.Int, error) {
	return _WooPPV4.Contract.Balance(&_WooPPV4.CallOpts, token)
}

// FeeAddr is a free data retrieval call binding the contract method 0x39e7fddc.
//
// Solidity: function feeAddr() view returns(address)
func (_WooPPV4 *WooPPV4Caller) FeeAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WooPPV4.contract.Call(opts, &out, "feeAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeAddr is a free data retrieval call binding the contract method 0x39e7fddc.
//
// Solidity: function feeAddr() view returns(address)
func (_WooPPV4 *WooPPV4Session) FeeAddr() (common.Address, error) {
	return _WooPPV4.Contract.FeeAddr(&_WooPPV4.CallOpts)
}

// FeeAddr is a free data retrieval call binding the contract method 0x39e7fddc.
//
// Solidity: function feeAddr() view returns(address)
func (_WooPPV4 *WooPPV4CallerSession) FeeAddr() (common.Address, error) {
	return _WooPPV4.Contract.FeeAddr(&_WooPPV4.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_WooPPV4 *WooPPV4Caller) IsAdmin(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _WooPPV4.contract.Call(opts, &out, "isAdmin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_WooPPV4 *WooPPV4Session) IsAdmin(arg0 common.Address) (bool, error) {
	return _WooPPV4.Contract.IsAdmin(&_WooPPV4.CallOpts, arg0)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_WooPPV4 *WooPPV4CallerSession) IsAdmin(arg0 common.Address) (bool, error) {
	return _WooPPV4.Contract.IsAdmin(&_WooPPV4.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WooPPV4 *WooPPV4Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WooPPV4.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WooPPV4 *WooPPV4Session) Owner() (common.Address, error) {
	return _WooPPV4.Contract.Owner(&_WooPPV4.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WooPPV4 *WooPPV4CallerSession) Owner() (common.Address, error) {
	return _WooPPV4.Contract.Owner(&_WooPPV4.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WooPPV4 *WooPPV4Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _WooPPV4.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WooPPV4 *WooPPV4Session) Paused() (bool, error) {
	return _WooPPV4.Contract.Paused(&_WooPPV4.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WooPPV4 *WooPPV4CallerSession) Paused() (bool, error) {
	return _WooPPV4.Contract.Paused(&_WooPPV4.CallOpts)
}

// PoolSize is a free data retrieval call binding the contract method 0xfa75d160.
//
// Solidity: function poolSize(address token) view returns(uint256)
func (_WooPPV4 *WooPPV4Caller) PoolSize(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WooPPV4.contract.Call(opts, &out, "poolSize", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolSize is a free data retrieval call binding the contract method 0xfa75d160.
//
// Solidity: function poolSize(address token) view returns(uint256)
func (_WooPPV4 *WooPPV4Session) PoolSize(token common.Address) (*big.Int, error) {
	return _WooPPV4.Contract.PoolSize(&_WooPPV4.CallOpts, token)
}

// PoolSize is a free data retrieval call binding the contract method 0xfa75d160.
//
// Solidity: function poolSize(address token) view returns(uint256)
func (_WooPPV4 *WooPPV4CallerSession) PoolSize(token common.Address) (*big.Int, error) {
	return _WooPPV4.Contract.PoolSize(&_WooPPV4.CallOpts, token)
}

// Query is a free data retrieval call binding the contract method 0xf58a435f.
//
// Solidity: function query(address fromToken, address toToken, uint256 fromAmount) view returns(uint256 toAmount)
func (_WooPPV4 *WooPPV4Caller) Query(opts *bind.CallOpts, fromToken common.Address, toToken common.Address, fromAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WooPPV4.contract.Call(opts, &out, "query", fromToken, toToken, fromAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Query is a free data retrieval call binding the contract method 0xf58a435f.
//
// Solidity: function query(address fromToken, address toToken, uint256 fromAmount) view returns(uint256 toAmount)
func (_WooPPV4 *WooPPV4Session) Query(fromToken common.Address, toToken common.Address, fromAmount *big.Int) (*big.Int, error) {
	return _WooPPV4.Contract.Query(&_WooPPV4.CallOpts, fromToken, toToken, fromAmount)
}

// Query is a free data retrieval call binding the contract method 0xf58a435f.
//
// Solidity: function query(address fromToken, address toToken, uint256 fromAmount) view returns(uint256 toAmount)
func (_WooPPV4 *WooPPV4CallerSession) Query(fromToken common.Address, toToken common.Address, fromAmount *big.Int) (*big.Int, error) {
	return _WooPPV4.Contract.Query(&_WooPPV4.CallOpts, fromToken, toToken, fromAmount)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_WooPPV4 *WooPPV4Caller) QuoteToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WooPPV4.contract.Call(opts, &out, "quoteToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_WooPPV4 *WooPPV4Session) QuoteToken() (common.Address, error) {
	return _WooPPV4.Contract.QuoteToken(&_WooPPV4.CallOpts)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_WooPPV4 *WooPPV4CallerSession) QuoteToken() (common.Address, error) {
	return _WooPPV4.Contract.QuoteToken(&_WooPPV4.CallOpts)
}

// TokenInfos is a free data retrieval call binding the contract method 0xba46ae72.
//
// Solidity: function tokenInfos(address ) view returns(uint192 reserve, uint16 feeRate)
func (_WooPPV4 *WooPPV4Caller) TokenInfos(opts *bind.CallOpts, arg0 common.Address) (struct {
	Reserve *big.Int
	FeeRate uint16
}, error) {
	var out []interface{}
	err := _WooPPV4.contract.Call(opts, &out, "tokenInfos", arg0)

	outstruct := new(struct {
		Reserve *big.Int
		FeeRate uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FeeRate = *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return *outstruct, err

}

// TokenInfos is a free data retrieval call binding the contract method 0xba46ae72.
//
// Solidity: function tokenInfos(address ) view returns(uint192 reserve, uint16 feeRate)
func (_WooPPV4 *WooPPV4Session) TokenInfos(arg0 common.Address) (struct {
	Reserve *big.Int
	FeeRate uint16
}, error) {
	return _WooPPV4.Contract.TokenInfos(&_WooPPV4.CallOpts, arg0)
}

// TokenInfos is a free data retrieval call binding the contract method 0xba46ae72.
//
// Solidity: function tokenInfos(address ) view returns(uint192 reserve, uint16 feeRate)
func (_WooPPV4 *WooPPV4CallerSession) TokenInfos(arg0 common.Address) (struct {
	Reserve *big.Int
	FeeRate uint16
}, error) {
	return _WooPPV4.Contract.TokenInfos(&_WooPPV4.CallOpts, arg0)
}

// TryQuery is a free data retrieval call binding the contract method 0xce824f19.
//
// Solidity: function tryQuery(address fromToken, address toToken, uint256 fromAmount) view returns(uint256 toAmount)
func (_WooPPV4 *WooPPV4Caller) TryQuery(opts *bind.CallOpts, fromToken common.Address, toToken common.Address, fromAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WooPPV4.contract.Call(opts, &out, "tryQuery", fromToken, toToken, fromAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TryQuery is a free data retrieval call binding the contract method 0xce824f19.
//
// Solidity: function tryQuery(address fromToken, address toToken, uint256 fromAmount) view returns(uint256 toAmount)
func (_WooPPV4 *WooPPV4Session) TryQuery(fromToken common.Address, toToken common.Address, fromAmount *big.Int) (*big.Int, error) {
	return _WooPPV4.Contract.TryQuery(&_WooPPV4.CallOpts, fromToken, toToken, fromAmount)
}

// TryQuery is a free data retrieval call binding the contract method 0xce824f19.
//
// Solidity: function tryQuery(address fromToken, address toToken, uint256 fromAmount) view returns(uint256 toAmount)
func (_WooPPV4 *WooPPV4CallerSession) TryQuery(fromToken common.Address, toToken common.Address, fromAmount *big.Int) (*big.Int, error) {
	return _WooPPV4.Contract.TryQuery(&_WooPPV4.CallOpts, fromToken, toToken, fromAmount)
}

// UnclaimedFee is a free data retrieval call binding the contract method 0xe4d43ec1.
//
// Solidity: function unclaimedFee() view returns(uint256)
func (_WooPPV4 *WooPPV4Caller) UnclaimedFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WooPPV4.contract.Call(opts, &out, "unclaimedFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnclaimedFee is a free data retrieval call binding the contract method 0xe4d43ec1.
//
// Solidity: function unclaimedFee() view returns(uint256)
func (_WooPPV4 *WooPPV4Session) UnclaimedFee() (*big.Int, error) {
	return _WooPPV4.Contract.UnclaimedFee(&_WooPPV4.CallOpts)
}

// UnclaimedFee is a free data retrieval call binding the contract method 0xe4d43ec1.
//
// Solidity: function unclaimedFee() view returns(uint256)
func (_WooPPV4 *WooPPV4CallerSession) UnclaimedFee() (*big.Int, error) {
	return _WooPPV4.Contract.UnclaimedFee(&_WooPPV4.CallOpts)
}

// Wooracle is a free data retrieval call binding the contract method 0xbc8530f9.
//
// Solidity: function wooracle() view returns(address)
func (_WooPPV4 *WooPPV4Caller) Wooracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WooPPV4.contract.Call(opts, &out, "wooracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wooracle is a free data retrieval call binding the contract method 0xbc8530f9.
//
// Solidity: function wooracle() view returns(address)
func (_WooPPV4 *WooPPV4Session) Wooracle() (common.Address, error) {
	return _WooPPV4.Contract.Wooracle(&_WooPPV4.CallOpts)
}

// Wooracle is a free data retrieval call binding the contract method 0xbc8530f9.
//
// Solidity: function wooracle() view returns(address)
func (_WooPPV4 *WooPPV4CallerSession) Wooracle() (common.Address, error) {
	return _WooPPV4.Contract.Wooracle(&_WooPPV4.CallOpts)
}

// ClaimFee is a paid mutator transaction binding the contract method 0x99d32fc4.
//
// Solidity: function claimFee() returns()
func (_WooPPV4 *WooPPV4Transactor) ClaimFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WooPPV4.contract.Transact(opts, "claimFee")
}

// ClaimFee is a paid mutator transaction binding the contract method 0x99d32fc4.
//
// Solidity: function claimFee() returns()
func (_WooPPV4 *WooPPV4Session) ClaimFee() (*types.Transaction, error) {
	return _WooPPV4.Contract.ClaimFee(&_WooPPV4.TransactOpts)
}

// ClaimFee is a paid mutator transaction binding the contract method 0x99d32fc4.
//
// Solidity: function claimFee() returns()
func (_WooPPV4 *WooPPV4TransactorSession) ClaimFee() (*types.Transaction, error) {
	return _WooPPV4.Contract.ClaimFee(&_WooPPV4.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns()
func (_WooPPV4 *WooPPV4Transactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WooPPV4.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns()
func (_WooPPV4 *WooPPV4Session) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WooPPV4.Contract.Deposit(&_WooPPV4.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns()
func (_WooPPV4 *WooPPV4TransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WooPPV4.Contract.Deposit(&_WooPPV4.TransactOpts, token, amount)
}

// DepositAll is a paid mutator transaction binding the contract method 0x9f0d5f27.
//
// Solidity: function depositAll(address token) returns()
func (_WooPPV4 *WooPPV4Transactor) DepositAll(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _WooPPV4.contract.Transact(opts, "depositAll", token)
}

// DepositAll is a paid mutator transaction binding the contract method 0x9f0d5f27.
//
// Solidity: function depositAll(address token) returns()
func (_WooPPV4 *WooPPV4Session) DepositAll(token common.Address) (*types.Transaction, error) {
	return _WooPPV4.Contract.DepositAll(&_WooPPV4.TransactOpts, token)
}

// DepositAll is a paid mutator transaction binding the contract method 0x9f0d5f27.
//
// Solidity: function depositAll(address token) returns()
func (_WooPPV4 *WooPPV4TransactorSession) DepositAll(token common.Address) (*types.Transaction, error) {
	return _WooPPV4.Contract.DepositAll(&_WooPPV4.TransactOpts, token)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address _wooracle, address _feeAddr) returns()
func (_WooPPV4 *WooPPV4Transactor) Init(opts *bind.TransactOpts, _wooracle common.Address, _feeAddr common.Address) (*types.Transaction, error) {
	return _WooPPV4.contract.Transact(opts, "init", _wooracle, _feeAddr)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address _wooracle, address _feeAddr) returns()
func (_WooPPV4 *WooPPV4Session) Init(_wooracle common.Address, _feeAddr common.Address) (*types.Transaction, error) {
	return _WooPPV4.Contract.Init(&_WooPPV4.TransactOpts, _wooracle, _feeAddr)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address _wooracle, address _feeAddr) returns()
func (_WooPPV4 *WooPPV4TransactorSession) Init(_wooracle common.Address, _feeAddr common.Address) (*types.Transaction, error) {
	return _WooPPV4.Contract.Init(&_WooPPV4.TransactOpts, _wooracle, _feeAddr)
}

// MigrateToNewPool is a paid mutator transaction binding the contract method 0x0426d975.
//
// Solidity: function migrateToNewPool(address token, address newPool) returns()
func (_WooPPV4 *WooPPV4Transactor) MigrateToNewPool(opts *bind.TransactOpts, token common.Address, newPool common.Address) (*types.Transaction, error) {
	return _WooPPV4.contract.Transact(opts, "migrateToNewPool", token, newPool)
}

// MigrateToNewPool is a paid mutator transaction binding the contract method 0x0426d975.
//
// Solidity: function migrateToNewPool(address token, address newPool) returns()
func (_WooPPV4 *WooPPV4Session) MigrateToNewPool(token common.Address, newPool common.Address) (*types.Transaction, error) {
	return _WooPPV4.Contract.MigrateToNewPool(&_WooPPV4.TransactOpts, token, newPool)
}

// MigrateToNewPool is a paid mutator transaction binding the contract method 0x0426d975.
//
// Solidity: function migrateToNewPool(address token, address newPool) returns()
func (_WooPPV4 *WooPPV4TransactorSession) MigrateToNewPool(token common.Address, newPool common.Address) (*types.Transaction, error) {
	return _WooPPV4.Contract.MigrateToNewPool(&_WooPPV4.TransactOpts, token, newPool)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WooPPV4 *WooPPV4Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WooPPV4.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WooPPV4 *WooPPV4Session) Pause() (*types.Transaction, error) {
	return _WooPPV4.Contract.Pause(&_WooPPV4.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WooPPV4 *WooPPV4TransactorSession) Pause() (*types.Transaction, error) {
	return _WooPPV4.Contract.Pause(&_WooPPV4.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WooPPV4 *WooPPV4Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WooPPV4.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WooPPV4 *WooPPV4Session) RenounceOwnership() (*types.Transaction, error) {
	return _WooPPV4.Contract.RenounceOwnership(&_WooPPV4.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WooPPV4 *WooPPV4TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WooPPV4.Contract.RenounceOwnership(&_WooPPV4.TransactOpts)
}

// RepayWeeklyLending is a paid mutator transaction binding the contract method 0xebc80f13.
//
// Solidity: function repayWeeklyLending(address lendManager) returns()
func (_WooPPV4 *WooPPV4Transactor) RepayWeeklyLending(opts *bind.TransactOpts, lendManager common.Address) (*types.Transaction, error) {
	return _WooPPV4.contract.Transact(opts, "repayWeeklyLending", lendManager)
}

// RepayWeeklyLending is a paid mutator transaction binding the contract method 0xebc80f13.
//
// Solidity: function repayWeeklyLending(address lendManager) returns()
func (_WooPPV4 *WooPPV4Session) RepayWeeklyLending(lendManager common.Address) (*types.Transaction, error) {
	return _WooPPV4.Contract.RepayWeeklyLending(&_WooPPV4.TransactOpts, lendManager)
}

// RepayWeeklyLending is a paid mutator transaction binding the contract method 0xebc80f13.
//
// Solidity: function repayWeeklyLending(address lendManager) returns()
func (_WooPPV4 *WooPPV4TransactorSession) RepayWeeklyLending(lendManager common.Address) (*types.Transaction, error) {
	return _WooPPV4.Contract.RepayWeeklyLending(&_WooPPV4.TransactOpts, lendManager)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address addr, bool flag) returns()
func (_WooPPV4 *WooPPV4Transactor) SetAdmin(opts *bind.TransactOpts, addr common.Address, flag bool) (*types.Transaction, error) {
	return _WooPPV4.contract.Transact(opts, "setAdmin", addr, flag)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address addr, bool flag) returns()
func (_WooPPV4 *WooPPV4Session) SetAdmin(addr common.Address, flag bool) (*types.Transaction, error) {
	return _WooPPV4.Contract.SetAdmin(&_WooPPV4.TransactOpts, addr, flag)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address addr, bool flag) returns()
func (_WooPPV4 *WooPPV4TransactorSession) SetAdmin(addr common.Address, flag bool) (*types.Transaction, error) {
	return _WooPPV4.Contract.SetAdmin(&_WooPPV4.TransactOpts, addr, flag)
}

// SetFeeAddr is a paid mutator transaction binding the contract method 0xb2855b4f.
//
// Solidity: function setFeeAddr(address _feeAddr) returns()
func (_WooPPV4 *WooPPV4Transactor) SetFeeAddr(opts *bind.TransactOpts, _feeAddr common.Address) (*types.Transaction, error) {
	return _WooPPV4.contract.Transact(opts, "setFeeAddr", _feeAddr)
}

// SetFeeAddr is a paid mutator transaction binding the contract method 0xb2855b4f.
//
// Solidity: function setFeeAddr(address _feeAddr) returns()
func (_WooPPV4 *WooPPV4Session) SetFeeAddr(_feeAddr common.Address) (*types.Transaction, error) {
	return _WooPPV4.Contract.SetFeeAddr(&_WooPPV4.TransactOpts, _feeAddr)
}

// SetFeeAddr is a paid mutator transaction binding the contract method 0xb2855b4f.
//
// Solidity: function setFeeAddr(address _feeAddr) returns()
func (_WooPPV4 *WooPPV4TransactorSession) SetFeeAddr(_feeAddr common.Address) (*types.Transaction, error) {
	return _WooPPV4.Contract.SetFeeAddr(&_WooPPV4.TransactOpts, _feeAddr)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x68078eed.
//
// Solidity: function setFeeRate(address token, uint16 rate) returns()
func (_WooPPV4 *WooPPV4Transactor) SetFeeRate(opts *bind.TransactOpts, token common.Address, rate uint16) (*types.Transaction, error) {
	return _WooPPV4.contract.Transact(opts, "setFeeRate", token, rate)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x68078eed.
//
// Solidity: function setFeeRate(address token, uint16 rate) returns()
func (_WooPPV4 *WooPPV4Session) SetFeeRate(token common.Address, rate uint16) (*types.Transaction, error) {
	return _WooPPV4.Contract.SetFeeRate(&_WooPPV4.TransactOpts, token, rate)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x68078eed.
//
// Solidity: function setFeeRate(address token, uint16 rate) returns()
func (_WooPPV4 *WooPPV4TransactorSession) SetFeeRate(token common.Address, rate uint16) (*types.Transaction, error) {
	return _WooPPV4.Contract.SetFeeRate(&_WooPPV4.TransactOpts, token, rate)
}

// SetWooracle is a paid mutator transaction binding the contract method 0xa1ae8490.
//
// Solidity: function setWooracle(address _wooracle) returns()
func (_WooPPV4 *WooPPV4Transactor) SetWooracle(opts *bind.TransactOpts, _wooracle common.Address) (*types.Transaction, error) {
	return _WooPPV4.contract.Transact(opts, "setWooracle", _wooracle)
}

// SetWooracle is a paid mutator transaction binding the contract method 0xa1ae8490.
//
// Solidity: function setWooracle(address _wooracle) returns()
func (_WooPPV4 *WooPPV4Session) SetWooracle(_wooracle common.Address) (*types.Transaction, error) {
	return _WooPPV4.Contract.SetWooracle(&_WooPPV4.TransactOpts, _wooracle)
}

// SetWooracle is a paid mutator transaction binding the contract method 0xa1ae8490.
//
// Solidity: function setWooracle(address _wooracle) returns()
func (_WooPPV4 *WooPPV4TransactorSession) SetWooracle(_wooracle common.Address) (*types.Transaction, error) {
	return _WooPPV4.Contract.SetWooracle(&_WooPPV4.TransactOpts, _wooracle)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address token) returns()
func (_WooPPV4 *WooPPV4Transactor) Skim(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _WooPPV4.contract.Transact(opts, "skim", token)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address token) returns()
func (_WooPPV4 *WooPPV4Session) Skim(token common.Address) (*types.Transaction, error) {
	return _WooPPV4.Contract.Skim(&_WooPPV4.TransactOpts, token)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address token) returns()
func (_WooPPV4 *WooPPV4TransactorSession) Skim(token common.Address) (*types.Transaction, error) {
	return _WooPPV4.Contract.Skim(&_WooPPV4.TransactOpts, token)
}

// SkimMulTokens is a paid mutator transaction binding the contract method 0x7f0eec02.
//
// Solidity: function skimMulTokens(address[] tokens) returns()
func (_WooPPV4 *WooPPV4Transactor) SkimMulTokens(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error) {
	return _WooPPV4.contract.Transact(opts, "skimMulTokens", tokens)
}

// SkimMulTokens is a paid mutator transaction binding the contract method 0x7f0eec02.
//
// Solidity: function skimMulTokens(address[] tokens) returns()
func (_WooPPV4 *WooPPV4Session) SkimMulTokens(tokens []common.Address) (*types.Transaction, error) {
	return _WooPPV4.Contract.SkimMulTokens(&_WooPPV4.TransactOpts, tokens)
}

// SkimMulTokens is a paid mutator transaction binding the contract method 0x7f0eec02.
//
// Solidity: function skimMulTokens(address[] tokens) returns()
func (_WooPPV4 *WooPPV4TransactorSession) SkimMulTokens(tokens []common.Address) (*types.Transaction, error) {
	return _WooPPV4.Contract.SkimMulTokens(&_WooPPV4.TransactOpts, tokens)
}

// Swap is a paid mutator transaction binding the contract method 0x7dc20382.
//
// Solidity: function swap(address fromToken, address toToken, uint256 fromAmount, uint256 minToAmount, address to, address rebateTo) returns(uint256 realToAmount)
func (_WooPPV4 *WooPPV4Transactor) Swap(opts *bind.TransactOpts, fromToken common.Address, toToken common.Address, fromAmount *big.Int, minToAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _WooPPV4.contract.Transact(opts, "swap", fromToken, toToken, fromAmount, minToAmount, to, rebateTo)
}

// Swap is a paid mutator transaction binding the contract method 0x7dc20382.
//
// Solidity: function swap(address fromToken, address toToken, uint256 fromAmount, uint256 minToAmount, address to, address rebateTo) returns(uint256 realToAmount)
func (_WooPPV4 *WooPPV4Session) Swap(fromToken common.Address, toToken common.Address, fromAmount *big.Int, minToAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _WooPPV4.Contract.Swap(&_WooPPV4.TransactOpts, fromToken, toToken, fromAmount, minToAmount, to, rebateTo)
}

// Swap is a paid mutator transaction binding the contract method 0x7dc20382.
//
// Solidity: function swap(address fromToken, address toToken, uint256 fromAmount, uint256 minToAmount, address to, address rebateTo) returns(uint256 realToAmount)
func (_WooPPV4 *WooPPV4TransactorSession) Swap(fromToken common.Address, toToken common.Address, fromAmount *big.Int, minToAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _WooPPV4.Contract.Swap(&_WooPPV4.TransactOpts, fromToken, toToken, fromAmount, minToAmount, to, rebateTo)
}

// Sync is a paid mutator transaction binding the contract method 0xa5841194.
//
// Solidity: function sync(address token) returns()
func (_WooPPV4 *WooPPV4Transactor) Sync(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _WooPPV4.contract.Transact(opts, "sync", token)
}

// Sync is a paid mutator transaction binding the contract method 0xa5841194.
//
// Solidity: function sync(address token) returns()
func (_WooPPV4 *WooPPV4Session) Sync(token common.Address) (*types.Transaction, error) {
	return _WooPPV4.Contract.Sync(&_WooPPV4.TransactOpts, token)
}

// Sync is a paid mutator transaction binding the contract method 0xa5841194.
//
// Solidity: function sync(address token) returns()
func (_WooPPV4 *WooPPV4TransactorSession) Sync(token common.Address) (*types.Transaction, error) {
	return _WooPPV4.Contract.Sync(&_WooPPV4.TransactOpts, token)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WooPPV4 *WooPPV4Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WooPPV4.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WooPPV4 *WooPPV4Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WooPPV4.Contract.TransferOwnership(&_WooPPV4.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WooPPV4 *WooPPV4TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WooPPV4.Contract.TransferOwnership(&_WooPPV4.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WooPPV4 *WooPPV4Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WooPPV4.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WooPPV4 *WooPPV4Session) Unpause() (*types.Transaction, error) {
	return _WooPPV4.Contract.Unpause(&_WooPPV4.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WooPPV4 *WooPPV4TransactorSession) Unpause() (*types.Transaction, error) {
	return _WooPPV4.Contract.Unpause(&_WooPPV4.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address token, uint256 amount) returns()
func (_WooPPV4 *WooPPV4Transactor) Withdraw(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WooPPV4.contract.Transact(opts, "withdraw", token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address token, uint256 amount) returns()
func (_WooPPV4 *WooPPV4Session) Withdraw(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WooPPV4.Contract.Withdraw(&_WooPPV4.TransactOpts, token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address token, uint256 amount) returns()
func (_WooPPV4 *WooPPV4TransactorSession) Withdraw(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WooPPV4.Contract.Withdraw(&_WooPPV4.TransactOpts, token, amount)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0xfa09e630.
//
// Solidity: function withdrawAll(address token) returns()
func (_WooPPV4 *WooPPV4Transactor) WithdrawAll(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _WooPPV4.contract.Transact(opts, "withdrawAll", token)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0xfa09e630.
//
// Solidity: function withdrawAll(address token) returns()
func (_WooPPV4 *WooPPV4Session) WithdrawAll(token common.Address) (*types.Transaction, error) {
	return _WooPPV4.Contract.WithdrawAll(&_WooPPV4.TransactOpts, token)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0xfa09e630.
//
// Solidity: function withdrawAll(address token) returns()
func (_WooPPV4 *WooPPV4TransactorSession) WithdrawAll(token common.Address) (*types.Transaction, error) {
	return _WooPPV4.Contract.WithdrawAll(&_WooPPV4.TransactOpts, token)
}

// WooPPV4AdminUpdatedIterator is returned from FilterAdminUpdated and is used to iterate over the raw logs and unpacked data for AdminUpdated events raised by the WooPPV4 contract.
type WooPPV4AdminUpdatedIterator struct {
	Event *WooPPV4AdminUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WooPPV4AdminUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WooPPV4AdminUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WooPPV4AdminUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WooPPV4AdminUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WooPPV4AdminUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WooPPV4AdminUpdated represents a AdminUpdated event raised by the WooPPV4 contract.
type WooPPV4AdminUpdated struct {
	Addr common.Address
	Flag bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAdminUpdated is a free log retrieval operation binding the contract event 0x235bc17e7930760029e9f4d860a2a8089976de5b381cf8380fc11c1d88a11133.
//
// Solidity: event AdminUpdated(address indexed addr, bool flag)
func (_WooPPV4 *WooPPV4Filterer) FilterAdminUpdated(opts *bind.FilterOpts, addr []common.Address) (*WooPPV4AdminUpdatedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _WooPPV4.contract.FilterLogs(opts, "AdminUpdated", addrRule)
	if err != nil {
		return nil, err
	}
	return &WooPPV4AdminUpdatedIterator{contract: _WooPPV4.contract, event: "AdminUpdated", logs: logs, sub: sub}, nil
}

// WatchAdminUpdated is a free log subscription operation binding the contract event 0x235bc17e7930760029e9f4d860a2a8089976de5b381cf8380fc11c1d88a11133.
//
// Solidity: event AdminUpdated(address indexed addr, bool flag)
func (_WooPPV4 *WooPPV4Filterer) WatchAdminUpdated(opts *bind.WatchOpts, sink chan<- *WooPPV4AdminUpdated, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _WooPPV4.contract.WatchLogs(opts, "AdminUpdated", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WooPPV4AdminUpdated)
				if err := _WooPPV4.contract.UnpackLog(event, "AdminUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminUpdated is a log parse operation binding the contract event 0x235bc17e7930760029e9f4d860a2a8089976de5b381cf8380fc11c1d88a11133.
//
// Solidity: event AdminUpdated(address indexed addr, bool flag)
func (_WooPPV4 *WooPPV4Filterer) ParseAdminUpdated(log types.Log) (*WooPPV4AdminUpdated, error) {
	event := new(WooPPV4AdminUpdated)
	if err := _WooPPV4.contract.UnpackLog(event, "AdminUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WooPPV4DepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the WooPPV4 contract.
type WooPPV4DepositIterator struct {
	Event *WooPPV4Deposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WooPPV4DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WooPPV4Deposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WooPPV4Deposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WooPPV4DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WooPPV4DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WooPPV4Deposit represents a Deposit event raised by the WooPPV4 contract.
type WooPPV4Deposit struct {
	Token  common.Address
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed token, address indexed sender, uint256 amount)
func (_WooPPV4 *WooPPV4Filterer) FilterDeposit(opts *bind.FilterOpts, token []common.Address, sender []common.Address) (*WooPPV4DepositIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WooPPV4.contract.FilterLogs(opts, "Deposit", tokenRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &WooPPV4DepositIterator{contract: _WooPPV4.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed token, address indexed sender, uint256 amount)
func (_WooPPV4 *WooPPV4Filterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *WooPPV4Deposit, token []common.Address, sender []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WooPPV4.contract.WatchLogs(opts, "Deposit", tokenRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WooPPV4Deposit)
				if err := _WooPPV4.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed token, address indexed sender, uint256 amount)
func (_WooPPV4 *WooPPV4Filterer) ParseDeposit(log types.Log) (*WooPPV4Deposit, error) {
	event := new(WooPPV4Deposit)
	if err := _WooPPV4.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WooPPV4FeeAddrUpdatedIterator is returned from FilterFeeAddrUpdated and is used to iterate over the raw logs and unpacked data for FeeAddrUpdated events raised by the WooPPV4 contract.
type WooPPV4FeeAddrUpdatedIterator struct {
	Event *WooPPV4FeeAddrUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WooPPV4FeeAddrUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WooPPV4FeeAddrUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WooPPV4FeeAddrUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WooPPV4FeeAddrUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WooPPV4FeeAddrUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WooPPV4FeeAddrUpdated represents a FeeAddrUpdated event raised by the WooPPV4 contract.
type WooPPV4FeeAddrUpdated struct {
	NewFeeAddr common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeeAddrUpdated is a free log retrieval operation binding the contract event 0x76f97b3f5ffcc8d8d9577d141751a7bed446b46d837fbc81b5c01c165bbfbbf4.
//
// Solidity: event FeeAddrUpdated(address indexed newFeeAddr)
func (_WooPPV4 *WooPPV4Filterer) FilterFeeAddrUpdated(opts *bind.FilterOpts, newFeeAddr []common.Address) (*WooPPV4FeeAddrUpdatedIterator, error) {

	var newFeeAddrRule []interface{}
	for _, newFeeAddrItem := range newFeeAddr {
		newFeeAddrRule = append(newFeeAddrRule, newFeeAddrItem)
	}

	logs, sub, err := _WooPPV4.contract.FilterLogs(opts, "FeeAddrUpdated", newFeeAddrRule)
	if err != nil {
		return nil, err
	}
	return &WooPPV4FeeAddrUpdatedIterator{contract: _WooPPV4.contract, event: "FeeAddrUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeAddrUpdated is a free log subscription operation binding the contract event 0x76f97b3f5ffcc8d8d9577d141751a7bed446b46d837fbc81b5c01c165bbfbbf4.
//
// Solidity: event FeeAddrUpdated(address indexed newFeeAddr)
func (_WooPPV4 *WooPPV4Filterer) WatchFeeAddrUpdated(opts *bind.WatchOpts, sink chan<- *WooPPV4FeeAddrUpdated, newFeeAddr []common.Address) (event.Subscription, error) {

	var newFeeAddrRule []interface{}
	for _, newFeeAddrItem := range newFeeAddr {
		newFeeAddrRule = append(newFeeAddrRule, newFeeAddrItem)
	}

	logs, sub, err := _WooPPV4.contract.WatchLogs(opts, "FeeAddrUpdated", newFeeAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WooPPV4FeeAddrUpdated)
				if err := _WooPPV4.contract.UnpackLog(event, "FeeAddrUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeeAddrUpdated is a log parse operation binding the contract event 0x76f97b3f5ffcc8d8d9577d141751a7bed446b46d837fbc81b5c01c165bbfbbf4.
//
// Solidity: event FeeAddrUpdated(address indexed newFeeAddr)
func (_WooPPV4 *WooPPV4Filterer) ParseFeeAddrUpdated(log types.Log) (*WooPPV4FeeAddrUpdated, error) {
	event := new(WooPPV4FeeAddrUpdated)
	if err := _WooPPV4.contract.UnpackLog(event, "FeeAddrUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WooPPV4MigrateIterator is returned from FilterMigrate and is used to iterate over the raw logs and unpacked data for Migrate events raised by the WooPPV4 contract.
type WooPPV4MigrateIterator struct {
	Event *WooPPV4Migrate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WooPPV4MigrateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WooPPV4Migrate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WooPPV4Migrate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WooPPV4MigrateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WooPPV4MigrateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WooPPV4Migrate represents a Migrate event raised by the WooPPV4 contract.
type WooPPV4Migrate struct {
	Token    common.Address
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMigrate is a free log retrieval operation binding the contract event 0x18df02dcc52b9c494f391df09661519c0069bd8540141946280399408205ca1a.
//
// Solidity: event Migrate(address indexed token, address indexed receiver, uint256 amount)
func (_WooPPV4 *WooPPV4Filterer) FilterMigrate(opts *bind.FilterOpts, token []common.Address, receiver []common.Address) (*WooPPV4MigrateIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _WooPPV4.contract.FilterLogs(opts, "Migrate", tokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &WooPPV4MigrateIterator{contract: _WooPPV4.contract, event: "Migrate", logs: logs, sub: sub}, nil
}

// WatchMigrate is a free log subscription operation binding the contract event 0x18df02dcc52b9c494f391df09661519c0069bd8540141946280399408205ca1a.
//
// Solidity: event Migrate(address indexed token, address indexed receiver, uint256 amount)
func (_WooPPV4 *WooPPV4Filterer) WatchMigrate(opts *bind.WatchOpts, sink chan<- *WooPPV4Migrate, token []common.Address, receiver []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _WooPPV4.contract.WatchLogs(opts, "Migrate", tokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WooPPV4Migrate)
				if err := _WooPPV4.contract.UnpackLog(event, "Migrate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMigrate is a log parse operation binding the contract event 0x18df02dcc52b9c494f391df09661519c0069bd8540141946280399408205ca1a.
//
// Solidity: event Migrate(address indexed token, address indexed receiver, uint256 amount)
func (_WooPPV4 *WooPPV4Filterer) ParseMigrate(log types.Log) (*WooPPV4Migrate, error) {
	event := new(WooPPV4Migrate)
	if err := _WooPPV4.contract.UnpackLog(event, "Migrate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WooPPV4OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WooPPV4 contract.
type WooPPV4OwnershipTransferredIterator struct {
	Event *WooPPV4OwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WooPPV4OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WooPPV4OwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WooPPV4OwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WooPPV4OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WooPPV4OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WooPPV4OwnershipTransferred represents a OwnershipTransferred event raised by the WooPPV4 contract.
type WooPPV4OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WooPPV4 *WooPPV4Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WooPPV4OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WooPPV4.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WooPPV4OwnershipTransferredIterator{contract: _WooPPV4.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WooPPV4 *WooPPV4Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WooPPV4OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WooPPV4.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WooPPV4OwnershipTransferred)
				if err := _WooPPV4.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WooPPV4 *WooPPV4Filterer) ParseOwnershipTransferred(log types.Log) (*WooPPV4OwnershipTransferred, error) {
	event := new(WooPPV4OwnershipTransferred)
	if err := _WooPPV4.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WooPPV4PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the WooPPV4 contract.
type WooPPV4PausedIterator struct {
	Event *WooPPV4Paused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WooPPV4PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WooPPV4Paused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WooPPV4Paused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WooPPV4PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WooPPV4PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WooPPV4Paused represents a Paused event raised by the WooPPV4 contract.
type WooPPV4Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WooPPV4 *WooPPV4Filterer) FilterPaused(opts *bind.FilterOpts) (*WooPPV4PausedIterator, error) {

	logs, sub, err := _WooPPV4.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &WooPPV4PausedIterator{contract: _WooPPV4.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WooPPV4 *WooPPV4Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *WooPPV4Paused) (event.Subscription, error) {

	logs, sub, err := _WooPPV4.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WooPPV4Paused)
				if err := _WooPPV4.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WooPPV4 *WooPPV4Filterer) ParsePaused(log types.Log) (*WooPPV4Paused, error) {
	event := new(WooPPV4Paused)
	if err := _WooPPV4.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WooPPV4UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the WooPPV4 contract.
type WooPPV4UnpausedIterator struct {
	Event *WooPPV4Unpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WooPPV4UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WooPPV4Unpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WooPPV4Unpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WooPPV4UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WooPPV4UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WooPPV4Unpaused represents a Unpaused event raised by the WooPPV4 contract.
type WooPPV4Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WooPPV4 *WooPPV4Filterer) FilterUnpaused(opts *bind.FilterOpts) (*WooPPV4UnpausedIterator, error) {

	logs, sub, err := _WooPPV4.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &WooPPV4UnpausedIterator{contract: _WooPPV4.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WooPPV4 *WooPPV4Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *WooPPV4Unpaused) (event.Subscription, error) {

	logs, sub, err := _WooPPV4.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WooPPV4Unpaused)
				if err := _WooPPV4.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WooPPV4 *WooPPV4Filterer) ParseUnpaused(log types.Log) (*WooPPV4Unpaused, error) {
	event := new(WooPPV4Unpaused)
	if err := _WooPPV4.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WooPPV4WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the WooPPV4 contract.
type WooPPV4WithdrawIterator struct {
	Event *WooPPV4Withdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WooPPV4WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WooPPV4Withdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WooPPV4Withdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WooPPV4WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WooPPV4WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WooPPV4Withdraw represents a Withdraw event raised by the WooPPV4 contract.
type WooPPV4Withdraw struct {
	Token    common.Address
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed token, address indexed receiver, uint256 amount)
func (_WooPPV4 *WooPPV4Filterer) FilterWithdraw(opts *bind.FilterOpts, token []common.Address, receiver []common.Address) (*WooPPV4WithdrawIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _WooPPV4.contract.FilterLogs(opts, "Withdraw", tokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &WooPPV4WithdrawIterator{contract: _WooPPV4.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed token, address indexed receiver, uint256 amount)
func (_WooPPV4 *WooPPV4Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *WooPPV4Withdraw, token []common.Address, receiver []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _WooPPV4.contract.WatchLogs(opts, "Withdraw", tokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WooPPV4Withdraw)
				if err := _WooPPV4.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed token, address indexed receiver, uint256 amount)
func (_WooPPV4 *WooPPV4Filterer) ParseWithdraw(log types.Log) (*WooPPV4Withdraw, error) {
	event := new(WooPPV4Withdraw)
	if err := _WooPPV4.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WooPPV4WooSwapIterator is returned from FilterWooSwap and is used to iterate over the raw logs and unpacked data for WooSwap events raised by the WooPPV4 contract.
type WooPPV4WooSwapIterator struct {
	Event *WooPPV4WooSwap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WooPPV4WooSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WooPPV4WooSwap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WooPPV4WooSwap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WooPPV4WooSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WooPPV4WooSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WooPPV4WooSwap represents a WooSwap event raised by the WooPPV4 contract.
type WooPPV4WooSwap struct {
	FromToken  common.Address
	ToToken    common.Address
	FromAmount *big.Int
	ToAmount   *big.Int
	From       common.Address
	To         common.Address
	RebateTo   common.Address
	SwapVol    *big.Int
	SwapFee    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWooSwap is a free log retrieval operation binding the contract event 0x0e8e403c2d36126272b08c75823e988381d9dc47f2f0a9a080d95f891d95c469.
//
// Solidity: event WooSwap(address indexed fromToken, address indexed toToken, uint256 fromAmount, uint256 toAmount, address from, address indexed to, address rebateTo, uint256 swapVol, uint256 swapFee)
func (_WooPPV4 *WooPPV4Filterer) FilterWooSwap(opts *bind.FilterOpts, fromToken []common.Address, toToken []common.Address, to []common.Address) (*WooPPV4WooSwapIterator, error) {

	var fromTokenRule []interface{}
	for _, fromTokenItem := range fromToken {
		fromTokenRule = append(fromTokenRule, fromTokenItem)
	}
	var toTokenRule []interface{}
	for _, toTokenItem := range toToken {
		toTokenRule = append(toTokenRule, toTokenItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WooPPV4.contract.FilterLogs(opts, "WooSwap", fromTokenRule, toTokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WooPPV4WooSwapIterator{contract: _WooPPV4.contract, event: "WooSwap", logs: logs, sub: sub}, nil
}

// WatchWooSwap is a free log subscription operation binding the contract event 0x0e8e403c2d36126272b08c75823e988381d9dc47f2f0a9a080d95f891d95c469.
//
// Solidity: event WooSwap(address indexed fromToken, address indexed toToken, uint256 fromAmount, uint256 toAmount, address from, address indexed to, address rebateTo, uint256 swapVol, uint256 swapFee)
func (_WooPPV4 *WooPPV4Filterer) WatchWooSwap(opts *bind.WatchOpts, sink chan<- *WooPPV4WooSwap, fromToken []common.Address, toToken []common.Address, to []common.Address) (event.Subscription, error) {

	var fromTokenRule []interface{}
	for _, fromTokenItem := range fromToken {
		fromTokenRule = append(fromTokenRule, fromTokenItem)
	}
	var toTokenRule []interface{}
	for _, toTokenItem := range toToken {
		toTokenRule = append(toTokenRule, toTokenItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WooPPV4.contract.WatchLogs(opts, "WooSwap", fromTokenRule, toTokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WooPPV4WooSwap)
				if err := _WooPPV4.contract.UnpackLog(event, "WooSwap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWooSwap is a log parse operation binding the contract event 0x0e8e403c2d36126272b08c75823e988381d9dc47f2f0a9a080d95f891d95c469.
//
// Solidity: event WooSwap(address indexed fromToken, address indexed toToken, uint256 fromAmount, uint256 toAmount, address from, address indexed to, address rebateTo, uint256 swapVol, uint256 swapFee)
func (_WooPPV4 *WooPPV4Filterer) ParseWooSwap(log types.Log) (*WooPPV4WooSwap, error) {
	event := new(WooPPV4WooSwap)
	if err := _WooPPV4.contract.UnpackLog(event, "WooSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WooPPV4WooracleUpdatedIterator is returned from FilterWooracleUpdated and is used to iterate over the raw logs and unpacked data for WooracleUpdated events raised by the WooPPV4 contract.
type WooPPV4WooracleUpdatedIterator struct {
	Event *WooPPV4WooracleUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WooPPV4WooracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WooPPV4WooracleUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WooPPV4WooracleUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WooPPV4WooracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WooPPV4WooracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WooPPV4WooracleUpdated represents a WooracleUpdated event raised by the WooPPV4 contract.
type WooPPV4WooracleUpdated struct {
	NewWooracle common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWooracleUpdated is a free log retrieval operation binding the contract event 0x59ddfdd1bf7e3ea08a07e8aaa1fe2ce9c840fab69fe5ede6ea727a45eb42fc66.
//
// Solidity: event WooracleUpdated(address indexed newWooracle)
func (_WooPPV4 *WooPPV4Filterer) FilterWooracleUpdated(opts *bind.FilterOpts, newWooracle []common.Address) (*WooPPV4WooracleUpdatedIterator, error) {

	var newWooracleRule []interface{}
	for _, newWooracleItem := range newWooracle {
		newWooracleRule = append(newWooracleRule, newWooracleItem)
	}

	logs, sub, err := _WooPPV4.contract.FilterLogs(opts, "WooracleUpdated", newWooracleRule)
	if err != nil {
		return nil, err
	}
	return &WooPPV4WooracleUpdatedIterator{contract: _WooPPV4.contract, event: "WooracleUpdated", logs: logs, sub: sub}, nil
}

// WatchWooracleUpdated is a free log subscription operation binding the contract event 0x59ddfdd1bf7e3ea08a07e8aaa1fe2ce9c840fab69fe5ede6ea727a45eb42fc66.
//
// Solidity: event WooracleUpdated(address indexed newWooracle)
func (_WooPPV4 *WooPPV4Filterer) WatchWooracleUpdated(opts *bind.WatchOpts, sink chan<- *WooPPV4WooracleUpdated, newWooracle []common.Address) (event.Subscription, error) {

	var newWooracleRule []interface{}
	for _, newWooracleItem := range newWooracle {
		newWooracleRule = append(newWooracleRule, newWooracleItem)
	}

	logs, sub, err := _WooPPV4.contract.WatchLogs(opts, "WooracleUpdated", newWooracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WooPPV4WooracleUpdated)
				if err := _WooPPV4.contract.UnpackLog(event, "WooracleUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWooracleUpdated is a log parse operation binding the contract event 0x59ddfdd1bf7e3ea08a07e8aaa1fe2ce9c840fab69fe5ede6ea727a45eb42fc66.
//
// Solidity: event WooracleUpdated(address indexed newWooracle)
func (_WooPPV4 *WooPPV4Filterer) ParseWooracleUpdated(log types.Log) (*WooPPV4WooracleUpdated, error) {
	event := new(WooPPV4WooracleUpdated)
	if err := _WooPPV4.contract.UnpackLog(event, "WooracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
