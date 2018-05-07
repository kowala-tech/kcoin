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
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/kowala-tech/kcoin/common"
)

type NodeRunner interface {
	Run(node *NodeSpec) error
	Log(nodeID NodeID) (string, error)
	IP(nodeID NodeID) (string, error)
	Exec(nodeID NodeID, command []string) (*ExecResponse, error)
	BuildDockerImage(tag, dockerFile string) error
}

type ExecResponse struct {
	StdOut string
	StdErr string
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

func (runner *dockerNodeRunner) Run(node *NodeSpec) error {
	_, err := runner.client.ContainerCreate(context.Background(), &container.Config{
		Image: node.Image,
		Cmd:   node.Cmd,
	}, nil, nil, node.ID.String())
	if err != nil {
		return err
	}

	for filename, contents := range node.Files {
		if err := runner.copyFile(node.ID, filename, contents); err != nil {
			return err
		}
	}

	err = runner.client.ContainerStart(context.Background(), node.ID.String(), types.ContainerStartOptions{})
	if err != nil {
		return err
	}

	if node.IsReadyFn != nil {
		return common.WaitFor("Node starts", 1*time.Second, 20*time.Second, func() bool {
			return node.IsReadyFn(runner)
		})
	}
	return nil
}

func (runner *dockerNodeRunner) Log(nodeID NodeID) (string, error) {
	log, err := runner.client.ContainerLogs(context.Background(), nodeID.String(), types.ContainerLogsOptions{
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

func (runner *dockerNodeRunner) IP(nodeID NodeID) (string, error) {
	container, err := runner.client.ContainerInspect(context.Background(), nodeID.String())
	if err != nil {
		return "", err
	}
	return container.NetworkSettings.IPAddress, nil
}

func (runner *dockerNodeRunner) Exec(nodeID NodeID, command []string) (*ExecResponse, error) {
	contExec, err := runner.client.ContainerExecCreate(context.Background(), nodeID.String(), types.ExecConfig{
		Cmd:          command,
		AttachStdout: true,
		AttachStderr: true,
	})
	if err != nil {
		return nil, err
	}
	res, err := runner.client.ContainerExecAttach(context.Background(), contExec.ID, types.ExecStartCheck{})
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

func (runner *dockerNodeRunner) copyFile(nodeID NodeID, filename string, contents []byte) error {
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
		nodeID.String(),
		dir,
		&buf,
		types.CopyToContainerOptions{})
	if err != nil {
		return err
	}

	return nil
}

func KcoinExecCommand(command string) []string {
	return []string{"./kcoin", "attach", "--exec", command}
}
