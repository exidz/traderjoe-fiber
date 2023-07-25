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

// ILBRouterLiquidityParameters is an auto generated low-level Go binding around an user-defined struct.
type ILBRouterLiquidityParameters struct {
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
	RefundTo        common.Address
	Deadline        *big.Int
}

// ILBRouterPath is an auto generated low-level Go binding around an user-defined struct.
type ILBRouterPath struct {
	PairBinSteps []*big.Int
	Versions     []uint8
	TokenPath    []common.Address
}

// LBRouterMetaData contains all meta data concerning the LBRouter contract.
var LBRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractILBFactory\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"contractIJoeFactory\",\"name\":\"factoryV1\",\"type\":\"address\"},{\"internalType\":\"contractILBLegacyFactory\",\"name\":\"legacyFactory\",\"type\":\"address\"},{\"internalType\":\"contractILBLegacyRouter\",\"name\":\"legacyRouter\",\"type\":\"address\"},{\"internalType\":\"contractIWNATIVE\",\"name\":\"wnative\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressHelper__CallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressHelper__NonContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"JoeLibrary__InsufficientAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"JoeLibrary__InsufficientLiquidity\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountSlippage\",\"type\":\"uint256\"}],\"name\":\"LBRouter__AmountSlippageBPTooBig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountXMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountYMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"name\":\"LBRouter__AmountSlippageCaught\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"LBRouter__BinReserveOverflows\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBRouter__BrokenSwapSafetyCheck\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentTimestamp\",\"type\":\"uint256\"}],\"name\":\"LBRouter__DeadlineExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LBRouter__FailedToSendNATIVE\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idSlippage\",\"type\":\"uint256\"}],\"name\":\"LBRouter__IdDesiredOverflows\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"id\",\"type\":\"int256\"}],\"name\":\"LBRouter__IdOverflows\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"activeIdDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idSlippage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activeId\",\"type\":\"uint256\"}],\"name\":\"LBRouter__IdSlippageCaught\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"LBRouter__InsufficientAmountOut\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wrongToken\",\"type\":\"address\"}],\"name\":\"LBRouter__InvalidTokenPath\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"LBRouter__InvalidVersion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBRouter__LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"LBRouter__MaxAmountInExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBRouter__NotFactoryOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"}],\"name\":\"LBRouter__PairNotCreated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBRouter__SenderIsNotWNATIVE\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"LBRouter__SwapOverflows\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"excess\",\"type\":\"uint256\"}],\"name\":\"LBRouter__TooMuchTokensIn\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserve\",\"type\":\"uint256\"}],\"name\":\"LBRouter__WrongAmounts\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"}],\"name\":\"LBRouter__WrongNativeLiquidityParameters\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBRouter__WrongTokenOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenHelper__TransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountXMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountYMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activeIdDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idSlippage\",\"type\":\"uint256\"},{\"internalType\":\"int256[]\",\"name\":\"deltaIds\",\"type\":\"int256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"distributionX\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"distributionY\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structILBRouter.LiquidityParameters\",\"name\":\"liquidityParameters\",\"type\":\"tuple\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountXAdded\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountYAdded\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountXLeft\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountYLeft\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"depositIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"liquidityMinted\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountXMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountYMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activeIdDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idSlippage\",\"type\":\"uint256\"},{\"internalType\":\"int256[]\",\"name\":\"deltaIds\",\"type\":\"int256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"distributionX\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"distributionY\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structILBRouter.LiquidityParameters\",\"name\":\"liquidityParameters\",\"type\":\"tuple\"}],\"name\":\"addLiquidityNATIVE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountXAdded\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountYAdded\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountXLeft\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountYLeft\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"depositIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"liquidityMinted\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"activeId\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"binStep\",\"type\":\"uint16\"}],\"name\":\"createLBPair\",\"outputs\":[{\"internalType\":\"contractILBPair\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFactory\",\"outputs\":[{\"internalType\":\"contractILBFactory\",\"name\":\"lbFactory\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILBPair\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"getIdFromPrice\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLegacyFactory\",\"outputs\":[{\"internalType\":\"contractILBLegacyFactory\",\"name\":\"legacyLBfactory\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLegacyRouter\",\"outputs\":[{\"internalType\":\"contractILBLegacyRouter\",\"name\":\"legacyRouter\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILBPair\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"id\",\"type\":\"uint24\"}],\"name\":\"getPriceFromId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILBPair\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amountOut\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"swapForY\",\"type\":\"bool\"}],\"name\":\"getSwapIn\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amountIn\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amountOutLeft\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"fee\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILBPair\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amountIn\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"swapForY\",\"type\":\"bool\"}],\"name\":\"getSwapOut\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amountInLeft\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amountOut\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"fee\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getV1Factory\",\"outputs\":[{\"internalType\":\"contractIJoeFactory\",\"name\":\"factoryV1\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWNATIVE\",\"outputs\":[{\"internalType\":\"contractIWNATIVE\",\"name\":\"wnative\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"binStep\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"amountXMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountYMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"binStep\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountNATIVEMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityNATIVE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountNATIVE\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"enumILBRouter.Version[]\",\"name\":\"versions\",\"type\":\"uint8[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"}],\"internalType\":\"structILBRouter.Path\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactNATIVEForTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"enumILBRouter.Version[]\",\"name\":\"versions\",\"type\":\"uint8[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"}],\"internalType\":\"structILBRouter.Path\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactNATIVEForTokensSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinNATIVE\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"enumILBRouter.Version[]\",\"name\":\"versions\",\"type\":\"uint8[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"}],\"internalType\":\"structILBRouter.Path\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForNATIVE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinNATIVE\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"enumILBRouter.Version[]\",\"name\":\"versions\",\"type\":\"uint8[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"}],\"internalType\":\"structILBRouter.Path\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForNATIVESupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"enumILBRouter.Version[]\",\"name\":\"versions\",\"type\":\"uint8[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"}],\"internalType\":\"structILBRouter.Path\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"enumILBRouter.Version[]\",\"name\":\"versions\",\"type\":\"uint8[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"}],\"internalType\":\"structILBRouter.Path\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"enumILBRouter.Version[]\",\"name\":\"versions\",\"type\":\"uint8[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"}],\"internalType\":\"structILBRouter.Path\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapNATIVEForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountsIn\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountNATIVEOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"enumILBRouter.Version[]\",\"name\":\"versions\",\"type\":\"uint8[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"}],\"internalType\":\"structILBRouter.Path\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactNATIVE\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountsIn\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"enumILBRouter.Version[]\",\"name\":\"versions\",\"type\":\"uint8[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"}],\"internalType\":\"structILBRouter.Path\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountsIn\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sweep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILBToken\",\"name\":\"lbToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"sweepLBToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// LBRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use LBRouterMetaData.ABI instead.
var LBRouterABI = LBRouterMetaData.ABI

// LBRouter is an auto generated Go binding around an Ethereum contract.
type LBRouter struct {
	LBRouterCaller     // Read-only binding to the contract
	LBRouterTransactor // Write-only binding to the contract
	LBRouterFilterer   // Log filterer for contract events
}

// LBRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type LBRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LBRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LBRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LBRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LBRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LBRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LBRouterSession struct {
	Contract     *LBRouter         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LBRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LBRouterCallerSession struct {
	Contract *LBRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// LBRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LBRouterTransactorSession struct {
	Contract     *LBRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LBRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type LBRouterRaw struct {
	Contract *LBRouter // Generic contract binding to access the raw methods on
}

// LBRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LBRouterCallerRaw struct {
	Contract *LBRouterCaller // Generic read-only contract binding to access the raw methods on
}

// LBRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LBRouterTransactorRaw struct {
	Contract *LBRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLBRouter creates a new instance of LBRouter, bound to a specific deployed contract.
func NewLBRouter(address common.Address, backend bind.ContractBackend) (*LBRouter, error) {
	contract, err := bindLBRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LBRouter{LBRouterCaller: LBRouterCaller{contract: contract}, LBRouterTransactor: LBRouterTransactor{contract: contract}, LBRouterFilterer: LBRouterFilterer{contract: contract}}, nil
}

// NewLBRouterCaller creates a new read-only instance of LBRouter, bound to a specific deployed contract.
func NewLBRouterCaller(address common.Address, caller bind.ContractCaller) (*LBRouterCaller, error) {
	contract, err := bindLBRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LBRouterCaller{contract: contract}, nil
}

// NewLBRouterTransactor creates a new write-only instance of LBRouter, bound to a specific deployed contract.
func NewLBRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*LBRouterTransactor, error) {
	contract, err := bindLBRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LBRouterTransactor{contract: contract}, nil
}

// NewLBRouterFilterer creates a new log filterer instance of LBRouter, bound to a specific deployed contract.
func NewLBRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*LBRouterFilterer, error) {
	contract, err := bindLBRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LBRouterFilterer{contract: contract}, nil
}

// bindLBRouter binds a generic wrapper to an already deployed contract.
func bindLBRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LBRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LBRouter *LBRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LBRouter.Contract.LBRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LBRouter *LBRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LBRouter.Contract.LBRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LBRouter *LBRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LBRouter.Contract.LBRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LBRouter *LBRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LBRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LBRouter *LBRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LBRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LBRouter *LBRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LBRouter.Contract.contract.Transact(opts, method, params...)
}

// GetFactory is a free data retrieval call binding the contract method 0x88cc58e4.
//
// Solidity: function getFactory() view returns(address lbFactory)
func (_LBRouter *LBRouterCaller) GetFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LBRouter.contract.Call(opts, &out, "getFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFactory is a free data retrieval call binding the contract method 0x88cc58e4.
//
// Solidity: function getFactory() view returns(address lbFactory)
func (_LBRouter *LBRouterSession) GetFactory() (common.Address, error) {
	return _LBRouter.Contract.GetFactory(&_LBRouter.CallOpts)
}

// GetFactory is a free data retrieval call binding the contract method 0x88cc58e4.
//
// Solidity: function getFactory() view returns(address lbFactory)
func (_LBRouter *LBRouterCallerSession) GetFactory() (common.Address, error) {
	return _LBRouter.Contract.GetFactory(&_LBRouter.CallOpts)
}

// GetIdFromPrice is a free data retrieval call binding the contract method 0xf96fe925.
//
// Solidity: function getIdFromPrice(address pair, uint256 price) view returns(uint24)
func (_LBRouter *LBRouterCaller) GetIdFromPrice(opts *bind.CallOpts, pair common.Address, price *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LBRouter.contract.Call(opts, &out, "getIdFromPrice", pair, price)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIdFromPrice is a free data retrieval call binding the contract method 0xf96fe925.
//
// Solidity: function getIdFromPrice(address pair, uint256 price) view returns(uint24)
func (_LBRouter *LBRouterSession) GetIdFromPrice(pair common.Address, price *big.Int) (*big.Int, error) {
	return _LBRouter.Contract.GetIdFromPrice(&_LBRouter.CallOpts, pair, price)
}

// GetIdFromPrice is a free data retrieval call binding the contract method 0xf96fe925.
//
// Solidity: function getIdFromPrice(address pair, uint256 price) view returns(uint24)
func (_LBRouter *LBRouterCallerSession) GetIdFromPrice(pair common.Address, price *big.Int) (*big.Int, error) {
	return _LBRouter.Contract.GetIdFromPrice(&_LBRouter.CallOpts, pair, price)
}

// GetLegacyFactory is a free data retrieval call binding the contract method 0x71d1974a.
//
// Solidity: function getLegacyFactory() view returns(address legacyLBfactory)
func (_LBRouter *LBRouterCaller) GetLegacyFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LBRouter.contract.Call(opts, &out, "getLegacyFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLegacyFactory is a free data retrieval call binding the contract method 0x71d1974a.
//
// Solidity: function getLegacyFactory() view returns(address legacyLBfactory)
func (_LBRouter *LBRouterSession) GetLegacyFactory() (common.Address, error) {
	return _LBRouter.Contract.GetLegacyFactory(&_LBRouter.CallOpts)
}

// GetLegacyFactory is a free data retrieval call binding the contract method 0x71d1974a.
//
// Solidity: function getLegacyFactory() view returns(address legacyLBfactory)
func (_LBRouter *LBRouterCallerSession) GetLegacyFactory() (common.Address, error) {
	return _LBRouter.Contract.GetLegacyFactory(&_LBRouter.CallOpts)
}

// GetLegacyRouter is a free data retrieval call binding the contract method 0xba846523.
//
// Solidity: function getLegacyRouter() view returns(address legacyRouter)
func (_LBRouter *LBRouterCaller) GetLegacyRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LBRouter.contract.Call(opts, &out, "getLegacyRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLegacyRouter is a free data retrieval call binding the contract method 0xba846523.
//
// Solidity: function getLegacyRouter() view returns(address legacyRouter)
func (_LBRouter *LBRouterSession) GetLegacyRouter() (common.Address, error) {
	return _LBRouter.Contract.GetLegacyRouter(&_LBRouter.CallOpts)
}

// GetLegacyRouter is a free data retrieval call binding the contract method 0xba846523.
//
// Solidity: function getLegacyRouter() view returns(address legacyRouter)
func (_LBRouter *LBRouterCallerSession) GetLegacyRouter() (common.Address, error) {
	return _LBRouter.Contract.GetLegacyRouter(&_LBRouter.CallOpts)
}

// GetPriceFromId is a free data retrieval call binding the contract method 0xd0e380f2.
//
// Solidity: function getPriceFromId(address pair, uint24 id) view returns(uint256)
func (_LBRouter *LBRouterCaller) GetPriceFromId(opts *bind.CallOpts, pair common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LBRouter.contract.Call(opts, &out, "getPriceFromId", pair, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriceFromId is a free data retrieval call binding the contract method 0xd0e380f2.
//
// Solidity: function getPriceFromId(address pair, uint24 id) view returns(uint256)
func (_LBRouter *LBRouterSession) GetPriceFromId(pair common.Address, id *big.Int) (*big.Int, error) {
	return _LBRouter.Contract.GetPriceFromId(&_LBRouter.CallOpts, pair, id)
}

// GetPriceFromId is a free data retrieval call binding the contract method 0xd0e380f2.
//
// Solidity: function getPriceFromId(address pair, uint24 id) view returns(uint256)
func (_LBRouter *LBRouterCallerSession) GetPriceFromId(pair common.Address, id *big.Int) (*big.Int, error) {
	return _LBRouter.Contract.GetPriceFromId(&_LBRouter.CallOpts, pair, id)
}

// GetSwapIn is a free data retrieval call binding the contract method 0x964f987c.
//
// Solidity: function getSwapIn(address pair, uint128 amountOut, bool swapForY) view returns(uint128 amountIn, uint128 amountOutLeft, uint128 fee)
func (_LBRouter *LBRouterCaller) GetSwapIn(opts *bind.CallOpts, pair common.Address, amountOut *big.Int, swapForY bool) (struct {
	AmountIn      *big.Int
	AmountOutLeft *big.Int
	Fee           *big.Int
}, error) {
	var out []interface{}
	err := _LBRouter.contract.Call(opts, &out, "getSwapIn", pair, amountOut, swapForY)

	outstruct := new(struct {
		AmountIn      *big.Int
		AmountOutLeft *big.Int
		Fee           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountIn = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountOutLeft = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSwapIn is a free data retrieval call binding the contract method 0x964f987c.
//
// Solidity: function getSwapIn(address pair, uint128 amountOut, bool swapForY) view returns(uint128 amountIn, uint128 amountOutLeft, uint128 fee)
func (_LBRouter *LBRouterSession) GetSwapIn(pair common.Address, amountOut *big.Int, swapForY bool) (struct {
	AmountIn      *big.Int
	AmountOutLeft *big.Int
	Fee           *big.Int
}, error) {
	return _LBRouter.Contract.GetSwapIn(&_LBRouter.CallOpts, pair, amountOut, swapForY)
}

// GetSwapIn is a free data retrieval call binding the contract method 0x964f987c.
//
// Solidity: function getSwapIn(address pair, uint128 amountOut, bool swapForY) view returns(uint128 amountIn, uint128 amountOutLeft, uint128 fee)
func (_LBRouter *LBRouterCallerSession) GetSwapIn(pair common.Address, amountOut *big.Int, swapForY bool) (struct {
	AmountIn      *big.Int
	AmountOutLeft *big.Int
	Fee           *big.Int
}, error) {
	return _LBRouter.Contract.GetSwapIn(&_LBRouter.CallOpts, pair, amountOut, swapForY)
}

// GetSwapOut is a free data retrieval call binding the contract method 0xa0d376cf.
//
// Solidity: function getSwapOut(address pair, uint128 amountIn, bool swapForY) view returns(uint128 amountInLeft, uint128 amountOut, uint128 fee)
func (_LBRouter *LBRouterCaller) GetSwapOut(opts *bind.CallOpts, pair common.Address, amountIn *big.Int, swapForY bool) (struct {
	AmountInLeft *big.Int
	AmountOut    *big.Int
	Fee          *big.Int
}, error) {
	var out []interface{}
	err := _LBRouter.contract.Call(opts, &out, "getSwapOut", pair, amountIn, swapForY)

	outstruct := new(struct {
		AmountInLeft *big.Int
		AmountOut    *big.Int
		Fee          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountInLeft = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountOut = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSwapOut is a free data retrieval call binding the contract method 0xa0d376cf.
//
// Solidity: function getSwapOut(address pair, uint128 amountIn, bool swapForY) view returns(uint128 amountInLeft, uint128 amountOut, uint128 fee)
func (_LBRouter *LBRouterSession) GetSwapOut(pair common.Address, amountIn *big.Int, swapForY bool) (struct {
	AmountInLeft *big.Int
	AmountOut    *big.Int
	Fee          *big.Int
}, error) {
	return _LBRouter.Contract.GetSwapOut(&_LBRouter.CallOpts, pair, amountIn, swapForY)
}

// GetSwapOut is a free data retrieval call binding the contract method 0xa0d376cf.
//
// Solidity: function getSwapOut(address pair, uint128 amountIn, bool swapForY) view returns(uint128 amountInLeft, uint128 amountOut, uint128 fee)
func (_LBRouter *LBRouterCallerSession) GetSwapOut(pair common.Address, amountIn *big.Int, swapForY bool) (struct {
	AmountInLeft *big.Int
	AmountOut    *big.Int
	Fee          *big.Int
}, error) {
	return _LBRouter.Contract.GetSwapOut(&_LBRouter.CallOpts, pair, amountIn, swapForY)
}

// GetV1Factory is a free data retrieval call binding the contract method 0xbb558a9f.
//
// Solidity: function getV1Factory() view returns(address factoryV1)
func (_LBRouter *LBRouterCaller) GetV1Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LBRouter.contract.Call(opts, &out, "getV1Factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetV1Factory is a free data retrieval call binding the contract method 0xbb558a9f.
//
// Solidity: function getV1Factory() view returns(address factoryV1)
func (_LBRouter *LBRouterSession) GetV1Factory() (common.Address, error) {
	return _LBRouter.Contract.GetV1Factory(&_LBRouter.CallOpts)
}

// GetV1Factory is a free data retrieval call binding the contract method 0xbb558a9f.
//
// Solidity: function getV1Factory() view returns(address factoryV1)
func (_LBRouter *LBRouterCallerSession) GetV1Factory() (common.Address, error) {
	return _LBRouter.Contract.GetV1Factory(&_LBRouter.CallOpts)
}

// GetWNATIVE is a free data retrieval call binding the contract method 0x6c9c0078.
//
// Solidity: function getWNATIVE() view returns(address wnative)
func (_LBRouter *LBRouterCaller) GetWNATIVE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LBRouter.contract.Call(opts, &out, "getWNATIVE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWNATIVE is a free data retrieval call binding the contract method 0x6c9c0078.
//
// Solidity: function getWNATIVE() view returns(address wnative)
func (_LBRouter *LBRouterSession) GetWNATIVE() (common.Address, error) {
	return _LBRouter.Contract.GetWNATIVE(&_LBRouter.CallOpts)
}

// GetWNATIVE is a free data retrieval call binding the contract method 0x6c9c0078.
//
// Solidity: function getWNATIVE() view returns(address wnative)
func (_LBRouter *LBRouterCallerSession) GetWNATIVE() (common.Address, error) {
	return _LBRouter.Contract.GetWNATIVE(&_LBRouter.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xa3c7271a.
//
// Solidity: function addLiquidity((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256[],uint256[],uint256[],address,address,uint256) liquidityParameters) returns(uint256 amountXAdded, uint256 amountYAdded, uint256 amountXLeft, uint256 amountYLeft, uint256[] depositIds, uint256[] liquidityMinted)
func (_LBRouter *LBRouterTransactor) AddLiquidity(opts *bind.TransactOpts, liquidityParameters ILBRouterLiquidityParameters) (*types.Transaction, error) {
	return _LBRouter.contract.Transact(opts, "addLiquidity", liquidityParameters)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xa3c7271a.
//
// Solidity: function addLiquidity((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256[],uint256[],uint256[],address,address,uint256) liquidityParameters) returns(uint256 amountXAdded, uint256 amountYAdded, uint256 amountXLeft, uint256 amountYLeft, uint256[] depositIds, uint256[] liquidityMinted)
func (_LBRouter *LBRouterSession) AddLiquidity(liquidityParameters ILBRouterLiquidityParameters) (*types.Transaction, error) {
	return _LBRouter.Contract.AddLiquidity(&_LBRouter.TransactOpts, liquidityParameters)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xa3c7271a.
//
// Solidity: function addLiquidity((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256[],uint256[],uint256[],address,address,uint256) liquidityParameters) returns(uint256 amountXAdded, uint256 amountYAdded, uint256 amountXLeft, uint256 amountYLeft, uint256[] depositIds, uint256[] liquidityMinted)
func (_LBRouter *LBRouterTransactorSession) AddLiquidity(liquidityParameters ILBRouterLiquidityParameters) (*types.Transaction, error) {
	return _LBRouter.Contract.AddLiquidity(&_LBRouter.TransactOpts, liquidityParameters)
}

// AddLiquidityNATIVE is a paid mutator transaction binding the contract method 0x8efc2b2c.
//
// Solidity: function addLiquidityNATIVE((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256[],uint256[],uint256[],address,address,uint256) liquidityParameters) payable returns(uint256 amountXAdded, uint256 amountYAdded, uint256 amountXLeft, uint256 amountYLeft, uint256[] depositIds, uint256[] liquidityMinted)
func (_LBRouter *LBRouterTransactor) AddLiquidityNATIVE(opts *bind.TransactOpts, liquidityParameters ILBRouterLiquidityParameters) (*types.Transaction, error) {
	return _LBRouter.contract.Transact(opts, "addLiquidityNATIVE", liquidityParameters)
}

// AddLiquidityNATIVE is a paid mutator transaction binding the contract method 0x8efc2b2c.
//
// Solidity: function addLiquidityNATIVE((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256[],uint256[],uint256[],address,address,uint256) liquidityParameters) payable returns(uint256 amountXAdded, uint256 amountYAdded, uint256 amountXLeft, uint256 amountYLeft, uint256[] depositIds, uint256[] liquidityMinted)
func (_LBRouter *LBRouterSession) AddLiquidityNATIVE(liquidityParameters ILBRouterLiquidityParameters) (*types.Transaction, error) {
	return _LBRouter.Contract.AddLiquidityNATIVE(&_LBRouter.TransactOpts, liquidityParameters)
}

// AddLiquidityNATIVE is a paid mutator transaction binding the contract method 0x8efc2b2c.
//
// Solidity: function addLiquidityNATIVE((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256[],uint256[],uint256[],address,address,uint256) liquidityParameters) payable returns(uint256 amountXAdded, uint256 amountYAdded, uint256 amountXLeft, uint256 amountYLeft, uint256[] depositIds, uint256[] liquidityMinted)
func (_LBRouter *LBRouterTransactorSession) AddLiquidityNATIVE(liquidityParameters ILBRouterLiquidityParameters) (*types.Transaction, error) {
	return _LBRouter.Contract.AddLiquidityNATIVE(&_LBRouter.TransactOpts, liquidityParameters)
}

// CreateLBPair is a paid mutator transaction binding the contract method 0x659ac74b.
//
// Solidity: function createLBPair(address tokenX, address tokenY, uint24 activeId, uint16 binStep) returns(address pair)
func (_LBRouter *LBRouterTransactor) CreateLBPair(opts *bind.TransactOpts, tokenX common.Address, tokenY common.Address, activeId *big.Int, binStep uint16) (*types.Transaction, error) {
	return _LBRouter.contract.Transact(opts, "createLBPair", tokenX, tokenY, activeId, binStep)
}

// CreateLBPair is a paid mutator transaction binding the contract method 0x659ac74b.
//
// Solidity: function createLBPair(address tokenX, address tokenY, uint24 activeId, uint16 binStep) returns(address pair)
func (_LBRouter *LBRouterSession) CreateLBPair(tokenX common.Address, tokenY common.Address, activeId *big.Int, binStep uint16) (*types.Transaction, error) {
	return _LBRouter.Contract.CreateLBPair(&_LBRouter.TransactOpts, tokenX, tokenY, activeId, binStep)
}

// CreateLBPair is a paid mutator transaction binding the contract method 0x659ac74b.
//
// Solidity: function createLBPair(address tokenX, address tokenY, uint24 activeId, uint16 binStep) returns(address pair)
func (_LBRouter *LBRouterTransactorSession) CreateLBPair(tokenX common.Address, tokenY common.Address, activeId *big.Int, binStep uint16) (*types.Transaction, error) {
	return _LBRouter.Contract.CreateLBPair(&_LBRouter.TransactOpts, tokenX, tokenY, activeId, binStep)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xc22159b6.
//
// Solidity: function removeLiquidity(address tokenX, address tokenY, uint16 binStep, uint256 amountXMin, uint256 amountYMin, uint256[] ids, uint256[] amounts, address to, uint256 deadline) returns(uint256 amountX, uint256 amountY)
func (_LBRouter *LBRouterTransactor) RemoveLiquidity(opts *bind.TransactOpts, tokenX common.Address, tokenY common.Address, binStep uint16, amountXMin *big.Int, amountYMin *big.Int, ids []*big.Int, amounts []*big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.contract.Transact(opts, "removeLiquidity", tokenX, tokenY, binStep, amountXMin, amountYMin, ids, amounts, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xc22159b6.
//
// Solidity: function removeLiquidity(address tokenX, address tokenY, uint16 binStep, uint256 amountXMin, uint256 amountYMin, uint256[] ids, uint256[] amounts, address to, uint256 deadline) returns(uint256 amountX, uint256 amountY)
func (_LBRouter *LBRouterSession) RemoveLiquidity(tokenX common.Address, tokenY common.Address, binStep uint16, amountXMin *big.Int, amountYMin *big.Int, ids []*big.Int, amounts []*big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.Contract.RemoveLiquidity(&_LBRouter.TransactOpts, tokenX, tokenY, binStep, amountXMin, amountYMin, ids, amounts, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xc22159b6.
//
// Solidity: function removeLiquidity(address tokenX, address tokenY, uint16 binStep, uint256 amountXMin, uint256 amountYMin, uint256[] ids, uint256[] amounts, address to, uint256 deadline) returns(uint256 amountX, uint256 amountY)
func (_LBRouter *LBRouterTransactorSession) RemoveLiquidity(tokenX common.Address, tokenY common.Address, binStep uint16, amountXMin *big.Int, amountYMin *big.Int, ids []*big.Int, amounts []*big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.Contract.RemoveLiquidity(&_LBRouter.TransactOpts, tokenX, tokenY, binStep, amountXMin, amountYMin, ids, amounts, to, deadline)
}

// RemoveLiquidityNATIVE is a paid mutator transaction binding the contract method 0x81c2fdfb.
//
// Solidity: function removeLiquidityNATIVE(address token, uint16 binStep, uint256 amountTokenMin, uint256 amountNATIVEMin, uint256[] ids, uint256[] amounts, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountNATIVE)
func (_LBRouter *LBRouterTransactor) RemoveLiquidityNATIVE(opts *bind.TransactOpts, token common.Address, binStep uint16, amountTokenMin *big.Int, amountNATIVEMin *big.Int, ids []*big.Int, amounts []*big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.contract.Transact(opts, "removeLiquidityNATIVE", token, binStep, amountTokenMin, amountNATIVEMin, ids, amounts, to, deadline)
}

// RemoveLiquidityNATIVE is a paid mutator transaction binding the contract method 0x81c2fdfb.
//
// Solidity: function removeLiquidityNATIVE(address token, uint16 binStep, uint256 amountTokenMin, uint256 amountNATIVEMin, uint256[] ids, uint256[] amounts, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountNATIVE)
func (_LBRouter *LBRouterSession) RemoveLiquidityNATIVE(token common.Address, binStep uint16, amountTokenMin *big.Int, amountNATIVEMin *big.Int, ids []*big.Int, amounts []*big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.Contract.RemoveLiquidityNATIVE(&_LBRouter.TransactOpts, token, binStep, amountTokenMin, amountNATIVEMin, ids, amounts, to, deadline)
}

// RemoveLiquidityNATIVE is a paid mutator transaction binding the contract method 0x81c2fdfb.
//
// Solidity: function removeLiquidityNATIVE(address token, uint16 binStep, uint256 amountTokenMin, uint256 amountNATIVEMin, uint256[] ids, uint256[] amounts, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountNATIVE)
func (_LBRouter *LBRouterTransactorSession) RemoveLiquidityNATIVE(token common.Address, binStep uint16, amountTokenMin *big.Int, amountNATIVEMin *big.Int, ids []*big.Int, amounts []*big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.Contract.RemoveLiquidityNATIVE(&_LBRouter.TransactOpts, token, binStep, amountTokenMin, amountNATIVEMin, ids, amounts, to, deadline)
}

// SwapExactNATIVEForTokens is a paid mutator transaction binding the contract method 0xb066ea7c.
//
// Solidity: function swapExactNATIVEForTokens(uint256 amountOutMin, (uint256[],uint8[],address[]) path, address to, uint256 deadline) payable returns(uint256 amountOut)
func (_LBRouter *LBRouterTransactor) SwapExactNATIVEForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.contract.Transact(opts, "swapExactNATIVEForTokens", amountOutMin, path, to, deadline)
}

// SwapExactNATIVEForTokens is a paid mutator transaction binding the contract method 0xb066ea7c.
//
// Solidity: function swapExactNATIVEForTokens(uint256 amountOutMin, (uint256[],uint8[],address[]) path, address to, uint256 deadline) payable returns(uint256 amountOut)
func (_LBRouter *LBRouterSession) SwapExactNATIVEForTokens(amountOutMin *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.Contract.SwapExactNATIVEForTokens(&_LBRouter.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactNATIVEForTokens is a paid mutator transaction binding the contract method 0xb066ea7c.
//
// Solidity: function swapExactNATIVEForTokens(uint256 amountOutMin, (uint256[],uint8[],address[]) path, address to, uint256 deadline) payable returns(uint256 amountOut)
func (_LBRouter *LBRouterTransactorSession) SwapExactNATIVEForTokens(amountOutMin *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.Contract.SwapExactNATIVEForTokens(&_LBRouter.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactNATIVEForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xe038e6dc.
//
// Solidity: function swapExactNATIVEForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, (uint256[],uint8[],address[]) path, address to, uint256 deadline) payable returns(uint256 amountOut)
func (_LBRouter *LBRouterTransactor) SwapExactNATIVEForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.contract.Transact(opts, "swapExactNATIVEForTokensSupportingFeeOnTransferTokens", amountOutMin, path, to, deadline)
}

// SwapExactNATIVEForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xe038e6dc.
//
// Solidity: function swapExactNATIVEForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, (uint256[],uint8[],address[]) path, address to, uint256 deadline) payable returns(uint256 amountOut)
func (_LBRouter *LBRouterSession) SwapExactNATIVEForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.Contract.SwapExactNATIVEForTokensSupportingFeeOnTransferTokens(&_LBRouter.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactNATIVEForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xe038e6dc.
//
// Solidity: function swapExactNATIVEForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, (uint256[],uint8[],address[]) path, address to, uint256 deadline) payable returns(uint256 amountOut)
func (_LBRouter *LBRouterTransactorSession) SwapExactNATIVEForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.Contract.SwapExactNATIVEForTokensSupportingFeeOnTransferTokens(&_LBRouter.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactTokensForNATIVE is a paid mutator transaction binding the contract method 0x9ab6156b.
//
// Solidity: function swapExactTokensForNATIVE(uint256 amountIn, uint256 amountOutMinNATIVE, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256 amountOut)
func (_LBRouter *LBRouterTransactor) SwapExactTokensForNATIVE(opts *bind.TransactOpts, amountIn *big.Int, amountOutMinNATIVE *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.contract.Transact(opts, "swapExactTokensForNATIVE", amountIn, amountOutMinNATIVE, path, to, deadline)
}

// SwapExactTokensForNATIVE is a paid mutator transaction binding the contract method 0x9ab6156b.
//
// Solidity: function swapExactTokensForNATIVE(uint256 amountIn, uint256 amountOutMinNATIVE, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256 amountOut)
func (_LBRouter *LBRouterSession) SwapExactTokensForNATIVE(amountIn *big.Int, amountOutMinNATIVE *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.Contract.SwapExactTokensForNATIVE(&_LBRouter.TransactOpts, amountIn, amountOutMinNATIVE, path, to, deadline)
}

// SwapExactTokensForNATIVE is a paid mutator transaction binding the contract method 0x9ab6156b.
//
// Solidity: function swapExactTokensForNATIVE(uint256 amountIn, uint256 amountOutMinNATIVE, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256 amountOut)
func (_LBRouter *LBRouterTransactorSession) SwapExactTokensForNATIVE(amountIn *big.Int, amountOutMinNATIVE *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.Contract.SwapExactTokensForNATIVE(&_LBRouter.TransactOpts, amountIn, amountOutMinNATIVE, path, to, deadline)
}

// SwapExactTokensForNATIVESupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x1a24f9a9.
//
// Solidity: function swapExactTokensForNATIVESupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMinNATIVE, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256 amountOut)
func (_LBRouter *LBRouterTransactor) SwapExactTokensForNATIVESupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMinNATIVE *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.contract.Transact(opts, "swapExactTokensForNATIVESupportingFeeOnTransferTokens", amountIn, amountOutMinNATIVE, path, to, deadline)
}

// SwapExactTokensForNATIVESupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x1a24f9a9.
//
// Solidity: function swapExactTokensForNATIVESupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMinNATIVE, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256 amountOut)
func (_LBRouter *LBRouterSession) SwapExactTokensForNATIVESupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMinNATIVE *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.Contract.SwapExactTokensForNATIVESupportingFeeOnTransferTokens(&_LBRouter.TransactOpts, amountIn, amountOutMinNATIVE, path, to, deadline)
}

// SwapExactTokensForNATIVESupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x1a24f9a9.
//
// Solidity: function swapExactTokensForNATIVESupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMinNATIVE, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256 amountOut)
func (_LBRouter *LBRouterTransactorSession) SwapExactTokensForNATIVESupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMinNATIVE *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.Contract.SwapExactTokensForNATIVESupportingFeeOnTransferTokens(&_LBRouter.TransactOpts, amountIn, amountOutMinNATIVE, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x2a443fae.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256 amountOut)
func (_LBRouter *LBRouterTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x2a443fae.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256 amountOut)
func (_LBRouter *LBRouterSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.Contract.SwapExactTokensForTokens(&_LBRouter.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x2a443fae.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256 amountOut)
func (_LBRouter *LBRouterTransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.Contract.SwapExactTokensForTokens(&_LBRouter.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x4b801870.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256 amountOut)
func (_LBRouter *LBRouterTransactor) SwapExactTokensForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.contract.Transact(opts, "swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x4b801870.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256 amountOut)
func (_LBRouter *LBRouterSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_LBRouter.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x4b801870.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256 amountOut)
func (_LBRouter *LBRouterTransactorSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_LBRouter.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapNATIVEForExactTokens is a paid mutator transaction binding the contract method 0x2075ad22.
//
// Solidity: function swapNATIVEForExactTokens(uint256 amountOut, (uint256[],uint8[],address[]) path, address to, uint256 deadline) payable returns(uint256[] amountsIn)
func (_LBRouter *LBRouterTransactor) SwapNATIVEForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.contract.Transact(opts, "swapNATIVEForExactTokens", amountOut, path, to, deadline)
}

// SwapNATIVEForExactTokens is a paid mutator transaction binding the contract method 0x2075ad22.
//
// Solidity: function swapNATIVEForExactTokens(uint256 amountOut, (uint256[],uint8[],address[]) path, address to, uint256 deadline) payable returns(uint256[] amountsIn)
func (_LBRouter *LBRouterSession) SwapNATIVEForExactTokens(amountOut *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.Contract.SwapNATIVEForExactTokens(&_LBRouter.TransactOpts, amountOut, path, to, deadline)
}

// SwapNATIVEForExactTokens is a paid mutator transaction binding the contract method 0x2075ad22.
//
// Solidity: function swapNATIVEForExactTokens(uint256 amountOut, (uint256[],uint8[],address[]) path, address to, uint256 deadline) payable returns(uint256[] amountsIn)
func (_LBRouter *LBRouterTransactorSession) SwapNATIVEForExactTokens(amountOut *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.Contract.SwapNATIVEForExactTokens(&_LBRouter.TransactOpts, amountOut, path, to, deadline)
}

// SwapTokensForExactNATIVE is a paid mutator transaction binding the contract method 0x3dc8f8ec.
//
// Solidity: function swapTokensForExactNATIVE(uint256 amountNATIVEOut, uint256 amountInMax, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256[] amountsIn)
func (_LBRouter *LBRouterTransactor) SwapTokensForExactNATIVE(opts *bind.TransactOpts, amountNATIVEOut *big.Int, amountInMax *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.contract.Transact(opts, "swapTokensForExactNATIVE", amountNATIVEOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactNATIVE is a paid mutator transaction binding the contract method 0x3dc8f8ec.
//
// Solidity: function swapTokensForExactNATIVE(uint256 amountNATIVEOut, uint256 amountInMax, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256[] amountsIn)
func (_LBRouter *LBRouterSession) SwapTokensForExactNATIVE(amountNATIVEOut *big.Int, amountInMax *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.Contract.SwapTokensForExactNATIVE(&_LBRouter.TransactOpts, amountNATIVEOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactNATIVE is a paid mutator transaction binding the contract method 0x3dc8f8ec.
//
// Solidity: function swapTokensForExactNATIVE(uint256 amountNATIVEOut, uint256 amountInMax, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256[] amountsIn)
func (_LBRouter *LBRouterTransactorSession) SwapTokensForExactNATIVE(amountNATIVEOut *big.Int, amountInMax *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.Contract.SwapTokensForExactNATIVE(&_LBRouter.TransactOpts, amountNATIVEOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x92fe8e70.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256[] amountsIn)
func (_LBRouter *LBRouterTransactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.contract.Transact(opts, "swapTokensForExactTokens", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x92fe8e70.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256[] amountsIn)
func (_LBRouter *LBRouterSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.Contract.SwapTokensForExactTokens(&_LBRouter.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x92fe8e70.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256[] amountsIn)
func (_LBRouter *LBRouterTransactorSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LBRouter.Contract.SwapTokensForExactTokens(&_LBRouter.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// Sweep is a paid mutator transaction binding the contract method 0x62c06767.
//
// Solidity: function sweep(address token, address to, uint256 amount) returns()
func (_LBRouter *LBRouterTransactor) Sweep(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LBRouter.contract.Transact(opts, "sweep", token, to, amount)
}

// Sweep is a paid mutator transaction binding the contract method 0x62c06767.
//
// Solidity: function sweep(address token, address to, uint256 amount) returns()
func (_LBRouter *LBRouterSession) Sweep(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LBRouter.Contract.Sweep(&_LBRouter.TransactOpts, token, to, amount)
}

// Sweep is a paid mutator transaction binding the contract method 0x62c06767.
//
// Solidity: function sweep(address token, address to, uint256 amount) returns()
func (_LBRouter *LBRouterTransactorSession) Sweep(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LBRouter.Contract.Sweep(&_LBRouter.TransactOpts, token, to, amount)
}

// SweepLBToken is a paid mutator transaction binding the contract method 0xe9361c08.
//
// Solidity: function sweepLBToken(address lbToken, address to, uint256[] ids, uint256[] amounts) returns()
func (_LBRouter *LBRouterTransactor) SweepLBToken(opts *bind.TransactOpts, lbToken common.Address, to common.Address, ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _LBRouter.contract.Transact(opts, "sweepLBToken", lbToken, to, ids, amounts)
}

// SweepLBToken is a paid mutator transaction binding the contract method 0xe9361c08.
//
// Solidity: function sweepLBToken(address lbToken, address to, uint256[] ids, uint256[] amounts) returns()
func (_LBRouter *LBRouterSession) SweepLBToken(lbToken common.Address, to common.Address, ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _LBRouter.Contract.SweepLBToken(&_LBRouter.TransactOpts, lbToken, to, ids, amounts)
}

// SweepLBToken is a paid mutator transaction binding the contract method 0xe9361c08.
//
// Solidity: function sweepLBToken(address lbToken, address to, uint256[] ids, uint256[] amounts) returns()
func (_LBRouter *LBRouterTransactorSession) SweepLBToken(lbToken common.Address, to common.Address, ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _LBRouter.Contract.SweepLBToken(&_LBRouter.TransactOpts, lbToken, to, ids, amounts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LBRouter *LBRouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LBRouter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LBRouter *LBRouterSession) Receive() (*types.Transaction, error) {
	return _LBRouter.Contract.Receive(&_LBRouter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LBRouter *LBRouterTransactorSession) Receive() (*types.Transaction, error) {
	return _LBRouter.Contract.Receive(&_LBRouter.TransactOpts)
}
