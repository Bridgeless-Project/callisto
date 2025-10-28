package bridge

import (
	"cosmossdk.io/errors"
	bridge "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	juno "github.com/forbole/juno/v4/types"
)

// handleMsgSetReferralRewards allows to properly handle a MsgAddReferralRewards
func (m *Module) handleMsgSetReferralRewards(_ *juno.Tx, msg *bridge.MsgSetReferralRewards) error {
	if err := m.db.SaveBridgeReferralRewards(
		&msg.Rewards,
	); err != nil {
		return errors.Wrap(err, "failed to save bridge ReferralRewards")
	}

	return nil
}

// handleMsgRemoveReferralRewards allows to properly handle a MsgRemoveReferralRewards
func (m *Module) handleMsgRemoveReferralRewards(_ *juno.Tx, msg *bridge.MsgRemoveReferralRewards) error {
	if err := m.db.RemoveBridgeReferralRewards(msg.ReferrerId, msg.TokenId); err != nil {
		return errors.Wrap(err, "failed to remove bridge token info")
	}

	return nil
}
