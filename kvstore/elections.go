package kvstore

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
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

// DoesGroupExist checks whether a group already exists
func DoesGroupExist(addr common.Address) bool {
	b, err := Exists(GROUPS, addr.String())
	if err != nil {
		log.Fatal(err)
	}

	return b
}

// SetGroup stores election group data at a given group address
func SetGroup(addr common.Address, data Group) (interface{}, error) {
	return Set(GROUPS, addr.String(), StringifyJSON(data))
}

// GetGroup retrieves election group data at a given group address
func GetGroup(addr common.Address) (string, error) {
	return Get(GROUPS, addr.String())
}

// DeleteGroup deletes election group data at a given group address
func DeleteGroup(addr common.Address) (interface{}, error) {
	return Delete(GROUPS, addr.String())
}

// DoesMemberExist checks whether a member already exists
func DoesMemberExist(addr common.Address) bool {
	b, err := Exists(MEMBERS, addr.String())
	if err != nil {
		log.Fatal(err)
	}

	return b
}

// SetMember stores election member data at a given member address
func SetMember(addr common.Address, data Member) (interface{}, error) {
	return Set(MEMBERS, addr.String(), StringifyJSON(data))
}

// GetMember retrieves election member data at a given member address
func GetMember(addr common.Address) (string, error) {
	return Get(MEMBERS, addr.String())
}

// DeleteMember deletes election member data at a given member address
func DeleteMember(addr common.Address) (interface{}, error) {
	return Delete(MEMBERS, addr.String())
}
