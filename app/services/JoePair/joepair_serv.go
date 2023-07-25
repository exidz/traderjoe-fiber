package joepair

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

type JoePair struct {
	instance contracts.Joepair
}
type JoePairReserve struct {
	ReserveX           *big.Int
	ReserveY           *big.Int
	BlockTimestampLast uint32
}

type PairTokens struct {
	TokenX common.Address
	TokenY common.Address
}

func InitJoePair(pair common.Address, client *ethclient.Client) (*JoePair, error) {

	pairinit, err := contracts.NewJoepair(pair, client)

	if err != nil {
		return nil, errors.New("RPC timeout: cannot initialized pair")
	}

	return &JoePair{
		instance: *pairinit,
	}, nil

}

func (j JoePair) GetJoePairReserve() (*JoePairReserve, error) {

	reserve, err := j.instance.GetReserves(&bind.CallOpts{})

	if err != nil {
		return nil, err
	}

	return &JoePairReserve{
		ReserveX:           reserve.Reserve0,
		ReserveY:           reserve.Reserve1,
		BlockTimestampLast: reserve.BlockTimestampLast,
	}, nil

}

func (j JoePair) GetJoePairTokens() (*PairTokens, error) {
	tokenYCh := make(chan common.Address)
	tokenXCh := make(chan common.Address)
	errCh := make(chan error)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		y, err := j.instance.Token1(&bind.CallOpts{})
		if err != nil {
			errCh <- fmt.Errorf("RPC timeout: cannot fetch quote asset token info: %w", err)
			return
		}

		tokenYCh <- y
	}()

	go func() {
		defer wg.Done()
		x, err := j.instance.Token0(&bind.CallOpts{})
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
		TokenX: tokenX,
		TokenY: tokenY,
	}
	return tokens, nil
}
