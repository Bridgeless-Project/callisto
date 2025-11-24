package bridge

import (
	bridge "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	v19 "github.com/forbole/bdjuno/v4/modules/bridge/migrations/v19"
	v21 "github.com/forbole/bdjuno/v4/modules/bridge/migrations/v21"
	v24 "github.com/forbole/bdjuno/v4/modules/bridge/migrations/v24"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v4/types"
)

func (m *Module) HandleMsgExec(index int, _ *authz.MsgExec, _ int, executedMsg sdk.Msg, tx *juno.Tx) error {
	return m.HandleMsg(index, executedMsg, tx)
}

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(_ int, msg sdk.Msg, tx *juno.Tx) error {
	logger := log.Debug().Str("module", "bridge")
	logger.Msg("handle msg")

	switch cosmosMsg := msg.(type) {
	// transactions
	case *bridge.MsgSubmitTransactions:
		return errors.Wrap(m.handleMsgSubmitBridgeTransactions(tx, cosmosMsg), "failed to handle msg submit transactions")
	case *bridge.MsgRemoveTransaction:
		return errors.Wrap(m.handleMsgRemoveTransaction(tx, cosmosMsg), "failed to handle msg remove transaction")
	case *v19.MsgSubmitTransactions:
		return errors.Wrap(m.handleMsgSubmitBridgeTransactions(tx, V19SubmitTxsToLatest(cosmosMsg)), "failed to handle msg submit transactions")

	// chains
	case *bridge.MsgDeleteChain:
		return errors.Wrap(m.handleMsgDeleteChain(tx, cosmosMsg), "failed to handle msg delete chain")
	case *bridge.MsgInsertChain:
		return errors.Wrap(m.handleMsgInsertChain(tx, cosmosMsg), "failed to handle msg insert chain")
	case *v19.MsgInsertChain:
		return errors.Wrap(m.handleMsgInsertChain(tx, V19InsertChainToLatest(cosmosMsg)), "failed to handle msg insert chain")

	// token info
	case *bridge.MsgAddTokenInfo:
		return errors.Wrap(m.handleMsgAddTokenInfo(tx, cosmosMsg), "failed to handle msg add token info")
	case *bridge.MsgRemoveTokenInfo:
		return errors.Wrap(m.handleMsgRemoveTokenInfo(tx, cosmosMsg), "failed to handle msg remove token info")
	case *v19.MsgAddTokenInfo:
		return errors.Wrap(m.handleMsgAddTokenInfo(tx, V19AddTokenInfoToLatest(cosmosMsg)), "failed to handle msg add token info")
	case *v21.MsgAddTokenInfo:
		return errors.Wrap(m.handleMsgAddTokenInfo(tx, V21AddTokenInfoToLatest(cosmosMsg)), "failed to handle msg add token info")
	case *v24.MsgAddTokenInfo:
		return errors.Wrap(m.handleMsgAddTokenInfo(tx, V24AddTokenInfoToLatest(cosmosMsg)), "failed to handle msg add token info")

	// token
	case *bridge.MsgUpdateToken:
		return errors.Wrap(m.handleMsgUpdateToken(tx, cosmosMsg), "failed to handle msg update token")
	case *bridge.MsgDeleteToken:
		return errors.Wrap(m.handleMsgDeleteToken(tx, cosmosMsg), "failed to handle msg delete token")
	case *bridge.MsgInsertToken:
		return errors.Wrap(m.handleMsgInsertToken(tx, cosmosMsg), "failed to handle msg insert token")
	case *v19.MsgInsertToken:
		return errors.Wrap(m.handleMsgInsertToken(tx, V19InsertTokenToLatest(cosmosMsg)), "failed to handle msg insert token")
	case *v19.MsgUpdateToken:
		return errors.Wrap(m.handleMsgUpdateToken(tx, V19UpdateTokenToLatest(cosmosMsg)), "failed to handle msg update token")
	case *v21.MsgInsertToken:
		return errors.Wrap(m.handleMsgInsertToken(tx, V21InsertTokenToLatest(cosmosMsg)), "failed to handle msg insert token")
	case *v24.MsgInsertToken:
		return errors.Wrap(m.handleMsgInsertToken(tx, V24InsertTokenToLatest(cosmosMsg)), "failed to handle msg insert token")

	// parties
	case *bridge.MsgSetParties:
		return errors.Wrap(m.handleMsgSetParties(tx, cosmosMsg), "failed to handle msg set parties")
	case *bridge.MsgSetTssThreshold:
		return errors.Wrap(m.handleMsgSetTssThreshold(tx, cosmosMsg), "failed to handle msg set tss threshold")

	// referrals
	case *bridge.MsgSetReferral:
		return errors.Wrap(m.handleMsgSetReferral(tx, cosmosMsg), "failed to handle msg set referral")
	case *bridge.MsgRemoveReferral:
		return errors.Wrap(m.handleMsgRemoveReferral(tx, cosmosMsg), "failed to handle msg remove referral")

	// referrals rewards
	case *bridge.MsgSetReferralRewards:
		return errors.Wrap(m.handleMsgSetReferralRewards(tx, cosmosMsg), "failed to handle msg set referral rewards")
	case *bridge.MsgRemoveReferralRewards:
		return errors.Wrap(m.handleMsgRemoveReferralRewards(tx, cosmosMsg), "failed to handle msg remove referral rewards")

	default:
		log.Error().Msgf("can not parse unknown msg: %#v", msg)
	}

	return nil
}
