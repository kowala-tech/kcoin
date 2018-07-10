package command

import (
	"context"
	"math/big"
	"testing"

	"github.com/kowala-tech/kcoin/wallet-backend/blockchain/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetBlockHeight(t *testing.T) {
	mockedClient := &mocks.Client{}
	ctx := context.Background()

	handler := GetBlockHeightHandler{
		Client: mockedClient,
	}

	expectedReturn := &BlockHeightResponse{
		big.NewInt(1234),
	}

	mockedClient.On("BlockNumber", ctx).
		Return(big.NewInt(1234), nil)

	blockHeight, err := handler.Handle(ctx)
	assert.NoError(t, err)

	assert.Equal(t, expectedReturn, blockHeight)
}
