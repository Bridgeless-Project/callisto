package bank

import (
	"github.com/cockroachdb/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m *Module) GetAccountBalance(address string, denom string, height int64) (sdk.Coin, error) {
	coins, err := m.keeper.GetAccountBalance(address, height)
	if err != nil {
		return sdk.Coin{}, errors.Wrap(err, "failed to get account balance")
	}

	for _, coin := range coins {
		if coin.Denom == denom {
			return coin, nil
		}
	}
	return sdk.Coin{}, errors.New("denom not found")
}
