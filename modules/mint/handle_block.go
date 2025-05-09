package mint

import (
	juno "github.com/forbole/juno/v4/types"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// HandleBlock implements modules.BlockModule
func (m *Module) HandleBlock(
	b *tmctypes.ResultBlock, blockResults *tmctypes.ResultBlockResults, txs []*juno.Tx, vals *tmctypes.ResultValidators,
) error {
	return m.UpdateParams(b.Block.Height)
}
