package util

// NetworkID stores the identifier of the network currently served by the server
var NetworkID = ""

// NetworkOrder stores the internal index/order of the currently selected network
var NetworkOrder = -1

// NetworkEndpoint stores the url endpoint of the currently selected network
var NetworkEndpoint = ""

// RegistryContractAddress stores the address of the contract holding the addresses of Celo contracts
const RegistryContractAddress = "0x000000000000000000000000000000000000ce10"

// SetNetworkID sets the networkID used by the backend server
func SetNetworkID(networkID string) {
	for idx, identifier := range NetworkIdentifiers {
		if identifier == networkID {
			NetworkOrder = idx
			NetworkID = identifier
			NetworkEndpoint = NetworkEndpoints[NetworkOrder]
			break
		}
	}

	if NetworkOrder == -1 {
		panic("Invalid or unsupported network")
	}
}
