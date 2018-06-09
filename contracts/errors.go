package contracts

import "errors"

var (
	ErrNoAddress = errors.New("there isn't an address for the provided chain ID")
)
