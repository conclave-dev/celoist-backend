package util

import (
	"errors"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/conclave-dev/go-celo/core/celo/common/registry"
)

// rpcClients stores the clients connected to supported networks
var rpcClients = make([]*ethclient.Client, 0)

// registryCallers stores the registry contract instance for supported networks
var registryCallers = make([]*registry.RegistryCaller, 0)

// SetupClients initializes and caches the RPC clients as well as the registry callers for all supported networks
func SetupClients() {
	for idx := range NetworkIdentifiers {
		client, err := ethclient.Dial(NetworkEndpoints[idx])
		if err != nil {
			log.Fatal(err)
		}

		rpcClients = append(rpcClients, client)

		address := common.HexToAddress(registryContractAddress)
		caller, err := registry.NewRegistryCaller(address, client)
		if err != nil {
			log.Fatal(err)
		}

		registryCallers = append(registryCallers, caller)
	}
}

// GetNetworkIndex finds the index of the specified networkID, returns -1 if invalid/not found
func GetNetworkIndex(networkID string) int {
	for idx, identifier := range NetworkIdentifiers {
		if identifier == networkID {
			return idx
		}
	}

	return -1
}

// GetNetworkClient returns the network RPC client
func GetNetworkClient(networkID string) (*ethclient.Client, error) {
	index := GetNetworkIndex(networkID)
	if index != -1 {
		return rpcClients[index], nil
	}

	return nil, errors.New("Invalid or unsupported network")
}

// GetNetworkRegistry returns the network registry caller
func GetNetworkRegistry(networkID string) (*registry.RegistryCaller, error) {
	index := GetNetworkIndex(networkID)
	if index != -1 {
		return registryCallers[index], nil
	}

	return nil, errors.New("Invalid or unsupported network")
}

// GetNetworkEndpoint returns the API URL for the specified network
func GetNetworkEndpoint(networkID string) (string, error) {
	index := GetNetworkIndex(networkID)
	if index != -1 {
		return NetworkEndpoints[index], nil
	}

	return "", errors.New("Invalid or unsupported network")
}
