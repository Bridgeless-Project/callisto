package bridge

import (
	juno "github.com/forbole/juno/v4/types"
	bridge "github.com/hyle-team/bridgeless-core/x/bridge/types"
)

// handleMsgAddTokenInfo allows to properly handle a MsgAddTokenInfo
func (m *Module) handleMsgAddTokenInfo(tx *juno.Tx, msg *bridge.MsgAddTokenInfo) error {
	return nil
}

// handleMsgRemoveTokenInfo allows to properly handle a MsgRemoveTokenInfo
func (m *Module) handleMsgRemoveTokenInfo(tx *juno.Tx, msg *bridge.MsgRemoveTokenInfo) error {
	return nil
}
