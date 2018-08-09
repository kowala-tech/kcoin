package main

import (
	testnet2 "github.com/kowala-tech/kcoin/toolbox/testnet"
	"docker.io/go-docker"
	"fmt"
	"os"
)

func main() {
	dockerClient, err := docker.NewEnvClient()
	if err != nil {
		fmt.Printf("error creating docker client. %s\n", err)
		os.Exit(1)
	}

	dockerEngine := testnet2.NewDockerEngine(dockerClient)

	testnet := testnet2.NewTestnet(dockerEngine)

	err = testnet.Start()
	if err != nil {
		fmt.Printf("error creating testnet. %s\n", err)
		os.Exit(1)
	}
}
