package nft

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

// RegisterPeriodicOperations implements modules.Module
func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	log.Debug().Str("module", "nft").Msg("setting up periodic tasks")

	if _, err := scheduler.Every(1).Minutes().Do(func() {
		height, err := m.db.GetLastBlockHeight()
		if err != nil {
			log.Error().Str("module", "nft").Err(err).Msg("unable to get last block height")
			return
		}

		nfts, err := m.db.GetNFTsToUpdate()
		if err != nil {
			log.Error().Str("module", "nft").Err(err).Msg("unable to get nfts to update")
			return
		}

		for _, nft := range nfts {
			updatedNFT, ok := m.keeper.GetNFT(nft.Address, height)
			if !ok {
				log.Error().Str("module", "nft").Str("address", nft.Address).Msg("nft does not exist, skipping")
				continue
			}

			if err = m.db.SaveNFT(nft.Address, nft.Owner, updatedNFT.AvailableToWithdraw, updatedNFT.LastVestingBlock, updatedNFT.VestingPeriodsCount, updatedNFT.RewardPerPeriod, updatedNFT.VestingCounter, nft.Denom); err != nil {
				log.Error().Str("module", "nft").Err(err).Msg("unable to save nft")
			}
		}

	}); err != nil {
		return fmt.Errorf("error while setting up bank periodic operation: %s", err)
	}

	return nil
}
