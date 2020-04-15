package celo

import (
	"encoding/json"
	"net/http"

	"github.com/conclave-dev/celoist-backend/kvstore"
	"github.com/conclave-dev/go-celo/types"
	"github.com/conclave-dev/go-celo/util"
)

const (
	celoRPCAPI = "https://geth.celoist.com"
)

func init() {
	util.SetupClients()
}

func handleElection(w http.ResponseWriter, r *http.Request) {
	var election []byte
	var err error

	opts := getCallOpts(w, r)
	epochNumber, err := getEpochNumber(opts)
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
		e, err := setElection(opts, ens)
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
