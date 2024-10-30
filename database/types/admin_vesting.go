package types

type AdminVestingRow struct {
	ID                  int64   `db:"id"`
	Address             string  `db:"address"`
	VestingPeriod       int64   `db:"vesting_period"`
	RewardsPerPeriod    DbCoins `db:"reward_per_period"`
	LastVestingTime     int64   `db:"last_vesting_time"`
	VestingCounter      int64   `db:"vesting_counter"`
	VestingPeriodsCount int64   `db:"vesting_periods_count"`
	Denom               string  `db:"denom"`
}

func NewAdminVestingRow(id int64, address string, vestingPeriod int64, rewardsPerPeriod DbCoins, lastVestingTime, vestingCounter int64, denom string) AdminVestingRow {
	return AdminVestingRow{
		ID:               id,
		Address:          address,
		VestingPeriod:    vestingPeriod,
		RewardsPerPeriod: rewardsPerPeriod,
		LastVestingTime:  lastVestingTime,
		VestingCounter:   vestingCounter,
		Denom:            denom,
	}
}

func (r AdminVestingRow) Equal(s AdminVestingRow) bool {
	return r.RewardsPerPeriod.Equal(&s.RewardsPerPeriod) &&
		r.ID == s.ID &&
		r.Address == s.Address
}
