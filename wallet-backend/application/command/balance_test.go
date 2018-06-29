package command

import (
	"context"
	"math/big"
	"testing"

	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/wallet-backend/blockchain/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetBalance(t *testing.T) {
	mockedClient := &mocks.Client{}

	handler := GetBalanceHandler{
		Client: mockedClient,
	}

	ctx := context.Background()
	address := common.HexToAddress("0xd6e579085c82329c89fca7a9f012be59028ed53f")

	expectedBalance := big.NewInt(1234)

	var blockNum *big.Int
	mockedClient.On("BalanceAt", ctx, address, blockNum).
		Return(expectedBalance, nil)

	balance, err := handler.Handle(ctx, address)
	if err != nil {
		t.Fatalf("Error %v", err)
	}

	expectedResponse := &BalanceResponse{
		Balance: expectedBalance,
	}
	assert.Equal(t, expectedResponse, balance)
}
