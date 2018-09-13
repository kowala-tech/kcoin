package main

import (
	"fmt"

	"github.com/kowala-tech/kcoin/client/contracts/bindings"
	"gopkg.in/urfave/cli.v1"
)

var (
	showAddressesCommand = cli.Command{
		Action:      showAddresses,
		Name:        "addresses",
		Usage:       "Show important addresses from the network",
		Description: "Write important network addresses, like the MultiSig and KNS related ones, to stdout",
	}
)

func showAddresses(ctx *cli.Context) error {
	fmt.Printf("MultiSig Wallet Address: %s\n", bindings.MultiSigWalletAddr.Hex())
	fmt.Printf("\n")
	fmt.Printf("Proxy Factory Address: %s\n", bindings.ProxyFactoryAddr.Hex())
	fmt.Printf("\n")
	fmt.Printf("Kns Registry Address: %s\n", bindings.ProxyKNSRegistryAddr.Hex())
	fmt.Printf("Kns Registrar Address: %s\n", bindings.ProxyRegistrarAddr.Hex())
	fmt.Printf("Kns Resolver Address: %s\n", bindings.ProxyResolverAddr.Hex())
	return nil
}
