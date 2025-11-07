package bridge

import (
	bridge "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	juno "github.com/forbole/juno/v4/types"
	"github.com/pkg/errors"
)

// handleMsgInsertToken allows to properly handle a MsgInsertToken
func (m *Module) handleMsgInsertToken(_ *juno.Tx, msg *bridge.MsgInsertToken) error {
	var (
		exists      bool
		err         error
		tokenInfoId int64
	)

	exists, err = m.tokenMetadataExists(msg.Token.Id)
	if err != nil {
		return errors.Wrap(err, "failed to check existence of token metadata")
	}

	if !exists {
		// save the token metadata only if it not exists in db
		if err = m.db.SaveBridgeTokenMetadata(
			msg.Token.Id,
			msg.Token.Metadata.Name,
			msg.Token.Metadata.Symbol,
			msg.Token.Metadata.Uri,
		); err != nil {
			return errors.Wrap(err, "failed to save bridge token metadata")
		}
	}

	for _, tokenInfo := range msg.Token.Info {
		tokenInfoId, exists, err = m.tokenInfoExists(tokenInfo.Address, tokenInfo.ChainId)
		if err != nil {
			return errors.Wrap(err, "failed to check existence of token info")
		}

		if !exists {
			tokenInfoId, err = m.db.SaveBridgeTokenInfo(
				tokenInfo.Address,
				tokenInfo.Decimals,
				tokenInfo.ChainId,
				tokenInfo.TokenId,
				tokenInfo.IsWrapped,
				tokenInfo.MinWithdrawalAmount,
				tokenInfo.CommissionRate,
			)
			if err != nil {
				return errors.Wrap(err, "failed to save bridge token info")
			}
		}

		exists, err = m.tokenExists(uint64(tokenInfoId), msg.Token.Id)
		if err != nil {
			return errors.Wrap(err, "failed to check existence of token")
		}

		if exists {
			continue
		}

		if err = m.db.SaveBridgeToken(tokenInfoId, msg.Token.Id); err != nil {
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

func (m *Module) tokenMetadataExists(tokenId uint64) (bool, error) {
	tokenMetadata, err := m.db.GetBridgeTokenMetadata(tokenId)
	if err != nil {
		return false, errors.Wrap(err, "failed to get bridge token metadata")
	}

	return tokenMetadata != nil, nil
}

func (m *Module) tokenExists(tokenId, metadataId uint64) (bool, error) {
	token, err := m.db.GetBridgeToken(tokenId, metadataId)
	if err != nil {
		return false, errors.Wrap(err, "failed to get bridge token")
	}

	return token != nil, nil
}
