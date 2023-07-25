package lbquoter

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/exidz/traderjoe-fiber/app/constants"
	"github.com/exidz/traderjoe-fiber/contracts"
)

type Quoter struct {
	Instance *contracts.LBQuoter
	Chain    int
}

type QuoteAbs struct {
	Quote *contracts.LBQuoterQuote
	Chain int
}

func QuoterInit(client *ethclient.Client, chain constants.Chain) (*Quoter, error) {

	var contractAddr common.Address

	if chain == 0 {
		contractAddr = common.HexToAddress(constants.LBQuoterV2_1)
	} else if chain == 1 {
		contractAddr = common.HexToAddress(constants.ARB_QUOTER_V21)
	} else if chain == 2 {
		contractAddr = common.HexToAddress(constants.BSC_QUOTER_V21)
	} else {
		return nil, errors.New("Please provide a valid chain (AVAX, ARB, BSC)")
	}
	instance, err := contracts.NewLBQuoter(contractAddr, client)

	if err != nil {
		return nil, errors.New("Contract error: quoter contract cannot initialized")
	}

	return &Quoter{
		Instance: instance,
		Chain:    int(chain),
	}, nil

}

func (q Quoter) GetQuote(tokenX common.Address, amountX *big.Int) (*QuoteAbs, error) {

	var xpath [4]common.Address
	xpath[0] = tokenX

	if q.Chain == 0 {
		xpath[1] = common.HexToAddress(constants.USDC)
		xpath[2] = common.HexToAddress(constants.USDT)
		xpath[3] = common.HexToAddress(constants.WAVAX)

	} else if q.Chain == 1 {
		xpath[1] = common.HexToAddress(constants.ARB_USDC)
		xpath[2] = common.HexToAddress(constants.ARB_USDT)
		xpath[3] = common.HexToAddress(constants.ARB_WETH)
	} else if q.Chain == 2 {
		xpath[1] = common.HexToAddress(constants.BSC_USDC)
		xpath[2] = common.HexToAddress(constants.BSC_USDT)
		xpath[3] = common.HexToAddress(constants.BSC_WBNB)
	} else {
		return nil, errors.New("Please provide a valid chain (AVAX, ARB, BSC)")
	}

	quote, err := q.Instance.FindBestPathFromAmountIn(&bind.CallOpts{}, xpath[:], amountX)
	if err != nil {
		return nil, errors.New("Contract error: cannot fetch quote data")
	}

	return &QuoteAbs{
		Quote: &quote,
		Chain: q.Chain,
	}, nil

}
