package genesis

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/kowala-tech/kcoin/client/accounts/abi"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/consensus"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/kns"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/oracle"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/ownership"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/proxy"
	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/core/vm/runtime"
	"github.com/kowala-tech/kcoin/client/crypto"
	kns2 "github.com/kowala-tech/kcoin/client/kns"
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

		proxyContractAddr, code, err := createProxyFromContract(
			common.HexToAddress("0x75AD571eFAcC241B23099c724c4A71FE30659145"),
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
			common.HexToAddress(consensus.MultiSigWalletAddr),
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
		runtimeCfg.Origin = common.HexToAddress(consensus.MultiSigWalletAddr)

		abi, err := abi.JSON(strings.NewReader(kns.KNSRegistryABI))
		if err != nil {
			return err
		}

		subnodeOwnerParams, err := abi.Pack(
			"setSubnodeOwner",
			[32]byte{},
			crypto.Keccak256Hash([]byte("kowala")),
			common.HexToAddress(consensus.ProxyRegistrarAddr),
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

		proxyContractAddr, code, err := createProxyFromContract(
			common.HexToAddress("0x2A4443ec27BF5F849B2Da15eB697d3Ef5302f186"),
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
			common.HexToAddress(consensus.ProxyKNSRegistryAddr),
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

		proxyContractAddr, code, err := createProxyFromContract(
			common.HexToAddress("0xA08fCF0425736Ae7c10c0D3FCfB45C65f5f78212"),
			*args.multiSigCreator,
			runtimeCfg,
		)
		if err != nil {
			return err
		}

		contract.address = *proxyContractAddr
		contract.code = code

		// Init
		validatorAddr := opts.prefundedAccounts[0].accountAddress

		runtimeCfg.Origin = *validatorAddr

		abi, err := abi.JSON(strings.NewReader(kns.PublicResolverABI))
		if err != nil {
			return err
		}

		initKnsParams, err := abi.Pack(
			"initialize",
			common.HexToAddress(consensus.ProxyKNSRegistryAddr),
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

	ret, _, err := runtime.Call(common.HexToAddress(consensus.ProxyFactoryAddr), createProxyArgs, runtimeCfg)
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

		return nil
	},
	postDeploy: func(contract *contract, opts *validGenesisOptions) error {
		return registerAddressToDomain(contract, opts, "multisig")
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
	postDeploy: func(contract *contract, opts *validGenesisOptions) error {
		err := mintTokens(contract, opts)
		if err != nil {
			return err
		}

		registerAddressToDomain(contract, opts, "miningtoken")

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
	postDeploy: func(contract *contract, opts *validGenesisOptions) error {
		err := registerAddressToDomain(contract, opts, "oraclemgr")
		if err != nil {
			return err
		}

		return nil
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

	_, _, err = runtime.Call(common.HexToAddress(consensus.ProxyRegistrarAddr), registerParams, runtimeCfg)
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
		kns2.NameHash(domain + ".kowala"),
		common.HexToAddress(consensus.ProxyResolverAddr),
	)
	if err != nil {
		return err
	}

	_, _, err = runtime.Call(common.HexToAddress(consensus.ProxyKNSRegistryAddr), setResolverParams, runtimeCfg)
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
		kns2.NameHash(domain + ".kowala"),
		contract.address,
	)
	if err != nil {
		return err
	}

	_, _, err = runtime.Call(common.HexToAddress(consensus.ProxyResolverAddr), setAddrParams, runtimeCfg)
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
	postDeploy: func(contract *contract, opts *validGenesisOptions) error {
		err := registerValidators(contract, opts)
		if err != nil {
			return err
		}

		return registerAddressToDomain(contract, opts, "validatormgr")
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
