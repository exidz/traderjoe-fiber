package legacyrouter

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/exidz/traderjoe-fiber/app/constants"
	"github.com/exidz/traderjoe-fiber/contracts"
	"github.com/exidz/traderjoe-fiber/utils"
)

type LegacyRouter struct {
	Instance contracts.Legacyrouter
	Chain    int
}

type SwapOut struct {
	AmountOut *big.Int
	FeesIn    *big.Int
}

func InitLegacyRouter(client *ethclient.Client, chain constants.Chain) (*LegacyRouter, error) {
	var contractAddr common.Address

	if chain == 0 {
		contractAddr = common.HexToAddress(constants.LBRouterV2)
	} else if chain == 1 {
		contractAddr = common.HexToAddress(constants.ARB_ROUTER_V2)
	} else if chain == 2 {
		contractAddr = common.HexToAddress(constants.BSC_ROUTER_V2)
	} else {
		return nil, errors.New("Please provide a valid chain (AVAX, ARB, BSC)")
	}
	instance, err := contracts.NewLegacyrouter(contractAddr, client)

	if err != nil {
		return nil, errors.New("Contract error: cannot initialized legacy router")
	}

	return &LegacyRouter{
		Instance: *instance,
		Chain:    int(chain),
	}, nil
}

func (r LegacyRouter) LegacySwapOut(pairAddr common.Address, decimal uint8, swapforY bool) (*SwapOut, error) {
	amountIn := utils.ToWei("1", int(decimal))

	swapOut, err := r.Instance.GetSwapOut(&bind.CallOpts{}, pairAddr, amountIn, swapforY)
	if err != nil {
		return nil, errors.New("Contract error:: cannot fetch legacy swapout")
	}

	return &SwapOut{
		AmountOut: swapOut.AmountOut,
		FeesIn:    swapOut.FeesIn,
	}, nil

}
