package accumulator

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

// RegisterPeriodicOperations implements modules.Module
func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	log.Debug().Str("module", m.Name()).Msg("setting up periodic tasks")

	logErr := log.Error().Str("module", m.Name())

	// TODO set up time before deploy
	if _, err := scheduler.Every(5).Minutes().Do(func() {
		height, err := m.db.GetLastBlockHeight()
		if err != nil {
			logErr.Err(err).Msg("unable to get last block height")
			return
		}

		admins, err := m.db.GetAdmins()
		if err != nil {
			logErr.Err(err).Msg("unable to get admins")
			return
		}

		for _, admin := range admins {
			vestingInfo, err := m.keeper.GetAdminByAddress(admin.Address, height)
			err = m.db.SaveAdmin(admin.Address, vestingInfo.VestingPeriodsCount, vestingInfo.LastVestingTime, vestingInfo.VestingPeriod, vestingInfo.RewardPerPeriod, admin.Denom)
			if err != nil {
				logErr.Err(err).Msg("unable to save admin vesting info")
			}
		}

	}); err != nil {
		return fmt.Errorf("error while setting up bank periodic operation: %s", err)
	}

	return nil
}
