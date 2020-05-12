package celo

import (
	"log"
	"math/big"

	"github.com/conclave-dev/celoist-backend/util"

	"github.com/conclave-dev/go-celo/core/celo"
	"github.com/conclave-dev/go-celo/core/celo/common/accounts"
	"github.com/conclave-dev/go-celo/core/celo/governance/election"
	"github.com/conclave-dev/go-celo/core/celo/governance/validators"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func getAccountsContract(networkID string) *accounts.Accounts {
	// Use the predefined address for the RegistryContract
	rpcClient, err := util.GetNetworkClient(networkID)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := getContractAddress(networkID, celo.Accounts)
	contract, err := accounts.NewAccounts(contractAddress, rpcClient)
	if err != nil {
		log.Fatal(err)
	}

	return contract
}

func getElectionContract(networkID string) *election.Election {
	// Use the predefined address for the RegistryContract
	rpcClient, err := util.GetNetworkClient(networkID)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := getContractAddress(networkID, celo.Election)
	contract, err := election.NewElection(contractAddress, rpcClient)
	if err != nil {
		log.Fatal(err)
	}

	return contract
}

func getValidatorsContract(networkID string) *validators.Validators {
	// Use the predefined address for the RegistryContract
	rpcClient, err := util.GetNetworkClient(networkID)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := getContractAddress(networkID, celo.Validators)
	contract, err := validators.NewValidators(contractAddress, rpcClient)
	if err != nil {
		log.Fatal(err)
	}

	return contract
}

func getEpochNumber(networkID string, opts *bind.CallOpts) (epochNumber *big.Int, err error) {
	contract := getValidatorsContract(networkID)
	n, err := contract.GetEpochNumber(opts)
	if err != nil {
		log.Fatal(err)
		return
	}

	return n.Sub(n, big.NewInt(1)), err
}
