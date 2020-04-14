package kvstore

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Election struct {
	GroupAddresses []common.Address `json:"groupAddresses"`
	GroupVotes     []*big.Int       `json:"groupVotes"`
}

type Group struct {
	Address         common.Address   `json:"address"`
	Votes           *big.Int         `json:"votes"`
	Commission      *big.Int         `json:"commission"`
	MemberAddresses []common.Address `json:"memberAddresses"`
	LastSlash       *big.Int         `json:"lastSlashed"`
	SlashMultiplier *big.Int         `json:"slashingMultiplier"`
}

type Member struct {
	Address        common.Address `json:"address"`
	Signer         common.Address `json:"signer"`
	Affiliation    common.Address `json:"affiliation"`
	Score          *big.Int       `json:"score"`
	ECDSAPublicKey []byte         `json:"ecdsaPublicKey"`
	BLSPublicKey   []byte         `json:"blsPublicKey"`
}
