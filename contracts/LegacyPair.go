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

// FeeHelperFeeParameters is an auto generated low-level Go binding around an user-defined struct.
type FeeHelperFeeParameters struct {
	BinStep                  uint16
	BaseFactor               uint16
	FilterPeriod             uint16
	DecayPeriod              uint16
	ReductionFactor          uint16
	VariableFeeControl       *big.Int
	ProtocolShare            uint16
	MaxVolatilityAccumulated *big.Int
	VolatilityAccumulated    *big.Int
	VolatilityReference      *big.Int
	IndexRef                 *big.Int
	Time                     *big.Int
}

// LegacypairMetaData contains all meta data concerning the Legacypair contract.
var LegacypairMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractILBFactory\",\"name\":\"_factory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bp\",\"type\":\"uint256\"}],\"name\":\"BinHelper__BinStepOverflows\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BinHelper__IdOverflows\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBPair__AddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBPair__AddressZeroOrThis\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBPair__AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"LBPair__CompositionFactorFlawed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBPair__DistributionsOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBPair__FlashLoanCallbackFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBPair__FlashLoanInvalidBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBPair__FlashLoanInvalidToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBPair__InsufficientAmounts\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"LBPair__InsufficientLiquidityBurned\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"LBPair__InsufficientLiquidityMinted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBPair__OnlyFactory\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"LBPair__OnlyFeeRecipient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBPair__OnlyStrictlyIncreasingId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oracleSize\",\"type\":\"uint256\"}],\"name\":\"LBPair__OracleNewSizeTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBPair__WrongLengths\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LBToken__BurnExceedsBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBToken__BurnFromAddress0\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"accountsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idsLength\",\"type\":\"uint256\"}],\"name\":\"LBToken__LengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBToken__MintToAddress0\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LBToken__SelfApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"LBToken__SpenderNotApproved\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LBToken__TransferExceedsBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBToken__TransferFromOrToAddress0\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBToken__TransferToSelf\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"y\",\"type\":\"int256\"}],\"name\":\"Math128x128__PowerUnderflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"prod1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"}],\"name\":\"Math512Bits__MulDivOverflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"prod1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"Math512Bits__MulShiftOverflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"Math512Bits__OffsetOverflows\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lookUpTimestamp\",\"type\":\"uint256\"}],\"name\":\"Oracle__LookUpTimestampTooOld\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Oracle__NotInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardUpgradeable__AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardUpgradeable__ReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"SafeCast__Exceeds112Bits\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"SafeCast__Exceeds128Bits\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"SafeCast__Exceeds24Bits\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"SafeCast__Exceeds40Bits\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenHelper__CallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenHelper__NonContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenHelper__TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TreeMath__ErrorDepthSearch\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesX\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesY\",\"type\":\"uint256\"}],\"name\":\"CompositionFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"name\":\"DepositedToBin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"name\":\"FeesCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractILBFlashLoanCallback\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"FlashLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSize\",\"type\":\"uint256\"}],\"name\":\"OracleSizeIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"name\":\"ProtocolFeesCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"swapForY\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"volatilityAccumulated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fees\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"name\":\"WithdrawnFromBin\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"batchBalances\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"}],\"name\":\"collectFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectProtocolFees\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amountX\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amountY\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"contractILBFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeParameters\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"binStep\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"baseFactor\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"filterPeriod\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"decayPeriod\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"reductionFactor\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"variableFeeControl\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"protocolShare\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"maxVolatilityAccumulated\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"volatilityAccumulated\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"volatilityReference\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"indexRef\",\"type\":\"uint24\"},{\"internalType\":\"uint40\",\"name\":\"time\",\"type\":\"uint40\"}],\"internalType\":\"structFeeHelper.FeeParameters\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_id\",\"type\":\"uint24\"},{\"internalType\":\"bool\",\"name\":\"_swapForY\",\"type\":\"bool\"}],\"name\":\"findFirstNonEmptyBinId\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILBFlashLoanCallback\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceDecay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_id\",\"type\":\"uint24\"}],\"name\":\"getBin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reserveX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveY\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGlobalFees\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"feesXTotal\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"feesYTotal\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"feesXProtocol\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"feesYProtocol\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOracleParameters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"oracleSampleLifetime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oracleSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oracleActiveSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oracleLastTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oracleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timeDelta\",\"type\":\"uint256\"}],\"name\":\"getOracleSampleFrom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cumulativeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeVolatilityAccumulated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeBinCrossed\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReservesAndId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reserveX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activeId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_newLength\",\"type\":\"uint16\"}],\"name\":\"increaseOracleLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_activeId\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"_sampleLifetime\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"_packedFeeParameters\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_distributionX\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_distributionY\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"liquidityMinted\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"}],\"name\":\"pendingFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_packedFeeParameters\",\"type\":\"bytes32\"}],\"name\":\"setFeesParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_swapForY\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountXOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountYOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenX\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenY\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LegacypairABI is the input ABI used to generate the binding from.
// Deprecated: Use LegacypairMetaData.ABI instead.
var LegacypairABI = LegacypairMetaData.ABI

// Legacypair is an auto generated Go binding around an Ethereum contract.
type Legacypair struct {
	LegacypairCaller     // Read-only binding to the contract
	LegacypairTransactor // Write-only binding to the contract
	LegacypairFilterer   // Log filterer for contract events
}

// LegacypairCaller is an auto generated read-only Go binding around an Ethereum contract.
type LegacypairCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacypairTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LegacypairTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacypairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LegacypairFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacypairSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LegacypairSession struct {
	Contract     *Legacypair       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LegacypairCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LegacypairCallerSession struct {
	Contract *LegacypairCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// LegacypairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LegacypairTransactorSession struct {
	Contract     *LegacypairTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// LegacypairRaw is an auto generated low-level Go binding around an Ethereum contract.
type LegacypairRaw struct {
	Contract *Legacypair // Generic contract binding to access the raw methods on
}

// LegacypairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LegacypairCallerRaw struct {
	Contract *LegacypairCaller // Generic read-only contract binding to access the raw methods on
}

// LegacypairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LegacypairTransactorRaw struct {
	Contract *LegacypairTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLegacypair creates a new instance of Legacypair, bound to a specific deployed contract.
func NewLegacypair(address common.Address, backend bind.ContractBackend) (*Legacypair, error) {
	contract, err := bindLegacypair(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Legacypair{LegacypairCaller: LegacypairCaller{contract: contract}, LegacypairTransactor: LegacypairTransactor{contract: contract}, LegacypairFilterer: LegacypairFilterer{contract: contract}}, nil
}

// NewLegacypairCaller creates a new read-only instance of Legacypair, bound to a specific deployed contract.
func NewLegacypairCaller(address common.Address, caller bind.ContractCaller) (*LegacypairCaller, error) {
	contract, err := bindLegacypair(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LegacypairCaller{contract: contract}, nil
}

// NewLegacypairTransactor creates a new write-only instance of Legacypair, bound to a specific deployed contract.
func NewLegacypairTransactor(address common.Address, transactor bind.ContractTransactor) (*LegacypairTransactor, error) {
	contract, err := bindLegacypair(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LegacypairTransactor{contract: contract}, nil
}

// NewLegacypairFilterer creates a new log filterer instance of Legacypair, bound to a specific deployed contract.
func NewLegacypairFilterer(address common.Address, filterer bind.ContractFilterer) (*LegacypairFilterer, error) {
	contract, err := bindLegacypair(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LegacypairFilterer{contract: contract}, nil
}

// bindLegacypair binds a generic wrapper to an already deployed contract.
func bindLegacypair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LegacypairMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Legacypair *LegacypairRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Legacypair.Contract.LegacypairCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Legacypair *LegacypairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Legacypair.Contract.LegacypairTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Legacypair *LegacypairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Legacypair.Contract.LegacypairTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Legacypair *LegacypairCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Legacypair.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Legacypair *LegacypairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Legacypair.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Legacypair *LegacypairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Legacypair.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address _account, uint256 _id) view returns(uint256)
func (_Legacypair *LegacypairCaller) BalanceOf(opts *bind.CallOpts, _account common.Address, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Legacypair.contract.Call(opts, &out, "balanceOf", _account, _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address _account, uint256 _id) view returns(uint256)
func (_Legacypair *LegacypairSession) BalanceOf(_account common.Address, _id *big.Int) (*big.Int, error) {
	return _Legacypair.Contract.BalanceOf(&_Legacypair.CallOpts, _account, _id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address _account, uint256 _id) view returns(uint256)
func (_Legacypair *LegacypairCallerSession) BalanceOf(_account common.Address, _id *big.Int) (*big.Int, error) {
	return _Legacypair.Contract.BalanceOf(&_Legacypair.CallOpts, _account, _id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] _accounts, uint256[] _ids) view returns(uint256[] batchBalances)
func (_Legacypair *LegacypairCaller) BalanceOfBatch(opts *bind.CallOpts, _accounts []common.Address, _ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Legacypair.contract.Call(opts, &out, "balanceOfBatch", _accounts, _ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] _accounts, uint256[] _ids) view returns(uint256[] batchBalances)
func (_Legacypair *LegacypairSession) BalanceOfBatch(_accounts []common.Address, _ids []*big.Int) ([]*big.Int, error) {
	return _Legacypair.Contract.BalanceOfBatch(&_Legacypair.CallOpts, _accounts, _ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] _accounts, uint256[] _ids) view returns(uint256[] batchBalances)
func (_Legacypair *LegacypairCallerSession) BalanceOfBatch(_accounts []common.Address, _ids []*big.Int) ([]*big.Int, error) {
	return _Legacypair.Contract.BalanceOfBatch(&_Legacypair.CallOpts, _accounts, _ids)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Legacypair *LegacypairCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Legacypair.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Legacypair *LegacypairSession) Factory() (common.Address, error) {
	return _Legacypair.Contract.Factory(&_Legacypair.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Legacypair *LegacypairCallerSession) Factory() (common.Address, error) {
	return _Legacypair.Contract.Factory(&_Legacypair.CallOpts)
}

// FeeParameters is a free data retrieval call binding the contract method 0x98c7adf3.
//
// Solidity: function feeParameters() view returns((uint16,uint16,uint16,uint16,uint16,uint24,uint16,uint24,uint24,uint24,uint24,uint40))
func (_Legacypair *LegacypairCaller) FeeParameters(opts *bind.CallOpts) (FeeHelperFeeParameters, error) {
	var out []interface{}
	err := _Legacypair.contract.Call(opts, &out, "feeParameters")

	if err != nil {
		return *new(FeeHelperFeeParameters), err
	}

	out0 := *abi.ConvertType(out[0], new(FeeHelperFeeParameters)).(*FeeHelperFeeParameters)

	return out0, err

}

// FeeParameters is a free data retrieval call binding the contract method 0x98c7adf3.
//
// Solidity: function feeParameters() view returns((uint16,uint16,uint16,uint16,uint16,uint24,uint16,uint24,uint24,uint24,uint24,uint40))
func (_Legacypair *LegacypairSession) FeeParameters() (FeeHelperFeeParameters, error) {
	return _Legacypair.Contract.FeeParameters(&_Legacypair.CallOpts)
}

// FeeParameters is a free data retrieval call binding the contract method 0x98c7adf3.
//
// Solidity: function feeParameters() view returns((uint16,uint16,uint16,uint16,uint16,uint24,uint16,uint24,uint24,uint24,uint24,uint40))
func (_Legacypair *LegacypairCallerSession) FeeParameters() (FeeHelperFeeParameters, error) {
	return _Legacypair.Contract.FeeParameters(&_Legacypair.CallOpts)
}

// FindFirstNonEmptyBinId is a free data retrieval call binding the contract method 0x8f919a83.
//
// Solidity: function findFirstNonEmptyBinId(uint24 _id, bool _swapForY) view returns(uint24)
func (_Legacypair *LegacypairCaller) FindFirstNonEmptyBinId(opts *bind.CallOpts, _id *big.Int, _swapForY bool) (*big.Int, error) {
	var out []interface{}
	err := _Legacypair.contract.Call(opts, &out, "findFirstNonEmptyBinId", _id, _swapForY)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FindFirstNonEmptyBinId is a free data retrieval call binding the contract method 0x8f919a83.
//
// Solidity: function findFirstNonEmptyBinId(uint24 _id, bool _swapForY) view returns(uint24)
func (_Legacypair *LegacypairSession) FindFirstNonEmptyBinId(_id *big.Int, _swapForY bool) (*big.Int, error) {
	return _Legacypair.Contract.FindFirstNonEmptyBinId(&_Legacypair.CallOpts, _id, _swapForY)
}

// FindFirstNonEmptyBinId is a free data retrieval call binding the contract method 0x8f919a83.
//
// Solidity: function findFirstNonEmptyBinId(uint24 _id, bool _swapForY) view returns(uint24)
func (_Legacypair *LegacypairCallerSession) FindFirstNonEmptyBinId(_id *big.Int, _swapForY bool) (*big.Int, error) {
	return _Legacypair.Contract.FindFirstNonEmptyBinId(&_Legacypair.CallOpts, _id, _swapForY)
}

// GetBin is a free data retrieval call binding the contract method 0x0abe9688.
//
// Solidity: function getBin(uint24 _id) view returns(uint256 reserveX, uint256 reserveY)
func (_Legacypair *LegacypairCaller) GetBin(opts *bind.CallOpts, _id *big.Int) (struct {
	ReserveX *big.Int
	ReserveY *big.Int
}, error) {
	var out []interface{}
	err := _Legacypair.contract.Call(opts, &out, "getBin", _id)

	outstruct := new(struct {
		ReserveX *big.Int
		ReserveY *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReserveX = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReserveY = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBin is a free data retrieval call binding the contract method 0x0abe9688.
//
// Solidity: function getBin(uint24 _id) view returns(uint256 reserveX, uint256 reserveY)
func (_Legacypair *LegacypairSession) GetBin(_id *big.Int) (struct {
	ReserveX *big.Int
	ReserveY *big.Int
}, error) {
	return _Legacypair.Contract.GetBin(&_Legacypair.CallOpts, _id)
}

// GetBin is a free data retrieval call binding the contract method 0x0abe9688.
//
// Solidity: function getBin(uint24 _id) view returns(uint256 reserveX, uint256 reserveY)
func (_Legacypair *LegacypairCallerSession) GetBin(_id *big.Int) (struct {
	ReserveX *big.Int
	ReserveY *big.Int
}, error) {
	return _Legacypair.Contract.GetBin(&_Legacypair.CallOpts, _id)
}

// GetGlobalFees is a free data retrieval call binding the contract method 0xa582cdaa.
//
// Solidity: function getGlobalFees() view returns(uint128 feesXTotal, uint128 feesYTotal, uint128 feesXProtocol, uint128 feesYProtocol)
func (_Legacypair *LegacypairCaller) GetGlobalFees(opts *bind.CallOpts) (struct {
	FeesXTotal    *big.Int
	FeesYTotal    *big.Int
	FeesXProtocol *big.Int
	FeesYProtocol *big.Int
}, error) {
	var out []interface{}
	err := _Legacypair.contract.Call(opts, &out, "getGlobalFees")

	outstruct := new(struct {
		FeesXTotal    *big.Int
		FeesYTotal    *big.Int
		FeesXProtocol *big.Int
		FeesYProtocol *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeesXTotal = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FeesYTotal = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FeesXProtocol = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FeesYProtocol = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetGlobalFees is a free data retrieval call binding the contract method 0xa582cdaa.
//
// Solidity: function getGlobalFees() view returns(uint128 feesXTotal, uint128 feesYTotal, uint128 feesXProtocol, uint128 feesYProtocol)
func (_Legacypair *LegacypairSession) GetGlobalFees() (struct {
	FeesXTotal    *big.Int
	FeesYTotal    *big.Int
	FeesXProtocol *big.Int
	FeesYProtocol *big.Int
}, error) {
	return _Legacypair.Contract.GetGlobalFees(&_Legacypair.CallOpts)
}

// GetGlobalFees is a free data retrieval call binding the contract method 0xa582cdaa.
//
// Solidity: function getGlobalFees() view returns(uint128 feesXTotal, uint128 feesYTotal, uint128 feesXProtocol, uint128 feesYProtocol)
func (_Legacypair *LegacypairCallerSession) GetGlobalFees() (struct {
	FeesXTotal    *big.Int
	FeesYTotal    *big.Int
	FeesXProtocol *big.Int
	FeesYProtocol *big.Int
}, error) {
	return _Legacypair.Contract.GetGlobalFees(&_Legacypair.CallOpts)
}

// GetOracleParameters is a free data retrieval call binding the contract method 0x55182894.
//
// Solidity: function getOracleParameters() view returns(uint256 oracleSampleLifetime, uint256 oracleSize, uint256 oracleActiveSize, uint256 oracleLastTimestamp, uint256 oracleId, uint256 min, uint256 max)
func (_Legacypair *LegacypairCaller) GetOracleParameters(opts *bind.CallOpts) (struct {
	OracleSampleLifetime *big.Int
	OracleSize           *big.Int
	OracleActiveSize     *big.Int
	OracleLastTimestamp  *big.Int
	OracleId             *big.Int
	Min                  *big.Int
	Max                  *big.Int
}, error) {
	var out []interface{}
	err := _Legacypair.contract.Call(opts, &out, "getOracleParameters")

	outstruct := new(struct {
		OracleSampleLifetime *big.Int
		OracleSize           *big.Int
		OracleActiveSize     *big.Int
		OracleLastTimestamp  *big.Int
		OracleId             *big.Int
		Min                  *big.Int
		Max                  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OracleSampleLifetime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.OracleSize = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.OracleActiveSize = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.OracleLastTimestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.OracleId = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Min = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Max = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetOracleParameters is a free data retrieval call binding the contract method 0x55182894.
//
// Solidity: function getOracleParameters() view returns(uint256 oracleSampleLifetime, uint256 oracleSize, uint256 oracleActiveSize, uint256 oracleLastTimestamp, uint256 oracleId, uint256 min, uint256 max)
func (_Legacypair *LegacypairSession) GetOracleParameters() (struct {
	OracleSampleLifetime *big.Int
	OracleSize           *big.Int
	OracleActiveSize     *big.Int
	OracleLastTimestamp  *big.Int
	OracleId             *big.Int
	Min                  *big.Int
	Max                  *big.Int
}, error) {
	return _Legacypair.Contract.GetOracleParameters(&_Legacypair.CallOpts)
}

// GetOracleParameters is a free data retrieval call binding the contract method 0x55182894.
//
// Solidity: function getOracleParameters() view returns(uint256 oracleSampleLifetime, uint256 oracleSize, uint256 oracleActiveSize, uint256 oracleLastTimestamp, uint256 oracleId, uint256 min, uint256 max)
func (_Legacypair *LegacypairCallerSession) GetOracleParameters() (struct {
	OracleSampleLifetime *big.Int
	OracleSize           *big.Int
	OracleActiveSize     *big.Int
	OracleLastTimestamp  *big.Int
	OracleId             *big.Int
	Min                  *big.Int
	Max                  *big.Int
}, error) {
	return _Legacypair.Contract.GetOracleParameters(&_Legacypair.CallOpts)
}

// GetOracleSampleFrom is a free data retrieval call binding the contract method 0xa21635a7.
//
// Solidity: function getOracleSampleFrom(uint256 _timeDelta) view returns(uint256 cumulativeId, uint256 cumulativeVolatilityAccumulated, uint256 cumulativeBinCrossed)
func (_Legacypair *LegacypairCaller) GetOracleSampleFrom(opts *bind.CallOpts, _timeDelta *big.Int) (struct {
	CumulativeId                    *big.Int
	CumulativeVolatilityAccumulated *big.Int
	CumulativeBinCrossed            *big.Int
}, error) {
	var out []interface{}
	err := _Legacypair.contract.Call(opts, &out, "getOracleSampleFrom", _timeDelta)

	outstruct := new(struct {
		CumulativeId                    *big.Int
		CumulativeVolatilityAccumulated *big.Int
		CumulativeBinCrossed            *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CumulativeId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CumulativeVolatilityAccumulated = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CumulativeBinCrossed = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetOracleSampleFrom is a free data retrieval call binding the contract method 0xa21635a7.
//
// Solidity: function getOracleSampleFrom(uint256 _timeDelta) view returns(uint256 cumulativeId, uint256 cumulativeVolatilityAccumulated, uint256 cumulativeBinCrossed)
func (_Legacypair *LegacypairSession) GetOracleSampleFrom(_timeDelta *big.Int) (struct {
	CumulativeId                    *big.Int
	CumulativeVolatilityAccumulated *big.Int
	CumulativeBinCrossed            *big.Int
}, error) {
	return _Legacypair.Contract.GetOracleSampleFrom(&_Legacypair.CallOpts, _timeDelta)
}

// GetOracleSampleFrom is a free data retrieval call binding the contract method 0xa21635a7.
//
// Solidity: function getOracleSampleFrom(uint256 _timeDelta) view returns(uint256 cumulativeId, uint256 cumulativeVolatilityAccumulated, uint256 cumulativeBinCrossed)
func (_Legacypair *LegacypairCallerSession) GetOracleSampleFrom(_timeDelta *big.Int) (struct {
	CumulativeId                    *big.Int
	CumulativeVolatilityAccumulated *big.Int
	CumulativeBinCrossed            *big.Int
}, error) {
	return _Legacypair.Contract.GetOracleSampleFrom(&_Legacypair.CallOpts, _timeDelta)
}

// GetReservesAndId is a free data retrieval call binding the contract method 0x1b05b83e.
//
// Solidity: function getReservesAndId() view returns(uint256 reserveX, uint256 reserveY, uint256 activeId)
func (_Legacypair *LegacypairCaller) GetReservesAndId(opts *bind.CallOpts) (struct {
	ReserveX *big.Int
	ReserveY *big.Int
	ActiveId *big.Int
}, error) {
	var out []interface{}
	err := _Legacypair.contract.Call(opts, &out, "getReservesAndId")

	outstruct := new(struct {
		ReserveX *big.Int
		ReserveY *big.Int
		ActiveId *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReserveX = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReserveY = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ActiveId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReservesAndId is a free data retrieval call binding the contract method 0x1b05b83e.
//
// Solidity: function getReservesAndId() view returns(uint256 reserveX, uint256 reserveY, uint256 activeId)
func (_Legacypair *LegacypairSession) GetReservesAndId() (struct {
	ReserveX *big.Int
	ReserveY *big.Int
	ActiveId *big.Int
}, error) {
	return _Legacypair.Contract.GetReservesAndId(&_Legacypair.CallOpts)
}

// GetReservesAndId is a free data retrieval call binding the contract method 0x1b05b83e.
//
// Solidity: function getReservesAndId() view returns(uint256 reserveX, uint256 reserveY, uint256 activeId)
func (_Legacypair *LegacypairCallerSession) GetReservesAndId() (struct {
	ReserveX *big.Int
	ReserveY *big.Int
	ActiveId *big.Int
}, error) {
	return _Legacypair.Contract.GetReservesAndId(&_Legacypair.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _spender) view returns(bool)
func (_Legacypair *LegacypairCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (bool, error) {
	var out []interface{}
	err := _Legacypair.contract.Call(opts, &out, "isApprovedForAll", _owner, _spender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _spender) view returns(bool)
func (_Legacypair *LegacypairSession) IsApprovedForAll(_owner common.Address, _spender common.Address) (bool, error) {
	return _Legacypair.Contract.IsApprovedForAll(&_Legacypair.CallOpts, _owner, _spender)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _spender) view returns(bool)
func (_Legacypair *LegacypairCallerSession) IsApprovedForAll(_owner common.Address, _spender common.Address) (bool, error) {
	return _Legacypair.Contract.IsApprovedForAll(&_Legacypair.CallOpts, _owner, _spender)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Legacypair *LegacypairCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Legacypair.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Legacypair *LegacypairSession) Name() (string, error) {
	return _Legacypair.Contract.Name(&_Legacypair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Legacypair *LegacypairCallerSession) Name() (string, error) {
	return _Legacypair.Contract.Name(&_Legacypair.CallOpts)
}

// PendingFees is a free data retrieval call binding the contract method 0xf7cff1f8.
//
// Solidity: function pendingFees(address _account, uint256[] _ids) view returns(uint256 amountX, uint256 amountY)
func (_Legacypair *LegacypairCaller) PendingFees(opts *bind.CallOpts, _account common.Address, _ids []*big.Int) (struct {
	AmountX *big.Int
	AmountY *big.Int
}, error) {
	var out []interface{}
	err := _Legacypair.contract.Call(opts, &out, "pendingFees", _account, _ids)

	outstruct := new(struct {
		AmountX *big.Int
		AmountY *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountX = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountY = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingFees is a free data retrieval call binding the contract method 0xf7cff1f8.
//
// Solidity: function pendingFees(address _account, uint256[] _ids) view returns(uint256 amountX, uint256 amountY)
func (_Legacypair *LegacypairSession) PendingFees(_account common.Address, _ids []*big.Int) (struct {
	AmountX *big.Int
	AmountY *big.Int
}, error) {
	return _Legacypair.Contract.PendingFees(&_Legacypair.CallOpts, _account, _ids)
}

// PendingFees is a free data retrieval call binding the contract method 0xf7cff1f8.
//
// Solidity: function pendingFees(address _account, uint256[] _ids) view returns(uint256 amountX, uint256 amountY)
func (_Legacypair *LegacypairCallerSession) PendingFees(_account common.Address, _ids []*big.Int) (struct {
	AmountX *big.Int
	AmountY *big.Int
}, error) {
	return _Legacypair.Contract.PendingFees(&_Legacypair.CallOpts, _account, _ids)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_Legacypair *LegacypairCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Legacypair.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_Legacypair *LegacypairSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _Legacypair.Contract.SupportsInterface(&_Legacypair.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_Legacypair *LegacypairCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _Legacypair.Contract.SupportsInterface(&_Legacypair.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Legacypair *LegacypairCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Legacypair.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Legacypair *LegacypairSession) Symbol() (string, error) {
	return _Legacypair.Contract.Symbol(&_Legacypair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Legacypair *LegacypairCallerSession) Symbol() (string, error) {
	return _Legacypair.Contract.Symbol(&_Legacypair.CallOpts)
}

// TokenX is a free data retrieval call binding the contract method 0x16dc165b.
//
// Solidity: function tokenX() view returns(address)
func (_Legacypair *LegacypairCaller) TokenX(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Legacypair.contract.Call(opts, &out, "tokenX")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenX is a free data retrieval call binding the contract method 0x16dc165b.
//
// Solidity: function tokenX() view returns(address)
func (_Legacypair *LegacypairSession) TokenX() (common.Address, error) {
	return _Legacypair.Contract.TokenX(&_Legacypair.CallOpts)
}

// TokenX is a free data retrieval call binding the contract method 0x16dc165b.
//
// Solidity: function tokenX() view returns(address)
func (_Legacypair *LegacypairCallerSession) TokenX() (common.Address, error) {
	return _Legacypair.Contract.TokenX(&_Legacypair.CallOpts)
}

// TokenY is a free data retrieval call binding the contract method 0xb7d19fc4.
//
// Solidity: function tokenY() view returns(address)
func (_Legacypair *LegacypairCaller) TokenY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Legacypair.contract.Call(opts, &out, "tokenY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenY is a free data retrieval call binding the contract method 0xb7d19fc4.
//
// Solidity: function tokenY() view returns(address)
func (_Legacypair *LegacypairSession) TokenY() (common.Address, error) {
	return _Legacypair.Contract.TokenY(&_Legacypair.CallOpts)
}

// TokenY is a free data retrieval call binding the contract method 0xb7d19fc4.
//
// Solidity: function tokenY() view returns(address)
func (_Legacypair *LegacypairCallerSession) TokenY() (common.Address, error) {
	return _Legacypair.Contract.TokenY(&_Legacypair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 _id) view returns(uint256)
func (_Legacypair *LegacypairCaller) TotalSupply(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Legacypair.contract.Call(opts, &out, "totalSupply", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 _id) view returns(uint256)
func (_Legacypair *LegacypairSession) TotalSupply(_id *big.Int) (*big.Int, error) {
	return _Legacypair.Contract.TotalSupply(&_Legacypair.CallOpts, _id)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 _id) view returns(uint256)
func (_Legacypair *LegacypairCallerSession) TotalSupply(_id *big.Int) (*big.Int, error) {
	return _Legacypair.Contract.TotalSupply(&_Legacypair.CallOpts, _id)
}

// Burn is a paid mutator transaction binding the contract method 0x0acd451d.
//
// Solidity: function burn(uint256[] _ids, uint256[] _amounts, address _to) returns(uint256 amountX, uint256 amountY)
func (_Legacypair *LegacypairTransactor) Burn(opts *bind.TransactOpts, _ids []*big.Int, _amounts []*big.Int, _to common.Address) (*types.Transaction, error) {
	return _Legacypair.contract.Transact(opts, "burn", _ids, _amounts, _to)
}

// Burn is a paid mutator transaction binding the contract method 0x0acd451d.
//
// Solidity: function burn(uint256[] _ids, uint256[] _amounts, address _to) returns(uint256 amountX, uint256 amountY)
func (_Legacypair *LegacypairSession) Burn(_ids []*big.Int, _amounts []*big.Int, _to common.Address) (*types.Transaction, error) {
	return _Legacypair.Contract.Burn(&_Legacypair.TransactOpts, _ids, _amounts, _to)
}

// Burn is a paid mutator transaction binding the contract method 0x0acd451d.
//
// Solidity: function burn(uint256[] _ids, uint256[] _amounts, address _to) returns(uint256 amountX, uint256 amountY)
func (_Legacypair *LegacypairTransactorSession) Burn(_ids []*big.Int, _amounts []*big.Int, _to common.Address) (*types.Transaction, error) {
	return _Legacypair.Contract.Burn(&_Legacypair.TransactOpts, _ids, _amounts, _to)
}

// CollectFees is a paid mutator transaction binding the contract method 0x225b20b9.
//
// Solidity: function collectFees(address _account, uint256[] _ids) returns(uint256 amountX, uint256 amountY)
func (_Legacypair *LegacypairTransactor) CollectFees(opts *bind.TransactOpts, _account common.Address, _ids []*big.Int) (*types.Transaction, error) {
	return _Legacypair.contract.Transact(opts, "collectFees", _account, _ids)
}

// CollectFees is a paid mutator transaction binding the contract method 0x225b20b9.
//
// Solidity: function collectFees(address _account, uint256[] _ids) returns(uint256 amountX, uint256 amountY)
func (_Legacypair *LegacypairSession) CollectFees(_account common.Address, _ids []*big.Int) (*types.Transaction, error) {
	return _Legacypair.Contract.CollectFees(&_Legacypair.TransactOpts, _account, _ids)
}

// CollectFees is a paid mutator transaction binding the contract method 0x225b20b9.
//
// Solidity: function collectFees(address _account, uint256[] _ids) returns(uint256 amountX, uint256 amountY)
func (_Legacypair *LegacypairTransactorSession) CollectFees(_account common.Address, _ids []*big.Int) (*types.Transaction, error) {
	return _Legacypair.Contract.CollectFees(&_Legacypair.TransactOpts, _account, _ids)
}

// CollectProtocolFees is a paid mutator transaction binding the contract method 0xa1af5b9a.
//
// Solidity: function collectProtocolFees() returns(uint128 amountX, uint128 amountY)
func (_Legacypair *LegacypairTransactor) CollectProtocolFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Legacypair.contract.Transact(opts, "collectProtocolFees")
}

// CollectProtocolFees is a paid mutator transaction binding the contract method 0xa1af5b9a.
//
// Solidity: function collectProtocolFees() returns(uint128 amountX, uint128 amountY)
func (_Legacypair *LegacypairSession) CollectProtocolFees() (*types.Transaction, error) {
	return _Legacypair.Contract.CollectProtocolFees(&_Legacypair.TransactOpts)
}

// CollectProtocolFees is a paid mutator transaction binding the contract method 0xa1af5b9a.
//
// Solidity: function collectProtocolFees() returns(uint128 amountX, uint128 amountY)
func (_Legacypair *LegacypairTransactorSession) CollectProtocolFees() (*types.Transaction, error) {
	return _Legacypair.Contract.CollectProtocolFees(&_Legacypair.TransactOpts)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address _receiver, address _token, uint256 _amount, bytes _data) returns()
func (_Legacypair *LegacypairTransactor) FlashLoan(opts *bind.TransactOpts, _receiver common.Address, _token common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _Legacypair.contract.Transact(opts, "flashLoan", _receiver, _token, _amount, _data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address _receiver, address _token, uint256 _amount, bytes _data) returns()
func (_Legacypair *LegacypairSession) FlashLoan(_receiver common.Address, _token common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _Legacypair.Contract.FlashLoan(&_Legacypair.TransactOpts, _receiver, _token, _amount, _data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address _receiver, address _token, uint256 _amount, bytes _data) returns()
func (_Legacypair *LegacypairTransactorSession) FlashLoan(_receiver common.Address, _token common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _Legacypair.Contract.FlashLoan(&_Legacypair.TransactOpts, _receiver, _token, _amount, _data)
}

// ForceDecay is a paid mutator transaction binding the contract method 0xd3b9fbe4.
//
// Solidity: function forceDecay() returns()
func (_Legacypair *LegacypairTransactor) ForceDecay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Legacypair.contract.Transact(opts, "forceDecay")
}

// ForceDecay is a paid mutator transaction binding the contract method 0xd3b9fbe4.
//
// Solidity: function forceDecay() returns()
func (_Legacypair *LegacypairSession) ForceDecay() (*types.Transaction, error) {
	return _Legacypair.Contract.ForceDecay(&_Legacypair.TransactOpts)
}

// ForceDecay is a paid mutator transaction binding the contract method 0xd3b9fbe4.
//
// Solidity: function forceDecay() returns()
func (_Legacypair *LegacypairTransactorSession) ForceDecay() (*types.Transaction, error) {
	return _Legacypair.Contract.ForceDecay(&_Legacypair.TransactOpts)
}

// IncreaseOracleLength is a paid mutator transaction binding the contract method 0xc7bd6586.
//
// Solidity: function increaseOracleLength(uint16 _newLength) returns()
func (_Legacypair *LegacypairTransactor) IncreaseOracleLength(opts *bind.TransactOpts, _newLength uint16) (*types.Transaction, error) {
	return _Legacypair.contract.Transact(opts, "increaseOracleLength", _newLength)
}

// IncreaseOracleLength is a paid mutator transaction binding the contract method 0xc7bd6586.
//
// Solidity: function increaseOracleLength(uint16 _newLength) returns()
func (_Legacypair *LegacypairSession) IncreaseOracleLength(_newLength uint16) (*types.Transaction, error) {
	return _Legacypair.Contract.IncreaseOracleLength(&_Legacypair.TransactOpts, _newLength)
}

// IncreaseOracleLength is a paid mutator transaction binding the contract method 0xc7bd6586.
//
// Solidity: function increaseOracleLength(uint16 _newLength) returns()
func (_Legacypair *LegacypairTransactorSession) IncreaseOracleLength(_newLength uint16) (*types.Transaction, error) {
	return _Legacypair.Contract.IncreaseOracleLength(&_Legacypair.TransactOpts, _newLength)
}

// Initialize is a paid mutator transaction binding the contract method 0xd32db437.
//
// Solidity: function initialize(address _tokenX, address _tokenY, uint24 _activeId, uint16 _sampleLifetime, bytes32 _packedFeeParameters) returns()
func (_Legacypair *LegacypairTransactor) Initialize(opts *bind.TransactOpts, _tokenX common.Address, _tokenY common.Address, _activeId *big.Int, _sampleLifetime uint16, _packedFeeParameters [32]byte) (*types.Transaction, error) {
	return _Legacypair.contract.Transact(opts, "initialize", _tokenX, _tokenY, _activeId, _sampleLifetime, _packedFeeParameters)
}

// Initialize is a paid mutator transaction binding the contract method 0xd32db437.
//
// Solidity: function initialize(address _tokenX, address _tokenY, uint24 _activeId, uint16 _sampleLifetime, bytes32 _packedFeeParameters) returns()
func (_Legacypair *LegacypairSession) Initialize(_tokenX common.Address, _tokenY common.Address, _activeId *big.Int, _sampleLifetime uint16, _packedFeeParameters [32]byte) (*types.Transaction, error) {
	return _Legacypair.Contract.Initialize(&_Legacypair.TransactOpts, _tokenX, _tokenY, _activeId, _sampleLifetime, _packedFeeParameters)
}

// Initialize is a paid mutator transaction binding the contract method 0xd32db437.
//
// Solidity: function initialize(address _tokenX, address _tokenY, uint24 _activeId, uint16 _sampleLifetime, bytes32 _packedFeeParameters) returns()
func (_Legacypair *LegacypairTransactorSession) Initialize(_tokenX common.Address, _tokenY common.Address, _activeId *big.Int, _sampleLifetime uint16, _packedFeeParameters [32]byte) (*types.Transaction, error) {
	return _Legacypair.Contract.Initialize(&_Legacypair.TransactOpts, _tokenX, _tokenY, _activeId, _sampleLifetime, _packedFeeParameters)
}

// Mint is a paid mutator transaction binding the contract method 0x714c8592.
//
// Solidity: function mint(uint256[] _ids, uint256[] _distributionX, uint256[] _distributionY, address _to) returns(uint256, uint256, uint256[] liquidityMinted)
func (_Legacypair *LegacypairTransactor) Mint(opts *bind.TransactOpts, _ids []*big.Int, _distributionX []*big.Int, _distributionY []*big.Int, _to common.Address) (*types.Transaction, error) {
	return _Legacypair.contract.Transact(opts, "mint", _ids, _distributionX, _distributionY, _to)
}

// Mint is a paid mutator transaction binding the contract method 0x714c8592.
//
// Solidity: function mint(uint256[] _ids, uint256[] _distributionX, uint256[] _distributionY, address _to) returns(uint256, uint256, uint256[] liquidityMinted)
func (_Legacypair *LegacypairSession) Mint(_ids []*big.Int, _distributionX []*big.Int, _distributionY []*big.Int, _to common.Address) (*types.Transaction, error) {
	return _Legacypair.Contract.Mint(&_Legacypair.TransactOpts, _ids, _distributionX, _distributionY, _to)
}

// Mint is a paid mutator transaction binding the contract method 0x714c8592.
//
// Solidity: function mint(uint256[] _ids, uint256[] _distributionX, uint256[] _distributionY, address _to) returns(uint256, uint256, uint256[] liquidityMinted)
func (_Legacypair *LegacypairTransactorSession) Mint(_ids []*big.Int, _distributionX []*big.Int, _distributionY []*big.Int, _to common.Address) (*types.Transaction, error) {
	return _Legacypair.Contract.Mint(&_Legacypair.TransactOpts, _ids, _distributionX, _distributionY, _to)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0xfba0ee64.
//
// Solidity: function safeBatchTransferFrom(address _from, address _to, uint256[] _ids, uint256[] _amounts) returns()
func (_Legacypair *LegacypairTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _ids []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _Legacypair.contract.Transact(opts, "safeBatchTransferFrom", _from, _to, _ids, _amounts)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0xfba0ee64.
//
// Solidity: function safeBatchTransferFrom(address _from, address _to, uint256[] _ids, uint256[] _amounts) returns()
func (_Legacypair *LegacypairSession) SafeBatchTransferFrom(_from common.Address, _to common.Address, _ids []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _Legacypair.Contract.SafeBatchTransferFrom(&_Legacypair.TransactOpts, _from, _to, _ids, _amounts)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0xfba0ee64.
//
// Solidity: function safeBatchTransferFrom(address _from, address _to, uint256[] _ids, uint256[] _amounts) returns()
func (_Legacypair *LegacypairTransactorSession) SafeBatchTransferFrom(_from common.Address, _to common.Address, _ids []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _Legacypair.Contract.SafeBatchTransferFrom(&_Legacypair.TransactOpts, _from, _to, _ids, _amounts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x0febdd49.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _id, uint256 _amount) returns()
func (_Legacypair *LegacypairTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Legacypair.contract.Transact(opts, "safeTransferFrom", _from, _to, _id, _amount)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x0febdd49.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _id, uint256 _amount) returns()
func (_Legacypair *LegacypairSession) SafeTransferFrom(_from common.Address, _to common.Address, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Legacypair.Contract.SafeTransferFrom(&_Legacypair.TransactOpts, _from, _to, _id, _amount)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x0febdd49.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _id, uint256 _amount) returns()
func (_Legacypair *LegacypairTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Legacypair.Contract.SafeTransferFrom(&_Legacypair.TransactOpts, _from, _to, _id, _amount)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _spender, bool _approved) returns()
func (_Legacypair *LegacypairTransactor) SetApprovalForAll(opts *bind.TransactOpts, _spender common.Address, _approved bool) (*types.Transaction, error) {
	return _Legacypair.contract.Transact(opts, "setApprovalForAll", _spender, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _spender, bool _approved) returns()
func (_Legacypair *LegacypairSession) SetApprovalForAll(_spender common.Address, _approved bool) (*types.Transaction, error) {
	return _Legacypair.Contract.SetApprovalForAll(&_Legacypair.TransactOpts, _spender, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _spender, bool _approved) returns()
func (_Legacypair *LegacypairTransactorSession) SetApprovalForAll(_spender common.Address, _approved bool) (*types.Transaction, error) {
	return _Legacypair.Contract.SetApprovalForAll(&_Legacypair.TransactOpts, _spender, _approved)
}

// SetFeesParameters is a paid mutator transaction binding the contract method 0x54b5fc87.
//
// Solidity: function setFeesParameters(bytes32 _packedFeeParameters) returns()
func (_Legacypair *LegacypairTransactor) SetFeesParameters(opts *bind.TransactOpts, _packedFeeParameters [32]byte) (*types.Transaction, error) {
	return _Legacypair.contract.Transact(opts, "setFeesParameters", _packedFeeParameters)
}

// SetFeesParameters is a paid mutator transaction binding the contract method 0x54b5fc87.
//
// Solidity: function setFeesParameters(bytes32 _packedFeeParameters) returns()
func (_Legacypair *LegacypairSession) SetFeesParameters(_packedFeeParameters [32]byte) (*types.Transaction, error) {
	return _Legacypair.Contract.SetFeesParameters(&_Legacypair.TransactOpts, _packedFeeParameters)
}

// SetFeesParameters is a paid mutator transaction binding the contract method 0x54b5fc87.
//
// Solidity: function setFeesParameters(bytes32 _packedFeeParameters) returns()
func (_Legacypair *LegacypairTransactorSession) SetFeesParameters(_packedFeeParameters [32]byte) (*types.Transaction, error) {
	return _Legacypair.Contract.SetFeesParameters(&_Legacypair.TransactOpts, _packedFeeParameters)
}

// Swap is a paid mutator transaction binding the contract method 0x53c059a0.
//
// Solidity: function swap(bool _swapForY, address _to) returns(uint256 amountXOut, uint256 amountYOut)
func (_Legacypair *LegacypairTransactor) Swap(opts *bind.TransactOpts, _swapForY bool, _to common.Address) (*types.Transaction, error) {
	return _Legacypair.contract.Transact(opts, "swap", _swapForY, _to)
}

// Swap is a paid mutator transaction binding the contract method 0x53c059a0.
//
// Solidity: function swap(bool _swapForY, address _to) returns(uint256 amountXOut, uint256 amountYOut)
func (_Legacypair *LegacypairSession) Swap(_swapForY bool, _to common.Address) (*types.Transaction, error) {
	return _Legacypair.Contract.Swap(&_Legacypair.TransactOpts, _swapForY, _to)
}

// Swap is a paid mutator transaction binding the contract method 0x53c059a0.
//
// Solidity: function swap(bool _swapForY, address _to) returns(uint256 amountXOut, uint256 amountYOut)
func (_Legacypair *LegacypairTransactorSession) Swap(_swapForY bool, _to common.Address) (*types.Transaction, error) {
	return _Legacypair.Contract.Swap(&_Legacypair.TransactOpts, _swapForY, _to)
}

// LegacypairApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Legacypair contract.
type LegacypairApprovalForAllIterator struct {
	Event *LegacypairApprovalForAll // Event containing the contract specifics and raw log

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
func (it *LegacypairApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacypairApprovalForAll)
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
		it.Event = new(LegacypairApprovalForAll)
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
func (it *LegacypairApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacypairApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacypairApprovalForAll represents a ApprovalForAll event raised by the Legacypair contract.
type LegacypairApprovalForAll struct {
	Account  common.Address
	Sender   common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed sender, bool approved)
func (_Legacypair *LegacypairFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, sender []common.Address) (*LegacypairApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Legacypair.contract.FilterLogs(opts, "ApprovalForAll", accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LegacypairApprovalForAllIterator{contract: _Legacypair.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed sender, bool approved)
func (_Legacypair *LegacypairFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *LegacypairApprovalForAll, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Legacypair.contract.WatchLogs(opts, "ApprovalForAll", accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacypairApprovalForAll)
				if err := _Legacypair.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed sender, bool approved)
func (_Legacypair *LegacypairFilterer) ParseApprovalForAll(log types.Log) (*LegacypairApprovalForAll, error) {
	event := new(LegacypairApprovalForAll)
	if err := _Legacypair.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacypairCompositionFeeIterator is returned from FilterCompositionFee and is used to iterate over the raw logs and unpacked data for CompositionFee events raised by the Legacypair contract.
type LegacypairCompositionFeeIterator struct {
	Event *LegacypairCompositionFee // Event containing the contract specifics and raw log

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
func (it *LegacypairCompositionFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacypairCompositionFee)
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
		it.Event = new(LegacypairCompositionFee)
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
func (it *LegacypairCompositionFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacypairCompositionFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacypairCompositionFee represents a CompositionFee event raised by the Legacypair contract.
type LegacypairCompositionFee struct {
	Sender    common.Address
	Recipient common.Address
	Id        *big.Int
	FeesX     *big.Int
	FeesY     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCompositionFee is a free log retrieval operation binding the contract event 0x56f8e764728c77dd99ffbc1b64e6d02e227e6ec8214f165d4ef31351de136a0d.
//
// Solidity: event CompositionFee(address indexed sender, address indexed recipient, uint256 indexed id, uint256 feesX, uint256 feesY)
func (_Legacypair *LegacypairFilterer) FilterCompositionFee(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address, id []*big.Int) (*LegacypairCompositionFeeIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Legacypair.contract.FilterLogs(opts, "CompositionFee", senderRule, recipientRule, idRule)
	if err != nil {
		return nil, err
	}
	return &LegacypairCompositionFeeIterator{contract: _Legacypair.contract, event: "CompositionFee", logs: logs, sub: sub}, nil
}

// WatchCompositionFee is a free log subscription operation binding the contract event 0x56f8e764728c77dd99ffbc1b64e6d02e227e6ec8214f165d4ef31351de136a0d.
//
// Solidity: event CompositionFee(address indexed sender, address indexed recipient, uint256 indexed id, uint256 feesX, uint256 feesY)
func (_Legacypair *LegacypairFilterer) WatchCompositionFee(opts *bind.WatchOpts, sink chan<- *LegacypairCompositionFee, sender []common.Address, recipient []common.Address, id []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Legacypair.contract.WatchLogs(opts, "CompositionFee", senderRule, recipientRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacypairCompositionFee)
				if err := _Legacypair.contract.UnpackLog(event, "CompositionFee", log); err != nil {
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

// ParseCompositionFee is a log parse operation binding the contract event 0x56f8e764728c77dd99ffbc1b64e6d02e227e6ec8214f165d4ef31351de136a0d.
//
// Solidity: event CompositionFee(address indexed sender, address indexed recipient, uint256 indexed id, uint256 feesX, uint256 feesY)
func (_Legacypair *LegacypairFilterer) ParseCompositionFee(log types.Log) (*LegacypairCompositionFee, error) {
	event := new(LegacypairCompositionFee)
	if err := _Legacypair.contract.UnpackLog(event, "CompositionFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacypairDepositedToBinIterator is returned from FilterDepositedToBin and is used to iterate over the raw logs and unpacked data for DepositedToBin events raised by the Legacypair contract.
type LegacypairDepositedToBinIterator struct {
	Event *LegacypairDepositedToBin // Event containing the contract specifics and raw log

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
func (it *LegacypairDepositedToBinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacypairDepositedToBin)
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
		it.Event = new(LegacypairDepositedToBin)
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
func (it *LegacypairDepositedToBinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacypairDepositedToBinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacypairDepositedToBin represents a DepositedToBin event raised by the Legacypair contract.
type LegacypairDepositedToBin struct {
	Sender    common.Address
	Recipient common.Address
	Id        *big.Int
	AmountX   *big.Int
	AmountY   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositedToBin is a free log retrieval operation binding the contract event 0x4216cc3bd0c40a90259d92f800c06ede5c47765f41a488072b7e7104a1f95841.
//
// Solidity: event DepositedToBin(address indexed sender, address indexed recipient, uint256 indexed id, uint256 amountX, uint256 amountY)
func (_Legacypair *LegacypairFilterer) FilterDepositedToBin(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address, id []*big.Int) (*LegacypairDepositedToBinIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Legacypair.contract.FilterLogs(opts, "DepositedToBin", senderRule, recipientRule, idRule)
	if err != nil {
		return nil, err
	}
	return &LegacypairDepositedToBinIterator{contract: _Legacypair.contract, event: "DepositedToBin", logs: logs, sub: sub}, nil
}

// WatchDepositedToBin is a free log subscription operation binding the contract event 0x4216cc3bd0c40a90259d92f800c06ede5c47765f41a488072b7e7104a1f95841.
//
// Solidity: event DepositedToBin(address indexed sender, address indexed recipient, uint256 indexed id, uint256 amountX, uint256 amountY)
func (_Legacypair *LegacypairFilterer) WatchDepositedToBin(opts *bind.WatchOpts, sink chan<- *LegacypairDepositedToBin, sender []common.Address, recipient []common.Address, id []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Legacypair.contract.WatchLogs(opts, "DepositedToBin", senderRule, recipientRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacypairDepositedToBin)
				if err := _Legacypair.contract.UnpackLog(event, "DepositedToBin", log); err != nil {
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

// ParseDepositedToBin is a log parse operation binding the contract event 0x4216cc3bd0c40a90259d92f800c06ede5c47765f41a488072b7e7104a1f95841.
//
// Solidity: event DepositedToBin(address indexed sender, address indexed recipient, uint256 indexed id, uint256 amountX, uint256 amountY)
func (_Legacypair *LegacypairFilterer) ParseDepositedToBin(log types.Log) (*LegacypairDepositedToBin, error) {
	event := new(LegacypairDepositedToBin)
	if err := _Legacypair.contract.UnpackLog(event, "DepositedToBin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacypairFeesCollectedIterator is returned from FilterFeesCollected and is used to iterate over the raw logs and unpacked data for FeesCollected events raised by the Legacypair contract.
type LegacypairFeesCollectedIterator struct {
	Event *LegacypairFeesCollected // Event containing the contract specifics and raw log

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
func (it *LegacypairFeesCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacypairFeesCollected)
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
		it.Event = new(LegacypairFeesCollected)
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
func (it *LegacypairFeesCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacypairFeesCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacypairFeesCollected represents a FeesCollected event raised by the Legacypair contract.
type LegacypairFeesCollected struct {
	Sender    common.Address
	Recipient common.Address
	AmountX   *big.Int
	AmountY   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeesCollected is a free log retrieval operation binding the contract event 0x28a87b6059180e46de5fb9ab35eb043e8fe00ab45afcc7789e3934ecbbcde3ea.
//
// Solidity: event FeesCollected(address indexed sender, address indexed recipient, uint256 amountX, uint256 amountY)
func (_Legacypair *LegacypairFilterer) FilterFeesCollected(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*LegacypairFeesCollectedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Legacypair.contract.FilterLogs(opts, "FeesCollected", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &LegacypairFeesCollectedIterator{contract: _Legacypair.contract, event: "FeesCollected", logs: logs, sub: sub}, nil
}

// WatchFeesCollected is a free log subscription operation binding the contract event 0x28a87b6059180e46de5fb9ab35eb043e8fe00ab45afcc7789e3934ecbbcde3ea.
//
// Solidity: event FeesCollected(address indexed sender, address indexed recipient, uint256 amountX, uint256 amountY)
func (_Legacypair *LegacypairFilterer) WatchFeesCollected(opts *bind.WatchOpts, sink chan<- *LegacypairFeesCollected, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Legacypair.contract.WatchLogs(opts, "FeesCollected", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacypairFeesCollected)
				if err := _Legacypair.contract.UnpackLog(event, "FeesCollected", log); err != nil {
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

// ParseFeesCollected is a log parse operation binding the contract event 0x28a87b6059180e46de5fb9ab35eb043e8fe00ab45afcc7789e3934ecbbcde3ea.
//
// Solidity: event FeesCollected(address indexed sender, address indexed recipient, uint256 amountX, uint256 amountY)
func (_Legacypair *LegacypairFilterer) ParseFeesCollected(log types.Log) (*LegacypairFeesCollected, error) {
	event := new(LegacypairFeesCollected)
	if err := _Legacypair.contract.UnpackLog(event, "FeesCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacypairFlashLoanIterator is returned from FilterFlashLoan and is used to iterate over the raw logs and unpacked data for FlashLoan events raised by the Legacypair contract.
type LegacypairFlashLoanIterator struct {
	Event *LegacypairFlashLoan // Event containing the contract specifics and raw log

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
func (it *LegacypairFlashLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacypairFlashLoan)
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
		it.Event = new(LegacypairFlashLoan)
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
func (it *LegacypairFlashLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacypairFlashLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacypairFlashLoan represents a FlashLoan event raised by the Legacypair contract.
type LegacypairFlashLoan struct {
	Sender   common.Address
	Receiver common.Address
	Token    common.Address
	Amount   *big.Int
	Fee      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFlashLoan is a free log retrieval operation binding the contract event 0x3659d15bd4bb92ab352a8d35bc3119ec6e7e0ab48e4d46201c8a28e02b6a8a86.
//
// Solidity: event FlashLoan(address indexed sender, address indexed receiver, address token, uint256 amount, uint256 fee)
func (_Legacypair *LegacypairFilterer) FilterFlashLoan(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*LegacypairFlashLoanIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Legacypair.contract.FilterLogs(opts, "FlashLoan", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &LegacypairFlashLoanIterator{contract: _Legacypair.contract, event: "FlashLoan", logs: logs, sub: sub}, nil
}

// WatchFlashLoan is a free log subscription operation binding the contract event 0x3659d15bd4bb92ab352a8d35bc3119ec6e7e0ab48e4d46201c8a28e02b6a8a86.
//
// Solidity: event FlashLoan(address indexed sender, address indexed receiver, address token, uint256 amount, uint256 fee)
func (_Legacypair *LegacypairFilterer) WatchFlashLoan(opts *bind.WatchOpts, sink chan<- *LegacypairFlashLoan, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Legacypair.contract.WatchLogs(opts, "FlashLoan", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacypairFlashLoan)
				if err := _Legacypair.contract.UnpackLog(event, "FlashLoan", log); err != nil {
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

// ParseFlashLoan is a log parse operation binding the contract event 0x3659d15bd4bb92ab352a8d35bc3119ec6e7e0ab48e4d46201c8a28e02b6a8a86.
//
// Solidity: event FlashLoan(address indexed sender, address indexed receiver, address token, uint256 amount, uint256 fee)
func (_Legacypair *LegacypairFilterer) ParseFlashLoan(log types.Log) (*LegacypairFlashLoan, error) {
	event := new(LegacypairFlashLoan)
	if err := _Legacypair.contract.UnpackLog(event, "FlashLoan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacypairOracleSizeIncreasedIterator is returned from FilterOracleSizeIncreased and is used to iterate over the raw logs and unpacked data for OracleSizeIncreased events raised by the Legacypair contract.
type LegacypairOracleSizeIncreasedIterator struct {
	Event *LegacypairOracleSizeIncreased // Event containing the contract specifics and raw log

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
func (it *LegacypairOracleSizeIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacypairOracleSizeIncreased)
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
		it.Event = new(LegacypairOracleSizeIncreased)
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
func (it *LegacypairOracleSizeIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacypairOracleSizeIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacypairOracleSizeIncreased represents a OracleSizeIncreased event raised by the Legacypair contract.
type LegacypairOracleSizeIncreased struct {
	PreviousSize *big.Int
	NewSize      *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOracleSizeIncreased is a free log retrieval operation binding the contract event 0x525a4241308ea122822834c841f67b00d5efc977ad9118724750f974f7f6531c.
//
// Solidity: event OracleSizeIncreased(uint256 previousSize, uint256 newSize)
func (_Legacypair *LegacypairFilterer) FilterOracleSizeIncreased(opts *bind.FilterOpts) (*LegacypairOracleSizeIncreasedIterator, error) {

	logs, sub, err := _Legacypair.contract.FilterLogs(opts, "OracleSizeIncreased")
	if err != nil {
		return nil, err
	}
	return &LegacypairOracleSizeIncreasedIterator{contract: _Legacypair.contract, event: "OracleSizeIncreased", logs: logs, sub: sub}, nil
}

// WatchOracleSizeIncreased is a free log subscription operation binding the contract event 0x525a4241308ea122822834c841f67b00d5efc977ad9118724750f974f7f6531c.
//
// Solidity: event OracleSizeIncreased(uint256 previousSize, uint256 newSize)
func (_Legacypair *LegacypairFilterer) WatchOracleSizeIncreased(opts *bind.WatchOpts, sink chan<- *LegacypairOracleSizeIncreased) (event.Subscription, error) {

	logs, sub, err := _Legacypair.contract.WatchLogs(opts, "OracleSizeIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacypairOracleSizeIncreased)
				if err := _Legacypair.contract.UnpackLog(event, "OracleSizeIncreased", log); err != nil {
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

// ParseOracleSizeIncreased is a log parse operation binding the contract event 0x525a4241308ea122822834c841f67b00d5efc977ad9118724750f974f7f6531c.
//
// Solidity: event OracleSizeIncreased(uint256 previousSize, uint256 newSize)
func (_Legacypair *LegacypairFilterer) ParseOracleSizeIncreased(log types.Log) (*LegacypairOracleSizeIncreased, error) {
	event := new(LegacypairOracleSizeIncreased)
	if err := _Legacypair.contract.UnpackLog(event, "OracleSizeIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacypairProtocolFeesCollectedIterator is returned from FilterProtocolFeesCollected and is used to iterate over the raw logs and unpacked data for ProtocolFeesCollected events raised by the Legacypair contract.
type LegacypairProtocolFeesCollectedIterator struct {
	Event *LegacypairProtocolFeesCollected // Event containing the contract specifics and raw log

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
func (it *LegacypairProtocolFeesCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacypairProtocolFeesCollected)
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
		it.Event = new(LegacypairProtocolFeesCollected)
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
func (it *LegacypairProtocolFeesCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacypairProtocolFeesCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacypairProtocolFeesCollected represents a ProtocolFeesCollected event raised by the Legacypair contract.
type LegacypairProtocolFeesCollected struct {
	Sender    common.Address
	Recipient common.Address
	AmountX   *big.Int
	AmountY   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeesCollected is a free log retrieval operation binding the contract event 0x26b782206d6b531bf95d487110cfefdc443291f176f1977e94abcb7e67bd1b79.
//
// Solidity: event ProtocolFeesCollected(address indexed sender, address indexed recipient, uint256 amountX, uint256 amountY)
func (_Legacypair *LegacypairFilterer) FilterProtocolFeesCollected(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*LegacypairProtocolFeesCollectedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Legacypair.contract.FilterLogs(opts, "ProtocolFeesCollected", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &LegacypairProtocolFeesCollectedIterator{contract: _Legacypair.contract, event: "ProtocolFeesCollected", logs: logs, sub: sub}, nil
}

// WatchProtocolFeesCollected is a free log subscription operation binding the contract event 0x26b782206d6b531bf95d487110cfefdc443291f176f1977e94abcb7e67bd1b79.
//
// Solidity: event ProtocolFeesCollected(address indexed sender, address indexed recipient, uint256 amountX, uint256 amountY)
func (_Legacypair *LegacypairFilterer) WatchProtocolFeesCollected(opts *bind.WatchOpts, sink chan<- *LegacypairProtocolFeesCollected, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Legacypair.contract.WatchLogs(opts, "ProtocolFeesCollected", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacypairProtocolFeesCollected)
				if err := _Legacypair.contract.UnpackLog(event, "ProtocolFeesCollected", log); err != nil {
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

// ParseProtocolFeesCollected is a log parse operation binding the contract event 0x26b782206d6b531bf95d487110cfefdc443291f176f1977e94abcb7e67bd1b79.
//
// Solidity: event ProtocolFeesCollected(address indexed sender, address indexed recipient, uint256 amountX, uint256 amountY)
func (_Legacypair *LegacypairFilterer) ParseProtocolFeesCollected(log types.Log) (*LegacypairProtocolFeesCollected, error) {
	event := new(LegacypairProtocolFeesCollected)
	if err := _Legacypair.contract.UnpackLog(event, "ProtocolFeesCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacypairSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Legacypair contract.
type LegacypairSwapIterator struct {
	Event *LegacypairSwap // Event containing the contract specifics and raw log

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
func (it *LegacypairSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacypairSwap)
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
		it.Event = new(LegacypairSwap)
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
func (it *LegacypairSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacypairSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacypairSwap represents a Swap event raised by the Legacypair contract.
type LegacypairSwap struct {
	Sender                common.Address
	Recipient             common.Address
	Id                    *big.Int
	SwapForY              bool
	AmountIn              *big.Int
	AmountOut             *big.Int
	VolatilityAccumulated *big.Int
	Fees                  *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xc528cda9e500228b16ce84fadae290d9a49aecb17483110004c5af0a07f6fd73.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, uint256 indexed id, bool swapForY, uint256 amountIn, uint256 amountOut, uint256 volatilityAccumulated, uint256 fees)
func (_Legacypair *LegacypairFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address, id []*big.Int) (*LegacypairSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Legacypair.contract.FilterLogs(opts, "Swap", senderRule, recipientRule, idRule)
	if err != nil {
		return nil, err
	}
	return &LegacypairSwapIterator{contract: _Legacypair.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xc528cda9e500228b16ce84fadae290d9a49aecb17483110004c5af0a07f6fd73.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, uint256 indexed id, bool swapForY, uint256 amountIn, uint256 amountOut, uint256 volatilityAccumulated, uint256 fees)
func (_Legacypair *LegacypairFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *LegacypairSwap, sender []common.Address, recipient []common.Address, id []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Legacypair.contract.WatchLogs(opts, "Swap", senderRule, recipientRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacypairSwap)
				if err := _Legacypair.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xc528cda9e500228b16ce84fadae290d9a49aecb17483110004c5af0a07f6fd73.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, uint256 indexed id, bool swapForY, uint256 amountIn, uint256 amountOut, uint256 volatilityAccumulated, uint256 fees)
func (_Legacypair *LegacypairFilterer) ParseSwap(log types.Log) (*LegacypairSwap, error) {
	event := new(LegacypairSwap)
	if err := _Legacypair.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacypairTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the Legacypair contract.
type LegacypairTransferBatchIterator struct {
	Event *LegacypairTransferBatch // Event containing the contract specifics and raw log

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
func (it *LegacypairTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacypairTransferBatch)
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
		it.Event = new(LegacypairTransferBatch)
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
func (it *LegacypairTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacypairTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacypairTransferBatch represents a TransferBatch event raised by the Legacypair contract.
type LegacypairTransferBatch struct {
	Sender  common.Address
	From    common.Address
	To      common.Address
	Ids     []*big.Int
	Amounts []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed sender, address indexed from, address indexed to, uint256[] ids, uint256[] amounts)
func (_Legacypair *LegacypairFilterer) FilterTransferBatch(opts *bind.FilterOpts, sender []common.Address, from []common.Address, to []common.Address) (*LegacypairTransferBatchIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Legacypair.contract.FilterLogs(opts, "TransferBatch", senderRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LegacypairTransferBatchIterator{contract: _Legacypair.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed sender, address indexed from, address indexed to, uint256[] ids, uint256[] amounts)
func (_Legacypair *LegacypairFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *LegacypairTransferBatch, sender []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Legacypair.contract.WatchLogs(opts, "TransferBatch", senderRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacypairTransferBatch)
				if err := _Legacypair.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed sender, address indexed from, address indexed to, uint256[] ids, uint256[] amounts)
func (_Legacypair *LegacypairFilterer) ParseTransferBatch(log types.Log) (*LegacypairTransferBatch, error) {
	event := new(LegacypairTransferBatch)
	if err := _Legacypair.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacypairTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the Legacypair contract.
type LegacypairTransferSingleIterator struct {
	Event *LegacypairTransferSingle // Event containing the contract specifics and raw log

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
func (it *LegacypairTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacypairTransferSingle)
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
		it.Event = new(LegacypairTransferSingle)
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
func (it *LegacypairTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacypairTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacypairTransferSingle represents a TransferSingle event raised by the Legacypair contract.
type LegacypairTransferSingle struct {
	Sender common.Address
	From   common.Address
	To     common.Address
	Id     *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed sender, address indexed from, address indexed to, uint256 id, uint256 amount)
func (_Legacypair *LegacypairFilterer) FilterTransferSingle(opts *bind.FilterOpts, sender []common.Address, from []common.Address, to []common.Address) (*LegacypairTransferSingleIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Legacypair.contract.FilterLogs(opts, "TransferSingle", senderRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LegacypairTransferSingleIterator{contract: _Legacypair.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed sender, address indexed from, address indexed to, uint256 id, uint256 amount)
func (_Legacypair *LegacypairFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *LegacypairTransferSingle, sender []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Legacypair.contract.WatchLogs(opts, "TransferSingle", senderRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacypairTransferSingle)
				if err := _Legacypair.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed sender, address indexed from, address indexed to, uint256 id, uint256 amount)
func (_Legacypair *LegacypairFilterer) ParseTransferSingle(log types.Log) (*LegacypairTransferSingle, error) {
	event := new(LegacypairTransferSingle)
	if err := _Legacypair.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacypairWithdrawnFromBinIterator is returned from FilterWithdrawnFromBin and is used to iterate over the raw logs and unpacked data for WithdrawnFromBin events raised by the Legacypair contract.
type LegacypairWithdrawnFromBinIterator struct {
	Event *LegacypairWithdrawnFromBin // Event containing the contract specifics and raw log

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
func (it *LegacypairWithdrawnFromBinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacypairWithdrawnFromBin)
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
		it.Event = new(LegacypairWithdrawnFromBin)
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
func (it *LegacypairWithdrawnFromBinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacypairWithdrawnFromBinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacypairWithdrawnFromBin represents a WithdrawnFromBin event raised by the Legacypair contract.
type LegacypairWithdrawnFromBin struct {
	Sender    common.Address
	Recipient common.Address
	Id        *big.Int
	AmountX   *big.Int
	AmountY   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnFromBin is a free log retrieval operation binding the contract event 0xda5e7177dface55f5e0eff7dfc67420a1db4243ddfcf0ecc84ed93e034dd8cc2.
//
// Solidity: event WithdrawnFromBin(address indexed sender, address indexed recipient, uint256 indexed id, uint256 amountX, uint256 amountY)
func (_Legacypair *LegacypairFilterer) FilterWithdrawnFromBin(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address, id []*big.Int) (*LegacypairWithdrawnFromBinIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Legacypair.contract.FilterLogs(opts, "WithdrawnFromBin", senderRule, recipientRule, idRule)
	if err != nil {
		return nil, err
	}
	return &LegacypairWithdrawnFromBinIterator{contract: _Legacypair.contract, event: "WithdrawnFromBin", logs: logs, sub: sub}, nil
}

// WatchWithdrawnFromBin is a free log subscription operation binding the contract event 0xda5e7177dface55f5e0eff7dfc67420a1db4243ddfcf0ecc84ed93e034dd8cc2.
//
// Solidity: event WithdrawnFromBin(address indexed sender, address indexed recipient, uint256 indexed id, uint256 amountX, uint256 amountY)
func (_Legacypair *LegacypairFilterer) WatchWithdrawnFromBin(opts *bind.WatchOpts, sink chan<- *LegacypairWithdrawnFromBin, sender []common.Address, recipient []common.Address, id []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Legacypair.contract.WatchLogs(opts, "WithdrawnFromBin", senderRule, recipientRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacypairWithdrawnFromBin)
				if err := _Legacypair.contract.UnpackLog(event, "WithdrawnFromBin", log); err != nil {
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

// ParseWithdrawnFromBin is a log parse operation binding the contract event 0xda5e7177dface55f5e0eff7dfc67420a1db4243ddfcf0ecc84ed93e034dd8cc2.
//
// Solidity: event WithdrawnFromBin(address indexed sender, address indexed recipient, uint256 indexed id, uint256 amountX, uint256 amountY)
func (_Legacypair *LegacypairFilterer) ParseWithdrawnFromBin(log types.Log) (*LegacypairWithdrawnFromBin, error) {
	event := new(LegacypairWithdrawnFromBin)
	if err := _Legacypair.contract.UnpackLog(event, "WithdrawnFromBin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
