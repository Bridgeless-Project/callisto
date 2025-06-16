package bridge

import (
	bridge "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	juno "github.com/forbole/juno/v4/types"
	"github.com/pkg/errors"
)

// handleMsgInsertToken allows to properly handle a MsgInsertToken
func (m *Module) handleMsgInsertToken(_ *juno.Tx, msg *bridge.MsgInsertToken) error {
	if err := m.db.SaveBridgeTokenMetadata(
		msg.Token.Id,
		msg.Token.Metadata.Name,
		msg.Token.Metadata.Symbol,
		msg.Token.Metadata.Uri,
	); err != nil {
		return errors.Wrap(err, "failed to save bridge token metadata")
	}

	for _, tokenInfo := range msg.Token.Info {
		tokenInfoId, err := m.db.SaveBridgeTokenInfo(
			tokenInfo.Address,
			tokenInfo.Decimals,
			tokenInfo.ChainId,
			tokenInfo.TokenId,
			tokenInfo.IsWrapped,
		)
		if err != nil {
			return errors.Wrap(err, "failed to save bridge token info")
		}

		if err = m.db.SaveBridgeToken(tokenInfoId, msg.Token.Id, msg.Token.CommissionRate); err != nil {
			return errors.Wrap(err, "failed to save bridge token")
		}
	}

	return nil
}

// handleMsgDeleteToken allows to properly handle a MsgDeleteToken
func (m *Module) handleMsgDeleteToken(_ *juno.Tx, msg *bridge.MsgDeleteToken) error {
	return errors.Wrap(m.db.RemoveBridgeToken(msg.TokenId), "failed to remove bridge token")
}

// handleMsgUpdateToken allows to properly handle a MsgUpdateToken
func (m *Module) handleMsgUpdateToken(_ *juno.Tx, msg *bridge.MsgUpdateToken) error {
	if err := m.db.SaveBridgeTokenMetadata(
		msg.TokenId,
		msg.Metadata.Name,
		msg.Metadata.Symbol,
		msg.Metadata.Uri,
	); err != nil {
		return errors.Wrap(err, "failed to save bridge token metadata")
	}

	return nil
}
