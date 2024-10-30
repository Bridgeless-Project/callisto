package remote

import (
	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	accumulatortypes "github.com/cosmos/cosmos-sdk/x/accumulator/types"
	"github.com/forbole/juno/v4/node/remote"

	accumulatorkeeper "github.com/forbole/bdjuno/v4/modules/accumulator/source"
)

var (
	_ accumulatorkeeper.Source = &Source{}
)

type Source struct {
	*remote.Source
	accumulatorClient accumulatortypes.QueryClient
}

// NewSource builds a new Source instance
func NewSource(source *remote.Source, accumulatorClient accumulatortypes.QueryClient) *Source {
	return &Source{
		Source:            source,
		accumulatorClient: accumulatorClient,
	}
}

func (s Source) GetAdmins(pagination *query.PageRequest, height int64) ([]accumulatortypes.Admin, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)

	response, err := s.accumulatorClient.GetAdmins(ctx, &accumulatortypes.QueryAdmins{
		Pagination: pagination,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to query all nfts")
	}

	return response.Admins, nil
}

func (s Source) GetAdminByAddress(address string, height int64) (*accumulatortypes.Admin, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)

	response, err := s.accumulatorClient.GetAdminByAddress(ctx, &accumulatortypes.QueryAdminByAddress{
		Address: address,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to query all nfts")
	}

	return &response.Admin, nil
}
