package legacyfactory

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/exidz/traderjoe-fiber/app/constants"
	"github.com/exidz/traderjoe-fiber/contracts"
)

type LegacyFactory struct {
	Instance contracts.Legacyfactory
	Chain    int
}

func LegacyFactoryInit(client *ethclient.Client, chain constants.Chain) (*LegacyFactory, error) {

	var contractAddr common.Address

	if chain == 0 {
		contractAddr = common.HexToAddress(constants.LBFactoryV2)
	} else if chain == 1 {
		contractAddr = common.HexToAddress(constants.ARB_FACTORY_V2)
	} else if chain == 2 {
		contractAddr = common.HexToAddress(constants.BSC_FACTORY_V2)
	} else {
		return nil, errors.New("Please provide a valid chain (AVAX, ARB, BSC)")
	}
	instance, err := contracts.NewLegacyfactory(contractAddr, client)
	if err != nil {
		return nil, errors.New("Contract error: cannot initialized legacy factory")
	}

	return &LegacyFactory{
		Instance: *instance,
	}, nil
}

func (f LegacyFactory) GetLegacyPairInfo(xAddr common.Address, yAddr common.Address, binstep int) (*contracts.ILBFactoryLBPairInformation, error) {
	pairInfo, err := f.Instance.GetLBPairInformation(&bind.CallOpts{}, xAddr, yAddr, big.NewInt(int64(binstep)))

	if err != nil {
		return nil, errors.New("Contract error: cannot fetch corresponding LB pair")
	}

	return &pairInfo, nil
}
