package extares

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kowala-tech/kcoin/mock-exchange/server"
)

func TestWeCanTransformADataRequestToExrateResponse(t *testing.T) {
	requestToTransform := server.FetchDataRequest{
		Sell: []server.Value{
			{
				Amount: 0.358,
				Rate:   6326.83689418,
			},
			{
				Amount: 0.1427,
				Rate:   6326.83689421,
			},
		},
		Buy: []server.Value{
			{
				Amount: 0.0021,
				Rate:   6214.3034165,
			},
			{
				Amount: 0.0029,
				Rate:   6203.01833171,
			},
		},
	}

	trans := Transformer{}
	reqTransformed, err := trans.Transform(requestToTransform)
	if err != nil {
		t.Fatalf("Failed to encode request: %s", err)
	}

	expectedTransformedResponse :=
		`{"SELL":[{"amount":0.358,"rate":6326.83689418},{"amount":0.1427,"rate":6326.83689421}],"BUY":[{"amount":0.0021,"rate":6214.3034165},{"amount":0.0029,"rate":6203.01833171}]}`

	assert.Equal(t, expectedTransformedResponse, reqTransformed)
}
