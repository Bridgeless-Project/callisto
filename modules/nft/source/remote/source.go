package remote

import (
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
