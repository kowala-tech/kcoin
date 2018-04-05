package main

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"github.com/spf13/viper"
	"strings"
	"strconv"
)

var (
	FileConfig string
	cmd *cobra.Command
)

func init() {
	cmd = &cobra.Command{
		Use:   "genesis",
		Short: "Generator of a genesis file.",
		Long:  `Generate a genesis.json file based on a config file or parameters.`,
		Run: func(cmd *cobra.Command, args []string) {
			if FileConfig != "" {
				viper.SetConfigFile(FileConfig)

				err := viper.ReadInConfig()
				if err != nil {
					panic(fmt.Errorf("Fatal error config file: %s \n", err))
				}
			}

			command := GenerateGenesisCommand{
				network: viper.GetString("genesis.network"),
				maxNumValidators: viper.GetString("genesis.maxNumValidators"),
				unbondingPeriod: viper.GetString("genesis.unbondingPeriod"),
				walletAddressGenesisValidator: viper.GetString("genesis.walletAddressGenesisValidator"),
				prefundedAccounts: parsePrefundedAccounts(viper.Get("prefundedAccounts")),
				consensusEngine: viper.GetString("genesis.consensusEngine"),
				smartContractsOwner: viper.GetString("genesis.smartContractsOwner"),
				extraData: viper.GetString("genesis.extraData"),
			}

			handler := GenerateGenesisCommandHandler{w:os.Stdout}
			err := handler.Handle(command)
			if err != nil {
				fmt.Printf("Error generating file: %s", err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringVarP(&FileConfig, "config", "c", "", "Use to load configuration from config file.")
	cmd.Flags().StringP("network", "n", "", "The network to use, test or main")
	viper.BindPFlag("genesis.network", cmd.Flags().Lookup("network"))
	cmd.Flags().StringP("maxNumValidators", "v", "", "The maximum num of validators.")
	viper.BindPFlag("genesis.maxNumValidators", cmd.Flags().Lookup("maxNumValidators"))
	cmd.Flags().StringP("unbondingPeriod", "p", "", "The unbonding period in days.")
	viper.BindPFlag("genesis.unbondingPeriod", cmd.Flags().Lookup("unbondingPeriod"))
	cmd.Flags().StringP("walletAddressGenesisValidator", "g", "", "The wallet address of the genesis validator.")
	viper.BindPFlag("genesis.walletAddressGenesisValidator", cmd.Flags().Lookup("walletAddressGenesisValidator"))
	cmd.Flags().StringP("consensusEngine", "e", "", "The consensus engine, right now only supports tendermint")
	viper.BindPFlag("genesis.consensusEngine", cmd.Flags().Lookup("consensusEngine"))
	cmd.Flags().StringP("smartContractsOwner", "s", "", "The address of the smart contracts owner.")
	viper.BindPFlag("genesis.smartContractsOwner", cmd.Flags().Lookup("smartContractsOwner"))
	cmd.Flags().StringP("extraData", "d", "", "Extra data")
	viper.BindPFlag("genesis.extraData", cmd.Flags().Lookup("extraData"))
	cmd.Flags().StringP("prefundedAccounts", "a", "", "The prefunded accounts in format 0x212121:12,0x212121:14")
	viper.BindPFlag("prefundedAccounts", cmd.Flags().Lookup("prefundedAccounts"))
}

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func parsePrefundedAccounts(accounts interface{}) []PrefundedAccount {
	prefundedAccounts := make([]PrefundedAccount, 0)

	switch accounts.(type) {
	case []interface{}:
		accountArray := accounts.([]interface{})
		for _, v := range accountArray {
			val := v.(map[string]interface{})

			prefundedAccount := PrefundedAccount{
				walletAddress: val["walletAddress"].(string),
				balance: val["balance"].(int64),
			}

			prefundedAccounts = append(prefundedAccounts, prefundedAccount)
		}
	case string:
		accountsString := accounts.(string)
		a := strings.Split(accountsString, ",")

		for _, v := range a {
			values := strings.Split(v, ":")
			balance, err := strconv.Atoi(values[1])
			if err != nil {
				balance = 0
			}

			prefundedAccount := PrefundedAccount{
				walletAddress: values[0],
				balance: int64(balance),
			}

			prefundedAccounts = append(prefundedAccounts, prefundedAccount)
		}

		fmt.Printf("%v", accounts)
	}

	return prefundedAccounts
}
