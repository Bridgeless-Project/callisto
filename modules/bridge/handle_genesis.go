package bridge

import (
	"encoding/json"
	bridgetypes "github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
	"github.com/pkg/errors"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/rs/zerolog/log"
)

// HandleGenesis implements modules.Module
func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", "bridge").Msg("parsing genesis")

	// Read the genesis state
	var genState bridgetypes.GenesisState
	err := m.cdc.UnmarshalJSON(appState[bridgetypes.ModuleName], &genState)
	if err != nil {
		return errors.Wrap(err, "error while reading bridge genesis data")
	}

	// Save tokens
	for _, token := range genState.Tokens {
		err = m.db.SaveBridgeTokenMetadata(token.Id, token.Metadata.Name, token.Metadata.Symbol, token.Metadata.Uri)
		if err != nil {
			return errors.Wrap(err, "error while storing genesis token metadata")
		}
		for _, tokenInfo := range token.Info {
			tokenInfoId, err := m.db.SaveBridgeTokenInfo(tokenInfo.Address, tokenInfo.Decimals, tokenInfo.ChainId, tokenInfo.TokenId, tokenInfo.IsWrapped)
			if err != nil {
				return errors.Wrap(err, "error while storing genesis token info")
			}
			if err = m.db.SaveBridgeToken(tokenInfoId, token.Id, token.CommissionRate); err != nil {
				return errors.Wrap(err, "error while storing genesis token")
			}
		}
	}

	for _, chain := range genState.Chains {
		if err = m.db.SaveBridgeChain(chain.Id, int32(chain.Type), chain.BridgeAddress, chain.Operator); err != nil {
			return errors.Wrap(err, "error while storing genesis chain")
		}
	}

	for _, tx := range genState.Transactions {
		if err = m.db.SaveBridgeTransaction(tx); err != nil {
			return errors.Wrap(err, "error while storing genesis transaction")
		}
	}

	for _, txsSubmissions := range genState.TransactionsSubmissions {
		if err = m.db.SaveBridgeTransactionSubmissions(&txsSubmissions); err != nil {
			return errors.Wrap(err, "error while storing genesis transaction submissions")
		}
	}

	if err = m.db.SaveBridgeParams(&genState.Params); err != nil {
		return errors.Wrap(err, "error while storing genesis params")
	}

	return nil
}
