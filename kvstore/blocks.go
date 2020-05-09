package kvstore

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

// DoesBlockExist checks whether a block already exists at a given number
func DoesBlockExist(networkID string, n *big.Int) bool {
	b, err := HExists(GetHashKey(networkID, BLOCKS), n.String())
	if err != nil {
		log.Fatal(err)
	}

	return b
}

// SetBlock stores block data at a given number
func SetBlock(networkID string, n *big.Int, data types.Block) (interface{}, error) {
	return HSet(GetHashKey(networkID, BLOCKS), n.String(), StringifyJSON(data))
}

// GetBlock retrieves block data at a given number
func GetBlock(networkID string, n *big.Int) (string, error) {
	return HGet(GetHashKey(networkID, BLOCKS), n.String())
}

// DeleteBlock deletes block data at a given number
func DeleteBlock(networkID string, n *big.Int) (interface{}, error) {
	return HDelete(GetHashKey(networkID, BLOCKS), n.String())
}

// SetSyncedBlockNumber stores block data at a given number
func SetSyncedBlockNumber(networkID string, n *big.Int) (interface{}, error) {
	return Set(GetHashKey(networkID, SYNCED_BLOCK_NUMBER), n.String())
}

// GetSyncedBlockNumber retrieves block data at a given number
func GetSyncedBlockNumber(networkID string) (string, error) {
	return Get(GetHashKey(networkID, SYNCED_BLOCK_NUMBER))
}
