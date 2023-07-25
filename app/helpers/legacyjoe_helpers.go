package helpers

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	legacyquoter "github.com/exidz/traderjoe-fiber/app/services/LegacyQuoter"
	legacyrouter "github.com/exidz/traderjoe-fiber/app/services/LegacyRouter"
	"github.com/exidz/traderjoe-fiber/app/types"
	"github.com/exidz/traderjoe-fiber/utils"
)

func getQuoteLegacyUnderlying(tokenX common.Address, decimalX uint8, quoteInstance *legacyquoter.LegacyQuoter) (*types.Quote, error) {
	amountX := utils.ToWei(1.0, int(decimalX))
	var quote types.Quote

	q1, err := quoteInstance.LegacyGetQuote(tokenX, amountX)
	if err != nil {

		return nil, err
	}

	for i := 1; i < len(q1.VirtualAmountsWithoutSlippage)-1; i++ {
		v := q1.VirtualAmountsWithoutSlippage[i]
		if v != big.NewInt(0) {
			decimal := 18
			switch i {
			case 1:
				if v != nil && v.Sign() > 0 {
					if quoteInstance.Chain == 3 {
						decimal = 18
						break
					}
					decimal = 6
					break
				}

			case 2:
				if v != nil && v.Sign() > 0 {
					if quoteInstance.Chain == 3 {
						decimal = 18
						break
					}
					decimal = 6
					break
				}

			case 3:
				if v != nil && v.Sign() > 0 {
					decimal = 18
					break
				}

			}

			fee := q1.Fees[i-1]

			feeInDecimal := utils.ToDecimal(fee, int(18))
			amountX := utils.ToDecimal(v, int(decimal))

			amountOutWithoutFee := feeInDecimal.Add(amountX)

			quote = types.Quote{
				AmountOut: amountX,
				Fee:       feeInDecimal,
				TotalAmt:  amountOutWithoutFee,
				Token:     i,
			}

			break
		}

	}

	return &quote, nil
}

func GetLegacyPairQuote(baseAsset common.Address, quoteAsset common.Address, baseDecimal uint8, quoteDecimal uint8, quoter legacyquoter.LegacyQuoter) (*types.QuoteBulk, error) {
	var xQoute *types.Quote
	var yQuote *types.Quote
	var error error

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		base, err := getQuoteLegacyUnderlying(baseAsset, baseDecimal, &quoter)
		if err != nil {

			error = err

			return
		}
		xQoute = base

	}()

	go func() {
		defer wg.Done()
		quote, err := getQuoteLegacyUnderlying(quoteAsset, quoteDecimal, &quoter)
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

func LegacyPairSwapOut(router legacyrouter.LegacyRouter, decimalX uint8, decimalY uint8, pairAddr common.Address) (*types.LegacyPairSwapOut, error) {
	var xSwapChan legacyrouter.SwapOut
	var ySwapChan legacyrouter.SwapOut
	var erroChan error

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		baseToQuote, err := router.LegacySwapOut(pairAddr, decimalX, true)
		if err != nil {
			erroChan = err
			return
		}
		xSwapChan = *baseToQuote

	}()

	go func() {
		defer wg.Done()
		quoteToBase, err := router.LegacySwapOut(pairAddr, decimalY, false)
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

	swapOut := types.LegacyPairSwapOut{
		BaseToQuote: xSwapChan,
		QuoteToBase: ySwapChan,
	}

	return &swapOut, nil

}
