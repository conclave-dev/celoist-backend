package celo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/conclave-dev/celoist-backend/util"

	"github.com/conclave-dev/celoist-backend/kvstore"
	"github.com/conclave-dev/go-celo/types"
)

func handleElection(w http.ResponseWriter, r *http.Request) {
	var election []byte
	var err error

	networkID, err := util.ParseNetworkID(r)
	if err != nil {
		util.RespondWithError(err, r, w)
		return
	}

	callOpts, err := getCallOpts(networkID, w, r)
	if err != nil {
		util.RespondWithError(err, r, w)
		return
	}

	epochNumber, err := getEpochNumber(networkID, callOpts)
	if err != nil {
		util.RespondWithError(err, r, w)
		return
	}

	ens := epochNumber.String()

	if kvstore.DoesElectionExist(networkID, ens) {
		election, err = getElection(networkID, ens)
		if err != nil {
			util.RespondWithError(err, r, w)
			return
		}
	} else {
		e, err := setElection(networkID, callOpts, ens)
		if err != nil {
			util.RespondWithError(err, r, w)
			return
		}

		election, err = json.Marshal(types.JSONResponse{
			Data: e,
		})
		if err != nil {
			util.RespondWithError(err, r, w)
			return
		}
	}

	util.RespondWithData(election, w)

	return
}

func handleBlock(w http.ResponseWriter, r *http.Request) {
	networkID, err := util.ParseNetworkID(r)
	if err != nil {
		util.RespondWithError(err, r, w)
		return
	}

	callOpts, err := getCallOpts(networkID, w, r)
	if err != nil {
		util.RespondWithError(err, r, w)
		return
	}

	fmt.Printf(" \n\n\n call opts %+v \n\n\n ", callOpts)

	block, err := getBlockByNumber(networkID, callOpts.BlockNumber)
	if err != nil {
		util.RespondWithError(err, r, w)
		return
	}

	d, err := json.Marshal(types.JSONResponse{
		Data: block,
	})
	if err != nil {
		util.RespondWithError(err, r, w)
		return
	}

	util.RespondWithData(d, w)
}
