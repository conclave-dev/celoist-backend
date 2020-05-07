package kvstore

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
)

// DoesAccountExist checks whether a block already exists at a given number
func DoesAccountExist(addr common.Address) bool {
	b, err := HExists(GetHashKeyForNetwork(BLOCKS), addr.String())
	if err != nil {
		log.Fatal(err)
	}

	return b
}

// SetAccount stores block data at a given number
func SetAccount(addr common.Address, data Account) (interface{}, error) {
	return HSet(GetHashKeyForNetwork(BLOCKS), addr.String(), StringifyJSON(data))
}

// GetAccount retrieves block data at a given number
func GetAccount(addr common.Address) (string, error) {
	return HGet(GetHashKeyForNetwork(BLOCKS), addr.String())
}

// DeleteAccount deletes block data at a given number
func DeleteAccount(addr common.Address) (interface{}, error) {
	return HDelete(GetHashKeyForNetwork(BLOCKS), addr.String())
}
