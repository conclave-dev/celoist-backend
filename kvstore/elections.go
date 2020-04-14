package kvstore

import (
	"log"
)

// DoesElectionExist checks whether an update already exists
func DoesElectionExist(election string) bool {
	b, err := Exists(ELECTIONS, election)
	if err != nil {
		log.Fatal(err)
	}

	return b
}

// SetElection stores election data at a given election number
func SetElection(num string, data Election) (interface{}, error) {
	return Set(ELECTIONS, num, StringifyJSON(data))
}

// GetElection retrieves election data at a given election number
func GetElection(num string) (string, error) {
	return Get(ELECTIONS, num)
}

// DeleteElection deletes election data at a given election number
func DeleteElection(num string) (interface{}, error) {
	return Delete(ELECTIONS, num)
}
