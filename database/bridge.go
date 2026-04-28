package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/big"

	bridgeTypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/forbole/bdjuno/v4/database/types"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

// SaveChain allows to save new Chain
func (db *Db) SaveBridgeChain(id string, chainType int32, bridgeAddress string, operator string, confirmations uint32, name string) error {
	query := `
		INSERT INTO bridge_chains(id, chain_type, bridge_address, operator, confirmations,name) 
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (id) DO UPDATE
		SET chain_type = excluded.chain_type,
			bridge_address = excluded.bridge_address,
			operator = excluded.operator,
		    confirmations = excluded.confirmations, 
			name = excluded.name;
	`
	_, err := db.SQL.Exec(query, id, chainType, bridgeAddress, operator, confirmations, name)
	if err != nil {
		return fmt.Errorf("error while storing chain: %s", err)
	}

	return nil
}

func (db *Db) UpdateBridgeChainAddress(id, bridgeAddress string) error {
	query := `
		UPDATE bridge_chains
		SET bridge_address = $2
		WHERE id = $1
	`
	_, err := db.SQL.Exec(query, id, bridgeAddress)
	if err != nil {
		return fmt.Errorf("error while updating chain bridge address: %s", err)
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
func (db *Db) SaveBridgeTokenInfo(address string, decimals uint64, chainID string, tokenID uint64, isWrapped bool, minWithdrawalAmount string, commissionRate string) (int64, error) {
	query := `
		INSERT INTO bridge_tokens_info(address, decimals, chain_id, token_id, is_wrapped, min_withdrawal_amount, commission_rate) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (address, chain_id) DO UPDATE
		SET chain_id = excluded.chain_id,
			token_id = excluded.token_id,
			is_wrapped = excluded.is_wrapped
		RETURNING id
	`

	var id int64
	err := db.SQL.QueryRow(query, address, decimals, chainID, tokenID, isWrapped, minWithdrawalAmount, commissionRate).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error while storing token info: %s", err)
	}

	return id, nil
}

// RemoveTokenInfo allows to remove the TokenInfo
func (db *Db) RemoveBridgeTokenInfo(id uint64) error {
	query := `
		DELETE FROM bridge_tokens_info WHERE id = $1
	`
	_, err := db.SQL.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error while removing token info: %s", err)
	}

	return nil
}

func (db *Db) GetTokenInfo(address, chainId string) (*types.BridgeTokenInfo, error) {
	query := `SELECT * FROM bridge_tokens_info WHERE address = $1 AND chain_id = $2`

	var tokenInfo types.BridgeTokenInfo
	if err := db.Sqlx.Get(&tokenInfo, query, address, chainId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error while getting token info: %s", err)
	}

	return &tokenInfo, nil
}

// -------------------------------------------------------------------------------------------------------------------

// SaveTokenMetadata allows to save new TokenMetadata
func (db *Db) SaveBridgeTokenMetadata(tokenID uint64, name, symbol, uri, dexName string) error {
	query := `
		INSERT INTO bridge_token_metadata(token_id, name, symbol, uri, dex_name) 
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (token_id) DO UPDATE
		SET name = excluded.name,
			symbol = excluded.symbol,
			uri = excluded.uri,
		    dex_name = excluded.dex_name;                     
	`

	_, err := db.SQL.Exec(query, tokenID, name, symbol, uri, dexName)
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

func (db *Db) GetBridgeTokenMetadata(tokenID uint64) (*types.BridgeTokenMetadata, error) {
	query := `SELECT * FROM bridge_token_metadata WHERE token_id = $1`

	var tokenMetadata types.BridgeTokenMetadata
	if err := db.Sqlx.Get(&tokenMetadata, query, tokenID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error while getting token metadata: %s", err)
	}

	return &tokenMetadata, nil
}

// -------------------------------------------------------------------------------------------------------------------

// SaveBridgeTokens allows to save new Tokens
func (db *Db) SaveBridgeToken(tokensInfoID int64, tokenMetadataID uint64) error {
	query := `
		INSERT INTO bridge_tokens(tokens_info_id, metadata_id) 
		VALUES ($1, $2) 
		ON CONFLICT (tokens_info_id, metadata_id) DO NOTHING

	`

	_, err := db.SQL.Exec(query, tokensInfoID, tokenMetadataID)
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

func (db *Db) GetBridgeToken(tokenID, metadataId uint64) (*types.BridgeToken, error) {
	query := `SELECT * FROM bridge_tokens WHERE tokens_info_id = $1 AND metadata_id = $2`

	var tokenInfo types.BridgeToken
	if err := db.Sqlx.Get(&tokenInfo, query, tokenID, metadataId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error while getting token info: %s", err)
	}

	return &tokenInfo, nil
}

// -------------------------------------------------------------------------------------------------------------------

func (db *Db) SaveBridgeTransaction(
	tx bridgeTypes.Transaction,
	timestamp string,
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
		    tx_data,
		    core_tx_timestamp,
		    referral_id,
			merkle_root) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17,$18,$19)
		RETURNING id
	`
	_, err := db.SQL.Exec(
		query,
		tx.DepositChainId,
		tx.DepositTxHash,
		big.NewInt(0).SetUint64(tx.DepositTxIndex).String(),
		big.NewInt(0).SetUint64(tx.DepositBlock).String(),
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
		timestamp,
		tx.ReferralId,
		tx.MerkleProof,
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

func (db *Db) GetBridgeTransaction(depositChainId string, depositTxHash string, depositTxNonce uint64) (*bridgeTypes.Transaction, error) {
	var tx types.Transaction
	err := db.Sqlx.Get(&tx, `SELECT * FROM bridge_transactions WHERE deposit_chain_id = $1 AND deposit_tx_hash = $2 AND deposit_tx_index = $3`, depositChainId, depositTxHash, depositTxNonce)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error while getting transaction: %s", err)
	}

	return types.ToBridgeTransaction(tx), nil
}

func (db *Db) RemoveBridgeTransaction(depositChainId string, depositTxHash string, depositTxNonce uint64) error {
	query := `
		DELETE FROM bridge_transactions WHERE deposit_chain_id = $1 AND deposit_tx_hash = $2 AND deposit_tx_index = $3
	`
	_, err := db.SQL.Exec(query, depositChainId, depositTxHash, depositTxNonce)
	if err != nil {
		return fmt.Errorf("error while removing transaction: %s", err)
	}

	return nil
}

func (db *Db) UpdateTransactionWithdrawalTxHash(depositChainId string, depositTxHash string, depositTxNonce uint64, withdrawalTxHash string) error {
	query := `
		UPDATE bridge_transactions
		SET withdrawal_tx_hash = $4
		WHERE deposit_chain_id = $1
		  AND deposit_tx_hash = $2
		  AND deposit_tx_index = $3
	`

	_, err := db.SQL.Exec(query, depositChainId, depositTxHash, depositTxNonce, withdrawalTxHash)
	if err != nil {
		return fmt.Errorf("error while updating withdrawal transaction hash: %s", err)
	}

	return nil
}

// -------------------------------------------------------------------------------------------------------------------

func (db *Db) SaveBridgeTransactionSubmissions(txSubmissions *bridgeTypes.Submissions) error {
	query := `INSERT INTO bridge_transaction_submissions (tx_hash,submitters) VALUES ($1, $2)
				ON CONFLICT (tx_hash) DO UPDATE
				SET submitters = excluded.submitters
				`

	_, err := db.SQL.Exec(query, txSubmissions.Hash, pq.Array(txSubmissions.Submitters))
	if err != nil {
		return fmt.Errorf("error while storing transaction submissions: %s", err)
	}

	return nil
}

func (db *Db) GetBridgeTransactionSubmissions(txHash string) (*bridgeTypes.Submissions, error) {
	var txSubmissions []types.TxSubmissions
	err := db.Sqlx.Select(&txSubmissions, `SELECT * FROM bridge_transaction_submissions WHERE tx_hash = $1`, txHash)

	if errors.Is(err, sql.ErrNoRows) || len(txSubmissions) == 0 {

		return &bridgeTypes.Submissions{
			Hash:       "",
			Submitters: nil,
		}, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error while getting transaction submissions: %s", err)
	}

	return types.ToTransactionSubmissions(txSubmissions[0]), nil
}

func (db *Db) RemoveBridgeTransactionSubmissions(txHash string) error {
	query := `
		DELETE FROM bridge_transaction_submissions WHERE tx_hash = $1
	`
	_, err := db.SQL.Exec(query, txHash)
	if err != nil {
		return fmt.Errorf("error while removing transaction submissions: %s", err)
	}

	return nil
}

func (d *Db) SetBridgeTransactionTokenId(tx bridgeTypes.Transaction) (uint64, error) {
	query := `
	UPDATE bridge_transactions AS bt
	SET token_id = bti.token_id
	FROM bridge_tokens_info AS bti
	WHERE
		bt.deposit_chain_id = bti.chain_id
		AND lower(bt.deposit_token) = lower(bti.address)
	
		AND bt.deposit_chain_id = $1
		AND lower(bt.deposit_tx_hash) = lower($2)
		AND bt.deposit_tx_index = $3
	
	RETURNING bt.token_id;
`
	var tokenId uint64
	err := d.SQL.Get(&tokenId, query, tx.DepositChainId, tx.DepositTxHash, tx.DepositTxIndex)
	if err != nil {
		return 0, fmt.Errorf("error while setting transaction token ID: %s", err)
	}

	return tokenId, nil
}

func (d *Db) SetBridgeTransactionDecimals(tx bridgeTypes.Transaction) (depositDecimals, withdrawalDecimals uint64, err error) {
	query := `
	UPDATE bridge_transactions AS bt
	SET deposit_decimals = bti.decimals
	FROM bridge_tokens_info AS bti
	WHERE
		bt.deposit_chain_id = bti.chain_id
		AND lower(bt.deposit_token) = lower(bti.address)
	
		AND bt.deposit_chain_id = $1
		AND lower(bt.deposit_tx_hash) = lower($2)
		AND bt.deposit_tx_index = $3
	
	RETURNING bt.deposit_decimals;
`

	err = d.SQL.Get(&depositDecimals, query, tx.DepositChainId, tx.DepositTxHash, tx.DepositTxIndex)
	if err != nil {
		return depositDecimals, withdrawalDecimals,
			fmt.Errorf("error while setting transaction deposit decimals: %s", err)
	}

	query = `
	UPDATE bridge_transactions AS bt
	SET withdrawal_decimals = bti.decimals
	FROM bridge_tokens_info AS bti
	WHERE
		bt.withdrawal_chain_id = bti.chain_id
		AND lower(bt.withdrawal_token) = lower(bti.address)
	
		AND bt.deposit_chain_id = $1
		AND lower(bt.deposit_tx_hash) = lower($2)
		AND bt.deposit_tx_index = $3
	
	RETURNING bt.withdrawal_decimals;
`
	err = d.SQL.Get(&withdrawalDecimals, query, tx.DepositChainId, tx.DepositTxHash, tx.DepositTxIndex)
	if err != nil {
		return depositDecimals, withdrawalDecimals,
			fmt.Errorf("error while setting transaction withdrawal decimals: %s", err)
	}

	return depositDecimals, withdrawalDecimals, nil
}

// -------------------------------------------------------------------------------------------------------------------

func (db *Db) SaveBridgeParams(params *bridgeTypes.Params) error {
	query := `INSERT INTO bridge_params (id, module_admin,parties,tss_threshold,relayer_accounts,epoch_id,supporting_time) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
				ON CONFLICT (id) DO UPDATE
				SET module_admin = excluded.module_admin,
				parties = excluded.parties,
				tss_threshold = excluded.tss_threshold,
				relayer_accounts = excluded.relayer_accounts,
				epoch_id = excluded.epoch_id,
				supporting_time = excluded.supporting_time`

	var parties []string
	for _, party := range params.Parties {
		parties = append(parties, party.Address)
	}
	_, err := db.SQL.Exec(query,
		1,
		params.ModuleAdmin,
		pq.StringArray(parties),
		int(params.TssThreshold),
		pq.StringArray(params.RelayerAccounts),
		params.Epoch,
		params.SupportingTime)
	if err != nil {
		return fmt.Errorf("error while storing bridge_params: %s", err)
	}

	return nil
}

func (db *Db) SaveBridgeEpoch(epoch *bridgeTypes.Epoch) error {
	query := `
		INSERT INTO bridge_epochs (
			id, status, start_block, end_block, parties, tss_threshold, tss_info, finalized_block
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (id) DO UPDATE
		SET status = excluded.status,
			start_block = excluded.start_block,
			end_block = excluded.end_block,
			parties = excluded.parties,
			tss_threshold = excluded.tss_threshold,
			tss_info = excluded.tss_info,
			finalized_block = excluded.finalized_block
	`

	var parties []string
	for _, party := range epoch.Parties {
		parties = append(parties, party.Address)
	}
	tssInfo, err := json.Marshal(epoch.TssInfo)
	if err != nil {
		return fmt.Errorf("error while marshaling epoch tss info: %s", err)
	}

	_, err = db.SQL.Exec(
		query,
		epoch.Id,
		int32(epoch.Status),
		epoch.StartBlock,
		epoch.EndBlock,
		pq.StringArray(parties),
		epoch.TssThreshold,
		tssInfo,
		epoch.FinalizedBlock,
	)
	if err != nil {
		return fmt.Errorf("error while storing bridge epoch: %s", err)
	}

	return nil
}

func (db *Db) GetBridgeEpoch(epochId uint32) (*bridgeTypes.Epoch, error) {
	var epoch types.BridgeEpoch
	err := db.Sqlx.Get(&epoch, `SELECT * FROM bridge_epochs WHERE id = $1`, epochId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error while getting bridge epoch: %s", err)
	}

	return types.ToBridgeEpoch(epoch)
}

func (db *Db) SaveBridgeEpochChainSignatures(signatures *bridgeTypes.EpochChainSignatures) error {
	query := `
		INSERT INTO bridge_epoch_chain_signatures (
			epoch_id,
			chain_type,
			added_mod,
			added_epoch_id,
			added_signature,
			added_new_signer,
			added_start_time,
			added_end_time,
			added_nonce,
			removed_mod,
			removed_epoch_id,
			removed_signature,
			removed_new_signer,
			removed_start_time,
			removed_end_time,
			removed_nonce
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
		ON CONFLICT (epoch_id, chain_type) DO UPDATE
		SET added_mod = excluded.added_mod,
			added_epoch_id = excluded.added_epoch_id,
			added_signature = excluded.added_signature,
			added_new_signer = excluded.added_new_signer,
			added_start_time = excluded.added_start_time,
			added_end_time = excluded.added_end_time,
			added_nonce = excluded.added_nonce,
			removed_mod = excluded.removed_mod,
			removed_epoch_id = excluded.removed_epoch_id,
			removed_signature = excluded.removed_signature,
			removed_new_signer = excluded.removed_new_signer,
			removed_start_time = excluded.removed_start_time,
			removed_end_time = excluded.removed_end_time,
			removed_nonce = excluded.removed_nonce
	`

	_, err := db.SQL.Exec(
		query,
		signatures.EpochId,
		int32(signatures.ChainType),
		epochSignatureMod(signatures.AddedSignature),
		epochSignatureEpochID(signatures.AddedSignature),
		epochSignatureValue(signatures.AddedSignature),
		epochSignatureNewSigner(signatures.AddedSignature),
		epochSignatureStartTime(signatures.AddedSignature),
		epochSignatureEndTime(signatures.AddedSignature),
		epochSignatureNonce(signatures.AddedSignature),
		epochSignatureMod(signatures.RemovedSignature),
		epochSignatureEpochID(signatures.RemovedSignature),
		epochSignatureValue(signatures.RemovedSignature),
		epochSignatureNewSigner(signatures.RemovedSignature),
		epochSignatureStartTime(signatures.RemovedSignature),
		epochSignatureEndTime(signatures.RemovedSignature),
		epochSignatureNonce(signatures.RemovedSignature),
	)
	if err != nil {
		return fmt.Errorf("error while storing bridge epoch chain signatures: %s", err)
	}

	return nil
}

func epochSignatureMod(signature *bridgeTypes.EpochSignature) interface{} {
	if signature == nil {
		return nil
	}
	return int32(signature.Mod)
}

func epochSignatureEpochID(signature *bridgeTypes.EpochSignature) interface{} {
	if signature == nil {
		return nil
	}
	return signature.EpochId
}

func epochSignatureValue(signature *bridgeTypes.EpochSignature) interface{} {
	if signature == nil {
		return nil
	}
	return signature.Signature
}

func epochSignatureNewSigner(signature *bridgeTypes.EpochSignature) interface{} {
	if signature == nil || signature.Data == nil {
		return nil
	}
	return signature.Data.NewSigner
}

func epochSignatureStartTime(signature *bridgeTypes.EpochSignature) interface{} {
	if signature == nil || signature.Data == nil {
		return nil
	}
	return signature.Data.StartTime
}

func epochSignatureEndTime(signature *bridgeTypes.EpochSignature) interface{} {
	if signature == nil || signature.Data == nil {
		return nil
	}
	return signature.Data.EndTime
}

func epochSignatureNonce(signature *bridgeTypes.EpochSignature) interface{} {
	if signature == nil || signature.Data == nil {
		return nil
	}
	return signature.Data.Nonce
}

func (db *Db) SaveBridgeEpochPubKey(epochId uint32, pubkey string) error {
	query := `
		INSERT INTO bridge_epoch_pubkeys (epoch_id, pubkey)
		VALUES ($1, $2)
		ON CONFLICT (epoch_id) DO UPDATE
		SET pubkey = excluded.pubkey
	`
	_, err := db.SQL.Exec(query, epochId, pubkey)
	if err != nil {
		return fmt.Errorf("error while storing bridge epoch pubkey: %s", err)
	}

	return nil
}

func (db *Db) GetBridgeEpochPubKeySubmissions(epochId uint32, hash string) (*types.BridgeEpochSubmissions, error) {
	return db.getBridgeEpochSubmissions("bridge_epoch_pubkey_submissions", epochId, hash)
}

func (db *Db) SaveBridgeEpochPubKeySubmissions(submissions *types.BridgeEpochSubmissions) error {
	return db.saveBridgeEpochSubmissions("bridge_epoch_pubkey_submissions", submissions)
}

func (db *Db) GetBridgeEpochSignatureSubmissions(epochId uint32, hash string) (*types.BridgeEpochSubmissions, error) {
	return db.getBridgeEpochSubmissions("bridge_epoch_signature_submissions", epochId, hash)
}

func (db *Db) SaveBridgeEpochSignatureSubmissions(submissions *types.BridgeEpochSubmissions) error {
	return db.saveBridgeEpochSubmissions("bridge_epoch_signature_submissions", submissions)
}

func (db *Db) getBridgeEpochSubmissions(table string, epochId uint32, hash string) (*types.BridgeEpochSubmissions, error) {
	var submissions types.BridgeEpochSubmissions
	err := db.Sqlx.Get(&submissions, fmt.Sprintf(`SELECT * FROM %s WHERE epoch_id = $1 AND hash = $2`, table), epochId, hash)
	if errors.Is(err, sql.ErrNoRows) {
		return &types.BridgeEpochSubmissions{
			EpochId: epochId,
			Hash:    hash,
		}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error while getting bridge epoch submissions: %s", err)
	}

	return &submissions, nil
}

func (db *Db) saveBridgeEpochSubmissions(table string, submissions *types.BridgeEpochSubmissions) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (epoch_id, hash, submitters)
		VALUES ($1, $2, $3)
		ON CONFLICT (epoch_id, hash) DO UPDATE
		SET submitters = excluded.submitters
	`, table)
	_, err := db.SQL.Exec(query, submissions.EpochId, submissions.Hash, pq.Array(submissions.Submitters))
	if err != nil {
		return fmt.Errorf("error while storing bridge epoch submissions: %s", err)
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

// -------------------------------------------------------------------------------------------------------------------

func (db *Db) SaveBridgeReferral(referral *bridgeTypes.Referral) error {
	query := `INSERT INTO referral (id, withdrawal_address, commission_rate) VALUES ($1, $2, $3)
				ON CONFLICT (id) DO UPDATE
				SET withdrawal_address = excluded.withdrawal_address,
				commission_rate = excluded.commission_rate,
			  `

	_, err := db.SQL.Exec(query, referral.Id, referral.WithdrawalAddress, referral.CommissionRate)
	if err != nil {
		return fmt.Errorf("error while storing referral: %s", err)
	}

	return nil
}

func (db *Db) GetBridgeReferrals() ([]bridgeTypes.Referral, error) {
	query := `SELECT * FROM referral`
	var refs []bridgeTypes.Referral
	err := db.Sqlx.Select(&refs, query)
	if err != nil {
		return nil, fmt.Errorf("error while getting referrals: %s", err)
	}

	return refs, nil
}

func (db *Db) GetBridgeReferralById(referralId uint32) (*bridgeTypes.Referral, error) {
	if referralId == 0 {
		return nil, fmt.Errorf("referral id cannot be zero")
	}

	query := `SELECT * FROM referral WHERE id = $1`
	var ref bridgeTypes.Referral
	err := db.Sqlx.Get(&ref, query, referralId)
	if err != nil {
		return nil, fmt.Errorf("error while getting referral by id: %s", err)
	}

	return &ref, nil
}

func (db *Db) RemoveBridgeReferral(referralId uint32) error {
	if referralId == 0 {
		return fmt.Errorf("referral id cannot be zero")
	}

	query := `DELETE FROM referral WHERE id = $1`
	_, err := db.SQL.Exec(query, referralId)
	if err != nil {
		return fmt.Errorf("error while removing referral: %s", err)
	}

	return nil
}

// -------------------------------------------------------------------------------------------------------------------

func (db *Db) SaveBridgeReferralRewards(rewards *bridgeTypes.ReferralRewards) error {
	if rewards == nil {
		return fmt.Errorf("rewards cannot be nil")
	}

	query := `INSERT INTO referral_rewards (referral_id,  token_id, total_claimed_amount, to_claim) VALUES ($1,$2, $3, $4)`
	_, err := db.SQL.Exec(query, rewards.ReferralId, rewards.TokenId, rewards.TotalClaimedAmount, rewards.ToClaim)
	if err != nil {
		return fmt.Errorf("error while storing referral rewards: %s", err)
	}

	return nil
}

func (db *Db) GetBridgeReferralRewards() ([]bridgeTypes.ReferralRewards, error) {
	query := `SELECT * FROM referral_rewards`
	var rewards []bridgeTypes.ReferralRewards
	err := db.Sqlx.Select(&rewards, query)
	if err != nil {
		return nil, fmt.Errorf("error while getting referral rewards: %s", err)
	}

	return rewards, nil
}

func (db *Db) GetBridgeReferralRewardsByReferralAndTokenIds(referralId uint32, tokenId uint64) (*bridgeTypes.ReferralRewards, error) {
	if referralId == 0 {
		return nil, fmt.Errorf("referral id cannot be zero")
	}
	if tokenId == 0 {
		return nil, fmt.Errorf("token id cannot be zero")
	}

	query := `SELECT * FROM referral_rewards WHERE referral_id = $1 AND token_id = $2`
	var referralRewards *bridgeTypes.ReferralRewards
	err := db.SQL.Select(referralRewards, query, referralId, tokenId)
	if err != nil {
		return nil, fmt.Errorf("error while getting referral rewards by referral and token ids: %s", err)
	}

	return referralRewards, nil
}

func (db *Db) RemoveBridgeReferralRewards(referralId uint32, tokenId uint64) error {
	if referralId == 0 {
		return fmt.Errorf("referral id cannot be zero")
	}
	if tokenId == 0 {
		return fmt.Errorf("token id cannot be zero")
	}

	query := `DELETE FROM referral_rewards WHERE referral_id = $1 AND token_id = $2`
	_, err := db.SQL.Exec(query, referralId, tokenId)
	if err != nil {
		return fmt.Errorf("error while removing referral rewards: %s", err)
	}

	return nil
}

// -------------------------------------------------------------------------------------------------------------------

func (db *Db) SetTokenVolume(volume *types.BridgeTokenVolume) error {
	query := `INSERT INTO bridge_tokens_volumes (
            		deposit_amount,
                	withdrawal_amount,
                    commission_amount,
                	token_id,
                    updated_at
                    ) VALUES ($1, $2, $3, $4, $5)
			ON CONFLICT (token_id,updated_at) DO UPDATE
			SET deposit_amount = bridge_tokens_volumes.deposit_amount + excluded.deposit_amount,
				withdrawal_amount = bridge_tokens_volumes.withdrawal_amount + excluded.withdrawal_amount,
				commission_amount = bridge_tokens_volumes.commission_amount + excluded.commission_amount,
				updated_at = excluded.updated_at
                    `

	_, err := db.SQL.Exec(query,
		volume.DepositAmount.String(),
		volume.WithdrawalAmount.String(),
		volume.CommissionAmount.String(),
		volume.TokenId,
		volume.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("error while storing token volume: %s", err)
	}

	return nil
}

func (db *Db) GetNativeDecimals(tokenId uint64) (uint64, error) {
	query := `
		SELECT decimals
		FROM bridge_tokens_info
		WHERE token_id = $1
		  AND is_wrapped = false
	`

	var decimals uint64
	err := db.Sqlx.Get(&decimals, query, tokenId)
	if err != nil {

		return decimals, fmt.Errorf("error while getting native decimals: %s", err)
	}

	return decimals, nil
}
