package local

import (
	multisigtypes "github.com/hyle-team/bridgeless-core/v12/x/multisig/types"

	"github.com/forbole/juno/v4/node/local"

	"github.com/forbole/bdjuno/v4/modules/multisig/source"
)

var (
	_ source.Source = &Source{}
)

// Source represents the implementation of the bank keeper that works on a local node
type Source struct {
	*local.Source
	q multisigtypes.QueryServer
}

func (s Source) Group(height int64, account string) (multisigtypes.Group, error) {
	//TODO implement me
	panic("implement me")
}

func (s Source) Proposal(height int64, id uint64) (multisigtypes.Proposal, error) {
	//TODO implement me
	panic("implement me")
}

func (s Source) ProposalAll(height int64) ([]multisigtypes.Proposal, error) {
	//TODO implement me
	panic("implement me")
}

func (s Source) Vote(height int64, proposalId uint64, voter string) (multisigtypes.Vote, error) {
	//TODO implement me
	panic("implement me")
}

func (s Source) Params(height int64) (multisigtypes.Params, error) {
	//TODO implement me
	panic("implement me")
}

// NewSource builds a new Source instance
func NewSource(source *local.Source, bk multisigtypes.QueryServer) *Source {
	return &Source{
		Source: source,
		q:      bk,
	}
}
