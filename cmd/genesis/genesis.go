package main

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"github.com/spf13/viper"
)

var (
	FileConfig string
)

func main() {
	cmd := createCommand()

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "genesis",
		Short: "Generator of a genesis file.",
		Long:  `Generate a genesis.json file based on a config file or parameters.`,
		Run: func(cmd *cobra.Command, args []string) {
			if FileConfig == "" {
				fmt.Println("Params usage.")
			} else {
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
			}

			parsePrefundedAccounts(viper.Get("prefundedAccounts"))

			handler := GenerateGenesisCommandHandler{w:os.Stdout}
			err := handler.Handle(command)
			if err != nil {
				fmt.Printf("Error generating file: %s", err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringVarP(&FileConfig, "config", "c", "", "Use to load configuration from config file.")

	return cmd
}
func parsePrefundedAccounts(accounts interface{}) []PrefundedAccount {
	prefundedAccounts := make([]PrefundedAccount, 0)

	accountArray := accounts.([]interface{})
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
