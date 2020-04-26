package kvstore

const (
	connProto = "tcp"
	connURL   = "localhost:6379"

	// Reusable hash key constants
	ELECTIONS           = "elections"
	GROUPS              = "groups"
	MEMBERS             = "members"
	ACCOUNTS            = "accounts"
	BLOCKS              = "blocks"
	SYNCED_BLOCK_NUMBER = "syncedBlockNumber"
)
