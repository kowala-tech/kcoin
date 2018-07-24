package sysvars

/*
func (suite *OracleMgrSuite) TestDeployOracleMgr_InitialPriceEqualsZero() {
	req := suite.Require()

	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		getAddress(governor): core.GenesisAccount{
			Balance: new(big.Int).Mul(new(big.Int).SetUint64(100), new(big.Int).SetUint64(params.Kcoin)),
		},
	})
	req.NotNil(backend)

	initialPrice := common.Big0
	baseDeposit := new(big.Int).SetUint64(100)
	maxNumOracles := new(big.Int).SetUint64(100)
	freezePeriod := new(big.Int).SetUint64(10)
	syncFrequency := new(big.Int).SetUint64(20)
	updatePeriod := new(big.Int).SetUint64(5)

	transactOpts := bind.NewKeyedTransactor(governor)
	_, _, _, err := oracle.DeployOracleMgr(transactOpts, backend, initialPrice, baseDeposit, maxNumOracles, freezePeriod, syncFrequency, updatePeriod, validatorMgrAddr)
	req.Error(err, "initial price cannot be zero")
}
*/