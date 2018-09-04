package testnet

import (
	"context"
	"fmt"
	"io"
	"os"

	"io/ioutil"

	"archive/tar"
	"bytes"

	"path/filepath"

	"strings"

	"os/exec"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/container"
	"docker.io/go-docker/api/types/network"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/docker/go-connections/nat"
)

//DockerEngine is a simplified interface for a docker client.
type DockerEngine interface {
	PullImage(image string) error
	BuildImage(dockerFilePath, imageTag string) error
	CreateContainer(image, containerName, networkName string, cmd []string, env []string, ports map[int32]int32) error
	StartContainer(containerID string) error
	CreateNetwork(networkName string) (string, error)
	RemoveNetwork(networkName string) error
	ContainerExists(containerID string) bool
	StopAndRemoveContainer(containerID string) error
	GetLogs(containerID string) (string, error)
	ContainerInspect(containerID string) (types.ContainerJSON, error)
	CopyToContainer(containerID, path string, content []byte) error
	Exec(containerID string, cmd []string) (*ExecResponse, error)
}

//ExecResponse encapsulates the response given by Exec function.
type ExecResponse struct {
	StdOut string
	StdErr string
}

type dockerEngine struct {
	client docker.APIClient
}

func (r *dockerEngine) BuildImage(dockerFilePath, imageTag string) error {
	cmd := exec.Command("docker", "build", "-t", imageTag, "-f", dockerFilePath, filepath.Dir(dockerFilePath))
	cmd.Env = os.Environ()
	if err := cmd.Run(); err != nil {
		return err
	}
	if !cmd.ProcessState.Success() {
		return fmt.Errorf("error building docker image")
	}
	return nil
}

func (r *dockerEngine) Exec(containerID string, cmd []string) (*ExecResponse, error) {
	contExec, err := r.client.ContainerExecCreate(context.Background(), containerID, types.ExecConfig{
		Cmd:          cmd,
		AttachStdout: true,
		AttachStderr: true,
	})
	if err != nil {
		return nil, err
	}
	res, err := r.client.ContainerExecAttach(context.Background(), contExec.ID, types.ExecConfig{})
	if err != nil {
		return nil, err
	}
	defer res.Close()

	var stdOut bytes.Buffer
	var stdErr bytes.Buffer

	_, err = stdcopy.StdCopy(&stdOut, &stdErr, res.Reader)
	if err != nil {
		return nil, err
	}

	return &ExecResponse{
		StdOut: strings.TrimSpace(stdOut.String()),
		StdErr: strings.TrimSpace(stdErr.String()),
	}, nil
}

func (r *dockerEngine) CopyToContainer(containerID, path string, content []byte) error {
	dir := filepath.Dir(path)
	file := filepath.Base(path)

	var buf bytes.Buffer

	w := tar.NewWriter(&buf)
	err := w.WriteHeader(&tar.Header{
		Name: file,
		Mode: 0600,
		Size: int64(len(content)),
	})
	if err != nil {
		return nil
	}

	_, err = w.Write(content)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return nil
	}

	err = r.client.CopyToContainer(context.Background(), containerID, dir, &buf, types.CopyToContainerOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (r *dockerEngine) ContainerInspect(containerID string) (types.ContainerJSON, error) {
	return r.client.ContainerInspect(context.Background(), containerID)
}

func (r *dockerEngine) GetLogs(containerID string) (string, error) {
	log, err := r.client.ContainerLogs(context.Background(), containerID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	})
	if err != nil {
		return "", err
	}
	defer log.Close()

	all, err := ioutil.ReadAll(log)
	if err != nil {
		return "", err
	}
	return string(all), nil
}

func (r *dockerEngine) RemoveNetwork(networkName string) error {
	err := r.client.NetworkRemove(context.Background(), networkName)
	if err != nil {
		return err
	}
	return nil
}

func (r *dockerEngine) StopAndRemoveContainer(containerID string) error {
	err := r.client.ContainerStop(context.Background(), containerID, nil)
	if err != nil {
		return err
	}

	err = r.client.ContainerRemove(context.Background(), containerID, types.ContainerRemoveOptions{})
	if err != nil {
		return err
	}

	return nil
}

//NewDockerEngine returns a simplified docker engine wrapper from docker client.
func NewDockerEngine(client docker.APIClient) DockerEngine {
	return &dockerEngine{
		client: client,
	}
}

//NewDockerEngineWithDefaultClient returns a simplified docker engine wrapper with default docker.APIClient
func NewDockerEngineWithDefaultClient() (DockerEngine, error) {
	client, err := docker.NewEnvClient()
	if err != nil {
		return nil, err
	}

	return NewDockerEngine(client), nil
}

func (r *dockerEngine) ContainerExists(containerID string) bool {
	resp, err := r.client.ContainerInspect(context.Background(), containerID)
	if err != nil {
		return false
	}

	if resp.ID != "" {
		return true
	}

	return false
}

func (r *dockerEngine) StartContainer(containerID string) error {
	err := r.client.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (r *dockerEngine) PullImage(image string) error {
	reader, err := r.client.ImagePull(context.Background(), image, types.ImagePullOptions{})
	if err != nil {
		return err
	}

	io.Copy(os.Stdout, reader)

	return nil
}

func (r *dockerEngine) CreateNetwork(networkName string) (string, error) {
	resp, err := r.client.NetworkCreate(context.Background(), networkName, types.NetworkCreate{})
	if err != nil {
		return "", err
	}

	return resp.ID, nil
}

func (r *dockerEngine) CreateContainer(image, containerName, networkName string, cmd []string, env []string, ports map[int32]int32) error {
	fmt.Println("Creating container.")

	portSpec := make([]string, 0)

	for hostPortRaw, containerPortRaw := range ports {
		portSpec = append(portSpec, fmt.Sprintf("%v:%v", hostPortRaw, containerPortRaw))
	}

	portSet, portMap, err := nat.ParsePortSpecs(portSpec)

	_, err = r.client.ContainerCreate(
		context.Background(),
		&container.Config{
			Image:        image,
			Cmd:          cmd,
			Tty:          true,
			Env:          env,
			ExposedPorts: portSet,
		},
		&container.HostConfig{
			PortBindings: portMap,
		},
		&network.NetworkingConfig{
			EndpointsConfig: map[string]*network.EndpointSettings{
				networkName: {
					NetworkID: networkName,
				},
			},
		},
		containerName,
	)

	if err != nil {
		return err
	}

	return nil
}
