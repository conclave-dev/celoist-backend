package celo

import (
	"log"
	"net/http"

	"github.com/conclave-dev/go-celo/util"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func getCallOpts(w http.ResponseWriter, r *http.Request) *bind.CallOpts {
	var args util.CallOpts
	err := util.ParseRequestParameters(w, r, &args)
	if err != nil {
		log.Fatal(err)
	}

	return args.CallOpts
}
