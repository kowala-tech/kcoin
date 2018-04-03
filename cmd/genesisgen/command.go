package main

import (
	"bytes"
	"github.com/kowala-tech/kcoin/accounts/abi"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/contracts/network/contracts"
	"github.com/kowala-tech/kcoin/core"
	"github.com/kowala-tech/kcoin/core/vm"
	"github.com/kowala-tech/kcoin/core/vm/runtime"
	"github.com/kowala-tech/kcoin/params"
	"github.com/pkg/errors"
	"io"
	"math/big"
	"math/rand"
	"strings"
	"time"
	"encoding/json"
)

var (
	ErrEmptyMaxNumValidators                        = errors.New("max number of validators is mandatory")
	ErrEmptyUnbondingPeriod                         = errors.New("unbonding period in days is mandatory")
	ErrEmptyWalletAddressValidator                  = errors.New("Wallet address of genesis validator is mandatory")
	ErrInvalidWalletAddressValidator                = errors.New("Wallet address of genesis validator is invalid")
	ErrEmptyPrefundedAccounts                       = errors.New("empty prefunded accounts, at least the validator wallet address should be included")
	ErrWalletAddressValidatorNotInPrefundedAccounts = errors.New("prefunded accounts should include genesis validator account")
	ErrInvalidAddressInPrefundedAccounts            = errors.New("address in prefunded accounts is invalid")
	ErrInvalidContractsOwnerAddress                 = errors.New("address used for smart contracts is invalid")
)

type GenerateGenesisCommand struct {
	network                       string
	maxNumValidators              string
	unbondingPeriod               string
	walletAddressGenesisValidator string
	prefundedAccounts             []PrefundedAccount
}

type PrefundedAccount struct {
	walletAddress string
	balance       int64
}

type validPrefundedAccount struct {
	walletAddress *common.Address
	balance       *big.Int
}

type GenerateGenesisCommandHandler struct {
	w io.Writer
}

func (h *GenerateGenesisCommandHandler) Handle(command GenerateGenesisCommand) error {
	network, err := NewNetwork(command.network)
	if err != nil {
		return err
	}

	maxNumValidators, err := h.createMaxNumValidators(command.maxNumValidators)
	if err != nil {
		return err
	}

	unbondingPeriod, err := h.createUnbondingPeriod(command.unbondingPeriod)
	if err != nil {
		return err
	}

	walletAddressValidator, err := h.createWalletAddress(command.walletAddressGenesisValidator)
	if err != nil {
		return err
	}

	validPrefundedAccounts, err := h.createPrefundedAccounts(command.prefundedAccounts)
	if err != nil {
		return err
	}

	if !h.prefundedIncludesValidatorWallet(validPrefundedAccounts, walletAddressValidator) {
		return ErrWalletAddressValidatorNotInPrefundedAccounts
	}

	genesis := &core.Genesis{
		Timestamp: uint64(time.Now().Unix()),
		GasLimit:  4700000,
		Alloc:     make(core.GenesisAlloc),
		Config:    &params.ChainConfig{},
	}

	switch network {
	case MainNetwork:
		genesis.Config.ChainID = params.MainnetChainConfig.ChainID
	case TestNetwork:
		genesis.Config.ChainID = params.TestnetChainConfig.ChainID
	case OtherNetwork:
		genesis.Config.ChainID = new(big.Int).SetUint64(uint64(rand.Intn(65536)))
	}

	//TODO: Add possibility to change it (optional)
	genesis.Config.Tendermint = &params.TendermintConfig{Rewarded: true}
	genesis.ExtraData = make([]byte, 32)

	//TODO: Default account for the network contracts, maybe it is good to unify in
	//a global constant this information.
	owner, err := h.createWalletAddress("0x259be75d96876f2ada3d202722523e9cd4dd917d")
	if err != nil {
		return ErrInvalidContractsOwnerAddress
	}

	genesis.Alloc[*owner] = core.GenesisAccount{Balance: new(big.Int).Mul(common.Big1, big.NewInt(params.Ether))}

	//TODO: This maybe will be need to be available to change by the parameters in the command.
	baseDeposit := common.Big0

	electionABI, err := abi.JSON(strings.NewReader(contracts.ElectionContractABI))
	if err != nil {
		return err
	}

	electionParams, err := electionABI.Pack(
		"",
		baseDeposit,
		maxNumValidators,
		unbondingPeriod,
		*walletAddressValidator,
	)
	if err != nil {
		return err
	}

	runtimeCfg := &runtime.Config{
		Origin: *owner,
		EVMConfig: vm.Config{
			Debug:  true,
			Tracer: newVmTracer(),
		},
	}

	contract, err := createContract(runtimeCfg, append(common.FromHex(contracts.ElectionContractBin), electionParams...))
	if err != nil {
		return err
	}

	genesis.Alloc[contract.addr] = core.GenesisAccount{
		Code:    contract.code,
		Storage: contract.storage,
		Balance: new(big.Int).Mul(baseDeposit, new(big.Int).SetUint64(params.Ether)),
	}

	for _, vAccount := range validPrefundedAccounts {
		genesis.Alloc[*vAccount.walletAddress] = core.GenesisAccount{
			Balance: vAccount.balance,
		}
	}

	// Add a batch of precompile balances to avoid them getting deleted
	for i := int64(0); i < 256; i++ {
		genesis.Alloc[common.BigToAddress(big.NewInt(i))] = core.GenesisAccount{Balance: big.NewInt(1)}
	}

	//TODO: This should be an optional param.
	extra := "Extradata"
	if len(extra) > 32 {
		extra = extra[:32]
	}

	genesis.ExtraData = append([]byte(extra), genesis.ExtraData[len(extra):]...)

	out, _ := json.MarshalIndent(genesis, "", "  ")

	_, err = h.w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (h *GenerateGenesisCommandHandler) createMaxNumValidators(s string) (*big.Int, error) {
	if s = strings.TrimSpace(s); s == "" {
		return nil, ErrEmptyMaxNumValidators
	}

	numValidators, ok := new(big.Int).SetString(s, 0)
	if !ok {
		//TODO: Create error
		return nil, errors.New("invalid max num of validators.")
	}

	return numValidators, nil
}

func (h *GenerateGenesisCommandHandler) createUnbondingPeriod(uP string) (*big.Int, error) {
	var text string
	if text = strings.TrimSpace(uP); text == "" {
		return nil, ErrEmptyUnbondingPeriod
	}

	unbondingPeriod, ok := new(big.Int).SetString(text, 0)
	if !ok {
		//TODO: Create error
		return nil, errors.New("invalid max num of validators.")
	}

	return unbondingPeriod, nil
}

func (h *GenerateGenesisCommandHandler) createWalletAddress(wA string) (*common.Address, error) {
	stringAddr := wA

	if text := strings.TrimSpace(wA); text == "" {
		return nil, ErrEmptyWalletAddressValidator
	}

	if strings.HasPrefix(stringAddr, "0x") {
		stringAddr = strings.TrimPrefix(stringAddr, "0x")
	}

	if len(stringAddr) != 40 {
		return nil, ErrInvalidWalletAddressValidator
	}

	bigaddr, _ := new(big.Int).SetString(stringAddr, 16)
	address := common.BigToAddress(bigaddr)

	return &address, nil
}

func (h *GenerateGenesisCommandHandler) createPrefundedAccounts(accounts []PrefundedAccount) ([]*validPrefundedAccount, error) {
	var validAccounts []*validPrefundedAccount

	if len(accounts) == 0 {
		return nil, ErrEmptyPrefundedAccounts
	}

	for _, a := range accounts {
		address, err := h.createWalletAddress(a.walletAddress)
		if err != nil {
			return nil, ErrInvalidAddressInPrefundedAccounts
		}

		balance := big.NewInt(a.balance)

		validAccount := &validPrefundedAccount{
			walletAddress: address,
			balance:       balance,
		}

		validAccounts = append(validAccounts, validAccount)
	}

	return validAccounts, nil
}

func (h *GenerateGenesisCommandHandler) prefundedIncludesValidatorWallet(
	accounts []*validPrefundedAccount,
	addresses *common.Address,
) bool {
	for _, account := range accounts {
		if bytes.Equal(account.walletAddress.Bytes(), addresses.Bytes()) {
			return true
		}
	}

	return false
}

func createContract(cfg *runtime.Config, code []byte) (*contractData, error) {
	out, addr, _, err := runtime.Create(code, cfg)
	if err != nil {
		return nil, err
	}
	return &contractData{
		addr:    addr,
		code:    out,
		storage: cfg.EVMConfig.Tracer.(*vmTracer).data[addr],
	}, nil
}

type contractData struct {
	addr    common.Address
	code    []byte
	storage map[common.Hash]common.Hash
}
