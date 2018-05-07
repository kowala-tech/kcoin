package cluster

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net"
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
	"github.com/docker/go-connections/nat"
	"github.com/kowala-tech/kcoin/common"
)

type NodeRunner interface {
	Run(node *NodeSpec) error
	Stop(nodeID NodeID) error
	StopAll() error
	Log(nodeID NodeID) (string, error)
	HostIP() string
	IP(nodeID NodeID) (string, error)
	Exec(nodeID NodeID, command []string) (*ExecResponse, error)
	BuildDockerImage(tag, dockerFile string) error
}

type ExecResponse struct {
	StdOut string
	StdErr string
}

type dockerNodeRunner struct {
	runningNodes map[NodeID]bool
	client       *client.Client
}

func NewDockerNodeRunner() (*dockerNodeRunner, error) {
	client, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}
	return &dockerNodeRunner{
		client:       client,
		runningNodes: make(map[NodeID]bool, 0),
	}, nil
}

func (runner *dockerNodeRunner) Run(node *NodeSpec) error {
	portSpec := make([]string, 0)

	for hostPortRaw, containerPortRaw := range node.PortMapping {
		portSpec = append(portSpec, fmt.Sprintf("%v:%v", hostPortRaw, containerPortRaw))
	}

	portSet, portMap, err := nat.ParsePortSpecs(portSpec)
	if err != nil {
		return err
	}
	_, err = runner.client.ContainerCreate(context.Background(), &container.Config{
		Image:        node.Image,
		Cmd:          node.Cmd,
		ExposedPorts: portSet,
	}, &container.HostConfig{
		PortBindings: portMap,
	}, nil, node.ID.String())
	if err != nil {
		return err
	}
	runner.runningNodes[node.ID] = true

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

func (runner *dockerNodeRunner) Stop(nodeID NodeID) error {
	runner.runningNodes[nodeID] = false
	return runner.client.ContainerRemove(context.Background(), nodeID.String(), types.ContainerRemoveOptions{Force: true})
}

func (runner *dockerNodeRunner) StopAll() error {
	for id, running := range runner.runningNodes {
		if running {
			if err := runner.Stop(id); err != nil {
				return err
			}
		}
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

func (runner *dockerNodeRunner) HostIP() string {
	hostIP := runner.client.DaemonHost()
	if addr := net.ParseIP(hostIP); addr == nil {
		// Assume non-ip based hosts mean localhost
		return "localhost"
	}
	return hostIP
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
	if _, err := tw.Write(contents); err != nil {
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
