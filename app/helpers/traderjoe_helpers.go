package helpers

import (
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	joepair "github.com/exidz/traderjoe-fiber/app/services/JoePair"
	joerouter "github.com/exidz/traderjoe-fiber/app/services/JoeRouter"
	"github.com/exidz/traderjoe-fiber/app/types"
	"github.com/exidz/traderjoe-fiber/utils"
	"github.com/shopspring/decimal"
)

func getPairSwapout(reserveX big.Int, reserveY big.Int, decimalX uint8, decimalY uint8, router joerouter.JoeRouter) (*types.JoePairSwapOut, error) {
	var swapOutX *decimal.Decimal
	var swapOutY *decimal.Decimal
	var errorChan error

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		swapOut, err := router.GetJoeRouterAmountOut(decimalX, reserveX, reserveY)

		if err != nil {
			errorChan = fmt.Errorf("RPC timeout: token contract cannot initialized: %w", err)
			return
		}

		amountOut := utils.ToDecimal(swapOut, int(decimalY))

		swapOutX = &amountOut

	}()

	go func() {
		defer wg.Done()

		swapOut, err := router.GetJoeRouterAmountOut(decimalY, reserveY, reserveX)

		if err != nil {
			errorChan = fmt.Errorf("RPC timeout: token contract cannot initialized: %w", err)
			return
		}

		amountOut := utils.ToDecimal(swapOut, int(decimalX))

		swapOutY = &amountOut

	}()

	wg.Wait()

	baseToQuote, quoteToBase, err := swapOutX, swapOutY, errorChan
	if err != nil {
		return nil, err
	}

	routerSwapout := types.JoePairSwapOut{
		XtoY: baseToQuote,
		YToX: quoteToBase,
	}

	return &routerSwapout, nil
}

func getJoeAmountOuts(tokenX common.Address, decimalX uint8, router joerouter.JoeRouter) (*types.JoeAmountOuts, error) {
	amountX := utils.ToWei(1.0, int(decimalX))
	var quote types.JoeAmountOuts

	q1, err := router.GetQuote(tokenX, amountX)
	if err != nil {
		return nil, err
	}

	quote = types.JoeAmountOuts{
		AmountOut:         q1.AmountOut,
		TokenQuoteDecimal: uint8(q1.TokenDecimal),
		TokenId:           q1.TokenId,
	}

	return &quote, nil
}

func GetPairData(pairInstance joepair.JoePair) (*types.PairInfo, error) {
	var tokensChan joepair.PairTokens
	var reserveChan joepair.JoePairReserve
	var erroChan error

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		tokens, err := pairInstance.GetJoePairTokens()
		if err != nil {
			erroChan = err
			return
		}
		tokensChan = *tokens

	}()

	go func() {
		defer wg.Done()
		reserve, err := pairInstance.GetJoePairReserve()
		if err != nil {
			erroChan = err
			return
		}
		reserveChan = *reserve

	}()

	wg.Wait()

	tokens, reserves, err := tokensChan, reserveChan, erroChan
	if err != nil {
		return nil, err
	}

	pairInfo := types.PairInfo{
		Tokens:  tokens,
		Reserve: reserves,
	}

	return &pairInfo, nil

}

func GetLiquidityAndSwapout(tokenDecimals types.TokenDecimals, pairToken types.PairInfo, router joerouter.JoeRouter) (*types.JoeSwapOutAndLiquidity, error) {

	var liquidityXChan *types.JoeAmountOuts
	var liquidityYChan *types.JoeAmountOuts
	var swapOutChan *types.JoePairSwapOut
	var errorChan error

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		liquidityX, err := getJoeAmountOuts(pairToken.Tokens.TokenX, tokenDecimals.BaseAssetDecimal, router)
		if err != nil {
			errorChan = err
			return
		}
		liquidityXChan = liquidityX

	}()

	go func() {
		defer wg.Done()
		liquidityY, err := getJoeAmountOuts(pairToken.Tokens.TokenY, tokenDecimals.QuoteAssetDecimal, router)
		if err != nil {
			errorChan = err
			return
		}
		liquidityYChan = liquidityY

	}()

	go func() {
		defer wg.Done()
		swapOut, err := getPairSwapout(*pairToken.Reserve.ReserveX, *pairToken.Reserve.ReserveY, tokenDecimals.BaseAssetDecimal, tokenDecimals.QuoteAssetDecimal, router)
		if err != nil {
			errorChan = err
			return
		}
		swapOutChan = swapOut

	}()

	wg.Wait()

	liquidityX, liquidityY, swapOut, err := liquidityXChan, liquidityYChan, swapOutChan, errorChan
	if err != nil {
		return nil, err
	}

	data := types.JoeSwapOutAndLiquidity{
		LiquidityX: *liquidityX,
		LiquidityY: *liquidityY,
		SwapOut:    *swapOut,
	}

	return &data, nil

}

func JoeArrangeBatch(data types.JoeBodyData) []types.JoeSerializedData {
	serializedData := make([]types.JoeSerializedData, len(data.Data))
	for i := 0; i < len(data.Data); i++ {
		req := data.Data[i]
		if !common.IsHexAddress(req.BaseAsset) {
			serializedData[i] = types.JoeSerializedData{Error: true, Data: &req, ErrorMessage: "Base asset param is invalid, must be a contract address"}
		} else if !common.IsHexAddress(req.QuoteAsset) {
			serializedData[i] = types.JoeSerializedData{Error: true, Data: &req, ErrorMessage: "Quote asset param is invalid, must be a contract address"}
		} else {
			serializedData[i] = types.JoeSerializedData{Error: false, Data: &req, ErrorMessage: ""}
		}
	}
	return serializedData
}

func GetJoeTotalLiquidity(tokenAQuote types.JoeAmountOuts, tokenBQuote types.JoeAmountOuts, reserveA decimal.Decimal, reserveB decimal.Decimal, chain int) error {
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
		if tokenAQuote.TokenId == 1 || tokenAQuote.TokenId == 2 || tokenAQuote.TokenId == 0 {
			tokenAwithReserve := tokenAQuote.AmountOut.Mul(reserveA)

			if tokenAwithReserve.LessThan(lowLiquidityTresholdStable) {
				if tokenBQuote.TokenId == 2 || tokenBQuote.TokenId == 1 {
					tokenBwithReserve := tokenBQuote.AmountOut.Mul(reserveB)

					if tokenBwithReserve.LessThan(lowLiquidityTresholdStable) {
						totalL := tokenAQuote.AmountOut.Add(tokenBQuote.AmountOut)

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
				if tokenBQuote.TokenId == 2 || tokenBQuote.TokenId == 1 {
					tokenBwithReserve := tokenBQuote.AmountOut.Mul(reserveB)
					if tokenBwithReserve.LessThan(lowLiquidityTresholdStable) {
						return errors.New("Low liquidity detected")

					}

				} else {
					tokenBwithReserve := tokenBQuote.AmountOut.Mul(reserveB)
					if tokenBwithReserve.LessThan(lowLiquidityTresholdNative) {
						totalL := tokenAQuote.AmountOut.Add(tokenBQuote.AmountOut)

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
