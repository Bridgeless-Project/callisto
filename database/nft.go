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
		return fmt.Errorf("error while storing community pool: %s", err)
	}

	return nil
}

// -------------------------------------------------------------------------------------------------------------------

// SaveNFT allows to save new NFT
func (db *Db) SaveNFT(address, owner string, availableAmount sdk.Coin) error {
	query := `
		INSERT INTO nft(address, owner, available_amount) 
		VALUES ($1, $2, $3)
		ON CONFLICT (address) DO UPDATE
	SET owner = excluded.owner,
		available_amount = excluded.available_amount
	WHERE nft.address <= excluded.address
	`
	_, err := db.SQL.Exec(query, address, owner, pq.Array(dbtypes.NewDbCoins(sdk.NewCoins(lockedAmount))), pq.Array(dbtypes.NewDbCoins(sdk.NewCoins(availableAmount))), delegations)
	if err != nil {
		return fmt.Errorf("error while storing community pool: %s", err)
	}

	return nil
}
