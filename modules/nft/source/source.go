package source

import (
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft/types"
)

type Source interface {
	GetNFT(address string, height int64) (val nfttypes.NFT, found bool)
}
