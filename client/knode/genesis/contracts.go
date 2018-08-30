package genesis

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/kowala-tech/kcoin/client/accounts/abi"
	"github.com/kowala-tech/kcoin/client/common"
	kns2 "github.com/kowala-tech/kcoin/client/common/kns"
	"github.com/kowala-tech/kcoin/client/contracts/bindings"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/consensus"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/kns"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/oracle"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/ownership"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/proxy"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/stability"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/sysvars"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/utils"
	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/core/vm/runtime"
	"github.com/kowala-tech/kcoin/client/crypto"
	"github.com/kowala-tech/kcoin/client/params"
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

var SystemVarsContract = &contract{
	name: "SystemVars",
	deploy: func(contract *contract, opts *validGenesisOptions) error {
		args := opts.sysvars

		systemVarsABI, err := abi.JSON(strings.NewReader(sysvars.SystemVarsABI))
		if err != nil {
			return err
		}

		systemVarsParams, err := systemVarsABI.Pack(
			"",
			args.initialPrice,
			args.initialSupply,
		)
		if err != nil {
			return err
		}

		runtimeCfg := contract.runtimeCfg
		runtimeCfg.Origin = args.owner
		contractCode, contractAddr, _, err := runtime.Create(append(common.FromHex(sysvars.SystemVarsBin), systemVarsParams...), runtimeCfg)
		if err != nil {
			return err
		}
		contract.code = contractCode
		contract.address = contractAddr

		opts.stability.systemVarsAddr = contractAddr

		return nil
	},
	postDeploy: func(contract *contract, opts *validGenesisOptions) error {
		return registerAddressToDomain(contract, opts, params.KNSDomains[params.SystemVarsDomain].Node())
	},
}

var StabilityContract = &contract{
	name: "Stability",
	deploy: func(contract *contract, opts *validGenesisOptions) error {
		args := opts.stability

		stabilityABI, err := abi.JSON(strings.NewReader(stability.StabilityABI))
		if err != nil {
			return err
		}

		stabilityParams, err := stabilityABI.Pack(
			"",
			args.minDeposit,
			args.systemVarsAddr,
		)
		if err != nil {
			return err
		}

		runtimeCfg := contract.runtimeCfg
		runtimeCfg.Origin = args.owner
		contractCode, contractAddr, _, err := runtime.Create(append(common.FromHex(stability.StabilityBin), stabilityParams...), runtimeCfg)
		if err != nil {
			return err
		}
		contract.code = contractCode
		contract.address = contractAddr

		return nil
	},
	postDeploy: func(contract *contract, opts *validGenesisOptions) error {
		return registerAddressToDomain(contract, opts, params.KNSDomains[params.StabilityDomain].Node())
	},
}

var KNSRegistry = &contract{
	name: "KNSRegistry",
	deploy: func(contract *contract, opts *validGenesisOptions) error {
		runtimeCfg := contract.runtimeCfg
		runtimeCfg.Origin = bindings.MultiSigWalletAddr
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

		proxyContractAddr, code, err := createProxyFromContract(
			common.HexToAddress("0xBFb47D8008d1ccDCAF3a36110a9338a274e86343"),
			*args.multiSigCreator,
			runtimeCfg,
		)
		if err != nil {
			return err
		}

		contract.address = *proxyContractAddr
		contract.code = code

		// Init Registry
		validatorAddr := opts.prefundedAccounts[0].accountAddress

		runtimeCfg.Origin = *validatorAddr
		abi, err := abi.JSON(strings.NewReader(kns.KNSRegistryABI))
		if err != nil {
			return err
		}

		initKnsParams, err := abi.Pack(
			"initialize",
			bindings.MultiSigWalletAddr,
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
	postDeploy: func(contract *contract, opts *validGenesisOptions) error {
		// call setSubnodeOwner
		runtimeCfg := contract.runtimeCfg
		runtimeCfg.Origin = bindings.MultiSigWalletAddr

		abi, err := abi.JSON(strings.NewReader(kns.KNSRegistryABI))
		if err != nil {
			return err
		}

		subnodeOwnerParams, err := abi.Pack(
			"setSubnodeOwner",
			[32]byte{},
			crypto.Keccak256Hash([]byte("kowala")),
			bindings.ProxyRegistrarAddr,
		)
		if err != nil {
			return err
		}

		_, _, err = runtime.Call(contract.address, subnodeOwnerParams, runtimeCfg)
		if err != nil {
			return fmt.Errorf("%s:%s", "Failed to set subnode owner in KNSRegistry.", err)
		}

		return nil
	},
}

var FIFSRegistrar = &contract{
	name: "FIFSRegistrar",
	deploy: func(contract *contract, opts *validGenesisOptions) error {
		runtimeCfg := contract.runtimeCfg
		runtimeCfg.Origin = bindings.MultiSigWalletAddr
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

		proxyContractAddr, code, err := createProxyFromContract(
			common.HexToAddress("0xe95d0D373E2FD320b84aaC705434b67B905092aE"),
			*args.multiSigCreator,
			runtimeCfg,
		)
		if err != nil {
			return err
		}

		contract.address = *proxyContractAddr
		contract.code = code

		validatorAddr := opts.prefundedAccounts[0].accountAddress
		runtimeCfg.Origin = *validatorAddr

		abi, err := abi.JSON(strings.NewReader(kns.FIFSRegistrarABI))
		if err != nil {
			return err
		}

		initKnsParams, err := abi.Pack(
			"initialize",
			bindings.ProxyKNSRegistryAddr,
			kns2.NameHash("kowala"),
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
		runtimeCfg := contract.runtimeCfg
		runtimeCfg.Origin = bindings.MultiSigWalletAddr
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

		proxyContractAddr, code, err := createProxyFromContract(
			common.HexToAddress("0x7F0de05687a7cb9a05399a26f4D1519Ba6Afc95F"),
			*args.multiSigCreator,
			runtimeCfg,
		)
		if err != nil {
			return err
		}

		contract.address = *proxyContractAddr
		contract.code = code

		runtimeCfg.Origin = bindings.MultiSigWalletAddr

		abi, err := abi.JSON(strings.NewReader(kns.PublicResolverABI))
		if err != nil {
			return err
		}

		initKnsParams, err := abi.Pack(
			"initialize",
			bindings.ProxyKNSRegistryAddr,
		)
		if err != nil {
			return err
		}

		_, _, err = runtime.Call(contract.address, initKnsParams, runtimeCfg)
		if err != nil {
			return fmt.Errorf("%s:%s", "Failed to initialize PublicResolver.", err)
		}

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

func createProxyFromContract(contractAddr common.Address, accountCreator common.Address, runtimeCfg *runtime.Config) (proxyContractAddr *common.Address, code []byte, err error) {
	runtimeCfg.Origin = accountCreator

	proxyABI, err := abi.JSON(strings.NewReader(proxy.UpgradeabilityProxyFactoryABI))
	if err != nil {
		return nil, nil, err
	}

	createProxyArgs, err := proxyABI.Pack(
		"createProxy",
		accountCreator,
		contractAddr,
	)
	if err != nil {
		return nil, nil, err
	}

	ret, _, err := runtime.Call(bindings.ProxyFactoryAddr, createProxyArgs, runtimeCfg)
	if err != nil {
		return nil, nil, fmt.Errorf("%s:%s", "Failed to create proxy for KNS", err)
	}

	knsProxiedAddress := common.BytesToAddress(ret)

	return &knsProxiedAddress, runtimeCfg.State.GetCode(knsProxiedAddress), nil
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
		opts.sysvars.owner = contractAddr
		opts.stability.owner = contractAddr

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

		return nil
	},
}

func mintTokens(contract *contract, opts *validGenesisOptions) error {
	args := opts.miningToken

	tokenABI, err := abi.JSON(strings.NewReader(consensus.MiningTokenABI))
	if err != nil {
		return err
	}

	runtimeCfg := contract.runtimeCfg
	runtimeCfg.Origin = bindings.MultiSigWalletAddr
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

var ProxiedMiningToken = &contract{
	name: "Proxied Mining Token",
	deploy: func(contract *contract, opts *validGenesisOptions) error {
		args := opts.miningToken

		runtimeCfg := contract.runtimeCfg

		proxyContractAddr, code, err := createProxyFromContract(
			common.HexToAddress("0x64c2a9CB0220D3e56783ed87cC1B20115Bc93F96"),
			*opts.multiSig.multiSigCreator,
			runtimeCfg,
		)
		if err != nil {
			return err
		}

		contract.address = *proxyContractAddr
		contract.code = code
		runtimeCfg.Origin = bindings.MultiSigWalletAddr
		abi, err := abi.JSON(strings.NewReader(consensus.MiningTokenABI))
		if err != nil {
			return err
		}

		initKnsParams, err := abi.Pack(
			"initialize",
			args.name,
			args.symbol,
			args.cap,
			uint8(args.decimals.Uint64()),
		)
		if err != nil {
			return err
		}

		_, _, err = runtime.Call(contract.address, initKnsParams, runtimeCfg)
		if err != nil {
			return fmt.Errorf("%s:%s", "Failed to initialize Proxied Mining Token.", err)
		}

		opts.validatorMgr.miningTokenAddr = *proxyContractAddr

		return nil
	},
	postDeploy: func(contract *contract, opts *validGenesisOptions) error {
		err := mintTokens(contract, opts)
		if err != nil {
			return err
		}

		return registerAddressToDomain(contract, opts, params.KNSDomains[params.MiningTokenDomain].Node())
	},
}

var StringsLibrary = &contract{
	name: "StringsLibrary",
	deploy: func(contract *contract, opts *validGenesisOptions) error {
		args := opts.multiSig

		runtimeCfg := contract.runtimeCfg
		runtimeCfg.Origin = *args.multiSigCreator
		contractCode, contractAddr, _, err := runtime.Create(common.FromHex(utils.StringsBin), runtimeCfg)
		if err != nil {
			return err
		}
		contract.code = contractCode
		contract.address = contractAddr

		return nil
	},
}

var NameHashLibrary = &contract{
	name: "NameHashLibrary",
	deploy: func(contract *contract, opts *validGenesisOptions) error {
		args := opts.multiSig

		runtimeCfg := contract.runtimeCfg
		runtimeCfg.Origin = *args.multiSigCreator
		contractCode, contractAddr, _, err := runtime.Create(common.FromHex(utils.NameHashBin), runtimeCfg)
		if err != nil {
			return err
		}
		contract.code = contractCode
		contract.address = contractAddr

		return nil
	},
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
			args.maxNumOracles,
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

var ProxiedOracleMgr = &contract{
	name: "Proxied Oracle Mgr",
	deploy: func(contract *contract, opts *validGenesisOptions) error {
		args := opts.oracleMgr

		runtimeCfg := contract.runtimeCfg

		proxyContractAddr, code, err := createProxyFromContract(
			common.HexToAddress("0x616a77BA32aDC911dBa37f5883e4013B5278a279"),
			*opts.multiSig.multiSigCreator,
			runtimeCfg,
		)
		if err != nil {
			return err
		}

		contract.address = *proxyContractAddr
		contract.code = code
		runtimeCfg.Origin = bindings.MultiSigWalletAddr
		abi, err := abi.JSON(strings.NewReader(oracle.OracleMgrABI))
		if err != nil {
			return err
		}

		initKnsParams, err := abi.Pack(
			"initialize",
			args.maxNumOracles,
			args.price.syncFrequency,
			args.price.updatePeriod,
			args.validatorMgrAddr,
		)
		if err != nil {
			return err
		}

		_, _, err = runtime.Call(contract.address, initKnsParams, runtimeCfg)
		if err != nil {
			return fmt.Errorf("%s:%s", "Failed to initialize Proxied Oracle Manager.", err)
		}

		return nil
	},
	postDeploy: func(contract *contract, opts *validGenesisOptions) error {
		return registerAddressToDomain(contract, opts, params.KNSDomains[params.OracleMgrDomain].Node())
	},
}

func registerAddressToDomain(contract *contract, opts *validGenesisOptions, domain string) error {
	validatorAddr := opts.prefundedAccounts[0].accountAddress

	runtimeCfg := contract.runtimeCfg
	runtimeCfg.Origin = *validatorAddr

	registrarABI, err := abi.JSON(strings.NewReader(kns.FIFSRegistrarABI))
	if err != nil {
		return err
	}

	registerParams, err := registrarABI.Pack(
		"register",
		crypto.Keccak256Hash([]byte(domain)),
		*validatorAddr,
	)
	if err != nil {
		return err
	}

	_, _, err = runtime.Call(bindings.ProxyRegistrarAddr, registerParams, runtimeCfg)
	if err != nil {
		return fmt.Errorf(
			"%s:%s",
			fmt.Sprintf("Failed to register domain %s in FIFSRegistrar.", domain),
			err,
		)
	}

	registryABI, err := abi.JSON(strings.NewReader(kns.KNSRegistryABI))
	if err != nil {
		return err
	}

	setResolverParams, err := registryABI.Pack(
		"setResolver",
		kns2.NameHash(domain+".kowala"),
		bindings.ProxyResolverAddr,
	)
	if err != nil {
		return err
	}

	_, _, err = runtime.Call(bindings.ProxyKNSRegistryAddr, setResolverParams, runtimeCfg)
	if err != nil {
		return fmt.Errorf(
			"%s:%s",
			fmt.Sprintf("Failed to set resolver for domain %s in Registry.", domain),
			err,
		)
	}

	resolverABI, err := abi.JSON(strings.NewReader(kns.PublicResolverABI))
	if err != nil {
		return err
	}

	setAddrParams, err := resolverABI.Pack(
		"setAddr",
		kns2.NameHash(domain+".kowala"),
		contract.address,
	)
	if err != nil {
		return err
	}

	_, _, err = runtime.Call(bindings.ProxyResolverAddr, setAddrParams, runtimeCfg)
	if err != nil {
		return fmt.Errorf(
			"%s:%s",
			fmt.Sprintf("Failed to set domain %s in resolver.", domain),
			err,
		)
	}

	return nil
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
			args.superNodeAmount,
			bindings.ProxyResolverAddr,
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

var ProxiedValidatorManager = &contract{
	name: "Proxied validator manager",
	deploy: func(contract *contract, opts *validGenesisOptions) error {
		args := opts.validatorMgr

		runtimeCfg := contract.runtimeCfg

		proxyContractAddr, code, err := createProxyFromContract(
			common.HexToAddress("0xb5822D5F8D221Ce2dc73e388629eCA256B0Aa4f2"),
			*opts.multiSig.multiSigCreator,
			runtimeCfg,
		)
		if err != nil {
			return err
		}

		contract.address = *proxyContractAddr
		contract.code = code
		runtimeCfg.Origin = bindings.MultiSigWalletAddr
		abi, err := abi.JSON(strings.NewReader(consensus.ValidatorMgrABI))
		if err != nil {
			return err
		}

		initKnsParams, err := abi.Pack(
			"initialize",
			args.baseDeposit,
			args.maxNumValidators,
			args.freezePeriod,
			args.superNodeAmount,
			bindings.ProxyResolverAddr,
		)
		if err != nil {
			return err
		}

		_, _, err = runtime.Call(contract.address, initKnsParams, runtimeCfg)
		if err != nil {
			return fmt.Errorf("%s:%s", "Failed to initialize Proxied Mining Token.", err)
		}

		opts.oracleMgr.validatorMgrAddr = *proxyContractAddr

		return nil
	},
	postDeploy: func(contract *contract, opts *validGenesisOptions) error {
		err := registerValidators(contract, opts)
		if err != nil {
			return err
		}

		return registerAddressToDomain(contract, opts, params.KNSDomains[params.ValidatorMgrDomain].Node())
	},
}
