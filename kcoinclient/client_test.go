package kcoinclient

import (
	"context"
	"testing"

	"math/big"

	"fmt"

	"github.com/kowala-tech/kcoin"
	"github.com/kowala-tech/kcoin/kcoinclient/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestClient_BlockNumber(t *testing.T) {
	client := Client{}
	ctx := context.Background()

	t.Run("It returns a big integer", func(t *testing.T) {
		mRpcClient := &mocks.RpcClient{}

		client.c = mRpcClient

		mRpcClient.On("CallContext", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("*string"), "eth_blockNumber").
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(*string)
				*arg = "0x12345"
			})

		b, err := client.BlockNumber(ctx)
		assert.NoError(t, err)

		expectedReturn, _ := new(big.Int).SetString("0x12345", 0)
		assert.Equal(t, expectedReturn, b)
	})

	t.Run("It fails when empty result is coming from client", func(t *testing.T) {
		mRpcClient := &mocks.RpcClient{}
		client.c = mRpcClient

		mRpcClient.On("CallContext", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("*string"), "eth_blockNumber").
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(*string)
				*arg = ""
			})

		_, err := client.BlockNumber(ctx)
		assert.Equal(t, err, kowala.NotFound)
	})

	t.Run("It fails when invalid number is coming from client", func(t *testing.T) {
		mRpcClient := &mocks.RpcClient{}
		client.c = mRpcClient

		mRpcClient.On("CallContext", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("*string"), "eth_blockNumber").
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(*string)
				*arg = "*^^^^^"
			})

		_, err := client.BlockNumber(ctx)
		assert.Equal(t, err, ErrInvalidBlockNumber)
	})
}

func TestClient(t *testing.T) {
	client, err := Dial("http://rpcnode.testnet.kowala.io:30503")
	if err != nil {
		t.Fatalf("Error %s", err)
	}

	block, err := client.BlockByNumber(context.Background(), big.NewInt(1995415))
	if err != nil {
		t.Fatalf("Error %s", err)
	}

	fmt.Printf("%v", block.Transactions()[0].Hash().String())
}
