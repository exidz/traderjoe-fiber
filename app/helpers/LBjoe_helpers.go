package helpers

import (
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	lbpair "github.com/exidz/traderjoe-fiber/app/services/LBPair"
	lbquoter "github.com/exidz/traderjoe-fiber/app/services/LBQuoter"
	token "github.com/exidz/traderjoe-fiber/app/services/Token"
	"github.com/exidz/traderjoe-fiber/app/types"
	"github.com/exidz/traderjoe-fiber/utils"
	"github.com/shopspring/decimal"
)

func getQuoteUnderlying(tokenX common.Address, decimalX uint8, quoteInstance *lbquoter.Quoter) (*types.Quote, error) {
	amountX := utils.ToWei(1.0, int(decimalX))

	var quote types.Quote

	q1, err := quoteInstance.GetQuote(tokenX, amountX)
	if err != nil {
		return nil, err
	}

loop:
	for i := 1; i < len(q1.Quote.VirtualAmountsWithoutSlippage)-1; i++ {
		v := q1.Quote.VirtualAmountsWithoutSlippage[i]

		if q1.Chain == 0 || q1.Chain == 1 {
			if v != big.NewInt(0) {
				decimal := 0
				if i == 1 || i == 2 {
					decimal = 6
				} else {
					decimal = 18
				}
				fee := q1.Quote.Fees[i-1]

				feeInDecimal := utils.ToDecimal(fee, int(18))
				amountX := utils.ToDecimal(v, int(decimal))

				amountOutWithoutFee := feeInDecimal.Add(amountX)

				quote = types.Quote{
					AmountOut: amountX,
					Fee:       feeInDecimal,
					TotalAmt:  amountOutWithoutFee,
					Token:     i,
				}
				break loop
			}
		} else {
			if v != big.NewInt(0) {
				decimal := 18
				fee := q1.Quote.Fees[i-1]

				feeInDecimal := utils.ToDecimal(fee, int(18))
				amountX := utils.ToDecimal(v, int(decimal))

				amountOutWithoutFee := feeInDecimal.Add(amountX)

				quote = types.Quote{
					AmountOut: amountX,
					Fee:       feeInDecimal,
					TotalAmt:  amountOutWithoutFee,
					Token:     i,
				}
				break loop
			}
		}
	}

	return &quote, nil
}

func GetTotalLiquidity(tokenAQuote types.Quote, tokenBQuote types.Quote, reserveA decimal.Decimal, reserveB decimal.Decimal, chain int) error {

	var lowLiquidityTresholdNative decimal.Decimal
	lowLiquidityTresholdStable := decimal.NewFromInt(10)

	if chain == 0 {
		lowLiquidityTresholdNative = decimal.NewFromInt(1)

	} else if chain == 2 {
		lowLiquidityTresholdNative = decimal.NewFromFloat(0.01)
	} else {
		lowLiquidityTresholdNative = decimal.NewFromFloat(0.15)
	}

	if tokenAQuote.AmountOut == decimal.NewFromInt(0) && tokenBQuote.AmountOut == decimal.NewFromInt(0) {
		return errors.New("Low liquidity detected")
	} else {
		if tokenAQuote.Token == 1 || tokenAQuote.Token == 2 {
			tokenAwithReserve := tokenAQuote.AmountOut.Mul(reserveA)
			if tokenAwithReserve.LessThan(lowLiquidityTresholdStable) {
				if tokenBQuote.Token == 2 || tokenBQuote.Token == 1 {
					tokenBwithReserve := tokenBQuote.AmountOut.Mul(reserveB)
					if tokenBwithReserve.LessThan(lowLiquidityTresholdStable) {
						totalL := tokenAwithReserve.Add(tokenBwithReserve)
						if totalL.LessThan(lowLiquidityTresholdStable) {
							return errors.New("Low liquidity detected")
						}
					}

				} else {
					tokenBwithReserve := tokenBQuote.AmountOut.Mul(reserveB)
					if tokenBwithReserve.LessThan(lowLiquidityTresholdNative) {
						return errors.New("Low liquidity detected")
					}
				}
			}
		} else {
			tokenAwithReserve := tokenAQuote.AmountOut.Mul(reserveA)
			if tokenAwithReserve.LessThan(lowLiquidityTresholdNative) {
				if tokenBQuote.Token == 2 || tokenBQuote.Token == 1 {
					tokenBwithReserve := tokenBQuote.AmountOut.Mul(reserveB)
					if tokenBwithReserve.LessThan(lowLiquidityTresholdStable) {
						return errors.New("Low liquidity detected")

					}

				} else {
					tokenBwithReserve := tokenBQuote.AmountOut.Mul(reserveB)
					if tokenBwithReserve.LessThan(lowLiquidityTresholdNative) {
						totalL := tokenAwithReserve.Add(tokenBwithReserve)
						if totalL.LessThan(lowLiquidityTresholdStable) {
							return errors.New("Low liquidity detected")
						}
					}
				}
			}

		}
	}

	return nil

}

func GetTokenDecimal(baseAsset common.Address, quoteAsset common.Address, client *ethclient.Client) (*types.TokenDecimals, error) {
	var xDecimal *uint8
	var yDecimal *uint8
	var errorChan error

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		y, err := token.InitToken(client, quoteAsset)

		if err != nil {
			errorChan = fmt.Errorf("RPC timeout: token contract cannot initialized: %w", err)
			return
		}

		ydec, err := y.GetDecimal()
		if err != nil {
			errorChan = fmt.Errorf("RPC timeout: cannot fetch token decimal: %w", err)
			return
		}

		yDecimal = ydec

	}()

	go func() {
		defer wg.Done()

		x, err := token.InitToken(client, baseAsset)

		if err != nil {
			errorChan = fmt.Errorf("RPC timeout: token contract cannot initialized: %w", err)
			return
		}

		xdec, err := x.GetDecimal()
		if err != nil {
			errorChan = fmt.Errorf("RPC timeout: cannot fetch token decimal: %w", err)
			return
		}

		xDecimal = xdec

	}()

	wg.Wait()

	if errorChan != nil {
		return nil, errorChan
	}

	tokenDecimal := types.TokenDecimals{
		BaseAssetDecimal:  *xDecimal,
		QuoteAssetDecimal: *yDecimal,
	}

	return &tokenDecimal, nil
}

func GetPairQuote(baseAsset common.Address, quoteAsset common.Address, baseDecimal uint8, quoteDecimal uint8, quoter lbquoter.Quoter) (*types.QuoteBulk, error) {
	var xQoute *types.Quote
	var yQuote *types.Quote
	var error error

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		base, err := getQuoteUnderlying(baseAsset, baseDecimal, &quoter)
		if err != nil {

			error = err

			return
		}
		xQoute = base

	}()

	go func() {
		defer wg.Done()
		quote, err := getQuoteUnderlying(quoteAsset, quoteDecimal, &quoter)
		if err != nil {
			error = err

			return
		}
		yQuote = quote

	}()

	wg.Wait()

	if error != nil {
		return nil, error
	}

	quoteBulk := types.QuoteBulk{
		BaseToQuote: *xQoute,
		QuoteToBase: *yQuote,
	}

	return &quoteBulk, nil

}

func PairSwapOut(pair lbpair.PairV2, decimalX uint8, decimalY uint8) (*types.PairSwapOut, error) {
	var xSwapChan lbpair.SwapOut
	var ySwapChan lbpair.SwapOut
	var erroChan error

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		baseToQuote, err := pair.GetSwapOut(decimalX, true)
		if err != nil {
			erroChan = err
			return
		}

		xSwapChan = *baseToQuote

	}()

	go func() {
		defer wg.Done()
		quoteToBase, err := pair.GetSwapOut(decimalY, false)
		if err != nil {
			erroChan = err
			return
		}

		ySwapChan = *quoteToBase

	}()

	wg.Wait()

	if erroChan != nil {
		return nil, erroChan
	}

	swapOut := types.PairSwapOut{
		BaseToQuote: xSwapChan,
		QuoteToBase: ySwapChan,
	}

	return &swapOut, nil

}

func ArrangeBatch(data types.BodyData) []types.SerializedData {
	serializedData := make([]types.SerializedData, len(data.Data))
	for i := 0; i < len(data.Data); i++ {
		req := data.Data[i]
		if !common.IsHexAddress(req.BaseAsset) {
			serializedData[i] = types.SerializedData{Error: true, Data: &req, ErrorMessage: "Base asset param is invalid, must be a contract address"}
		} else if !common.IsHexAddress(req.QuoteAsset) {
			serializedData[i] = types.SerializedData{Error: true, Data: &req, ErrorMessage: "Quote asset param is invalid, must be a contract address"}
		} else if !utils.IsInt(req.Binstep) {
			serializedData[i] = types.SerializedData{Error: true, Data: &req, ErrorMessage: "Binstep param is invalid, must be an integer/number"}
		} else {
			serializedData[i] = types.SerializedData{Error: false, Data: &req, ErrorMessage: ""}
		}
	}
	return serializedData
}
