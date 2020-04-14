package celo

import (
	"encoding/json"
	"math/big"
	"sync"

	"github.com/conclave-dev/celoist-backend/kvstore"
	"github.com/conclave-dev/go-celo/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

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
	g, err := getTotalVotesForEligibleValidatorGroups(opts)
	if err != nil {
		return kvstore.Election{}, err
	}

	groups, err := getGroups(opts, g.Groups)
	if err != nil {
		return kvstore.Election{}, err
	}

	data := kvstore.Election{
		GroupAddresses: g.Groups,
		GroupVotes:     g.Values,
		Groups:         groups,
	}

	_, err = kvstore.SetElection(election, data)
	if err != nil {
		return kvstore.Election{}, err
	}

	return data, nil
}

func getGroups(opts *bind.CallOpts, groupAddresses kvstore.GroupAddresses) (kvstore.Groups, error) {
	contract := getValidatorsContract()
	groups := make(kvstore.Groups)

	fetchGroup := func(groupAddress common.Address, wg *sync.WaitGroup, mu *sync.Mutex) (err error) {
		mu.Lock()

		defer wg.Done()

		memberAddresses, commission, nextCommission, nextCommissionBlock, _, slashMultiplier, lastSlash, err := contract.GetValidatorGroup(opts, groupAddress)
		if err != nil {
			return
		}

		members, err := getMembers(opts, memberAddresses)

		groups[groupAddress] = kvstore.Group{
			Address:             groupAddress,
			Commission:          commission,
			NextCommission:      nextCommission,
			NextCommissionBlock: nextCommissionBlock,
			MemberAddresses:     memberAddresses,
			Members:             members,
			LastSlash:           lastSlash,
			SlashMultiplier:     slashMultiplier,
		}

		mu.Unlock()
		return
	}

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, groupAddress := range groupAddresses {
		wg.Add(1)
		go fetchGroup(groupAddress, &wg, &mu)
	}

	wg.Wait()

	return groups, nil
}

func getMembers(opts *bind.CallOpts, memberAddresses kvstore.MemberAddresses) (kvstore.Members, error) {
	contract := getValidatorsContract()
	members := make(kvstore.Members)

	fetchGroup := func(memberAddress common.Address, wg *sync.WaitGroup, mu *sync.Mutex) (err error) {
		mu.Lock()

		defer wg.Done()

		validator, err := contract.GetValidator(opts, memberAddress)
		if err != nil {
			return
		}

		members[memberAddress] = kvstore.Member{
			Address:        memberAddress,
			ECDSAPublicKey: validator.EcdsaPublicKey,
			BLSPublicKey:   validator.BlsPublicKey,
			Score:          validator.Score,
			Signer:         validator.Signer,
		}

		mu.Unlock()
		return
	}

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, memberAddress := range memberAddresses {
		wg.Add(1)
		go fetchGroup(memberAddress, &wg, &mu)
	}

	wg.Wait()

	return members, nil
}
