package local

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft/types"
	"github.com/forbole/juno/v4/node/local"

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

// NewSource builds a new Source instance
func NewSource(source *local.Source, nk nfttypes.QueryServer) *Source {
	return &Source{
		Source: source,
		q:      nk,
	}
}

// GetNFT implements keeper.Source
// TODO return nil instead of empty object
func (s Source) GetNFT(address string, height int64) (val nfttypes.NFT, found bool) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return nfttypes.NFT{}, false
	}

	nft, err := s.q.GetNFTByAddress(sdk.WrapSDKContext(ctx), &nfttypes.QueryNFTByAddress{Address: address})
	if err != nil {
		return nfttypes.NFT{}, false
	}

	return *nft.Nft, true
}
