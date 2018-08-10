package consensus_test

import (
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/kowala-tech/kcoin/client/accounts/abi"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind/backends"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/consensus"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/consensus/testfiles"
	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/core/state"
	"github.com/kowala-tech/kcoin/client/core/vm"
	"github.com/kowala-tech/kcoin/client/core/vm/runtime"
	"github.com/kowala-tech/kcoin/client/crypto"
	"github.com/kowala-tech/kcoin/client/kcoindb"
	"github.com/kowala-tech/kcoin/client/params"
	"github.com/stretchr/testify/suite"
)

var (
	owner, _       = crypto.GenerateKey()
	user, _        = crypto.GenerateKey()
	secondsPerDay  = new(big.Int).SetUint64(86400)
	initialBalance = new(big.Int).Mul(new(big.Int).SetUint64(1000), new(big.Int).SetUint64(params.Kcoin))
)

func getDefaultOpts() genesis.Options {
	baseDeposit := uint64(20)
	superNodeAmount := uint64(6000000)
	tokenHolder := genesis.TokenHolder{
		Address:   getAddress(validator).Hex(),
		NumTokens: superNodeAmount,
	}

	opts := genesis.Options{
		Network: "test",
		Consensus: &genesis.ConsensusOpts{
			Engine:           "konsensus",
			MaxNumValidators: 10,
			FreezePeriod:     30,
			BaseDeposit:      baseDeposit,
			SuperNodeAmount:  superNodeAmount,
			Validators: []genesis.Validator{{
				Address: tokenHolder.Address,
				Deposit: tokenHolder.NumTokens,
			}},
			MiningToken: &genesis.MiningTokenOpts{
				Name:     "mUSD",
				Symbol:   "mUSD",
				Cap:      20000000,
				Decimals: 18,
				Holders:  []genesis.TokenHolder{tokenHolder, {Address: getAddress(user).Hex(), NumTokens: 10000000}},
			},
		},
		Governance: &genesis.GovernanceOpts{
			Origin:           getAddress(author).Hex(),
			Governors:        []string{getAddress(governor).Hex()},
			NumConfirmations: 1,
		},
		DataFeedSystem: &genesis.DataFeedSystemOpts{
			MaxNumOracles: 10,
			FreezePeriod:  0,
			BaseDeposit:   0,
			Price: genesis.PriceOpts{
				InitialPrice:  1,
				SyncFrequency: 600,
				UpdatePeriod:  30,
			},
		},
		PrefundedAccounts: []genesis.PrefundedAccount{
			{
				Address: tokenHolder.Address,
				Balance: 10,
			},
			{
				Address: getAddress(governor).Hex(),
				Balance: 10,
			},
			{
				Address: getAddress(user).Hex(),
				Balance: 10,
			},
			{
				Address: getAddress(deregistered).Hex(),
				Balance: 10,
			},
		},
	}

	return opts
}

type ValidatorMgrSuite struct {
	suite.Suite
	backend         *backends.SimulatedBackend
	validatorMgr    *consensus.ValidatorMgr
	tokenMock       *testfiles.TokenMock
	superNodeAmount *big.Int
	baseDeposit     *big.Int
}

func TestValidatorMgrSuite(t *testing.T) {
	suite.Run(t, new(ValidatorMgrSuite))
}

func (suite *ValidatorMgrSuite) BeforeTest(suiteName, testName string) {
	req := suite.Require()

	// we skip the following code to include to manipulate the genesis state
	// during the test
	if strings.Contains(testName, "TestIsGenesisValidator") {
		return
	}

	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		crypto.PubkeyToAddress(owner.PublicKey): core.GenesisAccount{
			Balance: initialBalance,
		},
		crypto.PubkeyToAddress(user.PublicKey): core.GenesisAccount{
			Balance: initialBalance,
		},
	})
	req.NotNil(backend)
	suite.backend = backend

	// we deploy the contracts during the test
	if strings.Contains(testName, "TestDeploy") {
		return
	}

	maxNumValidators := 100
	freezePeriod := 1000
	switch {
	case strings.Contains(testName, "_Full"):
		maxNumValidators = 1
		fallthrough
	case strings.Contains(testName, "_UnlockedDeposit"):
		freezePeriod = 0
	}

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy token mock
	mockAddr, _, tokenMock, err := testfiles.DeployTokenMock(transactOpts, suite.backend)
	req.NoError(err)
	req.NotNil(tokenMock)
	req.NotZero(mockAddr)
	suite.tokenMock = tokenMock

	suite.backend.Commit()

	// deploy validator mgr
	baseDeposit := new(big.Int).SetUint64(100)
	superNodeAmount := new(big.Int).SetUint64(200)
	_, _, validatorMgr, err := consensus.DeployValidatorMgr(transactOpts, suite.backend, baseDeposit, big.NewInt(int64(maxNumValidators)), big.NewInt(int64(freezePeriod)), mockAddr, superNodeAmount)
	req.NoError(err)
	req.NotNil(validatorMgr)
	suite.validatorMgr = validatorMgr
	suite.superNodeAmount = superNodeAmount
	suite.baseDeposit = baseDeposit

	suite.backend.Commit()

}

func (suite *ValidatorMgrSuite) TestDeploy() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy token mock
	mockAddr, _, tokenMock, err := testfiles.DeployTokenMock(transactOpts, suite.backend)
	req.NoError(err)
	req.NotNil(tokenMock)
	req.NotZero(mockAddr)

	suite.backend.Commit()

	baseDeposit := new(big.Int).SetUint64(100)
	maxNumValidators := new(big.Int).SetUint64(100)
	freezePeriod := new(big.Int).SetUint64(1000)
	superNodeAmount := new(big.Int).SetUint64(500000)
	_, _, validatorMgr, err := consensus.DeployValidatorMgr(transactOpts, suite.backend, baseDeposit, maxNumValidators, freezePeriod, mockAddr, superNodeAmount)
	req.NoError(err)
	req.NotNil(validatorMgr)

	suite.backend.Commit()

	storedBaseDeposit, err := validatorMgr.BaseDeposit(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedBaseDeposit)
	req.Equal(baseDeposit, storedBaseDeposit)

	storedFreezePeriod, err := validatorMgr.FreezePeriod(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedFreezePeriod)
	req.Equal(dtos(freezePeriod), storedFreezePeriod)

	storedMaxNumValidators, err := validatorMgr.MaxNumValidators(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedMaxNumValidators)
	req.Equal(maxNumValidators, storedMaxNumValidators)

	storedMiningTokenAddr, err := validatorMgr.MiningTokenAddr(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedMiningTokenAddr)
	req.Equal(mockAddr, storedMiningTokenAddr)

	storedSuperNodeAmount, err := validatorMgr.SuperNodeAmount(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedSuperNodeAmount)
	req.Equal(superNodeAmount, storedSuperNodeAmount)
}

func (suite *ValidatorMgrSuite) TestDeploy_MaxNumValidatorsZero() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy token mock
	mockAddr, _, tokenMock, err := testfiles.DeployTokenMock(transactOpts, suite.backend)
	req.NoError(err)
	req.NotNil(tokenMock)
	req.NotZero(mockAddr)

	suite.backend.Commit()

	baseDeposit := new(big.Int).SetUint64(100)
	maxNumValidators := common.Big0 // set max num validators to zero
	freezePeriod := new(big.Int).SetUint64(1000)
	superNodeAmount := new(big.Int).SetUint64(500000)
	_, _, _, err = consensus.DeployValidatorMgr(transactOpts, suite.backend, baseDeposit, maxNumValidators, freezePeriod, mockAddr, superNodeAmount)
	req.Error(err, "max number of validators cannot be zero")
}

func (suite *ValidatorMgrSuite) TestIsValidator() {
	req := suite.Require()

	// register validator 1
	registerOpts1 := bind.NewKeyedTransactor(owner)
	from1 := crypto.PubkeyToAddress(owner.PublicKey)
	value1 := new(big.Int).SetUint64(200)
	_, err := suite.validatorMgr.RegisterValidator(registerOpts1, from1, value1, []byte{})
	req.NoError(err)

	// register validator 2
	registerOpts2 := bind.NewKeyedTransactor(user)
	from2 := crypto.PubkeyToAddress(user.PublicKey)
	value2 := new(big.Int).SetUint64(200)
	_, err = suite.validatorMgr.RegisterValidator(registerOpts2, from2, value2, []byte{})
	req.NoError(err)

	// deregister validator 2
	_, err = suite.validatorMgr.DeregisterValidator(registerOpts2)
	req.NoError(err)

	// create a random user
	randomUser, err := crypto.GenerateKey()
	req.NoError(err)
	req.NotNil(randomUser)

	suite.backend.Commit()

	testCases := []struct {
		name   string
		input  common.Address
		output bool
	}{
		{
			name:   "validator",
			input:  from1,
			output: true,
		},
		{
			name:   "deregistered validator",
			input:  from2,
			output: false,
		},
		{
			name:   "random user",
			input:  crypto.PubkeyToAddress(randomUser.PublicKey),
			output: false,
		},
	}
	for _, tc := range testCases {
		suite.T().Run(fmt.Sprintf("role: %s, address: %s", tc.name, tc.input.Hex()), func(t *testing.T) {
			isValidator, err := suite.validatorMgr.IsValidator(&bind.CallOpts{}, tc.input)
			req.NoError(err)
			req.Equal(tc.output, isValidator)
		})
	}
}

type vmTracer struct {
	data map[common.Address]map[common.Hash]common.Hash
}

func newVMTracer() *vmTracer {
	return &vmTracer{
		data: make(map[common.Address]map[common.Hash]common.Hash, 1024),
	}
}

func (vmt *vmTracer) CaptureState(env *vm.EVM, pc uint64, op vm.OpCode, gas, cost uint64, memory *vm.Memory, stack *vm.Stack, contract *vm.Contract, depth int, err error) error {
	if err != nil {
		return err
	}
	if op == vm.SSTORE {
		s := stack.Data()

		addrStorage, ok := vmt.getAddrStorage(contract.Address())
		if !ok {
			addrStorage = make(map[common.Hash]common.Hash, 1024)
			vmt.setAddrStorage(contract.Address(), addrStorage)
		}
		addrStorage[common.BigToHash(s[len(s)-1])] = common.BigToHash(s[len(s)-2])
	}
	return nil
}

func (vmt *vmTracer) getAddrStorage(contractAddress common.Address) (addrStorage map[common.Hash]common.Hash, ok bool) {
	addrStorage, ok = vmt.data[contractAddress]
	return
}

func (vmt *vmTracer) setAddrStorage(contractAddress common.Address, addrStorage map[common.Hash]common.Hash) {
	vmt.data[contractAddress] = addrStorage
}

func (vmt *vmTracer) CaptureStart(from common.Address, to common.Address, call bool, input []byte, gas uint64, value *big.Int) error {
	return nil
}

func (vmt *vmTracer) CaptureEnd(output []byte, gasUsed uint64, t time.Duration, err error) error {
	return nil
}

func (vmt *vmTracer) CaptureFault(env *vm.EVM, pc uint64, op vm.OpCode, gas, cost uint64, memory *vm.Memory, stack *vm.Stack, contract *vm.Contract, depth int, err error) error {
	return nil
}

func (suite *ValidatorMgrSuite) TestIsGenesisValidator() {
	req := suite.Require()

	sharedState, err := state.New(common.Hash{}, state.NewDatabase(kcoindb.NewMemDatabase()))
	req.NoError(err)
	req.NotNil(sharedState)

	runtimeCfg := &runtime.Config{
		State:       sharedState,
		BlockNumber: common.Big0,
		EVMConfig: vm.Config{
			Debug:  true,
			Tracer: newVMTracer(),
		},
		Origin: crypto.PubkeyToAddress(owner.PublicKey),
	}

	// deploy token mock
	tokenMockABI, err := abi.JSON(strings.NewReader(testfiles.TokenMockABI))
	req.NoError(err)
	req.NotNil(tokenMockABI)

	tokenCode, tokenAddr, _, err := runtime.Create(common.FromHex(testfiles.TokenMockBin), runtimeCfg)
	req.NoError(err)
	req.NotZero(tokenCode)
	req.NotZero(tokenAddr)

	// deploy validator mgr
	validatorMgrABI, err := abi.JSON(strings.NewReader(consensus.ValidatorMgrABI))
	req.NoError(err)
	req.NotZero(validatorMgrABI)

	baseDeposit := new(big.Int).SetUint64(100)
	maxNumValidators := new(big.Int).SetUint64(100)
	freezePeriod := new(big.Int).SetUint64(10)
	superNodeAmount := new(big.Int).SetUint64(200)
	managerParams, err := validatorMgrABI.Pack(
		"",
		baseDeposit,
		maxNumValidators,
		freezePeriod,
		tokenAddr,
		superNodeAmount,
	)
	req.NoError(err)
	req.NotNil(managerParams)

	validatorMgrCode, validatorMgrAddr, _, err := runtime.Create(append(common.FromHex(consensus.ValidatorMgrBin), managerParams...), runtimeCfg)
	req.NoError(err)
	req.NotZero(validatorMgrAddr)
	req.NotZero(validatorMgrCode)

	// register genesis validator
	genesisValidator, err := crypto.GenerateKey()
	req.NoError(err)
	req.NotNil(genesisValidator)

	genesisValidatorAddr := crypto.PubkeyToAddress(genesisValidator.PublicKey)
	req.NotZero(genesisValidatorAddr)

	registrationParams, err := validatorMgrABI.Pack(
		"registerValidator",
		genesisValidatorAddr,
		baseDeposit,
		[]byte("not_zero"), // @NOTE (rgeraldes) - https://github.com/kowala-tech/kcoin/client/issues/285

	)
	req.NoError(err)
	req.NotZero(registrationParams)

	runtimeCfg.Origin = genesisValidatorAddr
	_, _, err = runtime.Call(validatorMgrAddr, registrationParams, runtimeCfg)
	req.NoError(err)

	// create backend
	userAddr := crypto.PubkeyToAddress(user.PublicKey)
	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		tokenAddr: core.GenesisAccount{
			Code:    tokenCode,
			Storage: runtimeCfg.EVMConfig.Tracer.(*vmTracer).data[tokenAddr],
			Balance: common.Big0,
		},
		validatorMgrAddr: core.GenesisAccount{
			Code:    validatorMgrCode,
			Storage: runtimeCfg.EVMConfig.Tracer.(*vmTracer).data[validatorMgrAddr],
			Balance: common.Big0,
		},
		userAddr: core.GenesisAccount{
			Balance: initialBalance,
		},
	})
	req.NotNil(backend)
	suite.backend = backend

	validatorMgr, err := consensus.NewValidatorMgr(validatorMgrAddr, suite.backend)
	req.NoError(err)
	req.NotNil(validatorMgr)
	suite.validatorMgr = validatorMgr

	// register a new validator (not genesis)
	registerOpts := bind.NewKeyedTransactor(user)
	from := crypto.PubkeyToAddress(user.PublicKey)
	value := baseDeposit
	_, err = suite.validatorMgr.RegisterValidator(registerOpts, from, value, []byte{})
	req.NoError(err)

	// create a random user
	randomUser, err := crypto.GenerateKey()
	req.NoError(err)
	req.NotNil(randomUser)

	suite.backend.Commit()

	testCases := []struct {
		name   string
		input  common.Address
		output bool
	}{
		{
			name:   "genesis validator",
			input:  genesisValidatorAddr,
			output: true,
		},
		{
			name:   "non-genesis validator",
			input:  userAddr,
			output: false,
		},
		{
			name:   "random user",
			input:  crypto.PubkeyToAddress(randomUser.PublicKey),
			output: false,
		},
	}
	for _, tc := range testCases {
		suite.T().Run(fmt.Sprintf("role: %s, address %s", tc.name, tc.input.Hex()), func(t *testing.T) {
			isValidator, err := suite.validatorMgr.IsGenesisValidator(&bind.CallOpts{}, tc.input)
			req.NoError(err)
			req.Equal(tc.output, isValidator)
		})
	}
}

func (suite *ValidatorMgrSuite) TestIsSuperNode() {
	req := suite.Require()

	// register validator as super node
	registerOpts1 := bind.NewKeyedTransactor(owner)
	from1 := crypto.PubkeyToAddress(owner.PublicKey)
	value1 := suite.superNodeAmount
	_, err := suite.validatorMgr.RegisterValidator(registerOpts1, from1, value1, []byte{})
	req.NoError(err)

	// register another validator (not super node)
	registerOpts2 := bind.NewKeyedTransactor(user)
	from2 := crypto.PubkeyToAddress(user.PublicKey)
	value2 := new(big.Int).Sub(suite.superNodeAmount, common.Big1)
	_, err = suite.validatorMgr.RegisterValidator(registerOpts2, from2, value2, []byte{})
	req.NoError(err)

	// create random user
	randomUser, err := crypto.GenerateKey()
	req.NoError(err)
	req.NotNil(randomUser)

	suite.backend.Commit()

	testCases := []struct {
		name   string
		input  common.Address
		output bool
	}{
		{
			name:   "super node",
			input:  from1,
			output: true,
		},
		{
			name:   "not super node",
			input:  from2,
			output: false,
		},
		{
			name:   "not validator - random user",
			input:  crypto.PubkeyToAddress(randomUser.PublicKey),
			output: false,
		},
	}
	for _, tc := range testCases {
		suite.T().Run(fmt.Sprintf("name: %s, address %s", tc.name, tc.input.Hex()), func(t *testing.T) {
			isSuperNode, err := suite.validatorMgr.IsSuperNode(&bind.CallOpts{}, tc.input)
			req.NoError(err)
			req.Equal(tc.output, isSuperNode)
		})
	}
}

func (suite *ValidatorMgrSuite) TestGetMinimumDeposit_NotFull() {
	req := suite.Require()

	storedMinDeposit, err := suite.validatorMgr.GetMinimumDeposit(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedMinDeposit)
	req.Equal(suite.baseDeposit, storedMinDeposit)
}

func (suite *ValidatorMgrSuite) TestGetMinimumDeposit_Full() {
	req := suite.Require()

	// register validator to match the maximum number of validators
	registerOpts := bind.NewKeyedTransactor(owner)
	from := crypto.PubkeyToAddress(owner.PublicKey)
	value := suite.baseDeposit
	_, err := suite.validatorMgr.RegisterValidator(registerOpts, from, value, []byte{})
	req.NoError(err)

	suite.backend.Commit()

	storedMinDeposit, err := suite.validatorMgr.GetMinimumDeposit(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedMinDeposit)
	req.Equal(new(big.Int).Add(suite.baseDeposit, common.Big1), storedMinDeposit)
}

func (suite *ValidatorMgrSuite) TestRegisterValidator_WhenPaused() {
	req := suite.Require()

	pauseOpts := bind.NewKeyedTransactor(owner)
	_, err := suite.validatorMgr.Pause(pauseOpts)
	req.NoError(err)

	suite.backend.Commit()

	// register validator
	registerOpts := bind.NewKeyedTransactor(owner)
	from := crypto.PubkeyToAddress(owner.PublicKey)
	value := suite.baseDeposit
	_, err = suite.validatorMgr.RegisterValidator(registerOpts, from, value, []byte{})
	req.Error(err, "cannot register the validator if the service is paused")
}

func (suite *ValidatorMgrSuite) TestRegisterValidator_Duplicate() {
	req := suite.Require()

	// register validator
	registerOpts := bind.NewKeyedTransactor(owner)
	from := crypto.PubkeyToAddress(owner.PublicKey)
	value := suite.baseDeposit
	_, err := suite.validatorMgr.RegisterValidator(registerOpts, from, value, []byte{})
	req.NoError(err)

	// register validator again
	_, err = suite.validatorMgr.RegisterValidator(registerOpts, from, value, []byte{})
	req.Error(err, "cannot register the same validator twice")
}

func (suite *ValidatorMgrSuite) TestRegisterValidator_WithoutMinDeposit() {
	req := suite.Require()

	// register validator
	registerOpts := bind.NewKeyedTransactor(owner)
	from := crypto.PubkeyToAddress(owner.PublicKey)
	value := new(big.Int).Sub(suite.baseDeposit, common.Big1) // set value to less than base deposit
	_, err := suite.validatorMgr.RegisterValidator(registerOpts, from, value, []byte{})
	req.Error(err, "requires the minimum deposit")
}

func (suite *ValidatorMgrSuite) TestRegister_NotPaused_NewCandidate_WithMinDeposit_NotFull() {
	req := suite.Require()

	// register validator
	registerOpts := bind.NewKeyedTransactor(owner)
	from := crypto.PubkeyToAddress(owner.PublicKey)
	value := suite.baseDeposit
	_, err := suite.validatorMgr.RegisterValidator(registerOpts, from, value, []byte{})
	req.NoError(err)

	suite.backend.Commit()

	validatorCount, err := suite.validatorMgr.GetValidatorCount(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(validatorCount)
	req.Equal(common.Big1, validatorCount)

	storedValidator, err := suite.validatorMgr.GetValidatorAtIndex(&bind.CallOpts{}, common.Big0)
	req.NoError(err)
	req.NotZero(storedValidator)
	req.Equal(storedValidator.Code, from)
	req.Equal(storedValidator.Deposit, value)

	depositCount, err := suite.validatorMgr.GetDepositCount(&bind.CallOpts{From: from})
	req.NoError(err)
	req.NotNil(depositCount)
	req.Equal(depositCount, common.Big1)

	deposit, err := suite.validatorMgr.GetDepositAtIndex(&bind.CallOpts{From: from}, common.Big0)
	req.NoError(err)
	req.NotZero(deposit)
	req.Zero(deposit.AvailableAt.Uint64())
	req.Equal(value, deposit.Amount)
}

func (suite *ValidatorMgrSuite) TestRegister_NotPaused_NewCandidate_WithMinDeposit_NotFull_DepositGreaterThan() {
	req := suite.Require()

	// register validator 1
	registerOpts1 := bind.NewKeyedTransactor(owner)
	from1 := crypto.PubkeyToAddress(owner.PublicKey)
	value1 := suite.baseDeposit
	_, err := suite.validatorMgr.RegisterValidator(registerOpts1, from1, value1, []byte{})
	req.NoError(err)

	// register validator 2
	registerOpts2 := bind.NewKeyedTransactor(user)
	from2 := crypto.PubkeyToAddress(user.PublicKey)
	value2 := new(big.Int).Add(value1, common.Big1)
	_, err = suite.validatorMgr.RegisterValidator(registerOpts2, from2, value2, []byte{})
	req.NoError(err)

	suite.backend.Commit()

	validatorCount, err := suite.validatorMgr.GetValidatorCount(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(validatorCount)
	req.Equal(common.Big2, validatorCount)

	highestBidder, err := suite.validatorMgr.GetValidatorAtIndex(&bind.CallOpts{}, common.Big0)
	req.NoError(err)
	req.NotZero(highestBidder)
	req.Equal(highestBidder.Code, from2)
	req.Equal(highestBidder.Deposit, value2)

	lowestBidder, err := suite.validatorMgr.GetValidatorAtIndex(&bind.CallOpts{}, common.Big1)
	req.NoError(err)
	req.NotZero(highestBidder)
	req.Equal(lowestBidder.Code, from1)
	req.Equal(lowestBidder.Deposit, value1)
}

func (suite *ValidatorMgrSuite) TestRegister_NotPaused_NewCandidate_WithMinDeposit_NotFull_DepositLessOrEqual() {
	req := suite.Require()

	// register validator 1
	registerOpts1 := bind.NewKeyedTransactor(owner)
	from1 := crypto.PubkeyToAddress(owner.PublicKey)
	value1 := suite.baseDeposit
	_, err := suite.validatorMgr.RegisterValidator(registerOpts1, from1, value1, []byte{})
	req.NoError(err)

	// register validator 2
	registerOpts2 := bind.NewKeyedTransactor(user)
	from2 := crypto.PubkeyToAddress(user.PublicKey)
	value2 := value1
	_, err = suite.validatorMgr.RegisterValidator(registerOpts2, from2, value2, []byte{})
	req.NoError(err)

	suite.backend.Commit()

	validatorCount, err := suite.validatorMgr.GetValidatorCount(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(validatorCount)
	req.Equal(common.Big2, validatorCount)

	highestBidder, err := suite.validatorMgr.GetValidatorAtIndex(&bind.CallOpts{}, common.Big0)
	req.NoError(err)
	req.NotZero(highestBidder)
	req.Equal(highestBidder.Code, from1)
	req.Equal(highestBidder.Deposit, value1)

	lowestBidder, err := suite.validatorMgr.GetValidatorAtIndex(&bind.CallOpts{}, common.Big1)
	req.NoError(err)
	req.NotZero(highestBidder)
	req.Equal(lowestBidder.Code, from2)
	req.Equal(lowestBidder.Deposit, value2)
}

func (suite *ValidatorMgrSuite) TestRegister_NotPaused_NewCandidate_WithMinDeposit_Full_Replacement() {
	req := suite.Require()

	// register validator to match the max number of validators
	registerOpts1 := bind.NewKeyedTransactor(owner)
	from1 := crypto.PubkeyToAddress(owner.PublicKey)
	value1 := suite.baseDeposit
	_, err := suite.validatorMgr.RegisterValidator(registerOpts1, from1, value1, []byte{})
	req.NoError(err)

	suite.backend.Commit()

	// register validator 2 based on the minimum deposit required to participate
	minDeposit, err := suite.validatorMgr.GetMinimumDeposit(&bind.CallOpts{})
	req.NoError(err)
	registerOpts2 := bind.NewKeyedTransactor(user)
	from2 := crypto.PubkeyToAddress(user.PublicKey)
	value2 := minDeposit
	_, err = suite.validatorMgr.RegisterValidator(registerOpts2, from2, value2, []byte{})
	req.NoError(err)

	suite.backend.Commit()

	// the new validator must replace the existing one
	validatorCount, err := suite.validatorMgr.GetValidatorCount(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(validatorCount)
	req.Equal(common.Big1, validatorCount)

	storedValidator, err := suite.validatorMgr.GetValidatorAtIndex(&bind.CallOpts{}, common.Big0)
	req.NoError(err)
	req.NotZero(storedValidator)
	req.Equal(storedValidator.Code, from2)
	req.Equal(storedValidator.Deposit, value2)
}

func (suite *ValidatorMgrSuite) TestDeregister_WhenPaused() {
	req := suite.Require()

	pauseOpts := bind.NewKeyedTransactor(owner)
	_, err := suite.validatorMgr.Pause(pauseOpts)
	req.NoError(err)

	suite.backend.Commit()

	deregisterOpts := bind.NewKeyedTransactor(owner)
	_, err = suite.validatorMgr.DeregisterValidator(deregisterOpts)
	req.Error(err, "cannot deregister because the service is paused")
}

func (suite *ValidatorMgrSuite) TestDeregister_NotPaused_NotValidator() {
	req := suite.Require()

	deregisterOpts := bind.NewKeyedTransactor(owner)
	_, err := suite.validatorMgr.DeregisterValidator(deregisterOpts)
	req.Error(err, "cannot deregister a validator that does not exist")
}

func (suite *ValidatorMgrSuite) TestDeregister_NotPaused_Validator() {
	req := suite.Require()

	// register validator
	registerOpts := bind.NewKeyedTransactor(owner)
	from := crypto.PubkeyToAddress(owner.PublicKey)
	value := suite.baseDeposit
	_, err := suite.validatorMgr.RegisterValidator(registerOpts, from, value, []byte{})
	req.NoError(err)

	suite.backend.Commit()

	_, err = suite.validatorMgr.DeregisterValidator(registerOpts)
	req.NoError(err)

	suite.backend.Commit()

	// there's a release date
	deposit, err := suite.validatorMgr.GetDepositAtIndex(&bind.CallOpts{From: from}, common.Big0)
	req.NoError(err)
	req.True(deposit.AvailableAt.Cmp(common.Big0) > 0)
}

func (suite *ValidatorMgrSuite) TestReleaseDeposits_WhenPaused() {
	req := suite.Require()

	pauseOpts := bind.NewKeyedTransactor(owner)
	_, err := suite.validatorMgr.Pause(pauseOpts)
	req.NoError(err)

	suite.backend.Commit()

	releaseOpts := bind.NewKeyedTransactor(owner)
	_, err = suite.validatorMgr.ReleaseDeposits(releaseOpts)
	req.Error(err, "cannot release deposits when the service is paused")
}

func (suite *ValidatorMgrSuite) TestReleaseDeposits_NotPaused_NoDeposits() {
	req := suite.Require()

	releaseOpts := bind.NewKeyedTransactor(owner)
	_, err := suite.validatorMgr.ReleaseDeposits(releaseOpts)
	req.NoError(err)

	suite.backend.Commit()

	mtokens, err := suite.tokenMock.BalanceOf(&bind.CallOpts{}, crypto.PubkeyToAddress(owner.PublicKey))
	req.NoError(err)
	req.Zero(mtokens.Uint64())
}

func (suite *ValidatorMgrSuite) TestReleaseDeposits_NotPaused_LockedDeposits() {
	req := suite.Require()

	// register validator
	registerOpts := bind.NewKeyedTransactor(owner)
	from := crypto.PubkeyToAddress(owner.PublicKey)
	value := suite.baseDeposit
	_, err := suite.validatorMgr.RegisterValidator(registerOpts, from, value, []byte{})
	req.NoError(err)

	suite.backend.Commit()

	// freeze period is > 0
	_, err = suite.validatorMgr.DeregisterValidator(registerOpts)
	req.NoError(err)

	suite.backend.Commit()

	releaseOpts := bind.NewKeyedTransactor(owner)
	_, err = suite.validatorMgr.ReleaseDeposits(releaseOpts)
	req.NoError(err)

	suite.backend.Commit()

	mtokens, err := suite.tokenMock.BalanceOf(&bind.CallOpts{}, crypto.PubkeyToAddress(owner.PublicKey))
	req.NoError(err)
	req.Zero(mtokens.Uint64())
}

func (suite *ValidatorMgrSuite) TestReleaseDeposits_UnlockedDeposit() {
	req := suite.Require()

	// register validator
	registerOpts := bind.NewKeyedTransactor(owner)
	from := crypto.PubkeyToAddress(owner.PublicKey)
	value := suite.baseDeposit
	_, err := suite.validatorMgr.RegisterValidator(registerOpts, from, value, []byte{})
	req.NoError(err)

	suite.backend.Commit()

	// freeze period is == 0
	_, err = suite.validatorMgr.DeregisterValidator(registerOpts)
	req.NoError(err)

	suite.backend.Commit()

	releaseOpts := bind.NewKeyedTransactor(owner)
	_, err = suite.validatorMgr.ReleaseDeposits(releaseOpts)
	req.NoError(err)

	suite.backend.Commit()

	mtokens, err := suite.tokenMock.BalanceOf(&bind.CallOpts{}, crypto.PubkeyToAddress(owner.PublicKey))
	req.NoError(err)
	req.Equal(value, mtokens)
}

// dtos converts days to seconds
func dtos(numberOfDays *big.Int) *big.Int {
	return new(big.Int).Mul(numberOfDays, secondsPerDay)
}
