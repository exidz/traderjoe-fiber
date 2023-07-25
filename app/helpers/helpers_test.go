package helpers

import (
	"testing"

	"github.com/exidz/traderjoe-fiber/app/types"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestArrangeBatch(t *testing.T) {
	input := []types.LbPayload{
		{BaseAsset: "0x1111111111111111111111111111111111111111",
			QuoteAsset: "0x2222222222222222222222222222222222222222",
			Binstep:    10,
		},
		{BaseAsset: "0x1111111111111111111111111111111111111112",
			QuoteAsset: "0x2222222222222222222222222222222222222222",
			Binstep:    10,
		},
		{BaseAsset: "0x1111111111111111111111111111111111111131",
			QuoteAsset: "0x2222222222222222222222222222222222222422",
			Binstep:    10,
		},
	}

	dataInput := types.BodyData{Data: input}
	// Test case with valid input
	data := types.BodyData{Data: dataInput.Data}

	serializedData := ArrangeBatch(data)
	assert.Len(t, serializedData, len(data.Data))

	for _, item := range serializedData {
		assert.False(t, item.Error)
		assert.NotNil(t, item.Data)
		assert.Empty(t, item.ErrorMessage)
	}
}

func TestJoeArrangeBatch(t *testing.T) {
	input := []types.JoePayload{
		{BaseAsset: "0x1111111111111111111111111111111111111111",
			QuoteAsset: "0x2222222222222222222222222222222222222222",
		},
		{BaseAsset: "0x1111111111111111111111111111111111111112",
			QuoteAsset: "0x2222222222222222222222222222222222222222",
		},
		{BaseAsset: "0x1111111111111111111111111111111111111131",
			QuoteAsset: "0x2222222222222222222222222222222222222422",
		},
	}

	dataInput := types.JoeBodyData{Data: input}
	// Test case with valid input
	data := types.JoeBodyData{Data: dataInput.Data}

	serializedData := JoeArrangeBatch(data)
	assert.Len(t, serializedData, len(data.Data))

	for _, item := range serializedData {
		assert.False(t, item.Error)
		assert.NotNil(t, item.Data)
		assert.Empty(t, item.ErrorMessage)
	}
}

func TestGetTotalLiquidity(t *testing.T) {
	// Test case where both tokenAQuote and tokenBQuote have amount out equal to 0.
	tokenAQuote := types.Quote{
		AmountOut: decimal.NewFromInt(0),
		Token:     1,
	}
	tokenBQuote := types.Quote{
		AmountOut: decimal.NewFromInt(0),
		Token:     2,
	}
	reserveA := decimal.NewFromInt(1)
	reserveB := decimal.NewFromInt(2)
	chain := 0

	err := GetTotalLiquidity(tokenAQuote, tokenBQuote, reserveA, reserveB, chain)
	assert.Error(t, err)
	assert.Equal(t, "Low liquidity detected", err.Error())

	// Test case where tokenAQuote has reserve amount out less than lowLiquidityTresholdNative.
	tokenAQuote.AmountOut = decimal.NewFromInt(1)
	err = GetTotalLiquidity(tokenAQuote, tokenBQuote, reserveA, reserveB, chain)
	assert.Error(t, err)
	assert.Equal(t, "Low liquidity detected", err.Error())

	// Test case where tokenBQuote has reserve amount out less than lowLiquidityTresholdStable.
	tokenAQuote.AmountOut = decimal.NewFromInt(10)
	tokenBQuote.AmountOut = decimal.NewFromInt(1)
	err = GetTotalLiquidity(tokenAQuote, tokenBQuote, reserveA, reserveB, chain)
	assert.NoError(t, err)

	// Test case where both tokenAQuote and tokenBQuote have enough liquidity.
	tokenBQuote.AmountOut = decimal.NewFromInt(100)
	err = GetTotalLiquidity(tokenAQuote, tokenBQuote, reserveA, reserveB, chain)
	assert.NoError(t, err)

	// Test case with different chain.
	chain = 2
	err = GetTotalLiquidity(tokenAQuote, tokenBQuote, reserveA, reserveB, chain)
	assert.NoError(t, err)
}
