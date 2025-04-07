package bridge

import (
	"cosmossdk.io/errors"
	juno "github.com/forbole/juno/v4/types"
	bridge "github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
)

// handleMsgAddTokenInfo allows to properly handle a MsgAddTokenInfo
func (m *Module) handleMsgAddTokenInfo(_ *juno.Tx, msg *bridge.MsgAddTokenInfo) error {
	if _, err := m.db.SaveBridgeTokenInfo(
		msg.Info.Address,
		msg.Info.Decimals,
		msg.Info.ChainId,
		msg.Info.TokenId,
		msg.Info.IsWrapped,
	); err != nil {
		return errors.Wrap(err, "failed to save bridge token info")
	}

	return nil
}
