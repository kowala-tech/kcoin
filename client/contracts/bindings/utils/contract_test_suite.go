package utils

import (
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind/backends"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/oracle/testfiles"
	"github.com/stretchr/testify/suite"
)

type ContractTestSuite struct {
	suite.Suite
	Backend *backends.SimulatedBackend
}

//DeployConsensusMock deploys a mock of the consensus contract.
func (suite *ContractTestSuite) DeployConsensusMock(opts *bind.TransactOpts, isSuperNode bool) common.Address {
	req := suite.Require()

	mockAddr, _, _, err := testfiles.DeployConsensusMock(opts, suite.Backend, isSuperNode)
	req.NoError(err)
	req.NotZero(mockAddr)
	suite.Backend.Commit()

	return mockAddr
}

//DeployStringsLibrary deploys the Strings.sol library.
func (suite *ContractTestSuite) DeployStringsLibrary(transactOpts *bind.TransactOpts) common.Address {
	req := suite.Require()

	stringsLibAddr, _, _, err := DeployStrings(transactOpts, suite.Backend)
	req.NoError(err)
	req.NotZero(stringsLibAddr)
	suite.Backend.Commit()

	return stringsLibAddr
}

//DeployNameHashLibrary deploys NameHash.sol library.
func (suite *ContractTestSuite) DeployNameHashLibrary(transactOpts *bind.TransactOpts) common.Address {
	req := suite.Require()

	nameHashLib, _, _, err := DeployNameHash(transactOpts, suite.Backend)
	req.NoError(err)
	req.NotZero(nameHashLib)
	suite.Backend.Commit()

	return nameHashLib
}
