package bridge

import (
	juno "github.com/forbole/juno/v4/types"
	bridge "github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
)

// handleMsgInsertChain allows to properly handle a MsgInsertChain
func (m *Module) handleMsgInsertChain(_ *juno.Tx, msg *bridge.MsgInsertChain) error {
	return m.db.SaveBridgeChain(
		msg.Chain.Id,
		int32(msg.Chain.Type),
		msg.Chain.BridgeAddress,
		msg.Chain.Operator,
	)
}

// handleMsgDeleteChain allows to properly handle a MsgDeleteChain
func (m *Module) handleMsgDeleteChain(_ *juno.Tx, msg *bridge.MsgDeleteChain) error {
	return m.db.RemoveBridgeChain(msg.ChainId)
}
