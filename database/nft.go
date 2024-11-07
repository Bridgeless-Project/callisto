package database

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	dbtypes "github.com/forbole/bdjuno/v4/database/types"
	"github.com/lib/pq"
)

// SaveNFTEvent allows to save new NFTEvent
func (db *Db) SaveNFTEvent(eventType string, nftAddress, validator, newValidator, newOwner, owner string, amount sdk.Coin) error {
	query := `
		INSERT INTO nft_events(event_type, nft_address, owner, new_owner, validator,new_validator, amount) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := db.SQL.Exec(query, eventType, nftAddress, owner, newOwner, validator, newValidator, pq.Array(dbtypes.NewDbCoins(sdk.NewCoins(amount))))
	if err != nil {
		return fmt.Errorf("error while storing nft: %s", err)
	}

	return nil
}

// -------------------------------------------------------------------------------------------------------------------

// SaveNFT allows to save new NFT
func (db *Db) SaveNFT(address, owner string, availableAmount sdk.Coin, lastVestingTime int64, vestingPeriod int64, rewardPerPeriod sdk.Coin, vestingCounter int64, denom string) error {
	query := `
		INSERT INTO nfts(address, owner, available_amount, vesting_period, reward_per_period, last_vesting_time, vesting_counter, denom) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (address) DO UPDATE
	SET owner = excluded.owner,
		available_amount = excluded.available_amount,
		last_vesting_time = excluded.last_vesting_time,
		vesting_counter = excluded.vesting_counter
	WHERE nfts.address <= excluded.address
	`
	_, err := db.SQL.Exec(query, address, owner, pq.Array(dbtypes.NewDbCoins(sdk.NewCoins(availableAmount))), vestingPeriod, pq.Array(dbtypes.NewDbCoins(sdk.NewCoins(rewardPerPeriod))), lastVestingTime, vestingCounter, denom)
	if err != nil {
		return fmt.Errorf("error while storing nft: %s", err)
	}

	return nil
}

// GetNFTsToUpdate returns all the nfts that are currently stored inside the database and should be updated.
func (db *Db) GetNFTsToUpdate() ([]dbtypes.NFTsRow, error) {
	var rows []dbtypes.NFTsRow
	err := db.Sqlx.Select(&rows, `SELECT * FROM nfts WHERE to_timestamp(last_vesting_time) + INTERVAL '1 second' * last_vesting_time < NOW();`)
	return rows, err
}
