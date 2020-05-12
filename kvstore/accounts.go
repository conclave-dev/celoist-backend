package kvstore

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
)

// DoesAccountExist checks whether a block already exists at a given number
func DoesAccountExist(networkID string, addr common.Address) bool {
	b, err := HExists(GetHashKey(networkID, BLOCKS), addr.String())
	if err != nil {
		log.Fatal(err)
	}

	return b
}

// SetAccount stores block data at a given number
func SetAccount(networkID string, addr common.Address, data Account) (interface{}, error) {
	return HSet(GetHashKey(networkID, BLOCKS), addr.String(), StringifyJSON(data))
}

// GetAccount retrieves block data at a given number
func GetAccount(networkID string, addr common.Address) (string, error) {
	return HGet(GetHashKey(networkID, BLOCKS), addr.String())
}

// DeleteAccount deletes block data at a given number
func DeleteAccount(networkID string, addr common.Address) (interface{}, error) {
	return HDelete(GetHashKey(networkID, BLOCKS), addr.String())
}
