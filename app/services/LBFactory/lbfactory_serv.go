package lbfactory

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/exidz/traderjoe-fiber/app/constants"
	"github.com/exidz/traderjoe-fiber/contracts"
)

type FactoryV2 struct {
	Instance *contracts.LBFactory
	Chain    int
}

func CreateFactory(client *ethclient.Client, chain constants.Chain) (*FactoryV2, error) {
	var contractAddr common.Address

	if chain == 0 {
		contractAddr = common.HexToAddress(constants.LBFactoryV2_1)
	} else if chain == 1 {
		contractAddr = common.HexToAddress(constants.ARB_FACTORY_V21)
	} else if chain == 2 {
		contractAddr = common.HexToAddress(constants.BSC_FACTORY_V21)
	} else {
		return nil, errors.New("Please provide a valid chain (AVAX, ARB, BSC)")
	}

	instance, err := contracts.NewLBFactory(contractAddr, client)

	if err != nil {
		return nil, errors.New("Contract error: factory contract cannot initialized")
	}

	return &FactoryV2{
		Instance: instance,
		Chain:    int(chain),
	}, nil

}

func (f FactoryV2) GetPairInfo(baseAsset common.Address, quoteAsset common.Address, binStep big.Int) (*contracts.ILBFactoryLBPairInformation, error) {
	pair, err := f.Instance.GetLBPairInformation(&bind.CallOpts{}, baseAsset, quoteAsset, &binStep)

	if err != nil {
		return nil, errors.New("Contract error: cannot fetch corresponding LB pair")
	}

	return &pair, nil

}
