package genesis

type Options struct {
	Network                        	string
	ValidatorMgrOpts		   		ValidatorMgrOpts
	OracleMgrOpts					OracleMgrOpts
	AccountAddressGenesisValidator string
	PrefundedAccounts              []PrefundedAccount
	ConsensusEngine                string
	MultiSigCreator string // should be the same every time (during testing) in order to maintain the same addresses for core contracts
	MultiSigOwners					[]string
	ExtraData                      string
}

type ValidatorMgrOpts struct {
	MaxNumValidators string
	UnbondingPeriod  string // in days
	BaseDeposit string // in mUSD
}

type OracleMgrOpts struct {
	MaxNumOracles string
	UnbondingPeriod string // in days
	baseDeposit string // in kUSD
}

type PrefundedAccount struct {
	AccountAddress string
	Balance        string
}

type validValidatorMgrOpts struct {
	maxNumValidators *big.Int
	unbondingPeriod  *big.Int
	baseDeposit *big.Int
}

type validOracleMgrOpts struct {
	maxNumValidators *big.Int
	unbondingPeriod  *big.Int
	baseDeposit *big.Int
}

type validPrefundedAccount struct {
	accountAddress *common.Address
	balance        *big.Int
}

type validGenesisOptions struct {
	network                        string
	validatorMgrOpts			    validValidatorMgrOpts
	OracleMgrOpts					validOracleMgrOpts
	accountAddressGenesisValidator *common.Address
	prefundedAccounts              []*validPrefundedAccount
	consensusEngine                string
	// @TODO (rgeraldes) - replace for non pointers
	multiSigCreator *common.Address
	multiSigOwners           []common.Address
}

func validateOptions(options Options) (*validGenesisOptions, error) {
	network, err := mapNetwork(options.Network)
	if err != nil {
		return nil, err
	}

	maxNumValidators, err := mapMaxNumValidators(options.MaxNumValidators)
	if err != nil {
		return nil, err
	}

	unbondingPeriod, err := mapUnbondingPeriod(options.UnbondingPeriod)
	if err != nil {
		return nil, err
	}

	walletAddressValidator, err := mapWalletAddress(options.AccountAddressGenesisValidator)
	if err != nil {
		return nil, err
	}

	validPrefundedAccounts, err := mapPrefundedAccounts(options.PrefundedAccounts)
	if err != nil {
		return nil, err
	}

	if !prefundedIncludesValidatorWallet(validPrefundedAccounts, walletAddressValidator) {
		return nil, ErrWalletAddressValidatorNotInPrefundedAccounts
	}

	consensusEngine := TendermintConsensus
	if options.ConsensusEngine != "" {
		consensusEngine, err = mapConsensusEngine(options.ConsensusEngine)
		if err != nil {
			return nil, err
		}
	}

	owner := &DefaultSmartContractsOwner
	if options.SmartContractsOwner != "" {
		strAddr := options.SmartContractsOwner

		owner, err = mapWalletAddress(strAddr)
		if err != nil {
			return nil, ErrInvalidContractsOwnerAddress
		}
	}

	return &validGenesisOptions{
		network:                        network,
		maxNumValidators:               maxNumValidators,
		unbondingPeriod:                unbondingPeriod,
		accountAddressGenesisValidator: walletAddressValidator,
		prefundedAccounts:              validPrefundedAccounts,
		consensusEngine:                consensusEngine,
		smartContractsOwner:            owner,
	}, nil
}