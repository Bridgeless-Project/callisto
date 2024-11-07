package nft

import (
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type StakingModule interface {
	GetDelegationByValidator(height int64, delegator string, validator string) (*stakingtypes.DelegationResponse, error)
}
