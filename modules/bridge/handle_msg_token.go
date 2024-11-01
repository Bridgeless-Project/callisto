package bridge

import (
	juno "github.com/forbole/juno/v4/types"
	bridge "github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
)

// handleMsgInsertToken allows to properly handle a MsgInsertToken
func (m *Module) handleMsgInsertToken(_ *juno.Tx, msg *bridge.MsgInsertToken) error {
	err := m.db.SaveBridgeTokenMetadata(msg.Token.Id, msg.Token.Metadata.Name, msg.Token.Metadata.Symbol, msg.Token.Metadata.Uri)
	if err != nil {
		return err
	}

	for _, tokenInfo := range msg.Token.Info {
		tokenInfoId, err := m.db.SaveBridgeTokenInfo(tokenInfo.Address, tokenInfo.Decimals, tokenInfo.ChainId, tokenInfo.TokenId, tokenInfo.IsWrapped)
		if err != nil {
			return err
		}

		if err = m.db.SaveBridgeToken(tokenInfoId, msg.Token.Id); err != nil {
			return err
		}
	}

	return err
}

// handleMsgDeleteToken allows to properly handle a MsgDeleteToken
func (m *Module) handleMsgDeleteToken(_ *juno.Tx, msg *bridge.MsgDeleteToken) error {
	return m.db.RemoveBridgeToken(msg.TokenId)
}

// handleMsgUpdateToken allows to properly handle a MsgUpdateToken
func (m *Module) handleMsgUpdateToken(_ *juno.Tx, msg *bridge.MsgUpdateToken) error {
	if err := m.db.SaveBridgeTokenMetadata(msg.TokenId, msg.Metadata.Name, msg.Metadata.Symbol, msg.Metadata.Uri); err != nil {
		return err
	}

	return nil
}
