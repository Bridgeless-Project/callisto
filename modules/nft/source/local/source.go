package local

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft/types"
	"github.com/forbole/juno/v4/node/local"
	"github.com/rs/zerolog/log"

	nftkeeper "github.com/forbole/bdjuno/v4/modules/nft/source"
)

var (
	_ nftkeeper.Source = &Source{}
)

// Source represents the implementation of the nft keeper that works on a local node
type Source struct {
	*local.Source
	q nfttypes.QueryServer
}

func (s Source) GetNFTsWithPagination(pagination *query.PageRequest, height int64) (val []nfttypes.NFT, pr *query.PageResponse, err error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error while loading height")
	}

	response, err := s.q.GetAllNFTs(sdk.WrapSDKContext(ctx), &nfttypes.QueryAllNFTs{Pagination: pagination})
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to query all nfts")
	}

	return response.Nft, response.Pagination, nil
}

// NewSource builds a new Source instance
func NewSource(source *local.Source, nk nfttypes.QueryServer) *Source {
	return &Source{
		Source: source,
		q:      nk,
	}
}

// GetNFT implements keeper.Source
// TODO return nil instead of empty object (update core)
func (s Source) GetNFT(address string, height int64) (val nfttypes.NFT, found bool) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		log.Err(err).Msg("error while loading height")
		return nfttypes.NFT{}, false
	}

	nft, err := s.q.GetNFTByAddress(sdk.WrapSDKContext(ctx), &nfttypes.QueryNFTByAddress{Address: address})
	if err != nil {
		log.Err(err).Msg("error while loading nft by height")
		return nfttypes.NFT{}, false
	}

	return *nft.Nft, true
}
