package bridge

import (
	"fmt"
	"strconv"

	bridge "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	juno "github.com/forbole/juno/v4/types"
	"github.com/rs/zerolog/log"
	abci "github.com/tendermint/tendermint/abci/types"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// HandleBlock implements modules.BlockModule.
func (m *Module) HandleBlock(
	block *tmctypes.ResultBlock, res *tmctypes.ResultBlockResults, _ []*juno.Tx, _ *tmctypes.ResultValidators,
) error {
	if err := m.handleBridgeABCI(block.Block.Height, res.BeginBlockEvents); err != nil {
		log.Error().Str("module", "bridge").Int64("height", block.Block.Height).
			Err(err).Msg("error while handling bridge abci")
	}

	return nil
}

func (m *Module) handleBridgeABCI(height int64, events []abci.Event) error {
	seen := make(map[string]bool)
	for _, event := range events {
		switch event.Type {
		case bridge.EventType_EPOCH_UPDATED.String():
			epochID, err := uint32EventAttr(event, bridge.AttributeEpochId)
			if err != nil {
				return err
			}
			isAdding, err := boolEventAttr(event, bridge.AttributeEpochSignatureMode)
			if err != nil {
				return err
			}
			key := fmt.Sprintf("%d/%t", epochID, isAdding)
			if seen[key] {
				continue
			}
			seen[key] = true
			if err = m.handleEpochUpdatedABCI(height, epochID, isAdding); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *Module) handleEpochUpdatedABCI(height int64, epochID uint32, isAdding bool) error {
	epoch, err := m.db.GetBridgeEpoch(epochID)
	if err != nil {
		return fmt.Errorf("get epoch %d: %w", epochID, err)
	}
	if epoch == nil {
		return nil
	}

	if !isAdding {
		epoch.Status = bridge.EpochStatus_UNSUPPORTED
		if err = m.db.SaveBridgeEpoch(epoch); err != nil {
			return fmt.Errorf("save epoch %d: %w", epochID, err)
		}
		return nil
	}

	params, err := m.db.GetBridgeParams()
	if err != nil {
		return fmt.Errorf("get bridge params: %w", err)
	}

	epoch.Status = bridge.EpochStatus_RUNNING
	epoch.FinalizedBlock = uint64(height)
	if err = m.db.SaveBridgeEpoch(epoch); err != nil {
		return fmt.Errorf("save epoch %d: %w", epochID, err)
	}

	if params.Epoch > 0 {
		prevEpoch, err := m.db.GetBridgeEpoch(params.Epoch)
		if err != nil {
			return fmt.Errorf("get previous epoch %d: %w", params.Epoch, err)
		}
		if prevEpoch != nil {
			prevEpoch.Status = bridge.EpochStatus_SHUTDOWN
			prevEpoch.EndBlock = uint64(height) + params.SupportingTime
			if err = m.db.SaveBridgeEpoch(prevEpoch); err != nil {
				return fmt.Errorf("save previous epoch %d: %w", params.Epoch, err)
			}
		}
	}

	params.Epoch = epoch.Id
	params.TssThreshold = epoch.TssThreshold
	params.Parties = epoch.Parties
	if err = m.db.SaveBridgeParams(params); err != nil {
		return fmt.Errorf("save bridge params: %w", err)
	}

	return nil
}

func uint32EventAttr(event abci.Event, key string) (uint32, error) {
	raw, err := eventAttr(event, key)
	if err != nil {
		return 0, err
	}
	value, err := strconv.ParseUint(raw, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("parse %s: %w", key, err)
	}

	return uint32(value), nil
}

func boolEventAttr(event abci.Event, key string) (bool, error) {
	raw, err := eventAttr(event, key)
	if err != nil {
		return false, err
	}
	value, err := strconv.ParseBool(raw)
	if err != nil {
		return false, fmt.Errorf("parse %s: %w", key, err)
	}

	return value, nil
}

func eventAttr(event abci.Event, key string) (string, error) {
	for _, attr := range event.Attributes {
		if string(attr.Key) == key {
			return string(attr.Value), nil
		}
	}

	return "", fmt.Errorf("event %s missing attribute %s", event.Type, key)
}
