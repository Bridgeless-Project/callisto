package bridge

import (
	juno "github.com/forbole/juno/v4/types"
	bridge "github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
)

// handleMsgSubmitTransactions allows to properly handle a MsgSubmitTransactions
func (m *Module) handleMsgSubmitTransactions(_ *juno.Tx, msg *bridge.MsgSubmitTransactions) error {

	for _, tx := range msg.Transactions {
		if err := m.db.SaveBridgeTransaction(tx); err != nil {
			return err
		}
	}
	return nil
}
