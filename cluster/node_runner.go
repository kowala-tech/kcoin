package cluster

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type NodeRunner interface {
	Run(node *Node) error
	Log(node *Node) (string, error)
	IP(node *Node) (string, error)
	BuildDockerImage(tag, dockerFile string) error
}

type dockerNodeRunner struct {
	client *client.Client
}

func NewDockerNodeRunner() (*dockerNodeRunner, error) {
	client, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}
	return &dockerNodeRunner{
		client: client,
	}, nil
}

func (runner *dockerNodeRunner) Run(node *Node) error {
	_, err := runner.client.ContainerCreate(context.Background(), &container.Config{
		Image: node.Image,
		Cmd:   node.Cmd,
	}, nil, nil, node.Name)
	if err != nil {
		return err
	}

	for filename, contents := range node.Files {
		if err := runner.copyFile(node, filename, contents); err != nil {
			return err
		}
	}

	err = runner.client.ContainerStart(context.Background(), node.Name, types.ContainerStartOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (runner *dockerNodeRunner) Log(node *Node) (string, error) {
	log, err := runner.client.ContainerLogs(context.Background(), node.Name, types.ContainerLogsOptions{
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

func (runner *dockerNodeRunner) IP(node *Node) (string, error) {
	container, err := runner.client.ContainerInspect(context.Background(), node.Name)
	if err != nil {
		return "", err
	}
	return container.NetworkSettings.IPAddress, nil
}

func (runner *dockerNodeRunner) BuildDockerImage(tag, dockerFile string) error {
	cmd := exec.Command("docker", "build", "-t", tag, "-f", path.Join(rootPath, dockerFile), rootPath)
	cmd.Env = os.Environ()
	if err := cmd.Run(); err != nil {
		return err
	}
	if !cmd.ProcessState.Success() {
		return fmt.Errorf("Error building docker image")
	}
	return nil
}

func (runner *dockerNodeRunner) copyFile(node *Node, filename string, contents []byte) error {
	dir := filepath.Dir(filename)
	file := filepath.Base(filename)

	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	hdr := &tar.Header{
		Name: file,
		Mode: 0600,
		Size: int64(len(contents)),
	}
	if err := tw.WriteHeader(hdr); err != nil {
		return err
	}
	if _, err := tw.Write([]byte(contents)); err != nil {
		return err
	}
	if err := tw.Close(); err != nil {
		return err
	}

	err := runner.client.CopyToContainer(context.Background(),
		node.Name,
		dir,
		&buf,
		types.CopyToContainerOptions{})
	if err != nil {
		return err
	}

	return nil
}
