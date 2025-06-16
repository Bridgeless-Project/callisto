package multisig

import (
	"fmt"
	multisigtypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/multisig/types"
	"github.com/forbole/bdjuno/v4/types"
)

func (m *Module) saveGroups(slice []multisigtypes.Group) error {
	// Save the groups
	groups := make([]*types.Group, len(slice))
	for i, group := range slice {
		groups[i] = types.GroupFromCore(group)
	}

	err := m.db.SaveGroups(groups)
	if err != nil {
		return fmt.Errorf("error while storing multisig groups: %s", err)
	}

	return nil
}

func (m *Module) saveGroup(height int64, account string) error {
	group, err := m.source.Group(height, account)
	if err != nil {
		return fmt.Errorf("error while getting group: %s", err)
	}

	return m.saveGroups([]multisigtypes.Group{group})
}
