package nft

import (
	nft "github.com/cosmos/cosmos-sdk/x/nft/types"
	juno "github.com/forbole/juno/v4/types"
	"github.com/pkg/errors"
)

// handleMsgDeposit allows to properly handle a handleMsgDeposit
func (m *Module) handleMsgDelegate(tx *juno.Tx, msg *nft.MsgDelegate) error {
	nft, ok := m.keeper.GetNFT(msg.Address, tx.Height)
	if !ok {
		return errors.New("nft does not exist")
	}

	return m.db.SaveNFTEvent(
		msg.Type(),
		nft.Address,
		"",
		msg.Validator,
		nft.Owner,
		nft.Owner,
		msg.Amount,
	)
}

// handleMsgRedelegate allows to properly handle a MsgRedelegate
func (m *Module) handleMsgRedelegate(tx *juno.Tx, msg *nft.MsgRedelegate) error {
	nft, ok := m.keeper.GetNFT(msg.Address, tx.Height)
	if !ok {
		return errors.New("nft does not exist")
	}

	return m.db.SaveNFTEvent(
		msg.Type(),
		nft.Address,
		msg.ValidatorSrc,
		msg.ValidatorNew,
		nft.Owner,
		nft.Owner,
		msg.Amount,
	)
}

// handleMsgUndelegate allows to properly handle a MsgUndelegate
func (m *Module) handleMsgUndelegate(tx *juno.Tx, msg *nft.MsgUndelegate) error {
	nft, ok := m.keeper.GetNFT(msg.Address, tx.Height)
	if !ok {
		return errors.New("nft does not exist")
	}

	delegation, err := m.stakingModule.GetDelegationByValidator(tx.Height, nft.Owner, msg.Validator)
	if err != nil {
		return errors.Wrap(err, "get delegation by validator")
	}

	newValidator := msg.Validator
	if delegation == nil {
		newValidator = ""
	}
	return m.db.SaveNFTEvent(
		msg.Type(),
		nft.Address,
		msg.Validator,
		newValidator,
		nft.Owner,
		nft.Owner,
		msg.Amount,
	)
}
