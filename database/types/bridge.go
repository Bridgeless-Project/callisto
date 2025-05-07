package types

import "github.com/hyle-team/bridgeless-core/v12/x/bridge/types"

// TxSubmissions adapts core TransactionSubmissions type for db operations
type TxSubmissions struct {
	TxHash     string   `db:"tx_hash"`
	Submitters []string `db:"submitters"`
}

// TxSubmissions adapts core bridge.Params type for db operations
type Params struct {
	Id           int    `db:"id"`
	ModuleAdmin  string `db:"module_admin"`
	Parties      string `db:"parties"`
	TssThreshold int    `db:"tss_threshold"`
}

func ToTransactionSubmissions(txSubmissions TxSubmissions) *types.TransactionSubmissions {
	return &types.TransactionSubmissions{
		TxHash:     txSubmissions.TxHash,
		Submitters: txSubmissions.Submitters,
	}
}
