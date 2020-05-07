package celo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/conclave-dev/celoist-backend/kvstore"
	"github.com/conclave-dev/go-celo/types"
	"github.com/conclave-dev/go-celo/util"
)

func handleElection(w http.ResponseWriter, r *http.Request) {
	var election []byte
	var err error

	callOpts, err := getCallOpts(w, r)
	if err != nil {
		util.RespondWithError(err, r, w)
	}

	epochNumber, err := getEpochNumber(callOpts)
	if err != nil {
		util.RespondWithError(err, r, w)
	}

	ens := epochNumber.String()

	if kvstore.DoesElectionExist(ens) {
		election, err = getElection(ens)
		if err != nil {
			util.RespondWithError(err, r, w)
		}
	} else {
		e, err := setElection(callOpts, ens)
		if err != nil {
			util.RespondWithError(err, r, w)
		}

		election, err = json.Marshal(types.JSONResponse{
			Data: e,
		})
		if err != nil {
			util.RespondWithError(err, r, w)
		}
	}

	util.RespondWithData(election, w)

	return
}

func handleBlock(w http.ResponseWriter, r *http.Request) {
	callOpts, err := getCallOpts(w, r)
	if err != nil {
		util.RespondWithError(err, r, w)
	}

	fmt.Printf(" \n\n\n call opts %+v \n\n\n ", callOpts)

	block, err := getBlockByNumber(callOpts.BlockNumber)
	if err != nil {
		util.RespondWithError(err, r, w)
	}

	d, err := json.Marshal(types.JSONResponse{
		Data: block,
	})
	if err != nil {
		util.RespondWithError(err, r, w)
	}

	util.RespondWithData(d, w)
}

func handleBlockNumber(w http.ResponseWriter, r *http.Request) {
	callOpts, err := getCallOpts(w, r)
	if err != nil {
		util.RespondWithError(err, r, w)
	}

	fmt.Printf(" \n\n\n call opts %+v \n\n\n ", callOpts)

	blockNumber, err := getBlockNumber()
	if err != nil {
		util.RespondWithError(err, r, w)
	}

	d, err := json.Marshal(types.JSONResponse{
		Data: blockNumber,
	})
	if err != nil {
		util.RespondWithError(err, r, w)
	}

	util.RespondWithData(d, w)
}
