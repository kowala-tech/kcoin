package main

import (
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/kowala-tech/kUSD/cmd/kusd/cmds"
)

func main() {
	KowalaCmd := cmd.KowalaCmd
	KowalaCmd.AddCommand(
		cmd.ConfigCmd,
		/*
			cmd.InitCmd,
			cmd.AccountCmd,
			cmd.KeyCmd,
			cmd.ValidatorCmd,
			// Misc
			cmd.VersionCmd,
			cmd.LicenseCmd,
		*/
	)

	if err := cmd.Execute(); err != nil {
		utils.Fatalf(err)
	}
}
