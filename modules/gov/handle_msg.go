package gov

import (
	"fmt"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/gogo/protobuf/proto"
	"strings"
	"time"

	"strconv"

	"github.com/cosmos/cosmos-sdk/x/authz"

	"github.com/forbole/bdjuno/v4/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
	govtypesv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	juno "github.com/forbole/juno/v4/types"
)

// HandleMsgExec implements modules.AuthzMessageModule
func (m *Module) HandleMsgExec(index int, _ *authz.MsgExec, _ int, executedMsg sdk.Msg, tx *juno.Tx) error {
	return m.HandleMsg(index, executedMsg, tx)
}

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	if len(tx.Logs) == 0 {
		return nil
	}

	switch cosmosMsg := msg.(type) {
	case *govtypesv1beta1.MsgSubmitProposal:
		return m.handleMsgSubmitProposal(tx, index, cosmosMsg)

	case *govtypesv1.MsgSubmitProposal:
		return m.handleMsgSubmitV1Proposal(tx, index, cosmosMsg)

	case *govtypesv1.MsgDeposit:
		return m.handleMsgDeposit(tx, cosmosMsg.Depositor, cosmosMsg.ProposalId)

	case *govtypesv1beta1.MsgDeposit:
		return m.handleMsgDeposit(tx, cosmosMsg.Depositor, cosmosMsg.ProposalId)

	case *govtypesv1.MsgVote:
		return m.handleMsgVote(tx, cosmosMsg.ProposalId, cosmosMsg.Voter, int32(cosmosMsg.Option))

	case *govtypesv1beta1.MsgVote:
		return m.handleMsgVote(tx, cosmosMsg.ProposalId, cosmosMsg.Voter, int32(cosmosMsg.Option))
	}

	return nil
}

// handleMsgSubmitProposal allows to properly handle a handleMsgSubmitProposal
func (m *Module) handleMsgSubmitProposal(tx *juno.Tx, index int, msg *govtypesv1beta1.MsgSubmitProposal) error {
	// Get the proposal id
	event, err := tx.FindEventByType(index, gov.EventTypeSubmitProposal)
	if err != nil {
		return fmt.Errorf("error while searching for EventTypeSubmitProposal: %s", err)
	}

	id, err := tx.FindAttributeByKey(event, gov.AttributeKeyProposalID)
	if err != nil {
		return fmt.Errorf("error while searching for AttributeKeyProposalID: %s", err)
	}

	proposalID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return fmt.Errorf("error while parsing proposal id: %s", err)
	}

	// Get the proposal
	proposal, err := m.source.Proposal(tx.Height, proposalID)
	if err != nil {
		return fmt.Errorf("error while getting proposal: %s", err)
	}

	// Unpack the content
	var content govtypesv1beta1.Content
	err = m.cdc.UnpackAny(proposal.Content, &content)
	if err != nil {
		return fmt.Errorf("error while unpacking proposal content: %s", err)
	}

	// Encode the content properly
	protoContent, ok := content.(proto.Message)
	if !ok {
		return fmt.Errorf("invalid proposal content types: %T", proposal.Content)
	}

	anyContent, err := codectypes.NewAnyWithValue(protoContent)
	if err != nil {
		return fmt.Errorf("error while wrapping proposal proto content: %s", err)
	}

	contentBz, err := m.db.EncodingConfig.Codec.MarshalJSON(anyContent)
	if err != nil {
		return fmt.Errorf("error while marshaling proposal content: %s", err)
	}

	// Store the proposal
	proposalObj := types.NewProposal(
		proposal.ProposalId,
		msg.GetContent().ProposalRoute(),
		msg.GetContent().ProposalType(),
		msg.GetContent().GetTitle(),
		msg.GetContent().GetDescription(),
		string(contentBz),
		proposal.Status.String(),
		proposal.SubmitTime,
		proposal.DepositEndTime,
		proposal.VotingStartTime,
		proposal.VotingEndTime,
		msg.Proposer,
		"",
	)

	err = m.db.SaveProposals([]types.Proposal{proposalObj})
	if err != nil {
		return err
	}

	txTimestamp, err := time.Parse(time.RFC3339, tx.Timestamp)
	if err != nil {
		return fmt.Errorf("error while parsing time: %s", err)
	}

	// Store the deposit
	deposit := types.NewDeposit(proposal.ProposalId, msg.Proposer, msg.InitialDeposit, txTimestamp, tx.Height)
	return m.db.SaveDeposits([]types.Deposit{deposit})
}

// handleMsgSubmitProposal allows to properly handle a handleMsgSubmitProposal
func (m *Module) handleMsgSubmitV1Proposal(tx *juno.Tx, index int, msg *govtypesv1.MsgSubmitProposal) error {
	// Get the proposal id
	event, err := tx.FindEventByType(index, gov.EventTypeSubmitProposal)
	if err != nil {
		return fmt.Errorf("error while searching for EventTypeSubmitProposal: %s", err)
	}

	id, err := tx.FindAttributeByKey(event, gov.AttributeKeyProposalID)
	if err != nil {
		return fmt.Errorf("error while searching for AttributeKeyProposalID: %s", err)
	}

	proposalID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return fmt.Errorf("error while parsing proposal id: %s", err)
	}

	// Get the proposal
	proposal, err := m.source.Proposal(tx.Height, proposalID)
	if err != nil {
		return fmt.Errorf("error while getting proposal: %s", err)
	}

	var sb strings.Builder

	for i, proposalMsg := range msg.Messages {
		if i == 0 {
			sb.WriteString("[")
		}

		bz, err := m.db.EncodingConfig.Codec.MarshalJSON(proposalMsg)
		if err != nil {
			return fmt.Errorf("error while marshaling proposal content: %s", err)
		}

		sb.Write(bz)

		if (i + 1) == len(msg.Messages) {
			sb.WriteString("]")
		} else {
			sb.WriteString(",")
		}
	}

	// Store the proposal
	proposalObj := types.NewProposal(
		proposal.ProposalId,
		"",
		"",
		"",
		"",
		sb.String(),
		proposal.Status.String(),
		proposal.SubmitTime,
		proposal.DepositEndTime,
		proposal.VotingStartTime,
		proposal.VotingEndTime,
		msg.Proposer,
		msg.Metadata,
	)

	// As govtypesv1.MsgSubmitProposal does not support Title and Description fields
	// it is passed with separate metadata.json and needs to be fetched.
	FetchIPFSProposalMetadata(&proposalObj)

	err = m.db.SaveProposals([]types.Proposal{proposalObj})
	if err != nil {
		return err
	}

	txTimestamp, err := time.Parse(time.RFC3339, tx.Timestamp)
	if err != nil {
		return fmt.Errorf("error while parsing time: %s", err)
	}

	// Store the deposit
	deposit := types.NewDeposit(proposal.ProposalId, msg.Proposer, msg.InitialDeposit, txTimestamp, tx.Height)
	return m.db.SaveDeposits([]types.Deposit{deposit})
}

// handleMsgDeposit allows to properly handle a handleMsgDeposit
func (m *Module) handleMsgDeposit(tx *juno.Tx, depositor string, proposalId uint64) error {
	deposit, err := m.source.ProposalDeposit(tx.Height, proposalId, depositor)
	if err != nil {
		return fmt.Errorf("error while getting proposal deposit: %s", err)
	}
	txTimestamp, err := time.Parse(time.RFC3339, tx.Timestamp)
	if err != nil {
		return fmt.Errorf("error while parsing time: %s", err)
	}

	return m.db.SaveDeposits([]types.Deposit{
		types.NewDeposit(proposalId, depositor, deposit.Amount, txTimestamp, tx.Height),
	})
}

// handleMsgVote allows to properly handle a handleMsgVote
func (m *Module) handleMsgVote(tx *juno.Tx, proposalId uint64, voter string, option int32) error {
	txTimestamp, err := time.Parse(time.RFC3339, tx.Timestamp)
	if err != nil {
		return fmt.Errorf("error while parsing time: %s", err)
	}

	vote := types.NewVote(proposalId, voter, option, txTimestamp, tx.Height)
	return m.db.SaveVote(vote)
}
