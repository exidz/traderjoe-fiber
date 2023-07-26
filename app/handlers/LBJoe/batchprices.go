package lbjoe

import (
	"math/big"
	"os"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/exidz/traderjoe-fiber/app/constants"
	"github.com/exidz/traderjoe-fiber/app/helpers"
	lbfactory "github.com/exidz/traderjoe-fiber/app/services/LBFactory"
	lbpair "github.com/exidz/traderjoe-fiber/app/services/LBPair"
	lbquoter "github.com/exidz/traderjoe-fiber/app/services/LBQuoter"
	"github.com/exidz/traderjoe-fiber/app/types"
	"github.com/exidz/traderjoe-fiber/utils"
	"github.com/gofiber/fiber/v2"
)

// BatchLBPrice retrieves batch pool prices for multiple pairs on different chains.
// @Summary Get batch v2.1 liquidity book pool prices
// @Description Retrieves batch pool prices for multiple pairs on different chains.
// @Tags V2.1 Batch Liquidity Book Pool price
// @Accept json
// @Produce json
// @Param payload body types.BodyData true "Batch Liquidity book pool prices request data"
// @Param chain path string true "Chain Name (avax, arb or bsc)"
// @Success 200 {object} types.BatchResponse
// @Failure 400 {object} types.BatchErrorResponse
// @Failure 404 {object} types.BatchErrorResponse
// @Failure 500 {object} types.BatchErrorResponsefmt
// @Router /{chain}/v2_1/batch-prices [post]
func BatchLBPrice(c *fiber.Ctx) error {
	payload := &types.BodyData{}
	if err := c.BodyParser(payload); err != nil {
		return c.Status(400).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}
	chain := c.AllParams()["chain"]
	var chainId int
	var rpcUrl string

	if chain == "avax" {
		rpcUrl = os.Getenv("AVAX_RPC")
		chainId = 0
	} else if chain == "arb" {
		rpcUrl = os.Getenv("ARB_RPC")
		chainId = 1
	} else if chain == "bsc" {
		rpcUrl = os.Getenv("BSC_RPC")
		chainId = 2
	} else {
		return c.Status(fiber.StatusNotFound).JSON(types.ErrorResponse{
			Error: "Endpoint not found",
		})
	}

	rpc := rpcUrl

	serData := helpers.ArrangeBatch(*payload)

	quoteData := make([]interface{}, 20)
	batchLimit := 20
	if len(serData) < 20 {
		quoteData = make([]interface{}, len(serData))
		batchLimit = len(serData)

	}
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return c.Status(500).JSON(&types.ErrorResponse{
			Error: "RPC timeout: cannot dial rpc",
		})
	}

	factory, err := lbfactory.CreateFactory(client, constants.Chain(chainId))
	if err != nil {
		return c.Status(500).JSON(&types.ErrorResponse{
			Error: err.Error(),
		})
	}
	quoteInstance, err := lbquoter.QuoterInit(client, constants.Chain(chainId))
	if err != nil {
		return c.Status(500).JSON(&types.ErrorResponse{
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
				quoteData[i] = types.BatchErrorResponse{
					Error: dataInit.ErrorMessage,
					Params: types.BatchParams{
						BaseAsset:  dataInit.Data.BaseAsset,
						QuoteAsset: dataInit.Data.QuoteAsset,
						BinStep:    dataInit.Data.Binstep,
					},
				}
				return
			} else {

				pairInfo, _ := factory.GetPairInfo(common.HexToAddress(dataInit.Data.BaseAsset), common.HexToAddress(dataInit.Data.QuoteAsset), *big.NewInt(int64(dataInit.Data.Binstep)))

				if pairInfo.LBPair == common.HexToAddress("0x0000000000000000000000000000000000000000") {

					quoteData[i] = types.BatchErrorResponse{
						Error: "Lb pool pair not found",
						Params: types.BatchParams{
							BaseAsset:  dataInit.Data.BaseAsset,
							QuoteAsset: dataInit.Data.QuoteAsset,
							BinStep:    dataInit.Data.Binstep,
						},
					}
					return
				}
				pair, err := lbpair.InitPair(client, pairInfo.LBPair)
				if err != nil {
					quoteData[i] = types.BatchErrorResponse{
						Error: err.Error(),
						Params: types.BatchParams{
							BaseAsset:  dataInit.Data.BaseAsset,
							QuoteAsset: dataInit.Data.QuoteAsset,
							BinStep:    dataInit.Data.Binstep,
						},
					}
					return
				}

				reserve, err := pair.GetPairReserve()
				if err != nil {
					quoteData[i] = types.BatchErrorResponse{
						Error: err.Error(),
						Params: types.BatchParams{
							BaseAsset:  dataInit.Data.BaseAsset,
							QuoteAsset: dataInit.Data.QuoteAsset,
							BinStep:    dataInit.Data.Binstep,
						},
					}
					return
				}

				if reserve.ReserveX == big.NewInt(0) || reserve.ReserveY == big.NewInt(0) {
					quoteData[i] = types.BatchErrorResponse{
						Error: "Pair reach the low liquidity threshold.",
						Params: types.BatchParams{
							BaseAsset:  dataInit.Data.BaseAsset,
							QuoteAsset: dataInit.Data.QuoteAsset,
							BinStep:    dataInit.Data.Binstep,
						},
					}
					return
				}

				tokens, err := pair.GetPairTokens()

				if err != nil {
					quoteData[i] = types.BatchErrorResponse{
						Error: err.Error(),
						Params: types.BatchParams{
							BaseAsset:  dataInit.Data.BaseAsset,
							QuoteAsset: dataInit.Data.QuoteAsset,
							BinStep:    dataInit.Data.Binstep,
						},
					}
					return
				}

				tokenX := tokens.BaseToken
				tokenY := tokens.QuoteToken

				pairDecimals, err := helpers.GetTokenDecimal(tokenX, tokenY, client)

				if err != nil {
					quoteData[i] = types.BatchErrorResponse{
						Error: err.Error(),
						Params: types.BatchParams{
							BaseAsset:  dataInit.Data.BaseAsset,
							QuoteAsset: dataInit.Data.QuoteAsset,
							BinStep:    dataInit.Data.Binstep,
						},
					}
					return
				}
				decimalx := pairDecimals.BaseAssetDecimal
				decimaly := pairDecimals.QuoteAssetDecimal

				bulkQuote, err := helpers.GetPairQuote(tokenX, tokenY, decimalx, decimaly, *quoteInstance)
				if err != nil {
					quoteData[i] = types.BatchErrorResponse{
						Error: err.Error(),
						Params: types.BatchParams{
							BaseAsset:  dataInit.Data.BaseAsset,
							QuoteAsset: dataInit.Data.QuoteAsset,
							BinStep:    dataInit.Data.Binstep,
						},
					}
					return
				}
				tokenAQuote := bulkQuote.BaseToQuote
				tokenBQuote := bulkQuote.QuoteToBase

				reserveADecimal := utils.ToDecimal(reserve.ReserveX, int(decimalx))
				reserveBDecimal := utils.ToDecimal(reserve.ReserveY, int(decimaly))

				lowLiquidityPair := helpers.GetTotalLiquidity(tokenAQuote, tokenBQuote, reserveADecimal, reserveBDecimal, chainId)

				if lowLiquidityPair != nil {
					quoteData[i] = types.BatchErrorResponse{
						Error: "Pair reach the low liquidity threshold.",
						Params: types.BatchParams{
							BaseAsset:  dataInit.Data.BaseAsset,
							QuoteAsset: dataInit.Data.QuoteAsset,
							BinStep:    dataInit.Data.Binstep,
						},
					}
					return
				}

				swapOut, err := helpers.PairSwapOut(*pair, decimalx, decimaly)
				if err != nil {
					quoteData[i] = types.BatchErrorResponse{
						Error: err.Error(),
						Params: types.BatchParams{
							BaseAsset:  dataInit.Data.BaseAsset,
							QuoteAsset: dataInit.Data.QuoteAsset,
							BinStep:    dataInit.Data.Binstep,
						},
					}
					return
				}

				baseToQuote := swapOut.BaseToQuote
				quoteToBase := swapOut.QuoteToBase

				baseQuoteWithDecimal := utils.ToDecimal(baseToQuote.AmountOut, int(decimaly))
				quoteBaseWithDecimal := utils.ToDecimal(quoteToBase.AmountOut, int(decimalx))

				if err != nil {
					quoteData[i] = types.BatchErrorResponse{
						Error: err.Error(),
						Params: types.BatchParams{
							BaseAsset:  dataInit.Data.BaseAsset,
							QuoteAsset: dataInit.Data.QuoteAsset,
							BinStep:    dataInit.Data.Binstep,
						},
					}
					return
				}
				quoteAsset := tokenY.String()
				baseAsset := tokenX.String()

				if common.HexToAddress(dataInit.Data.BaseAsset) != tokenX || common.HexToAddress(dataInit.Data.QuoteAsset) != tokenY {
					quoteAsset = tokenX.String()
					baseAsset = tokenY.String()
					x := baseQuoteWithDecimal
					y := quoteBaseWithDecimal
					baseQuoteWithDecimal = y
					quoteBaseWithDecimal = x

				}

				quoteData[i] = types.BatchSuccessResponse{
					BaseToQuote: baseQuoteWithDecimal,
					QuoteToBase: quoteBaseWithDecimal,
					PairAddress: pairInfo.LBPair,
					QuoteAsset:  quoteAsset,
					BaseAsset:   baseAsset,
					BinStep:     dataInit.Data.Binstep,

					Params: types.BatchParams{
						BaseAsset:  dataInit.Data.BaseAsset,
						QuoteAsset: dataInit.Data.QuoteAsset,
						BinStep:    dataInit.Data.Binstep,
					},
				}

			}

		}(i)

	}
	wg.Wait()

	return c.Status(200).JSON(types.BatchResponse{
		Message: "Batch pool price retrieved successfully",
		Success: true,
		Data:    quoteData,
	})

}
