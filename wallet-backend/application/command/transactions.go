package command

import (
	"context"
	"math/big"

	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/notifications/protocolbuffer"
	"github.com/kowala-tech/wallet-backend/blockchain"
)

//GetTransactions represents the parameters needed to perform the use case for getting the transactions of a
//given account.
type GetTransactions struct {
	Address common.Address
	From    *big.Int
	To      *big.Int
}

//GetTransactionsHandler represents the use case of getting the transactions sent or received from a given account.
type GetTransactionsHandler struct {
	Client protocolbuffer.TransactionServiceClient
}

//Handle executes the use case of getting the transactions sent or received from a given account. It returns error
// or TransactionsResponse with the information.
func (h *GetTransactionsHandler) Handle(ctx context.Context, cmd GetTransactions) (*TransactionsResponse, error) {
	req := &protocolbuffer.GetTransactionsRequest{
		Account: cmd.Address.String(),
	}

	txsResp, err := h.Client.GetTransactions(ctx, req)
	if err != nil {
		return nil, err
	}

	txs := make([]*blockchain.Transaction, 0)
	for _, tx := range txsResp.Transactions {
		txs = append(
			txs,
			&blockchain.Transaction{
				Hash:        tx.Hash,
				From:        tx.From,
				To:          tx.To,
				Amount:      big.NewInt(tx.Amount),
				Timestamp:   big.NewInt(tx.Timestamp),
				BlockHeight: big.NewInt(tx.BlockHeight),
				GasUsed:     big.NewInt(tx.GasUsed),
				GasPrice:    big.NewInt(tx.GasPrice),
			},
		)
	}

	rangeIsSpecified := cmd.From != nil && cmd.To != nil
	if rangeIsSpecified {
		txs = filterTxsByRange(txs, cmd.From, cmd.To)
	}

	resp := &TransactionsResponse{
		Transactions: txs,
	}

	return resp, nil
}

func filterTxsByRange(txs []*blockchain.Transaction, from *big.Int, to *big.Int) []*blockchain.Transaction {
	filteredTransactions := make([]*blockchain.Transaction, 0)

	for _, tx := range txs {
		isBiggerOrEqualThanFrom := tx.BlockHeight.Cmp(from) >= 0
		isSmallerOrEqualThanTo := tx.BlockHeight.Cmp(to) <= 0

		if isBiggerOrEqualThanFrom && isSmallerOrEqualThanTo {
			filteredTransactions = append(filteredTransactions, tx)
		}
	}

	return filteredTransactions
}

//TransactionsResponse represents the response with the transactions sent or received from a given account.
type TransactionsResponse struct {
	Transactions []*blockchain.Transaction `json:"transactions"`
}
