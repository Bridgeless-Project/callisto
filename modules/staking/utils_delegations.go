package staking

import (
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func (m *Module) GetDelegationByValidator(height int64, delegator string, validator string) (*stakingtypes.DelegationResponse, error) {
	res, err := m.source.GetDelegationByValidator(
		height, delegator, validator,
	)
	if err != nil {
		return nil, err
	}

	return res.DelegationResponse, nil
}
