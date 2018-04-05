package main

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"github.com/spf13/viper"
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
	cmd.Flags().StringP("maxNumValidators", "v", "", "The network to use, test or main")
	viper.BindPFlag("genesis.maxNumValidators", cmd.Flags().Lookup("maxNumValidators"))
	cmd.Flags().StringP("unbondingPeriod", "p", "", "The network to use, test or main")
	viper.BindPFlag("genesis.unbondingPeriod", cmd.Flags().Lookup("unbondingPeriod"))
	cmd.Flags().StringP("walletAddressGenesisValidator", "g", "", "The network to use, test or main")
	viper.BindPFlag("genesis.walletAddressGenesisValidator", cmd.Flags().Lookup("walletAddressGenesisValidator"))
	cmd.Flags().StringP("consensusEngine", "e", "", "The network to use, test or main")
	viper.BindPFlag("genesis.consensusEngine", cmd.Flags().Lookup("consensusEngine"))
	cmd.Flags().StringP("smartContractsOwner", "s", "", "The network to use, test or main")
	viper.BindPFlag("genesis.smartContractsOwner", cmd.Flags().Lookup("smartContractsOwner"))
	cmd.Flags().StringP("extraData", "d", "", "The network to use, test or main")
	viper.BindPFlag("genesis.extraData", cmd.Flags().Lookup("extraData"))


}

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func parsePrefundedAccounts(accounts interface{}) []PrefundedAccount {
	prefundedAccounts := make([]PrefundedAccount, 0)

	accountArray, ok := accounts.([]interface{})
	if !ok {
		return []PrefundedAccount{}
	}

	for _, v := range accountArray {
		val := v.(map[string]interface{})

		prefundedAccount := PrefundedAccount{
			walletAddress: val["walletAddress"].(string),
			balance: val["balance"].(int64),
		}

		prefundedAccounts = append(prefundedAccounts, prefundedAccount)
	}

	return prefundedAccounts
}
