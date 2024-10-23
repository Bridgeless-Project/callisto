package nft

import (
	"github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	nft "github.com/cosmos/cosmos-sdk/x/nft/types"
	juno "github.com/forbole/juno/v4/types"
)

// HandleMsgExec implements modules.AuthzMessageModule
func (m *Module) HandleMsgExec(index int, _ *authz.MsgExec, _ int, executedMsg sdk.Msg, tx *juno.Tx) error {
	return m.HandleMsg(index, executedMsg, tx)
}

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(_ int, msg sdk.Msg, tx *juno.Tx) error {
	log.Debug().Str("module", "nft").Msg("handle msg")

	switch cosmosMsg := msg.(type) {
	case *nft.MsgDelegate:
		return m.handleMsgDelegate(tx, cosmosMsg)
	case *nft.MsgRedelegate:
		return m.handleMsgRedelegate(tx, cosmosMsg)
	case *nft.MsgUndelegate:
		return m.handleMsgUndelegate(tx, cosmosMsg)
	case *nft.MsgSend:
		return m.handleMsgSend(tx, cosmosMsg)
	case *nft.MsgWithdrawal:
		return m.handleMsgWithdrawal(cosmosMsg)
	default:
		break
	}

	return nil
}

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
