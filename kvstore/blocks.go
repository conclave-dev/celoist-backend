package kvstore

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

// DoesBlockExist checks whether a block already exists at a given number
func DoesBlockExist(n *big.Int) bool {
	b, err := HExists(BLOCKS, n.String())
	if err != nil {
		log.Fatal(err)
	}

	return b
}

// SetBlock stores block data at a given number
func SetBlock(n *big.Int, data types.Block) (interface{}, error) {
	return HSet(BLOCKS, n.String(), StringifyJSON(data))
}

// GetBlock retrieves block data at a given number
func GetBlock(n *big.Int) (string, error) {
	return HGet(BLOCKS, n.String())
}

// DeleteBlock deletes block data at a given number
func DeleteBlock(n *big.Int) (interface{}, error) {
	return HDelete(BLOCKS, n.String())
}

// SetBlock stores block data at a given number
func SetSyncedBlockNumber(n *big.Int) (interface{}, error) {
	return Set(SYNCED_BLOCK_NUMBER, n.String())
}

// GetBlock retrieves block data at a given number
func GetSyncedBlockNumber() (string, error) {
	return Get(SYNCED_BLOCK_NUMBER)
}
