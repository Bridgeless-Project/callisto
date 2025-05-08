package database

import (
	"database/sql"
	"fmt"
	"github.com/forbole/bdjuno/v4/database/types"
	bridgeTypes "github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

// SaveChain allows to save new Chain
func (db *Db) SaveBridgeChain(id string, chainType int32, bridgeAddress string, operator string) error {
	query := `
		INSERT INTO bridge_chains(id, chain_type, bridge_address, operator) 
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (id) DO UPDATE
		SET chain_type = excluded.chain_type,
			bridge_address = excluded.bridge_address,
			operator = excluded.operator
	`
	_, err := db.SQL.Exec(query, id, chainType, bridgeAddress, operator)
	if err != nil {
		return fmt.Errorf("error while storing chain: %s", err)
	}

	return nil
}

// RemoveChain allows to remove the Chain
func (db *Db) RemoveBridgeChain(id string) error {
	query := `
		DELETE FROM bridge_chains WHERE id = $1
	`
	_, err := db.SQL.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error while removing chain: %s", err)
	}

	return nil
}

// -------------------------------------------------------------------------------------------------------------------

// SaveTokenInfo allows to save new TokenInfo
func (db *Db) SaveBridgeTokenInfo(address string, decimals uint64, chainID string, tokenID uint64, isWrapped bool) (int64, error) {
	query := `
		INSERT INTO bridge_tokens_info(address, decimals, chain_id, token_id, is_wrapped) 
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (address) DO UPDATE
		SET chain_id = excluded.chain_id,
			token_id = excluded.token_id,
			is_wrapped = excluded.is_wrapped
		RETURNING id
	`

	var id int64
	err := db.SQL.QueryRow(query, address, decimals, chainID, tokenID, isWrapped).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error while storing token info: %s", err)
	}

	return id, nil
}

// RemoveTokenInfo allows to remove the TokenInfo
func (db *Db) RemoveBridgeTokenInfo(id int64) error {
	query := `
		DELETE FROM bridge_tokens_info WHERE id = $1
	`
	_, err := db.SQL.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error while removing token info: %s", err)
	}

	return nil
}

// -------------------------------------------------------------------------------------------------------------------

// SaveTokenMetadata allows to save new TokenMetadata
func (db *Db) SaveBridgeTokenMetadata(tokenID uint64, name, symbol, uri string) error {
	query := `
		INSERT INTO bridge_token_metadata(token_id, name, symbol, uri) 
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (token_id) DO UPDATE
		SET name = excluded.name,
			symbol = excluded.symbol,
			uri = excluded.uri
	`

	_, err := db.SQL.Exec(query, tokenID, name, symbol, uri)
	if err != nil {
		return fmt.Errorf("error while storing token metadata: %s", err)
	}

	return nil
}

// RemoveTokenMetadata allows to remove the TokenMetadata
func (db *Db) RemoveBridgeTokenMetadata(id int64) error {
	query := `
		DELETE FROM bridge_token_metadata WHERE id = $1
	`
	_, err := db.SQL.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error while removing token metadata: %s", err)
	}

	return nil
}

// -------------------------------------------------------------------------------------------------------------------

// SaveBridgeTokens allows to save new Tokens
func (db *Db) SaveBridgeToken(tokensInfoID int64, tokenMetadataID uint64, commissionRate string) error {
	query := `
		INSERT INTO bridge_tokens(tokens_info_id, metadata_id, commission_rate) 
		VALUES ($1, $2, $3) 
		ON CONFLICT (tokens_info_id, metadata_id) DO NOTHING

	`

	_, err := db.SQL.Exec(query, tokensInfoID, tokenMetadataID, commissionRate)
	if err != nil {
		return fmt.Errorf("error while storing token: %s", err)
	}
	return nil
}

// RemoveBridgeTokens allows to remove the Tokens
func (db *Db) RemoveBridgeToken(tokenID uint64) error {
	query := `
		DELETE FROM bridge_tokens WHERE metadata_id = $1;
		DELETE FROM bridge_tokens_info WHERE id = $1;
		DELETE FROM bridge_token_metadata WHERE token_id = $1;
	`
	_, err := db.SQL.Exec(query, tokenID)
	if err != nil {
		return fmt.Errorf("error while removing token: %s", err)
	}

	return nil
}

// -------------------------------------------------------------------------------------------------------------------

func (db *Db) SaveBridgeTransaction(
	tx bridgeTypes.Transaction,
) error {
	query := `
		INSERT INTO bridge_transactions(
			deposit_chain_id, 
			deposit_tx_hash, 
			deposit_tx_index,
			deposit_block, 
			deposit_token, 
			depositor,
			receiver,
			withdrawal_chain_id,
			withdrawal_tx_hash,
			withdrawal_token, 
			signature,
			is_wrapped,
			deposit_amount,
		    withdrawal_amount,
			commission_amount,
		    tx_data
	 	) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) RETURNING id
	`
	_, err := db.SQL.Exec(
		query,
		tx.DepositChainId,
		tx.DepositTxHash,
		tx.DepositTxIndex,
		tx.DepositBlock,
		tx.DepositToken,
		tx.Depositor,
		tx.Receiver,
		tx.WithdrawalChainId,
		tx.WithdrawalTxHash,
		tx.WithdrawalToken,
		tx.Signature,
		tx.IsWrapped,
		tx.DepositAmount,
		tx.WithdrawalAmount,
		tx.CommissionAmount,
		tx.TxData,
	)
	if err != nil {
		return fmt.Errorf("error while storing transaction: %s", err)
	}

	return nil
}

func (db *Db) GetBridgeTransactions() ([]bridgeTypes.Transaction, error) {
	var (
		txs []types.Transaction
		res []bridgeTypes.Transaction
	)

	err := db.Sqlx.Select(&txs, `SELECT * FROM bridge_transactions`)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) || len(txs) == 0 {

			return res, nil
		}

		return nil, fmt.Errorf("error while getting transactions: %s", err)
	}

	for _, tx := range txs {
		res = append(res, *types.ToBridgeTransaction(tx))
	}

	return res, nil

}

// -------------------------------------------------------------------------------------------------------------------

func (db *Db) SaveBridgeTransactionSubmissions(txSubmissions *bridgeTypes.TransactionSubmissions) error {
	query := `INSERT INTO bridge_transaction_submissions (tx_hash,submitters) VALUES ($1, $2)
				ON CONFLICT (tx_hash) DO UPDATE
				SET submitters = excluded.submitters
				`

	_, err := db.SQL.Exec(query, txSubmissions.TxHash, pq.Array(txSubmissions.Submitters))
	if err != nil {
		return fmt.Errorf("error while storing transaction submissions: %s", err)
	}

	return nil
}

func (db *Db) GetBridgeTransactionSubmissions(txHash string) (*bridgeTypes.TransactionSubmissions, error) {
	var txSubmissions []types.TxSubmissions
	err := db.Sqlx.Select(&txSubmissions, `SELECT * FROM bridge_transaction_submissions WHERE tx_hash = $1`, txHash)

	if errors.Is(err, sql.ErrNoRows) || len(txSubmissions) == 0 {

		return &bridgeTypes.TransactionSubmissions{
			TxHash:     "",
			Submitters: nil,
		}, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error while getting transaction submissions: %s", err)
	}

	return types.ToTransactionSubmissions(txSubmissions[0]), nil
}

// -------------------------------------------------------------------------------------------------------------------

func (db *Db) SaveBridgeParams(params *bridgeTypes.Params) error {
	query := `INSERT INTO bridge_params (id, module_admin,parties,tss_threshold) VALUES ($1, $2, $3, $4)
				ON CONFLICT (id) DO UPDATE
				SET module_admin = excluded.module_admin,
				parties = excluded.parties,
				tss_threshold = excluded.tss_threshold`

	var parties []string
	for _, party := range params.Parties {
		parties = append(parties, party.Address)
	}
	_, err := db.SQL.Exec(query, 1, params.ModuleAdmin, pq.StringArray(parties), int(params.TssThreshold))
	if err != nil {
		return fmt.Errorf("error while storing bridge_params: %s", err)
	}

	return nil
}

func (db *Db) GetBridgeParams() (*bridgeTypes.Params, error) {
	var params []types.Params

	err := db.Sqlx.Select(&params, `SELECT * FROM bridge_params`)
	if err != nil {
		return nil, fmt.Errorf("error while getting bridge_params: %s", err)
	}

	if len(params) == 0 {
		return nil, fmt.Errorf("error while getting bridge_params: no params found")
	}

	if len(params) > 1 {
		return nil, fmt.Errorf("error while getting bridge_params: more than one param found")
	}

	return types.ToBridgeParams(params[0]), nil
}
