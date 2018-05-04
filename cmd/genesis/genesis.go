package main

import (
	"fmt"
	"os"
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
				PrefundedAccounts: parsePrefundedAccounts(viper.Get("genesis.prefundedAccounts")),
				Consensus: &genesis.ConsensusOpts{
					Engine:           viper.GetString("consensus.engine"),
					MaxNumValidators: uint64(viper.GetInt64("consensus.maxNumValidators")),
					FreezePeriod:     uint64(viper.GetInt64("consensus.freezePeriod")),
					BaseDeposit:      uint64(viper.GetInt64("consensus.baseDeposit")),
					Validators:       viper.GetStringSlice("consensus.validators"),
				},
				DataFeedSystem: &genesis.DataFeedSystemOpts{
					MaxNumOracles: uint64(viper.GetInt64("datafeed.maxNumOracles")),
					FreezePeriod:  uint64(viper.GetInt64("datafeed.freezePeriod")),
					BaseDeposit:   uint64(viper.GetInt64("datafeed.baseDeposit")),
				},
				Governance: &genesis.GovernanceOpts{
					Origin:           viper.GetString("governance.origin"),
					Governors:        viper.GetStringSlice("governance.governors"),
					NumConfirmations: uint64(viper.GetInt64("governance.numConfirmations")),
				},
				ExtraData: viper.GetString("genesis.extraData"),
			}

			fileName := viper.GetString("genesis.fileName")

			file, err := os.Create(fileName)
			if err != nil {
				fmt.Printf("Error generating file: %s", err)
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
	cmd.Flags().StringP("fileName", "o", "genesis.json", "The output filename (default:genesis.json).")
	viper.BindPFlag("genesis.fileName", cmd.Flags().Lookup("fileName"))

	// governance
	cmd.Flags().StringP("origin", "", "", "The creator's address")
	viper.BindPFlag("governance.origin", cmd.Flags().Lookup("origin"))
	cmd.Flags().StringSliceP("governors", "", []string{}, "Kowala blockchain governors")
	viper.BindPFlag("genesis.governors", cmd.Flags().Lookup("governors"))
	cmd.Flags().Uint64P("numConfirmations", "", 0, "Number of required confirmations to post a transaction")
	viper.BindPFlag("genesis.numConfirmations", cmd.Flags().Lookup("numConfirmations"))

	// system
	cmd.Flags().StringP("network", "n", "", "The network to use, test or main")
	viper.BindPFlag("genesis.network", cmd.Flags().Lookup("network"))
	cmd.Flags().StringP("prefundedAccounts", "a", "", "The prefunded accounts in format 0x212121:12,0x212121:14")
	viper.BindPFlag("prefundedAccounts", cmd.Flags().Lookup("prefundedAccounts"))

	// consensus
	cmd.Flags().StringP("engine", "e", "", "The consensus engine, right now, tendermint is the only available option")
	viper.BindPFlag("consensus.engine", cmd.Flags().Lookup("consensusEngine"))
	cmd.Flags().Uint64P("maxNumValidators", "v", 100, "The maximum number of validators.")
	viper.BindPFlag("consensus.maxNumValidators", cmd.Flags().Lookup("maxNumValidators"))
	cmd.Flags().Uint64P("consensusFreeze", "", 0, "The consensus's deposit freeze period in days.")
	viper.BindPFlag("consensus.freezePeriod", cmd.Flags().Lookup("consensusFreeze"))
	cmd.Flags().Uint64P("consensusBaseDeposit", "", 0, "Base deposit for the consensus")
	viper.BindPFlag("consensus.baseDeposit", cmd.Flags().Lookup("consensusBaseDeposit"))
	cmd.Flags().StringSliceP("validators", "", []string{}, "List of consensus validators")
	viper.BindPFlag("consensus.validators", cmd.Flags().Lookup("validators"))

	// data feed system
	cmd.Flags().Uint64P("maxNumOracles", "o", 0, "The maximum num of oracles.")
	viper.BindPFlag("datafeed.maxNumOracles", cmd.Flags().Lookup("maxNumOracles"))
	cmd.Flags().Uint64P("oracleFreezePeriod", "", 0, "The oracle's deposit freeze period in days.")
	viper.BindPFlag("datafeed.freezePeriod", cmd.Flags().Lookup("oracleUnbondingPeriod"))
	cmd.Flags().StringP("oracleBaseDeposit", "", "", "Base deposit for the oracle activity")
	viper.BindPFlag("datafeed.baseDeposit", cmd.Flags().Lookup("oracleBaseDeposit"))

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

func prefundAccountsFromConfigFile(accounts interface{}) []genesis.PrefundedAccount {
	prefundedAccounts := make([]genesis.PrefundedAccount, 0)

	accountArray := accounts.([]interface{})

	for _, v := range accountArray {
		val := v.(map[string]interface{})

		prefundedAccount := genesis.PrefundedAccount{
			AccountAddress: val["accountAddress"].(string),
			Balance:        val["balance"].(string),
		}

		prefundedAccounts = append(prefundedAccounts, prefundedAccount)
	}

	return prefundedAccounts
}
