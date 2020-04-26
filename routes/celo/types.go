package celo

import "github.com/ethereum/go-ethereum/core/types"

// JSONRPCResponse is the response from a Web3 JSON RPC call
type JSONRPCResponse struct {
	ID      int    `json:"id"`
	JSONRPC string `json:"jsonrpc"`
}

type GetBlockByNumberResponse struct {
	JSONRPCResponse
	Result *types.Header `json:"result"`
}

type GetBlockNumberResponse struct {
	JSONRPCResponse
	Result string `json:"result"`
}

type HandleBlockResponse struct {
	BlockNumber int `json:"blockNumber"`
}
