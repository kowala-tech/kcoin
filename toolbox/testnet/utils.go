package testnet

import (
	"context"
	"encoding/hex"

	"math/rand"
	"time"

	"docker.io/go-docker"
)

func createDockerClient() (*docker.Client, error) {
	dockerClient, err := docker.NewEnvClient()
	if err != nil {
		return nil, err
	}

	return dockerClient, nil
}

type testContext struct {
	context      context.Context
	dockerEngine DockerEngine
	networkID    string
	nodes        []Node
}

func cleanContext(c testContext) error {
	for _, node := range c.nodes {
		err := node.Stop()
		if err != nil {
			return err
		}
	}

	err := c.dockerEngine.RemoveNetwork(c.networkID)
	if err != nil {
		return err
	}

	return nil
}

var src = rand.New(rand.NewSource(time.Now().UnixNano()))

func randStringBytes(n int) string {
	b := make([]byte, n/2)

	if _, err := src.Read(b); err != nil {
		panic(err)
	}

	return hex.EncodeToString(b)[:n]
}
