package lbjoe

import (
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/exidz/traderjoe-fiber/app/constants"
	"github.com/exidz/traderjoe-fiber/app/helpers"
	lbfactory "github.com/exidz/traderjoe-fiber/app/services/LBFactory"
	lbpair "github.com/exidz/traderjoe-fiber/app/services/LBPair"
	lbquoter "github.com/exidz/traderjoe-fiber/app/services/LBQuoter"
	"github.com/exidz/traderjoe-fiber/app/types"
	"github.com/exidz/traderjoe-fiber/config"
	"github.com/exidz/traderjoe-fiber/utils"
	"github.com/gofiber/fiber/v2"
)

// GetLbPoolPrice retrieves the price data for a traderjoe liquidity book pool pair on different chains.
//
// @Summary Get v2.1 liquidity book pool pair price
// @Description Retrieves the price data for a v2.1 liquidity book pool pair on different chains.
// @Tags V2.1 Get Liquidity Book Pool price
// @Accept json
// @Produce json
// @Param base path string true "The base asset contract address"
// @Param quote path string true "The quote asset contract address"
// @Param binstep path string true "The binstep value"
// @Param chain path string true "The chain name (avax, arb, bsc)"
// @Success 200 {object} types.Response
// @Failure 400 {object} types.ErrorResponse
// @Failure 404 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /{chain}/v2_1/prices/{base}/{quote}/{binstep} [get]
func GetLbPoolPrice(c *fiber.Ctx) error {
	baseAsset := c.AllParams()["base"]
	quoteAsset := c.AllParams()["quote"]
	binstep := c.AllParams()["binstep"]
	chain := c.AllParams()["chain"]

	var chainId int
	var rpcUrl string

	if chain == "avax" {
		rpc, err := config.GoDotEnvVariable("AVAX_RPC")
		if err != nil {
			rpcUrl = constants.AVAX_RPC
		} else {
			rpcUrl = *rpc

		}
		chainId = 0
	} else if chain == "arb" {
		rpc, err := config.GoDotEnvVariable("ARB_RPC")
		if err != nil {
			rpcUrl = constants.ARB_RPC
		} else {
			rpcUrl = *rpc

		}

		chainId = 1
	} else if chain == "bsc" {
		rpc, err := config.GoDotEnvVariable("BSC_RPC")
		if err != nil {
			rpcUrl = constants.BSC_RPC
		} else {
			rpcUrl = *rpc

		}
		chainId = 2
	} else {
		return c.Status(fiber.StatusNotFound).JSON(types.ErrorResponse{
			Error: "Endpoint not found",
		})
	}

	rpc := rpcUrl

	lowLiquidity := types.ErrorResponse{
		Error: "Pair reach the low liquidity threshold.",
	}

	if !common.IsHexAddress(baseAsset) {
		return c.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Error: "Base asset param is invalid, must be a contract address",
		})

	} else if !common.IsHexAddress(quoteAsset) {
		return c.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Error: "Quote asset param is invalid, must be a contract address",
		})

	} else if !utils.IsNumeric(binstep) {
		return c.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Error: "Binstep param is invalid, must be a integer/number",
		})

	}

	bin, _ := strconv.Atoi(binstep)

	client, err := ethclient.Dial(rpc)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Error: "RPC timeout: cannot dial rpc",
		})
	}

	factory, err := lbfactory.CreateFactory(client, constants.Chain(chainId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	pairInfo, _ := factory.GetPairInfo(common.HexToAddress(baseAsset), common.HexToAddress(quoteAsset), *big.NewInt(int64(bin)))

	if pairInfo.LBPair == common.HexToAddress("0x0000000000000000000000000000000000000000") {
		return c.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Error: "Lb pool pair not found",
		})
	}
	pair, err := lbpair.InitPair(client, pairInfo.LBPair)
	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	reserve, err := pair.GetPairReserve()
	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	if reserve.ReserveX == big.NewInt(0) || reserve.ReserveY == big.NewInt(0) {
		return c.Status(fiber.StatusBadRequest).JSON(lowLiquidity)
	}

	tokens, err := pair.GetPairTokens()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	tokenX := tokens.BaseToken
	tokenY := tokens.QuoteToken

	pairDecimals, err := helpers.GetTokenDecimal(tokenX, tokenY, client)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}
	decimalx := pairDecimals.BaseAssetDecimal
	decimaly := pairDecimals.QuoteAssetDecimal

	quoteInstance, err := lbquoter.QuoterInit(client, constants.Chain(chainId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	bulkQuote, err := helpers.GetPairQuote(tokenX, tokenY, decimalx, decimaly, *quoteInstance)
	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	tokenAQuote := bulkQuote.BaseToQuote
	tokenBQuote := bulkQuote.QuoteToBase

	reserveADecimal := utils.ToDecimal(reserve.ReserveX, int(decimalx))
	reserveBDecimal := utils.ToDecimal(reserve.ReserveY, int(decimaly))

	lowLiquidityPair := helpers.GetTotalLiquidity(tokenAQuote, tokenBQuote, reserveADecimal, reserveBDecimal, chainId)

	if lowLiquidityPair != nil {
		return c.Status(fiber.StatusBadRequest).JSON(lowLiquidity)
	}

	swapOut, err := helpers.PairSwapOut(*pair, decimalx, decimaly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	baseToQuote := swapOut.BaseToQuote
	quoteToBase := swapOut.QuoteToBase

	baseQuoteWithDecimal := utils.ToDecimal(baseToQuote.AmountOut, int(decimaly))
	quoteBaseWithDecimal := utils.ToDecimal(quoteToBase.AmountOut, int(decimalx))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	if common.HexToAddress(baseAsset) != tokenX || common.HexToAddress(quoteAsset) != tokenY {
		quoteAsset = tokenX.String()
		baseAsset = tokenY.String()
		x := baseQuoteWithDecimal
		y := quoteBaseWithDecimal
		baseQuoteWithDecimal = y
		quoteBaseWithDecimal = x

	}

	response := types.PriceResponse{
		BaseToQuote: baseQuoteWithDecimal,
		QuoteToBase: quoteBaseWithDecimal,
		PairAddress: pairInfo.LBPair,
		BaseAsset:   baseAsset,
		QuoteAsset:  quoteAsset,
		BinStep:     bin,
	}

	defer client.Close()
	return c.Status(fiber.StatusOK).JSON(types.Response{
		Data:    response,
		Message: "Pool pair price data successfully retrieved",
		Success: true,
	})
}
