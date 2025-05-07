package bridge

import (
	"cosmossdk.io/errors"
	"github.com/ethereum/go-ethereum/crypto"
	juno "github.com/forbole/juno/v4/types"
	bridge "github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
)

// handleMsgSubmitBridgeTransactions allows to properly handle a MsgSubmitTransactions
func (m *Module) handleMsgSubmitBridgeTransactions(_ *juno.Tx, msg *bridge.MsgSubmitTransactions) error {
	for _, tx := range msg.Transactions {
		txSubmissions, err := m.db.GetBridgeTransactionSubmissions(crypto.Keccak256Hash(m.cdc.MustMarshal(&tx)).String())
		if err != nil {
			return errors.Wrap(err, "failed to get transaction submissions")
		}

		if isSubmitter(txSubmissions.Submitters, msg.Submitter) {
			return nil
		}

		params, err := m.db.GetBridgeParams()
		if err != nil {
			return errors.Wrap(err, "failed to get bridge params")
		}
		txSubmissions.Submitters = append(txSubmissions.Submitters, msg.Submitter)

		if err = m.db.SaveBridgeTransactionSubmissions(txSubmissions); err != nil {
			return errors.Wrap(err, "failed to save bridge transaction submissions")
		}

		if len(txSubmissions.Submitters) == int(params.TssThreshold+1) {
			if err := m.db.SaveBridgeTransaction(tx); err != nil {
				return errors.Wrap(err, "failed to save bridge transaction")
			}
		}

	}

	return nil
}

func isSubmitter(submitters []string, submitter string) bool {
	for _, s := range submitters {
		if submitter == s {
			return true
		}
	}

	return false
}
