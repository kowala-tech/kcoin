package genesis

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/kowala-tech/kcoin/client/accounts/abi"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/consensus"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/oracle"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/ownership"
	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/core/vm/runtime"
)

type contract struct {
	name       string
	runtimeCfg *runtime.Config
	address    common.Address
	storage    map[common.Hash]common.Hash
	code       []byte
	deploy     func(contract *contract, opts *validGenesisOptions) error
	postDeploy func(contract *contract, opts *validGenesisOptions) error
}

func (contract *contract) AsGenesisAccount() core.GenesisAccount {
	return core.GenesisAccount{
		Code:    contract.code,
		Storage: contract.storage,
		Balance: new(big.Int),
	}
}

var MultiSigContract = &contract{
	name: "MultiSigWallet",
	deploy: func(contract *contract, opts *validGenesisOptions) error {
		args := opts.multiSig

		multiSigWalletABI, err := abi.JSON(strings.NewReader(ownership.MultiSigWalletABI))
		if err != nil {
			return err
		}

		multiSigParams, err := multiSigWalletABI.Pack(
			"",
			args.multiSigOwners,
			args.numConfirmations,
		)
		if err != nil {
			return err
		}

		runtimeCfg := contract.runtimeCfg
		runtimeCfg.Origin = *args.multiSigCreator
		contractCode, contractAddr, _, err := runtime.Create(append(common.FromHex(ownership.MultiSigWalletBin), multiSigParams...), runtimeCfg)
		if err != nil {
			return err
		}
		contract.code = contractCode
		contract.address = contractAddr

		opts.miningToken.owner = contractAddr
		opts.validatorMgr.owner = contractAddr
		opts.oracleMgr.owner = contractAddr

		return nil
	},
}

var MiningTokenContract = &contract{
	name: "Mining Token",
	deploy: func(contract *contract, opts *validGenesisOptions) error {
		args := opts.miningToken

		tokenABI, err := abi.JSON(strings.NewReader(consensus.MiningTokenABI))
		if err != nil {
			return err
		}

		tokenParams, err := tokenABI.Pack(
			"",
			args.name,
			args.symbol,
			args.cap,
			uint8(args.decimals.Uint64()),
		)
		if err != nil {
			return err
		}

		runtimeCfg := contract.runtimeCfg
		runtimeCfg.Origin = args.owner
		contractCode, contractAddr, _, err := runtime.Create(append(common.FromHex(consensus.MiningTokenBin), tokenParams...), runtimeCfg)
		if err != nil {
			return err
		}
		contract.code = contractCode
		contract.address = contractAddr

		opts.validatorMgr.miningTokenAddr = contractAddr

		return nil
	},
	postDeploy: mintTokens,
}

func mintTokens(contract *contract, opts *validGenesisOptions) error {
	args := opts.miningToken

	tokenABI, err := abi.JSON(strings.NewReader(consensus.MiningTokenABI))
	if err != nil {
		return err
	}

	runtimeCfg := contract.runtimeCfg
	runtimeCfg.Origin = args.owner
	for _, holder := range args.holders {
		mintParams, err := tokenABI.Pack(
			"mint",
			holder.address,
			holder.balance,
		)
		if err != nil {
			return err
		}

		_, _, err = runtime.Call(contract.address, mintParams, runtimeCfg)
		if err != nil {
			return fmt.Errorf("%s:%s", "Failed to mint tokens", err)
		}
	}

	return nil
}

var OracleMgrContract = &contract{
	name: "Oracle Manager",
	deploy: func(contract *contract, opts *validGenesisOptions) error {
		args := opts.oracleMgr

		managerABI, err := abi.JSON(strings.NewReader(oracle.OracleMgrABI))
		if err != nil {
			return err
		}

		managerParams, err := managerABI.Pack(
			"",
			args.price.initialPrice,
			args.baseDeposit,
			args.maxNumOracles,
			args.freezePeriod,
			args.price.syncFrequency,
			args.price.updatePeriod,
			args.validatorMgrAddr,
		)
		if err != nil {
			return err
		}

		runtimeCfg := contract.runtimeCfg
		runtimeCfg.Origin = args.owner
		contractCode, contractAddr, _, err := runtime.Create(append(common.FromHex(oracle.OracleMgrBin), managerParams...), runtimeCfg)
		if err != nil {
			return err
		}
		contract.code = contractCode
		contract.address = contractAddr

		return nil
	},
}

var ValidatorMgrContract = &contract{
	name: "Validator Manager",
	deploy: func(contract *contract, opts *validGenesisOptions) error {
		args := opts.validatorMgr

		managerABI, err := abi.JSON(strings.NewReader(consensus.ValidatorMgrABI))
		if err != nil {
			return err
		}

		managerParams, err := managerABI.Pack(
			"",
			args.baseDeposit,
			args.maxNumValidators,
			args.freezePeriod,
			args.miningTokenAddr,
			args.superNodeAmount,
		)
		if err != nil {
			return err
		}

		runtimeCfg := contract.runtimeCfg
		runtimeCfg.Origin = args.owner
		contractCode, contractAddr, _, err := runtime.Create(append(common.FromHex(consensus.ValidatorMgrBin), managerParams...), runtimeCfg)
		if err != nil {
			return err
		}
		contract.code = contractCode
		contract.address = contractAddr

		opts.oracleMgr.validatorMgrAddr = contractAddr

		return nil
	},
	postDeploy: registerValidators,
}

func registerValidators(contract *contract, opts *validGenesisOptions) error {
	args := opts.validatorMgr

	// register genesis validators
	tokenABI, err := abi.JSON(strings.NewReader(consensus.MiningTokenABI))
	if err != nil {
		return err
	}

	runtimeCfg := contract.runtimeCfg
	for _, validator := range args.validators {
		runtimeCfg.Origin = validator.address

		registrationParams, err := tokenABI.Pack(
			"transfer",
			contract.address,
			validator.deposit,
			[]byte("not_zero"), // @NOTE (rgeraldes) - https://github.com/kowala-tech/kcoin/client/issues/285
			consensus.RegistrationHandler,
		)
		if err != nil {
			return err
		}

		_, _, err = runtime.Call(args.miningTokenAddr, registrationParams, runtimeCfg)
		if err != nil {
			return fmt.Errorf("%s:%s", "Failed to register validator", err)
		}
	}

	return nil
}
