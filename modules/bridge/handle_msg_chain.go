package bridge

import (
	juno "github.com/forbole/juno/v4/types"
	bridge "github.com/hyle-team/bridgeless-core/x/bridge/types"
)

// handleMsgInsertChain allows to properly handle a MsgRemoveTokenInfo
func (m *Module) handleMsgInsertChain(tx *juno.Tx, msg *bridge.MsgInsertChain) error {
	return nil
}

// handleMsgDeleteChain allows to properly handle a MsgDeleteChain
func (m *Module) handleMsgDeleteChain(tx *juno.Tx, msg *bridge.MsgDeleteChain) error {
	return nil
}
