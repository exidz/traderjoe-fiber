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
type LegacyQuoterQuote struct {
	Route                         []common.Address
	Pairs                         []common.Address
	BinSteps                      []*big.Int
	Amounts                       []*big.Int
	VirtualAmountsWithoutSlippage []*big.Int
	Fees                          []*big.Int
}

// LegacyQuoterMetaData contains all meta data concerning the LegacyQuoter contract.
var LegacyQuoterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_routerV2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factoryV1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factoryV2\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bp\",\"type\":\"uint256\"}],\"name\":\"BinHelper__BinStepOverflows\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BinHelper__IdOverflows\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"JoeLibrary__AddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"JoeLibrary__IdenticalAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"JoeLibrary__InsufficientAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"JoeLibrary__InsufficientLiquidity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBQuoter_InvalidLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"y\",\"type\":\"int256\"}],\"name\":\"Math128x128__PowerUnderflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"prod1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"}],\"name\":\"Math512Bits__MulDivOverflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"prod1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"Math512Bits__MulShiftOverflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"Math512Bits__OffsetOverflows\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"factoryV1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factoryV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_route\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"}],\"name\":\"findBestPathFromAmountIn\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"route\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"pairs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"binSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"virtualAmountsWithoutSlippage\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"fees\",\"type\":\"uint256[]\"}],\"internalType\":\"structLBQuoter.Quote\",\"name\":\"quote\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_route\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"}],\"name\":\"findBestPathFromAmountOut\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"route\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"pairs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"binSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"virtualAmountsWithoutSlippage\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"fees\",\"type\":\"uint256[]\"}],\"internalType\":\"structLBQuoter.Quote\",\"name\":\"quote\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"routerV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LegacyQuoterABI is the input ABI used to generate the binding from.
// Deprecated: Use LegacyQuoterMetaData.ABI instead.
var LegacyQuoterABI = LegacyQuoterMetaData.ABI

// LegacyQuoter is an auto generated Go binding around an Ethereum contract.
type LegacyQuoter struct {
	LegacyQuoterCaller     // Read-only binding to the contract
	LegacyQuoterTransactor // Write-only binding to the contract
	LegacyQuoterFilterer   // Log filterer for contract events
}

// LegacyQuoterCaller is an auto generated read-only Go binding around an Ethereum contract.
type LegacyQuoterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyQuoterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LegacyQuoterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyQuoterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LegacyQuoterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyQuoterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LegacyQuoterSession struct {
	Contract     *LegacyQuoter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LegacyQuoterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LegacyQuoterCallerSession struct {
	Contract *LegacyQuoterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// LegacyQuoterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LegacyQuoterTransactorSession struct {
	Contract     *LegacyQuoterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LegacyQuoterRaw is an auto generated low-level Go binding around an Ethereum contract.
type LegacyQuoterRaw struct {
	Contract *LegacyQuoter // Generic contract binding to access the raw methods on
}

// LegacyQuoterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LegacyQuoterCallerRaw struct {
	Contract *LegacyQuoterCaller // Generic read-only contract binding to access the raw methods on
}

// LegacyQuoterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LegacyQuoterTransactorRaw struct {
	Contract *LegacyQuoterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLegacyQuoter creates a new instance of LegacyQuoter, bound to a specific deployed contract.
func NewLegacyQuoter(address common.Address, backend bind.ContractBackend) (*LegacyQuoter, error) {
	contract, err := bindLegacyQuoter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LegacyQuoter{LegacyQuoterCaller: LegacyQuoterCaller{contract: contract}, LegacyQuoterTransactor: LegacyQuoterTransactor{contract: contract}, LegacyQuoterFilterer: LegacyQuoterFilterer{contract: contract}}, nil
}

// NewLegacyQuoterCaller creates a new read-only instance of LegacyQuoter, bound to a specific deployed contract.
func NewLegacyQuoterCaller(address common.Address, caller bind.ContractCaller) (*LegacyQuoterCaller, error) {
	contract, err := bindLegacyQuoter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LegacyQuoterCaller{contract: contract}, nil
}

// NewLegacyQuoterTransactor creates a new write-only instance of LegacyQuoter, bound to a specific deployed contract.
func NewLegacyQuoterTransactor(address common.Address, transactor bind.ContractTransactor) (*LegacyQuoterTransactor, error) {
	contract, err := bindLegacyQuoter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LegacyQuoterTransactor{contract: contract}, nil
}

// NewLegacyQuoterFilterer creates a new log filterer instance of LegacyQuoter, bound to a specific deployed contract.
func NewLegacyQuoterFilterer(address common.Address, filterer bind.ContractFilterer) (*LegacyQuoterFilterer, error) {
	contract, err := bindLegacyQuoter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LegacyQuoterFilterer{contract: contract}, nil
}

// bindLegacyQuoter binds a generic wrapper to an already deployed contract.
func bindLegacyQuoter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LegacyQuoterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LegacyQuoter *LegacyQuoterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LegacyQuoter.Contract.LegacyQuoterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LegacyQuoter *LegacyQuoterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LegacyQuoter.Contract.LegacyQuoterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LegacyQuoter *LegacyQuoterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LegacyQuoter.Contract.LegacyQuoterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LegacyQuoter *LegacyQuoterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LegacyQuoter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LegacyQuoter *LegacyQuoterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LegacyQuoter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LegacyQuoter *LegacyQuoterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LegacyQuoter.Contract.contract.Transact(opts, method, params...)
}

// FactoryV1 is a free data retrieval call binding the contract method 0x3957f453.
//
// Solidity: function factoryV1() view returns(address)
func (_LegacyQuoter *LegacyQuoterCaller) FactoryV1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LegacyQuoter.contract.Call(opts, &out, "factoryV1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryV1 is a free data retrieval call binding the contract method 0x3957f453.
//
// Solidity: function factoryV1() view returns(address)
func (_LegacyQuoter *LegacyQuoterSession) FactoryV1() (common.Address, error) {
	return _LegacyQuoter.Contract.FactoryV1(&_LegacyQuoter.CallOpts)
}

// FactoryV1 is a free data retrieval call binding the contract method 0x3957f453.
//
// Solidity: function factoryV1() view returns(address)
func (_LegacyQuoter *LegacyQuoterCallerSession) FactoryV1() (common.Address, error) {
	return _LegacyQuoter.Contract.FactoryV1(&_LegacyQuoter.CallOpts)
}

// FactoryV2 is a free data retrieval call binding the contract method 0x68e0d4e1.
//
// Solidity: function factoryV2() view returns(address)
func (_LegacyQuoter *LegacyQuoterCaller) FactoryV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LegacyQuoter.contract.Call(opts, &out, "factoryV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryV2 is a free data retrieval call binding the contract method 0x68e0d4e1.
//
// Solidity: function factoryV2() view returns(address)
func (_LegacyQuoter *LegacyQuoterSession) FactoryV2() (common.Address, error) {
	return _LegacyQuoter.Contract.FactoryV2(&_LegacyQuoter.CallOpts)
}

// FactoryV2 is a free data retrieval call binding the contract method 0x68e0d4e1.
//
// Solidity: function factoryV2() view returns(address)
func (_LegacyQuoter *LegacyQuoterCallerSession) FactoryV2() (common.Address, error) {
	return _LegacyQuoter.Contract.FactoryV2(&_LegacyQuoter.CallOpts)
}

// FindBestPathFromAmountIn is a free data retrieval call binding the contract method 0xc1b7687b.
//
// Solidity: function findBestPathFromAmountIn(address[] _route, uint256 _amountIn) view returns((address[],address[],uint256[],uint256[],uint256[],uint256[]) quote)
func (_LegacyQuoter *LegacyQuoterCaller) FindBestPathFromAmountIn(opts *bind.CallOpts, _route []common.Address, _amountIn *big.Int) (LegacyQuoterQuote, error) {
	var out []interface{}
	err := _LegacyQuoter.contract.Call(opts, &out, "findBestPathFromAmountIn", _route, _amountIn)

	if err != nil {
		return *new(LegacyQuoterQuote), err
	}

	out0 := *abi.ConvertType(out[0], new(LegacyQuoterQuote)).(*LegacyQuoterQuote)

	return out0, err

}

// FindBestPathFromAmountIn is a free data retrieval call binding the contract method 0xc1b7687b.
//
// Solidity: function findBestPathFromAmountIn(address[] _route, uint256 _amountIn) view returns((address[],address[],uint256[],uint256[],uint256[],uint256[]) quote)
func (_LegacyQuoter *LegacyQuoterSession) FindBestPathFromAmountIn(_route []common.Address, _amountIn *big.Int) (LegacyQuoterQuote, error) {
	return _LegacyQuoter.Contract.FindBestPathFromAmountIn(&_LegacyQuoter.CallOpts, _route, _amountIn)
}

// FindBestPathFromAmountIn is a free data retrieval call binding the contract method 0xc1b7687b.
//
// Solidity: function findBestPathFromAmountIn(address[] _route, uint256 _amountIn) view returns((address[],address[],uint256[],uint256[],uint256[],uint256[]) quote)
func (_LegacyQuoter *LegacyQuoterCallerSession) FindBestPathFromAmountIn(_route []common.Address, _amountIn *big.Int) (LegacyQuoterQuote, error) {
	return _LegacyQuoter.Contract.FindBestPathFromAmountIn(&_LegacyQuoter.CallOpts, _route, _amountIn)
}

// FindBestPathFromAmountOut is a free data retrieval call binding the contract method 0x7000af3f.
//
// Solidity: function findBestPathFromAmountOut(address[] _route, uint256 _amountOut) view returns((address[],address[],uint256[],uint256[],uint256[],uint256[]) quote)
func (_LegacyQuoter *LegacyQuoterCaller) FindBestPathFromAmountOut(opts *bind.CallOpts, _route []common.Address, _amountOut *big.Int) (LBQuoterQuote, error) {
	var out []interface{}
	err := _LegacyQuoter.contract.Call(opts, &out, "findBestPathFromAmountOut", _route, _amountOut)

	if err != nil {
		return *new(LBQuoterQuote), err
	}

	out0 := *abi.ConvertType(out[0], new(LBQuoterQuote)).(*LBQuoterQuote)

	return out0, err

}

// FindBestPathFromAmountOut is a free data retrieval call binding the contract method 0x7000af3f.
//
// Solidity: function findBestPathFromAmountOut(address[] _route, uint256 _amountOut) view returns((address[],address[],uint256[],uint256[],uint256[],uint256[]) quote)
func (_LegacyQuoter *LegacyQuoterSession) FindBestPathFromAmountOut(_route []common.Address, _amountOut *big.Int) (LBQuoterQuote, error) {
	return _LegacyQuoter.Contract.FindBestPathFromAmountOut(&_LegacyQuoter.CallOpts, _route, _amountOut)
}

// FindBestPathFromAmountOut is a free data retrieval call binding the contract method 0x7000af3f.
//
// Solidity: function findBestPathFromAmountOut(address[] _route, uint256 _amountOut) view returns((address[],address[],uint256[],uint256[],uint256[],uint256[]) quote)
func (_LegacyQuoter *LegacyQuoterCallerSession) FindBestPathFromAmountOut(_route []common.Address, _amountOut *big.Int) (LBQuoterQuote, error) {
	return _LegacyQuoter.Contract.FindBestPathFromAmountOut(&_LegacyQuoter.CallOpts, _route, _amountOut)
}

// RouterV2 is a free data retrieval call binding the contract method 0x502f7446.
//
// Solidity: function routerV2() view returns(address)
func (_LegacyQuoter *LegacyQuoterCaller) RouterV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LegacyQuoter.contract.Call(opts, &out, "routerV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RouterV2 is a free data retrieval call binding the contract method 0x502f7446.
//
// Solidity: function routerV2() view returns(address)
func (_LegacyQuoter *LegacyQuoterSession) RouterV2() (common.Address, error) {
	return _LegacyQuoter.Contract.RouterV2(&_LegacyQuoter.CallOpts)
}

// RouterV2 is a free data retrieval call binding the contract method 0x502f7446.
//
// Solidity: function routerV2() view returns(address)
func (_LegacyQuoter *LegacyQuoterCallerSession) RouterV2() (common.Address, error) {
	return _LegacyQuoter.Contract.RouterV2(&_LegacyQuoter.CallOpts)
}
