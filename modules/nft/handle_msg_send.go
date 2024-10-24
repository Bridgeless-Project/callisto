package nft

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	nft "github.com/cosmos/cosmos-sdk/x/nft/types"
	juno "github.com/forbole/juno/v4/types"
	"github.com/pkg/errors"
)

// handleMsgSend allows to properly handle a MsgSend
func (m *Module) handleMsgSend(tx *juno.Tx, msg *nft.MsgSend) error {
	nft, ok := m.keeper.GetNFT(msg.Address, tx.Height)
	if !ok {
		return errors.New("nft does not exist")
	}

	return m.db.SaveNFTEvent(
		msg.Type(),
		nft.Address,
		"",
		"",
		msg.Creator,
		msg.Recipient,
		sdk.NewCoin(
			nft.Denom,
			sdk.ZeroInt(),
		),
	)
}
