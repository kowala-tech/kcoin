package types

import (
	"crypto/ecdsa"

	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/crypto"
)

type Validator struct {
	address   common.Address
	publicKey *ecdsa.PublicKey
}

func NewValidator(pubKey *ecdsa.PublicKey) *Validator {
	return &Validator{
		publicKey: pubKey,
		address:   crypto.PubkeyToAddress(pubKey),
	}
}

type Validators []*Validator
