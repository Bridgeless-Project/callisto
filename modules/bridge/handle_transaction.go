package bridge

import (
	"math/big"

	"cosmossdk.io/errors"
	bridge "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/forbole/bdjuno/v4/database/types"
	juno "github.com/forbole/juno/v4/types"
)

// handleMsgSubmitBridgeTransactions allows to properly handle a MsgSubmitTransactions
func (m *Module) handleMsgSubmitBridgeTransactions(junotx *juno.Tx, msg *bridge.MsgSubmitTransactions) error {
	txs, err := m.db.GetBridgeTransactions()
	if err != nil {
		return errors.Wrap(err, "failed to get transactions")
	}

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

		if len(txSubmissions.TxHash) == 0 {
			txSubmissions.TxHash = crypto.Keccak256Hash(m.cdc.MustMarshal(&tx)).String()
		}
		txSubmissions.Submitters = append(txSubmissions.Submitters, msg.Submitter)

		if err = m.db.SaveBridgeTransactionSubmissions(txSubmissions); err != nil {
			return errors.Wrap(err, "failed to save bridge transaction submissions")
		}

		if len(txSubmissions.Submitters) == int(params.TssThreshold+1) && !m.isTxSaved(&tx, txs) {
			if err := m.db.SaveBridgeTransaction(tx, junotx.Timestamp); err != nil {
				return errors.Wrap(err, "failed to save bridge transaction")
			}

			tokenId, err := m.db.SetBridgeTransactionTokenId(tx)
			if err != nil {
				return errors.Wrap(err, "failed to set bridge transaction token id")
			}

			depositDecimals, withdrawalDecimals, err := m.db.SetBridgeTransactionDecimals(tx)
			if err != nil {
				return errors.Wrap(err, "failed to set bridge transaction decimals")
			}

			if err := m.UpdateTokenVolume(&tx, tokenId, depositDecimals, withdrawalDecimals, junotx.Timestamp); err != nil {
				return errors.Wrap(err, "failed to update token volume")
			}

		}

	}

	return nil
}

func (m *Module) handleMsgRemoveTransaction(_ *juno.Tx, msg *bridge.MsgRemoveTransaction) error {
	tx, err := m.db.GetBridgeTransaction(msg.DepositChainId, msg.DepositTxHash, msg.DepositTxIndex)
	if err != nil {
		return errors.Wrap(err, "failed to get bridge transaction")
	}
	if tx == nil {
		return nil
	}

	err = m.db.RemoveBridgeTransaction(msg.DepositChainId, msg.DepositTxHash, msg.DepositTxIndex)
	if err != nil {
		return errors.Wrap(err, "failed to remove bridge transaction")
	}

	txHash := crypto.Keccak256Hash(m.cdc.MustMarshal(tx)).String()
	err = m.db.RemoveBridgeTransactionSubmissions(txHash)
	if err != nil {
		return errors.Wrap(err, "failed to remove bridge transaction submissions")
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

func (m *Module) isTxSaved(tx *bridge.Transaction, savedTxs []bridge.Transaction) bool {
	for _, transaction := range savedTxs {
		if crypto.Keccak256Hash(m.cdc.MustMarshal(tx)).String() == crypto.Keccak256Hash(m.cdc.
			MustMarshal(&transaction)).String() {
			return true
		}
	}

	return false
}

func (m *Module) UpdateTokenVolume(tx *bridge.Transaction, tokenId uint64, depositDecimals, withdrawalDecimals uint64,
	timestamp string) error {
	nativeDecimals, err := m.db.GetNativeDecimals(tokenId)
	if err != nil {
		return errors.Wrap(err, "failed to get native decimals")
	}

	currentVolume := &types.BridgeTokenVolume{
		DepositAmount:    transformAmount(tx.DepositAmount, depositDecimals, nativeDecimals),
		WithdrawalAmount: transformAmount(tx.WithdrawalAmount, withdrawalDecimals, nativeDecimals),
		CommissionAmount: transformAmount(tx.CommissionAmount, withdrawalDecimals, nativeDecimals),
		TokenId:          tokenId,
		UpdatedAt:        timestamp,
	}

	err = m.db.SetTokenVolume(currentVolume)
	if err != nil {
		return errors.Wrap(err, "failed to save token volume")
	}

	return nil
}

func transformAmount(amount string, currentDecimals, targetDecimals uint64) *big.Int {
	result, _ := new(big.Int).SetString(amount, 10)

	if currentDecimals == targetDecimals {
		return result
	}

	if currentDecimals < targetDecimals {
		for i := uint64(0); i < targetDecimals-currentDecimals; i++ {
			result.Mul(result, new(big.Int).SetInt64(10))
		}
	} else {
		for i := uint64(0); i < currentDecimals-targetDecimals; i++ {
			result.Div(result, new(big.Int).SetInt64(10))
		}
	}

	return result
}
