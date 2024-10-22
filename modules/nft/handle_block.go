package nft

import (
	juno "github.com/forbole/juno/v4/types"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// HandleBlock implements BlockModule
func (m *Module) HandleBlock(
	block *tmctypes.ResultBlock, results *tmctypes.ResultBlockResults, tx []*juno.Tx, resVal *tmctypes.ResultValidators,
) error {
	// TODO update vesting a block

	return nil
}
