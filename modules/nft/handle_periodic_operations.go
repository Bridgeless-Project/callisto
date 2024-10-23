package nft

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
	"math"
)

// RegisterPeriodicOperations implements modules.Module
func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	log.Debug().Str("module", "nft").Msg("setting up periodic tasks")
	pagination := &query.PageRequest{
		Limit: math.MaxInt32,
	}

	if _, err := scheduler.Every(5).Minutes().Do(func() {
		height, err := m.db.GetLastBlockHeight()
		if err != nil {
			log.Error().Str("module", "nft").Err(err).Msg("unable to get last block height")
			return
		}

		val, _, err := m.keeper.GetNFTsWithPagination(pagination, height)
		if err != nil {
			log.Error().Str("module", "nft").Err(err).Msg("unable to get nfts")
			return
		}

		for _, nft := range val {
			if err = m.db.SaveNFT(nft.Address, nft.Owner, nft.AvailableToWithdraw); err != nil {
				log.Error().Str("module", "nft").Err(err).Msg("unable to save nft")
			}
		}

	}); err != nil {
		return fmt.Errorf("error while setting up bank periodic operation: %s", err)
	}

	return nil
}
