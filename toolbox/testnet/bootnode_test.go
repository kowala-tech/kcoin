package testnet

import (
	"testing"

	"context"

	"github.com/stretchr/testify/assert"
)

func TestCreateBootNode(t *testing.T) {
	dockerClient, err := createDockerClient()
	if err != nil {
		t.Fatalf("Error creating docker client. %s", err)
	}

	dockerEngine := NewDockerEngine(dockerClient)

	_, err = dockerEngine.CreateNetwork("integration")
	if err != nil {
		t.Fatalf("Error creating network. %s", err)
	}

	testContext := testContext{
		context:      context.Background(),
		networkID:    "integration",
		dockerEngine: dockerEngine,
	}

	t.Run("Create bootnode", func(t *testing.T) {
		g, err := NewBootNode(dockerEngine, testContext.networkID)
		if err != nil {
			t.Fatalf("Error creating boot node. %s", err)
		}

		err = g.Start()
		if err != nil {
			t.Fatalf("Error starting the boot node. %s", err)
		}

		testContext.nodes = append(testContext.nodes, g)

		assert.True(t, dockerEngine.ContainerExists(g.ID()))
		assert.Contains(t, g.Enode(), ":33445")
	})

	err = cleanContext(testContext)
	if err != nil {
		t.Fatalf("Error cleaning context %v", err)
	}
}
