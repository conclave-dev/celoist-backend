package celo

import (
	"github.com/gorilla/mux"
)

// AddRoutes includes routes related to the `eth` endpoints on the RPC
func AddRoutes(router *mux.Router) {
	router.HandleFunc("/celo/election", handleElection)
	router.HandleFunc("/celo/block", handleBlock)
}
