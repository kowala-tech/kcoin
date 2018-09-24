package testnet

import (
	"testing"

	"context"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/stretchr/testify/assert"
)

func TestCreateGenesisValidator(t *testing.T) {
	dockerClient, err := createDockerClient()
	if err != nil {
		t.Fatalf("Error creating docker client. %s", err)
	}

	dockerEngine := NewDockerEngine(dockerClient)

	networkID, err := dockerEngine.CreateNetwork("validator")
	if err != nil {
		t.Fatalf("Error creating network. %s", err)
	}

	testContext := testContext{
		context:      context.Background(),
		networkID:    networkID,
		dockerEngine: dockerEngine,
	}

	t.Run("Create genesis validator.", func(t *testing.T) {
		theGenesisContent := []byte("Content")

		g, err := NewGenesisValidator(
			dockerEngine,
			testContext.networkID,
			"",
			theGenesisContent,
			common.HexToAddress("0x09438E46Ea66647EA65E4b104C125c82076FDcE5"),
			[]byte("theRawAccount"),
		)
		if err != nil {
			t.Fatalf("Error creating genesis validator. %s", err)
		}

		err = g.Start()
		if err != nil {
			t.Fatalf("Error starting the Genesis Validator. %s", err)
		}

		testContext.nodes = append(testContext.nodes, g)

		assert.True(t, dockerEngine.ContainerExists(g.ID()))
	})

	err = cleanContext(testContext)
	if err != nil {
		t.Fatalf("Error cleaning context %v", err)
	}
}
