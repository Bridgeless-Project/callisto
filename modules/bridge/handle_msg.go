package bridge

import (
	"github.com/cosmos/cosmos-sdk/x/authz"
	bridge "github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
	"github.com/rs/zerolog/log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v4/types"
)

func (m *Module) HandleMsgExec(index int, _ *authz.MsgExec, _ int, executedMsg sdk.Msg, tx *juno.Tx) error {
	return m.HandleMsg(index, executedMsg, tx)
}

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(_ int, msg sdk.Msg, tx *juno.Tx) error {
	log.Debug().Str("module", "nft").Msg("handle msg")

	switch cosmosMsg := msg.(type) {
	case *bridge.MsgSubmitTransactions:
		return m.handleMsgSubmitTransactions(tx, cosmosMsg)
		// chains
	case *bridge.MsgDeleteChain:
		return m.handleMsgDeleteChain(tx, cosmosMsg)
	case *bridge.MsgInsertChain:
		return m.handleMsgInsertChain(tx, cosmosMsg)

		// token info
	case *bridge.MsgAddTokenInfo:
		return m.handleMsgAddTokenInfo(tx, cosmosMsg)
	case *bridge.MsgRemoveTokenInfo:
		return m.handleMsgRemoveTokenInfo(tx, cosmosMsg)

		// token
	case *bridge.MsgUpdateToken:
		return m.handleMsgUpdateToken(tx, cosmosMsg)
	case *bridge.MsgDeleteToken:
		return m.handleMsgDeleteToken(tx, cosmosMsg)
	case *bridge.MsgInsertToken:
		return m.handleMsgInsertToken(tx, cosmosMsg)

	default:
		break
	}

	return nil
}
