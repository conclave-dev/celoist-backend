package kvstore

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Election map[*big.Int]Group

type Group struct {
	Address         common.Address `json:"address"`
	Votes           *big.Int       `json:"votes"`
	Name            string         `json:"name,omitempty"`
	Commission      *big.Int       `json:"commission,omitempty"`
	Members         []Member       `json:"members,omitempty"`
	LastSlash       *big.Int       `json:"lastSlashed,omitempty"`
	SlashMultiplier *big.Int       `json:"slashingMultiplier,omitempty"`
}

// Member is a member that the user has added
type Member struct {
	Address        common.Address `json:"address"`
	Score          *big.Int       `json:"score"`
	Signer         common.Address `json:"signer"`
	Affiliation    common.Address `json:"affiliation"`
	ECDSAPublicKey []byte         `json:"ecdsaPublicKey"`
	BLSPublicKey   []byte         `json:"blsPublicKey"`
	Name           string         `json:"name,omitempty"`
}
