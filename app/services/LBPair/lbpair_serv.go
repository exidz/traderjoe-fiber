package lbpair

import (
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/exidz/traderjoe-fiber/contracts"
	"github.com/exidz/traderjoe-fiber/utils"
)

type PairV2 struct {
	instance *contracts.LBPair
	pair     common.Address
}

type PairTokens struct {
	BaseToken  common.Address
	QuoteToken common.Address
}

type SwapOut struct {
	AmountInLeft *big.Int
	AmountOut    *big.Int
	Fee          *big.Int
}

func InitPair(client *ethclient.Client, pair common.Address) (*PairV2, error) {

	instance, err := contracts.NewLBPair(pair, client)

	if err != nil {
		return nil, errors.New("Contract error: pair contract cannot initialized")
	}

	return &PairV2{
		instance: instance,
		pair:     pair,
	}, nil

}

func (p PairV2) GetPairReserve() (*struct {
	ReserveX *big.Int
	ReserveY *big.Int
}, error) {

	reserves, err := p.instance.GetReserves(&bind.CallOpts{})

	if err != nil {
		return nil, errors.New("Contract error: cannot fetch pair reserve")
	}

	return &reserves, nil

}

func (p PairV2) GetTokenX() (*common.Address, error) {

	x, err := p.instance.GetTokenX(&bind.CallOpts{})
	if err != nil {
		return nil, errors.New("Contract error: cannot fetch base asset token info")
	}

	return &x, nil
}

func (p PairV2) GetTokenY() (*common.Address, error) {

	y, err := p.instance.GetTokenY(&bind.CallOpts{})
	if err != nil {
		return nil, errors.New("RPC timeout: cannot fetch quote asset token info")
	}

	return &y, nil
}

func (p PairV2) GetPairTokens() (*PairTokens, error) {
	tokenYCh := make(chan common.Address)
	tokenXCh := make(chan common.Address)
	errCh := make(chan error)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		y, err := p.instance.GetTokenY(&bind.CallOpts{})
		if err != nil {
			errCh <- fmt.Errorf("RPC timeout: cannot fetch quote asset token info: %w", err)
			return
		}

		tokenYCh <- y
	}()

	go func() {
		defer wg.Done()
		x, err := p.instance.GetTokenX(&bind.CallOpts{})
		if err != nil {
			errCh <- fmt.Errorf("RPC timeout: cannot fetch base asset token info: %w", err)
			return
		}

		tokenXCh <- x
	}()

	go func() {
		wg.Wait()
		close(tokenXCh)
		close(tokenYCh)
		close(errCh)
	}()

	tokenX, tokenY, err := <-tokenXCh, <-tokenYCh, <-errCh
	if err != nil {
		return nil, err
	}

	tokens := &PairTokens{
		BaseToken:  tokenX,
		QuoteToken: tokenY,
	}
	return tokens, nil
}

func (p PairV2) GetSwapOut(decimal uint8, swapToy bool) (*SwapOut, error) {
	amountX := utils.ToWei(1.0, int(decimal))
	x, err := p.instance.GetSwapOut(&bind.CallOpts{}, amountX, swapToy)
	if err != nil {
		return nil, errors.New("RPC timeout: cannot fetch swap out data ")
	}

	swapout := SwapOut{
		AmountInLeft: x.AmountInLeft,
		AmountOut:    x.AmountOut,
		Fee:          x.Fee,
	}

	return &swapout, nil
}
