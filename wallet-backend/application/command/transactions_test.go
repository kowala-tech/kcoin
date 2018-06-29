package command

import (
	"context"
	"testing"

	"math/big"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/notifications/protocolbuffer"
	"github.com/kowala-tech/kcoin/notifications/protocolbuffer/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetTransactions(t *testing.T) {
	addr := common.HexToAddress("0xd6e579085c82329c89fca7a9f012be59028ed53f")

	t.Run("Get transactions without range selection", func(t *testing.T) {
		mockedClient := &mocks.TransactionServiceClient{}

		handl := GetTransactionsHandler{
			Client: mockedClient,
		}

		cmd := GetTransactions{
			Address: addr,
		}

		req := &protocolbuffer.GetTransactionsRequest{
			Account: addr.String(),
		}

		mockedResponse := &protocolbuffer.GetTransactionsReply{
			Transactions: []*protocolbuffer.Transaction{
				{
					From: addr.String(),
					To:   "0xdbdfdbce9a34c3ac5546657f651146d88d1b639a",
				},
			},
		}

		mockedClient.On("GetTransactions", context.Background(), req).
			Return(mockedResponse, nil)

		resp, err := handl.Handle(context.Background(), cmd)
		if err != nil {
			t.Fatalf("%v", err)
		}

		assert.Len(t, resp.Transactions, 1)

		tx := resp.Transactions[0]

		assert.Equal(t, addr.String(), tx.From)
		assert.Equal(t, "0xdbdfdbce9a34c3ac5546657f651146d88d1b639a", tx.To)
	})

	t.Run("We can specify a range from block to block", func(t *testing.T) {
		mockedClient := &mocks.TransactionServiceClient{}

		handl := GetTransactionsHandler{
			Client: mockedClient,
		}

		cmd := GetTransactions{
			Address: addr,
			From:    big.NewInt(100),
			To:      big.NewInt(150),
		}

		req := &protocolbuffer.GetTransactionsRequest{
			Account: addr.String(),
		}

		mockedResponse := &protocolbuffer.GetTransactionsReply{
			Transactions: []*protocolbuffer.Transaction{
				{
					From:        addr.String(),
					To:          "0xdbdfdbce9a34c3ac5546657f651146d88d1b639a",
					BlockHeight: 99,
				},
				{
					From:        addr.String(),
					To:          "0xdbdfdbce9a34c3ac5546657f651146d88d1b639a",
					BlockHeight: 102,
				},
				{
					From:        addr.String(),
					To:          "0xdbdfdbce9a34c3ac5546657f651146d88d1b639a",
					BlockHeight: 151,
				},
			},
		}

		mockedClient.On("GetTransactions", context.Background(), req).
			Return(mockedResponse, nil)

		resp, err := handl.Handle(context.Background(), cmd)
		if err != nil {
			t.Fatalf("%v", err)
		}

		assert.Len(t, resp.Transactions, 1)
		tx := resp.Transactions[0]

		assert.Equal(t, addr.String(), tx.From)
		assert.Equal(t, "0xdbdfdbce9a34c3ac5546657f651146d88d1b639a", tx.To)
		assert.Equal(t, big.NewInt(102), tx.BlockHeight)
	})
}
