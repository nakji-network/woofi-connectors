// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package WooracleV2

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

// IWooracleV2State is an auto generated low-level Go binding around an user-defined struct.
type IWooracleV2State struct {
	Price      *big.Int
	Spread     uint64
	Coeff      uint64
	WoFeasible bool
}

// WooracleV2MetaData contains all meta data concerning the WooracleV2 contract.
var WooracleV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bound\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"clOracles\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"decimal\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"cloPreferred\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"}],\"name\":\"cloAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"clo\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"}],\"name\":\"cloPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"}],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"infos\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"coeff\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"spread\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"}],\"name\":\"isWoFeasible\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"newPrice\",\"type\":\"uint128\"}],\"name\":\"postPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"bases\",\"type\":\"address[]\"},{\"internalType\":\"uint128[]\",\"name\":\"newPrices\",\"type\":\"uint128[]\"}],\"name\":\"postPriceList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"newSpread\",\"type\":\"uint64\"}],\"name\":\"postSpread\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"bases\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"newSpreads\",\"type\":\"uint64[]\"}],\"name\":\"postSpreadList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"newPrice\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"newSpread\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"newCoeff\",\"type\":\"uint64\"}],\"name\":\"postState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"bases\",\"type\":\"address[]\"},{\"internalType\":\"uint128[]\",\"name\":\"newPrices\",\"type\":\"uint128[]\"},{\"internalType\":\"uint64[]\",\"name\":\"newSpreads\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64[]\",\"name\":\"newCoeffs\",\"type\":\"uint64[]\"}],\"name\":\"postStateList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"}],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"priceOut\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"feasible\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_bound\",\"type\":\"uint64\"}],\"name\":\"setBound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_cloPreferred\",\"type\":\"bool\"}],\"name\":\"setCLOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_cloPreferred\",\"type\":\"bool\"}],\"name\":\"setCloPreferred\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_quote\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"setQuoteToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newStaleDuration\",\"type\":\"uint256\"}],\"name\":\"setStaleDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staleDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"}],\"name\":\"state\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"spread\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"coeff\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"woFeasible\",\"type\":\"bool\"}],\"internalType\":\"structIWooracleV2.State\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"}],\"name\":\"woCoeff\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"}],\"name\":\"woPrice\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"priceOut\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"priceTimestampOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"}],\"name\":\"woSpread\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"}],\"name\":\"woState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"spread\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"coeff\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"woFeasible\",\"type\":\"bool\"}],\"internalType\":\"structIWooracleV2.State\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// WooracleV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use WooracleV2MetaData.ABI instead.
var WooracleV2ABI = WooracleV2MetaData.ABI

// WooracleV2 is an auto generated Go binding around an Ethereum contract.
type WooracleV2 struct {
	WooracleV2Caller     // Read-only binding to the contract
	WooracleV2Transactor // Write-only binding to the contract
	WooracleV2Filterer   // Log filterer for contract events
}

// WooracleV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type WooracleV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WooracleV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type WooracleV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WooracleV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WooracleV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WooracleV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WooracleV2Session struct {
	Contract     *WooracleV2       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WooracleV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WooracleV2CallerSession struct {
	Contract *WooracleV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// WooracleV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WooracleV2TransactorSession struct {
	Contract     *WooracleV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// WooracleV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type WooracleV2Raw struct {
	Contract *WooracleV2 // Generic contract binding to access the raw methods on
}

// WooracleV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WooracleV2CallerRaw struct {
	Contract *WooracleV2Caller // Generic read-only contract binding to access the raw methods on
}

// WooracleV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WooracleV2TransactorRaw struct {
	Contract *WooracleV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewWooracleV2 creates a new instance of WooracleV2, bound to a specific deployed contract.
func NewWooracleV2(address common.Address, backend bind.ContractBackend) (*WooracleV2, error) {
	contract, err := bindWooracleV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WooracleV2{WooracleV2Caller: WooracleV2Caller{contract: contract}, WooracleV2Transactor: WooracleV2Transactor{contract: contract}, WooracleV2Filterer: WooracleV2Filterer{contract: contract}}, nil
}

// NewWooracleV2Caller creates a new read-only instance of WooracleV2, bound to a specific deployed contract.
func NewWooracleV2Caller(address common.Address, caller bind.ContractCaller) (*WooracleV2Caller, error) {
	contract, err := bindWooracleV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WooracleV2Caller{contract: contract}, nil
}

// NewWooracleV2Transactor creates a new write-only instance of WooracleV2, bound to a specific deployed contract.
func NewWooracleV2Transactor(address common.Address, transactor bind.ContractTransactor) (*WooracleV2Transactor, error) {
	contract, err := bindWooracleV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WooracleV2Transactor{contract: contract}, nil
}

// NewWooracleV2Filterer creates a new log filterer instance of WooracleV2, bound to a specific deployed contract.
func NewWooracleV2Filterer(address common.Address, filterer bind.ContractFilterer) (*WooracleV2Filterer, error) {
	contract, err := bindWooracleV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WooracleV2Filterer{contract: contract}, nil
}

// bindWooracleV2 binds a generic wrapper to an already deployed contract.
func bindWooracleV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WooracleV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WooracleV2 *WooracleV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WooracleV2.Contract.WooracleV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WooracleV2 *WooracleV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WooracleV2.Contract.WooracleV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WooracleV2 *WooracleV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WooracleV2.Contract.WooracleV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WooracleV2 *WooracleV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WooracleV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WooracleV2 *WooracleV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WooracleV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WooracleV2 *WooracleV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WooracleV2.Contract.contract.Transact(opts, method, params...)
}

// Bound is a free data retrieval call binding the contract method 0xbe4df7d6.
//
// Solidity: function bound() view returns(uint64)
func (_WooracleV2 *WooracleV2Caller) Bound(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _WooracleV2.contract.Call(opts, &out, "bound")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Bound is a free data retrieval call binding the contract method 0xbe4df7d6.
//
// Solidity: function bound() view returns(uint64)
func (_WooracleV2 *WooracleV2Session) Bound() (uint64, error) {
	return _WooracleV2.Contract.Bound(&_WooracleV2.CallOpts)
}

// Bound is a free data retrieval call binding the contract method 0xbe4df7d6.
//
// Solidity: function bound() view returns(uint64)
func (_WooracleV2 *WooracleV2CallerSession) Bound() (uint64, error) {
	return _WooracleV2.Contract.Bound(&_WooracleV2.CallOpts)
}

// ClOracles is a free data retrieval call binding the contract method 0xf2030e73.
//
// Solidity: function clOracles(address ) view returns(address oracle, uint8 decimal, bool cloPreferred)
func (_WooracleV2 *WooracleV2Caller) ClOracles(opts *bind.CallOpts, arg0 common.Address) (struct {
	Oracle       common.Address
	Decimal      uint8
	CloPreferred bool
}, error) {
	var out []interface{}
	err := _WooracleV2.contract.Call(opts, &out, "clOracles", arg0)

	outstruct := new(struct {
		Oracle       common.Address
		Decimal      uint8
		CloPreferred bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Oracle = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Decimal = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.CloPreferred = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// ClOracles is a free data retrieval call binding the contract method 0xf2030e73.
//
// Solidity: function clOracles(address ) view returns(address oracle, uint8 decimal, bool cloPreferred)
func (_WooracleV2 *WooracleV2Session) ClOracles(arg0 common.Address) (struct {
	Oracle       common.Address
	Decimal      uint8
	CloPreferred bool
}, error) {
	return _WooracleV2.Contract.ClOracles(&_WooracleV2.CallOpts, arg0)
}

// ClOracles is a free data retrieval call binding the contract method 0xf2030e73.
//
// Solidity: function clOracles(address ) view returns(address oracle, uint8 decimal, bool cloPreferred)
func (_WooracleV2 *WooracleV2CallerSession) ClOracles(arg0 common.Address) (struct {
	Oracle       common.Address
	Decimal      uint8
	CloPreferred bool
}, error) {
	return _WooracleV2.Contract.ClOracles(&_WooracleV2.CallOpts, arg0)
}

// CloAddress is a free data retrieval call binding the contract method 0x96bb520f.
//
// Solidity: function cloAddress(address base) view returns(address clo)
func (_WooracleV2 *WooracleV2Caller) CloAddress(opts *bind.CallOpts, base common.Address) (common.Address, error) {
	var out []interface{}
	err := _WooracleV2.contract.Call(opts, &out, "cloAddress", base)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CloAddress is a free data retrieval call binding the contract method 0x96bb520f.
//
// Solidity: function cloAddress(address base) view returns(address clo)
func (_WooracleV2 *WooracleV2Session) CloAddress(base common.Address) (common.Address, error) {
	return _WooracleV2.Contract.CloAddress(&_WooracleV2.CallOpts, base)
}

// CloAddress is a free data retrieval call binding the contract method 0x96bb520f.
//
// Solidity: function cloAddress(address base) view returns(address clo)
func (_WooracleV2 *WooracleV2CallerSession) CloAddress(base common.Address) (common.Address, error) {
	return _WooracleV2.Contract.CloAddress(&_WooracleV2.CallOpts, base)
}

// CloPrice is a free data retrieval call binding the contract method 0x0b7841f5.
//
// Solidity: function cloPrice(address base) view returns(uint256, uint256)
func (_WooracleV2 *WooracleV2Caller) CloPrice(opts *bind.CallOpts, base common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _WooracleV2.contract.Call(opts, &out, "cloPrice", base)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// CloPrice is a free data retrieval call binding the contract method 0x0b7841f5.
//
// Solidity: function cloPrice(address base) view returns(uint256, uint256)
func (_WooracleV2 *WooracleV2Session) CloPrice(base common.Address) (*big.Int, *big.Int, error) {
	return _WooracleV2.Contract.CloPrice(&_WooracleV2.CallOpts, base)
}

// CloPrice is a free data retrieval call binding the contract method 0x0b7841f5.
//
// Solidity: function cloPrice(address base) view returns(uint256, uint256)
func (_WooracleV2 *WooracleV2CallerSession) CloPrice(base common.Address) (*big.Int, *big.Int, error) {
	return _WooracleV2.Contract.CloPrice(&_WooracleV2.CallOpts, base)
}

// Decimals is a free data retrieval call binding the contract method 0xd449a832.
//
// Solidity: function decimals(address base) view returns(uint8)
func (_WooracleV2 *WooracleV2Caller) Decimals(opts *bind.CallOpts, base common.Address) (uint8, error) {
	var out []interface{}
	err := _WooracleV2.contract.Call(opts, &out, "decimals", base)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0xd449a832.
//
// Solidity: function decimals(address base) view returns(uint8)
func (_WooracleV2 *WooracleV2Session) Decimals(base common.Address) (uint8, error) {
	return _WooracleV2.Contract.Decimals(&_WooracleV2.CallOpts, base)
}

// Decimals is a free data retrieval call binding the contract method 0xd449a832.
//
// Solidity: function decimals(address base) view returns(uint8)
func (_WooracleV2 *WooracleV2CallerSession) Decimals(base common.Address) (uint8, error) {
	return _WooracleV2.Contract.Decimals(&_WooracleV2.CallOpts, base)
}

// Infos is a free data retrieval call binding the contract method 0xc6ddb642.
//
// Solidity: function infos(address ) view returns(uint128 price, uint64 coeff, uint64 spread)
func (_WooracleV2 *WooracleV2Caller) Infos(opts *bind.CallOpts, arg0 common.Address) (struct {
	Price  *big.Int
	Coeff  uint64
	Spread uint64
}, error) {
	var out []interface{}
	err := _WooracleV2.contract.Call(opts, &out, "infos", arg0)

	outstruct := new(struct {
		Price  *big.Int
		Coeff  uint64
		Spread uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Coeff = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Spread = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// Infos is a free data retrieval call binding the contract method 0xc6ddb642.
//
// Solidity: function infos(address ) view returns(uint128 price, uint64 coeff, uint64 spread)
func (_WooracleV2 *WooracleV2Session) Infos(arg0 common.Address) (struct {
	Price  *big.Int
	Coeff  uint64
	Spread uint64
}, error) {
	return _WooracleV2.Contract.Infos(&_WooracleV2.CallOpts, arg0)
}

// Infos is a free data retrieval call binding the contract method 0xc6ddb642.
//
// Solidity: function infos(address ) view returns(uint128 price, uint64 coeff, uint64 spread)
func (_WooracleV2 *WooracleV2CallerSession) Infos(arg0 common.Address) (struct {
	Price  *big.Int
	Coeff  uint64
	Spread uint64
}, error) {
	return _WooracleV2.Contract.Infos(&_WooracleV2.CallOpts, arg0)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_WooracleV2 *WooracleV2Caller) IsAdmin(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _WooracleV2.contract.Call(opts, &out, "isAdmin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_WooracleV2 *WooracleV2Session) IsAdmin(arg0 common.Address) (bool, error) {
	return _WooracleV2.Contract.IsAdmin(&_WooracleV2.CallOpts, arg0)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_WooracleV2 *WooracleV2CallerSession) IsAdmin(arg0 common.Address) (bool, error) {
	return _WooracleV2.Contract.IsAdmin(&_WooracleV2.CallOpts, arg0)
}

// IsWoFeasible is a free data retrieval call binding the contract method 0x37e257fd.
//
// Solidity: function isWoFeasible(address base) view returns(bool)
func (_WooracleV2 *WooracleV2Caller) IsWoFeasible(opts *bind.CallOpts, base common.Address) (bool, error) {
	var out []interface{}
	err := _WooracleV2.contract.Call(opts, &out, "isWoFeasible", base)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWoFeasible is a free data retrieval call binding the contract method 0x37e257fd.
//
// Solidity: function isWoFeasible(address base) view returns(bool)
func (_WooracleV2 *WooracleV2Session) IsWoFeasible(base common.Address) (bool, error) {
	return _WooracleV2.Contract.IsWoFeasible(&_WooracleV2.CallOpts, base)
}

// IsWoFeasible is a free data retrieval call binding the contract method 0x37e257fd.
//
// Solidity: function isWoFeasible(address base) view returns(bool)
func (_WooracleV2 *WooracleV2CallerSession) IsWoFeasible(base common.Address) (bool, error) {
	return _WooracleV2.Contract.IsWoFeasible(&_WooracleV2.CallOpts, base)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WooracleV2 *WooracleV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WooracleV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WooracleV2 *WooracleV2Session) Owner() (common.Address, error) {
	return _WooracleV2.Contract.Owner(&_WooracleV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WooracleV2 *WooracleV2CallerSession) Owner() (common.Address, error) {
	return _WooracleV2.Contract.Owner(&_WooracleV2.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xaea91078.
//
// Solidity: function price(address base) view returns(uint256 priceOut, bool feasible)
func (_WooracleV2 *WooracleV2Caller) Price(opts *bind.CallOpts, base common.Address) (struct {
	PriceOut *big.Int
	Feasible bool
}, error) {
	var out []interface{}
	err := _WooracleV2.contract.Call(opts, &out, "price", base)

	outstruct := new(struct {
		PriceOut *big.Int
		Feasible bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PriceOut = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Feasible = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// Price is a free data retrieval call binding the contract method 0xaea91078.
//
// Solidity: function price(address base) view returns(uint256 priceOut, bool feasible)
func (_WooracleV2 *WooracleV2Session) Price(base common.Address) (struct {
	PriceOut *big.Int
	Feasible bool
}, error) {
	return _WooracleV2.Contract.Price(&_WooracleV2.CallOpts, base)
}

// Price is a free data retrieval call binding the contract method 0xaea91078.
//
// Solidity: function price(address base) view returns(uint256 priceOut, bool feasible)
func (_WooracleV2 *WooracleV2CallerSession) Price(base common.Address) (struct {
	PriceOut *big.Int
	Feasible bool
}, error) {
	return _WooracleV2.Contract.Price(&_WooracleV2.CallOpts, base)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_WooracleV2 *WooracleV2Caller) QuoteToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WooracleV2.contract.Call(opts, &out, "quoteToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_WooracleV2 *WooracleV2Session) QuoteToken() (common.Address, error) {
	return _WooracleV2.Contract.QuoteToken(&_WooracleV2.CallOpts)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_WooracleV2 *WooracleV2CallerSession) QuoteToken() (common.Address, error) {
	return _WooracleV2.Contract.QuoteToken(&_WooracleV2.CallOpts)
}

// StaleDuration is a free data retrieval call binding the contract method 0xcc6864b1.
//
// Solidity: function staleDuration() view returns(uint256)
func (_WooracleV2 *WooracleV2Caller) StaleDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WooracleV2.contract.Call(opts, &out, "staleDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StaleDuration is a free data retrieval call binding the contract method 0xcc6864b1.
//
// Solidity: function staleDuration() view returns(uint256)
func (_WooracleV2 *WooracleV2Session) StaleDuration() (*big.Int, error) {
	return _WooracleV2.Contract.StaleDuration(&_WooracleV2.CallOpts)
}

// StaleDuration is a free data retrieval call binding the contract method 0xcc6864b1.
//
// Solidity: function staleDuration() view returns(uint256)
func (_WooracleV2 *WooracleV2CallerSession) StaleDuration() (*big.Int, error) {
	return _WooracleV2.Contract.StaleDuration(&_WooracleV2.CallOpts)
}

// State is a free data retrieval call binding the contract method 0x31e658a5.
//
// Solidity: function state(address base) view returns((uint128,uint64,uint64,bool))
func (_WooracleV2 *WooracleV2Caller) State(opts *bind.CallOpts, base common.Address) (IWooracleV2State, error) {
	var out []interface{}
	err := _WooracleV2.contract.Call(opts, &out, "state", base)

	if err != nil {
		return *new(IWooracleV2State), err
	}

	out0 := *abi.ConvertType(out[0], new(IWooracleV2State)).(*IWooracleV2State)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0x31e658a5.
//
// Solidity: function state(address base) view returns((uint128,uint64,uint64,bool))
func (_WooracleV2 *WooracleV2Session) State(base common.Address) (IWooracleV2State, error) {
	return _WooracleV2.Contract.State(&_WooracleV2.CallOpts, base)
}

// State is a free data retrieval call binding the contract method 0x31e658a5.
//
// Solidity: function state(address base) view returns((uint128,uint64,uint64,bool))
func (_WooracleV2 *WooracleV2CallerSession) State(base common.Address) (IWooracleV2State, error) {
	return _WooracleV2.Contract.State(&_WooracleV2.CallOpts, base)
}

// Timestamp is a free data retrieval call binding the contract method 0xb80777ea.
//
// Solidity: function timestamp() view returns(uint256)
func (_WooracleV2 *WooracleV2Caller) Timestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WooracleV2.contract.Call(opts, &out, "timestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Timestamp is a free data retrieval call binding the contract method 0xb80777ea.
//
// Solidity: function timestamp() view returns(uint256)
func (_WooracleV2 *WooracleV2Session) Timestamp() (*big.Int, error) {
	return _WooracleV2.Contract.Timestamp(&_WooracleV2.CallOpts)
}

// Timestamp is a free data retrieval call binding the contract method 0xb80777ea.
//
// Solidity: function timestamp() view returns(uint256)
func (_WooracleV2 *WooracleV2CallerSession) Timestamp() (*big.Int, error) {
	return _WooracleV2.Contract.Timestamp(&_WooracleV2.CallOpts)
}

// WoCoeff is a free data retrieval call binding the contract method 0xc3c0993b.
//
// Solidity: function woCoeff(address base) view returns(uint64)
func (_WooracleV2 *WooracleV2Caller) WoCoeff(opts *bind.CallOpts, base common.Address) (uint64, error) {
	var out []interface{}
	err := _WooracleV2.contract.Call(opts, &out, "woCoeff", base)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// WoCoeff is a free data retrieval call binding the contract method 0xc3c0993b.
//
// Solidity: function woCoeff(address base) view returns(uint64)
func (_WooracleV2 *WooracleV2Session) WoCoeff(base common.Address) (uint64, error) {
	return _WooracleV2.Contract.WoCoeff(&_WooracleV2.CallOpts, base)
}

// WoCoeff is a free data retrieval call binding the contract method 0xc3c0993b.
//
// Solidity: function woCoeff(address base) view returns(uint64)
func (_WooracleV2 *WooracleV2CallerSession) WoCoeff(base common.Address) (uint64, error) {
	return _WooracleV2.Contract.WoCoeff(&_WooracleV2.CallOpts, base)
}

// WoPrice is a free data retrieval call binding the contract method 0xa2c5d01a.
//
// Solidity: function woPrice(address base) view returns(uint128 priceOut, uint256 priceTimestampOut)
func (_WooracleV2 *WooracleV2Caller) WoPrice(opts *bind.CallOpts, base common.Address) (struct {
	PriceOut          *big.Int
	PriceTimestampOut *big.Int
}, error) {
	var out []interface{}
	err := _WooracleV2.contract.Call(opts, &out, "woPrice", base)

	outstruct := new(struct {
		PriceOut          *big.Int
		PriceTimestampOut *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PriceOut = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PriceTimestampOut = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// WoPrice is a free data retrieval call binding the contract method 0xa2c5d01a.
//
// Solidity: function woPrice(address base) view returns(uint128 priceOut, uint256 priceTimestampOut)
func (_WooracleV2 *WooracleV2Session) WoPrice(base common.Address) (struct {
	PriceOut          *big.Int
	PriceTimestampOut *big.Int
}, error) {
	return _WooracleV2.Contract.WoPrice(&_WooracleV2.CallOpts, base)
}

// WoPrice is a free data retrieval call binding the contract method 0xa2c5d01a.
//
// Solidity: function woPrice(address base) view returns(uint128 priceOut, uint256 priceTimestampOut)
func (_WooracleV2 *WooracleV2CallerSession) WoPrice(base common.Address) (struct {
	PriceOut          *big.Int
	PriceTimestampOut *big.Int
}, error) {
	return _WooracleV2.Contract.WoPrice(&_WooracleV2.CallOpts, base)
}

// WoSpread is a free data retrieval call binding the contract method 0xd09a568d.
//
// Solidity: function woSpread(address base) view returns(uint64)
func (_WooracleV2 *WooracleV2Caller) WoSpread(opts *bind.CallOpts, base common.Address) (uint64, error) {
	var out []interface{}
	err := _WooracleV2.contract.Call(opts, &out, "woSpread", base)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// WoSpread is a free data retrieval call binding the contract method 0xd09a568d.
//
// Solidity: function woSpread(address base) view returns(uint64)
func (_WooracleV2 *WooracleV2Session) WoSpread(base common.Address) (uint64, error) {
	return _WooracleV2.Contract.WoSpread(&_WooracleV2.CallOpts, base)
}

// WoSpread is a free data retrieval call binding the contract method 0xd09a568d.
//
// Solidity: function woSpread(address base) view returns(uint64)
func (_WooracleV2 *WooracleV2CallerSession) WoSpread(base common.Address) (uint64, error) {
	return _WooracleV2.Contract.WoSpread(&_WooracleV2.CallOpts, base)
}

// WoState is a free data retrieval call binding the contract method 0xc16116d4.
//
// Solidity: function woState(address base) view returns((uint128,uint64,uint64,bool))
func (_WooracleV2 *WooracleV2Caller) WoState(opts *bind.CallOpts, base common.Address) (IWooracleV2State, error) {
	var out []interface{}
	err := _WooracleV2.contract.Call(opts, &out, "woState", base)

	if err != nil {
		return *new(IWooracleV2State), err
	}

	out0 := *abi.ConvertType(out[0], new(IWooracleV2State)).(*IWooracleV2State)

	return out0, err

}

// WoState is a free data retrieval call binding the contract method 0xc16116d4.
//
// Solidity: function woState(address base) view returns((uint128,uint64,uint64,bool))
func (_WooracleV2 *WooracleV2Session) WoState(base common.Address) (IWooracleV2State, error) {
	return _WooracleV2.Contract.WoState(&_WooracleV2.CallOpts, base)
}

// WoState is a free data retrieval call binding the contract method 0xc16116d4.
//
// Solidity: function woState(address base) view returns((uint128,uint64,uint64,bool))
func (_WooracleV2 *WooracleV2CallerSession) WoState(base common.Address) (IWooracleV2State, error) {
	return _WooracleV2.Contract.WoState(&_WooracleV2.CallOpts, base)
}

// PostPrice is a paid mutator transaction binding the contract method 0xd5bade07.
//
// Solidity: function postPrice(address base, uint128 newPrice) returns()
func (_WooracleV2 *WooracleV2Transactor) PostPrice(opts *bind.TransactOpts, base common.Address, newPrice *big.Int) (*types.Transaction, error) {
	return _WooracleV2.contract.Transact(opts, "postPrice", base, newPrice)
}

// PostPrice is a paid mutator transaction binding the contract method 0xd5bade07.
//
// Solidity: function postPrice(address base, uint128 newPrice) returns()
func (_WooracleV2 *WooracleV2Session) PostPrice(base common.Address, newPrice *big.Int) (*types.Transaction, error) {
	return _WooracleV2.Contract.PostPrice(&_WooracleV2.TransactOpts, base, newPrice)
}

// PostPrice is a paid mutator transaction binding the contract method 0xd5bade07.
//
// Solidity: function postPrice(address base, uint128 newPrice) returns()
func (_WooracleV2 *WooracleV2TransactorSession) PostPrice(base common.Address, newPrice *big.Int) (*types.Transaction, error) {
	return _WooracleV2.Contract.PostPrice(&_WooracleV2.TransactOpts, base, newPrice)
}

// PostPriceList is a paid mutator transaction binding the contract method 0x49230eab.
//
// Solidity: function postPriceList(address[] bases, uint128[] newPrices) returns()
func (_WooracleV2 *WooracleV2Transactor) PostPriceList(opts *bind.TransactOpts, bases []common.Address, newPrices []*big.Int) (*types.Transaction, error) {
	return _WooracleV2.contract.Transact(opts, "postPriceList", bases, newPrices)
}

// PostPriceList is a paid mutator transaction binding the contract method 0x49230eab.
//
// Solidity: function postPriceList(address[] bases, uint128[] newPrices) returns()
func (_WooracleV2 *WooracleV2Session) PostPriceList(bases []common.Address, newPrices []*big.Int) (*types.Transaction, error) {
	return _WooracleV2.Contract.PostPriceList(&_WooracleV2.TransactOpts, bases, newPrices)
}

// PostPriceList is a paid mutator transaction binding the contract method 0x49230eab.
//
// Solidity: function postPriceList(address[] bases, uint128[] newPrices) returns()
func (_WooracleV2 *WooracleV2TransactorSession) PostPriceList(bases []common.Address, newPrices []*big.Int) (*types.Transaction, error) {
	return _WooracleV2.Contract.PostPriceList(&_WooracleV2.TransactOpts, bases, newPrices)
}

// PostSpread is a paid mutator transaction binding the contract method 0xba1eba68.
//
// Solidity: function postSpread(address base, uint64 newSpread) returns()
func (_WooracleV2 *WooracleV2Transactor) PostSpread(opts *bind.TransactOpts, base common.Address, newSpread uint64) (*types.Transaction, error) {
	return _WooracleV2.contract.Transact(opts, "postSpread", base, newSpread)
}

// PostSpread is a paid mutator transaction binding the contract method 0xba1eba68.
//
// Solidity: function postSpread(address base, uint64 newSpread) returns()
func (_WooracleV2 *WooracleV2Session) PostSpread(base common.Address, newSpread uint64) (*types.Transaction, error) {
	return _WooracleV2.Contract.PostSpread(&_WooracleV2.TransactOpts, base, newSpread)
}

// PostSpread is a paid mutator transaction binding the contract method 0xba1eba68.
//
// Solidity: function postSpread(address base, uint64 newSpread) returns()
func (_WooracleV2 *WooracleV2TransactorSession) PostSpread(base common.Address, newSpread uint64) (*types.Transaction, error) {
	return _WooracleV2.Contract.PostSpread(&_WooracleV2.TransactOpts, base, newSpread)
}

// PostSpreadList is a paid mutator transaction binding the contract method 0x18e07221.
//
// Solidity: function postSpreadList(address[] bases, uint64[] newSpreads) returns()
func (_WooracleV2 *WooracleV2Transactor) PostSpreadList(opts *bind.TransactOpts, bases []common.Address, newSpreads []uint64) (*types.Transaction, error) {
	return _WooracleV2.contract.Transact(opts, "postSpreadList", bases, newSpreads)
}

// PostSpreadList is a paid mutator transaction binding the contract method 0x18e07221.
//
// Solidity: function postSpreadList(address[] bases, uint64[] newSpreads) returns()
func (_WooracleV2 *WooracleV2Session) PostSpreadList(bases []common.Address, newSpreads []uint64) (*types.Transaction, error) {
	return _WooracleV2.Contract.PostSpreadList(&_WooracleV2.TransactOpts, bases, newSpreads)
}

// PostSpreadList is a paid mutator transaction binding the contract method 0x18e07221.
//
// Solidity: function postSpreadList(address[] bases, uint64[] newSpreads) returns()
func (_WooracleV2 *WooracleV2TransactorSession) PostSpreadList(bases []common.Address, newSpreads []uint64) (*types.Transaction, error) {
	return _WooracleV2.Contract.PostSpreadList(&_WooracleV2.TransactOpts, bases, newSpreads)
}

// PostState is a paid mutator transaction binding the contract method 0x71ea9205.
//
// Solidity: function postState(address base, uint128 newPrice, uint64 newSpread, uint64 newCoeff) returns()
func (_WooracleV2 *WooracleV2Transactor) PostState(opts *bind.TransactOpts, base common.Address, newPrice *big.Int, newSpread uint64, newCoeff uint64) (*types.Transaction, error) {
	return _WooracleV2.contract.Transact(opts, "postState", base, newPrice, newSpread, newCoeff)
}

// PostState is a paid mutator transaction binding the contract method 0x71ea9205.
//
// Solidity: function postState(address base, uint128 newPrice, uint64 newSpread, uint64 newCoeff) returns()
func (_WooracleV2 *WooracleV2Session) PostState(base common.Address, newPrice *big.Int, newSpread uint64, newCoeff uint64) (*types.Transaction, error) {
	return _WooracleV2.Contract.PostState(&_WooracleV2.TransactOpts, base, newPrice, newSpread, newCoeff)
}

// PostState is a paid mutator transaction binding the contract method 0x71ea9205.
//
// Solidity: function postState(address base, uint128 newPrice, uint64 newSpread, uint64 newCoeff) returns()
func (_WooracleV2 *WooracleV2TransactorSession) PostState(base common.Address, newPrice *big.Int, newSpread uint64, newCoeff uint64) (*types.Transaction, error) {
	return _WooracleV2.Contract.PostState(&_WooracleV2.TransactOpts, base, newPrice, newSpread, newCoeff)
}

// PostStateList is a paid mutator transaction binding the contract method 0x7967d37e.
//
// Solidity: function postStateList(address[] bases, uint128[] newPrices, uint64[] newSpreads, uint64[] newCoeffs) returns()
func (_WooracleV2 *WooracleV2Transactor) PostStateList(opts *bind.TransactOpts, bases []common.Address, newPrices []*big.Int, newSpreads []uint64, newCoeffs []uint64) (*types.Transaction, error) {
	return _WooracleV2.contract.Transact(opts, "postStateList", bases, newPrices, newSpreads, newCoeffs)
}

// PostStateList is a paid mutator transaction binding the contract method 0x7967d37e.
//
// Solidity: function postStateList(address[] bases, uint128[] newPrices, uint64[] newSpreads, uint64[] newCoeffs) returns()
func (_WooracleV2 *WooracleV2Session) PostStateList(bases []common.Address, newPrices []*big.Int, newSpreads []uint64, newCoeffs []uint64) (*types.Transaction, error) {
	return _WooracleV2.Contract.PostStateList(&_WooracleV2.TransactOpts, bases, newPrices, newSpreads, newCoeffs)
}

// PostStateList is a paid mutator transaction binding the contract method 0x7967d37e.
//
// Solidity: function postStateList(address[] bases, uint128[] newPrices, uint64[] newSpreads, uint64[] newCoeffs) returns()
func (_WooracleV2 *WooracleV2TransactorSession) PostStateList(bases []common.Address, newPrices []*big.Int, newSpreads []uint64, newCoeffs []uint64) (*types.Transaction, error) {
	return _WooracleV2.Contract.PostStateList(&_WooracleV2.TransactOpts, bases, newPrices, newSpreads, newCoeffs)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WooracleV2 *WooracleV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WooracleV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WooracleV2 *WooracleV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _WooracleV2.Contract.RenounceOwnership(&_WooracleV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WooracleV2 *WooracleV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WooracleV2.Contract.RenounceOwnership(&_WooracleV2.TransactOpts)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address addr, bool flag) returns()
func (_WooracleV2 *WooracleV2Transactor) SetAdmin(opts *bind.TransactOpts, addr common.Address, flag bool) (*types.Transaction, error) {
	return _WooracleV2.contract.Transact(opts, "setAdmin", addr, flag)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address addr, bool flag) returns()
func (_WooracleV2 *WooracleV2Session) SetAdmin(addr common.Address, flag bool) (*types.Transaction, error) {
	return _WooracleV2.Contract.SetAdmin(&_WooracleV2.TransactOpts, addr, flag)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address addr, bool flag) returns()
func (_WooracleV2 *WooracleV2TransactorSession) SetAdmin(addr common.Address, flag bool) (*types.Transaction, error) {
	return _WooracleV2.Contract.SetAdmin(&_WooracleV2.TransactOpts, addr, flag)
}

// SetBound is a paid mutator transaction binding the contract method 0x1142e753.
//
// Solidity: function setBound(uint64 _bound) returns()
func (_WooracleV2 *WooracleV2Transactor) SetBound(opts *bind.TransactOpts, _bound uint64) (*types.Transaction, error) {
	return _WooracleV2.contract.Transact(opts, "setBound", _bound)
}

// SetBound is a paid mutator transaction binding the contract method 0x1142e753.
//
// Solidity: function setBound(uint64 _bound) returns()
func (_WooracleV2 *WooracleV2Session) SetBound(_bound uint64) (*types.Transaction, error) {
	return _WooracleV2.Contract.SetBound(&_WooracleV2.TransactOpts, _bound)
}

// SetBound is a paid mutator transaction binding the contract method 0x1142e753.
//
// Solidity: function setBound(uint64 _bound) returns()
func (_WooracleV2 *WooracleV2TransactorSession) SetBound(_bound uint64) (*types.Transaction, error) {
	return _WooracleV2.Contract.SetBound(&_WooracleV2.TransactOpts, _bound)
}

// SetCLOracle is a paid mutator transaction binding the contract method 0xa4a2a8c5.
//
// Solidity: function setCLOracle(address token, address _oracle, bool _cloPreferred) returns()
func (_WooracleV2 *WooracleV2Transactor) SetCLOracle(opts *bind.TransactOpts, token common.Address, _oracle common.Address, _cloPreferred bool) (*types.Transaction, error) {
	return _WooracleV2.contract.Transact(opts, "setCLOracle", token, _oracle, _cloPreferred)
}

// SetCLOracle is a paid mutator transaction binding the contract method 0xa4a2a8c5.
//
// Solidity: function setCLOracle(address token, address _oracle, bool _cloPreferred) returns()
func (_WooracleV2 *WooracleV2Session) SetCLOracle(token common.Address, _oracle common.Address, _cloPreferred bool) (*types.Transaction, error) {
	return _WooracleV2.Contract.SetCLOracle(&_WooracleV2.TransactOpts, token, _oracle, _cloPreferred)
}

// SetCLOracle is a paid mutator transaction binding the contract method 0xa4a2a8c5.
//
// Solidity: function setCLOracle(address token, address _oracle, bool _cloPreferred) returns()
func (_WooracleV2 *WooracleV2TransactorSession) SetCLOracle(token common.Address, _oracle common.Address, _cloPreferred bool) (*types.Transaction, error) {
	return _WooracleV2.Contract.SetCLOracle(&_WooracleV2.TransactOpts, token, _oracle, _cloPreferred)
}

// SetCloPreferred is a paid mutator transaction binding the contract method 0x1ffabeb8.
//
// Solidity: function setCloPreferred(address token, bool _cloPreferred) returns()
func (_WooracleV2 *WooracleV2Transactor) SetCloPreferred(opts *bind.TransactOpts, token common.Address, _cloPreferred bool) (*types.Transaction, error) {
	return _WooracleV2.contract.Transact(opts, "setCloPreferred", token, _cloPreferred)
}

// SetCloPreferred is a paid mutator transaction binding the contract method 0x1ffabeb8.
//
// Solidity: function setCloPreferred(address token, bool _cloPreferred) returns()
func (_WooracleV2 *WooracleV2Session) SetCloPreferred(token common.Address, _cloPreferred bool) (*types.Transaction, error) {
	return _WooracleV2.Contract.SetCloPreferred(&_WooracleV2.TransactOpts, token, _cloPreferred)
}

// SetCloPreferred is a paid mutator transaction binding the contract method 0x1ffabeb8.
//
// Solidity: function setCloPreferred(address token, bool _cloPreferred) returns()
func (_WooracleV2 *WooracleV2TransactorSession) SetCloPreferred(token common.Address, _cloPreferred bool) (*types.Transaction, error) {
	return _WooracleV2.Contract.SetCloPreferred(&_WooracleV2.TransactOpts, token, _cloPreferred)
}

// SetQuoteToken is a paid mutator transaction binding the contract method 0x6e27fcd6.
//
// Solidity: function setQuoteToken(address _quote, address _oracle) returns()
func (_WooracleV2 *WooracleV2Transactor) SetQuoteToken(opts *bind.TransactOpts, _quote common.Address, _oracle common.Address) (*types.Transaction, error) {
	return _WooracleV2.contract.Transact(opts, "setQuoteToken", _quote, _oracle)
}

// SetQuoteToken is a paid mutator transaction binding the contract method 0x6e27fcd6.
//
// Solidity: function setQuoteToken(address _quote, address _oracle) returns()
func (_WooracleV2 *WooracleV2Session) SetQuoteToken(_quote common.Address, _oracle common.Address) (*types.Transaction, error) {
	return _WooracleV2.Contract.SetQuoteToken(&_WooracleV2.TransactOpts, _quote, _oracle)
}

// SetQuoteToken is a paid mutator transaction binding the contract method 0x6e27fcd6.
//
// Solidity: function setQuoteToken(address _quote, address _oracle) returns()
func (_WooracleV2 *WooracleV2TransactorSession) SetQuoteToken(_quote common.Address, _oracle common.Address) (*types.Transaction, error) {
	return _WooracleV2.Contract.SetQuoteToken(&_WooracleV2.TransactOpts, _quote, _oracle)
}

// SetStaleDuration is a paid mutator transaction binding the contract method 0x99235fd4.
//
// Solidity: function setStaleDuration(uint256 newStaleDuration) returns()
func (_WooracleV2 *WooracleV2Transactor) SetStaleDuration(opts *bind.TransactOpts, newStaleDuration *big.Int) (*types.Transaction, error) {
	return _WooracleV2.contract.Transact(opts, "setStaleDuration", newStaleDuration)
}

// SetStaleDuration is a paid mutator transaction binding the contract method 0x99235fd4.
//
// Solidity: function setStaleDuration(uint256 newStaleDuration) returns()
func (_WooracleV2 *WooracleV2Session) SetStaleDuration(newStaleDuration *big.Int) (*types.Transaction, error) {
	return _WooracleV2.Contract.SetStaleDuration(&_WooracleV2.TransactOpts, newStaleDuration)
}

// SetStaleDuration is a paid mutator transaction binding the contract method 0x99235fd4.
//
// Solidity: function setStaleDuration(uint256 newStaleDuration) returns()
func (_WooracleV2 *WooracleV2TransactorSession) SetStaleDuration(newStaleDuration *big.Int) (*types.Transaction, error) {
	return _WooracleV2.Contract.SetStaleDuration(&_WooracleV2.TransactOpts, newStaleDuration)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WooracleV2 *WooracleV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WooracleV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WooracleV2 *WooracleV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WooracleV2.Contract.TransferOwnership(&_WooracleV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WooracleV2 *WooracleV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WooracleV2.Contract.TransferOwnership(&_WooracleV2.TransactOpts, newOwner)
}

// WooracleV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WooracleV2 contract.
type WooracleV2OwnershipTransferredIterator struct {
	Event *WooracleV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WooracleV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WooracleV2OwnershipTransferred)
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
		it.Event = new(WooracleV2OwnershipTransferred)
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
func (it *WooracleV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WooracleV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WooracleV2OwnershipTransferred represents a OwnershipTransferred event raised by the WooracleV2 contract.
type WooracleV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WooracleV2 *WooracleV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WooracleV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WooracleV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WooracleV2OwnershipTransferredIterator{contract: _WooracleV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WooracleV2 *WooracleV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WooracleV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WooracleV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WooracleV2OwnershipTransferred)
				if err := _WooracleV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_WooracleV2 *WooracleV2Filterer) ParseOwnershipTransferred(log types.Log) (*WooracleV2OwnershipTransferred, error) {
	event := new(WooracleV2OwnershipTransferred)
	if err := _WooracleV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
