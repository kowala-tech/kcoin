package testnet

import (
	"testing"

	"docker.io/go-docker"
	"github.com/stretchr/testify/assert"
)

func TestCreationTestNet(t *testing.T) {
	dockerClient, err := docker.NewEnvClient()
	if err != nil {
		t.Fatalf("Error connecting to docker. %s", err)
	}

	dockerEngine := NewDockerEngine(dockerClient)

	testnet := NewTestnet(dockerEngine)
	err = testnet.Start()
	if err != nil {
		t.Fatalf("Error %s", err)
	}

	assert.True(t, testnet.IsValidating())

	testnet.Stop()
}
