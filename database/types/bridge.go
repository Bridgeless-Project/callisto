package types

import (
	"math/big"

	bridgeTypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/forbole/bdjuno/v4/modules/actions/types"
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

type Referral struct {
	Id                uint32 `db:"id"`
	WithdrawalAddress string `db:"withdrawal_address"`
	CommissionRate    int32  `db:"commission_rate"`
}

type ReferralRewards struct {
	ReferralId           uint32     `db:"referral_id"`
	TokenId              uint64     `db:"token_id"`
	ToClaim              types.Coin `db:"to_claim"`
	TotalCollectedAmount types.Coin `db:"total_collected_amount"`
}

func ToTransactionSubmissions(txSubmissions TxSubmissions) *bridgeTypes.TransactionSubmissions {
	return &bridgeTypes.TransactionSubmissions{
		TxHash:     txSubmissions.TxHash,
		Submitters: txSubmissions.Submitters,
	}
}

func ToBridgeParams(params Params) *bridgeTypes.Params {
	var parties []*bridgeTypes.Party
	for _, party := range params.Parties {
		parties = append(parties, &bridgeTypes.Party{
			Address: party,
		})
	}
	return &bridgeTypes.Params{
		ModuleAdmin:  params.ModuleAdmin,
		Parties:      parties,
		TssThreshold: params.TssThreshold,
	}
}

func ToBridgeTransaction(transaction Transaction) *bridgeTypes.Transaction {
	return &bridgeTypes.Transaction{
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

func ToReferral(referral Referral) *bridgeTypes.Referral {
	return &bridgeTypes.Referral{
		Id:                referral.Id,
		WithdrawalAddress: referral.WithdrawalAddress,
		CommissionRate:    referral.CommissionRate,
	}
}

func ToReferralRewards(rewards ReferralRewards) *bridgeTypes.ReferralRewards {
	rawToClaim, found := big.NewInt(0).SetString(rewards.ToClaim.Amount, 10)
	if !found {
		rawToClaim = big.NewInt(0)
	}
	toClaim := cosmostypes.Coin{
		Denom:  rewards.ToClaim.Denom,
		Amount: cosmostypes.NewIntFromBigInt(rawToClaim),
	}

	rawToTotalCollectedAmount, found := big.NewInt(0).SetString(rewards.TotalCollectedAmount.Amount, 10)
	if !found {
		rawToTotalCollectedAmount = big.NewInt(0)
	}
	toTotalColectedAmount := cosmostypes.Coin{
		Denom:  rewards.TotalCollectedAmount.Denom,
		Amount: cosmostypes.NewIntFromBigInt(rawToTotalCollectedAmount),
	}
	return &bridgeTypes.ReferralRewards{
		ReferralId:           rewards.ReferralId,
		TokenId:              rewards.TokenId,
		ToClaim:              toClaim,
		TotalCollectedAmount: toTotalColectedAmount,
	}
}
