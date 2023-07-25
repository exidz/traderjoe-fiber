package traderjoe

import (
	"math/big"
	"sync"

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

// BatchJoePrice retrieves v1 batch pool prices for multiple pairs on different chains using the JoeSwap protocol.
//
// @Summary Get batch pool prices
// @Description Retrieves batch pool prices for multiple pairs on different chains using the JoeSwap protocol.
// @Tags v1 Batch pool price
// @Accept json
// @Produce json
// @Param payload body types.JoeBodyData true "Batch pool request data"
// @Param chain path string true "The chain name (avax, arb, bsc)"
// @Success 200 {object} types.Response
// @Failure 400 {object} types.JoeBatchErrorResponse
// @Failure 404 {object} types.JoeBatchErrorResponse
// @Failure 500 {object} types.JoeBatchErrorResponse
// @Router /{chain}/v1/batch-prices [post]
func BatchJoePrice(c *fiber.Ctx) error {
	payload := &types.JoeBodyData{}
	if err := c.BodyParser(payload); err != nil {
		return c.Status(400).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}
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

	serData := helpers.JoeArrangeBatch(*payload)
	quoteData := make([]interface{}, 20)
	batchLimit := 20
	if len(serData) < 20 {
		quoteData = make([]interface{}, len(serData))
		batchLimit = len(serData)

	}
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{
			Error: "RPC timeout: cannot dial rpc",
		})
	}
	defer client.Close()

	factory, err := joefactory.JoeFactoryInit(client, constants.Chain(chainId))

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

	var wg sync.WaitGroup

	for i := 0; i < batchLimit; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			dataInit := serData[i]
			if dataInit.Error {
				quoteData[i] = types.JoeBatchErrorResponse{
					Error: dataInit.ErrorMessage,
					Params: types.JoeBatchParams{
						BaseAsset:  dataInit.Data.BaseAsset,
						QuoteAsset: dataInit.Data.QuoteAsset,
					},
				}
				return
			} else {
				pair, err := factory.GetJoePair(common.HexToAddress(dataInit.Data.BaseAsset), common.HexToAddress(dataInit.Data.QuoteAsset))

				if err != nil {
					quoteData[i] = types.JoeBatchErrorResponse{
						Error: dataInit.ErrorMessage,
						Params: types.JoeBatchParams{
							BaseAsset:  dataInit.Data.BaseAsset,
							QuoteAsset: dataInit.Data.QuoteAsset,
						},
					}
					return
				}
				initPair, err := joepair.InitJoePair(*pair, client)

				if err != nil {
					quoteData[i] = types.JoeBatchErrorResponse{
						Error: err.Error(),
						Params: types.JoeBatchParams{
							BaseAsset:  dataInit.Data.BaseAsset,
							QuoteAsset: dataInit.Data.QuoteAsset,
						},
					}
					return
				}

				pairToken, err := helpers.GetPairData(*initPair)
				if err != nil {
					quoteData[i] = types.JoeBatchErrorResponse{
						Error: err.Error(),
						Params: types.JoeBatchParams{
							BaseAsset:  dataInit.Data.BaseAsset,
							QuoteAsset: dataInit.Data.QuoteAsset,
						},
					}
					return
				}

				if pairToken.Reserve.ReserveX == big.NewInt(0) || pairToken.Reserve.ReserveY == big.NewInt(0) {
					quoteData[i] = types.JoeBatchErrorResponse{
						Error: err.Error(),
						Params: types.JoeBatchParams{
							BaseAsset:  dataInit.Data.BaseAsset,
							QuoteAsset: dataInit.Data.QuoteAsset,
						},
					}
					return
				}

				tokenDecimals, err := helpers.GetTokenDecimal(pairToken.Tokens.TokenX, pairToken.Tokens.TokenY, client)
				if err != nil {
					quoteData[i] = types.JoeBatchErrorResponse{
						Error: err.Error(),
						Params: types.JoeBatchParams{
							BaseAsset:  dataInit.Data.BaseAsset,
							QuoteAsset: dataInit.Data.QuoteAsset,
						},
					}
					return
				}

				swapOutAndLiquidity, err := helpers.GetLiquidityAndSwapout(*tokenDecimals, *pairToken, *router)
				if err != nil {
					quoteData[i] = types.JoeBatchErrorResponse{
						Error: err.Error(),
						Params: types.JoeBatchParams{
							BaseAsset:  dataInit.Data.BaseAsset,
							QuoteAsset: dataInit.Data.QuoteAsset,
						},
					}
					return
				}

				reserveA := utils.ToDecimal(pairToken.Reserve.ReserveX, int(tokenDecimals.BaseAssetDecimal))
				reserveB := utils.ToDecimal(pairToken.Reserve.ReserveY, int(tokenDecimals.QuoteAssetDecimal))

				liquidityErr := helpers.GetJoeTotalLiquidity(swapOutAndLiquidity.LiquidityX, swapOutAndLiquidity.LiquidityY, reserveA, reserveB, chainId)

				if liquidityErr != nil {

					quoteData[i] = types.JoeBatchErrorResponse{
						Error: "Pair reach the low liquidity threshold.",
						Params: types.JoeBatchParams{
							BaseAsset:  dataInit.Data.BaseAsset,
							QuoteAsset: dataInit.Data.QuoteAsset,
						},
					}
					return

				}

				baseQuoteWithDecimal := swapOutAndLiquidity.SwapOut.XtoY
				quoteBaseWithDecimal := swapOutAndLiquidity.SwapOut.YToX
				quoteAsset := dataInit.Data.QuoteAsset
				baseAsset := dataInit.Data.BaseAsset

				if common.HexToAddress(dataInit.Data.BaseAsset) != pairToken.Tokens.TokenX || common.HexToAddress(dataInit.Data.QuoteAsset) != pairToken.Tokens.TokenY {
					quoteAsset = pairToken.Tokens.TokenX.String()
					baseAsset = pairToken.Tokens.TokenY.String()
					x := swapOutAndLiquidity.SwapOut.XtoY
					y := swapOutAndLiquidity.SwapOut.YToX
					baseQuoteWithDecimal = y
					quoteBaseWithDecimal = x

				}

				quoteData[i] = types.JoeBatchSuccessResponse{
					BaseToQuote: *baseQuoteWithDecimal,
					QuoteToBase: *quoteBaseWithDecimal,
					PairAddress: *pair,
					BaseAsset:   baseAsset,
					QuoteAsset:  quoteAsset,
					Params: types.JoeBatchParams{
						BaseAsset:  dataInit.Data.BaseAsset,
						QuoteAsset: dataInit.Data.QuoteAsset,
					},
				}

			}

		}(i)
	}

	wg.Wait()

	return c.Status(200).JSON(types.Response{
		Data:    quoteData,
		Message: "Batch pool price retrieved successfully",
		Success: true,
	})

}
