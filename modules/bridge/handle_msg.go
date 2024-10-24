package bridge

import (
	"github.com/cosmos/cosmos-sdk/x/authz"
	bridge "github.com/hyle-team/bridgeless-core/x/bridge/types"
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
		return nil
		// chains
	case *bridge.MsgDeleteChain:
		return nil
	case *bridge.MsgInsertChain:
		return nil

		// token info
	case *bridge.MsgAddTokenInfo:
		return nil
	case *bridge.MsgRemoveTokenInfo:
		return nil

		// token
	case *bridge.MsgUpdateToken:
		return nil
	case *bridge.MsgDeleteToken:
		return nil
	case *bridge.MsgInsertToken:
		return nil

	default:
		break
	}

	return nil
}
