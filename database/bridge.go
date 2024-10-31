package database

import (
	"fmt"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
)

// SaveChain allows to save new Chain
func (db *Db) SaveBridgeChain(id string, chainType int32, bridgeAddress string, operator string) error {
	query := `
		INSERT INTO chains(id, chain_type, bridge_address, operator) 
		VALUES ($1, $2, $3, $4) RETURNING id
		ON CONFLICT (id) DO UPDATE
	SET chain_type = excluded.chain_type,
		bridge_address = excluded.bridge_address,
		operator = excluded.operator
	WHERE chains.id <= excluded.id
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
		DELETE FROM chains WHERE id = $1
	`
	_, err := db.SQL.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error while removing chain: %s", err)
	}

	return nil
}

// -------------------------------------------------------------------------------------------------------------------

// SaveTokenInfo allows to save new TokenInfo
func (db *Db) SaveBridgeTokenInfo(address string, decimals uint64, chaiID string, tokenID uint64, isWrapped bool) (int64, error) {
	query := `
		INSERT INTO tokens_info( address, decimals, chaiId, tokenId, isWrapped) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id
		ON CONFLICT (address) DO UPDATE
	SET chaiId = excluded.chaiId,
		tokenId = excluded.tokenId,
	WHERE tokens_info.address <= excluded.address
	`

	var id int64
	err := db.SQL.QueryRow(query, address, decimals, chaiID, tokenID, isWrapped)
	if err != nil {
		return emptyIndex, fmt.Errorf("error while storing token info: %s", err)
	}

	return id, nil
}

// RemoveTokenInfo allows to remove the TokenInfo
func (db *Db) RemoveBridgeTokenInfo(id string) error {
	query := `
		DELETE FROM tokens_info WHERE id = $1
	`
	_, err := db.SQL.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error while removing token info: %s", err)
	}

	return nil
}

// -------------------------------------------------------------------------------------------------------------------

// SaveTokenMetadata allows to save new TokenMetadata
func (db *Db) SaveBridgeTokenMetadata(tokenID uint64, name, symbol, uri string) (int64, error) {
	query := `
		INSERT INTO token_metadata(tokenID, name, symbol, uri) 
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (name) DO UPDATE
		SET uri = excluded.uri
		WHERE tokens_info.name <= excluded.name
	`

	var id int64
	err := db.SQL.QueryRow(query, tokenID, name, symbol, uri).Scan(&id)
	if err != nil {
		return emptyIndex, fmt.Errorf("error while storing token metadata: %s", err)
	}

	return id, nil
}

// RemoveTokenMetadata allows to remove the TokenMetadata
func (db *Db) RemoveBridgeTokenMetadata(id string) error {
	query := `
		DELETE FROM token_metadata WHERE id = $1
	`
	_, err := db.SQL.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error while removing token metadata: %s", err)
	}

	return nil
}

// -------------------------------------------------------------------------------------------------------------------

func (db *Db) SaveBridgeTokens(tokensInfoID, tokenMetadataID int64) error {
	query := `
		INSERT INTO tokens(tokens_info_id, token_metadata_id) 
		VALUES ($1, $2) RETURNING id
	`

	_, err := db.SQL.Exec(query, tokensInfoID, tokenMetadataID)
	if err != nil {
		return fmt.Errorf("error while removing token: %s", err)
	}
	return err
}

func (db *Db) RemoveBridgeTokens(tokenID uint64) error {
	query := `
		DELETE FROM tokens WHERE metadata_id = $1;
		DELETE FORM tokens_info WHERE id = $1;
		DELETE FORM token_metadata WHERE token_id = $1;
	`
	_, err := db.SQL.Exec(query, tokenID)
	if err != nil {
		return fmt.Errorf("error while removing token: %s", err)
	}

	return nil
}

// -------------------------------------------------------------------------------------------------------------------

func (db *Db) SaveBridgeTransaction(
	tx types.Transaction,
) error {
	query := `
		INSERT INTO transactions(
			deposit_chain_id, 
			deposit_tx_hash, 
			deposit_tx_index,
			deposit_block, 
			deposit_token, 
			amount,
			depositor,
			receiver,
			withdrawal_chain_id,
			withdrawal_tx_hash,
			withdrawal_token, 
			signature,
			isWrapped
	 	) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id
	`
	_, err := db.SQL.Exec(
		query,
		tx.DepositToken,
		tx.DepositTxHash,
		tx.DepositTxIndex,
		tx.DepositBlock,
		tx.DepositToken,
		tx.Amount,
		tx.Depositor,
		tx.Receiver,
		tx.WithdrawalChainId,
		tx.WithdrawalTxHash,
		tx.WithdrawalToken,
		tx.Signature,
		tx.IsWrapped,
	)
	if err != nil {
		return fmt.Errorf("error while storing transaction: %s", err)
	}

	return nil
}
