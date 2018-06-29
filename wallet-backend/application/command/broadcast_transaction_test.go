package command

import (
	"testing"

	"context"

	"github.com/kowala-tech/kcoin/wallet-backend/blockchain/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBroadcastTransaction(t *testing.T) {
	mockedClient := &mocks.Client{}

	hdl := BroadcastTransactionHandler{
		Client: mockedClient,
	}

	mockedTrans := []byte("mockedTransBytes")
	ctx := context.Background()

	mockedClient.On("SendRawTransaction", ctx, mockedTrans).
		Return(nil)

	expectedResponse := &BroadcastTransactionResponse{
		Status: "ok",
	}

	resp, err := hdl.Handle(ctx, mockedTrans)
	if err != nil {
		t.Fatalf("Error testing broadcast transaction.")
	}

	assert.Equal(t, expectedResponse, resp)
}
