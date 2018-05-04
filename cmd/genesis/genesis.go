package main

import (
	"fmt"
	"github.com/kowala-tech/kcoin/kcoin/genesis"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
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
				Network:                        viper.GetString("genesis.network"),
				ConsensusEngine:                viper.GetString("genesis.consensusEngine"),
				PrefundedAccounts:              parsePrefundedAccounts(viper.Get("prefundedAccounts")),				
				Consensus:						&genesis.ConsensusOpts{
					MaxNumValidators:               viper.GetString("consensus.maxNumValidators"),
					UnbondingPeriod:                viper.GetString("consensus.UnbondingPeriod"),
					BaseDeposit:					viper.GetString("consensus.BaseDeposit"),
					Validators: []string{viper.GetString("consensus.Validators")},
				},
				DataFeedSystem: &genesis.DataFeedSystemOpts{
					MaxNumOracles:		viper.GetString("datafeed.MaxNumOracles")
					UnbondingPeriod:                viper.GetString("datafeed.oracleUnbondingPeriod"),
					BaseDeposit:					viper.GetString("datafeed.oracleBaseDeposit"),
				},
				Governance: &genesis.MultiSigOpts{
					Origin:            viper.GetString("governance.origin"),
					Governors:					[]string{},
				},
				ExtraData:                      viper.GetString("genesis.extraData"),
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
	
	// system
	cmd.Flags().StringP("network", "n", "", "The network to use, test or main")
	viper.BindPFlag("genesis.network", cmd.Flags().Lookup("network"))
	cmd.Flags().StringP("consensusEngine", "e", "", "The consensus engine, right now only supports tendermint")
	viper.BindPFlag("genesis.consensusEngine", cmd.Flags().Lookup("consensusEngine"))
	cmd.Flags().StringP("prefundedAccounts", "a", "", "The prefunded accounts in format 0x212121:12,0x212121:14")
	viper.BindPFlag("prefundedAccounts", cmd.Flags().Lookup("prefundedAccounts"))	
	
	// validator mgr
	cmd.Flags().StringP("maxNumValidators", "v", "", "The maximum num of validators.")
	viper.BindPFlag("genesis.maxNumValidators", cmd.Flags().Lookup("maxNumValidators"))
	cmd.Flags().StringP("consensusUnbondingPeriod", "p", "", "The consensus unbounding period in days.")
	viper.BindPFlag("genesis.consensusUnbondingPeriod", cmd.Flags().Lookup("consensusUnbondingPeriod"))
	cmd.Flags().StringP("consensusBaseDeposit", "", "", "Base deposit for the consensus")
	viper.BindPFlag("genesis.consensusBaseDeposit", cmd.Flags().Lookup("consensusBaseDeposit"))
	cmd.Flags().StringP("accountAddressGenesisValidator", "g", "", "The wallet address of the genesis validator.")
	viper.BindPFlag("genesis.accountAddressGenesisValidator", cmd.Flags().Lookup("accountAddressGenesisValidator"))

	// oracle mgr
	cmd.Flags().StringP("maxNumOracles", "o", "", "The maximum num of oracles.")
	viper.BindPFlag("genesis.maxNumOracles", cmd.Flags().Lookup("maxNumOracles"))
	cmd.Flags().StringP("oracleUnbondingPeriod", "", "", "The oracle unbounding period in days.")
	viper.BindPFlag("genesis.oracleUnbondingPeriod", cmd.Flags().Lookup("oracleUnbondingPeriod"))
	cmd.Flags().StringP("oracleBaseDeposit", "", "", "Base deposit for the oracle activity")
	viper.BindPFlag("genesis.oracleBaseDeposit", cmd.Flags().Lookup("oracleBaseDeposit"))

	// multi sig
	cmd.Flags().StringP("multiSigOnwers", "s", "", "The addresses of core multi signature wallet owners.")
	viper.BindPFlag("genesis.multiSigOwners", cmd.Flags().Lookup("multiSigOwners"))
	cmd.Flags().StringP("multiSigCreator", "s", "", "The address of core multi signature wallet creator.")
	viper.BindPFlag("genesis.multiSigOwners", cmd.Flags().Lookup("multiSigOwners"))

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
