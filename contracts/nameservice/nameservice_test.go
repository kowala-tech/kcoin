package nameservice

import (
	"crypto/ecdsa"

	"github.com/kowala-tech/kcoin/accounts/abi/bind/backends"
	"github.com/stretchr/testify/suite"
)

type NameServiceSuite struct {
	suite.Suite
	backend           *backends.SimulatedBackend
	owner, randomUser *ecdsa.PrivateKey
}