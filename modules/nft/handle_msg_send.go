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

	// Update the nft by setting a new owner
	err := m.db.SaveNFT(nft.Address, nft.Owner, nft.AvailableToWithdraw, nft.LastVestingBlock, nft.VestingPeriodsCount, nft.RewardPerPeriod, nft.VestingPeriodsCount, nft.Denom)
	if err != nil {
		return errors.Wrap(err, "error while saving nft")
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
