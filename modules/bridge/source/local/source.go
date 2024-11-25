package local

import (
	bridgetypes "github.com/hyle-team/bridgeless-core/v12/x/bridge/types"

	"github.com/forbole/juno/v4/node/local"

	"github.com/forbole/bdjuno/v4/modules/bridge/source"
)

var (
	_ source.Source = &Source{}
)

// Source represents the implementation of the bank keeper that works on a local node
type Source struct {
	*local.Source
	q bridgetypes.QueryServer
}

// NewSource builds a new Source instance
func NewSource(source *local.Source, bk bridgetypes.QueryServer) *Source {
	return &Source{
		Source: source,
		q:      bk,
	}
}
