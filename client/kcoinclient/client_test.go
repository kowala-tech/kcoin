package kcoinclient

import (
	"context"
	"testing"

	"math/big"

	"github.com/kowala-tech/kcoin/client"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/kcoinclient/mocks"
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

func TestClient_Coinbase(t *testing.T) {
	client := Client{}
	ctx := context.Background()

	t.Run("It returns an account", func(t *testing.T) {
		mRpcClient := &mocks.RpcClient{}

		client.c = mRpcClient

		mRpcClient.On("CallContext", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("*string"), "eth_coinbase").
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(*string)
				*arg = "0x1234567890123456789012345678901234567890"
			})

		addr, err := client.Coinbase(ctx)
		assert.NoError(t, err)

		assert.Equal(t, common.HexToAddress("0x1234567890123456789012345678901234567890"), *addr)
	})

	t.Run("It returns nil if there's no coinbase", func(t *testing.T) {
		mRpcClient := &mocks.RpcClient{}

		client.c = mRpcClient

		mRpcClient.On("CallContext", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("*string"), "eth_coinbase").
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(*string)
				*arg = ""
			})

		addr, err := client.Coinbase(ctx)
		assert.NoError(t, err)
		assert.Nil(t, addr)
	})

}
