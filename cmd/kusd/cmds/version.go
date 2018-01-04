package cmd

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/kowala-tech/kUSD/params"
	"github.com/spf13/cobra"
)

// VersionCmd prints version numbers
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version numbers",
	Run:   version,
	Args:  cobra.NoArgs,
}

func version(cmd *cobra.Command, args []String) {
	fmt.Println(strings.Title(cmd.Parent().Name())
	fmt.Println("Version:", params.Version)
	if gitCommit != "" {
		fmt.Println("Git Commit:", gitCommit)
	}
	fmt.Println("Architecture:", runtime.GOARCH)
	fmt.Println("Consensus Protocol Versions:", consensus.ProtocolVersions)
	fmt.Println("Network Id:", consensus.DefaultConfig.NetworkId)
	fmt.Println("Go Version:", runtime.Version())
	fmt.Println("Operating System:", runtime.GOOS)
	fmt.Printf("GOPATH=%s\n", os.Getenv("GOPATH"))
	fmt.Printf("GOROOT=%s\n", runtime.GOROOT())
}
