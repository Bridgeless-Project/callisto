package nft

import (
	"encoding/json"
	"fmt"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft/types"

	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/rs/zerolog/log"
)

// HandleGenesis implements modules.Module
func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", "nft").Msg("parsing genesis")

	// Read the genesis state
	var genState nfttypes.GenesisState
	err := m.cdc.UnmarshalJSON(appState[nfttypes.ModuleName], &genState)
	if err != nil {
		return fmt.Errorf("error while reading nft genesis data: %s", err)
	}

	// Save NFTs
	for _, nft := range genState.Nfts {
		err = m.db.SaveNFT(nft.Address, nft.Owner, nft.AvailableToWithdraw, nft.LastVestingTime, nft.VestingPeriod, nft.RewardPerPeriod, nft.VestingPeriodsCount, nft.Denom)
		if err != nil {
			return fmt.Errorf("error while storing genesis nft: %s", err)
		}
	}
	return nil
}
