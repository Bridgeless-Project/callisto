package bridge

import (
	"cosmossdk.io/errors"
	bridge "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	juno "github.com/forbole/juno/v4/types"
)

// handleMsgAddTokenInfo allows to properly handle a MsgAddTokenInfo
func (m *Module) handleMsgAddTokenInfo(_ *juno.Tx, msg *bridge.MsgAddTokenInfo) error {
	if _, err := m.db.SaveBridgeTokenInfo(
		msg.Info.Address,
		msg.Info.Decimals,
		msg.Info.ChainId,
		msg.Info.TokenId,
		msg.Info.IsWrapped,
		msg.Info.CommissionRate,
	); err != nil {
		return errors.Wrap(err, "failed to save bridge token info")
	}

	return nil
}

// handleMsgRemoveTokenInfo allows to properly handle a MsgRemoveTokenInfo
func (m *Module) handleMsgRemoveTokenInfo(_ *juno.Tx, msg *bridge.MsgRemoveTokenInfo) error {
	if err := m.db.RemoveBridgeTokenInfo(msg.TokenId); err != nil {
		return errors.Wrap(err, "failed to remove bridge token info")
	}

	return nil
}
