package genesis

import (
	"math/big"

	"github.com/kowala-tech/kcoin/common"
)

// @TODO (rgeraldes) - verify if it matters having pointer or not

type Options struct {
	Network           string
	Governance        GovernanceOpts
	Consensus         ConsensusOpts
	DataFeedSystem    DataFeedSystemOpts
	PrefundedAccounts []PrefundedAccount
	ExtraData         string
}

type ConsensusOpts struct {
	Paused           bool
	MaxNumValidators uint64
	UnbondingPeriod  string
	BaseDeposit      string
	Validators       []string
}

type GovernanceOpts struct {
	Origin           string
	Governors        []string
	NumConfirmations uint64
}

type DataFeedSystemOpts struct {
	Paused          bool
	MaxNumOracles   string
	UnbondingPeriod string // in days
	baseDeposit     string // in kUSD
}

type PrefundedAccount struct {
	AccountAddress string
	Balance        string
}

type validValidatorMgrOpts struct {
	paused           bool
	maxNumValidators *big.Int
	unbondingPeriod  *big.Int
	baseDeposit      *big.Int
	validators       []common.Address
}

type validOracleMgrOpts struct {
	paused          *big.Int
	maxNumOracles   *big.Int
	unbondingPeriod *big.Int
	baseDeposit     *big.Int
}

type validMultiSigOpts struct {
	multiSigCreator  *common.Address
	multiSigOwners   []common.Address
	numConfirmations *big.Int
}

type validPrefundedAccount struct {
	accountAddress *common.Address
	balance        *big.Int
}

type validGenesisOptions struct {
	network           string
	validatorMgr      *validValidatorMgrOpts
	oracleMgr         *validOracleMgrOpts
	multiSig          *validMultiSigOpts
	prefundedAccounts []*validPrefundedAccount
	consensusEngine   string
}

func validateOptions(options Options) (*validGenesisOptions, error) {
	network, err := mapNetwork(options.Network)
	if err != nil {
		return nil, err
	}

	consensusEngine := TendermintConsensus
	if options.ConsensusEngine != "" {
		consensusEngine, err = mapConsensusEngine(options.ConsensusEngine)
		if err != nil {
			return nil, err
		}
	}

	// governance
	multiSigCreator, err := getAddress(options.Governance.Origin)
	if err != nil {
		return nil, err
	}

	multiSigOwners = make([]common.Address, len(options.Governance.Governors))
	for _, governor := range options.Governance.Governors {
		multiSigOwners = append(multiSigOwners, governor)
	}

	numConfirmations := new(big.Int).SetUint64(options.Governance.NumConfirmations)

	// consensus
	maxNumValidators := new(big.Int).SetUint64(options.Consensus.MaxNumValidators)

	consensusBaseDeposit := new(big.Int).SetUint64(options.Consensus.BaseDeposit)

	consensusUnbondingPeriod := new(big.Int).SetUint64(options.Consensus.UnbondingPeriod)

	validators = make([]common.Address, len(options.Consensus.Validators))
	for _, validator := range options.Consensus.Validators {
		validator, err := getAddress(validator)
		if err != nil {
			return nil, err
		}
	}

	// data feed system
	maxNumOracles := new(big.Int).SetUint64(options.DataFeedSystem.MaxNumOracles)

	oracleBaseDeposit := new(big.Int).SetUint64(options.DataFeedSystem.BaseDeposit)

	oracleUnbondingPeriod := new(big.Int).SetUint64(options.DataFeedSystem.UnbondingPeriod)

	// optional
	validPrefundedAccounts, err := mapPrefundedAccounts(options.PrefundedAccounts)
	if err != nil {
		return nil, err
	}

	if !prefundedIncludesValidatorWallet(validPrefundedAccounts, walletAddressValidator) {
		return nil, ErrWalletAddressValidatorNotInPrefundedAccounts
	}

	return &validGenesisOptions{
		network:         network,
		consensusEngine: consensusEngine,
		multiSig: &validMultiSigOpts{
			multiSigCreator: multiSigCreator,
			multiSigOwners, multiSigOwners,
			numConfirmations: numConfirmations,
		},
		validatorMgr: &validValidatorMgrOpts{
			paused:           options.Consensus.Paused,
			maxNumValidators: maxNumValidators,
			unbondingPeriod:  consensusUnbondingPeriod,
			baseDeposit:      consensusBaseDeposit,
			validators:       validators,
		},
		oracleMgrOpts: &validOracleMgrOpts{
			paused:          options.DataFeedSystem.Paused,
			maxNumOracles:   maxNumOracles,
			unbondingPeriod: oracleUnbondingPeriod,
			baseDeposit:     oracleBaseDeposit,
		},
		prefundedAccounts:   validPrefundedAccounts,
		smartContractsOwner: owner,
	}, nil
}

func getAddress(s string) (*common.Address, error) {
	if !common.IsHexAddress(s) {
		return nil, ErrInvalidAddress
	}
	return &common.HexToAddress(s), nil
}
