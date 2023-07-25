package traderjoe

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/exidz/traderjoe-fiber/app/constants"
	"github.com/exidz/traderjoe-fiber/app/helpers"
	joefactory "github.com/exidz/traderjoe-fiber/app/services/JoeFactory"
	joepair "github.com/exidz/traderjoe-fiber/app/services/JoePair"
	joerouter "github.com/exidz/traderjoe-fiber/app/services/JoeRouter"
	"github.com/exidz/traderjoe-fiber/app/types"
	"github.com/exidz/traderjoe-fiber/config"
	"github.com/exidz/traderjoe-fiber/utils"
	"github.com/gofiber/fiber/v2"
)

// JoePrice retrieves the price data for a traderjoe v1 pool pair on different chains.
//
// @Summary Get v1 pool pair price
// @Description Retrieves the price data for a v1 pool pair on different chains.
// @Tags V1 Get Pool price
// @Accept json
// @Produce json
// @Param base path string true "The base asset contract address"
// @Param quote path string true "The quote asset contract address"
// @Param chain path string true "The chain name (avax, arb, bsc)"
// @Success 200 {object} types.Response
// @Failure 400 {object} types.ErrorResponse
// @Failure 404 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /{chain}/v1/prices/{base}/{quote} [get]
func JoePrice(c *fiber.Ctx) error {

	baseAsset := c.AllParams()["base"]
	quoteAsset := c.AllParams()["quote"]
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

	if !common.IsHexAddress(baseAsset) {
		return c.Status(400).JSON(types.ErrorResponse{
			Error: "Base asset param is invalid, must be a contract address",
		})

	} else if !common.IsHexAddress(quoteAsset) {
		return c.Status(400).JSON(types.ErrorResponse{
			Error: "Quote asset param is invalid, must be a contract address",
		})

	}

	client, err := ethclient.Dial(rpcUrl)

	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{
			Error: "RPC Timeout: cannot dial rpc",
		})
	}
	defer client.Close()
	factory, err := joefactory.JoeFactoryInit(client, constants.Chain(chainId))

	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	pair, err := factory.GetJoePair(common.HexToAddress(baseAsset), common.HexToAddress(quoteAsset))

	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}
	initPair, err := joepair.InitJoePair(*pair, client)

	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	router, err := joerouter.JoeRouterInit(client, constants.Chain(chainId))
	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	pairToken, err := helpers.GetPairData(*initPair)
	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	if pairToken.Reserve.ReserveX == big.NewInt(0) || pairToken.Reserve.ReserveY == big.NewInt(0) {
		return c.Status(400).JSON(
			types.ErrorResponse{
				Error: "Pair reach the low liquidity threshold.",
			})
	}

	tokenDecimals, err := helpers.GetTokenDecimal(pairToken.Tokens.TokenX, pairToken.Tokens.TokenY, client)
	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	swapOutAndLiquidity, err := helpers.GetLiquidityAndSwapout(*tokenDecimals, *pairToken, *router)
	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	reserveA := utils.ToDecimal(pairToken.Reserve.ReserveX, int(tokenDecimals.BaseAssetDecimal))
	reserveB := utils.ToDecimal(pairToken.Reserve.ReserveY, int(tokenDecimals.QuoteAssetDecimal))

	liquidityErr := helpers.GetJoeTotalLiquidity(swapOutAndLiquidity.LiquidityX, swapOutAndLiquidity.LiquidityY, reserveA, reserveB, chainId)

	if liquidityErr != nil {
		return c.Status(400).JSON(
			types.ErrorResponse{
				Error: "Pair reach the low liquidity threshold.",
			})

	}

	baseQuoteWithDecimal := swapOutAndLiquidity.SwapOut.XtoY
	quoteBaseWithDecimal := swapOutAndLiquidity.SwapOut.YToX

	if common.HexToAddress(baseAsset) != pairToken.Tokens.TokenX || common.HexToAddress(quoteAsset) != pairToken.Tokens.TokenY {
		quoteAsset = pairToken.Tokens.TokenX.String()
		baseAsset = pairToken.Tokens.TokenY.String()
		x := swapOutAndLiquidity.SwapOut.XtoY
		y := swapOutAndLiquidity.SwapOut.YToX
		baseQuoteWithDecimal = y
		quoteBaseWithDecimal = x

	}

	priceResponse := types.JoePriceResponse{
		BaseToQuote: baseQuoteWithDecimal,
		QuoteToBase: quoteBaseWithDecimal,
		PairAddress: *pair,
		BaseAsset:   baseAsset,
		QuoteAsset:  quoteAsset,
	}
	defer client.Close()

	return c.Status(200).JSON(types.Response{
		Data:    priceResponse,
		Message: "Pool pair price data successfully retrieved",
		Success: true,
	})

}
