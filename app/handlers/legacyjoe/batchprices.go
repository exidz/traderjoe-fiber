package legacyjoe

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/exidz/traderjoe-fiber/app/constants"
	"github.com/exidz/traderjoe-fiber/app/helpers"
	legacyfactory "github.com/exidz/traderjoe-fiber/app/services/LegacyFactory"
	legacypair "github.com/exidz/traderjoe-fiber/app/services/LegacyPair"
	legacyquoter "github.com/exidz/traderjoe-fiber/app/services/LegacyQuoter"
	legacyrouter "github.com/exidz/traderjoe-fiber/app/services/LegacyRouter"
	"github.com/exidz/traderjoe-fiber/app/types"
	"github.com/exidz/traderjoe-fiber/config"
	"github.com/exidz/traderjoe-fiber/utils"
	"github.com/gofiber/fiber/v2"
)

func BatchLegacyPrice(c *fiber.Ctx) error {
	payload := &types.BodyData{}
	if err := c.BodyParser(payload); err != nil {
		return c.Status(400).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	serData := helpers.ArrangeBatch(*payload)
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

	quoteData := make([]interface{}, 20)
	batchLimit := 20
	if len(serData) < 20 {
		quoteData = make([]interface{}, len(serData))
		batchLimit = len(serData)

	}
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return c.Status(500).JSON(&types.ErrorResponse{
			Error: "RPC timeout: cannot dial rpc",
		})
	}

	factory, err := legacyfactory.LegacyFactoryInit(client, constants.Chain(chainId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}
	router, err := legacyrouter.InitLegacyRouter(client, constants.Chain(chainId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}
	quoteInstance, err := legacyquoter.InitLegacyQuoter(client, constants.Chain(chainId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
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

				pairInfo, _ := factory.GetLegacyPairInfo(common.HexToAddress(dataInit.Data.BaseAsset), common.HexToAddress(dataInit.Data.QuoteAsset), dataInit.Data.Binstep)

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
				pair, err := legacypair.LegacyPairInit(pairInfo.LBPair, client)
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

				tokens, err := pair.GetLegacyPairTokens()

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

				bulkQuote, err := helpers.GetLegacyPairQuote(tokenX, tokenY, decimalx, decimaly, *quoteInstance)
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

				swapOut, err := helpers.LegacyPairSwapOut(*router, decimalx, decimaly, pairInfo.LBPair)
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
