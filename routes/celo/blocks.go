package celo

import (
	"fmt"
	"math/big"
)

func getBlockByNumber(num *big.Int) (GetBlockByNumberResponse, error) {
	d := []byte(fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["0x%x", true],"id":1}`, num.Int64()))

	var v GetBlockByNumberResponse
	err := callJSONRPC(d, &v)
	if err != nil {
		return v, err
	}

	return v, nil
}

func getBlockNumber() (GetBlockNumberResponse, error) {
	d := []byte(`{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}`)
	var v GetBlockNumberResponse
	err := callJSONRPC(d, &v)
	if err != nil {
		return v, err
	}

	return v, nil
}
