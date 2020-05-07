package kvstore

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

func init() {
	// Get synced block number on load and set to 0 if does not exist
	_, err := GetSyncedBlockNumber()
	if err != nil {
		_, err := SetSyncedBlockNumber(big.NewInt(0))
		if err != nil {
			log.Fatal(err)
		}
	}
}

// DoesBlockExist checks whether a block already exists at a given number
func DoesBlockExist(n *big.Int) bool {
	b, err := HExists(GetHashKeyForNetwork(BLOCKS), n.String())
	if err != nil {
		log.Fatal(err)
	}

	return b
}

// SetBlock stores block data at a given number
func SetBlock(n *big.Int, data types.Block) (interface{}, error) {
	return HSet(GetHashKeyForNetwork(BLOCKS), n.String(), StringifyJSON(data))
}

// GetBlock retrieves block data at a given number
func GetBlock(n *big.Int) (string, error) {
	return HGet(GetHashKeyForNetwork(BLOCKS), n.String())
}

// DeleteBlock deletes block data at a given number
func DeleteBlock(n *big.Int) (interface{}, error) {
	return HDelete(GetHashKeyForNetwork(BLOCKS), n.String())
}

// SetSyncedBlockNumber stores block data at a given number
func SetSyncedBlockNumber(n *big.Int) (interface{}, error) {
	return Set(GetHashKeyForNetwork(SYNCED_BLOCK_NUMBER), n.String())
}

// GetSyncedBlockNumber retrieves block data at a given number
func GetSyncedBlockNumber() (string, error) {
	return Get(GetHashKeyForNetwork(SYNCED_BLOCK_NUMBER))
}
