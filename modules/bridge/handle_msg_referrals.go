package bridge

import (
	"cosmossdk.io/errors"
	bridge "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	juno "github.com/forbole/juno/v4/types"
)

// handleMsgSetReferral allows to properly handle a MsgAddReferral
func (m *Module) handleMsgSetReferral(_ *juno.Tx, msg *bridge.MsgSetReferral) error {
	if err := m.db.SaveBridgeReferral(
		&msg.Referral,
	); err != nil {
		return errors.Wrap(err, "failed to save bridge referral")
	}

	return nil
}

// handleMsgRemoveReferral allows to properly handle a MsgRemoveReferral
func (m *Module) handleMsgRemoveReferral(_ *juno.Tx, msg *bridge.MsgRemoveReferral) error {
	if err := m.db.RemoveBridgeReferral(msg.Id); err != nil {
		return errors.Wrap(err, "failed to remove bridge token info")
	}

	return nil
}
