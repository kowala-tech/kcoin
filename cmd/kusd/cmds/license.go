package cmd

import "github.com/spf13/cobra"

var LicenseCmd = &cobra.Command{
	Short: "Display license information",
	Run:   printLicense,
	Args:  cobra.NoArgs,
}

// @TODO(rgeraldes)
// printLicense
func printLicense(cmd *cobra.Command, args []string) {}
