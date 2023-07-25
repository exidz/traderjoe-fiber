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

// ILBFactoryLBPairInformation is an auto generated low-level Go binding around an user-defined struct.
type ILegacyFactoryLBPairInformation struct {
	BinStep           uint16
	LBPair            common.Address
	CreatedByOwner    bool
	IgnoredForRouting bool
}

// LegacyfactoryMetaData contains all meta data concerning the Legacyfactory contract.
var LegacyfactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_flashLoanFee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bp\",\"type\":\"uint256\"}],\"name\":\"BinHelper__BinStepOverflows\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BinHelper__IdOverflows\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBFactory__AddressZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"}],\"name\":\"LBFactory__BinStepHasNoPreset\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lowerBound\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"binStep\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"higherBound\",\"type\":\"uint256\"}],\"name\":\"LBFactory__BinStepRequirementsBreached\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"filterPeriod\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"decayPeriod\",\"type\":\"uint16\"}],\"name\":\"LBFactory__DecreasingPeriods\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBFactory__FactoryLockIsAlreadyInTheSameState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFees\",\"type\":\"uint256\"}],\"name\":\"LBFactory__FeesAboveMax\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFees\",\"type\":\"uint256\"}],\"name\":\"LBFactory__FlashLoanFeeAboveMax\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"LBFactory__FunctionIsLockedForUsers\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"LBFactory__IdenticalAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBFactory__ImplementationNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_binStep\",\"type\":\"uint256\"}],\"name\":\"LBFactory__LBPairAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBFactory__LBPairIgnoredIsAlreadyInTheSameState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"}],\"name\":\"LBFactory__LBPairNotCreated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"LBPairImplementation\",\"type\":\"address\"}],\"name\":\"LBFactory__LBPairSafetyCheckFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"protocolShare\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"LBFactory__ProtocolShareOverflows\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"quoteAsset\",\"type\":\"address\"}],\"name\":\"LBFactory__QuoteAssetAlreadyWhitelisted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"quoteAsset\",\"type\":\"address\"}],\"name\":\"LBFactory__QuoteAssetNotWhitelisted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"reductionFactor\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"LBFactory__ReductionFactorOverflows\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"LBFactory__SameFeeRecipient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"flashLoanFee\",\"type\":\"uint256\"}],\"name\":\"LBFactory__SameFlashLoanFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"LBPairImplementation\",\"type\":\"address\"}],\"name\":\"LBFactory__SameImplementation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"y\",\"type\":\"int256\"}],\"name\":\"Math128x128__PowerUnderflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingOwnable__AddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingOwnable__NoPendingOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingOwnable__NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingOwnable__NotPendingOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingOwnable__PendingOwnerAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"SafeCast__Exceeds16Bits\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"unlocked\",\"type\":\"bool\"}],\"name\":\"FactoryLockedStatusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractILBPair\",\"name\":\"LBPair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseFactor\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"filterPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"decayPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reductionFactor\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"variableFeeControl\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolShare\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxVolatilityAccumulated\",\"type\":\"uint256\"}],\"name\":\"FeeParametersSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRecipient\",\"type\":\"address\"}],\"name\":\"FeeRecipientSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldFlashLoanFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFlashLoanFee\",\"type\":\"uint256\"}],\"name\":\"FlashLoanFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractILBPair\",\"name\":\"LBPair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"}],\"name\":\"LBPairCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractILBPair\",\"name\":\"LBPair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"ignored\",\"type\":\"bool\"}],\"name\":\"LBPairIgnoredStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldLBPairImplementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"LBPairImplementation\",\"type\":\"address\"}],\"name\":\"LBPairImplementationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"PendingOwnerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"}],\"name\":\"PresetRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseFactor\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"filterPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"decayPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reductionFactor\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"variableFeeControl\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolShare\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxVolatilityAccumulated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sampleLifetime\",\"type\":\"uint256\"}],\"name\":\"PresetSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"quoteAsset\",\"type\":\"address\"}],\"name\":\"QuoteAssetAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"quoteAsset\",\"type\":\"address\"}],\"name\":\"QuoteAssetRemoved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LBPairImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BIN_STEP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_PROTOCOL_SHARE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_BIN_STEP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_quoteAsset\",\"type\":\"address\"}],\"name\":\"addQuoteAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allLBPairs\",\"outputs\":[{\"internalType\":\"contractILBPair\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"becomeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_activeId\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"_binStep\",\"type\":\"uint16\"}],\"name\":\"createLBPair\",\"outputs\":[{\"internalType\":\"contractILBPair\",\"name\":\"_LBPair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creationUnlocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flashLoanFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILBPair\",\"name\":\"_LBPair\",\"type\":\"address\"}],\"name\":\"forceDecay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllBinSteps\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"presetsBinStep\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_tokenY\",\"type\":\"address\"}],\"name\":\"getAllLBPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"binStep\",\"type\":\"uint16\"},{\"internalType\":\"contractILBPair\",\"name\":\"LBPair\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"createdByOwner\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"ignoredForRouting\",\"type\":\"bool\"}],\"internalType\":\"structILBFactory.LBPairInformation[]\",\"name\":\"LBPairsAvailable\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_binStep\",\"type\":\"uint256\"}],\"name\":\"getLBPairInformation\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"binStep\",\"type\":\"uint16\"},{\"internalType\":\"contractILBPair\",\"name\":\"LBPair\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"createdByOwner\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"ignoredForRouting\",\"type\":\"bool\"}],\"internalType\":\"structILBFactory.LBPairInformation\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberOfLBPairs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberOfQuoteAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_binStep\",\"type\":\"uint16\"}],\"name\":\"getPreset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"filterPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decayPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reductionFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableFeeControl\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVolatilityAccumulated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sampleLifetime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getQuoteAsset\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"isQuoteAsset\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_binStep\",\"type\":\"uint16\"}],\"name\":\"removePreset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_quoteAsset\",\"type\":\"address\"}],\"name\":\"removeQuoteAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokePendingOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_locked\",\"type\":\"bool\"}],\"name\":\"setFactoryLockedState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeRecipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_tokenY\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_binStep\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_baseFactor\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_filterPeriod\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_decayPeriod\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_reductionFactor\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"_variableFeeControl\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"_protocolShare\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"_maxVolatilityAccumulated\",\"type\":\"uint24\"}],\"name\":\"setFeesParametersOnPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_flashLoanFee\",\"type\":\"uint256\"}],\"name\":\"setFlashLoanFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_tokenY\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_binStep\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_ignored\",\"type\":\"bool\"}],\"name\":\"setLBPairIgnored\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_LBPairImplementation\",\"type\":\"address\"}],\"name\":\"setLBPairImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner_\",\"type\":\"address\"}],\"name\":\"setPendingOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_binStep\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_baseFactor\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_filterPeriod\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_decayPeriod\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_reductionFactor\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"_variableFeeControl\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"_protocolShare\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"_maxVolatilityAccumulated\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"_sampleLifetime\",\"type\":\"uint16\"}],\"name\":\"setPreset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LegacyfactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use LegacyfactoryMetaData.ABI instead.
var LegacyfactoryABI = LegacyfactoryMetaData.ABI

// Legacyfactory is an auto generated Go binding around an Ethereum contract.
type Legacyfactory struct {
	LegacyfactoryCaller     // Read-only binding to the contract
	LegacyfactoryTransactor // Write-only binding to the contract
	LegacyfactoryFilterer   // Log filterer for contract events
}

// LegacyfactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type LegacyfactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyfactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LegacyfactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyfactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LegacyfactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyfactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LegacyfactorySession struct {
	Contract     *Legacyfactory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LegacyfactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LegacyfactoryCallerSession struct {
	Contract *LegacyfactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// LegacyfactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LegacyfactoryTransactorSession struct {
	Contract     *LegacyfactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// LegacyfactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type LegacyfactoryRaw struct {
	Contract *Legacyfactory // Generic contract binding to access the raw methods on
}

// LegacyfactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LegacyfactoryCallerRaw struct {
	Contract *LegacyfactoryCaller // Generic read-only contract binding to access the raw methods on
}

// LegacyfactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LegacyfactoryTransactorRaw struct {
	Contract *LegacyfactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLegacyfactory creates a new instance of Legacyfactory, bound to a specific deployed contract.
func NewLegacyfactory(address common.Address, backend bind.ContractBackend) (*Legacyfactory, error) {
	contract, err := bindLegacyfactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Legacyfactory{LegacyfactoryCaller: LegacyfactoryCaller{contract: contract}, LegacyfactoryTransactor: LegacyfactoryTransactor{contract: contract}, LegacyfactoryFilterer: LegacyfactoryFilterer{contract: contract}}, nil
}

// NewLegacyfactoryCaller creates a new read-only instance of Legacyfactory, bound to a specific deployed contract.
func NewLegacyfactoryCaller(address common.Address, caller bind.ContractCaller) (*LegacyfactoryCaller, error) {
	contract, err := bindLegacyfactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LegacyfactoryCaller{contract: contract}, nil
}

// NewLegacyfactoryTransactor creates a new write-only instance of Legacyfactory, bound to a specific deployed contract.
func NewLegacyfactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*LegacyfactoryTransactor, error) {
	contract, err := bindLegacyfactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LegacyfactoryTransactor{contract: contract}, nil
}

// NewLegacyfactoryFilterer creates a new log filterer instance of Legacyfactory, bound to a specific deployed contract.
func NewLegacyfactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*LegacyfactoryFilterer, error) {
	contract, err := bindLegacyfactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LegacyfactoryFilterer{contract: contract}, nil
}

// bindLegacyfactory binds a generic wrapper to an already deployed contract.
func bindLegacyfactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LegacyfactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Legacyfactory *LegacyfactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Legacyfactory.Contract.LegacyfactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Legacyfactory *LegacyfactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Legacyfactory.Contract.LegacyfactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Legacyfactory *LegacyfactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Legacyfactory.Contract.LegacyfactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Legacyfactory *LegacyfactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Legacyfactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Legacyfactory *LegacyfactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Legacyfactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Legacyfactory *LegacyfactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Legacyfactory.Contract.contract.Transact(opts, method, params...)
}

// LBPairImplementation is a free data retrieval call binding the contract method 0x509ceb90.
//
// Solidity: function LBPairImplementation() view returns(address)
func (_Legacyfactory *LegacyfactoryCaller) LBPairImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Legacyfactory.contract.Call(opts, &out, "LBPairImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LBPairImplementation is a free data retrieval call binding the contract method 0x509ceb90.
//
// Solidity: function LBPairImplementation() view returns(address)
func (_Legacyfactory *LegacyfactorySession) LBPairImplementation() (common.Address, error) {
	return _Legacyfactory.Contract.LBPairImplementation(&_Legacyfactory.CallOpts)
}

// LBPairImplementation is a free data retrieval call binding the contract method 0x509ceb90.
//
// Solidity: function LBPairImplementation() view returns(address)
func (_Legacyfactory *LegacyfactoryCallerSession) LBPairImplementation() (common.Address, error) {
	return _Legacyfactory.Contract.LBPairImplementation(&_Legacyfactory.CallOpts)
}

// MAXBINSTEP is a free data retrieval call binding the contract method 0x10e9ec4a.
//
// Solidity: function MAX_BIN_STEP() view returns(uint256)
func (_Legacyfactory *LegacyfactoryCaller) MAXBINSTEP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Legacyfactory.contract.Call(opts, &out, "MAX_BIN_STEP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBINSTEP is a free data retrieval call binding the contract method 0x10e9ec4a.
//
// Solidity: function MAX_BIN_STEP() view returns(uint256)
func (_Legacyfactory *LegacyfactorySession) MAXBINSTEP() (*big.Int, error) {
	return _Legacyfactory.Contract.MAXBINSTEP(&_Legacyfactory.CallOpts)
}

// MAXBINSTEP is a free data retrieval call binding the contract method 0x10e9ec4a.
//
// Solidity: function MAX_BIN_STEP() view returns(uint256)
func (_Legacyfactory *LegacyfactoryCallerSession) MAXBINSTEP() (*big.Int, error) {
	return _Legacyfactory.Contract.MAXBINSTEP(&_Legacyfactory.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_Legacyfactory *LegacyfactoryCaller) MAXFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Legacyfactory.contract.Call(opts, &out, "MAX_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_Legacyfactory *LegacyfactorySession) MAXFEE() (*big.Int, error) {
	return _Legacyfactory.Contract.MAXFEE(&_Legacyfactory.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_Legacyfactory *LegacyfactoryCallerSession) MAXFEE() (*big.Int, error) {
	return _Legacyfactory.Contract.MAXFEE(&_Legacyfactory.CallOpts)
}

// MAXPROTOCOLSHARE is a free data retrieval call binding the contract method 0xa931208f.
//
// Solidity: function MAX_PROTOCOL_SHARE() view returns(uint256)
func (_Legacyfactory *LegacyfactoryCaller) MAXPROTOCOLSHARE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Legacyfactory.contract.Call(opts, &out, "MAX_PROTOCOL_SHARE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPROTOCOLSHARE is a free data retrieval call binding the contract method 0xa931208f.
//
// Solidity: function MAX_PROTOCOL_SHARE() view returns(uint256)
func (_Legacyfactory *LegacyfactorySession) MAXPROTOCOLSHARE() (*big.Int, error) {
	return _Legacyfactory.Contract.MAXPROTOCOLSHARE(&_Legacyfactory.CallOpts)
}

// MAXPROTOCOLSHARE is a free data retrieval call binding the contract method 0xa931208f.
//
// Solidity: function MAX_PROTOCOL_SHARE() view returns(uint256)
func (_Legacyfactory *LegacyfactoryCallerSession) MAXPROTOCOLSHARE() (*big.Int, error) {
	return _Legacyfactory.Contract.MAXPROTOCOLSHARE(&_Legacyfactory.CallOpts)
}

// MINBINSTEP is a free data retrieval call binding the contract method 0x7df880e3.
//
// Solidity: function MIN_BIN_STEP() view returns(uint256)
func (_Legacyfactory *LegacyfactoryCaller) MINBINSTEP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Legacyfactory.contract.Call(opts, &out, "MIN_BIN_STEP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBINSTEP is a free data retrieval call binding the contract method 0x7df880e3.
//
// Solidity: function MIN_BIN_STEP() view returns(uint256)
func (_Legacyfactory *LegacyfactorySession) MINBINSTEP() (*big.Int, error) {
	return _Legacyfactory.Contract.MINBINSTEP(&_Legacyfactory.CallOpts)
}

// MINBINSTEP is a free data retrieval call binding the contract method 0x7df880e3.
//
// Solidity: function MIN_BIN_STEP() view returns(uint256)
func (_Legacyfactory *LegacyfactoryCallerSession) MINBINSTEP() (*big.Int, error) {
	return _Legacyfactory.Contract.MINBINSTEP(&_Legacyfactory.CallOpts)
}

// AllLBPairs is a free data retrieval call binding the contract method 0x72e47b8c.
//
// Solidity: function allLBPairs(uint256 ) view returns(address)
func (_Legacyfactory *LegacyfactoryCaller) AllLBPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Legacyfactory.contract.Call(opts, &out, "allLBPairs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllLBPairs is a free data retrieval call binding the contract method 0x72e47b8c.
//
// Solidity: function allLBPairs(uint256 ) view returns(address)
func (_Legacyfactory *LegacyfactorySession) AllLBPairs(arg0 *big.Int) (common.Address, error) {
	return _Legacyfactory.Contract.AllLBPairs(&_Legacyfactory.CallOpts, arg0)
}

// AllLBPairs is a free data retrieval call binding the contract method 0x72e47b8c.
//
// Solidity: function allLBPairs(uint256 ) view returns(address)
func (_Legacyfactory *LegacyfactoryCallerSession) AllLBPairs(arg0 *big.Int) (common.Address, error) {
	return _Legacyfactory.Contract.AllLBPairs(&_Legacyfactory.CallOpts, arg0)
}

// CreationUnlocked is a free data retrieval call binding the contract method 0x5c779d6d.
//
// Solidity: function creationUnlocked() view returns(bool)
func (_Legacyfactory *LegacyfactoryCaller) CreationUnlocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Legacyfactory.contract.Call(opts, &out, "creationUnlocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CreationUnlocked is a free data retrieval call binding the contract method 0x5c779d6d.
//
// Solidity: function creationUnlocked() view returns(bool)
func (_Legacyfactory *LegacyfactorySession) CreationUnlocked() (bool, error) {
	return _Legacyfactory.Contract.CreationUnlocked(&_Legacyfactory.CallOpts)
}

// CreationUnlocked is a free data retrieval call binding the contract method 0x5c779d6d.
//
// Solidity: function creationUnlocked() view returns(bool)
func (_Legacyfactory *LegacyfactoryCallerSession) CreationUnlocked() (bool, error) {
	return _Legacyfactory.Contract.CreationUnlocked(&_Legacyfactory.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Legacyfactory *LegacyfactoryCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Legacyfactory.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Legacyfactory *LegacyfactorySession) FeeRecipient() (common.Address, error) {
	return _Legacyfactory.Contract.FeeRecipient(&_Legacyfactory.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Legacyfactory *LegacyfactoryCallerSession) FeeRecipient() (common.Address, error) {
	return _Legacyfactory.Contract.FeeRecipient(&_Legacyfactory.CallOpts)
}

// FlashLoanFee is a free data retrieval call binding the contract method 0x4847cdc8.
//
// Solidity: function flashLoanFee() view returns(uint256)
func (_Legacyfactory *LegacyfactoryCaller) FlashLoanFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Legacyfactory.contract.Call(opts, &out, "flashLoanFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlashLoanFee is a free data retrieval call binding the contract method 0x4847cdc8.
//
// Solidity: function flashLoanFee() view returns(uint256)
func (_Legacyfactory *LegacyfactorySession) FlashLoanFee() (*big.Int, error) {
	return _Legacyfactory.Contract.FlashLoanFee(&_Legacyfactory.CallOpts)
}

// FlashLoanFee is a free data retrieval call binding the contract method 0x4847cdc8.
//
// Solidity: function flashLoanFee() view returns(uint256)
func (_Legacyfactory *LegacyfactoryCallerSession) FlashLoanFee() (*big.Int, error) {
	return _Legacyfactory.Contract.FlashLoanFee(&_Legacyfactory.CallOpts)
}

// GetAllBinSteps is a free data retrieval call binding the contract method 0x5b35875c.
//
// Solidity: function getAllBinSteps() view returns(uint256[] presetsBinStep)
func (_Legacyfactory *LegacyfactoryCaller) GetAllBinSteps(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Legacyfactory.contract.Call(opts, &out, "getAllBinSteps")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAllBinSteps is a free data retrieval call binding the contract method 0x5b35875c.
//
// Solidity: function getAllBinSteps() view returns(uint256[] presetsBinStep)
func (_Legacyfactory *LegacyfactorySession) GetAllBinSteps() ([]*big.Int, error) {
	return _Legacyfactory.Contract.GetAllBinSteps(&_Legacyfactory.CallOpts)
}

// GetAllBinSteps is a free data retrieval call binding the contract method 0x5b35875c.
//
// Solidity: function getAllBinSteps() view returns(uint256[] presetsBinStep)
func (_Legacyfactory *LegacyfactoryCallerSession) GetAllBinSteps() ([]*big.Int, error) {
	return _Legacyfactory.Contract.GetAllBinSteps(&_Legacyfactory.CallOpts)
}

// GetAllLBPairs is a free data retrieval call binding the contract method 0x6622e0d7.
//
// Solidity: function getAllLBPairs(address _tokenX, address _tokenY) view returns((uint16,address,bool,bool)[] LBPairsAvailable)
func (_Legacyfactory *LegacyfactoryCaller) GetAllLBPairs(opts *bind.CallOpts, _tokenX common.Address, _tokenY common.Address) ([]ILBFactoryLBPairInformation, error) {
	var out []interface{}
	err := _Legacyfactory.contract.Call(opts, &out, "getAllLBPairs", _tokenX, _tokenY)

	if err != nil {
		return *new([]ILBFactoryLBPairInformation), err
	}

	out0 := *abi.ConvertType(out[0], new([]ILBFactoryLBPairInformation)).(*[]ILBFactoryLBPairInformation)

	return out0, err

}

// GetAllLBPairs is a free data retrieval call binding the contract method 0x6622e0d7.
//
// Solidity: function getAllLBPairs(address _tokenX, address _tokenY) view returns((uint16,address,bool,bool)[] LBPairsAvailable)
func (_Legacyfactory *LegacyfactorySession) GetAllLBPairs(_tokenX common.Address, _tokenY common.Address) ([]ILBFactoryLBPairInformation, error) {
	return _Legacyfactory.Contract.GetAllLBPairs(&_Legacyfactory.CallOpts, _tokenX, _tokenY)
}

// GetAllLBPairs is a free data retrieval call binding the contract method 0x6622e0d7.
//
// Solidity: function getAllLBPairs(address _tokenX, address _tokenY) view returns((uint16,address,bool,bool)[] LBPairsAvailable)
func (_Legacyfactory *LegacyfactoryCallerSession) GetAllLBPairs(_tokenX common.Address, _tokenY common.Address) ([]ILBFactoryLBPairInformation, error) {
	return _Legacyfactory.Contract.GetAllLBPairs(&_Legacyfactory.CallOpts, _tokenX, _tokenY)
}

// GetLBPairInformation is a free data retrieval call binding the contract method 0x704037bd.
//
// Solidity: function getLBPairInformation(address _tokenA, address _tokenB, uint256 _binStep) view returns((uint16,address,bool,bool))
func (_Legacyfactory *LegacyfactoryCaller) GetLBPairInformation(opts *bind.CallOpts, _tokenA common.Address, _tokenB common.Address, _binStep *big.Int) (ILBFactoryLBPairInformation, error) {
	var out []interface{}
	err := _Legacyfactory.contract.Call(opts, &out, "getLBPairInformation", _tokenA, _tokenB, _binStep)

	if err != nil {
		return *new(ILBFactoryLBPairInformation), err
	}

	out0 := *abi.ConvertType(out[0], new(ILBFactoryLBPairInformation)).(*ILBFactoryLBPairInformation)

	return out0, err

}

// GetLBPairInformation is a free data retrieval call binding the contract method 0x704037bd.
//
// Solidity: function getLBPairInformation(address _tokenA, address _tokenB, uint256 _binStep) view returns((uint16,address,bool,bool))
func (_Legacyfactory *LegacyfactorySession) GetLBPairInformation(_tokenA common.Address, _tokenB common.Address, _binStep *big.Int) (ILBFactoryLBPairInformation, error) {
	return _Legacyfactory.Contract.GetLBPairInformation(&_Legacyfactory.CallOpts, _tokenA, _tokenB, _binStep)
}

// GetLBPairInformation is a free data retrieval call binding the contract method 0x704037bd.
//
// Solidity: function getLBPairInformation(address _tokenA, address _tokenB, uint256 _binStep) view returns((uint16,address,bool,bool))
func (_Legacyfactory *LegacyfactoryCallerSession) GetLBPairInformation(_tokenA common.Address, _tokenB common.Address, _binStep *big.Int) (ILBFactoryLBPairInformation, error) {
	return _Legacyfactory.Contract.GetLBPairInformation(&_Legacyfactory.CallOpts, _tokenA, _tokenB, _binStep)
}

// GetNumberOfLBPairs is a free data retrieval call binding the contract method 0x4e937c3a.
//
// Solidity: function getNumberOfLBPairs() view returns(uint256)
func (_Legacyfactory *LegacyfactoryCaller) GetNumberOfLBPairs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Legacyfactory.contract.Call(opts, &out, "getNumberOfLBPairs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberOfLBPairs is a free data retrieval call binding the contract method 0x4e937c3a.
//
// Solidity: function getNumberOfLBPairs() view returns(uint256)
func (_Legacyfactory *LegacyfactorySession) GetNumberOfLBPairs() (*big.Int, error) {
	return _Legacyfactory.Contract.GetNumberOfLBPairs(&_Legacyfactory.CallOpts)
}

// GetNumberOfLBPairs is a free data retrieval call binding the contract method 0x4e937c3a.
//
// Solidity: function getNumberOfLBPairs() view returns(uint256)
func (_Legacyfactory *LegacyfactoryCallerSession) GetNumberOfLBPairs() (*big.Int, error) {
	return _Legacyfactory.Contract.GetNumberOfLBPairs(&_Legacyfactory.CallOpts)
}

// GetNumberOfQuoteAssets is a free data retrieval call binding the contract method 0x80c5061e.
//
// Solidity: function getNumberOfQuoteAssets() view returns(uint256)
func (_Legacyfactory *LegacyfactoryCaller) GetNumberOfQuoteAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Legacyfactory.contract.Call(opts, &out, "getNumberOfQuoteAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberOfQuoteAssets is a free data retrieval call binding the contract method 0x80c5061e.
//
// Solidity: function getNumberOfQuoteAssets() view returns(uint256)
func (_Legacyfactory *LegacyfactorySession) GetNumberOfQuoteAssets() (*big.Int, error) {
	return _Legacyfactory.Contract.GetNumberOfQuoteAssets(&_Legacyfactory.CallOpts)
}

// GetNumberOfQuoteAssets is a free data retrieval call binding the contract method 0x80c5061e.
//
// Solidity: function getNumberOfQuoteAssets() view returns(uint256)
func (_Legacyfactory *LegacyfactoryCallerSession) GetNumberOfQuoteAssets() (*big.Int, error) {
	return _Legacyfactory.Contract.GetNumberOfQuoteAssets(&_Legacyfactory.CallOpts)
}

// GetPreset is a free data retrieval call binding the contract method 0x935ea51b.
//
// Solidity: function getPreset(uint16 _binStep) view returns(uint256 baseFactor, uint256 filterPeriod, uint256 decayPeriod, uint256 reductionFactor, uint256 variableFeeControl, uint256 protocolShare, uint256 maxVolatilityAccumulated, uint256 sampleLifetime)
func (_Legacyfactory *LegacyfactoryCaller) GetPreset(opts *bind.CallOpts, _binStep uint16) (struct {
	BaseFactor               *big.Int
	FilterPeriod             *big.Int
	DecayPeriod              *big.Int
	ReductionFactor          *big.Int
	VariableFeeControl       *big.Int
	ProtocolShare            *big.Int
	MaxVolatilityAccumulated *big.Int
	SampleLifetime           *big.Int
}, error) {
	var out []interface{}
	err := _Legacyfactory.contract.Call(opts, &out, "getPreset", _binStep)

	outstruct := new(struct {
		BaseFactor               *big.Int
		FilterPeriod             *big.Int
		DecayPeriod              *big.Int
		ReductionFactor          *big.Int
		VariableFeeControl       *big.Int
		ProtocolShare            *big.Int
		MaxVolatilityAccumulated *big.Int
		SampleLifetime           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BaseFactor = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FilterPeriod = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DecayPeriod = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ReductionFactor = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VariableFeeControl = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.ProtocolShare = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.MaxVolatilityAccumulated = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.SampleLifetime = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPreset is a free data retrieval call binding the contract method 0x935ea51b.
//
// Solidity: function getPreset(uint16 _binStep) view returns(uint256 baseFactor, uint256 filterPeriod, uint256 decayPeriod, uint256 reductionFactor, uint256 variableFeeControl, uint256 protocolShare, uint256 maxVolatilityAccumulated, uint256 sampleLifetime)
func (_Legacyfactory *LegacyfactorySession) GetPreset(_binStep uint16) (struct {
	BaseFactor               *big.Int
	FilterPeriod             *big.Int
	DecayPeriod              *big.Int
	ReductionFactor          *big.Int
	VariableFeeControl       *big.Int
	ProtocolShare            *big.Int
	MaxVolatilityAccumulated *big.Int
	SampleLifetime           *big.Int
}, error) {
	return _Legacyfactory.Contract.GetPreset(&_Legacyfactory.CallOpts, _binStep)
}

// GetPreset is a free data retrieval call binding the contract method 0x935ea51b.
//
// Solidity: function getPreset(uint16 _binStep) view returns(uint256 baseFactor, uint256 filterPeriod, uint256 decayPeriod, uint256 reductionFactor, uint256 variableFeeControl, uint256 protocolShare, uint256 maxVolatilityAccumulated, uint256 sampleLifetime)
func (_Legacyfactory *LegacyfactoryCallerSession) GetPreset(_binStep uint16) (struct {
	BaseFactor               *big.Int
	FilterPeriod             *big.Int
	DecayPeriod              *big.Int
	ReductionFactor          *big.Int
	VariableFeeControl       *big.Int
	ProtocolShare            *big.Int
	MaxVolatilityAccumulated *big.Int
	SampleLifetime           *big.Int
}, error) {
	return _Legacyfactory.Contract.GetPreset(&_Legacyfactory.CallOpts, _binStep)
}

// GetQuoteAsset is a free data retrieval call binding the contract method 0xf89a4cd5.
//
// Solidity: function getQuoteAsset(uint256 _index) view returns(address)
func (_Legacyfactory *LegacyfactoryCaller) GetQuoteAsset(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Legacyfactory.contract.Call(opts, &out, "getQuoteAsset", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetQuoteAsset is a free data retrieval call binding the contract method 0xf89a4cd5.
//
// Solidity: function getQuoteAsset(uint256 _index) view returns(address)
func (_Legacyfactory *LegacyfactorySession) GetQuoteAsset(_index *big.Int) (common.Address, error) {
	return _Legacyfactory.Contract.GetQuoteAsset(&_Legacyfactory.CallOpts, _index)
}

// GetQuoteAsset is a free data retrieval call binding the contract method 0xf89a4cd5.
//
// Solidity: function getQuoteAsset(uint256 _index) view returns(address)
func (_Legacyfactory *LegacyfactoryCallerSession) GetQuoteAsset(_index *big.Int) (common.Address, error) {
	return _Legacyfactory.Contract.GetQuoteAsset(&_Legacyfactory.CallOpts, _index)
}

// IsQuoteAsset is a free data retrieval call binding the contract method 0x27721842.
//
// Solidity: function isQuoteAsset(address _token) view returns(bool)
func (_Legacyfactory *LegacyfactoryCaller) IsQuoteAsset(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _Legacyfactory.contract.Call(opts, &out, "isQuoteAsset", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsQuoteAsset is a free data retrieval call binding the contract method 0x27721842.
//
// Solidity: function isQuoteAsset(address _token) view returns(bool)
func (_Legacyfactory *LegacyfactorySession) IsQuoteAsset(_token common.Address) (bool, error) {
	return _Legacyfactory.Contract.IsQuoteAsset(&_Legacyfactory.CallOpts, _token)
}

// IsQuoteAsset is a free data retrieval call binding the contract method 0x27721842.
//
// Solidity: function isQuoteAsset(address _token) view returns(bool)
func (_Legacyfactory *LegacyfactoryCallerSession) IsQuoteAsset(_token common.Address) (bool, error) {
	return _Legacyfactory.Contract.IsQuoteAsset(&_Legacyfactory.CallOpts, _token)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Legacyfactory *LegacyfactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Legacyfactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Legacyfactory *LegacyfactorySession) Owner() (common.Address, error) {
	return _Legacyfactory.Contract.Owner(&_Legacyfactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Legacyfactory *LegacyfactoryCallerSession) Owner() (common.Address, error) {
	return _Legacyfactory.Contract.Owner(&_Legacyfactory.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Legacyfactory *LegacyfactoryCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Legacyfactory.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Legacyfactory *LegacyfactorySession) PendingOwner() (common.Address, error) {
	return _Legacyfactory.Contract.PendingOwner(&_Legacyfactory.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Legacyfactory *LegacyfactoryCallerSession) PendingOwner() (common.Address, error) {
	return _Legacyfactory.Contract.PendingOwner(&_Legacyfactory.CallOpts)
}

// AddQuoteAsset is a paid mutator transaction binding the contract method 0x5a440923.
//
// Solidity: function addQuoteAsset(address _quoteAsset) returns()
func (_Legacyfactory *LegacyfactoryTransactor) AddQuoteAsset(opts *bind.TransactOpts, _quoteAsset common.Address) (*types.Transaction, error) {
	return _Legacyfactory.contract.Transact(opts, "addQuoteAsset", _quoteAsset)
}

// AddQuoteAsset is a paid mutator transaction binding the contract method 0x5a440923.
//
// Solidity: function addQuoteAsset(address _quoteAsset) returns()
func (_Legacyfactory *LegacyfactorySession) AddQuoteAsset(_quoteAsset common.Address) (*types.Transaction, error) {
	return _Legacyfactory.Contract.AddQuoteAsset(&_Legacyfactory.TransactOpts, _quoteAsset)
}

// AddQuoteAsset is a paid mutator transaction binding the contract method 0x5a440923.
//
// Solidity: function addQuoteAsset(address _quoteAsset) returns()
func (_Legacyfactory *LegacyfactoryTransactorSession) AddQuoteAsset(_quoteAsset common.Address) (*types.Transaction, error) {
	return _Legacyfactory.Contract.AddQuoteAsset(&_Legacyfactory.TransactOpts, _quoteAsset)
}

// BecomeOwner is a paid mutator transaction binding the contract method 0xf9dca989.
//
// Solidity: function becomeOwner() returns()
func (_Legacyfactory *LegacyfactoryTransactor) BecomeOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Legacyfactory.contract.Transact(opts, "becomeOwner")
}

// BecomeOwner is a paid mutator transaction binding the contract method 0xf9dca989.
//
// Solidity: function becomeOwner() returns()
func (_Legacyfactory *LegacyfactorySession) BecomeOwner() (*types.Transaction, error) {
	return _Legacyfactory.Contract.BecomeOwner(&_Legacyfactory.TransactOpts)
}

// BecomeOwner is a paid mutator transaction binding the contract method 0xf9dca989.
//
// Solidity: function becomeOwner() returns()
func (_Legacyfactory *LegacyfactoryTransactorSession) BecomeOwner() (*types.Transaction, error) {
	return _Legacyfactory.Contract.BecomeOwner(&_Legacyfactory.TransactOpts)
}

// CreateLBPair is a paid mutator transaction binding the contract method 0x659ac74b.
//
// Solidity: function createLBPair(address _tokenX, address _tokenY, uint24 _activeId, uint16 _binStep) returns(address _LBPair)
func (_Legacyfactory *LegacyfactoryTransactor) CreateLBPair(opts *bind.TransactOpts, _tokenX common.Address, _tokenY common.Address, _activeId *big.Int, _binStep uint16) (*types.Transaction, error) {
	return _Legacyfactory.contract.Transact(opts, "createLBPair", _tokenX, _tokenY, _activeId, _binStep)
}

// CreateLBPair is a paid mutator transaction binding the contract method 0x659ac74b.
//
// Solidity: function createLBPair(address _tokenX, address _tokenY, uint24 _activeId, uint16 _binStep) returns(address _LBPair)
func (_Legacyfactory *LegacyfactorySession) CreateLBPair(_tokenX common.Address, _tokenY common.Address, _activeId *big.Int, _binStep uint16) (*types.Transaction, error) {
	return _Legacyfactory.Contract.CreateLBPair(&_Legacyfactory.TransactOpts, _tokenX, _tokenY, _activeId, _binStep)
}

// CreateLBPair is a paid mutator transaction binding the contract method 0x659ac74b.
//
// Solidity: function createLBPair(address _tokenX, address _tokenY, uint24 _activeId, uint16 _binStep) returns(address _LBPair)
func (_Legacyfactory *LegacyfactoryTransactorSession) CreateLBPair(_tokenX common.Address, _tokenY common.Address, _activeId *big.Int, _binStep uint16) (*types.Transaction, error) {
	return _Legacyfactory.Contract.CreateLBPair(&_Legacyfactory.TransactOpts, _tokenX, _tokenY, _activeId, _binStep)
}

// ForceDecay is a paid mutator transaction binding the contract method 0x3c78a941.
//
// Solidity: function forceDecay(address _LBPair) returns()
func (_Legacyfactory *LegacyfactoryTransactor) ForceDecay(opts *bind.TransactOpts, _LBPair common.Address) (*types.Transaction, error) {
	return _Legacyfactory.contract.Transact(opts, "forceDecay", _LBPair)
}

// ForceDecay is a paid mutator transaction binding the contract method 0x3c78a941.
//
// Solidity: function forceDecay(address _LBPair) returns()
func (_Legacyfactory *LegacyfactorySession) ForceDecay(_LBPair common.Address) (*types.Transaction, error) {
	return _Legacyfactory.Contract.ForceDecay(&_Legacyfactory.TransactOpts, _LBPair)
}

// ForceDecay is a paid mutator transaction binding the contract method 0x3c78a941.
//
// Solidity: function forceDecay(address _LBPair) returns()
func (_Legacyfactory *LegacyfactoryTransactorSession) ForceDecay(_LBPair common.Address) (*types.Transaction, error) {
	return _Legacyfactory.Contract.ForceDecay(&_Legacyfactory.TransactOpts, _LBPair)
}

// RemovePreset is a paid mutator transaction binding the contract method 0xe203a31f.
//
// Solidity: function removePreset(uint16 _binStep) returns()
func (_Legacyfactory *LegacyfactoryTransactor) RemovePreset(opts *bind.TransactOpts, _binStep uint16) (*types.Transaction, error) {
	return _Legacyfactory.contract.Transact(opts, "removePreset", _binStep)
}

// RemovePreset is a paid mutator transaction binding the contract method 0xe203a31f.
//
// Solidity: function removePreset(uint16 _binStep) returns()
func (_Legacyfactory *LegacyfactorySession) RemovePreset(_binStep uint16) (*types.Transaction, error) {
	return _Legacyfactory.Contract.RemovePreset(&_Legacyfactory.TransactOpts, _binStep)
}

// RemovePreset is a paid mutator transaction binding the contract method 0xe203a31f.
//
// Solidity: function removePreset(uint16 _binStep) returns()
func (_Legacyfactory *LegacyfactoryTransactorSession) RemovePreset(_binStep uint16) (*types.Transaction, error) {
	return _Legacyfactory.Contract.RemovePreset(&_Legacyfactory.TransactOpts, _binStep)
}

// RemoveQuoteAsset is a paid mutator transaction binding the contract method 0xddbfd941.
//
// Solidity: function removeQuoteAsset(address _quoteAsset) returns()
func (_Legacyfactory *LegacyfactoryTransactor) RemoveQuoteAsset(opts *bind.TransactOpts, _quoteAsset common.Address) (*types.Transaction, error) {
	return _Legacyfactory.contract.Transact(opts, "removeQuoteAsset", _quoteAsset)
}

// RemoveQuoteAsset is a paid mutator transaction binding the contract method 0xddbfd941.
//
// Solidity: function removeQuoteAsset(address _quoteAsset) returns()
func (_Legacyfactory *LegacyfactorySession) RemoveQuoteAsset(_quoteAsset common.Address) (*types.Transaction, error) {
	return _Legacyfactory.Contract.RemoveQuoteAsset(&_Legacyfactory.TransactOpts, _quoteAsset)
}

// RemoveQuoteAsset is a paid mutator transaction binding the contract method 0xddbfd941.
//
// Solidity: function removeQuoteAsset(address _quoteAsset) returns()
func (_Legacyfactory *LegacyfactoryTransactorSession) RemoveQuoteAsset(_quoteAsset common.Address) (*types.Transaction, error) {
	return _Legacyfactory.Contract.RemoveQuoteAsset(&_Legacyfactory.TransactOpts, _quoteAsset)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Legacyfactory *LegacyfactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Legacyfactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Legacyfactory *LegacyfactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _Legacyfactory.Contract.RenounceOwnership(&_Legacyfactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Legacyfactory *LegacyfactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Legacyfactory.Contract.RenounceOwnership(&_Legacyfactory.TransactOpts)
}

// RevokePendingOwner is a paid mutator transaction binding the contract method 0x67ab8a4e.
//
// Solidity: function revokePendingOwner() returns()
func (_Legacyfactory *LegacyfactoryTransactor) RevokePendingOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Legacyfactory.contract.Transact(opts, "revokePendingOwner")
}

// RevokePendingOwner is a paid mutator transaction binding the contract method 0x67ab8a4e.
//
// Solidity: function revokePendingOwner() returns()
func (_Legacyfactory *LegacyfactorySession) RevokePendingOwner() (*types.Transaction, error) {
	return _Legacyfactory.Contract.RevokePendingOwner(&_Legacyfactory.TransactOpts)
}

// RevokePendingOwner is a paid mutator transaction binding the contract method 0x67ab8a4e.
//
// Solidity: function revokePendingOwner() returns()
func (_Legacyfactory *LegacyfactoryTransactorSession) RevokePendingOwner() (*types.Transaction, error) {
	return _Legacyfactory.Contract.RevokePendingOwner(&_Legacyfactory.TransactOpts)
}

// SetFactoryLockedState is a paid mutator transaction binding the contract method 0x22f3fe14.
//
// Solidity: function setFactoryLockedState(bool _locked) returns()
func (_Legacyfactory *LegacyfactoryTransactor) SetFactoryLockedState(opts *bind.TransactOpts, _locked bool) (*types.Transaction, error) {
	return _Legacyfactory.contract.Transact(opts, "setFactoryLockedState", _locked)
}

// SetFactoryLockedState is a paid mutator transaction binding the contract method 0x22f3fe14.
//
// Solidity: function setFactoryLockedState(bool _locked) returns()
func (_Legacyfactory *LegacyfactorySession) SetFactoryLockedState(_locked bool) (*types.Transaction, error) {
	return _Legacyfactory.Contract.SetFactoryLockedState(&_Legacyfactory.TransactOpts, _locked)
}

// SetFactoryLockedState is a paid mutator transaction binding the contract method 0x22f3fe14.
//
// Solidity: function setFactoryLockedState(bool _locked) returns()
func (_Legacyfactory *LegacyfactoryTransactorSession) SetFactoryLockedState(_locked bool) (*types.Transaction, error) {
	return _Legacyfactory.Contract.SetFactoryLockedState(&_Legacyfactory.TransactOpts, _locked)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _feeRecipient) returns()
func (_Legacyfactory *LegacyfactoryTransactor) SetFeeRecipient(opts *bind.TransactOpts, _feeRecipient common.Address) (*types.Transaction, error) {
	return _Legacyfactory.contract.Transact(opts, "setFeeRecipient", _feeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _feeRecipient) returns()
func (_Legacyfactory *LegacyfactorySession) SetFeeRecipient(_feeRecipient common.Address) (*types.Transaction, error) {
	return _Legacyfactory.Contract.SetFeeRecipient(&_Legacyfactory.TransactOpts, _feeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _feeRecipient) returns()
func (_Legacyfactory *LegacyfactoryTransactorSession) SetFeeRecipient(_feeRecipient common.Address) (*types.Transaction, error) {
	return _Legacyfactory.Contract.SetFeeRecipient(&_Legacyfactory.TransactOpts, _feeRecipient)
}

// SetFeesParametersOnPair is a paid mutator transaction binding the contract method 0x093ff769.
//
// Solidity: function setFeesParametersOnPair(address _tokenX, address _tokenY, uint16 _binStep, uint16 _baseFactor, uint16 _filterPeriod, uint16 _decayPeriod, uint16 _reductionFactor, uint24 _variableFeeControl, uint16 _protocolShare, uint24 _maxVolatilityAccumulated) returns()
func (_Legacyfactory *LegacyfactoryTransactor) SetFeesParametersOnPair(opts *bind.TransactOpts, _tokenX common.Address, _tokenY common.Address, _binStep uint16, _baseFactor uint16, _filterPeriod uint16, _decayPeriod uint16, _reductionFactor uint16, _variableFeeControl *big.Int, _protocolShare uint16, _maxVolatilityAccumulated *big.Int) (*types.Transaction, error) {
	return _Legacyfactory.contract.Transact(opts, "setFeesParametersOnPair", _tokenX, _tokenY, _binStep, _baseFactor, _filterPeriod, _decayPeriod, _reductionFactor, _variableFeeControl, _protocolShare, _maxVolatilityAccumulated)
}

// SetFeesParametersOnPair is a paid mutator transaction binding the contract method 0x093ff769.
//
// Solidity: function setFeesParametersOnPair(address _tokenX, address _tokenY, uint16 _binStep, uint16 _baseFactor, uint16 _filterPeriod, uint16 _decayPeriod, uint16 _reductionFactor, uint24 _variableFeeControl, uint16 _protocolShare, uint24 _maxVolatilityAccumulated) returns()
func (_Legacyfactory *LegacyfactorySession) SetFeesParametersOnPair(_tokenX common.Address, _tokenY common.Address, _binStep uint16, _baseFactor uint16, _filterPeriod uint16, _decayPeriod uint16, _reductionFactor uint16, _variableFeeControl *big.Int, _protocolShare uint16, _maxVolatilityAccumulated *big.Int) (*types.Transaction, error) {
	return _Legacyfactory.Contract.SetFeesParametersOnPair(&_Legacyfactory.TransactOpts, _tokenX, _tokenY, _binStep, _baseFactor, _filterPeriod, _decayPeriod, _reductionFactor, _variableFeeControl, _protocolShare, _maxVolatilityAccumulated)
}

// SetFeesParametersOnPair is a paid mutator transaction binding the contract method 0x093ff769.
//
// Solidity: function setFeesParametersOnPair(address _tokenX, address _tokenY, uint16 _binStep, uint16 _baseFactor, uint16 _filterPeriod, uint16 _decayPeriod, uint16 _reductionFactor, uint24 _variableFeeControl, uint16 _protocolShare, uint24 _maxVolatilityAccumulated) returns()
func (_Legacyfactory *LegacyfactoryTransactorSession) SetFeesParametersOnPair(_tokenX common.Address, _tokenY common.Address, _binStep uint16, _baseFactor uint16, _filterPeriod uint16, _decayPeriod uint16, _reductionFactor uint16, _variableFeeControl *big.Int, _protocolShare uint16, _maxVolatilityAccumulated *big.Int) (*types.Transaction, error) {
	return _Legacyfactory.Contract.SetFeesParametersOnPair(&_Legacyfactory.TransactOpts, _tokenX, _tokenY, _binStep, _baseFactor, _filterPeriod, _decayPeriod, _reductionFactor, _variableFeeControl, _protocolShare, _maxVolatilityAccumulated)
}

// SetFlashLoanFee is a paid mutator transaction binding the contract method 0xe92d0d5d.
//
// Solidity: function setFlashLoanFee(uint256 _flashLoanFee) returns()
func (_Legacyfactory *LegacyfactoryTransactor) SetFlashLoanFee(opts *bind.TransactOpts, _flashLoanFee *big.Int) (*types.Transaction, error) {
	return _Legacyfactory.contract.Transact(opts, "setFlashLoanFee", _flashLoanFee)
}

// SetFlashLoanFee is a paid mutator transaction binding the contract method 0xe92d0d5d.
//
// Solidity: function setFlashLoanFee(uint256 _flashLoanFee) returns()
func (_Legacyfactory *LegacyfactorySession) SetFlashLoanFee(_flashLoanFee *big.Int) (*types.Transaction, error) {
	return _Legacyfactory.Contract.SetFlashLoanFee(&_Legacyfactory.TransactOpts, _flashLoanFee)
}

// SetFlashLoanFee is a paid mutator transaction binding the contract method 0xe92d0d5d.
//
// Solidity: function setFlashLoanFee(uint256 _flashLoanFee) returns()
func (_Legacyfactory *LegacyfactoryTransactorSession) SetFlashLoanFee(_flashLoanFee *big.Int) (*types.Transaction, error) {
	return _Legacyfactory.Contract.SetFlashLoanFee(&_Legacyfactory.TransactOpts, _flashLoanFee)
}

// SetLBPairIgnored is a paid mutator transaction binding the contract method 0x200aa7e3.
//
// Solidity: function setLBPairIgnored(address _tokenX, address _tokenY, uint256 _binStep, bool _ignored) returns()
func (_Legacyfactory *LegacyfactoryTransactor) SetLBPairIgnored(opts *bind.TransactOpts, _tokenX common.Address, _tokenY common.Address, _binStep *big.Int, _ignored bool) (*types.Transaction, error) {
	return _Legacyfactory.contract.Transact(opts, "setLBPairIgnored", _tokenX, _tokenY, _binStep, _ignored)
}

// SetLBPairIgnored is a paid mutator transaction binding the contract method 0x200aa7e3.
//
// Solidity: function setLBPairIgnored(address _tokenX, address _tokenY, uint256 _binStep, bool _ignored) returns()
func (_Legacyfactory *LegacyfactorySession) SetLBPairIgnored(_tokenX common.Address, _tokenY common.Address, _binStep *big.Int, _ignored bool) (*types.Transaction, error) {
	return _Legacyfactory.Contract.SetLBPairIgnored(&_Legacyfactory.TransactOpts, _tokenX, _tokenY, _binStep, _ignored)
}

// SetLBPairIgnored is a paid mutator transaction binding the contract method 0x200aa7e3.
//
// Solidity: function setLBPairIgnored(address _tokenX, address _tokenY, uint256 _binStep, bool _ignored) returns()
func (_Legacyfactory *LegacyfactoryTransactorSession) SetLBPairIgnored(_tokenX common.Address, _tokenY common.Address, _binStep *big.Int, _ignored bool) (*types.Transaction, error) {
	return _Legacyfactory.Contract.SetLBPairIgnored(&_Legacyfactory.TransactOpts, _tokenX, _tokenY, _binStep, _ignored)
}

// SetLBPairImplementation is a paid mutator transaction binding the contract method 0xb0384781.
//
// Solidity: function setLBPairImplementation(address _LBPairImplementation) returns()
func (_Legacyfactory *LegacyfactoryTransactor) SetLBPairImplementation(opts *bind.TransactOpts, _LBPairImplementation common.Address) (*types.Transaction, error) {
	return _Legacyfactory.contract.Transact(opts, "setLBPairImplementation", _LBPairImplementation)
}

// SetLBPairImplementation is a paid mutator transaction binding the contract method 0xb0384781.
//
// Solidity: function setLBPairImplementation(address _LBPairImplementation) returns()
func (_Legacyfactory *LegacyfactorySession) SetLBPairImplementation(_LBPairImplementation common.Address) (*types.Transaction, error) {
	return _Legacyfactory.Contract.SetLBPairImplementation(&_Legacyfactory.TransactOpts, _LBPairImplementation)
}

// SetLBPairImplementation is a paid mutator transaction binding the contract method 0xb0384781.
//
// Solidity: function setLBPairImplementation(address _LBPairImplementation) returns()
func (_Legacyfactory *LegacyfactoryTransactorSession) SetLBPairImplementation(_LBPairImplementation common.Address) (*types.Transaction, error) {
	return _Legacyfactory.Contract.SetLBPairImplementation(&_Legacyfactory.TransactOpts, _LBPairImplementation)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address pendingOwner_) returns()
func (_Legacyfactory *LegacyfactoryTransactor) SetPendingOwner(opts *bind.TransactOpts, pendingOwner_ common.Address) (*types.Transaction, error) {
	return _Legacyfactory.contract.Transact(opts, "setPendingOwner", pendingOwner_)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address pendingOwner_) returns()
func (_Legacyfactory *LegacyfactorySession) SetPendingOwner(pendingOwner_ common.Address) (*types.Transaction, error) {
	return _Legacyfactory.Contract.SetPendingOwner(&_Legacyfactory.TransactOpts, pendingOwner_)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address pendingOwner_) returns()
func (_Legacyfactory *LegacyfactoryTransactorSession) SetPendingOwner(pendingOwner_ common.Address) (*types.Transaction, error) {
	return _Legacyfactory.Contract.SetPendingOwner(&_Legacyfactory.TransactOpts, pendingOwner_)
}

// SetPreset is a paid mutator transaction binding the contract method 0x0af97c9a.
//
// Solidity: function setPreset(uint16 _binStep, uint16 _baseFactor, uint16 _filterPeriod, uint16 _decayPeriod, uint16 _reductionFactor, uint24 _variableFeeControl, uint16 _protocolShare, uint24 _maxVolatilityAccumulated, uint16 _sampleLifetime) returns()
func (_Legacyfactory *LegacyfactoryTransactor) SetPreset(opts *bind.TransactOpts, _binStep uint16, _baseFactor uint16, _filterPeriod uint16, _decayPeriod uint16, _reductionFactor uint16, _variableFeeControl *big.Int, _protocolShare uint16, _maxVolatilityAccumulated *big.Int, _sampleLifetime uint16) (*types.Transaction, error) {
	return _Legacyfactory.contract.Transact(opts, "setPreset", _binStep, _baseFactor, _filterPeriod, _decayPeriod, _reductionFactor, _variableFeeControl, _protocolShare, _maxVolatilityAccumulated, _sampleLifetime)
}

// SetPreset is a paid mutator transaction binding the contract method 0x0af97c9a.
//
// Solidity: function setPreset(uint16 _binStep, uint16 _baseFactor, uint16 _filterPeriod, uint16 _decayPeriod, uint16 _reductionFactor, uint24 _variableFeeControl, uint16 _protocolShare, uint24 _maxVolatilityAccumulated, uint16 _sampleLifetime) returns()
func (_Legacyfactory *LegacyfactorySession) SetPreset(_binStep uint16, _baseFactor uint16, _filterPeriod uint16, _decayPeriod uint16, _reductionFactor uint16, _variableFeeControl *big.Int, _protocolShare uint16, _maxVolatilityAccumulated *big.Int, _sampleLifetime uint16) (*types.Transaction, error) {
	return _Legacyfactory.Contract.SetPreset(&_Legacyfactory.TransactOpts, _binStep, _baseFactor, _filterPeriod, _decayPeriod, _reductionFactor, _variableFeeControl, _protocolShare, _maxVolatilityAccumulated, _sampleLifetime)
}

// SetPreset is a paid mutator transaction binding the contract method 0x0af97c9a.
//
// Solidity: function setPreset(uint16 _binStep, uint16 _baseFactor, uint16 _filterPeriod, uint16 _decayPeriod, uint16 _reductionFactor, uint24 _variableFeeControl, uint16 _protocolShare, uint24 _maxVolatilityAccumulated, uint16 _sampleLifetime) returns()
func (_Legacyfactory *LegacyfactoryTransactorSession) SetPreset(_binStep uint16, _baseFactor uint16, _filterPeriod uint16, _decayPeriod uint16, _reductionFactor uint16, _variableFeeControl *big.Int, _protocolShare uint16, _maxVolatilityAccumulated *big.Int, _sampleLifetime uint16) (*types.Transaction, error) {
	return _Legacyfactory.Contract.SetPreset(&_Legacyfactory.TransactOpts, _binStep, _baseFactor, _filterPeriod, _decayPeriod, _reductionFactor, _variableFeeControl, _protocolShare, _maxVolatilityAccumulated, _sampleLifetime)
}

// LegacyfactoryFactoryLockedStatusUpdatedIterator is returned from FilterFactoryLockedStatusUpdated and is used to iterate over the raw logs and unpacked data for FactoryLockedStatusUpdated events raised by the Legacyfactory contract.
type LegacyfactoryFactoryLockedStatusUpdatedIterator struct {
	Event *LegacyfactoryFactoryLockedStatusUpdated // Event containing the contract specifics and raw log

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
func (it *LegacyfactoryFactoryLockedStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyfactoryFactoryLockedStatusUpdated)
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
		it.Event = new(LegacyfactoryFactoryLockedStatusUpdated)
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
func (it *LegacyfactoryFactoryLockedStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyfactoryFactoryLockedStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyfactoryFactoryLockedStatusUpdated represents a FactoryLockedStatusUpdated event raised by the Legacyfactory contract.
type LegacyfactoryFactoryLockedStatusUpdated struct {
	Unlocked bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFactoryLockedStatusUpdated is a free log retrieval operation binding the contract event 0xcdee7bf87b7a743b4cbe1d2d534c5248621b76f58460337e7fda92d5d23f4124.
//
// Solidity: event FactoryLockedStatusUpdated(bool unlocked)
func (_Legacyfactory *LegacyfactoryFilterer) FilterFactoryLockedStatusUpdated(opts *bind.FilterOpts) (*LegacyfactoryFactoryLockedStatusUpdatedIterator, error) {

	logs, sub, err := _Legacyfactory.contract.FilterLogs(opts, "FactoryLockedStatusUpdated")
	if err != nil {
		return nil, err
	}
	return &LegacyfactoryFactoryLockedStatusUpdatedIterator{contract: _Legacyfactory.contract, event: "FactoryLockedStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchFactoryLockedStatusUpdated is a free log subscription operation binding the contract event 0xcdee7bf87b7a743b4cbe1d2d534c5248621b76f58460337e7fda92d5d23f4124.
//
// Solidity: event FactoryLockedStatusUpdated(bool unlocked)
func (_Legacyfactory *LegacyfactoryFilterer) WatchFactoryLockedStatusUpdated(opts *bind.WatchOpts, sink chan<- *LegacyfactoryFactoryLockedStatusUpdated) (event.Subscription, error) {

	logs, sub, err := _Legacyfactory.contract.WatchLogs(opts, "FactoryLockedStatusUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyfactoryFactoryLockedStatusUpdated)
				if err := _Legacyfactory.contract.UnpackLog(event, "FactoryLockedStatusUpdated", log); err != nil {
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

// ParseFactoryLockedStatusUpdated is a log parse operation binding the contract event 0xcdee7bf87b7a743b4cbe1d2d534c5248621b76f58460337e7fda92d5d23f4124.
//
// Solidity: event FactoryLockedStatusUpdated(bool unlocked)
func (_Legacyfactory *LegacyfactoryFilterer) ParseFactoryLockedStatusUpdated(log types.Log) (*LegacyfactoryFactoryLockedStatusUpdated, error) {
	event := new(LegacyfactoryFactoryLockedStatusUpdated)
	if err := _Legacyfactory.contract.UnpackLog(event, "FactoryLockedStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacyfactoryFeeParametersSetIterator is returned from FilterFeeParametersSet and is used to iterate over the raw logs and unpacked data for FeeParametersSet events raised by the Legacyfactory contract.
type LegacyfactoryFeeParametersSetIterator struct {
	Event *LegacyfactoryFeeParametersSet // Event containing the contract specifics and raw log

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
func (it *LegacyfactoryFeeParametersSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyfactoryFeeParametersSet)
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
		it.Event = new(LegacyfactoryFeeParametersSet)
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
func (it *LegacyfactoryFeeParametersSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyfactoryFeeParametersSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyfactoryFeeParametersSet represents a FeeParametersSet event raised by the Legacyfactory contract.
type LegacyfactoryFeeParametersSet struct {
	Sender                   common.Address
	LBPair                   common.Address
	BinStep                  *big.Int
	BaseFactor               *big.Int
	FilterPeriod             *big.Int
	DecayPeriod              *big.Int
	ReductionFactor          *big.Int
	VariableFeeControl       *big.Int
	ProtocolShare            *big.Int
	MaxVolatilityAccumulated *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterFeeParametersSet is a free log retrieval operation binding the contract event 0x63a7af39b7b68b9c3f2dfe93e5f32d9faecb4c6c98733bb608f757e62f816c0d.
//
// Solidity: event FeeParametersSet(address indexed sender, address indexed LBPair, uint256 binStep, uint256 baseFactor, uint256 filterPeriod, uint256 decayPeriod, uint256 reductionFactor, uint256 variableFeeControl, uint256 protocolShare, uint256 maxVolatilityAccumulated)
func (_Legacyfactory *LegacyfactoryFilterer) FilterFeeParametersSet(opts *bind.FilterOpts, sender []common.Address, LBPair []common.Address) (*LegacyfactoryFeeParametersSetIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var LBPairRule []interface{}
	for _, LBPairItem := range LBPair {
		LBPairRule = append(LBPairRule, LBPairItem)
	}

	logs, sub, err := _Legacyfactory.contract.FilterLogs(opts, "FeeParametersSet", senderRule, LBPairRule)
	if err != nil {
		return nil, err
	}
	return &LegacyfactoryFeeParametersSetIterator{contract: _Legacyfactory.contract, event: "FeeParametersSet", logs: logs, sub: sub}, nil
}

// WatchFeeParametersSet is a free log subscription operation binding the contract event 0x63a7af39b7b68b9c3f2dfe93e5f32d9faecb4c6c98733bb608f757e62f816c0d.
//
// Solidity: event FeeParametersSet(address indexed sender, address indexed LBPair, uint256 binStep, uint256 baseFactor, uint256 filterPeriod, uint256 decayPeriod, uint256 reductionFactor, uint256 variableFeeControl, uint256 protocolShare, uint256 maxVolatilityAccumulated)
func (_Legacyfactory *LegacyfactoryFilterer) WatchFeeParametersSet(opts *bind.WatchOpts, sink chan<- *LegacyfactoryFeeParametersSet, sender []common.Address, LBPair []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var LBPairRule []interface{}
	for _, LBPairItem := range LBPair {
		LBPairRule = append(LBPairRule, LBPairItem)
	}

	logs, sub, err := _Legacyfactory.contract.WatchLogs(opts, "FeeParametersSet", senderRule, LBPairRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyfactoryFeeParametersSet)
				if err := _Legacyfactory.contract.UnpackLog(event, "FeeParametersSet", log); err != nil {
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

// ParseFeeParametersSet is a log parse operation binding the contract event 0x63a7af39b7b68b9c3f2dfe93e5f32d9faecb4c6c98733bb608f757e62f816c0d.
//
// Solidity: event FeeParametersSet(address indexed sender, address indexed LBPair, uint256 binStep, uint256 baseFactor, uint256 filterPeriod, uint256 decayPeriod, uint256 reductionFactor, uint256 variableFeeControl, uint256 protocolShare, uint256 maxVolatilityAccumulated)
func (_Legacyfactory *LegacyfactoryFilterer) ParseFeeParametersSet(log types.Log) (*LegacyfactoryFeeParametersSet, error) {
	event := new(LegacyfactoryFeeParametersSet)
	if err := _Legacyfactory.contract.UnpackLog(event, "FeeParametersSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacyfactoryFeeRecipientSetIterator is returned from FilterFeeRecipientSet and is used to iterate over the raw logs and unpacked data for FeeRecipientSet events raised by the Legacyfactory contract.
type LegacyfactoryFeeRecipientSetIterator struct {
	Event *LegacyfactoryFeeRecipientSet // Event containing the contract specifics and raw log

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
func (it *LegacyfactoryFeeRecipientSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyfactoryFeeRecipientSet)
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
		it.Event = new(LegacyfactoryFeeRecipientSet)
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
func (it *LegacyfactoryFeeRecipientSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyfactoryFeeRecipientSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyfactoryFeeRecipientSet represents a FeeRecipientSet event raised by the Legacyfactory contract.
type LegacyfactoryFeeRecipientSet struct {
	OldRecipient common.Address
	NewRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeeRecipientSet is a free log retrieval operation binding the contract event 0x15d80a013f22151bc7246e3bc132e12828cde19de98870475e3fa70840152721.
//
// Solidity: event FeeRecipientSet(address oldRecipient, address newRecipient)
func (_Legacyfactory *LegacyfactoryFilterer) FilterFeeRecipientSet(opts *bind.FilterOpts) (*LegacyfactoryFeeRecipientSetIterator, error) {

	logs, sub, err := _Legacyfactory.contract.FilterLogs(opts, "FeeRecipientSet")
	if err != nil {
		return nil, err
	}
	return &LegacyfactoryFeeRecipientSetIterator{contract: _Legacyfactory.contract, event: "FeeRecipientSet", logs: logs, sub: sub}, nil
}

// WatchFeeRecipientSet is a free log subscription operation binding the contract event 0x15d80a013f22151bc7246e3bc132e12828cde19de98870475e3fa70840152721.
//
// Solidity: event FeeRecipientSet(address oldRecipient, address newRecipient)
func (_Legacyfactory *LegacyfactoryFilterer) WatchFeeRecipientSet(opts *bind.WatchOpts, sink chan<- *LegacyfactoryFeeRecipientSet) (event.Subscription, error) {

	logs, sub, err := _Legacyfactory.contract.WatchLogs(opts, "FeeRecipientSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyfactoryFeeRecipientSet)
				if err := _Legacyfactory.contract.UnpackLog(event, "FeeRecipientSet", log); err != nil {
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

// ParseFeeRecipientSet is a log parse operation binding the contract event 0x15d80a013f22151bc7246e3bc132e12828cde19de98870475e3fa70840152721.
//
// Solidity: event FeeRecipientSet(address oldRecipient, address newRecipient)
func (_Legacyfactory *LegacyfactoryFilterer) ParseFeeRecipientSet(log types.Log) (*LegacyfactoryFeeRecipientSet, error) {
	event := new(LegacyfactoryFeeRecipientSet)
	if err := _Legacyfactory.contract.UnpackLog(event, "FeeRecipientSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacyfactoryFlashLoanFeeSetIterator is returned from FilterFlashLoanFeeSet and is used to iterate over the raw logs and unpacked data for FlashLoanFeeSet events raised by the Legacyfactory contract.
type LegacyfactoryFlashLoanFeeSetIterator struct {
	Event *LegacyfactoryFlashLoanFeeSet // Event containing the contract specifics and raw log

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
func (it *LegacyfactoryFlashLoanFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyfactoryFlashLoanFeeSet)
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
		it.Event = new(LegacyfactoryFlashLoanFeeSet)
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
func (it *LegacyfactoryFlashLoanFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyfactoryFlashLoanFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyfactoryFlashLoanFeeSet represents a FlashLoanFeeSet event raised by the Legacyfactory contract.
type LegacyfactoryFlashLoanFeeSet struct {
	OldFlashLoanFee *big.Int
	NewFlashLoanFee *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFlashLoanFeeSet is a free log retrieval operation binding the contract event 0x5c34e91c94c78b662a45d0bd4a25a4e32c584c54a45a76e4a4d43be27ba40e50.
//
// Solidity: event FlashLoanFeeSet(uint256 oldFlashLoanFee, uint256 newFlashLoanFee)
func (_Legacyfactory *LegacyfactoryFilterer) FilterFlashLoanFeeSet(opts *bind.FilterOpts) (*LegacyfactoryFlashLoanFeeSetIterator, error) {

	logs, sub, err := _Legacyfactory.contract.FilterLogs(opts, "FlashLoanFeeSet")
	if err != nil {
		return nil, err
	}
	return &LegacyfactoryFlashLoanFeeSetIterator{contract: _Legacyfactory.contract, event: "FlashLoanFeeSet", logs: logs, sub: sub}, nil
}

// WatchFlashLoanFeeSet is a free log subscription operation binding the contract event 0x5c34e91c94c78b662a45d0bd4a25a4e32c584c54a45a76e4a4d43be27ba40e50.
//
// Solidity: event FlashLoanFeeSet(uint256 oldFlashLoanFee, uint256 newFlashLoanFee)
func (_Legacyfactory *LegacyfactoryFilterer) WatchFlashLoanFeeSet(opts *bind.WatchOpts, sink chan<- *LegacyfactoryFlashLoanFeeSet) (event.Subscription, error) {

	logs, sub, err := _Legacyfactory.contract.WatchLogs(opts, "FlashLoanFeeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyfactoryFlashLoanFeeSet)
				if err := _Legacyfactory.contract.UnpackLog(event, "FlashLoanFeeSet", log); err != nil {
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

// ParseFlashLoanFeeSet is a log parse operation binding the contract event 0x5c34e91c94c78b662a45d0bd4a25a4e32c584c54a45a76e4a4d43be27ba40e50.
//
// Solidity: event FlashLoanFeeSet(uint256 oldFlashLoanFee, uint256 newFlashLoanFee)
func (_Legacyfactory *LegacyfactoryFilterer) ParseFlashLoanFeeSet(log types.Log) (*LegacyfactoryFlashLoanFeeSet, error) {
	event := new(LegacyfactoryFlashLoanFeeSet)
	if err := _Legacyfactory.contract.UnpackLog(event, "FlashLoanFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacyfactoryLBPairCreatedIterator is returned from FilterLBPairCreated and is used to iterate over the raw logs and unpacked data for LBPairCreated events raised by the Legacyfactory contract.
type LegacyfactoryLBPairCreatedIterator struct {
	Event *LegacyfactoryLBPairCreated // Event containing the contract specifics and raw log

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
func (it *LegacyfactoryLBPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyfactoryLBPairCreated)
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
		it.Event = new(LegacyfactoryLBPairCreated)
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
func (it *LegacyfactoryLBPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyfactoryLBPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyfactoryLBPairCreated represents a LBPairCreated event raised by the Legacyfactory contract.
type LegacyfactoryLBPairCreated struct {
	TokenX  common.Address
	TokenY  common.Address
	BinStep *big.Int
	LBPair  common.Address
	Pid     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLBPairCreated is a free log retrieval operation binding the contract event 0x2c8d104b27c6b7f4492017a6f5cf3803043688934ebcaa6a03540beeaf976aff.
//
// Solidity: event LBPairCreated(address indexed tokenX, address indexed tokenY, uint256 indexed binStep, address LBPair, uint256 pid)
func (_Legacyfactory *LegacyfactoryFilterer) FilterLBPairCreated(opts *bind.FilterOpts, tokenX []common.Address, tokenY []common.Address, binStep []*big.Int) (*LegacyfactoryLBPairCreatedIterator, error) {

	var tokenXRule []interface{}
	for _, tokenXItem := range tokenX {
		tokenXRule = append(tokenXRule, tokenXItem)
	}
	var tokenYRule []interface{}
	for _, tokenYItem := range tokenY {
		tokenYRule = append(tokenYRule, tokenYItem)
	}
	var binStepRule []interface{}
	for _, binStepItem := range binStep {
		binStepRule = append(binStepRule, binStepItem)
	}

	logs, sub, err := _Legacyfactory.contract.FilterLogs(opts, "LBPairCreated", tokenXRule, tokenYRule, binStepRule)
	if err != nil {
		return nil, err
	}
	return &LegacyfactoryLBPairCreatedIterator{contract: _Legacyfactory.contract, event: "LBPairCreated", logs: logs, sub: sub}, nil
}

// WatchLBPairCreated is a free log subscription operation binding the contract event 0x2c8d104b27c6b7f4492017a6f5cf3803043688934ebcaa6a03540beeaf976aff.
//
// Solidity: event LBPairCreated(address indexed tokenX, address indexed tokenY, uint256 indexed binStep, address LBPair, uint256 pid)
func (_Legacyfactory *LegacyfactoryFilterer) WatchLBPairCreated(opts *bind.WatchOpts, sink chan<- *LegacyfactoryLBPairCreated, tokenX []common.Address, tokenY []common.Address, binStep []*big.Int) (event.Subscription, error) {

	var tokenXRule []interface{}
	for _, tokenXItem := range tokenX {
		tokenXRule = append(tokenXRule, tokenXItem)
	}
	var tokenYRule []interface{}
	for _, tokenYItem := range tokenY {
		tokenYRule = append(tokenYRule, tokenYItem)
	}
	var binStepRule []interface{}
	for _, binStepItem := range binStep {
		binStepRule = append(binStepRule, binStepItem)
	}

	logs, sub, err := _Legacyfactory.contract.WatchLogs(opts, "LBPairCreated", tokenXRule, tokenYRule, binStepRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyfactoryLBPairCreated)
				if err := _Legacyfactory.contract.UnpackLog(event, "LBPairCreated", log); err != nil {
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

// ParseLBPairCreated is a log parse operation binding the contract event 0x2c8d104b27c6b7f4492017a6f5cf3803043688934ebcaa6a03540beeaf976aff.
//
// Solidity: event LBPairCreated(address indexed tokenX, address indexed tokenY, uint256 indexed binStep, address LBPair, uint256 pid)
func (_Legacyfactory *LegacyfactoryFilterer) ParseLBPairCreated(log types.Log) (*LegacyfactoryLBPairCreated, error) {
	event := new(LegacyfactoryLBPairCreated)
	if err := _Legacyfactory.contract.UnpackLog(event, "LBPairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacyfactoryLBPairIgnoredStateChangedIterator is returned from FilterLBPairIgnoredStateChanged and is used to iterate over the raw logs and unpacked data for LBPairIgnoredStateChanged events raised by the Legacyfactory contract.
type LegacyfactoryLBPairIgnoredStateChangedIterator struct {
	Event *LegacyfactoryLBPairIgnoredStateChanged // Event containing the contract specifics and raw log

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
func (it *LegacyfactoryLBPairIgnoredStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyfactoryLBPairIgnoredStateChanged)
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
		it.Event = new(LegacyfactoryLBPairIgnoredStateChanged)
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
func (it *LegacyfactoryLBPairIgnoredStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyfactoryLBPairIgnoredStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyfactoryLBPairIgnoredStateChanged represents a LBPairIgnoredStateChanged event raised by the Legacyfactory contract.
type LegacyfactoryLBPairIgnoredStateChanged struct {
	LBPair  common.Address
	Ignored bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLBPairIgnoredStateChanged is a free log retrieval operation binding the contract event 0x44cf35361c9ff3c8c1397ec6410d5495cc481feaef35c9af11da1a637107de4f.
//
// Solidity: event LBPairIgnoredStateChanged(address indexed LBPair, bool ignored)
func (_Legacyfactory *LegacyfactoryFilterer) FilterLBPairIgnoredStateChanged(opts *bind.FilterOpts, LBPair []common.Address) (*LegacyfactoryLBPairIgnoredStateChangedIterator, error) {

	var LBPairRule []interface{}
	for _, LBPairItem := range LBPair {
		LBPairRule = append(LBPairRule, LBPairItem)
	}

	logs, sub, err := _Legacyfactory.contract.FilterLogs(opts, "LBPairIgnoredStateChanged", LBPairRule)
	if err != nil {
		return nil, err
	}
	return &LegacyfactoryLBPairIgnoredStateChangedIterator{contract: _Legacyfactory.contract, event: "LBPairIgnoredStateChanged", logs: logs, sub: sub}, nil
}

// WatchLBPairIgnoredStateChanged is a free log subscription operation binding the contract event 0x44cf35361c9ff3c8c1397ec6410d5495cc481feaef35c9af11da1a637107de4f.
//
// Solidity: event LBPairIgnoredStateChanged(address indexed LBPair, bool ignored)
func (_Legacyfactory *LegacyfactoryFilterer) WatchLBPairIgnoredStateChanged(opts *bind.WatchOpts, sink chan<- *LegacyfactoryLBPairIgnoredStateChanged, LBPair []common.Address) (event.Subscription, error) {

	var LBPairRule []interface{}
	for _, LBPairItem := range LBPair {
		LBPairRule = append(LBPairRule, LBPairItem)
	}

	logs, sub, err := _Legacyfactory.contract.WatchLogs(opts, "LBPairIgnoredStateChanged", LBPairRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyfactoryLBPairIgnoredStateChanged)
				if err := _Legacyfactory.contract.UnpackLog(event, "LBPairIgnoredStateChanged", log); err != nil {
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

// ParseLBPairIgnoredStateChanged is a log parse operation binding the contract event 0x44cf35361c9ff3c8c1397ec6410d5495cc481feaef35c9af11da1a637107de4f.
//
// Solidity: event LBPairIgnoredStateChanged(address indexed LBPair, bool ignored)
func (_Legacyfactory *LegacyfactoryFilterer) ParseLBPairIgnoredStateChanged(log types.Log) (*LegacyfactoryLBPairIgnoredStateChanged, error) {
	event := new(LegacyfactoryLBPairIgnoredStateChanged)
	if err := _Legacyfactory.contract.UnpackLog(event, "LBPairIgnoredStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacyfactoryLBPairImplementationSetIterator is returned from FilterLBPairImplementationSet and is used to iterate over the raw logs and unpacked data for LBPairImplementationSet events raised by the Legacyfactory contract.
type LegacyfactoryLBPairImplementationSetIterator struct {
	Event *LegacyfactoryLBPairImplementationSet // Event containing the contract specifics and raw log

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
func (it *LegacyfactoryLBPairImplementationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyfactoryLBPairImplementationSet)
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
		it.Event = new(LegacyfactoryLBPairImplementationSet)
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
func (it *LegacyfactoryLBPairImplementationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyfactoryLBPairImplementationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyfactoryLBPairImplementationSet represents a LBPairImplementationSet event raised by the Legacyfactory contract.
type LegacyfactoryLBPairImplementationSet struct {
	OldLBPairImplementation common.Address
	LBPairImplementation    common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterLBPairImplementationSet is a free log retrieval operation binding the contract event 0x900d0e3d359f50e4f923ecdc06b401e07dbb9f485e17b07bcfc91a13000b277e.
//
// Solidity: event LBPairImplementationSet(address oldLBPairImplementation, address LBPairImplementation)
func (_Legacyfactory *LegacyfactoryFilterer) FilterLBPairImplementationSet(opts *bind.FilterOpts) (*LegacyfactoryLBPairImplementationSetIterator, error) {

	logs, sub, err := _Legacyfactory.contract.FilterLogs(opts, "LBPairImplementationSet")
	if err != nil {
		return nil, err
	}
	return &LegacyfactoryLBPairImplementationSetIterator{contract: _Legacyfactory.contract, event: "LBPairImplementationSet", logs: logs, sub: sub}, nil
}

// WatchLBPairImplementationSet is a free log subscription operation binding the contract event 0x900d0e3d359f50e4f923ecdc06b401e07dbb9f485e17b07bcfc91a13000b277e.
//
// Solidity: event LBPairImplementationSet(address oldLBPairImplementation, address LBPairImplementation)
func (_Legacyfactory *LegacyfactoryFilterer) WatchLBPairImplementationSet(opts *bind.WatchOpts, sink chan<- *LegacyfactoryLBPairImplementationSet) (event.Subscription, error) {

	logs, sub, err := _Legacyfactory.contract.WatchLogs(opts, "LBPairImplementationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyfactoryLBPairImplementationSet)
				if err := _Legacyfactory.contract.UnpackLog(event, "LBPairImplementationSet", log); err != nil {
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

// ParseLBPairImplementationSet is a log parse operation binding the contract event 0x900d0e3d359f50e4f923ecdc06b401e07dbb9f485e17b07bcfc91a13000b277e.
//
// Solidity: event LBPairImplementationSet(address oldLBPairImplementation, address LBPairImplementation)
func (_Legacyfactory *LegacyfactoryFilterer) ParseLBPairImplementationSet(log types.Log) (*LegacyfactoryLBPairImplementationSet, error) {
	event := new(LegacyfactoryLBPairImplementationSet)
	if err := _Legacyfactory.contract.UnpackLog(event, "LBPairImplementationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacyfactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Legacyfactory contract.
type LegacyfactoryOwnershipTransferredIterator struct {
	Event *LegacyfactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LegacyfactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyfactoryOwnershipTransferred)
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
		it.Event = new(LegacyfactoryOwnershipTransferred)
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
func (it *LegacyfactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyfactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyfactoryOwnershipTransferred represents a OwnershipTransferred event raised by the Legacyfactory contract.
type LegacyfactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Legacyfactory *LegacyfactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LegacyfactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Legacyfactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LegacyfactoryOwnershipTransferredIterator{contract: _Legacyfactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Legacyfactory *LegacyfactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LegacyfactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Legacyfactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyfactoryOwnershipTransferred)
				if err := _Legacyfactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Legacyfactory *LegacyfactoryFilterer) ParseOwnershipTransferred(log types.Log) (*LegacyfactoryOwnershipTransferred, error) {
	event := new(LegacyfactoryOwnershipTransferred)
	if err := _Legacyfactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacyfactoryPendingOwnerSetIterator is returned from FilterPendingOwnerSet and is used to iterate over the raw logs and unpacked data for PendingOwnerSet events raised by the Legacyfactory contract.
type LegacyfactoryPendingOwnerSetIterator struct {
	Event *LegacyfactoryPendingOwnerSet // Event containing the contract specifics and raw log

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
func (it *LegacyfactoryPendingOwnerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyfactoryPendingOwnerSet)
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
		it.Event = new(LegacyfactoryPendingOwnerSet)
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
func (it *LegacyfactoryPendingOwnerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyfactoryPendingOwnerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyfactoryPendingOwnerSet represents a PendingOwnerSet event raised by the Legacyfactory contract.
type LegacyfactoryPendingOwnerSet struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPendingOwnerSet is a free log retrieval operation binding the contract event 0x68f49b346b94582a8b5f9d10e3fe3365318fe8f191ff8dce7c59c6cad06b02f5.
//
// Solidity: event PendingOwnerSet(address indexed pendingOwner)
func (_Legacyfactory *LegacyfactoryFilterer) FilterPendingOwnerSet(opts *bind.FilterOpts, pendingOwner []common.Address) (*LegacyfactoryPendingOwnerSetIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Legacyfactory.contract.FilterLogs(opts, "PendingOwnerSet", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LegacyfactoryPendingOwnerSetIterator{contract: _Legacyfactory.contract, event: "PendingOwnerSet", logs: logs, sub: sub}, nil
}

// WatchPendingOwnerSet is a free log subscription operation binding the contract event 0x68f49b346b94582a8b5f9d10e3fe3365318fe8f191ff8dce7c59c6cad06b02f5.
//
// Solidity: event PendingOwnerSet(address indexed pendingOwner)
func (_Legacyfactory *LegacyfactoryFilterer) WatchPendingOwnerSet(opts *bind.WatchOpts, sink chan<- *LegacyfactoryPendingOwnerSet, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Legacyfactory.contract.WatchLogs(opts, "PendingOwnerSet", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyfactoryPendingOwnerSet)
				if err := _Legacyfactory.contract.UnpackLog(event, "PendingOwnerSet", log); err != nil {
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

// ParsePendingOwnerSet is a log parse operation binding the contract event 0x68f49b346b94582a8b5f9d10e3fe3365318fe8f191ff8dce7c59c6cad06b02f5.
//
// Solidity: event PendingOwnerSet(address indexed pendingOwner)
func (_Legacyfactory *LegacyfactoryFilterer) ParsePendingOwnerSet(log types.Log) (*LegacyfactoryPendingOwnerSet, error) {
	event := new(LegacyfactoryPendingOwnerSet)
	if err := _Legacyfactory.contract.UnpackLog(event, "PendingOwnerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacyfactoryPresetRemovedIterator is returned from FilterPresetRemoved and is used to iterate over the raw logs and unpacked data for PresetRemoved events raised by the Legacyfactory contract.
type LegacyfactoryPresetRemovedIterator struct {
	Event *LegacyfactoryPresetRemoved // Event containing the contract specifics and raw log

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
func (it *LegacyfactoryPresetRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyfactoryPresetRemoved)
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
		it.Event = new(LegacyfactoryPresetRemoved)
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
func (it *LegacyfactoryPresetRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyfactoryPresetRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyfactoryPresetRemoved represents a PresetRemoved event raised by the Legacyfactory contract.
type LegacyfactoryPresetRemoved struct {
	BinStep *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPresetRemoved is a free log retrieval operation binding the contract event 0xdd86b848bb56ff540caa68683fa467d0e7eb5f8b2d44e4ee435742eeeae9be13.
//
// Solidity: event PresetRemoved(uint256 indexed binStep)
func (_Legacyfactory *LegacyfactoryFilterer) FilterPresetRemoved(opts *bind.FilterOpts, binStep []*big.Int) (*LegacyfactoryPresetRemovedIterator, error) {

	var binStepRule []interface{}
	for _, binStepItem := range binStep {
		binStepRule = append(binStepRule, binStepItem)
	}

	logs, sub, err := _Legacyfactory.contract.FilterLogs(opts, "PresetRemoved", binStepRule)
	if err != nil {
		return nil, err
	}
	return &LegacyfactoryPresetRemovedIterator{contract: _Legacyfactory.contract, event: "PresetRemoved", logs: logs, sub: sub}, nil
}

// WatchPresetRemoved is a free log subscription operation binding the contract event 0xdd86b848bb56ff540caa68683fa467d0e7eb5f8b2d44e4ee435742eeeae9be13.
//
// Solidity: event PresetRemoved(uint256 indexed binStep)
func (_Legacyfactory *LegacyfactoryFilterer) WatchPresetRemoved(opts *bind.WatchOpts, sink chan<- *LegacyfactoryPresetRemoved, binStep []*big.Int) (event.Subscription, error) {

	var binStepRule []interface{}
	for _, binStepItem := range binStep {
		binStepRule = append(binStepRule, binStepItem)
	}

	logs, sub, err := _Legacyfactory.contract.WatchLogs(opts, "PresetRemoved", binStepRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyfactoryPresetRemoved)
				if err := _Legacyfactory.contract.UnpackLog(event, "PresetRemoved", log); err != nil {
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

// ParsePresetRemoved is a log parse operation binding the contract event 0xdd86b848bb56ff540caa68683fa467d0e7eb5f8b2d44e4ee435742eeeae9be13.
//
// Solidity: event PresetRemoved(uint256 indexed binStep)
func (_Legacyfactory *LegacyfactoryFilterer) ParsePresetRemoved(log types.Log) (*LegacyfactoryPresetRemoved, error) {
	event := new(LegacyfactoryPresetRemoved)
	if err := _Legacyfactory.contract.UnpackLog(event, "PresetRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacyfactoryPresetSetIterator is returned from FilterPresetSet and is used to iterate over the raw logs and unpacked data for PresetSet events raised by the Legacyfactory contract.
type LegacyfactoryPresetSetIterator struct {
	Event *LegacyfactoryPresetSet // Event containing the contract specifics and raw log

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
func (it *LegacyfactoryPresetSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyfactoryPresetSet)
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
		it.Event = new(LegacyfactoryPresetSet)
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
func (it *LegacyfactoryPresetSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyfactoryPresetSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyfactoryPresetSet represents a PresetSet event raised by the Legacyfactory contract.
type LegacyfactoryPresetSet struct {
	BinStep                  *big.Int
	BaseFactor               *big.Int
	FilterPeriod             *big.Int
	DecayPeriod              *big.Int
	ReductionFactor          *big.Int
	VariableFeeControl       *big.Int
	ProtocolShare            *big.Int
	MaxVolatilityAccumulated *big.Int
	SampleLifetime           *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterPresetSet is a free log retrieval operation binding the contract event 0x2f6cfdcc0e02e7355350f527dd3b5a957787b96f231165e48a3fdf90332a40cb.
//
// Solidity: event PresetSet(uint256 indexed binStep, uint256 baseFactor, uint256 filterPeriod, uint256 decayPeriod, uint256 reductionFactor, uint256 variableFeeControl, uint256 protocolShare, uint256 maxVolatilityAccumulated, uint256 sampleLifetime)
func (_Legacyfactory *LegacyfactoryFilterer) FilterPresetSet(opts *bind.FilterOpts, binStep []*big.Int) (*LegacyfactoryPresetSetIterator, error) {

	var binStepRule []interface{}
	for _, binStepItem := range binStep {
		binStepRule = append(binStepRule, binStepItem)
	}

	logs, sub, err := _Legacyfactory.contract.FilterLogs(opts, "PresetSet", binStepRule)
	if err != nil {
		return nil, err
	}
	return &LegacyfactoryPresetSetIterator{contract: _Legacyfactory.contract, event: "PresetSet", logs: logs, sub: sub}, nil
}

// WatchPresetSet is a free log subscription operation binding the contract event 0x2f6cfdcc0e02e7355350f527dd3b5a957787b96f231165e48a3fdf90332a40cb.
//
// Solidity: event PresetSet(uint256 indexed binStep, uint256 baseFactor, uint256 filterPeriod, uint256 decayPeriod, uint256 reductionFactor, uint256 variableFeeControl, uint256 protocolShare, uint256 maxVolatilityAccumulated, uint256 sampleLifetime)
func (_Legacyfactory *LegacyfactoryFilterer) WatchPresetSet(opts *bind.WatchOpts, sink chan<- *LegacyfactoryPresetSet, binStep []*big.Int) (event.Subscription, error) {

	var binStepRule []interface{}
	for _, binStepItem := range binStep {
		binStepRule = append(binStepRule, binStepItem)
	}

	logs, sub, err := _Legacyfactory.contract.WatchLogs(opts, "PresetSet", binStepRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyfactoryPresetSet)
				if err := _Legacyfactory.contract.UnpackLog(event, "PresetSet", log); err != nil {
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

// ParsePresetSet is a log parse operation binding the contract event 0x2f6cfdcc0e02e7355350f527dd3b5a957787b96f231165e48a3fdf90332a40cb.
//
// Solidity: event PresetSet(uint256 indexed binStep, uint256 baseFactor, uint256 filterPeriod, uint256 decayPeriod, uint256 reductionFactor, uint256 variableFeeControl, uint256 protocolShare, uint256 maxVolatilityAccumulated, uint256 sampleLifetime)
func (_Legacyfactory *LegacyfactoryFilterer) ParsePresetSet(log types.Log) (*LegacyfactoryPresetSet, error) {
	event := new(LegacyfactoryPresetSet)
	if err := _Legacyfactory.contract.UnpackLog(event, "PresetSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacyfactoryQuoteAssetAddedIterator is returned from FilterQuoteAssetAdded and is used to iterate over the raw logs and unpacked data for QuoteAssetAdded events raised by the Legacyfactory contract.
type LegacyfactoryQuoteAssetAddedIterator struct {
	Event *LegacyfactoryQuoteAssetAdded // Event containing the contract specifics and raw log

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
func (it *LegacyfactoryQuoteAssetAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyfactoryQuoteAssetAdded)
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
		it.Event = new(LegacyfactoryQuoteAssetAdded)
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
func (it *LegacyfactoryQuoteAssetAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyfactoryQuoteAssetAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyfactoryQuoteAssetAdded represents a QuoteAssetAdded event raised by the Legacyfactory contract.
type LegacyfactoryQuoteAssetAdded struct {
	QuoteAsset common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterQuoteAssetAdded is a free log retrieval operation binding the contract event 0x84cc2115995684dcb0cd3d3a9565e3d32f075de81db70c8dc3a719b2a47af67e.
//
// Solidity: event QuoteAssetAdded(address indexed quoteAsset)
func (_Legacyfactory *LegacyfactoryFilterer) FilterQuoteAssetAdded(opts *bind.FilterOpts, quoteAsset []common.Address) (*LegacyfactoryQuoteAssetAddedIterator, error) {

	var quoteAssetRule []interface{}
	for _, quoteAssetItem := range quoteAsset {
		quoteAssetRule = append(quoteAssetRule, quoteAssetItem)
	}

	logs, sub, err := _Legacyfactory.contract.FilterLogs(opts, "QuoteAssetAdded", quoteAssetRule)
	if err != nil {
		return nil, err
	}
	return &LegacyfactoryQuoteAssetAddedIterator{contract: _Legacyfactory.contract, event: "QuoteAssetAdded", logs: logs, sub: sub}, nil
}

// WatchQuoteAssetAdded is a free log subscription operation binding the contract event 0x84cc2115995684dcb0cd3d3a9565e3d32f075de81db70c8dc3a719b2a47af67e.
//
// Solidity: event QuoteAssetAdded(address indexed quoteAsset)
func (_Legacyfactory *LegacyfactoryFilterer) WatchQuoteAssetAdded(opts *bind.WatchOpts, sink chan<- *LegacyfactoryQuoteAssetAdded, quoteAsset []common.Address) (event.Subscription, error) {

	var quoteAssetRule []interface{}
	for _, quoteAssetItem := range quoteAsset {
		quoteAssetRule = append(quoteAssetRule, quoteAssetItem)
	}

	logs, sub, err := _Legacyfactory.contract.WatchLogs(opts, "QuoteAssetAdded", quoteAssetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyfactoryQuoteAssetAdded)
				if err := _Legacyfactory.contract.UnpackLog(event, "QuoteAssetAdded", log); err != nil {
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

// ParseQuoteAssetAdded is a log parse operation binding the contract event 0x84cc2115995684dcb0cd3d3a9565e3d32f075de81db70c8dc3a719b2a47af67e.
//
// Solidity: event QuoteAssetAdded(address indexed quoteAsset)
func (_Legacyfactory *LegacyfactoryFilterer) ParseQuoteAssetAdded(log types.Log) (*LegacyfactoryQuoteAssetAdded, error) {
	event := new(LegacyfactoryQuoteAssetAdded)
	if err := _Legacyfactory.contract.UnpackLog(event, "QuoteAssetAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacyfactoryQuoteAssetRemovedIterator is returned from FilterQuoteAssetRemoved and is used to iterate over the raw logs and unpacked data for QuoteAssetRemoved events raised by the Legacyfactory contract.
type LegacyfactoryQuoteAssetRemovedIterator struct {
	Event *LegacyfactoryQuoteAssetRemoved // Event containing the contract specifics and raw log

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
func (it *LegacyfactoryQuoteAssetRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyfactoryQuoteAssetRemoved)
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
		it.Event = new(LegacyfactoryQuoteAssetRemoved)
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
func (it *LegacyfactoryQuoteAssetRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyfactoryQuoteAssetRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyfactoryQuoteAssetRemoved represents a QuoteAssetRemoved event raised by the Legacyfactory contract.
type LegacyfactoryQuoteAssetRemoved struct {
	QuoteAsset common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterQuoteAssetRemoved is a free log retrieval operation binding the contract event 0x0b767739217755d8af5a2ba75b181a19fa1750f8bb701f09311cb19a90140cb3.
//
// Solidity: event QuoteAssetRemoved(address indexed quoteAsset)
func (_Legacyfactory *LegacyfactoryFilterer) FilterQuoteAssetRemoved(opts *bind.FilterOpts, quoteAsset []common.Address) (*LegacyfactoryQuoteAssetRemovedIterator, error) {

	var quoteAssetRule []interface{}
	for _, quoteAssetItem := range quoteAsset {
		quoteAssetRule = append(quoteAssetRule, quoteAssetItem)
	}

	logs, sub, err := _Legacyfactory.contract.FilterLogs(opts, "QuoteAssetRemoved", quoteAssetRule)
	if err != nil {
		return nil, err
	}
	return &LegacyfactoryQuoteAssetRemovedIterator{contract: _Legacyfactory.contract, event: "QuoteAssetRemoved", logs: logs, sub: sub}, nil
}

// WatchQuoteAssetRemoved is a free log subscription operation binding the contract event 0x0b767739217755d8af5a2ba75b181a19fa1750f8bb701f09311cb19a90140cb3.
//
// Solidity: event QuoteAssetRemoved(address indexed quoteAsset)
func (_Legacyfactory *LegacyfactoryFilterer) WatchQuoteAssetRemoved(opts *bind.WatchOpts, sink chan<- *LegacyfactoryQuoteAssetRemoved, quoteAsset []common.Address) (event.Subscription, error) {

	var quoteAssetRule []interface{}
	for _, quoteAssetItem := range quoteAsset {
		quoteAssetRule = append(quoteAssetRule, quoteAssetItem)
	}

	logs, sub, err := _Legacyfactory.contract.WatchLogs(opts, "QuoteAssetRemoved", quoteAssetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyfactoryQuoteAssetRemoved)
				if err := _Legacyfactory.contract.UnpackLog(event, "QuoteAssetRemoved", log); err != nil {
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

// ParseQuoteAssetRemoved is a log parse operation binding the contract event 0x0b767739217755d8af5a2ba75b181a19fa1750f8bb701f09311cb19a90140cb3.
//
// Solidity: event QuoteAssetRemoved(address indexed quoteAsset)
func (_Legacyfactory *LegacyfactoryFilterer) ParseQuoteAssetRemoved(log types.Log) (*LegacyfactoryQuoteAssetRemoved, error) {
	event := new(LegacyfactoryQuoteAssetRemoved)
	if err := _Legacyfactory.contract.UnpackLog(event, "QuoteAssetRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
