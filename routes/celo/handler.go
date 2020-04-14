package celo

import (
	"encoding/json"
	"log"
	"math/big"
	"net/http"

	"github.com/conclave-dev/celoist-backend/kvstore"
	"github.com/conclave-dev/go-celo/client"
	"github.com/conclave-dev/go-celo/core/celo"
	"github.com/conclave-dev/go-celo/core/celo/governance/election"
	"github.com/conclave-dev/go-celo/core/celo/governance/validators"
	"github.com/conclave-dev/go-celo/types"
	"github.com/conclave-dev/go-celo/util"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

const (
	celoRPCAPI = "https://geth.celoist.com"
)

func init() {
	util.SetupClients()
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

func getTotalVotesForEligibleValidatorGroups(opts *bind.CallOpts) (struct {
	Groups []common.Address
	Values []*big.Int
}, error) {
	contract := getElectionContract()
	return contract.GetTotalVotesForEligibleValidatorGroups(opts)
}

func getElection(election string) ([]byte, error) {
	e, err := kvstore.GetElection(election)
	if err != nil {
		return nil, err
	}

	var electionData kvstore.Election
	err = json.Unmarshal([]byte(e), &electionData)
	if err != nil {
		return nil, err
	}

	return json.Marshal(types.JSONResponse{
		Data: electionData,
	})
}

func setElection(opts *bind.CallOpts, election string) (kvstore.Election, error) {
	groups, err := getTotalVotesForEligibleValidatorGroups(opts)
	if err != nil {
		return kvstore.Election{}, err
	}

	data := kvstore.Election{
		GroupAddresses: groups.Groups,
		GroupVotes:     groups.Values,
	}

	_, err = kvstore.SetElection(election, data)
	if err != nil {
		return kvstore.Election{}, err
	}

	return data, nil
}

func handleElection(w http.ResponseWriter, r *http.Request) {
	opts := getCallOpts(w, r)
	epochNumber, err := getEpochNumber(opts)
	if err != nil {
		util.RespondWithError(err, r, w)
	}

	ens := epochNumber.String()

	if kvstore.DoesElectionExist(ens) {
		election, err := getElection(ens)
		if err != nil {
			util.RespondWithError(err, r, w)
		}

		util.RespondWithData(election, w)
		return
	}

	election, err := setElection(opts, ens)
	if err != nil {
		util.RespondWithError(err, r, w)
	}

	res, err := json.Marshal(types.JSONResponse{
		Data: struct {
			Election kvstore.Election `json:"election"`
		}{
			Election: election,
		},
	})
	if err != nil {
		util.RespondWithError(err, r, w)
	}

	util.RespondWithData(res, w)
	return
}
