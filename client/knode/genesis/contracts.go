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
	"github.com/kowala-tech/kcoin/client/contracts/bindings/proxy"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/kns"
	kns2 "github.com/kowala-tech/kcoin/client/kns"
)

var ProxyFactoryAddr = "0xfE9bed356E7bC4f7a8fC48CC19C958f4e640AC62"
var ProxyKNSRegistryAddr = "0x6f04441A6eD440Cc139a4E33402b438C27E97F4B"

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

var KNSRegistry = &contract{
	name: "KNSRegistry",
	deploy: func(contract *contract, opts *validGenesisOptions) error {
		args := opts.multiSig

		runtimeCfg := contract.runtimeCfg
		runtimeCfg.Origin = *args.multiSigCreator
		contractCode, contractAddr, _, err := runtime.Create(common.FromHex(kns.KNSRegistryBin), runtimeCfg)
		if err != nil {
			return err
		}

		contract.code = contractCode
		contract.address = contractAddr

		return nil
	},
}

var ProxiedKNSRegistry = &contract{
	name: "ProxiedKNSRegistry",
	deploy: func(contract *contract, opts *validGenesisOptions) error {
		args := opts.multiSig

		runtimeCfg := contract.runtimeCfg
		runtimeCfg.Origin = *args.multiSigCreator

		proxyABI, err := abi.JSON(strings.NewReader(proxy.UpgradeabilityProxyFactoryABI))
		if err != nil {
			return err
		}

		createProxyArgs, err := proxyABI.Pack(
			"createProxy",
			*args.multiSigCreator,
			common.HexToAddress("0x1582aEd4A8156325e28ef9eF075Da1E1D44AA56E"),
		)
		if err != nil{
			return err
		}

		ret, _, err := runtime.Call(common.HexToAddress(ProxyFactoryAddr), createProxyArgs, runtimeCfg)
		if err != nil {
			return fmt.Errorf("%s:%s", "Failed to create proxy for KNS", err)
		}

		knsProxiedAddress := common.BytesToAddress(ret)
		contract.address = knsProxiedAddress
		contract.code = contract.runtimeCfg.State.GetCode(knsProxiedAddress)

		return nil
	},
	postDeploy: func(contract *contract, opts *validGenesisOptions) error {
		validatorAddr := opts.prefundedAccounts[0].accountAddress

		runtimeCfg := contract.runtimeCfg
		runtimeCfg.Origin = *validatorAddr

		abi, err := abi.JSON(strings.NewReader(kns.KNSRegistryABI))
		if err != nil {
			return err
		}

		//TODO (jgimeno) for now is the validator coming from the testnet.
		initKnsParams, err := abi.Pack(
			"initialize",
			*validatorAddr,
		)
		if err != nil {
			return err
		}

		_, _, err = runtime.Call(contract.address, initKnsParams, runtimeCfg)
		if err != nil {
			return fmt.Errorf("%s:%s", "Failed to initialize KNSRegistry.", err)
		}

		return nil
	},
}

var FIFSRegistrar = &contract{
	name: "FIFSRegistrar",
	deploy: func(contract *contract, opts *validGenesisOptions) error {
		args := opts.multiSig

		runtimeCfg := contract.runtimeCfg
		runtimeCfg.Origin = *args.multiSigCreator
		contractCode, contractAddr, _, err := runtime.Create(common.FromHex(kns.FIFSRegistrarBin), runtimeCfg)
		if err != nil {
			return err
		}

		contract.code = contractCode
		contract.address = contractAddr

		return nil
	},
}

var ProxiedFIFSRegistrar = &contract{
	name: "ProxiedFIFSRegistrar",
	deploy: func(contract *contract, opts *validGenesisOptions) error {
		args := opts.multiSig

		runtimeCfg := contract.runtimeCfg
		runtimeCfg.Origin = *args.multiSigCreator

		proxyABI, err := abi.JSON(strings.NewReader(proxy.UpgradeabilityProxyFactoryABI))
		if err != nil {
			return err
		}

		createProxyArgs, err := proxyABI.Pack(
			"createProxy",
			*args.multiSigCreator,
			common.HexToAddress("0x75AD571eFAcC241B23099c724c4A71FE30659145"),
		)
		if err != nil{
			return err
		}

		ret, _, err := runtime.Call(common.HexToAddress(ProxyFactoryAddr), createProxyArgs, runtimeCfg)
		if err != nil {
			return fmt.Errorf("%s:%s", "Failed to create proxy for FIFSRegistrar", err)
		}

		registrarProxiedAddr := common.BytesToAddress(ret)
		contract.address = registrarProxiedAddr
		contract.code = contract.runtimeCfg.State.GetCode(registrarProxiedAddr)

		return nil
	},
	postDeploy: func(contract *contract, opts *validGenesisOptions) error {
		validatorAddr := opts.prefundedAccounts[0].accountAddress

		runtimeCfg := contract.runtimeCfg
		runtimeCfg.Origin = *validatorAddr

		abi, err := abi.JSON(strings.NewReader(kns.FIFSRegistrarABI))
		if err != nil {
			return err
		}

		//TODO (jgimeno) for now is the validator coming from the testnet.
		initKnsParams, err := abi.Pack(
			"initialize",
			common.HexToAddress(ProxyKNSRegistryAddr),
			kns2.NameHash("tech"),
		)
		if err != nil {
			return err
		}

		_, _, err = runtime.Call(contract.address, initKnsParams, runtimeCfg)
		if err != nil {
			return fmt.Errorf("%s:%s", "Failed to initialize KNSRegistry.", err)
		}

		return nil
	},
}

var PublicResolver = &contract{
	name: "PublicResolver",
	deploy: func(contract *contract, opts *validGenesisOptions) error {
		args := opts.multiSig

		runtimeCfg := contract.runtimeCfg
		runtimeCfg.Origin = *args.multiSigCreator
		contractCode, contractAddr, _, err := runtime.Create(common.FromHex(kns.PublicResolverBin), runtimeCfg)
		if err != nil {
			return err
		}

		contract.code = contractCode
		contract.address = contractAddr

		return nil
	},
}

var ProxiedPublicResolver = &contract{
	name: "ProxiedPublicResolver",
	deploy: func(contract *contract, opts *validGenesisOptions) error {
		args := opts.multiSig

		runtimeCfg := contract.runtimeCfg
		runtimeCfg.Origin = *args.multiSigCreator

		proxyABI, err := abi.JSON(strings.NewReader(proxy.UpgradeabilityProxyFactoryABI))
		if err != nil {
			return err
		}

		createProxyArgs, err := proxyABI.Pack(
			"createProxy",
			*args.multiSigCreator,
			common.HexToAddress("0x2A4443ec27BF5F849B2Da15eB697d3Ef5302f186"),
		)
		if err != nil{
			return err
		}

		ret, _, err := runtime.Call(common.HexToAddress(ProxyFactoryAddr), createProxyArgs, runtimeCfg)
		if err != nil {
			return fmt.Errorf("%s:%s", "Failed to create proxy for PublicResolver", err)
		}

		registrarProxiedAddr := common.BytesToAddress(ret)
		contract.address = registrarProxiedAddr
		contract.code = contract.runtimeCfg.State.GetCode(registrarProxiedAddr)

		return nil
	},
}

var UpgradeabilityProxyFactoryContract = &contract{
	name: "UpgradeabilityProxyFactoryContract",
	deploy: func(contract *contract, opts *validGenesisOptions) error {
		args := opts.multiSig

		runtimeCfg := contract.runtimeCfg
		runtimeCfg.Origin = *args.multiSigCreator
		contractCode, contractAddr, _, err := runtime.Create(common.FromHex(proxy.UpgradeabilityProxyFactoryBin), runtimeCfg)
		if err != nil {
			return err
		}

		contract.code = contractCode
		contract.address = contractAddr

		return nil
	},
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
