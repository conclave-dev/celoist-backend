package kvstore

import (
	"encoding/json"
	"log"
)

// StringifyJSON converts the JSON object into a string
func StringifyJSON(data interface{}) string {
	b, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	return string(b[:])
}

// GetHashKeyForNetwork returns the hash key for the specified base key and networkID
func GetHashKeyForNetwork(baseKey string, networkID string) string {
	var supportedNetwork = false

	for _, identifier := range NETWORK_IDENTIFIERS {
		if identifier == networkID {
			supportedNetwork = true
			break
		}
	}

	if !supportedNetwork {
		panic("Invalid or unsupported network")
	}

	return baseKey + "-" + networkID
}
