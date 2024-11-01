package types

import (
	"math/big"
)

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
	Address         string  `db:"address"`
	Owner           string  `db:"owner"`
	AvailableAmount DbCoins `db:"available_amount"`
	VestingPeriod   int     `db:"vesting_period"`
	RewardPerPeriod DbCoins `db:"reward_per_period"`
	LastVestingTime int     `db:"last_vesting_time"`
	VestingCounter  int16   `db:"vesting_counter"`
	Denom           string  `db:"denom"`
}

// NewNFTsRow creates a new instance of NFTsRow
func NewNFTsRow(address, owner string, availableAmount, rewardPerPeriod DbCoins, vestingPeriod, lastVestingTime int, vestingCounter int16, denom string) *NFTsRow {
	return &NFTsRow{
		Address:         address,
		Owner:           owner,
		AvailableAmount: availableAmount,
		VestingPeriod:   vestingPeriod,
		RewardPerPeriod: rewardPerPeriod,
		LastVestingTime: lastVestingTime,
		VestingCounter:  vestingCounter,
		Denom:           denom,
	}
}
