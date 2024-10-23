package source

import (
	"github.com/cosmos/cosmos-sdk/types/query"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft/types"
)

type Source interface {
	GetNFT(address string, height int64) (val nfttypes.NFT, found bool)
	GetNFTsWithPagination(pagination *query.PageRequest, height int64) (val []nfttypes.NFT, pr *query.PageResponse, err error)
}
