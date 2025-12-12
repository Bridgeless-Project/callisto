package bridge

import (
	"math/big"

	"cosmossdk.io/errors"
	bridge "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/forbole/bdjuno/v4/database/types"
	juno "github.com/forbole/juno/v4/types"
	"github.com/rs/zerolog/log"
)

// handleMsgSubmitBridgeTransactions allows to properly handle a MsgSubmitTransactions
func (m *Module) handleMsgSubmitBridgeTransactions(junotx *juno.Tx, msg *bridge.MsgSubmitTransactions) error {
	log.Debug().Str("module", "bridge").Msg("Getting Bridge txs")
	txs, err := m.db.GetBridgeTransactions()
	if err != nil {
		return errors.Wrap(err, "failed to get transactions")
	}

	log.Debug().Str("module", "bridge").Msg("Start iterating")
	for _, tx := range msg.Transactions {
		txBytes, err := m.cdc.Marshal(&tx)
		if err != nil {
			return errors.Wrap(err, "failed to marshal tx")
		}
		log.Debug().Str("module", "bridge").Msg("Getting Bridge tx submissions")
		txSubmissions, err := m.db.GetBridgeTransactionSubmissions(crypto.Keccak256Hash(txBytes).String())
		if err != nil {
			return errors.Wrap(err, "failed to get transaction submissions")
		}

		log.Debug().Str("module", "bridge").Msg("Checking if submitter")
		if isSubmitter(txSubmissions.Submitters, msg.Submitter) {
			return nil
		}

		log.Debug().Str("module", "bridge").Msg("Getting bridge tx params")
		params, err := m.db.GetBridgeParams()
		if err != nil {
			return errors.Wrap(err, "failed to get bridge params")
		}

		if len(txSubmissions.TxHash) == 0 {
			txSubmissions.TxHash = crypto.Keccak256Hash(txBytes).String()
		}
		txSubmissions.Submitters = append(txSubmissions.Submitters, msg.Submitter)

		log.Debug().Str("module", "bridge").Msg("Saving submissions")
		if err = m.db.SaveBridgeTransactionSubmissions(txSubmissions); err != nil {
			return errors.Wrap(err, "failed to save bridge transaction submissions")
		}

		log.Debug().Str("module", "bridge").Msg("Checking if saved")
		saved, err := m.isTxSaved(&tx, txs)
		if err != nil {
			return errors.Wrap(err, "failed to check tx saved")
		}

		log.Debug().Str("module", "bridge").Msg("Saving tx")
		if len(txSubmissions.Submitters) == int(params.TssThreshold+1) && !saved {
			if err := m.db.SaveBridgeTransaction(tx, junotx.Timestamp); err != nil {
				return errors.Wrap(err, "failed to save bridge transaction")
			}
			log.Debug().Str("module", "bridge").Msg("Saving tx token id")
			tokenId, err := m.db.SetBridgeTransactionTokenId(tx)
			if err != nil {
				return errors.Wrap(err, "failed to set bridge transaction token id")
			}

			log.Debug().Str("module", "bridge").Msg("Saving tx dec")
			depositDecimals, withdrawalDecimals, err := m.db.SetBridgeTransactionDecimals(tx)
			if err != nil {
				return errors.Wrap(err, "failed to set bridge transaction decimals")
			}

			log.Debug().Str("module", "bridge").Msg("Saving tx volume")
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

	txHashBytes, err := m.cdc.Marshal(tx)
	if err != nil {
		return errors.Wrap(err, "failed to marshal tx")
	}

	txHash := crypto.Keccak256Hash(txHashBytes).String()
	err = m.db.RemoveBridgeTransactionSubmissions(txHash)
	if err != nil {
		return errors.Wrap(err, "failed to remove bridge transaction submissions")
	}

	return nil
}

func (m *Module) handleMsgUpdateTransaction(_ *juno.Tx, msg *bridge.MsgUpdateTransaction) error {
	tx, err := m.db.GetBridgeTransaction(msg.Transaction.DepositChainId, msg.Transaction.DepositTxHash, msg.Transaction.DepositTxIndex)
	if err != nil {
		return errors.Wrap(err, "failed to get bridge transaction")
	}
	if tx == nil {
		return nil
	}

	if err = compareTxs(msg.Transaction, *tx); err != nil {
		return errors.Wrap(err, "failed to compare transactions")
	}

	err = m.db.UpdateTransactionWithdrawalTxHash(
		msg.Transaction.DepositChainId,
		msg.Transaction.DepositTxHash,
		msg.Transaction.DepositTxIndex,
		msg.Transaction.WithdrawalTxHash,
	)
	if err != nil {
		return errors.Wrap(err, "failed to update withdrawal transaction hash")
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

func (m *Module) isTxSaved(tx *bridge.Transaction, savedTxs []bridge.Transaction) (bool, error) {
	txBytes, err := m.cdc.Marshal(tx)
	if err != nil {
		return false, errors.Wrap(err, "failed to marshal tx")
	}

	for _, transaction := range savedTxs {
		transactionBytes, err := m.cdc.Marshal(&transaction)
		if err != nil {
			return false, errors.Wrap(err, "failed to marshal transaction")
		}

		if crypto.Keccak256Hash(txBytes).String() == crypto.Keccak256Hash(transactionBytes).String() {
			return true, nil
		}
	}

	return false, nil
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
