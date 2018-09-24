package testnet

import (
	"context"
	"testing"

	"bufio"

	"bytes"

	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/container"
	"docker.io/go-docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"github.com/kowala-tech/toolbox/testnet/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestContainerExists(t *testing.T) {
	t.Run("Returns true when container information is returned and no error.", func(t *testing.T) {
		containerID := "b0044f048d3d"

		mockedClient := &mocks.APIClient{}
		mockedClient.On("ContainerInspect", context.Background(), containerID).
			Return(
				types.ContainerJSON{
					ContainerJSONBase: &types.ContainerJSONBase{
						ID: containerID,
					},
				},
				nil,
			)

		dockerEngine := NewDockerEngine(mockedClient)

		exist := dockerEngine.ContainerExists(containerID)

		assert.True(t, exist)
	})

	t.Run("Returns false when empty container info and an error.", func(t *testing.T) {
		containerID := "b0044f048d3d"

		mockedClient := &mocks.APIClient{}
		mockedClient.On("ContainerInspect", context.Background(), containerID).
			Return(
				types.ContainerJSON{},
				errors.New("random error"),
			)

		dockerEngine := NewDockerEngine(mockedClient)

		exist := dockerEngine.ContainerExists("b0044f048d3d")

		assert.False(t, exist)
	})
}

func TestCreateContainer(t *testing.T) {
	t.Run("Creates a container with a network name", func(t *testing.T) {
		mockedClient := &mocks.APIClient{}
		dockerEngine := NewDockerEngine(mockedClient)

		spec := &NodeSpec{
			Image:     "image/theimage",
			NetworkID: "theNetworkName",
			ID:        "theContainer",
			Cmd: []string{
				"--nodekeyhex", "abcde",
			},
		}

		envVars := []string{
			"HOLA=hola",
		}

		mockedClient.On(
			"ContainerCreate",
			context.Background(),
			&container.Config{
				Image:        spec.Image,
				Cmd:          spec.Cmd,
				Tty:          true,
				Env:          envVars,
				ExposedPorts: nat.PortSet{"8080/tcp": struct{}{}},
			},
			mock.Anything,
			&network.NetworkingConfig{
				EndpointsConfig: map[string]*network.EndpointSettings{
					spec.NetworkID: {
						NetworkID: spec.NetworkID,
					},
				},
			},
			spec.ID,
		).Return(
			container.ContainerCreateCreatedBody{
				ID: "anId",
			},
			nil,
		)

		err := dockerEngine.CreateContainer(spec.Image, spec.ID, spec.NetworkID, spec.Cmd, envVars, map[int32]int32{80: 8080})
		assert.NoError(t, err)
	})
}

func TestExec(t *testing.T) {
	mockedClient := &mocks.APIClient{}
	cmd := []string{"./kcoin", "attach", "--exec", "eth.blockNumber"}

	engine := NewDockerEngine(mockedClient)
	mockedClient.On(
		"ContainerExecCreate",
		context.Background(),
		"validator",
		types.ExecConfig{
			Cmd:          cmd,
			AttachStderr: true,
			AttachStdout: true,
		},
	).Return(
		types.IDResponse{
			ID: "1",
		},
		nil,
	)

	buf := new(bytes.Buffer)
	standardPrefixConsole := []byte{
		0x01,
		0x02,
		0x05,
		0x05,
		0x05,
		0x05,
		0x05,
		0x05,
	}
	buf.Write(append(standardPrefixConsole, []byte("Hello!")...))
	reader := bufio.NewReader(buf)

	mockedConn := &mocks.Conn{}
	mockedClient.On(
		"ContainerExecAttach",
		context.Background(),
		"1",
		types.ExecConfig{},
	).Return(
		types.HijackedResponse{
			Reader: reader,
			Conn:   mockedConn,
		},
		nil,
	)

	mockedConn.On("Close").Return(nil)

	resp, err := engine.Exec("validator", cmd)
	if err != nil {
		t.Fatalf("Error executing command %s", err)
	}

	assert.Equal(t, "", resp.StdOut)
}
