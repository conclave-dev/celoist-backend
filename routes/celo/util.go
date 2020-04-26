package celo

import (
	"net/http"

	"github.com/conclave-dev/celoist-backend/util"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Define default values for callOpts struct's fields
func setCallOptsDefaults() (callOpts *bind.CallOpts, err error) {
	j, err := getBlockNumber()
	if err != nil {
		return
	}

	// Decode hex-encoded block number
	b, err := hexutil.DecodeBig(j.Result)
	if err != nil {
		return
	}

	callOpts = &bind.CallOpts{
		BlockNumber: b,
	}

	return callOpts, err
}

func getCallOpts(w http.ResponseWriter, r *http.Request) (callOpts *bind.CallOpts, err error) {
	err = util.ParseResponse(r.Body, &callOpts)
	if err != nil {
		return
	}

	if callOpts == nil {
		return setCallOptsDefaults()
	}

	return
}
