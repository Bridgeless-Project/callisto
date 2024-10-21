package nft

import (
	"fmt"
	juno "github.com/forbole/juno/v4/types"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// HandleBlock implements BlockModule
func (m *Module) HandleBlock(
	block *tmctypes.ResultBlock, results *tmctypes.ResultBlockResults, tx []*juno.Tx, resVal *tmctypes.ResultValidators,
) error {

	fmt.Println("results: ", results)
	fmt.Println("tx: ", tx)
	fmt.Println("resVal: ", resVal)
	fmt.Println("block: ", block)

	// TODO update implementation

	return nil
}
