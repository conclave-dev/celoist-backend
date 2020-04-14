package celo

import (
	"log"

	"github.com/conclave-dev/go-celo/client"
	"github.com/conclave-dev/go-celo/core/celo"
	"github.com/conclave-dev/go-celo/core/celo/governance/election"
	"github.com/conclave-dev/go-celo/core/celo/governance/validators"
)

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
