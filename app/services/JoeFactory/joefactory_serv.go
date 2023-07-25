package joefactory

import (
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/exidz/traderjoe-fiber/app/constants"
	"github.com/exidz/traderjoe-fiber/contracts"
)

type JoeFactory struct {
	Instance contracts.Joefactory
	Chain    int
}

func JoeFactoryInit(client *ethclient.Client, chain constants.Chain) (*JoeFactory, error) {
	var contractAddr common.Address

	if chain == 0 {
		contractAddr = common.HexToAddress(constants.JoeFactory)
	} else if chain == 1 {
		contractAddr = common.HexToAddress(constants.ARB_FACTORY_V1)
	} else if chain == 2 {
		contractAddr = common.HexToAddress(constants.BSC_FACTORY_V1)
	} else {
		return nil, errors.New("Please provide a valid chain (AVAX, ARB, BSC)")
	}
	factory, err := contracts.NewJoefactory(contractAddr, client)

	if err != nil {
		return nil, errors.New("Contract error: cannot initialized joe factory")
	}

	return &JoeFactory{
		Instance: *factory,
		Chain:    int(chain),
	}, nil

}

func (f JoeFactory) GetJoePair(tokenX common.Address, tokenY common.Address) (*common.Address, error) {

	pair, err := f.Instance.GetPair(&bind.CallOpts{}, tokenX, tokenY)

	if err != nil {
		return nil, errors.New("Contract error: cannot get joe pair")
	}

	return &pair, nil

}
