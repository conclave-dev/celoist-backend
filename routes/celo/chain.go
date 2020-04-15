package celo

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func getEpochNumber(opts *bind.CallOpts) (epochNumber *big.Int, err error) {
	contract := getValidatorsContract()
	n, err := contract.GetEpochNumber(opts)
	if err != nil {
		return
	}

	return n.Sub(n, big.NewInt(1)), err
}
