package source

import (
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/x/accumulator/types"
)

type Source interface {
	GetAdmins(req *query.PageRequest, height int64) ([]types.Admin, error)
}
