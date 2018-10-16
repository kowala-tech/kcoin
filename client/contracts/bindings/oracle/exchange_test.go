package oracle_test

import (
	"testing"

	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind/backends"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/oracle/testfiles"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/utils"
	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/crypto"
	"github.com/stretchr/testify/suite"
)

var (
	exchange = "coinbase.com" // exchange name
)

type ExchangeMgrSuite struct {
	utils.ContractTestSuite
	exchangeMgr *testfiles.ExchangeMgr
}

func TestExchamgeMgrSuite(t *testing.T) {
	suite.Run(t, new(ExchangeMgrSuite))
}

func (suite *ExchangeMgrSuite) BeforeTest(suiteName, testName string) {
	req := suite.Require()

	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		crypto.PubkeyToAddress(owner.PublicKey): core.GenesisAccount{
			Balance: initialBalance,
		},
		crypto.PubkeyToAddress(user.PublicKey): core.GenesisAccount{
			Balance: initialBalance,
		},
	})
	req.NotNil(backend)
	suite.Backend = backend

	// deploy exchange mgr contract
	transactOpts := bind.NewKeyedTransactor(owner)
	_, _, exchangeMgrContract, err := testfiles.DeployExchangeMgr(transactOpts, suite.Backend)
	req.NoError(err)
	req.NotNil(exchangeMgrContract)
	suite.exchangeMgr = exchangeMgrContract

	suite.Backend.Commit()
}

func (suite *ExchangeMgrSuite) TestAddExchange_Paused() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// pause service
	_, err := suite.exchangeMgr.Pause(transactOpts)
	req.NoError(err)
	suite.Backend.Commit()

	// add exchange
	_, err = suite.exchangeMgr.AddExchange(transactOpts, exchange)
	req.Error(err, "service is paused")
}

func (suite *ExchangeMgrSuite) TestAddExchange_NotPaused_NotOwner() {
	req := suite.Require()

	// add exchange
	additionOpts := bind.NewKeyedTransactor(user)
	_, err := suite.exchangeMgr.AddExchange(additionOpts, exchange)
	req.Error(err, "transactor is not the owner")
}

func (suite *ExchangeMgrSuite) TestAddExchange_NotPaused_Owner_NewExchange() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// add exchange
	_, err := suite.exchangeMgr.AddExchange(transactOpts, exchange)
	req.NoError(err)
	suite.Backend.Commit()

	// added and whitelisted
	isWhitelistedExchange, err := suite.exchangeMgr.IsWhitelistedExchange(new(bind.CallOpts), exchange)
	req.NoError(err)
	req.True(isWhitelistedExchange)
}

func (suite *ExchangeMgrSuite) TestAddExchange_NotPaused_Owner_NotNewExchange() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// add exchange
	_, err := suite.exchangeMgr.AddExchange(transactOpts, exchange)
	req.NoError(err)
	suite.Backend.Commit()

	// add the same exchange again
	_, err = suite.exchangeMgr.AddExchange(transactOpts, exchange)
	req.Error(err, "exchange already exists")
}

func (suite *ExchangeMgrSuite) TestRemoveExchange_Paused() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// add exchange
	_, err := suite.exchangeMgr.AddExchange(transactOpts, exchange)
	req.NoError(err)
	suite.Backend.Commit()

	// pause service
	_, err = suite.exchangeMgr.Pause(transactOpts)
	req.NoError(err)
	suite.Backend.Commit()

	// remove exchange
	_, err = suite.exchangeMgr.RemoveExchange(transactOpts, exchange)
	req.Error(err, "service is paused")
}

func (suite *ExchangeMgrSuite) TestRemoveExchange_NotPaused_NotOwner() {
	req := suite.Require()

	// add exchange
	additionOpts := bind.NewKeyedTransactor(owner)
	_, err := suite.exchangeMgr.AddExchange(additionOpts, exchange)
	req.NoError(err)
	suite.Backend.Commit()

	// remove exchange
	removalOpts := bind.NewKeyedTransactor(user)
	_, err = suite.exchangeMgr.RemoveExchange(removalOpts, exchange)
	req.Error(err, "transactor is not the owner")
}

func (suite *ExchangeMgrSuite) TestRemoveExchange_NotPaused_Owner_NotExchange() {
	req := suite.Require()

	// remove exchange
	removalOpts := bind.NewKeyedTransactor(owner)
	_, err := suite.exchangeMgr.RemoveExchange(removalOpts, exchange)
	req.Error(err, "exchange does not exist")
}

func (suite *ExchangeMgrSuite) TestRemoveExchange_NotPaused_Owner_Exchange() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// add exchange
	_, err := suite.exchangeMgr.AddExchange(transactOpts, exchange)
	req.NoError(err)
	suite.Backend.Commit()

	// remove exchange
	_, err = suite.exchangeMgr.RemoveExchange(transactOpts, exchange)
	req.NoError(err)
	suite.Backend.Commit()

	// must not be an exchange
	isExchange, err := suite.exchangeMgr.IsExchange(nil, exchange)
	req.NoError(err)
	req.False(isExchange)
}

func (suite *ExchangeMgrSuite) TestWhitelistExchange_Paused() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// add exchange
	_, err := suite.exchangeMgr.AddExchange(transactOpts, exchange)
	req.NoError(err)
	suite.Backend.Commit()

	// blacklist exchange
	_, err = suite.exchangeMgr.BlacklistExchange(transactOpts, exchange)
	req.NoError(err)
	suite.Backend.Commit()

	// pause service
	_, err = suite.exchangeMgr.Pause(transactOpts)
	req.NoError(err)
	suite.Backend.Commit()

	// whitelist exchange
	_, err = suite.exchangeMgr.WhitelistExchange(transactOpts, exchange)
	req.Error(err, "service is paused")
}

func (suite *ExchangeMgrSuite) TestWhitelistExchange_NotPaused_NotOwner() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// add exchange
	_, err := suite.exchangeMgr.AddExchange(transactOpts, exchange)
	req.NoError(err)
	suite.Backend.Commit()

	// blacklist exchange
	_, err = suite.exchangeMgr.BlacklistExchange(transactOpts, exchange)
	req.NoError(err)
	suite.Backend.Commit()

	// whitelist exchange
	blacklistOpts := bind.NewKeyedTransactor(user)
	_, err = suite.exchangeMgr.WhitelistExchange(blacklistOpts, exchange)
	req.Error(err, "transactor is not the owner")
}

func (suite *ExchangeMgrSuite) TestWhitelistExchange_NotPaused_Owner_NotBlacklistedExchange() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// whitelist exchange
	_, err := suite.exchangeMgr.WhitelistExchange(transactOpts, "exchange")
	req.Error(err, "exchange does not exist")
}

func (suite *ExchangeMgrSuite) TestWhitelistExchange_NotPaused_Owner_BlacklistedExchange() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// add exchange
	_, err := suite.exchangeMgr.AddExchange(transactOpts, exchange)
	req.NoError(err)
	suite.Backend.Commit()

	// blacklist exchange
	_, err = suite.exchangeMgr.BlacklistExchange(transactOpts, exchange)
	req.NoError(err)
	suite.Backend.Commit()

	// whitelist exchange
	_, err = suite.exchangeMgr.WhitelistExchange(transactOpts, exchange)
	req.NoError(err)
	suite.Backend.Commit()

	// whitelisted
	isWhitelistedExchange, err := suite.exchangeMgr.IsWhitelistedExchange(nil, exchange)
	req.NoError(err)
	req.True(isWhitelistedExchange)
}

func (suite *ExchangeMgrSuite) TestBlacklistExchange_Paused() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// add exchange
	_, err := suite.exchangeMgr.AddExchange(transactOpts, exchange)
	req.NoError(err)
	suite.Backend.Commit()

	// pause service
	_, err = suite.exchangeMgr.Pause(transactOpts)
	req.NoError(err)
	suite.Backend.Commit()

	// blacklist exchange
	_, err = suite.exchangeMgr.BlacklistExchange(transactOpts, exchange)
	req.Error(err, "service is paused")
}

func (suite *ExchangeMgrSuite) TestBlacklistExchange_NotPaused_NotOwner() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// add exchange
	_, err := suite.exchangeMgr.AddExchange(transactOpts, exchange)
	req.NoError(err)
	suite.Backend.Commit()

	// blacklist exchange
	blacklistOpts := bind.NewKeyedTransactor(user)
	_, err = suite.exchangeMgr.BlacklistExchange(blacklistOpts, exchange)
	req.Error(err, "transactor is not the owner")
}

func (suite *ExchangeMgrSuite) TestBlacklistExchange_NotPaused_Owner_NotWhitelistedExchange() {
	req := suite.Require()

	// blacklist exchange
	transactOpts := bind.NewKeyedTransactor(owner)
	_, err := suite.exchangeMgr.BlacklistExchange(transactOpts, exchange)
	req.Error(err, "exchange does not exist")
}

func (suite *ExchangeMgrSuite) TestBlacklistExchange_NotPaused_Owner_WhitelistedExchange() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// add exchange
	_, err := suite.exchangeMgr.AddExchange(transactOpts, exchange)
	req.NoError(err)
	suite.Backend.Commit()

	// blacklist exchange
	_, err = suite.exchangeMgr.BlacklistExchange(transactOpts, exchange)
	req.NoError(err)
	suite.Backend.Commit()

	// most not be an whitelisted exchange
	isBlacklistedExchange, err := suite.exchangeMgr.IsBlacklistedExchange(nil, exchange)
	req.NoError(err)
	req.True(isBlacklistedExchange)
}
