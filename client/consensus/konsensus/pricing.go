package Konsensus

import "math/big"

type Pricing interface {
	SetPrince(price *big.Int)
}

type PricingFunc func(*big.Int)

type PricingMiddleware func(Pricing) Pricing
