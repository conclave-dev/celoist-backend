package kvstore

import (
	"encoding/json"
	"log"

	"github.com/conclave-dev/celoist-backend/util"
)

// StringifyJSON converts the JSON object into a string
func StringifyJSON(data interface{}) string {
	b, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	return string(b[:])
}

// GetHashKey returns the hash key for the specified base key and currently set networkID
func GetHashKey(baseKey string) string {
	return baseKey + "-" + util.NetworkID
}
