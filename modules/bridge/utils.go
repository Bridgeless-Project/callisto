package bridge

import (
	"reflect"

	errorsmod "cosmossdk.io/errors"
	bridgetypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/ibc-go/v6/modules/light-clients/06-solomachine/types"
	v19 "github.com/forbole/bdjuno/v4/modules/bridge/migrations/v19"
	v21 "github.com/forbole/bdjuno/v4/modules/bridge/migrations/v21"
	v24 "github.com/forbole/bdjuno/v4/modules/bridge/migrations/v24"
)

func V19InsertChainToLatest(msg *v19.MsgInsertChain) *bridgetypes.MsgInsertChain {
	return &bridgetypes.MsgInsertChain{
		Creator: msg.Creator,
		Chain: bridgetypes.Chain{
			Id:            msg.Chain.Id,
			Type:          bridgetypes.ChainType(msg.Chain.Type),
			BridgeAddress: msg.Chain.BridgeAddress,
			Operator:      msg.Chain.Operator,
			Confirmations: 0,
			Name:          "",
		},
	}
}

func V19InsertTokenToLatest(msg *v19.MsgInsertToken) *bridgetypes.MsgInsertToken {
	info := make([]bridgetypes.TokenInfo, len(msg.Token.Info))

	for i, tokenInfo := range msg.Token.Info {
		info[i] = bridgetypes.TokenInfo{
			Address:             tokenInfo.Address,
			Decimals:            tokenInfo.Decimals,
			ChainId:             tokenInfo.ChainId,
			TokenId:             tokenInfo.TokenId,
			IsWrapped:           tokenInfo.IsWrapped,
			MinWithdrawalAmount: "0",
			CommissionRate:      msg.Token.CommissionRate,
		}
	}

	return &bridgetypes.MsgInsertToken{
		Creator: msg.Creator,
		Token: bridgetypes.Token{
			Id: msg.Token.Id,
			Metadata: bridgetypes.TokenMetadata{
				Name:    msg.Token.Metadata.Name,
				Symbol:  msg.Token.Metadata.Symbol,
				Uri:     msg.Token.Metadata.Uri,
				DexName: "",
			},
			Info: info,
		},
	}
}

func V19UpdateTokenToLatest(msg *v19.MsgUpdateToken) *bridgetypes.MsgUpdateToken {
	return &bridgetypes.MsgUpdateToken{
		Creator: msg.Creator,
		TokenId: msg.TokenId,
		Metadata: bridgetypes.TokenMetadata{
			Name:    msg.Metadata.Name,
			Symbol:  msg.Metadata.Symbol,
			Uri:     msg.Metadata.Uri,
			DexName: "",
		},
	}
}

func V19AddTokenInfoToLatest(msg *v19.MsgAddTokenInfo) *bridgetypes.MsgAddTokenInfo {
	return &bridgetypes.MsgAddTokenInfo{
		Creator: msg.Creator,
		Info: bridgetypes.TokenInfo{
			Address:             msg.Info.Address,
			Decimals:            msg.Info.Decimals,
			ChainId:             msg.Info.ChainId,
			TokenId:             msg.Info.TokenId,
			IsWrapped:           msg.Info.IsWrapped,
			MinWithdrawalAmount: "0",
			CommissionRate:      "0",
		},
	}
}

func V19SubmitTxsToLatest(msg *v19.MsgSubmitTransactions) *bridgetypes.MsgSubmitTransactions {
	txs := make([]bridgetypes.Transaction, len(msg.Transactions), len(msg.Transactions))
	for i, tx := range msg.Transactions {
		txs[i] = bridgetypes.Transaction{
			DepositChainId:    tx.DepositChainId,
			DepositTxHash:     tx.DepositTxHash,
			DepositTxIndex:    tx.DepositTxIndex,
			DepositBlock:      tx.DepositBlock,
			DepositToken:      tx.DepositToken,
			DepositAmount:     tx.DepositAmount,
			Depositor:         tx.Depositor,
			Receiver:          tx.Receiver,
			WithdrawalChainId: tx.WithdrawalChainId,
			WithdrawalTxHash:  tx.WithdrawalTxHash,
			WithdrawalToken:   tx.WithdrawalToken,
			Signature:         tx.Signature,
			IsWrapped:         tx.IsWrapped,
			WithdrawalAmount:  tx.WithdrawalAmount,
			CommissionAmount:  tx.CommissionAmount,
			TxData:            tx.TxData,
			ReferralId:        0,
		}
	}

	return &bridgetypes.MsgSubmitTransactions{
		Submitter:    msg.Submitter,
		Transactions: txs,
	}
}

func V21InsertTokenToLatest(msg *v21.MsgInsertToken) *bridgetypes.MsgInsertToken {
	info := make([]bridgetypes.TokenInfo, len(msg.Token.Info))

	for i, tokenInfo := range msg.Token.Info {
		info[i] = bridgetypes.TokenInfo{
			Address:             tokenInfo.Address,
			Decimals:            tokenInfo.Decimals,
			ChainId:             tokenInfo.ChainId,
			TokenId:             tokenInfo.TokenId,
			IsWrapped:           tokenInfo.IsWrapped,
			MinWithdrawalAmount: "0",
			CommissionRate:      msg.Token.CommissionRate,
		}
	}

	return &bridgetypes.MsgInsertToken{
		Creator: msg.Creator,
		Token: bridgetypes.Token{
			Id: msg.Token.Id,
			Metadata: bridgetypes.TokenMetadata{
				Name:    msg.Token.Metadata.Name,
				Symbol:  msg.Token.Metadata.Symbol,
				Uri:     msg.Token.Metadata.Uri,
				DexName: msg.Token.Metadata.DexName,
			},
			Info: info,
		},
	}
}

func V21AddTokenInfoToLatest(msg *v21.MsgAddTokenInfo) *bridgetypes.MsgAddTokenInfo {
	return &bridgetypes.MsgAddTokenInfo{
		Creator: msg.Creator,
		Info: bridgetypes.TokenInfo{
			Address:             msg.Info.Address,
			Decimals:            msg.Info.Decimals,
			ChainId:             msg.Info.ChainId,
			TokenId:             msg.Info.TokenId,
			IsWrapped:           msg.Info.IsWrapped,
			MinWithdrawalAmount: "0",
			CommissionRate:      "0",
		},
	}
}

func V24InsertTokenToLatest(msg *v24.MsgInsertToken) *bridgetypes.MsgInsertToken {
	info := make([]bridgetypes.TokenInfo, len(msg.Token.Info))

	for i, tokenInfo := range msg.Token.Info {
		info[i] = bridgetypes.TokenInfo{
			Address:             tokenInfo.Address,
			Decimals:            tokenInfo.Decimals,
			ChainId:             tokenInfo.ChainId,
			TokenId:             tokenInfo.TokenId,
			IsWrapped:           tokenInfo.IsWrapped,
			MinWithdrawalAmount: tokenInfo.MinWithdrawalAmount,
			CommissionRate:      msg.Token.CommissionRate,
		}
	}

	return &bridgetypes.MsgInsertToken{
		Creator: msg.Creator,
		Token: bridgetypes.Token{
			Id: msg.Token.Id,
			Metadata: bridgetypes.TokenMetadata{
				Name:    msg.Token.Metadata.Name,
				Symbol:  msg.Token.Metadata.Symbol,
				Uri:     msg.Token.Metadata.Uri,
				DexName: msg.Token.Metadata.DexName,
			},
			Info: info,
		},
	}
}

func V24AddTokenInfoToLatest(msg *v24.MsgAddTokenInfo) *bridgetypes.MsgAddTokenInfo {
	return &bridgetypes.MsgAddTokenInfo{
		Creator: msg.Creator,
		Info: bridgetypes.TokenInfo{
			Address:             msg.Info.Address,
			Decimals:            msg.Info.Decimals,
			ChainId:             msg.Info.ChainId,
			TokenId:             msg.Info.TokenId,
			IsWrapped:           msg.Info.IsWrapped,
			MinWithdrawalAmount: msg.Info.MinWithdrawalAmount,
			CommissionRate:      "0",
		},
	}
}
func isInList(val string, list []string) bool {
	for _, v := range list {
		if v == val {
			return true
		}
	}

	return false
}

func compareTxs(tx, tx2 bridgetypes.Transaction) error {
	txValue := reflect.ValueOf(tx)
	tx2Value := reflect.ValueOf(tx2)

	txTypes := reflect.TypeOf(tx)

	if txValue.NumField() != tx2Value.NumField() {
		return errorsmod.Wrap(types.ErrInvalidDataType, "transactions have different number of fields")
	}

	for i := 0; i < txValue.NumField(); i++ {
		if txTypes.Field(i).Name == "WithdrawalTxHash" {
			continue
		}

		if txValue.Field(i).Interface() != tx2Value.Field(i).Interface() {
			return errorsmod.Wrapf(types.ErrInvalidDataType, "field %s is different: %v != %v", txTypes.Field(i).Name, txValue.Field(i).Interface(), tx2Value.Field(i).Interface())
		}
	}

	return nil
}
