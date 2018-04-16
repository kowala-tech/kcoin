package kcoinclient

import (
	"context"
	"testing"

	"math/big"

	"github.com/kowala-tech/kcoin/kcoinclient/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestClient_BlockNumber(t *testing.T) {
	mRpcClient := &mocks.RpcClient{}
	client := Client{c: mRpcClient}
	ctx := context.Background()

	mRpcClient.On("CallContext", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("*string"), "eth_blockNumber").
		Return(nil).
		Run(func(args mock.Arguments) {
			arg := args.Get(1).(*string)
			*arg = "0x12345"
		})

	b, err := client.BlockNumber(ctx)
	if err != nil {
		t.Fatalf("Error %s", err)
	}

	expectedReturn, _ := new(big.Int).SetString("0x12345", 0)
	assert.Equal(t, expectedReturn, b)
}
