package joerouter

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/exidz/traderjoe-fiber/app/constants"
	"github.com/exidz/traderjoe-fiber/contracts"
	"github.com/exidz/traderjoe-fiber/utils"
	"github.com/shopspring/decimal"
)

type JoeRouter struct {
	Instance *contracts.Joerouter
	Chain    int
}

type Quote struct {
	AmountOut    decimal.Decimal
	TokenDecimal int
	TokenId      int
}

func JoeRouterInit(client *ethclient.Client, chain constants.Chain) (*JoeRouter, error) {
	var contractAddr common.Address

	if chain == 0 {
		contractAddr = common.HexToAddress(constants.JoeRouter)
	} else if chain == 1 {
		contractAddr = common.HexToAddress(constants.ARB_ROUTER_V1)
	} else if chain == 2 {
		contractAddr = common.HexToAddress(constants.BSC_ROUTER_V1)
	} else {
		return nil, errors.New("Please provide a valid chain (AVAX, ARB, BSC)")
	}

	router, err := contracts.NewJoerouter(contractAddr, client)

	if err != nil {
		return nil, errors.New("RPC timeout: cannot initialized router")
	}

	return &JoeRouter{
		Instance: router,
		Chain:    int(chain),
	}, nil

}

func (r JoeRouter) GetJoeRouterAmountOut(tokenDecimal uint8, reserveIn big.Int, reserveOut big.Int) (*big.Int, error) {

	amountIn := utils.ToWei(1.0, int(tokenDecimal))

	amountOut, err := r.Instance.GetAmountOut(&bind.CallOpts{}, amountIn, &reserveIn, &reserveOut)

	if err != nil {
		return nil, errors.New("RPC timeout: cannot get amount out")
	}

	return amountOut, nil

}

func (r JoeRouter) GetQuote(tokenX common.Address, amountX *big.Int) (*Quote, error) {
	var xpath [4]common.Address
	xpath[0] = tokenX

	if r.Chain == 0 {
		xpath[1] = common.HexToAddress(constants.USDC)
		xpath[2] = common.HexToAddress(constants.USDT)
		xpath[3] = common.HexToAddress(constants.WAVAX)

	} else if r.Chain == 1 {
		xpath[1] = common.HexToAddress(constants.ARB_USDC)
		xpath[2] = common.HexToAddress(constants.ARB_USDT)
		xpath[3] = common.HexToAddress(constants.ARB_WETH)
	} else if r.Chain == 2 {
		xpath[1] = common.HexToAddress(constants.BSC_USDC)
		xpath[2] = common.HexToAddress(constants.BSC_USDT)
		xpath[3] = common.HexToAddress(constants.BSC_WBNB)
	} else {
		return nil, errors.New("Please provide a valid chain (AVAX, ARB, BSC)")
	}

	path := utils.RemoveFirstIfIdentical(xpath[:])

	quote, err := r.Instance.GetAmountsOut(&bind.CallOpts{}, amountX, path[:])
	if err != nil {
		var pathX [2]common.Address
		pathX[0] = tokenX
		if r.Chain == 0 {
			pathX[1] = common.HexToAddress(constants.WAVAX)
		} else if r.Chain == 1 {

			pathX[1] = common.HexToAddress(constants.ARB_WETH)
		} else if r.Chain == 2 {

			pathX[1] = common.HexToAddress(constants.BSC_WBNB)
		} else {
			return nil, errors.New("Please provide a valid chain (AVAX, ARB, BSC)")
		}

		avaxQuote, err := r.Instance.GetAmountsOut(&bind.CallOpts{}, amountX, pathX[:])
		if err != nil {
			return nil, errors.New("Contract timeout: cannot fetch quote data")
		}
		amountX := utils.ToDecimal(*avaxQuote[1], int(18))
		return &Quote{
			AmountOut:    amountX,
			TokenDecimal: 18,
			TokenId:      3,
		}, nil

	}
	var quoted Quote

loop:
	for i := 1; i < len(quote); i++ {
		v := quote[i]
		if r.Chain == 0 || r.Chain == 1 {
			if v != big.NewInt(0) {
				decimal := 0
				if i == 1 || i == 2 {
					decimal = 6
				} else {
					decimal = 18
				}
				amountX := utils.ToDecimal(v, int(decimal))
				quoted = Quote{
					AmountOut:    amountX,
					TokenDecimal: decimal,
					TokenId:      i,
				}
				break loop
			}
		} else {
			if v != big.NewInt(0) {
				decimal := 18
				amountX := utils.ToDecimal(v, int(decimal))

				quoted = Quote{
					AmountOut:    amountX,
					TokenDecimal: 18,
					TokenId:      i,
				}
				break loop
			}
		}

	}

	return &quoted, nil

}
