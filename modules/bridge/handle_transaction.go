package bridge

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
	juno "github.com/forbole/juno/v4/types"
	bridge "github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
)

// handleMsgSubmitBridgeTransactions allows to properly handle a MsgSubmitTransactions
func (m *Module) handleMsgSubmitBridgeTransactions(_ *juno.Tx, msg *bridge.MsgSubmitTransactions) error {
	for _, tx := range msg.Transactions {
		if err := m.db.SaveBridgeTransaction(tx); err != nil {
			return errors.Wrap(err, "failed to save bridge transaction")
		}
	}
	
	return nil
}
