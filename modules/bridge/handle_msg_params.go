package bridge

import (
	"fmt"
	juno "github.com/forbole/juno/v4/types"
	bridge "github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
)

func (m *Module) handleMsgSetParties(_ *juno.Tx, msg *bridge.MsgSetParties) error {
	params, err := m.db.GetBridgeParams()
	if err != nil {
		return fmt.Errorf("handleMsgSetParties: get params: %w", err)
	}

	params.Parties = msg.Parties

	if err = m.db.SaveBridgeParams(params); err != nil {
		return fmt.Errorf("handleMsgSetParties: save params: %w", err)
	}

	return nil
}

func (m *Module) handleMsgSetTssThreshold(_ *juno.Tx, msg *bridge.MsgSetTssThreshold) error {
	params, err := m.db.GetBridgeParams()
	if err != nil {
		return fmt.Errorf("handleMsgSetTssThreshold: get params: %w", err)
	}

	params.TssThreshold = msg.Threshold

	if err = m.db.SaveBridgeParams(params); err != nil {
		return fmt.Errorf("handleMsgSetTssThreshold: save params: %w", err)
	}
	return nil
}
