package types

import (
	accumulatortypes "github.com/cosmos/cosmos-sdk/x/accumulator/types"
)

// AccumulatorParams represents the x/mint parameters
type AccumulatorParams struct {
	accumulatortypes.Params
	Height int64
}

// NewAccumulatorParams allows to build a new MintParams instance
func NewAccumulatorParams(params accumulatortypes.Params, height int64) *AccumulatorParams {
	return &AccumulatorParams{
		Params: params,
		Height: height,
	}
}
