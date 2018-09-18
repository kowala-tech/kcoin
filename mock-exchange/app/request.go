package app

type Request struct {
	Sell []RateValue `json:"sell"`
	Buy  []RateValue `json:"buy"`
}

type RateValue struct {
	Amount float64
	Rate   float64
}
