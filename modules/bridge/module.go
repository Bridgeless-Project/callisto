package bridge

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/forbole/bdjuno/v4/database"
	junomessages "github.com/forbole/juno/v4/modules/messages"

	"github.com/forbole/juno/v4/modules"
)

var (
	_ modules.Module        = &Module{}
	_ modules.MessageModule = &Module{}
	_ modules.GenesisModule = &Module{}
)

// Module represents the x/bridge module
type Module struct {
	cdc codec.Codec
	db  *database.Db

	messageParser junomessages.MessageAddressesParser
}

// NewModule returns a new Module instance
func NewModule(
	messageParser junomessages.MessageAddressesParser, cdc codec.Codec, db *database.Db,
) *Module {
	return &Module{
		cdc:           cdc,
		db:            db,
		messageParser: messageParser,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "bridge"
}
