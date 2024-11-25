package bridge

import (
	"github.com/cosmos/cosmos-sdk/x/authz"
	bridge "github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
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
	case *bridge.MsgSubmitTransactions:
		return errors.Wrap(m.handleMsgSubmitBridgeTransactions(tx, cosmosMsg), "failed to handle msg submit transactions")

		// chains
	case *bridge.MsgDeleteChain:
		return errors.Wrap(m.handleMsgDeleteChain(tx, cosmosMsg), "failed to handle msg delete chain")
	case *bridge.MsgInsertChain:
		return errors.Wrap(m.handleMsgInsertChain(tx, cosmosMsg), "failed to handle msg insert chain")

		// token info
	case *bridge.MsgAddTokenInfo:
		return errors.Wrap(m.handleMsgAddTokenInfo(tx, cosmosMsg), "failed to handle msg add token info")

		// token
	case *bridge.MsgUpdateToken:
		return errors.Wrap(m.handleMsgUpdateToken(tx, cosmosMsg), "failed to handle msg update token")
	case *bridge.MsgDeleteToken:
		return errors.Wrap(m.handleMsgDeleteToken(tx, cosmosMsg), "failed to handle msg delete token")
	case *bridge.MsgInsertToken:
		return errors.Wrap(m.handleMsgInsertToken(tx, cosmosMsg), "failed to handle msg insert token")

	default:
		break
	}

	return nil
}
