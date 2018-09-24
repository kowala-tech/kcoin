package testnet

import (
	"math/big"

	"encoding/json"

	"fmt"

	"strconv"
	"strings"
	"time"

	"io/ioutil"
	"math/rand"

	"github.com/kowala-tech/kcoin/client/accounts"
	"github.com/kowala-tech/kcoin/client/accounts/keystore"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/knode/genesis"
)

const defaultNetworkName = "integration"

func init() {
	rand.Seed(time.Now().UnixNano())
}

//Testnet is the interface representing a complete testnet fully operative.
type Testnet interface {
	Start() error
	Stop() error
	IsValidating() bool
	GetNetworkID() string
	GetValidatorID() string
	GetKeyStore() *keystore.KeyStore
	GetGenesisValidatorAccount() accounts.Account
}

type testnet struct {
	dockerEngine DockerEngine

	bootNode         BootNode
	genesisValidator GenesisValidator
	keyStore         *keystore.KeyStore
	chainID          *big.Int
	genesis          []byte

	genesisValidatorAccount accounts.Account
	seederAccount           accounts.Account
	mtokensSeederAccount    accounts.Account

	validating  bool
	networkName string
}

//GetGenesisValidatorAccount returns the genesis validator account used when building the genesis block.
func (t *testnet) GetGenesisValidatorAccount() accounts.Account {
	return t.genesisValidatorAccount
}

//GetKeyStore returns the keystore that the executed testnet has saved all its needed accounts.
func (t *testnet) GetKeyStore() *keystore.KeyStore {
	return t.keyStore
}

//GetValidatorID returns the docker ID of the container of the genesis validator node.
func (t *testnet) GetValidatorID() string {
	return t.genesisValidator.ID()
}

//NewTestnet creates a Testnet with a validator node and a bootnode.
func NewTestnet(dockerEngine DockerEngine) Testnet {
	tmpdir, _ := ioutil.TempDir("", "eth-keystore-test")
	accountsStorage := keystore.NewKeyStore(tmpdir, 2, 1)

	return &testnet{
		dockerEngine: dockerEngine,
		keyStore:     accountsStorage,
		networkName:  defaultNetworkName + strconv.Itoa(rand.Int()),
	}
}

//IsValidating returns if the testnet is validating blocks or not.
func (t *testnet) IsValidating() bool {
	return t.validating
}

//Start prepares the testnet and starts validating blocks.
func (t *testnet) Start() error {
	err := t.createNetwork()
	if err != nil {
		return err
	}

	err = t.createAccounts()
	if err != nil {
		return err
	}

	bootNode, err := NewBootNode(t.dockerEngine, t.networkName)
	if err != nil {
		return err
	}

	err = t.buildGenesis()
	if err != nil {
		return err
	}

	raw, err := t.keyStore.Export(t.genesisValidatorAccount, "test", "test")
	if err != nil {
		return err
	}

	t.bootNode = bootNode
	err = t.bootNode.Start()
	if err != nil {
		return err
	}

	genesisValidator, err := NewGenesisValidator(
		t.dockerEngine,
		t.networkName,
		bootNode.Enode(),
		t.genesis,
		t.genesisValidatorAccount.Address,
		raw,
	)
	if err != nil {
		return err
	}

	t.genesisValidator = genesisValidator
	err = t.genesisValidator.Start()
	if err != nil {
		return err
	}

	err = t.triggerValidation()
	if err != nil {
		return err
	}

	return nil
}

//Stop finishes the testnet and removes all its containers.
func (t *testnet) Stop() error {
	if t.bootNode != nil {
		t.bootNode.Stop()
	}

	if t.genesisValidator != nil {
		t.genesisValidator.Stop()
	}

	err := t.dockerEngine.RemoveNetwork(t.networkName)
	if err != nil {
		return err
	}

	return nil
}

//GetNetworkID returns the docker network that all the elements of the testnet pertains.
func (t *testnet) GetNetworkID() string {
	return t.networkName
}

func (t *testnet) createNetwork() error {
	_, err := t.dockerEngine.CreateNetwork(t.networkName)
	if err != nil {
		return err
	}

	return nil
}

func (t *testnet) createAccounts() error {
	seederAccount, err := t.newAccount()
	if err != nil {
		return err
	}
	t.seederAccount = *seederAccount

	mtokensSeederAccount, err := t.newAccount()
	if err != nil {
		return err
	}
	t.mtokensSeederAccount = *mtokensSeederAccount

	genesisValidatorAccount, err := t.newAccount()
	if err != nil {
		return err
	}
	t.genesisValidatorAccount = *genesisValidatorAccount
	fmt.Printf("The genesis validator address %s\n", t.genesisValidatorAccount.Address.Hex())

	return nil
}

func (t *testnet) newAccount() (*accounts.Account, error) {
	acc, err := t.keyStore.NewAccount("test")
	if err != nil {
		return nil, err
	}
	if err := t.keyStore.Unlock(acc, "test"); err != nil {
		return nil, err
	}
	return &acc, nil
}

func (t *testnet) buildGenesis() error {
	genesisValidatorAddr := t.genesisValidatorAccount.Address.Hex()
	baseDeposit := uint64(1)

	newGenesis, err := genesis.Generate(genesis.Options{
		Network:     "test",
		BlockNumber: 0,
		ExtraData:   "Kowala's first block",
		SystemVars: &genesis.SystemVarsOpts{
			InitialPrice: 1,
		},
		Consensus: &genesis.ConsensusOpts{
			Engine:           genesis.KonsensusConsensus,
			MaxNumValidators: 10,
			FreezePeriod:     5,
			BaseDeposit:      baseDeposit,
			Validators: []genesis.Validator{{
				Address: genesisValidatorAddr,
				Deposit: baseDeposit,
			}},
			MiningToken: &genesis.MiningTokenOpts{
				Name:     "mUSD",
				Symbol:   "mUSD",
				Cap:      100000,
				Decimals: 18,
				Holders: []genesis.TokenHolder{
					{
						Address:   genesisValidatorAddr,
						NumTokens: baseDeposit * 100,
					},
					{
						Address:   t.mtokensSeederAccount.Address.String(),
						NumTokens: baseDeposit * 100,
					},
				},
			},
		},
		Governance: &genesis.GovernanceOpts{
			Origin:           "0xFF9DFBD395cD1C4a4F23C16aa8a5c44109Bc17DF",
			Governors:        []string{"0x259be75d96876f2ada3d202722523e9cd4dd917d"},
			NumConfirmations: 1,
		},
		StabilityContract: &genesis.StabilityContractOpts{
			MinDeposit: 50,
		},
		DataFeedSystem: &genesis.DataFeedSystemOpts{
			MaxNumOracles: 10,
			Price: genesis.PriceOpts{
				SyncFrequency: 600,
				UpdatePeriod:  30,
			},
		},
		PrefundedAccounts: []genesis.PrefundedAccount{
			{
				Address: t.genesisValidatorAccount.Address.Hex(),
				Balance: baseDeposit * 100,
			},
			{
				Address: "0x259be75d96876f2ada3d202722523e9cd4dd917d",
				Balance: baseDeposit * 100,
			},
			{
				Address: t.seederAccount.Address.Hex(),
				Balance: 1000000000000000000,
			},
			{
				Address: t.mtokensSeederAccount.Address.Hex(),
				Balance: baseDeposit * 10000,
			},
		},
	})
	if err != nil {
		return err
	}

	rawJSON, err := json.Marshal(newGenesis)
	if err != nil {
		return err
	}

	t.genesis = rawJSON

	return nil
}

func (t *testnet) triggerValidation() error {
	command := fmt.Sprintf(`
		personal.unlockAccount(eth.coinbase, "test");
		eth.sendTransaction({from:eth.coinbase,to: "%v",value: 1})
	`, t.seederAccount.Address.Hex())
	_, err := t.dockerEngine.Exec(t.genesisValidator.ID(), ConsoleExecCommand(command))
	if err != nil {
		return err
	}

	return common.WaitFor("validation starts", 2*time.Second, 60*time.Second, func() error {
		res, err := t.dockerEngine.Exec(t.genesisValidator.ID(), ConsoleExecCommand("eth.blockNumber"))
		if err != nil {
			return err
		}

		parsed, err := strconv.Atoi(strings.TrimSpace(res.StdOut))
		if err != nil {
			return err
		}

		if parsed <= 0 {
			return fmt.Errorf("can't start validation %q", res.StdOut)
		}

		t.validating = true

		return nil
	})
}

//ConsoleExecCommand wraps a command with ./kcoin attach --exec so it executes
// in the kcoin console.
func ConsoleExecCommand(command string) []string {
	return []string{"./kcoin", "attach", "/root/.kcoin/kusd/kcoin.ipc", "--exec", command}
}
