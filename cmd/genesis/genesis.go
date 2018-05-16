package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kowala-tech/kcoin/kcoin/genesis"
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
		Run: func(cmd *cobra.Command, args []string) {
			loadFromFileConfigIfAvailable()

			options := genesis.Options{
				Network:           viper.GetString("genesis.network"),
				PrefundedAccounts: parsePrefundedAccounts(viper.Get("prefundedAccounts")),
				Consensus: &genesis.ConsensusOpts{
					Engine:           viper.GetString("genesis.consensus.engine"),
					MaxNumValidators: uint64(viper.GetInt64("genesis.consensus.maxNumValidators")),
					FreezePeriod:     uint64(viper.GetInt64("genesis.consensus.freezePeriod")),
					BaseDeposit:      uint64(viper.GetInt64("genesis.consensus.baseDeposit")),
					Validators:       parseValidators(viper.Get("genesis.consensus.validators")),
					MiningToken: &genesis.MiningTokenOpts{
						Name:     viper.GetString("genesis.consensus.token.name"),
						Symbol:   viper.GetString("genesis.consensus.token.symbol"),
						Cap:      uint64(viper.GetInt64("genesis.consensus.token.cap")),
						Decimals: uint64(viper.GetInt64("genesis.consensus.token.decimals")),
						Holders:  parseTokenHolders(viper.Get("genesis.consensus.token.holders")),
					},
				},
				DataFeedSystem: &genesis.DataFeedSystemOpts{
					MaxNumOracles: uint64(viper.GetInt64("genesis.datafeed.maxNumOracles")),
					FreezePeriod:  uint64(viper.GetInt64("genesis.datafeed.freezePeriod")),
					BaseDeposit:   uint64(viper.GetInt64("genesis.datafeed.baseDeposit")),
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
				fmt.Printf("Error during file creation: %s", err)
				os.Exit(1)
			}

			handler := generateGenesisFileCommandHandler{w: file}
			err = handler.handle(options)
			if err != nil {
				fmt.Printf("Error generating file: %s", err)
				os.Exit(1)
			}

			fmt.Println("Genesis file generated.")
		},
	}

	cmd.Flags().StringP("config", "c", "", "Use to load configuration from config file.")
	cmd.Flags().StringP("fileName", "f", "genesis.json", "The output filename (default:genesis.json).")
	viper.BindPFlag("genesis.fileName", cmd.Flags().Lookup("fileName"))

	// governance
	cmd.Flags().StringP("origin", "", "", "The creator's address")
	viper.BindPFlag("genesis.governance.origin", cmd.Flags().Lookup("origin"))
	cmd.Flags().StringSliceP("governors", "", []string{}, "Kowala blockchain governors")
	viper.BindPFlag("genesis.governance.governors", cmd.Flags().Lookup("governors"))
	cmd.Flags().Uint64P("numConfirmations", "", 0, "Number of required confirmations to post a transaction")
	viper.BindPFlag("genesis.governance.numConfirmations", cmd.Flags().Lookup("numConfirmations"))

	// system
	cmd.Flags().StringP("network", "n", "test", "The network to use, test or main")
	viper.BindPFlag("genesis.network", cmd.Flags().Lookup("network"))
	cmd.Flags().StringP("prefundedAccounts", "a", "", "The prefunded accounts in format 0x212121:12,0x212121:14")
	viper.BindPFlag("prefundedAccounts", cmd.Flags().Lookup("prefundedAccounts"))

	// consensus
	cmd.Flags().StringP("engine", "e", "", "The consensus engine, right now, tendermint is the only available option")
	viper.BindPFlag("genesis.consensus.engine", cmd.Flags().Lookup("engine"))
	cmd.Flags().Uint64P("maxNumValidators", "v", 100, "The maximum number of validators.")
	viper.BindPFlag("genesis.consensus.maxNumValidators", cmd.Flags().Lookup("maxNumValidators"))
	cmd.Flags().Uint64P("consensusFreeze", "", 0, "The consensus's deposit freeze period in days.")
	viper.BindPFlag("genesis.consensus.freezePeriod", cmd.Flags().Lookup("consensusFreeze"))
	cmd.Flags().Uint64P("consensusBaseDeposit", "", 0, "Base deposit for the consensus")
	viper.BindPFlag("genesis.consensus.baseDeposit", cmd.Flags().Lookup("consensusBaseDeposit"))
	cmd.Flags().StringP("validators", "", "", "The genesis validators in format (address):(deposit),0x212121:14")
	viper.BindPFlag("genesis.consensus.validators", cmd.Flags().Lookup("validators"))

	// mining token
	cmd.Flags().StringP("token.name", "", "", "The mining token name")
	viper.BindPFlag("genesis.consensus.token.name", cmd.Flags().Lookup("token.name"))
	cmd.Flags().StringP("token.symbol", "", "", "The mining token symbol")
	viper.BindPFlag("genesis.consensus.token.symbol", cmd.Flags().Lookup("token.symbol"))
	cmd.Flags().Uint64P("token.cap", "", 0, "The mining token cap")
	viper.BindPFlag("genesis.consensus.token.cap", cmd.Flags().Lookup("token.cap"))
	cmd.Flags().Uint64P("token.decimals", "", 0, "The mining token decimals")
	viper.BindPFlag("genesis.consensus.token.decimals", cmd.Flags().Lookup("token.decimals"))
	cmd.Flags().StringP("holders", "", "", "The token holders in format (address):(numTokens),0x212121:14")
	viper.BindPFlag("genesis.consensus.token.holders", cmd.Flags().Lookup("holders"))

	// data feed system
	cmd.Flags().Uint64P("maxNumOracles", "o", 0, "The maximum num of oracles.")
	viper.BindPFlag("genesis.datafeed.maxNumOracles", cmd.Flags().Lookup("maxNumOracles"))
	cmd.Flags().Uint64P("oracleFreezePeriod", "", 0, "The oracle's deposit freeze period in days.")
	viper.BindPFlag("genesis.datafeed.freezePeriod", cmd.Flags().Lookup("oracleFreezePeriod"))
	cmd.Flags().StringP("oracleBaseDeposit", "", "", "Base deposit for the oracle activity")
	viper.BindPFlag("genesis.datafeed.baseDeposit", cmd.Flags().Lookup("oracleBaseDeposit"))

	// other
	cmd.Flags().StringP("extraData", "d", "", "Extra data")
	viper.BindPFlag("genesis.extraData", cmd.Flags().Lookup("extraData"))
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

func parsePrefundedAccounts(accounts interface{}) []genesis.PrefundedAccount {
	var prefundedAccounts []genesis.PrefundedAccount

	switch accounts.(type) {
	case []interface{}:
		prefundedAccounts = prefundAccountsFromConfigFile(accounts)
	case string:
		prefundedAccounts = prefundAccountsFromCommandLine(accounts)
	}

	return prefundedAccounts
}

func parseValidators(input interface{}) []genesis.Validator {
	var validators []genesis.Validator

	switch input.(type) {
	case []interface{}:
		validators = validatorsFromConfigFile(input)
	case string:
		validators = validatorsFromCommandLine(input)
	}

	return validators
}

func parseTokenHolders(input interface{}) []genesis.TokenHolder {
	var holders []genesis.TokenHolder

	switch input.(type) {
	case []interface{}:
		holders = tokenHoldersFromConfigFile(input)
	case string:
		holders = tokenHoldersFromCommandLine(input)
	}

	return holders
}

func prefundAccountsFromCommandLine(accounts interface{}) []genesis.PrefundedAccount {
	prefundedAccounts := make([]genesis.PrefundedAccount, 0)

	accountsString := accounts.(string)
	if accountsString == "" {
		return nil
	}

	a := strings.Split(accountsString, ",")
	for _, v := range a {
		values := strings.Split(v, ":")
		balance := values[1]

		prefundedAccount := genesis.PrefundedAccount{
			AccountAddress: values[0],
			Balance:        balance,
		}

		prefundedAccounts = append(prefundedAccounts, prefundedAccount)
	}

	return prefundedAccounts
}

func validatorsFromCommandLine(input interface{}) []genesis.Validator {
	validators := make([]genesis.Validator, 0)

	validatorsStr := input.(string)
	if validatorsStr == "" {
		return nil
	}

	values := strings.Split(validatorsStr, ",")
	for _, value := range values {
		parts := strings.Split(value, ":")

		deposit, err := strconv.ParseUint(parts[1], 10, 64)
		if err != nil {
			panic(fmt.Errorf("Fatal error parsing deposit: %s \n", err))
		}
		validator := genesis.Validator{
			Address: parts[0],
			Deposit: deposit,
		}

		validators = append(validators, validator)
	}

	return validators
}

func tokenHoldersFromCommandLine(input interface{}) []genesis.TokenHolder {
	holders := make([]genesis.TokenHolder, 0)

	holdersStr := input.(string)
	if holdersStr == "" {
		return nil
	}

	values := strings.Split(holdersStr, ",")
	for _, value := range values {
		parts := strings.Split(value, ":")

		numTokens, err := strconv.ParseUint(parts[1], 10, 64)
		if err != nil {
			panic(fmt.Errorf("Fatal error parsing numTokens: %s \n", err))
		}
		holders = append(holders, genesis.TokenHolder{
			Address:   parts[0],
			NumTokens: numTokens,
		})
	}

	return holders
}

func prefundAccountsFromConfigFile(accounts interface{}) []genesis.PrefundedAccount {
	prefundedAccounts := make([]genesis.PrefundedAccount, 0)

	accountArray := accounts.([]interface{})
	for _, v := range accountArray {
		val := v.(map[string]interface{})
		prefundedAccounts = append(prefundedAccounts, genesis.PrefundedAccount{
			AccountAddress: val["accountAddress"].(string),
			Balance:        val["balance"].(string),
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
			Deposit: parts["deposit"].(uint64),
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
			NumTokens: parts["numTokens"].(uint64),
		})
	}

	return holders
}
