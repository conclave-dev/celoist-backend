package celo

import (
	"bytes"
	"log"
	"net/http"

	solsha3 "github.com/miguelmota/go-solidity-sha3"

	"github.com/conclave-dev/celoist-backend/util"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Define default values for callOpts struct's fields
func setCallOptsDefaults(networkID string) (callOpts *bind.CallOpts, err error) {
	j, err := getBlockNumber(networkID)
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

func getCallOpts(networkID string, w http.ResponseWriter, r *http.Request) (callOpts *bind.CallOpts, err error) {
	err = util.ParseResponse(r.Body, &callOpts)
	if err != nil {
		return
	}

	if callOpts == nil {
		return setCallOptsDefaults(networkID)
	}

	return
}

func callJSONRPC(networkID string, data []byte, v interface{}) error {
	endpoint, err := util.GetNetworkEndpoint(networkID)
	if err != nil {
		return err
	}

	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = util.ParseResponse(resp.Body, &v)
	if err != nil {
		return err
	}

	return err
}

func getContractAddress(networkID string, contract string) common.Address {
	hash := solsha3.SoliditySHA3([]string{"string"}, []interface{}{contract})

	var rawAddress [32]byte
	copy(rawAddress[:], hash)

	registryCaller, err := util.GetNetworkRegistry(networkID)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress, err := registryCaller.GetAddressFor(nil, rawAddress)
	if err != nil {
		log.Fatal(err)
	}

	return contractAddress
}
