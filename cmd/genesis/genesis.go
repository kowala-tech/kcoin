package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kowala-tech/kcoin/knode/genesis"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cmd *cobra.Command
)

func init() {
	cmd = &cobra.Command{
		Use:   "genesis",
		Short: "Generator of a genesis file.",
		Long:  `Generate a genesis.json file based on a config file or parameters.`,
		RunE:  createGenesis,
	}

	cmd.Flags().StringP("config", "c", "", "Use to load configuration from config file.")
}

func loadFromFileConfigIfAvailable() {
	fileConfig, _ := cmd.Flags().GetString("config")
	if fileConfig != "" {
		viper.SetConfigFile(fileConfig)

		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}
}

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func parsePrefundedAccounts(accounts interface{}) ([]genesis.PrefundedAccount, error) {
	var err error
	var prefundedAccounts []genesis.PrefundedAccount

	switch accounts.(type) {
	case []interface{}:
		prefundedAccounts = prefundAccountsFromConfigFile(accounts)
	case string:
		prefundedAccounts, err = prefundAccountsFromCommandLine(accounts)
		if err != nil {
			return nil, err
		}
	}

	return prefundedAccounts, nil
}

func parseValidators(input interface{}) ([]genesis.Validator, error) {
	var err error
	var validators []genesis.Validator

	switch input.(type) {
	case []interface{}:
		validators = validatorsFromConfigFile(input)
	case string:
		validators, err = validatorsFromCommandLine(input)
		if err != nil {
			return nil, err
		}
	}

	return validators, nil
}

func parseTokenHolders(input interface{}) ([]genesis.TokenHolder, error) {
	var err error
	var holders []genesis.TokenHolder

	switch input.(type) {
	case []interface{}:
		holders = tokenHoldersFromConfigFile(input)
	case string:
		holders, err = tokenHoldersFromCommandLine(input)
		if err != nil {
			return nil, err
		}
	}

	return holders, nil
}

func prefundAccountsFromCommandLine(accounts interface{}) ([]genesis.PrefundedAccount, error) {
	prefundedAccounts := make([]genesis.PrefundedAccount, 0)

	accountsString := accounts.(string)
	if accountsString == "" {
		return nil, nil
	}

	a := strings.Split(accountsString, ",")
	for _, v := range a {
		values := strings.Split(v, ":")
		balance, err := strconv.ParseUint(values[1], 10, 64)
		if err != nil {
			return nil, err
		}

		prefundedAccount := genesis.PrefundedAccount{
			Address: values[0],
			Balance: balance,
		}

		prefundedAccounts = append(prefundedAccounts, prefundedAccount)
	}

	return prefundedAccounts, nil
}

func validatorsFromCommandLine(input interface{}) ([]genesis.Validator, error) {
	validators := make([]genesis.Validator, 0)

	validatorsStr := input.(string)
	if validatorsStr == "" {
		return nil, nil
	}

	values := strings.Split(validatorsStr, ",")
	for _, value := range values {
		parts := strings.Split(value, ":")

		deposit, err := strconv.ParseUint(parts[1], 10, 64)
		if err != nil {
			return nil, err
		}
		validator := genesis.Validator{
			Address: parts[0],
			Deposit: deposit,
		}

		validators = append(validators, validator)
	}

	return validators, nil
}

func tokenHoldersFromCommandLine(input interface{}) ([]genesis.TokenHolder, error) {
	holders := make([]genesis.TokenHolder, 0)

	holdersStr := input.(string)
	if holdersStr == "" {
		return nil, nil
	}

	values := strings.Split(holdersStr, ",")
	for _, value := range values {
		parts := strings.Split(value, ":")

		numTokens, err := strconv.ParseUint(parts[1], 10, 64)
		if err != nil {
			return nil, err
		}
		holders = append(holders, genesis.TokenHolder{
			Address:   parts[0],
			NumTokens: numTokens,
		})
	}

	return holders, nil
}

func prefundAccountsFromConfigFile(accounts interface{}) []genesis.PrefundedAccount {
	prefundedAccounts := make([]genesis.PrefundedAccount, 0)

	accountArray := accounts.([]interface{})
	for _, v := range accountArray {
		val := v.(map[string]interface{})
		prefundedAccounts = append(prefundedAccounts, genesis.PrefundedAccount{
			Address: val["accountAddress"].(string),
			Balance: uint64(val["balance"].(int64)),
		})
	}

	return prefundedAccounts
}

func validatorsFromConfigFile(input interface{}) []genesis.Validator {
	validators := make([]genesis.Validator, 0)

	validatorArray := input.([]interface{})
	for _, value := range validatorArray {
		parts := value.(map[string]interface{})
		validators = append(validators, genesis.Validator{
			Address: parts["address"].(string),
			Deposit: uint64(parts["deposit"].(int64)),
		})
	}

	return validators
}

func tokenHoldersFromConfigFile(input interface{}) []genesis.TokenHolder {
	holders := make([]genesis.TokenHolder, 0)

	holdersArray := input.([]interface{})
	for _, value := range holdersArray {
		parts := value.(map[string]interface{})
		holders = append(holders, genesis.TokenHolder{
			Address:   parts["address"].(string),
			NumTokens: uint64(parts["numTokens"].(int64)),
		})
	}

	return holders
}

func createGenesis(cmd *cobra.Command, args []string) error {
	loadFromFileConfigIfAvailable()

	prefundedAccounts, err := parsePrefundedAccounts(viper.Get("prefundedAccounts"))
	if err != nil {
		return err
	}

	validators, err := parseValidators(viper.Get("genesis.consensus.validators"))
	if err != nil {
		return err
	}

	tokenHolders, err := parseTokenHolders(viper.Get("genesis.consensus.token.holders"))
	if err != nil {
		return err
	}

	options := genesis.Options{
		Network:           viper.GetString("genesis.network"),
		PrefundedAccounts: prefundedAccounts,
		Consensus: &genesis.ConsensusOpts{
			Engine:           viper.GetString("genesis.consensus.engine"),
			MaxNumValidators: uint64(viper.GetInt64("genesis.consensus.maxNumValidators")),
			FreezePeriod:     uint64(viper.GetInt64("genesis.consensus.freezePeriod")),
			BaseDeposit:      uint64(viper.GetInt64("genesis.consensus.baseDeposit")),
			Validators:       validators,
			MiningToken: &genesis.MiningTokenOpts{
				Name:     viper.GetString("genesis.consensus.token.name"),
				Symbol:   viper.GetString("genesis.consensus.token.symbol"),
				Cap:      uint64(viper.GetInt64("genesis.consensus.token.cap")),
				Decimals: uint64(viper.GetInt64("genesis.consensus.token.decimals")),
				Holders:  tokenHolders,
			},
		},
		DataFeedSystem: &genesis.DataFeedSystemOpts{
			MaxNumOracles: uint64(viper.GetInt64("genesis.datafeed.maxNumOracles")),
			FreezePeriod:  uint64(viper.GetInt64("genesis.datafeed.freezePeriod")),
			BaseDeposit:   uint64(viper.GetInt64("genesis.datafeed.baseDeposit")),
			Price: genesis.PriceOpts{
				InitialPrice:  viper.GetFloat64("genesis.datafeed.price.initialPrice"),
				SyncFrequency: uint64(viper.GetInt64("genesis.datafeed.price.syncFrequency")),
				UpdatePeriod:  uint64(viper.GetInt64("genesis.datafeed.price.updatePeriod")),
			},
		},
		Governance: &genesis.GovernanceOpts{
			Origin:           viper.GetString("genesis.governance.origin"),
			Governors:        viper.GetStringSlice("genesis.governance.governors"),
			NumConfirmations: uint64(viper.GetInt64("genesis.governance.numConfirmations")),
		},
		ExtraData: viper.GetString("genesis.extraData"),
	}

	fileName := viper.GetString("genesis.fileName")

	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("Error during file creation: %s", err)
	}

	handler := generateGenesisFileCommandHandler{w: file}
	err = handler.handle(options)
	if err != nil {
		return fmt.Errorf("Error generating file: %s", err)
	}

	fmt.Println("Genesis file generated.")

	return nil
}
