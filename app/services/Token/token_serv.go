package token

import (
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/exidz/traderjoe-fiber/contracts"
)

type Token struct {
	instance *contracts.Token
	address  common.Address
}

func InitToken(client *ethclient.Client, address common.Address) (*Token, error) {
	instance, err := contracts.NewToken(address, client)

	if err != nil {
		return nil, errors.New("Contract error: token contract cannot initialized")
	}

	return &Token{
		instance: instance,
		address:  address,
	}, nil
}

func (t Token) GetDecimal() (*uint8, error) {
	decimal, err := t.instance.Decimals(&bind.CallOpts{})
	if err != nil {
		return nil, errors.New("Contract error: cannot fetch token decimal")
	}

	return &decimal, nil
}
