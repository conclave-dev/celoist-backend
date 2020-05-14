package kvstore

import (
	"log"
)

// DoesElectionExist checks whether an update already exists
func DoesElectionExist(networkID string, election string) bool {
	b, err := HExists(GetHashKey(networkID, ELECTIONS), election)
	if err != nil {
		log.Fatal(err)
	}

	return b
}

// SetElection stores election data at a given election number
func SetElection(networkID string, num string, data Election) (interface{}, error) {
	return HSet(GetHashKey(networkID, ELECTIONS), num, StringifyJSON(data))
}

// GetElection retrieves election data at a given election number
func GetElection(networkID string, num string) (string, error) {
	return HGet(GetHashKey(networkID, ELECTIONS), num)
}

// DeleteElection deletes election data at a given election number
func DeleteElection(networkID string, num string) (interface{}, error) {
	return HDelete(GetHashKey(networkID, ELECTIONS), num)
}
