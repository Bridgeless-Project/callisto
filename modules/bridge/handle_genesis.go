package bridge

import (
	"encoding/json"
	"time"

	bridgetypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/pkg/errors"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/rs/zerolog/log"
)

// HandleGenesis implements modules.Module
func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", "bridge").Msg("parsing genesis")

	// Read the genesis state
	var (
		genState    bridgetypes.GenesisState
		exists      bool
		err         error
		tokenInfoId int64
	)

	err = m.cdc.UnmarshalJSON(appState[bridgetypes.ModuleName], &genState)
	if err != nil {
		return errors.Wrap(err, "error while reading bridge genesis data")
	}

	// Save tokens
	for _, token := range genState.Tokens {
		exists, err = m.tokenMetadataExists(token.Id)
		if err != nil {
			return errors.Wrap(err, "error while checking if token metadata exists")
		}

		if !exists {
			err = m.db.SaveBridgeTokenMetadata(token.Id, token.Metadata.Name, token.Metadata.Symbol, token.Metadata.Uri)
			if err != nil {
				return errors.Wrap(err, "error while storing genesis token metadata")
			}
		}

		for _, tokenInfo := range token.Info {
			tokenInfoId, exists, err = m.tokenInfoExists(tokenInfo.Address, tokenInfo.ChainId)
			if err != nil {
				return errors.Wrap(err, "failed to check token info existence")
			}

			if !exists {
				tokenInfoId, err = m.db.SaveBridgeTokenInfo(tokenInfo.Address, tokenInfo.Decimals, tokenInfo.ChainId, tokenInfo.TokenId, tokenInfo.IsWrapped, tokenInfo.MinWithdrawalAmount, tokenInfo.CommissionRate)
				if err != nil {
					return errors.Wrap(err, "error while storing genesis token info")
				}
			}

			exists, err = m.tokenExists(uint64(tokenInfoId), token.Id)
			if err != nil {
				return errors.Wrap(err, "failed to check existence of token")
			}

			if exists {
				continue
			}

			if err = m.db.SaveBridgeToken(tokenInfoId, token.Id); err != nil {
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
		if err = m.db.SaveBridgeTransaction(tx, doc.GenesisTime.Format(time.RFC3339Nano)); err != nil {
			return errors.Wrap(err, "error while storing genesis transaction")
		}

		tokenId, err := m.db.SetBridgeTransactionTokenId(tx)
		if err != nil {
			return errors.Wrap(err, "failed to set bridge transaction token id")
		}

		depositDecimals, withdrawalDecimals, err := m.db.SetBridgeTransactionDecimals(tx)
		if err != nil {
			return errors.Wrap(err, "failed to set bridge transaction decimals")
		}

		if err := m.UpdateTokenVolume(&tx, tokenId, depositDecimals, withdrawalDecimals,
			doc.GenesisTime.Format(time.RFC3339Nano)); err != nil {
			return errors.Wrap(err, "failed to update token volume")
		}
	}

	for _, txsSubmissions := range genState.TransactionsSubmissions {
		if err = m.db.SaveBridgeTransactionSubmissions(&txsSubmissions); err != nil {
			return errors.Wrap(err, "error while storing genesis transaction submissions")
		}
	}

	for _, referral := range genState.Referrals {
		if err = m.db.SaveBridgeReferral(&referral); err != nil {
			return errors.Wrap(err, "error while storing genesis tss party")
		}
	}

	for _, referralRewards := range genState.ReferralsRewards {
		if err = m.db.SaveBridgeReferralRewards(&referralRewards); err != nil {
			return errors.Wrap(err, "error while storing genesis referral rewards")
		}
	}

	if err = m.db.SaveBridgeParams(&genState.Params); err != nil {
		return errors.Wrap(err, "error while storing genesis params")
	}

	return nil
}
