package local

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	accumulatortypes "github.com/cosmos/cosmos-sdk/x/accumulator/types"
	accumulatorkeeper "github.com/forbole/bdjuno/v4/modules/accumulator/source"
	"github.com/forbole/juno/v4/node/local"
)

var (
	_ accumulatorkeeper.Source = &Source{}
)

// Source represents the implementation of the nft keeper that works on a local node
type Source struct {
	*local.Source
	q accumulatortypes.QueryServer
}

// NewSource builds a new Source instance
func NewSource(source *local.Source, nk accumulatortypes.QueryServer) *Source {
	return &Source{
		Source: source,
		q:      nk,
	}
}

func (s Source) GetAdminByAddress(address string, height int64) (*accumulatortypes.Admin, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get height context")
	}

	response, err := s.q.GetAdminByAddress(ctx, &accumulatortypes.QueryAdminByAddress{Address: address})
	if err != nil {
		return nil, errors.Wrap(err, "failed to query all nfts")
	}

	return &response.Admin, nil
}

func (s Source) GetAdmins(pagination *query.PageRequest, height int64) ([]accumulatortypes.Admin, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get height context")
	}

	response, err := s.q.GetAdmins(ctx, &accumulatortypes.QueryAdmins{
		Pagination: pagination,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to query all nfts")
	}

	return response.Admins, nil
}
