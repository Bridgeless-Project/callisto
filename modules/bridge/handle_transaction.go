package bridge

import (
	juno "github.com/forbole/juno/v4/types"
	bridge "github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
)

// handleMsgSubmitTransactions allows to properly handle a MsgSubmitTransactions
func (m *Module) handleMsgSubmitTransactions(tx *juno.Tx, msg *bridge.MsgSubmitTransactions) error {
	return nil
}
