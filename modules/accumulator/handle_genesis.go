package accumulator

import (
	"encoding/json"
	"fmt"
	accumulatortypes "github.com/cosmos/cosmos-sdk/x/accumulator/types"

	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/forbole/bdjuno/v4/types"

	"github.com/rs/zerolog/log"
)

// HandleGenesis implements modules.Module
func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", "accumulator").Msg("parsing genesis")

	// Read the genesis state
	var genState accumulatortypes.GenesisState
	err := m.cdc.UnmarshalJSON(appState[accumulatortypes.ModuleName], &genState)
	if err != nil {
		return fmt.Errorf("error while reading accumulator genesis data: %s", err)
	}

	// Save the params
	err = m.db.SaveAccumulatorParams(types.NewAccumulatorParams(genState.Params, doc.InitialHeight))
	if err != nil {
		return fmt.Errorf("error while storing genesis accumulator params: %s", err)
	}

	return nil
}
