package exrates

import (
	"encoding/json"

	"github.com/kowala-tech/kcoin/mock-exchange/app"
)

type Transformer struct {
}

func (*Transformer) Transform(request app.Request) (string, error) {
	response := Response{
		Sell: make([]Value, 0),
		Buy:  make([]Value, 0),
	}

	for _, s := range request.Sell {
		response.Sell = append(response.Sell, Value{Amount: s.Amount, Rate: s.Rate})
	}

	for _, s := range request.Buy {
		response.Buy = append(response.Buy, Value{Amount: s.Amount, Rate: s.Rate})
	}

	encodedResp, err := json.Marshal(response)
	if err != nil {
		return "", err
	}

	return string(encodedResp), nil
}

type Response struct {
	Sell []Value `json:"SELL"`
	Buy  []Value `json:"BUY"`
}

type Value struct {
	Amount float64 `json:"amount"`
	Rate   float64 `json:"rate"`
}
