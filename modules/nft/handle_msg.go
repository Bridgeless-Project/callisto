package nft

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/authz"

	sdk "github.com/cosmos/cosmos-sdk/types"
	nft "github.com/cosmos/cosmos-sdk/x/nft/types"
	juno "github.com/forbole/juno/v4/types"
)

// HandleMsgExec implements modules.AuthzMessageModule
func (m *Module) HandleMsgExec(_ int, _ *authz.MsgExec, _ int, executedMsg sdk.Msg, tx *juno.Tx) error {
	return m.HandleMsg(executedMsg, tx)
}

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(msg sdk.Msg, tx *juno.Tx) error {
	if len(tx.Logs) == 0 {
		return nil
	}

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
		return fmt.Errorf("nft does not exist")
	}

	return m.db.SaveNFTEvent(
		msg.Type(),
		nft.Address,
		"",
		msg.Validator,
		nft.Owner,
		nft.Owner,
		sdk.NewCoins(
			msg.Amount,
		),
	)
}

// handleMsgRedelegate allows to properly handle a MsgRedelegate
func (m *Module) handleMsgRedelegate(tx *juno.Tx, msg *nft.MsgRedelegate) error {
	nft, ok := m.keeper.GetNFT(msg.Address, tx.Height)
	if !ok {
		return fmt.Errorf("nft does not exist")
	}

	return m.db.SaveNFTEvent(
		msg.Type(),
		nft.Address,
		msg.ValidatorSrc,
		msg.ValidatorNew,
		nft.Owner,
		nft.Owner,
		sdk.NewCoins(
			msg.Amount,
		),
	)
}

// handleMsgUndelegate allows to properly handle a MsgUndelegate
func (m *Module) handleMsgUndelegate(tx *juno.Tx, msg *nft.MsgUndelegate) error {
	nft, ok := m.keeper.GetNFT(msg.Address, tx.Height)
	if !ok {
		return fmt.Errorf("nft does not exist")
	}

	// TODO validate the undelegation

	return m.db.SaveNFTEvent(
		msg.Type(),
		nft.Address,
		msg.Validator,
		"",
		nft.Owner,
		nft.Owner,
		sdk.NewCoins(
			msg.Amount,
		),
	)
}

// handleMsgSend allows to properly handle a MsgSend
func (m *Module) handleMsgSend(tx *juno.Tx, msg *nft.MsgSend) error {
	nft, ok := m.keeper.GetNFT(msg.Address, tx.Height)
	if !ok {
		return fmt.Errorf("nft does not exist")
	}

	return m.db.SaveNFTEvent(
		msg.Type(),
		nft.Address,
		"",
		"",
		msg.Creator,
		msg.Recipient,
		sdk.NewCoins(
			sdk.NewCoin(
				nft.Denom,
				sdk.ZeroInt(),
			),
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
		sdk.NewCoins(
			msg.Amount,
		),
	)
}
