package types

import (
	"math/big"
	"time"

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
	Id              int            `db:"id"`
	ModuleAdmin     string         `db:"module_admin"`
	Parties         pq.StringArray `db:"parties"`
	TssThreshold    uint32         `db:"tss_threshold"`
	RelayerAccounts pq.StringArray `db:"relayer_accounts"`
}

type Transaction struct {
	Id                 int    `db:"id"`
	DepositChainId     string `db:"deposit_chain_id"`
	DepositTxHash      string `db:"deposit_tx_hash"`
	DepositTxIndex     uint64 `db:"deposit_tx_index"`
	DepositBlock       uint64 `db:"deposit_block"`
	DepositToken       string `db:"deposit_token"`
	DepositAmount      string `db:"deposit_amount"`
	Depositor          string `db:"depositor"`
	Receiver           string `db:"receiver"`
	WithdrawalChainId  string `db:"withdrawal_chain_id"`
	WithdrawalTxHash   string `db:"withdrawal_tx_hash"`
	WithdrawalToken    string `db:"withdrawal_token"`
	Signature          string `db:"signature"`
	IsWrapped          bool   `db:"is_wrapped"`
	WithdrawalAmount   string `db:"withdrawal_amount"`
	CommissionAmount   string `db:"commission_amount"`
	TxData             string `db:"tx_data"`
	ReferralId         uint32 `db:"referral_id"`
	TokenId            int    `db:"token_id"`
	DepositDecimals    uint32 `db:"deposit_decimals"`
	WithdrawalDecimals uint32 `db:"withdrawal_decimals"`
	MerkleRoot         string `db:"merkle_root"`

	CoreTxTimestamp time.Time `db:"core_tx_timestamp"`
}

type Referral struct {
	Id                uint32 `db:"id"`
	WithdrawalAddress string `db:"withdrawal_address"`
	CommissionRate    string `db:"commission_rate"`
}

type ReferralRewards struct {
	ReferralId         uint32     `db:"referral_id"`
	TokenId            uint64     `db:"token_id"`
	ToClaim            types.Coin `db:"to_claim"`
	TotalClaimedAmount types.Coin `db:"total_claimed_amount"`
}

type BridgeTokenVolume struct {
	Id               uint64   `db:"id"`
	DepositAmount    *big.Int `db:"deposit_amount"`
	WithdrawalAmount *big.Int `db:"withdrawal_amount"`
	CommissionAmount *big.Int `db:"commission_amount"`
	TokenId          uint64   `db:"token_id"`
	UpdatedAt        string   `db:"updated_at"`
}
type BridgeToken struct {
	MetadataId     uint64 `db:"metadata_id"`
	TokenInfoId    int64  `db:"tokens_info_id"`
	CommissionRate string `db:"commission_rate"`
}

type BridgeTokenMetadata struct {
	TokenId uint64 `db:"token_id"`
	Name    string `db:"name"`
	Symbol  string `db:"symbol"`
	Uri     string `db:"uri"`
	DexName string `db:"dex_name"`
}

type BridgeTokenInfo struct {
	Id                  int64  `db:"id"`
	Address             string `db:"address"`
	Decimals            uint64 `db:"decimals"`
	ChainId             string `db:"chain_id"`
	TokenId             uint64 `db:"token_id"`
	IsWrapped           bool   `db:"is_wrapped"`
	MinWithdrawalAmount string `db:"min_withdrawal_amount"`
	CommissionRate      string `db:"commission_rate"`
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
		ModuleAdmin:     params.ModuleAdmin,
		Parties:         parties,
		TssThreshold:    params.TssThreshold,
		RelayerAccounts: params.RelayerAccounts,
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
		ReferralId:        transaction.ReferralId,
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
	rawToClaim, ok := big.NewInt(0).SetString(rewards.ToClaim.Amount, 10)
	if !ok {
		rawToClaim = big.NewInt(0)
	}
	toClaim := cosmostypes.Coin{
		Denom:  rewards.ToClaim.Denom,
		Amount: cosmostypes.NewIntFromBigInt(rawToClaim),
	}

	rawTotalClaimedAmount, ok := big.NewInt(0).SetString(rewards.TotalClaimedAmount.Amount, 10)
	if !ok {
		rawTotalClaimedAmount = big.NewInt(0)
	}
	totalClaimedAmount := cosmostypes.Coin{
		Denom:  rewards.TotalClaimedAmount.Denom,
		Amount: cosmostypes.NewIntFromBigInt(rawTotalClaimedAmount),
	}
	return &bridgeTypes.ReferralRewards{
		ReferralId:         rewards.ReferralId,
		TokenId:            rewards.TokenId,
		ToClaim:            toClaim.String(),
		TotalClaimedAmount: totalClaimedAmount.String(),
	}
}
