package bridge

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"

	bridge "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	juno "github.com/forbole/juno/v4/types"
)

func (m *Module) handleMsgStartEpoch(tx *juno.Tx, msg *bridge.MsgStartEpoch) error {
	params, err := m.db.GetBridgeParams()
	if err != nil {
		return fmt.Errorf("handleMsgStartEpoch: get params: %w", err)
	}

	parties, err := determineEpochParties(params.Parties, msg.Info)
	if err != nil {
		return fmt.Errorf("handleMsgStartEpoch: determine epoch parties: %w", err)
	}

	epoch := &bridge.Epoch{
		Id:           msg.EpochId,
		Status:       bridge.EpochStatus_INITIATED,
		StartBlock:   uint64(tx.Height),
		Parties:      parties,
		TssThreshold: msg.TssThreshold,
		TssInfo:      msg.Info,
	}

	if err = m.db.SaveBridgeEpoch(epoch); err != nil {
		return fmt.Errorf("handleMsgStartEpoch: save epoch: %w", err)
	}

	return nil
}

func (m *Module) handleMsgSetEpochPubKey(_ *juno.Tx, msg *bridge.MsgSetEpochPubKey) error {
	epoch, err := m.db.GetBridgeEpoch(msg.EpochId)
	if err != nil {
		return fmt.Errorf("handleMsgSetEpochPubKey: get epoch: %w", err)
	}
	if epoch == nil {
		return nil
	}

	submissions, err := m.db.GetBridgeEpochPubKeySubmissions(msg.EpochId, msg.Pubkey)
	if err != nil {
		return fmt.Errorf("handleMsgSetEpochPubKey: get submissions: %w", err)
	}
	if isSubmitter(submissions.Submitters, msg.Creator) {
		return nil
	}

	submissions.EpochId = msg.EpochId
	submissions.Hash = msg.Pubkey
	submissions.Submitters = append(submissions.Submitters, msg.Creator)
	if err = m.db.SaveBridgeEpochPubKeySubmissions(submissions); err != nil {
		return fmt.Errorf("handleMsgSetEpochPubKey: save submissions: %w", err)
	}

	if len(submissions.Submitters) == int(epoch.TssThreshold+1) {
		if err = m.db.SaveBridgeEpochPubKey(msg.EpochId, msg.Pubkey); err != nil {
			return fmt.Errorf("handleMsgSetEpochPubKey: save pubkey: %w", err)
		}
	}

	return nil
}

func (m *Module) handleMsgSetEpochSignature(_ *juno.Tx, msg *bridge.MsgSetEpochSignature) error {
	params, err := m.db.GetBridgeParams()
	if err != nil {
		return fmt.Errorf("handleMsgSetEpochSignature: get params: %w", err)
	}

	hash, err := epochSignaturesHash(msg.EpochChainSignatures)
	if err != nil {
		return fmt.Errorf("handleMsgSetEpochSignature: hash signatures: %w", err)
	}

	submissions, err := m.db.GetBridgeEpochSignatureSubmissions(msg.EpochId, hash)
	if err != nil {
		return fmt.Errorf("handleMsgSetEpochSignature: get submissions: %w", err)
	}
	if isSubmitter(submissions.Submitters, msg.Creator) {
		return nil
	}

	submissions.EpochId = msg.EpochId
	submissions.Hash = hash
	submissions.Submitters = append(submissions.Submitters, msg.Creator)
	if err = m.db.SaveBridgeEpochSignatureSubmissions(submissions); err != nil {
		return fmt.Errorf("handleMsgSetEpochSignature: save submissions: %w", err)
	}

	if len(submissions.Submitters) != int(params.TssThreshold+1) {
		return nil
	}

	for _, sig := range msg.EpochChainSignatures {
		if err = m.db.SaveBridgeEpochChainSignatures(&sig); err != nil {
			return fmt.Errorf("handleMsgSetEpochSignature: save chain signatures: %w", err)
		}
	}

	for _, address := range msg.Addresses {
		if err = m.db.UpdateBridgeChainAddress(address.ChainId, address.Address); err != nil {
			return fmt.Errorf("handleMsgSetEpochSignature: update bridge address: %w", err)
		}
	}

	epoch, err := m.db.GetBridgeEpoch(msg.EpochId)
	if err != nil {
		return fmt.Errorf("handleMsgSetEpochSignature: get epoch: %w", err)
	}
	if epoch == nil {
		return nil
	}
	epoch.Status = bridge.EpochStatus_MIGRATION_FINALIZING
	if err = m.db.SaveBridgeEpoch(epoch); err != nil {
		return fmt.Errorf("handleMsgSetEpochSignature: save epoch: %w", err)
	}

	return nil
}

func determineEpochParties(tssParties []*bridge.Party, tssInfo []bridge.TSSInfo) ([]*bridge.Party, error) {
	parties := append([]*bridge.Party(nil), tssParties...)
	for _, info := range tssInfo {
		found := false
		for i, party := range parties {
			if party.Address != info.Address {
				continue
			}
			found = true
			if info.Active {
				return nil, fmt.Errorf("duplicate active party found: %s", info.Address)
			}
			parties = append(parties[:i], parties[i+1:]...)
			break
		}
		if !found {
			parties = append(parties, &bridge.Party{Address: info.Address})
		}
	}

	return parties, nil
}

func epochSignaturesHash(signatures []bridge.EpochChainSignatures) (string, error) {
	bz, err := json.Marshal(signatures)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(bz)
	return hex.EncodeToString(sum[:]), nil
}
