package multisig

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/bdjuno/v4/database"
	"github.com/forbole/bdjuno/v4/modules/auth"

	multisig "github.com/forbole/bdjuno/v4/modules/multisig/source"

	"github.com/forbole/juno/v4/modules"
)

var (
	_ modules.Module        = &Module{}
	_ modules.GenesisModule = &Module{}
	_ modules.MessageModule = &Module{}
	_ modules.BlockModule   = &Module{}
)

// Module represents the x/multisig module
type Module struct {
	cdc    codec.Codec
	db     *database.Db
	source multisig.Source
	auth   *auth.Module
}

// NewModule builds a new Module instance
func NewModule(source multisig.Source, cdc codec.Codec, db *database.Db, auth *auth.Module) *Module {
	return &Module{
		source: source,
		cdc:    cdc,
		db:     db,
		auth:   auth,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "multisig"
}
