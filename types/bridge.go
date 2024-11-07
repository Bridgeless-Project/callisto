package types

import (
	bridgetypes "github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
)

// BridgeStore represents the x/bridge store
type BridgeStore struct {
	Params       bridgetypes.Params
	Chains       []bridgetypes.Chain
	Tokens       []bridgetypes.Token
	Transactions []bridgetypes.Transaction
	Height       int64
}

// NewAccumulatorParams allows to build a new MintParams instance
func NewBridgeStore(params bridgetypes.Params, chains []bridgetypes.Chain, tokens []bridgetypes.Token, transactions []bridgetypes.Transaction, height int64) *BridgeStore {
	return &BridgeStore{
		Params:       params,
		Chains:       chains,
		Tokens:       tokens,
		Transactions: transactions,
		Height:       height,
	}
}
