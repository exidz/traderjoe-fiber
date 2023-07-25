package types

import (
	"github.com/ethereum/go-ethereum/common"
	joepair "github.com/exidz/traderjoe-fiber/app/services/JoePair"
	lbpair "github.com/exidz/traderjoe-fiber/app/services/LBPair"
	legacyrouter "github.com/exidz/traderjoe-fiber/app/services/LegacyRouter"
	"github.com/shopspring/decimal"
)

type LbPayload struct {
	BaseAsset  string `json:"baseAsset"`
	QuoteAsset string `json:"quoteAsset"`
	Binstep    int    `json:"binStep"`
}

type BodyData struct {
	Data []LbPayload `json:"data"`
}

type SerializedData struct {
	Error        bool
	Data         *LbPayload
	ErrorMessage string
}
type JoeSerializedData struct {
	Error        bool
	Data         *JoePayload
	ErrorMessage string
}

type BatchResponse struct {
	Success bool        `json:"success" default:"true"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type BatchErrorResponse struct {
	Error  string      `json:"error"`
	Params BatchParams `json:"params"`
}

type BatchSuccessResponse struct {
	BaseToQuote decimal.Decimal `json:"baseToQuote"`
	QuoteToBase decimal.Decimal `json:"quoteToBase"`
	PairAddress common.Address  `json:"pairContractAddress"`
	QuoteAsset  string          `json:"quoteAsset"`
	BaseAsset   string          `json:"baseAsset"`
	BinStep     int             `json:"binStep"`
	Params      BatchParams     `json:"params"`
}

type BatchParams struct {
	QuoteAsset string `json:"quoteAsset"`
	BaseAsset  string `json:"baseAsset"`
	BinStep    int    `json:"binStep"`
}

type BatchPriceResponse struct {
	BaseToQuote decimal.Decimal `json:"baseToQuote"`
	QuoteToBase decimal.Decimal `json:"quoteToBase"`
	PairAddress common.Address  `json:"pairContractAddress"`
	QuoteAsset  string          `json:"quoteAsset"`
	BaseAsset   string          `json:"baseAsset"`
}

type PriceResponse struct {
	BaseToQuote decimal.Decimal `json:"baseToQuote"`
	QuoteToBase decimal.Decimal `json:"quoteToBase"`
	QuoteAsset  string          `json:"quoteAsset"`
	BaseAsset   string          `json:"baseAsset"`
	BinStep     int             `json:"binStep"`
	PairAddress common.Address  `json:"pairContractAddress"`
}

type ErrorResponse struct {
	Error   string      `json:"error"`
	Success bool        `json:"success" default:"false"`
	Data    interface{} `json:"data"`
}

type Quote struct {
	AmountOut decimal.Decimal
	Fee       decimal.Decimal
	TotalAmt  decimal.Decimal
	Token     int
}

type QuoteBulk struct {
	BaseToQuote Quote
	QuoteToBase Quote
}

type TokenDecimals struct {
	BaseAssetDecimal  uint8
	QuoteAssetDecimal uint8
}

type PairSwapOut struct {
	BaseToQuote lbpair.SwapOut
	QuoteToBase lbpair.SwapOut
}

type LegacyPairSwapOut struct {
	BaseToQuote legacyrouter.SwapOut
	QuoteToBase legacyrouter.SwapOut
}

type JoePriceResponse struct {
	BaseToQuote *decimal.Decimal `json:"baseToQuote"`
	QuoteToBase *decimal.Decimal `json:"quoteToBase"`
	QuoteAsset  string           `json:"quoteAsset"`
	BaseAsset   string           `json:"baseAsset"`
	PairAddress common.Address   `json:"pairContractAddress"`
}

type JoePairSwapOut struct {
	XtoY *decimal.Decimal
	YToX *decimal.Decimal
}

type JoeAmountOuts struct {
	AmountOut         decimal.Decimal
	TokenQuoteDecimal uint8
	TokenId           int
}

type PairInfo struct {
	Tokens  joepair.PairTokens
	Reserve joepair.JoePairReserve
}

type JoeSwapOutAndLiquidity struct {
	LiquidityX JoeAmountOuts
	LiquidityY JoeAmountOuts
	SwapOut    JoePairSwapOut
}

type JoePayload struct {
	BaseAsset  string `json:"baseAsset"`
	QuoteAsset string `json:"quoteAsset"`
}

type JoeBodyData struct {
	Data []JoePayload `json:"data"`
}

type JoeBatchParams struct {
	QuoteAsset string `json:"quoteAsset"`
	BaseAsset  string `json:"baseAsset"`
}

type JoeBatchErrorResponse struct {
	Error  string         `json:"error"`
	Params JoeBatchParams `json:"params"`
}

type JoeBatchSuccessResponse struct {
	BaseToQuote decimal.Decimal `json:"baseToQuote"`
	QuoteToBase decimal.Decimal `json:"quoteToBase"`
	QuoteAsset  string          `json:"quoteAsset"`
	BaseAsset   string          `json:"baseAsset"`
	PairAddress common.Address  `json:"pairContractAddress"`
	Params      JoeBatchParams  `json:"params"`
}

type Response struct {
	Success bool        `json:"success" default:"true"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
