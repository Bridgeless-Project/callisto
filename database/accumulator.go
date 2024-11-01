package database

import (
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	dbtypes "github.com/forbole/bdjuno/v4/database/types"
	"github.com/forbole/bdjuno/v4/types"
	"github.com/lib/pq"
)

// SaveAdmin allows to save new Admin
func (db *Db) SaveAdmin(address string, vestingCount, lastVestingTime, vestingPeriod int64, rewardPerPeriod sdk.Coin, denom string) error {
	query := `
		INSERT INTO admins_vesting(address, vesting_period, reward_per_period, last_vesting_time, vesting_counter, denom) 
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (address) DO UPDATE
	SET vesting_counter = excluded.vesting_counter,
		last_vesting_time = excluded.last_vesting_time
	WHERE admins_vesting.address <= excluded.address
	`
	_, err := db.SQL.Exec(query, address, vestingPeriod, pq.Array(dbtypes.NewDbCoins(sdk.NewCoins(rewardPerPeriod))), lastVestingTime, vestingCount, denom)
	if err != nil {
		return fmt.Errorf("error while storing admin vesting info: %s", err)
	}

	return nil
}

// -------------------------------------------------------------------------------------------------------------------

// SaveAccumulatorParams allows to store the given params inside the database
func (db *Db) SaveAccumulatorParams(params *types.AccumulatorParams) error {
	paramsBz, err := json.Marshal(&params.Params)
	if err != nil {
		return fmt.Errorf("error while marshaling accumulator params: %s", err)
	}

	stmt := `
		INSERT INTO accumulator_params (params, height) 
		VALUES ($1, $2)
		ON CONFLICT (one_row_id) DO UPDATE 
			SET params = excluded.params,
				height = excluded.height
		WHERE accumulator_params.height <= excluded.height
	`

	_, err = db.SQL.Exec(stmt, string(paramsBz), params.Height)
	if err != nil {
		return fmt.Errorf("error while storing accumulator params: %s", err)
	}

	return nil
}

// GetAdmins returns all the admins that are currently stored inside the database.
func (db *Db) GetAdmins() ([]dbtypes.AdminVestingRow, error) {
	var rows []dbtypes.AdminVestingRow
	err := db.Sqlx.Select(&rows, `SELECT * FROM admins_vesting WHERE to_timestamp(last_vesting_time) + INTERVAL '1 second' * last_vesting_time < NOW();`)
	return rows, err
}
