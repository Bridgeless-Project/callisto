package types

import "math/big"

// NFTEventsRow represents a single row inside the nft_events table
type NFTEventsRow struct {
	ID         uint64   `db:"id"`
	EventType  string   `db:"event_type"`
	NftAddress string   `db:"nft_address"`
	Owner      string   `db:"owner"`
	NewOwner   string   `db:"new_owner"`
	Validator  string   `db:"validator"`
	Amount     *big.Int `db:"amount"`
}

func NewNFTEventsRow(id uint64, eventType string, nftAddress string, owner string) *NFTEventsRow {
	return &NFTEventsRow{
		ID:         id,
		EventType:  eventType,
		NftAddress: nftAddress,
		Owner:      owner,
	}
}

// NFTsRow represents a single row inside the nfts table
type NFTsRow struct {
	Address     string          `db:"address"`
	Owner       string          `db:"owner"`
	Delegations []NFTDelegation `db:"delegations"`
}

func NewNFTsRow(address, owner string) *NFTsRow {
	return &NFTsRow{
		Address: address,
		Owner:   owner,
	}
}

// NFTDelegation represent delegation type
type NFTDelegation struct {
	Validator string   `db:"validator"`
	Amount    *big.Int `db:"amount"`
	Timestamp uint64   `db:"timestamp"`
}

func NewDelegations(validator string, timestamp uint64, amount *big.Int) *NFTDelegation {
	return &NFTDelegation{
		Validator: validator,
		Timestamp: timestamp,
		Amount:    amount,
	}
}
