package nft

import (
	nft "github.com/cosmos/cosmos-sdk/x/nft/types"
)

// handleMsgSend allows to properly handle a MsgSend
func (m *Module) handleMsgWithdrawal(msg *nft.MsgWithdrawal) error {
	return m.db.SaveNFTEvent(
		msg.Type(),
		msg.Address,
		"",
		"",
		msg.Creator,
		msg.Creator,
		msg.Amount,
	)
}
