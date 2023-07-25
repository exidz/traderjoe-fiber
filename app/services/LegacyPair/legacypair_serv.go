package legacypair

import (
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/exidz/traderjoe-fiber/contracts"
)

type LegacyPair struct {
	instance contracts.Legacypair
	pair     common.Address
}

type PairTokens struct {
	BaseToken  common.Address
	QuoteToken common.Address
}

func LegacyPairInit(pairAddr common.Address, client *ethclient.Client) (*LegacyPair, error) {
	instance, err := contracts.NewLegacypair(pairAddr, client)

	if err != nil {
		return nil, errors.New("Contract error:: cannot fecth pair")
	}

	return &LegacyPair{
		instance: *instance,
		pair:     pairAddr,
	}, nil
}

func (p LegacyPair) GetLegacyPairTokens() (*PairTokens, error) {
	var tokenYCh common.Address
	var tokenXCh common.Address
	var errCh error

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		y, err := p.instance.TokenY(&bind.CallOpts{})
		if err != nil {
			errCh = fmt.Errorf("Contract error: cannot fetch quote asset token info: %w", err)
			return
		}

		tokenYCh = y
	}()

	go func() {
		defer wg.Done()
		x, err := p.instance.TokenX(&bind.CallOpts{})
		if err != nil {
			errCh = fmt.Errorf("Contract error:: cannot fetch base asset token info: %w", err)
			return
		}

		tokenXCh = x
	}()

	wg.Wait()

	tokenX, tokenY, err := tokenXCh, tokenYCh, errCh
	if err != nil {
		return nil, err
	}

	tokens := &PairTokens{
		BaseToken:  tokenX,
		QuoteToken: tokenY,
	}
	return tokens, nil
}

func (p LegacyPair) GetPairReserve() (*struct {
	ReserveX *big.Int
	ReserveY *big.Int
	ActiveId *big.Int
}, error) {

	reserves, err := p.instance.GetReservesAndId(&bind.CallOpts{})

	if err != nil {
		return nil, errors.New("Contract error:: cannot fetch pair reserve")
	}

	return &reserves, nil

}
