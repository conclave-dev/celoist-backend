package kvstore

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type GroupAddresses []common.Address
type MemberAddresses []common.Address
type Groups map[common.Address]Group
type Members map[common.Address]Member

type Election struct {
	GroupAddresses GroupAddresses `json:"groupAddresses"`
	GroupVotes     []*big.Int     `json:"groupVotes"`
	Groups         Groups         `json:"groups"`
}

type Group struct {
	Name                string          `json:"name"`
	Address             common.Address  `json:"address"`
	Commission          *big.Int        `json:"commission"`
	NextCommission      *big.Int        `json:"nextCommission"`
	NextCommissionBlock *big.Int        `json:"nextCommissionBlock"`
	MemberAddresses     MemberAddresses `json:"memberAddresses"`
	Members             Members         `json:"members"`
	LastSlash           *big.Int        `json:"lastSlashed"`
	SlashMultiplier     *big.Int        `json:"slashingMultiplier"`
	Capacity            *big.Int        `json:"capacity"`
}

type Member struct {
	Name           string         `json:"name"`
	Address        common.Address `json:"address"`
	Signer         common.Address `json:"signer"`
	Score          *big.Int       `json:"score"`
	ECDSAPublicKey []byte         `json:"ecdsaPublicKey"`
	BLSPublicKey   []byte         `json:"blsPublicKey"`
}
