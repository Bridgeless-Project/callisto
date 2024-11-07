package nft

import (
	"github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/rs/zerolog/log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	nft "github.com/cosmos/cosmos-sdk/x/nft/types"
	juno "github.com/forbole/juno/v4/types"
)

// HandleMsgExec implements modules.AuthzMessageModule
func (m *Module) HandleMsgExec(index int, _ *authz.MsgExec, _ int, executedMsg sdk.Msg, tx *juno.Tx) error {
	return m.HandleMsg(index, executedMsg, tx)
}

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(_ int, msg sdk.Msg, tx *juno.Tx) error {
	log.Debug().Str("module", "nft").Msg("handle msg")

	switch cosmosMsg := msg.(type) {
	case *nft.MsgDelegate:
		return m.handleMsgDelegate(tx, cosmosMsg)
	case *nft.MsgRedelegate:
		return m.handleMsgRedelegate(tx, cosmosMsg)
	case *nft.MsgUndelegate:
		return m.handleMsgUndelegate(tx, cosmosMsg)
	case *nft.MsgSend:
		return m.handleMsgSend(tx, cosmosMsg)
	case *nft.MsgWithdrawal:
		return m.handleMsgWithdrawal(tx, cosmosMsg)
	default:
		break
	}

	return nil
}
