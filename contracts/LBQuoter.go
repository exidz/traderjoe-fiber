// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
	_ = abi.ConvertType
)

// LBQuoterQuote is an auto generated low-level Go binding around an user-defined struct.
type LBQuoterQuote struct {
	Route                         []common.Address
	Pairs                         []common.Address
	BinSteps                      []*big.Int
	Versions                      []uint8
	Amounts                       []*big.Int
	VirtualAmountsWithoutSlippage []*big.Int
	Fees                          []*big.Int
}

// LBQuoterMetaData contains all meta data concerning the LBQuoter contract.
var LBQuoterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factoryV1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"legacyFactoryV2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"factoryV2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"legacyRouterV2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"routerV2\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"JoeLibrary__AddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"JoeLibrary__IdenticalAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"JoeLibrary__InsufficientAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"JoeLibrary__InsufficientLiquidity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBQuoter_InvalidLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeCast__Exceeds128Bits\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeCast__Exceeds24Bits\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"y\",\"type\":\"int256\"}],\"name\":\"Uint128x128Math__PowUnderflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Uint256x256Math__MulDivOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Uint256x256Math__MulShiftOverflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"route\",\"type\":\"address[]\"},{\"internalType\":\"uint128\",\"name\":\"amountIn\",\"type\":\"uint128\"}],\"name\":\"findBestPathFromAmountIn\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"route\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"pairs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"binSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"enumILBRouter.Version[]\",\"name\":\"versions\",\"type\":\"uint8[]\"},{\"internalType\":\"uint128[]\",\"name\":\"amounts\",\"type\":\"uint128[]\"},{\"internalType\":\"uint128[]\",\"name\":\"virtualAmountsWithoutSlippage\",\"type\":\"uint128[]\"},{\"internalType\":\"uint128[]\",\"name\":\"fees\",\"type\":\"uint128[]\"}],\"internalType\":\"structLBQuoter.Quote\",\"name\":\"quote\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"route\",\"type\":\"address[]\"},{\"internalType\":\"uint128\",\"name\":\"amountOut\",\"type\":\"uint128\"}],\"name\":\"findBestPathFromAmountOut\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"route\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"pairs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"binSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"enumILBRouter.Version[]\",\"name\":\"versions\",\"type\":\"uint8[]\"},{\"internalType\":\"uint128[]\",\"name\":\"amounts\",\"type\":\"uint128[]\"},{\"internalType\":\"uint128[]\",\"name\":\"virtualAmountsWithoutSlippage\",\"type\":\"uint128[]\"},{\"internalType\":\"uint128[]\",\"name\":\"fees\",\"type\":\"uint128[]\"}],\"internalType\":\"structLBQuoter.Quote\",\"name\":\"quote\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFactoryV1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"factoryV1\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFactoryV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"factoryV2\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLegacyFactoryV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"legacyFactoryV2\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLegacyRouterV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"legacyRouterV2\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouterV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"routerV2\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LBQuoterABI is the input ABI used to generate the binding from.
// Deprecated: Use LBQuoterMetaData.ABI instead.
var LBQuoterABI = LBQuoterMetaData.ABI

// LBQuoter is an auto generated Go binding around an Ethereum contract.
type LBQuoter struct {
	LBQuoterCaller     // Read-only binding to the contract
	LBQuoterTransactor // Write-only binding to the contract
	LBQuoterFilterer   // Log filterer for contract events
}

// LBQuoterCaller is an auto generated read-only Go binding around an Ethereum contract.
type LBQuoterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LBQuoterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LBQuoterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LBQuoterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LBQuoterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LBQuoterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LBQuoterSession struct {
	Contract     *LBQuoter         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LBQuoterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LBQuoterCallerSession struct {
	Contract *LBQuoterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// LBQuoterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LBQuoterTransactorSession struct {
	Contract     *LBQuoterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LBQuoterRaw is an auto generated low-level Go binding around an Ethereum contract.
type LBQuoterRaw struct {
	Contract *LBQuoter // Generic contract binding to access the raw methods on
}

// LBQuoterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LBQuoterCallerRaw struct {
	Contract *LBQuoterCaller // Generic read-only contract binding to access the raw methods on
}

// LBQuoterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LBQuoterTransactorRaw struct {
	Contract *LBQuoterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLBQuoter creates a new instance of LBQuoter, bound to a specific deployed contract.
func NewLBQuoter(address common.Address, backend bind.ContractBackend) (*LBQuoter, error) {
	contract, err := bindLBQuoter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LBQuoter{LBQuoterCaller: LBQuoterCaller{contract: contract}, LBQuoterTransactor: LBQuoterTransactor{contract: contract}, LBQuoterFilterer: LBQuoterFilterer{contract: contract}}, nil
}

// NewLBQuoterCaller creates a new read-only instance of LBQuoter, bound to a specific deployed contract.
func NewLBQuoterCaller(address common.Address, caller bind.ContractCaller) (*LBQuoterCaller, error) {
	contract, err := bindLBQuoter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LBQuoterCaller{contract: contract}, nil
}

// NewLBQuoterTransactor creates a new write-only instance of LBQuoter, bound to a specific deployed contract.
func NewLBQuoterTransactor(address common.Address, transactor bind.ContractTransactor) (*LBQuoterTransactor, error) {
	contract, err := bindLBQuoter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LBQuoterTransactor{contract: contract}, nil
}

// NewLBQuoterFilterer creates a new log filterer instance of LBQuoter, bound to a specific deployed contract.
func NewLBQuoterFilterer(address common.Address, filterer bind.ContractFilterer) (*LBQuoterFilterer, error) {
	contract, err := bindLBQuoter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LBQuoterFilterer{contract: contract}, nil
}

// bindLBQuoter binds a generic wrapper to an already deployed contract.
func bindLBQuoter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LBQuoterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LBQuoter *LBQuoterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LBQuoter.Contract.LBQuoterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LBQuoter *LBQuoterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LBQuoter.Contract.LBQuoterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LBQuoter *LBQuoterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LBQuoter.Contract.LBQuoterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LBQuoter *LBQuoterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LBQuoter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LBQuoter *LBQuoterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LBQuoter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LBQuoter *LBQuoterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LBQuoter.Contract.contract.Transact(opts, method, params...)
}

// FindBestPathFromAmountIn is a free data retrieval call binding the contract method 0x0f902a40.
//
// Solidity: function findBestPathFromAmountIn(address[] route, uint128 amountIn) view returns((address[],address[],uint256[],uint8[],uint128[],uint128[],uint128[]) quote)
func (_LBQuoter *LBQuoterCaller) FindBestPathFromAmountIn(opts *bind.CallOpts, route []common.Address, amountIn *big.Int) (LBQuoterQuote, error) {
	var out []interface{}
	err := _LBQuoter.contract.Call(opts, &out, "findBestPathFromAmountIn", route, amountIn)

	if err != nil {
		return *new(LBQuoterQuote), err
	}

	out0 := *abi.ConvertType(out[0], new(LBQuoterQuote)).(*LBQuoterQuote)

	return out0, err

}

// FindBestPathFromAmountIn is a free data retrieval call binding the contract method 0x0f902a40.
//
// Solidity: function findBestPathFromAmountIn(address[] route, uint128 amountIn) view returns((address[],address[],uint256[],uint8[],uint128[],uint128[],uint128[]) quote)
func (_LBQuoter *LBQuoterSession) FindBestPathFromAmountIn(route []common.Address, amountIn *big.Int) (LBQuoterQuote, error) {
	return _LBQuoter.Contract.FindBestPathFromAmountIn(&_LBQuoter.CallOpts, route, amountIn)
}

// FindBestPathFromAmountIn is a free data retrieval call binding the contract method 0x0f902a40.
//
// Solidity: function findBestPathFromAmountIn(address[] route, uint128 amountIn) view returns((address[],address[],uint256[],uint8[],uint128[],uint128[],uint128[]) quote)
func (_LBQuoter *LBQuoterCallerSession) FindBestPathFromAmountIn(route []common.Address, amountIn *big.Int) (LBQuoterQuote, error) {
	return _LBQuoter.Contract.FindBestPathFromAmountIn(&_LBQuoter.CallOpts, route, amountIn)
}

// FindBestPathFromAmountOut is a free data retrieval call binding the contract method 0x59214226.
//
// Solidity: function findBestPathFromAmountOut(address[] route, uint128 amountOut) view returns((address[],address[],uint256[],uint8[],uint128[],uint128[],uint128[]) quote)
func (_LBQuoter *LBQuoterCaller) FindBestPathFromAmountOut(opts *bind.CallOpts, route []common.Address, amountOut *big.Int) (LBQuoterQuote, error) {
	var out []interface{}
	err := _LBQuoter.contract.Call(opts, &out, "findBestPathFromAmountOut", route, amountOut)

	if err != nil {
		return *new(LBQuoterQuote), err
	}

	out0 := *abi.ConvertType(out[0], new(LBQuoterQuote)).(*LBQuoterQuote)

	return out0, err

}

// FindBestPathFromAmountOut is a free data retrieval call binding the contract method 0x59214226.
//
// Solidity: function findBestPathFromAmountOut(address[] route, uint128 amountOut) view returns((address[],address[],uint256[],uint8[],uint128[],uint128[],uint128[]) quote)
func (_LBQuoter *LBQuoterSession) FindBestPathFromAmountOut(route []common.Address, amountOut *big.Int) (LBQuoterQuote, error) {
	return _LBQuoter.Contract.FindBestPathFromAmountOut(&_LBQuoter.CallOpts, route, amountOut)
}

// FindBestPathFromAmountOut is a free data retrieval call binding the contract method 0x59214226.
//
// Solidity: function findBestPathFromAmountOut(address[] route, uint128 amountOut) view returns((address[],address[],uint256[],uint8[],uint128[],uint128[],uint128[]) quote)
func (_LBQuoter *LBQuoterCallerSession) FindBestPathFromAmountOut(route []common.Address, amountOut *big.Int) (LBQuoterQuote, error) {
	return _LBQuoter.Contract.FindBestPathFromAmountOut(&_LBQuoter.CallOpts, route, amountOut)
}

// GetFactoryV1 is a free data retrieval call binding the contract method 0x07da8f57.
//
// Solidity: function getFactoryV1() view returns(address factoryV1)
func (_LBQuoter *LBQuoterCaller) GetFactoryV1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LBQuoter.contract.Call(opts, &out, "getFactoryV1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFactoryV1 is a free data retrieval call binding the contract method 0x07da8f57.
//
// Solidity: function getFactoryV1() view returns(address factoryV1)
func (_LBQuoter *LBQuoterSession) GetFactoryV1() (common.Address, error) {
	return _LBQuoter.Contract.GetFactoryV1(&_LBQuoter.CallOpts)
}

// GetFactoryV1 is a free data retrieval call binding the contract method 0x07da8f57.
//
// Solidity: function getFactoryV1() view returns(address factoryV1)
func (_LBQuoter *LBQuoterCallerSession) GetFactoryV1() (common.Address, error) {
	return _LBQuoter.Contract.GetFactoryV1(&_LBQuoter.CallOpts)
}

// GetFactoryV2 is a free data retrieval call binding the contract method 0xca56fc72.
//
// Solidity: function getFactoryV2() view returns(address factoryV2)
func (_LBQuoter *LBQuoterCaller) GetFactoryV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LBQuoter.contract.Call(opts, &out, "getFactoryV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFactoryV2 is a free data retrieval call binding the contract method 0xca56fc72.
//
// Solidity: function getFactoryV2() view returns(address factoryV2)
func (_LBQuoter *LBQuoterSession) GetFactoryV2() (common.Address, error) {
	return _LBQuoter.Contract.GetFactoryV2(&_LBQuoter.CallOpts)
}

// GetFactoryV2 is a free data retrieval call binding the contract method 0xca56fc72.
//
// Solidity: function getFactoryV2() view returns(address factoryV2)
func (_LBQuoter *LBQuoterCallerSession) GetFactoryV2() (common.Address, error) {
	return _LBQuoter.Contract.GetFactoryV2(&_LBQuoter.CallOpts)
}

// GetLegacyFactoryV2 is a free data retrieval call binding the contract method 0x8fe4b3ad.
//
// Solidity: function getLegacyFactoryV2() view returns(address legacyFactoryV2)
func (_LBQuoter *LBQuoterCaller) GetLegacyFactoryV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LBQuoter.contract.Call(opts, &out, "getLegacyFactoryV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLegacyFactoryV2 is a free data retrieval call binding the contract method 0x8fe4b3ad.
//
// Solidity: function getLegacyFactoryV2() view returns(address legacyFactoryV2)
func (_LBQuoter *LBQuoterSession) GetLegacyFactoryV2() (common.Address, error) {
	return _LBQuoter.Contract.GetLegacyFactoryV2(&_LBQuoter.CallOpts)
}

// GetLegacyFactoryV2 is a free data retrieval call binding the contract method 0x8fe4b3ad.
//
// Solidity: function getLegacyFactoryV2() view returns(address legacyFactoryV2)
func (_LBQuoter *LBQuoterCallerSession) GetLegacyFactoryV2() (common.Address, error) {
	return _LBQuoter.Contract.GetLegacyFactoryV2(&_LBQuoter.CallOpts)
}

// GetLegacyRouterV2 is a free data retrieval call binding the contract method 0x23229d6d.
//
// Solidity: function getLegacyRouterV2() view returns(address legacyRouterV2)
func (_LBQuoter *LBQuoterCaller) GetLegacyRouterV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LBQuoter.contract.Call(opts, &out, "getLegacyRouterV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLegacyRouterV2 is a free data retrieval call binding the contract method 0x23229d6d.
//
// Solidity: function getLegacyRouterV2() view returns(address legacyRouterV2)
func (_LBQuoter *LBQuoterSession) GetLegacyRouterV2() (common.Address, error) {
	return _LBQuoter.Contract.GetLegacyRouterV2(&_LBQuoter.CallOpts)
}

// GetLegacyRouterV2 is a free data retrieval call binding the contract method 0x23229d6d.
//
// Solidity: function getLegacyRouterV2() view returns(address legacyRouterV2)
func (_LBQuoter *LBQuoterCallerSession) GetLegacyRouterV2() (common.Address, error) {
	return _LBQuoter.Contract.GetLegacyRouterV2(&_LBQuoter.CallOpts)
}

// GetRouterV2 is a free data retrieval call binding the contract method 0x33141d3e.
//
// Solidity: function getRouterV2() view returns(address routerV2)
func (_LBQuoter *LBQuoterCaller) GetRouterV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LBQuoter.contract.Call(opts, &out, "getRouterV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRouterV2 is a free data retrieval call binding the contract method 0x33141d3e.
//
// Solidity: function getRouterV2() view returns(address routerV2)
func (_LBQuoter *LBQuoterSession) GetRouterV2() (common.Address, error) {
	return _LBQuoter.Contract.GetRouterV2(&_LBQuoter.CallOpts)
}

// GetRouterV2 is a free data retrieval call binding the contract method 0x33141d3e.
//
// Solidity: function getRouterV2() view returns(address routerV2)
func (_LBQuoter *LBQuoterCallerSession) GetRouterV2() (common.Address, error) {
	return _LBQuoter.Contract.GetRouterV2(&_LBQuoter.CallOpts)
}
