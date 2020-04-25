package celo

import (
	"log"
	"math/big"

	"github.com/conclave-dev/go-celo/client"
	"github.com/conclave-dev/go-celo/core/celo"
	"github.com/conclave-dev/go-celo/core/celo/common/accounts"
	"github.com/conclave-dev/go-celo/core/celo/governance/election"
	"github.com/conclave-dev/go-celo/core/celo/governance/validators"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func getAccountsContract() *accounts.Accounts {
	// Use the predefined address for the RegistryContract
	contractAddress := celo.GetContractAddress(celo.Accounts, client.EthClient)
	contract, err := accounts.NewAccounts(contractAddress, client.EthClient)
	if err != nil {
		log.Fatal(err)
	}

	return contract
}

func getElectionContract() *election.Election {
	// Use the predefined address for the RegistryContract
	contractAddress := celo.GetContractAddress(celo.Election, client.EthClient)
	contract, err := election.NewElection(contractAddress, client.EthClient)
	if err != nil {
		log.Fatal(err)
	}

	return contract
}

func getValidatorsContract() *validators.Validators {
	// Use the predefined address for the RegistryContract
	contractAddress := celo.GetContractAddress(celo.Validators, client.EthClient)
	contract, err := validators.NewValidators(contractAddress, client.EthClient)
	if err != nil {
		log.Fatal(err)
	}

	return contract
}

func getEpochNumber(opts *bind.CallOpts) (epochNumber *big.Int, err error) {
	contract := getValidatorsContract()
	n, err := contract.GetEpochNumber(opts)
	if err != nil {
		return
	}

	return n.Sub(n, big.NewInt(1)), err
}
