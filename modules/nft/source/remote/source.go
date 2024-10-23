package remote

import (
	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft/types"

	"github.com/forbole/juno/v4/node/remote"

	nftkeeper "github.com/forbole/bdjuno/v4/modules/nft/source"
)

var (
	_ nftkeeper.Source = &Source{}
)

type Source struct {
	*remote.Source
	nftClient nfttypes.QueryClient
}

func (s Source) GetNFTsWithPagination(pagination *query.PageRequest, height int64) (val []nfttypes.NFT, pr *query.PageResponse, err error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)

	response, err := s.nftClient.GetAllNFTs(ctx, &nfttypes.QueryAllNFTs{Pagination: pagination})
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to query all nfts")
	}

	return response.Nft, response.Pagination, nil
}

// NewSource builds a new Source instance
func NewSource(source *remote.Source, nftClient nfttypes.QueryClient) *Source {
	return &Source{
		Source:    source,
		nftClient: nftClient,
	}
}

func (s Source) GetNFT(address string, height int64) (val nfttypes.NFT, found bool) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	nft, err := s.nftClient.GetNFTByAddress(ctx, &nfttypes.QueryNFTByAddress{Address: address})
	if err != nil {
		return nfttypes.NFT{}, false
	}

	return *nft.Nft, true

}
