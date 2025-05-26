package types

import (
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
	"github.com/lib/pq"
)

// TxSubmissions adapts core TransactionSubmissions type for db operations
type TxSubmissions struct {
	TxHash     string         `db:"tx_hash"`
	Submitters pq.StringArray `db:"submitters"`
}

// Params adapts core bridge.Params type for db operations
type Params struct {
	Id           int            `db:"id"`
	ModuleAdmin  string         `db:"module_admin"`
	Parties      pq.StringArray `db:"parties"`
	TssThreshold uint32         `db:"tss_threshold"`
}

type Transaction struct {
	Id                int    `db:"id"`
	DepositChainId    string `db:"deposit_chain_id"`
	DepositTxHash     string `db:"deposit_tx_hash"`
	DepositTxIndex    uint64 `db:"deposit_tx_index"`
	DepositBlock      uint64 `db:"deposit_block"`
	DepositToken      string `db:"deposit_token"`
	DepositAmount     string `db:"deposit_amount"`
	Depositor         string `db:"depositor"`
	Receiver          string `db:"receiver"`
	WithdrawalChainId string `db:"withdrawal_chain_id"`
	WithdrawalTxHash  string `db:"withdrawal_tx_hash"`
	WithdrawalToken   string `db:"withdrawal_token"`
	Signature         string `db:"signature"`
	IsWrapped         bool   `db:"is_wrapped"`
	WithdrawalAmount  string `db:"withdrawal_amount"`
	CommissionAmount  string `db:"commission_amount"`
	TxData            string `db:"tx_data"`
}

func ToTransactionSubmissions(txSubmissions TxSubmissions) *types.TransactionSubmissions {
	return &types.TransactionSubmissions{
		TxHash:     txSubmissions.TxHash,
		Submitters: txSubmissions.Submitters,
	}
}

func ToBridgeParams(params Params) *types.Params {
	var parties []*types.Party
	for _, party := range params.Parties {
		parties = append(parties, &types.Party{
			Address: party,
		})
	}
	return &types.Params{
		ModuleAdmin:  params.ModuleAdmin,
		Parties:      parties,
		TssThreshold: params.TssThreshold,
	}
}

func ToBridgeTransaction(transaction Transaction) *types.Transaction {
	return &types.Transaction{
		DepositChainId:    transaction.DepositChainId,
		DepositTxHash:     transaction.DepositTxHash,
		DepositTxIndex:    transaction.DepositTxIndex,
		DepositBlock:      transaction.DepositBlock,
		DepositToken:      transaction.DepositToken,
		DepositAmount:     transaction.DepositAmount,
		Depositor:         transaction.Depositor,
		Receiver:          transaction.Receiver,
		WithdrawalChainId: transaction.WithdrawalChainId,
		WithdrawalTxHash:  transaction.WithdrawalTxHash,
		WithdrawalToken:   transaction.WithdrawalToken,
		Signature:         transaction.Signature,
		IsWrapped:         transaction.IsWrapped,
		WithdrawalAmount:  transaction.WithdrawalAmount,
		CommissionAmount:  transaction.CommissionAmount,
		TxData:            transaction.TxData,
	}
}
