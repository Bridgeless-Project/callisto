package accumulator

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
	"math"
)

// RegisterPeriodicOperations implements modules.Module
func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	log.Debug().Str("module", m.Name()).Msg("setting up periodic tasks")
	pagination := &query.PageRequest{
		Limit: math.MaxInt32,
	}

	logErr := log.Error().Str("module", m.Name())

	// TODO set up time before deploy
	if _, err := scheduler.Every(5).Minutes().Do(func() {
		height, err := m.db.GetLastBlockHeight()
		if err != nil {
			logErr.Err(err).Msg("unable to get last block height")
			return
		}

		admins, err := m.keeper.GetAdmins(pagination, height)
		if err != nil {
			logErr.Err(err).Msg("unable to get nfts")
			return
		}

		for _, admin := range admins {
			err = m.db.SaveAdmin(admin.Address, admin.VestingPeriodsCount, admin.LastVestingTime, admin.VestingPeriod, admin.RewardPerPeriod, admin.Denom)
			if err != nil {
				logErr.Err(err).Msg("unable to save nft")
			}
		}

	}); err != nil {
		return fmt.Errorf("error while setting up bank periodic operation: %s", err)
	}

	return nil
}
