package remote

import (
	bridgetypes "github.com/hyle-team/bridgeless-core/v12/x/bridge/types"

	"github.com/forbole/juno/v4/node/remote"

	bridgekeeper "github.com/forbole/bdjuno/v4/modules/bridge/source"
)

var (
	_ bridgekeeper.Source = &Source{}
)

type Source struct {
	*remote.Source
	client bridgetypes.QueryClient
}

// NewSource builds a new Source instance
func NewSource(source *remote.Source, bridgeClient bridgetypes.QueryClient) *Source {
	return &Source{
		Source: source,
		client: bridgeClient,
	}
}
