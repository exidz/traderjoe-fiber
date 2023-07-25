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

// LegacyRouterLiquidityParameters is an auto generated low-level Go binding around an user-defined struct.
type ILegacyRouterLiquidityParameters struct {
	TokenX          common.Address
	TokenY          common.Address
	BinStep         *big.Int
	AmountX         *big.Int
	AmountY         *big.Int
	AmountXMin      *big.Int
	AmountYMin      *big.Int
	ActiveIdDesired *big.Int
	IdSlippage      *big.Int
	DeltaIds        []*big.Int
	DistributionX   []*big.Int
	DistributionY   []*big.Int
	To              common.Address
	Deadline        *big.Int
}

// LegacyrouterMetaData contains all meta data concerning the Legacyrouter contract.
var LegacyrouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractILBFactory\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"contractIJoeFactory\",\"name\":\"_oldFactory\",\"type\":\"address\"},{\"internalType\":\"contractIWAVAX\",\"name\":\"_wavax\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bp\",\"type\":\"uint256\"}],\"name\":\"BinHelper__BinStepOverflows\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BinHelper__IdOverflows\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"JoeLibrary__InsufficientAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"JoeLibrary__InsufficientLiquidity\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountXMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountYMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"name\":\"LBRouter__AmountSlippageCaught\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"LBRouter__BinReserveOverflows\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBRouter__BrokenSwapSafetyCheck\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentTimestamp\",\"type\":\"uint256\"}],\"name\":\"LBRouter__DeadlineExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LBRouter__FailedToSendAVAX\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idSlippage\",\"type\":\"uint256\"}],\"name\":\"LBRouter__IdDesiredOverflows\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"id\",\"type\":\"int256\"}],\"name\":\"LBRouter__IdOverflows\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"activeIdDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idSlippage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activeId\",\"type\":\"uint256\"}],\"name\":\"LBRouter__IdSlippageCaught\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"LBRouter__InsufficientAmountOut\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wrongToken\",\"type\":\"address\"}],\"name\":\"LBRouter__InvalidTokenPath\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBRouter__LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"LBRouter__MaxAmountInExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBRouter__NotFactoryOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"}],\"name\":\"LBRouter__PairNotCreated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBRouter__SenderIsNotWAVAX\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"LBRouter__SwapOverflows\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"excess\",\"type\":\"uint256\"}],\"name\":\"LBRouter__TooMuchTokensIn\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserve\",\"type\":\"uint256\"}],\"name\":\"LBRouter__WrongAmounts\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"}],\"name\":\"LBRouter__WrongAvaxLiquidityParameters\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBRouter__WrongTokenOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Math128x128__LogUnderflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"y\",\"type\":\"int256\"}],\"name\":\"Math128x128__PowerUnderflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"prod1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"}],\"name\":\"Math512Bits__MulDivOverflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"prod1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"Math512Bits__MulShiftOverflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"Math512Bits__OffsetOverflows\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"SafeCast__Exceeds128Bits\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"SafeCast__Exceeds40Bits\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenHelper__CallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenHelper__NonContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenHelper__TransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountXMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountYMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activeIdDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idSlippage\",\"type\":\"uint256\"},{\"internalType\":\"int256[]\",\"name\":\"deltaIds\",\"type\":\"int256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"distributionX\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"distributionY\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structILBRouter.LiquidityParameters\",\"name\":\"_liquidityParameters\",\"type\":\"tuple\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"depositIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"liquidityMinted\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountXMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountYMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activeIdDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idSlippage\",\"type\":\"uint256\"},{\"internalType\":\"int256[]\",\"name\":\"deltaIds\",\"type\":\"int256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"distributionX\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"distributionY\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structILBRouter.LiquidityParameters\",\"name\":\"_liquidityParameters\",\"type\":\"tuple\"}],\"name\":\"addLiquidityAVAX\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"depositIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"liquidityMinted\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_activeId\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"_binStep\",\"type\":\"uint16\"}],\"name\":\"createLBPair\",\"outputs\":[{\"internalType\":\"contractILBPair\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"contractILBFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILBPair\",\"name\":\"_LBPair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"getIdFromPrice\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILBPair\",\"name\":\"_LBPair\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_id\",\"type\":\"uint24\"}],\"name\":\"getPriceFromId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILBPair\",\"name\":\"_LBPair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_swapForY\",\"type\":\"bool\"}],\"name\":\"getSwapIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feesIn\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILBPair\",\"name\":\"_LBPair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_swapForY\",\"type\":\"bool\"}],\"name\":\"getSwapOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feesIn\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oldFactory\",\"outputs\":[{\"internalType\":\"contractIJoeFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_tokenY\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_binStep\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_amountXMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountYMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_binStep\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountAVAXMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityAVAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAX\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"_tokenPath\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"swapAVAXForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountsIn\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"_tokenPath\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactAVAXForTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"_tokenPath\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactAVAXForTokensSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountOutMinAVAX\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"_tokenPath\",\"type\":\"address[]\"},{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForAVAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountOutMinAVAX\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"_tokenPath\",\"type\":\"address[]\"},{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForAVAXSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"_tokenPath\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"_tokenPath\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountAVAXOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"_tokenPath\",\"type\":\"address[]\"},{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactAVAX\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountsIn\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"_tokenPath\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountsIn\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"sweep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILBToken\",\"name\":\"_lbToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"sweepLBToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wavax\",\"outputs\":[{\"internalType\":\"contractIWAVAX\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// LegacyrouterABI is the input ABI used to generate the binding from.
// Deprecated: Use LegacyrouterMetaData.ABI instead.
var LegacyrouterABI = LegacyrouterMetaData.ABI

// Legacyrouter is an auto generated Go binding around an Ethereum contract.
type Legacyrouter struct {
	LegacyrouterCaller     // Read-only binding to the contract
	LegacyrouterTransactor // Write-only binding to the contract
	LegacyrouterFilterer   // Log filterer for contract events
}

// LegacyrouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type LegacyrouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyrouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LegacyrouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyrouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LegacyrouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyrouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LegacyrouterSession struct {
	Contract     *Legacyrouter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LegacyrouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LegacyrouterCallerSession struct {
	Contract *LegacyrouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// LegacyrouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LegacyrouterTransactorSession struct {
	Contract     *LegacyrouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LegacyrouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type LegacyrouterRaw struct {
	Contract *Legacyrouter // Generic contract binding to access the raw methods on
}

// LegacyrouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LegacyrouterCallerRaw struct {
	Contract *LegacyrouterCaller // Generic read-only contract binding to access the raw methods on
}

// LegacyrouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LegacyrouterTransactorRaw struct {
	Contract *LegacyrouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLegacyrouter creates a new instance of Legacyrouter, bound to a specific deployed contract.
func NewLegacyrouter(address common.Address, backend bind.ContractBackend) (*Legacyrouter, error) {
	contract, err := bindLegacyrouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Legacyrouter{LegacyrouterCaller: LegacyrouterCaller{contract: contract}, LegacyrouterTransactor: LegacyrouterTransactor{contract: contract}, LegacyrouterFilterer: LegacyrouterFilterer{contract: contract}}, nil
}

// NewLegacyrouterCaller creates a new read-only instance of Legacyrouter, bound to a specific deployed contract.
func NewLegacyrouterCaller(address common.Address, caller bind.ContractCaller) (*LegacyrouterCaller, error) {
	contract, err := bindLegacyrouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LegacyrouterCaller{contract: contract}, nil
}

// NewLegacyrouterTransactor creates a new write-only instance of Legacyrouter, bound to a specific deployed contract.
func NewLegacyrouterTransactor(address common.Address, transactor bind.ContractTransactor) (*LegacyrouterTransactor, error) {
	contract, err := bindLegacyrouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LegacyrouterTransactor{contract: contract}, nil
}

// NewLegacyrouterFilterer creates a new log filterer instance of Legacyrouter, bound to a specific deployed contract.
func NewLegacyrouterFilterer(address common.Address, filterer bind.ContractFilterer) (*LegacyrouterFilterer, error) {
	contract, err := bindLegacyrouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LegacyrouterFilterer{contract: contract}, nil
}

// bindLegacyrouter binds a generic wrapper to an already deployed contract.
func bindLegacyrouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LegacyrouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Legacyrouter *LegacyrouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Legacyrouter.Contract.LegacyrouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Legacyrouter *LegacyrouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Legacyrouter.Contract.LegacyrouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Legacyrouter *LegacyrouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Legacyrouter.Contract.LegacyrouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Legacyrouter *LegacyrouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Legacyrouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Legacyrouter *LegacyrouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Legacyrouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Legacyrouter *LegacyrouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Legacyrouter.Contract.contract.Transact(opts, method, params...)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Legacyrouter *LegacyrouterCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Legacyrouter.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Legacyrouter *LegacyrouterSession) Factory() (common.Address, error) {
	return _Legacyrouter.Contract.Factory(&_Legacyrouter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Legacyrouter *LegacyrouterCallerSession) Factory() (common.Address, error) {
	return _Legacyrouter.Contract.Factory(&_Legacyrouter.CallOpts)
}

// GetIdFromPrice is a free data retrieval call binding the contract method 0xf96fe925.
//
// Solidity: function getIdFromPrice(address _LBPair, uint256 _price) view returns(uint24)
func (_Legacyrouter *LegacyrouterCaller) GetIdFromPrice(opts *bind.CallOpts, _LBPair common.Address, _price *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Legacyrouter.contract.Call(opts, &out, "getIdFromPrice", _LBPair, _price)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIdFromPrice is a free data retrieval call binding the contract method 0xf96fe925.
//
// Solidity: function getIdFromPrice(address _LBPair, uint256 _price) view returns(uint24)
func (_Legacyrouter *LegacyrouterSession) GetIdFromPrice(_LBPair common.Address, _price *big.Int) (*big.Int, error) {
	return _Legacyrouter.Contract.GetIdFromPrice(&_Legacyrouter.CallOpts, _LBPair, _price)
}

// GetIdFromPrice is a free data retrieval call binding the contract method 0xf96fe925.
//
// Solidity: function getIdFromPrice(address _LBPair, uint256 _price) view returns(uint24)
func (_Legacyrouter *LegacyrouterCallerSession) GetIdFromPrice(_LBPair common.Address, _price *big.Int) (*big.Int, error) {
	return _Legacyrouter.Contract.GetIdFromPrice(&_Legacyrouter.CallOpts, _LBPair, _price)
}

// GetPriceFromId is a free data retrieval call binding the contract method 0xd0e380f2.
//
// Solidity: function getPriceFromId(address _LBPair, uint24 _id) view returns(uint256)
func (_Legacyrouter *LegacyrouterCaller) GetPriceFromId(opts *bind.CallOpts, _LBPair common.Address, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Legacyrouter.contract.Call(opts, &out, "getPriceFromId", _LBPair, _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriceFromId is a free data retrieval call binding the contract method 0xd0e380f2.
//
// Solidity: function getPriceFromId(address _LBPair, uint24 _id) view returns(uint256)
func (_Legacyrouter *LegacyrouterSession) GetPriceFromId(_LBPair common.Address, _id *big.Int) (*big.Int, error) {
	return _Legacyrouter.Contract.GetPriceFromId(&_Legacyrouter.CallOpts, _LBPair, _id)
}

// GetPriceFromId is a free data retrieval call binding the contract method 0xd0e380f2.
//
// Solidity: function getPriceFromId(address _LBPair, uint24 _id) view returns(uint256)
func (_Legacyrouter *LegacyrouterCallerSession) GetPriceFromId(_LBPair common.Address, _id *big.Int) (*big.Int, error) {
	return _Legacyrouter.Contract.GetPriceFromId(&_Legacyrouter.CallOpts, _LBPair, _id)
}

// GetSwapIn is a free data retrieval call binding the contract method 0x5bdd4b7c.
//
// Solidity: function getSwapIn(address _LBPair, uint256 _amountOut, bool _swapForY) view returns(uint256 amountIn, uint256 feesIn)
func (_Legacyrouter *LegacyrouterCaller) GetSwapIn(opts *bind.CallOpts, _LBPair common.Address, _amountOut *big.Int, _swapForY bool) (struct {
	AmountIn *big.Int
	FeesIn   *big.Int
}, error) {
	var out []interface{}
	err := _Legacyrouter.contract.Call(opts, &out, "getSwapIn", _LBPair, _amountOut, _swapForY)

	outstruct := new(struct {
		AmountIn *big.Int
		FeesIn   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountIn = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FeesIn = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSwapIn is a free data retrieval call binding the contract method 0x5bdd4b7c.
//
// Solidity: function getSwapIn(address _LBPair, uint256 _amountOut, bool _swapForY) view returns(uint256 amountIn, uint256 feesIn)
func (_Legacyrouter *LegacyrouterSession) GetSwapIn(_LBPair common.Address, _amountOut *big.Int, _swapForY bool) (struct {
	AmountIn *big.Int
	FeesIn   *big.Int
}, error) {
	return _Legacyrouter.Contract.GetSwapIn(&_Legacyrouter.CallOpts, _LBPair, _amountOut, _swapForY)
}

// GetSwapIn is a free data retrieval call binding the contract method 0x5bdd4b7c.
//
// Solidity: function getSwapIn(address _LBPair, uint256 _amountOut, bool _swapForY) view returns(uint256 amountIn, uint256 feesIn)
func (_Legacyrouter *LegacyrouterCallerSession) GetSwapIn(_LBPair common.Address, _amountOut *big.Int, _swapForY bool) (struct {
	AmountIn *big.Int
	FeesIn   *big.Int
}, error) {
	return _Legacyrouter.Contract.GetSwapIn(&_Legacyrouter.CallOpts, _LBPair, _amountOut, _swapForY)
}

// GetSwapOut is a free data retrieval call binding the contract method 0x2004b724.
//
// Solidity: function getSwapOut(address _LBPair, uint256 _amountIn, bool _swapForY) view returns(uint256 amountOut, uint256 feesIn)
func (_Legacyrouter *LegacyrouterCaller) GetSwapOut(opts *bind.CallOpts, _LBPair common.Address, _amountIn *big.Int, _swapForY bool) (struct {
	AmountOut *big.Int
	FeesIn    *big.Int
}, error) {
	var out []interface{}
	err := _Legacyrouter.contract.Call(opts, &out, "getSwapOut", _LBPair, _amountIn, _swapForY)

	outstruct := new(struct {
		AmountOut *big.Int
		FeesIn    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountOut = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FeesIn = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSwapOut is a free data retrieval call binding the contract method 0x2004b724.
//
// Solidity: function getSwapOut(address _LBPair, uint256 _amountIn, bool _swapForY) view returns(uint256 amountOut, uint256 feesIn)
func (_Legacyrouter *LegacyrouterSession) GetSwapOut(_LBPair common.Address, _amountIn *big.Int, _swapForY bool) (struct {
	AmountOut *big.Int
	FeesIn    *big.Int
}, error) {
	return _Legacyrouter.Contract.GetSwapOut(&_Legacyrouter.CallOpts, _LBPair, _amountIn, _swapForY)
}

// GetSwapOut is a free data retrieval call binding the contract method 0x2004b724.
//
// Solidity: function getSwapOut(address _LBPair, uint256 _amountIn, bool _swapForY) view returns(uint256 amountOut, uint256 feesIn)
func (_Legacyrouter *LegacyrouterCallerSession) GetSwapOut(_LBPair common.Address, _amountIn *big.Int, _swapForY bool) (struct {
	AmountOut *big.Int
	FeesIn    *big.Int
}, error) {
	return _Legacyrouter.Contract.GetSwapOut(&_Legacyrouter.CallOpts, _LBPair, _amountIn, _swapForY)
}

// OldFactory is a free data retrieval call binding the contract method 0x1bd6dfe1.
//
// Solidity: function oldFactory() view returns(address)
func (_Legacyrouter *LegacyrouterCaller) OldFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Legacyrouter.contract.Call(opts, &out, "oldFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OldFactory is a free data retrieval call binding the contract method 0x1bd6dfe1.
//
// Solidity: function oldFactory() view returns(address)
func (_Legacyrouter *LegacyrouterSession) OldFactory() (common.Address, error) {
	return _Legacyrouter.Contract.OldFactory(&_Legacyrouter.CallOpts)
}

// OldFactory is a free data retrieval call binding the contract method 0x1bd6dfe1.
//
// Solidity: function oldFactory() view returns(address)
func (_Legacyrouter *LegacyrouterCallerSession) OldFactory() (common.Address, error) {
	return _Legacyrouter.Contract.OldFactory(&_Legacyrouter.CallOpts)
}

// Wavax is a free data retrieval call binding the contract method 0x117be4c2.
//
// Solidity: function wavax() view returns(address)
func (_Legacyrouter *LegacyrouterCaller) Wavax(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Legacyrouter.contract.Call(opts, &out, "wavax")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wavax is a free data retrieval call binding the contract method 0x117be4c2.
//
// Solidity: function wavax() view returns(address)
func (_Legacyrouter *LegacyrouterSession) Wavax() (common.Address, error) {
	return _Legacyrouter.Contract.Wavax(&_Legacyrouter.CallOpts)
}

// Wavax is a free data retrieval call binding the contract method 0x117be4c2.
//
// Solidity: function wavax() view returns(address)
func (_Legacyrouter *LegacyrouterCallerSession) Wavax() (common.Address, error) {
	return _Legacyrouter.Contract.Wavax(&_Legacyrouter.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe324a3e4.
//
// Solidity: function addLiquidity((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256[],uint256[],uint256[],address,uint256) _liquidityParameters) returns(uint256[] depositIds, uint256[] liquidityMinted)
func (_Legacyrouter *LegacyrouterTransactor) AddLiquidity(opts *bind.TransactOpts, _liquidityParameters ILBRouterLiquidityParameters) (*types.Transaction, error) {
	return _Legacyrouter.contract.Transact(opts, "addLiquidity", _liquidityParameters)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe324a3e4.
//
// Solidity: function addLiquidity((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256[],uint256[],uint256[],address,uint256) _liquidityParameters) returns(uint256[] depositIds, uint256[] liquidityMinted)
func (_Legacyrouter *LegacyrouterSession) AddLiquidity(_liquidityParameters ILBRouterLiquidityParameters) (*types.Transaction, error) {
	return _Legacyrouter.Contract.AddLiquidity(&_Legacyrouter.TransactOpts, _liquidityParameters)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe324a3e4.
//
// Solidity: function addLiquidity((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256[],uint256[],uint256[],address,uint256) _liquidityParameters) returns(uint256[] depositIds, uint256[] liquidityMinted)
func (_Legacyrouter *LegacyrouterTransactorSession) AddLiquidity(_liquidityParameters ILBRouterLiquidityParameters) (*types.Transaction, error) {
	return _Legacyrouter.Contract.AddLiquidity(&_Legacyrouter.TransactOpts, _liquidityParameters)
}

// AddLiquidityAVAX is a paid mutator transaction binding the contract method 0xea8f43d8.
//
// Solidity: function addLiquidityAVAX((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256[],uint256[],uint256[],address,uint256) _liquidityParameters) payable returns(uint256[] depositIds, uint256[] liquidityMinted)
func (_Legacyrouter *LegacyrouterTransactor) AddLiquidityAVAX(opts *bind.TransactOpts, _liquidityParameters ILBRouterLiquidityParameters) (*types.Transaction, error) {
	return _Legacyrouter.contract.Transact(opts, "addLiquidityAVAX", _liquidityParameters)
}

// AddLiquidityAVAX is a paid mutator transaction binding the contract method 0xea8f43d8.
//
// Solidity: function addLiquidityAVAX((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256[],uint256[],uint256[],address,uint256) _liquidityParameters) payable returns(uint256[] depositIds, uint256[] liquidityMinted)
func (_Legacyrouter *LegacyrouterSession) AddLiquidityAVAX(_liquidityParameters ILBRouterLiquidityParameters) (*types.Transaction, error) {
	return _Legacyrouter.Contract.AddLiquidityAVAX(&_Legacyrouter.TransactOpts, _liquidityParameters)
}

// AddLiquidityAVAX is a paid mutator transaction binding the contract method 0xea8f43d8.
//
// Solidity: function addLiquidityAVAX((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256[],uint256[],uint256[],address,uint256) _liquidityParameters) payable returns(uint256[] depositIds, uint256[] liquidityMinted)
func (_Legacyrouter *LegacyrouterTransactorSession) AddLiquidityAVAX(_liquidityParameters ILBRouterLiquidityParameters) (*types.Transaction, error) {
	return _Legacyrouter.Contract.AddLiquidityAVAX(&_Legacyrouter.TransactOpts, _liquidityParameters)
}

// CreateLBPair is a paid mutator transaction binding the contract method 0x659ac74b.
//
// Solidity: function createLBPair(address _tokenX, address _tokenY, uint24 _activeId, uint16 _binStep) returns(address pair)
func (_Legacyrouter *LegacyrouterTransactor) CreateLBPair(opts *bind.TransactOpts, _tokenX common.Address, _tokenY common.Address, _activeId *big.Int, _binStep uint16) (*types.Transaction, error) {
	return _Legacyrouter.contract.Transact(opts, "createLBPair", _tokenX, _tokenY, _activeId, _binStep)
}

// CreateLBPair is a paid mutator transaction binding the contract method 0x659ac74b.
//
// Solidity: function createLBPair(address _tokenX, address _tokenY, uint24 _activeId, uint16 _binStep) returns(address pair)
func (_Legacyrouter *LegacyrouterSession) CreateLBPair(_tokenX common.Address, _tokenY common.Address, _activeId *big.Int, _binStep uint16) (*types.Transaction, error) {
	return _Legacyrouter.Contract.CreateLBPair(&_Legacyrouter.TransactOpts, _tokenX, _tokenY, _activeId, _binStep)
}

// CreateLBPair is a paid mutator transaction binding the contract method 0x659ac74b.
//
// Solidity: function createLBPair(address _tokenX, address _tokenY, uint24 _activeId, uint16 _binStep) returns(address pair)
func (_Legacyrouter *LegacyrouterTransactorSession) CreateLBPair(_tokenX common.Address, _tokenY common.Address, _activeId *big.Int, _binStep uint16) (*types.Transaction, error) {
	return _Legacyrouter.Contract.CreateLBPair(&_Legacyrouter.TransactOpts, _tokenX, _tokenY, _activeId, _binStep)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xc22159b6.
//
// Solidity: function removeLiquidity(address _tokenX, address _tokenY, uint16 _binStep, uint256 _amountXMin, uint256 _amountYMin, uint256[] _ids, uint256[] _amounts, address _to, uint256 _deadline) returns(uint256 amountX, uint256 amountY)
func (_Legacyrouter *LegacyrouterTransactor) RemoveLiquidity(opts *bind.TransactOpts, _tokenX common.Address, _tokenY common.Address, _binStep uint16, _amountXMin *big.Int, _amountYMin *big.Int, _ids []*big.Int, _amounts []*big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.contract.Transact(opts, "removeLiquidity", _tokenX, _tokenY, _binStep, _amountXMin, _amountYMin, _ids, _amounts, _to, _deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xc22159b6.
//
// Solidity: function removeLiquidity(address _tokenX, address _tokenY, uint16 _binStep, uint256 _amountXMin, uint256 _amountYMin, uint256[] _ids, uint256[] _amounts, address _to, uint256 _deadline) returns(uint256 amountX, uint256 amountY)
func (_Legacyrouter *LegacyrouterSession) RemoveLiquidity(_tokenX common.Address, _tokenY common.Address, _binStep uint16, _amountXMin *big.Int, _amountYMin *big.Int, _ids []*big.Int, _amounts []*big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.Contract.RemoveLiquidity(&_Legacyrouter.TransactOpts, _tokenX, _tokenY, _binStep, _amountXMin, _amountYMin, _ids, _amounts, _to, _deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xc22159b6.
//
// Solidity: function removeLiquidity(address _tokenX, address _tokenY, uint16 _binStep, uint256 _amountXMin, uint256 _amountYMin, uint256[] _ids, uint256[] _amounts, address _to, uint256 _deadline) returns(uint256 amountX, uint256 amountY)
func (_Legacyrouter *LegacyrouterTransactorSession) RemoveLiquidity(_tokenX common.Address, _tokenY common.Address, _binStep uint16, _amountXMin *big.Int, _amountYMin *big.Int, _ids []*big.Int, _amounts []*big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.Contract.RemoveLiquidity(&_Legacyrouter.TransactOpts, _tokenX, _tokenY, _binStep, _amountXMin, _amountYMin, _ids, _amounts, _to, _deadline)
}

// RemoveLiquidityAVAX is a paid mutator transaction binding the contract method 0xbcb1c957.
//
// Solidity: function removeLiquidityAVAX(address _token, uint16 _binStep, uint256 _amountTokenMin, uint256 _amountAVAXMin, uint256[] _ids, uint256[] _amounts, address _to, uint256 _deadline) returns(uint256 amountToken, uint256 amountAVAX)
func (_Legacyrouter *LegacyrouterTransactor) RemoveLiquidityAVAX(opts *bind.TransactOpts, _token common.Address, _binStep uint16, _amountTokenMin *big.Int, _amountAVAXMin *big.Int, _ids []*big.Int, _amounts []*big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.contract.Transact(opts, "removeLiquidityAVAX", _token, _binStep, _amountTokenMin, _amountAVAXMin, _ids, _amounts, _to, _deadline)
}

// RemoveLiquidityAVAX is a paid mutator transaction binding the contract method 0xbcb1c957.
//
// Solidity: function removeLiquidityAVAX(address _token, uint16 _binStep, uint256 _amountTokenMin, uint256 _amountAVAXMin, uint256[] _ids, uint256[] _amounts, address _to, uint256 _deadline) returns(uint256 amountToken, uint256 amountAVAX)
func (_Legacyrouter *LegacyrouterSession) RemoveLiquidityAVAX(_token common.Address, _binStep uint16, _amountTokenMin *big.Int, _amountAVAXMin *big.Int, _ids []*big.Int, _amounts []*big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.Contract.RemoveLiquidityAVAX(&_Legacyrouter.TransactOpts, _token, _binStep, _amountTokenMin, _amountAVAXMin, _ids, _amounts, _to, _deadline)
}

// RemoveLiquidityAVAX is a paid mutator transaction binding the contract method 0xbcb1c957.
//
// Solidity: function removeLiquidityAVAX(address _token, uint16 _binStep, uint256 _amountTokenMin, uint256 _amountAVAXMin, uint256[] _ids, uint256[] _amounts, address _to, uint256 _deadline) returns(uint256 amountToken, uint256 amountAVAX)
func (_Legacyrouter *LegacyrouterTransactorSession) RemoveLiquidityAVAX(_token common.Address, _binStep uint16, _amountTokenMin *big.Int, _amountAVAXMin *big.Int, _ids []*big.Int, _amounts []*big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.Contract.RemoveLiquidityAVAX(&_Legacyrouter.TransactOpts, _token, _binStep, _amountTokenMin, _amountAVAXMin, _ids, _amounts, _to, _deadline)
}

// SwapAVAXForExactTokens is a paid mutator transaction binding the contract method 0x42f564a0.
//
// Solidity: function swapAVAXForExactTokens(uint256 _amountOut, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) payable returns(uint256[] amountsIn)
func (_Legacyrouter *LegacyrouterTransactor) SwapAVAXForExactTokens(opts *bind.TransactOpts, _amountOut *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.contract.Transact(opts, "swapAVAXForExactTokens", _amountOut, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapAVAXForExactTokens is a paid mutator transaction binding the contract method 0x42f564a0.
//
// Solidity: function swapAVAXForExactTokens(uint256 _amountOut, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) payable returns(uint256[] amountsIn)
func (_Legacyrouter *LegacyrouterSession) SwapAVAXForExactTokens(_amountOut *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.Contract.SwapAVAXForExactTokens(&_Legacyrouter.TransactOpts, _amountOut, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapAVAXForExactTokens is a paid mutator transaction binding the contract method 0x42f564a0.
//
// Solidity: function swapAVAXForExactTokens(uint256 _amountOut, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) payable returns(uint256[] amountsIn)
func (_Legacyrouter *LegacyrouterTransactorSession) SwapAVAXForExactTokens(_amountOut *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.Contract.SwapAVAXForExactTokens(&_Legacyrouter.TransactOpts, _amountOut, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactAVAXForTokens is a paid mutator transaction binding the contract method 0x264bb94e.
//
// Solidity: function swapExactAVAXForTokens(uint256 _amountOutMin, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) payable returns(uint256 amountOut)
func (_Legacyrouter *LegacyrouterTransactor) SwapExactAVAXForTokens(opts *bind.TransactOpts, _amountOutMin *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.contract.Transact(opts, "swapExactAVAXForTokens", _amountOutMin, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactAVAXForTokens is a paid mutator transaction binding the contract method 0x264bb94e.
//
// Solidity: function swapExactAVAXForTokens(uint256 _amountOutMin, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) payable returns(uint256 amountOut)
func (_Legacyrouter *LegacyrouterSession) SwapExactAVAXForTokens(_amountOutMin *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.Contract.SwapExactAVAXForTokens(&_Legacyrouter.TransactOpts, _amountOutMin, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactAVAXForTokens is a paid mutator transaction binding the contract method 0x264bb94e.
//
// Solidity: function swapExactAVAXForTokens(uint256 _amountOutMin, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) payable returns(uint256 amountOut)
func (_Legacyrouter *LegacyrouterTransactorSession) SwapExactAVAXForTokens(_amountOutMin *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.Contract.SwapExactAVAXForTokens(&_Legacyrouter.TransactOpts, _amountOutMin, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactAVAXForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x440830bd.
//
// Solidity: function swapExactAVAXForTokensSupportingFeeOnTransferTokens(uint256 _amountOutMin, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) payable returns(uint256 amountOut)
func (_Legacyrouter *LegacyrouterTransactor) SwapExactAVAXForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, _amountOutMin *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.contract.Transact(opts, "swapExactAVAXForTokensSupportingFeeOnTransferTokens", _amountOutMin, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactAVAXForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x440830bd.
//
// Solidity: function swapExactAVAXForTokensSupportingFeeOnTransferTokens(uint256 _amountOutMin, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) payable returns(uint256 amountOut)
func (_Legacyrouter *LegacyrouterSession) SwapExactAVAXForTokensSupportingFeeOnTransferTokens(_amountOutMin *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.Contract.SwapExactAVAXForTokensSupportingFeeOnTransferTokens(&_Legacyrouter.TransactOpts, _amountOutMin, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactAVAXForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x440830bd.
//
// Solidity: function swapExactAVAXForTokensSupportingFeeOnTransferTokens(uint256 _amountOutMin, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) payable returns(uint256 amountOut)
func (_Legacyrouter *LegacyrouterTransactorSession) SwapExactAVAXForTokensSupportingFeeOnTransferTokens(_amountOutMin *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.Contract.SwapExactAVAXForTokensSupportingFeeOnTransferTokens(&_Legacyrouter.TransactOpts, _amountOutMin, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactTokensForAVAX is a paid mutator transaction binding the contract method 0xfb321c70.
//
// Solidity: function swapExactTokensForAVAX(uint256 _amountIn, uint256 _amountOutMinAVAX, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256 amountOut)
func (_Legacyrouter *LegacyrouterTransactor) SwapExactTokensForAVAX(opts *bind.TransactOpts, _amountIn *big.Int, _amountOutMinAVAX *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.contract.Transact(opts, "swapExactTokensForAVAX", _amountIn, _amountOutMinAVAX, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactTokensForAVAX is a paid mutator transaction binding the contract method 0xfb321c70.
//
// Solidity: function swapExactTokensForAVAX(uint256 _amountIn, uint256 _amountOutMinAVAX, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256 amountOut)
func (_Legacyrouter *LegacyrouterSession) SwapExactTokensForAVAX(_amountIn *big.Int, _amountOutMinAVAX *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.Contract.SwapExactTokensForAVAX(&_Legacyrouter.TransactOpts, _amountIn, _amountOutMinAVAX, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactTokensForAVAX is a paid mutator transaction binding the contract method 0xfb321c70.
//
// Solidity: function swapExactTokensForAVAX(uint256 _amountIn, uint256 _amountOutMinAVAX, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256 amountOut)
func (_Legacyrouter *LegacyrouterTransactorSession) SwapExactTokensForAVAX(_amountIn *big.Int, _amountOutMinAVAX *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.Contract.SwapExactTokensForAVAX(&_Legacyrouter.TransactOpts, _amountIn, _amountOutMinAVAX, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactTokensForAVAXSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x9a17e820.
//
// Solidity: function swapExactTokensForAVAXSupportingFeeOnTransferTokens(uint256 _amountIn, uint256 _amountOutMinAVAX, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256 amountOut)
func (_Legacyrouter *LegacyrouterTransactor) SwapExactTokensForAVAXSupportingFeeOnTransferTokens(opts *bind.TransactOpts, _amountIn *big.Int, _amountOutMinAVAX *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.contract.Transact(opts, "swapExactTokensForAVAXSupportingFeeOnTransferTokens", _amountIn, _amountOutMinAVAX, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactTokensForAVAXSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x9a17e820.
//
// Solidity: function swapExactTokensForAVAXSupportingFeeOnTransferTokens(uint256 _amountIn, uint256 _amountOutMinAVAX, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256 amountOut)
func (_Legacyrouter *LegacyrouterSession) SwapExactTokensForAVAXSupportingFeeOnTransferTokens(_amountIn *big.Int, _amountOutMinAVAX *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.Contract.SwapExactTokensForAVAXSupportingFeeOnTransferTokens(&_Legacyrouter.TransactOpts, _amountIn, _amountOutMinAVAX, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactTokensForAVAXSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x9a17e820.
//
// Solidity: function swapExactTokensForAVAXSupportingFeeOnTransferTokens(uint256 _amountIn, uint256 _amountOutMinAVAX, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256 amountOut)
func (_Legacyrouter *LegacyrouterTransactorSession) SwapExactTokensForAVAXSupportingFeeOnTransferTokens(_amountIn *big.Int, _amountOutMinAVAX *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.Contract.SwapExactTokensForAVAXSupportingFeeOnTransferTokens(&_Legacyrouter.TransactOpts, _amountIn, _amountOutMinAVAX, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x6d0ff495.
//
// Solidity: function swapExactTokensForTokens(uint256 _amountIn, uint256 _amountOutMin, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256 amountOut)
func (_Legacyrouter *LegacyrouterTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, _amountIn *big.Int, _amountOutMin *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.contract.Transact(opts, "swapExactTokensForTokens", _amountIn, _amountOutMin, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x6d0ff495.
//
// Solidity: function swapExactTokensForTokens(uint256 _amountIn, uint256 _amountOutMin, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256 amountOut)
func (_Legacyrouter *LegacyrouterSession) SwapExactTokensForTokens(_amountIn *big.Int, _amountOutMin *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.Contract.SwapExactTokensForTokens(&_Legacyrouter.TransactOpts, _amountIn, _amountOutMin, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x6d0ff495.
//
// Solidity: function swapExactTokensForTokens(uint256 _amountIn, uint256 _amountOutMin, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256 amountOut)
func (_Legacyrouter *LegacyrouterTransactorSession) SwapExactTokensForTokens(_amountIn *big.Int, _amountOutMin *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.Contract.SwapExactTokensForTokens(&_Legacyrouter.TransactOpts, _amountIn, _amountOutMin, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x212a1d94.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 _amountIn, uint256 _amountOutMin, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256 amountOut)
func (_Legacyrouter *LegacyrouterTransactor) SwapExactTokensForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, _amountIn *big.Int, _amountOutMin *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.contract.Transact(opts, "swapExactTokensForTokensSupportingFeeOnTransferTokens", _amountIn, _amountOutMin, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x212a1d94.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 _amountIn, uint256 _amountOutMin, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256 amountOut)
func (_Legacyrouter *LegacyrouterSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(_amountIn *big.Int, _amountOutMin *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_Legacyrouter.TransactOpts, _amountIn, _amountOutMin, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x212a1d94.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 _amountIn, uint256 _amountOutMin, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256 amountOut)
func (_Legacyrouter *LegacyrouterTransactorSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(_amountIn *big.Int, _amountOutMin *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_Legacyrouter.TransactOpts, _amountIn, _amountOutMin, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapTokensForExactAVAX is a paid mutator transaction binding the contract method 0x6d3420ed.
//
// Solidity: function swapTokensForExactAVAX(uint256 _amountAVAXOut, uint256 _amountInMax, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256[] amountsIn)
func (_Legacyrouter *LegacyrouterTransactor) SwapTokensForExactAVAX(opts *bind.TransactOpts, _amountAVAXOut *big.Int, _amountInMax *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.contract.Transact(opts, "swapTokensForExactAVAX", _amountAVAXOut, _amountInMax, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapTokensForExactAVAX is a paid mutator transaction binding the contract method 0x6d3420ed.
//
// Solidity: function swapTokensForExactAVAX(uint256 _amountAVAXOut, uint256 _amountInMax, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256[] amountsIn)
func (_Legacyrouter *LegacyrouterSession) SwapTokensForExactAVAX(_amountAVAXOut *big.Int, _amountInMax *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.Contract.SwapTokensForExactAVAX(&_Legacyrouter.TransactOpts, _amountAVAXOut, _amountInMax, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapTokensForExactAVAX is a paid mutator transaction binding the contract method 0x6d3420ed.
//
// Solidity: function swapTokensForExactAVAX(uint256 _amountAVAXOut, uint256 _amountInMax, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256[] amountsIn)
func (_Legacyrouter *LegacyrouterTransactorSession) SwapTokensForExactAVAX(_amountAVAXOut *big.Int, _amountInMax *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.Contract.SwapTokensForExactAVAX(&_Legacyrouter.TransactOpts, _amountAVAXOut, _amountInMax, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0xa7b856d3.
//
// Solidity: function swapTokensForExactTokens(uint256 _amountOut, uint256 _amountInMax, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256[] amountsIn)
func (_Legacyrouter *LegacyrouterTransactor) SwapTokensForExactTokens(opts *bind.TransactOpts, _amountOut *big.Int, _amountInMax *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.contract.Transact(opts, "swapTokensForExactTokens", _amountOut, _amountInMax, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0xa7b856d3.
//
// Solidity: function swapTokensForExactTokens(uint256 _amountOut, uint256 _amountInMax, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256[] amountsIn)
func (_Legacyrouter *LegacyrouterSession) SwapTokensForExactTokens(_amountOut *big.Int, _amountInMax *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.Contract.SwapTokensForExactTokens(&_Legacyrouter.TransactOpts, _amountOut, _amountInMax, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0xa7b856d3.
//
// Solidity: function swapTokensForExactTokens(uint256 _amountOut, uint256 _amountInMax, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256[] amountsIn)
func (_Legacyrouter *LegacyrouterTransactorSession) SwapTokensForExactTokens(_amountOut *big.Int, _amountInMax *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.Contract.SwapTokensForExactTokens(&_Legacyrouter.TransactOpts, _amountOut, _amountInMax, _pairBinSteps, _tokenPath, _to, _deadline)
}

// Sweep is a paid mutator transaction binding the contract method 0x62c06767.
//
// Solidity: function sweep(address _token, address _to, uint256 _amount) returns()
func (_Legacyrouter *LegacyrouterTransactor) Sweep(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.contract.Transact(opts, "sweep", _token, _to, _amount)
}

// Sweep is a paid mutator transaction binding the contract method 0x62c06767.
//
// Solidity: function sweep(address _token, address _to, uint256 _amount) returns()
func (_Legacyrouter *LegacyrouterSession) Sweep(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.Contract.Sweep(&_Legacyrouter.TransactOpts, _token, _to, _amount)
}

// Sweep is a paid mutator transaction binding the contract method 0x62c06767.
//
// Solidity: function sweep(address _token, address _to, uint256 _amount) returns()
func (_Legacyrouter *LegacyrouterTransactorSession) Sweep(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Legacyrouter.Contract.Sweep(&_Legacyrouter.TransactOpts, _token, _to, _amount)
}

// SweepLBToken is a paid mutator transaction binding the contract method 0xe9361c08.
//
// Solidity: function sweepLBToken(address _lbToken, address _to, uint256[] _ids, uint256[] _amounts) returns()
func (_Legacyrouter *LegacyrouterTransactor) SweepLBToken(opts *bind.TransactOpts, _lbToken common.Address, _to common.Address, _ids []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _Legacyrouter.contract.Transact(opts, "sweepLBToken", _lbToken, _to, _ids, _amounts)
}

// SweepLBToken is a paid mutator transaction binding the contract method 0xe9361c08.
//
// Solidity: function sweepLBToken(address _lbToken, address _to, uint256[] _ids, uint256[] _amounts) returns()
func (_Legacyrouter *LegacyrouterSession) SweepLBToken(_lbToken common.Address, _to common.Address, _ids []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _Legacyrouter.Contract.SweepLBToken(&_Legacyrouter.TransactOpts, _lbToken, _to, _ids, _amounts)
}

// SweepLBToken is a paid mutator transaction binding the contract method 0xe9361c08.
//
// Solidity: function sweepLBToken(address _lbToken, address _to, uint256[] _ids, uint256[] _amounts) returns()
func (_Legacyrouter *LegacyrouterTransactorSession) SweepLBToken(_lbToken common.Address, _to common.Address, _ids []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _Legacyrouter.Contract.SweepLBToken(&_Legacyrouter.TransactOpts, _lbToken, _to, _ids, _amounts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Legacyrouter *LegacyrouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Legacyrouter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Legacyrouter *LegacyrouterSession) Receive() (*types.Transaction, error) {
	return _Legacyrouter.Contract.Receive(&_Legacyrouter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Legacyrouter *LegacyrouterTransactorSession) Receive() (*types.Transaction, error) {
	return _Legacyrouter.Contract.Receive(&_Legacyrouter.TransactOpts)
}
