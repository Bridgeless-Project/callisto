package bridge

import (
	bridge "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
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

	// chains
	case *bridge.MsgDeleteChain:
		return errors.Wrap(m.handleMsgDeleteChain(tx, cosmosMsg), "failed to handle msg delete chain")
	case *bridge.MsgInsertChain:
		return errors.Wrap(m.handleMsgInsertChain(tx, cosmosMsg), "failed to handle msg insert chain")

	// token info
	case *bridge.MsgAddTokenInfo:
		return errors.Wrap(m.handleMsgAddTokenInfo(tx, cosmosMsg), "failed to handle msg add token info")
	case *bridge.MsgRemoveTokenInfo:
		return errors.Wrap(m.handleMsgRemoveTokenInfo(tx, cosmosMsg), "failed to handle msg remove token info")

	// token
	case *bridge.MsgUpdateToken:
		return errors.Wrap(m.handleMsgUpdateToken(tx, cosmosMsg), "failed to handle msg update token")
	case *bridge.MsgDeleteToken:
		return errors.Wrap(m.handleMsgDeleteToken(tx, cosmosMsg), "failed to handle msg delete token")
	case *bridge.MsgInsertToken:
		return errors.Wrap(m.handleMsgInsertToken(tx, cosmosMsg), "failed to handle msg insert token")
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
		break
	}

	return nil
}
