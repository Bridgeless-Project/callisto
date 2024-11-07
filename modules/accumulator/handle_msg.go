package accumulator

import (
	"github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/rs/zerolog/log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	accumulator "github.com/cosmos/cosmos-sdk/x/accumulator/types"
	juno "github.com/forbole/juno/v4/types"
)

// HandleMsgExec implements modules.AuthzMessageModule
func (m *Module) HandleMsgExec(index int, _ *authz.MsgExec, _ int, executedMsg sdk.Msg, tx *juno.Tx) error {
	return m.HandleMsg(index, executedMsg, tx)
}

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(_ int, msg sdk.Msg, tx *juno.Tx) error {
	log.Debug().Str("module", "accumulator").Msg("handle msg")

	switch cosmosMsg := msg.(type) {
	case *accumulator.MsgAddAdmin:
		return m.handleMsgAddAdmin(tx, cosmosMsg)
	default:
		break
	}

	return nil
}

// handleMsgAddAdmin save a new admin to db
func (m *Module) handleMsgAddAdmin(_ *juno.Tx, msg *accumulator.MsgAddAdmin) error {
	return m.db.SaveAdmin(
		msg.Address,
		msg.VestingPeriodsCount,
		0,
		msg.VestingPeriod,
		msg.RewardPerPeriod,
		msg.Denom,
	)
}
