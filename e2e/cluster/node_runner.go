package cluster

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/docker/go-connections/nat"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/log"
)

type NodeRunner interface {
	Run(node *NodeSpec) error
	Stop(nodeID NodeID) error
	StopAll() error
	Log(nodeID NodeID) (string, error)
	HostIP() string
	IP(nodeID NodeID) (string, error)
	Exec(nodeID NodeID, command []string) (*ExecResponse, error)
}

type ExecResponse struct {
	StdOut string
	StdErr string
}

type dockerNodeRunner struct {
	pullMtx sync.Mutex

	runningNodes map[NodeID]bool
	client       *client.Client
	logsToStdout bool
	logsDir      string
	logPrefix    string
}

type NewNodeRunnerOpts struct {
	Prefix       string
	LogsToStdout bool
	LogsDir      string
}

func NewDockerNodeRunner(opts *NewNodeRunnerOpts) (*dockerNodeRunner, error) {
	client, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}
	return &dockerNodeRunner{
		client:       client,
		runningNodes: make(map[NodeID]bool, 0),
		logsDir:      opts.LogsDir,
		logsToStdout: opts.LogsToStdout,
		logPrefix:    opts.Prefix,
	}, nil
}

func (runner *dockerNodeRunner) Run(node *NodeSpec) error {
	if err := runner.pullIfNecessary(node.Image); err != nil {
		return err
	}

	portSpec := make([]string, 0)
	for hostPortRaw, containerPortRaw := range node.PortMapping {
		portSpec = append(portSpec, fmt.Sprintf("%v:%v", hostPortRaw, containerPortRaw))
	}

	portSet, portMap, err := nat.ParsePortSpecs(portSpec)
	if err != nil {
		return err
	}

	config := &container.Config{
		Image:        node.Image,
		Cmd:          node.Cmd,
		Env:          node.Env,
		ExposedPorts: portSet,
	}
	hostConfig := &container.HostConfig{
		PortBindings: portMap,
	}

	_, err = runner.client.ContainerCreate(context.Background(), config, hostConfig, nil, node.ID.String())
	if err != nil {
		return err
	}
	runner.runningNodes[node.ID] = true

	if node.Files != nil {
		for filename, contents := range node.Files {
			if err := runner.copyFile(node.ID, filename, contents); err != nil {
				return err
			}
		}
	}

	err = runner.client.ContainerStart(context.Background(), node.ID.String(), types.ContainerStartOptions{})
	if err != nil {
		return err
	}

	logStream := os.Stdout
	if !runner.logsToStdout {
		logFilename := filepath.Join(runner.logsDir, fmt.Sprintf("%s-%v.log", runner.logPrefix, node.ID))
		logFile, err := os.OpenFile(logFilename, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0777)
		if err != nil {
			log.Error(fmt.Sprintf("error creating container logs file %q: %s", logFilename, err))
			return err
		}
		logStream = logFile
	}

	go func() {
		defer logStream.Close()
		err := runner.logStream(node.ID, logStream)
		if err != nil {
			log.Error(fmt.Sprintf("error saving container logs to file: %s", err))
		}
	}()

	if node.IsReadyFn != nil {
		return common.WaitFor("Node starts", 1*time.Second, 20*time.Second, func() error {
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
	ip := os.Getenv("DOCKER_PUBLIC_IP")
	if ip == "" {
		parsed, err := url.Parse(os.Getenv("DOCKER_HOST"))
		if err == nil {
			ip = parsed.Hostname()
		}
	}
	if ip == "" {
		ip = "localhost"
	}
	return ip
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
	res, err := runner.client.ContainerExecAttach(context.Background(), contExec.ID, types.ExecConfig{})
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
	return []string{"./kcoin", "attach", "/root/.kcoin/kusd/kcoin.ipc", "--exec", command}
}

func (runner *dockerNodeRunner) logStream(nodeID NodeID, w io.Writer) error {
	log, err := runner.client.ContainerLogs(context.Background(), nodeID.String(), types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
	})
	if err != nil {
		return err
	}
	defer log.Close()

	_, err = stdcopy.StdCopy(w, w, log)
	return err
}

func (runner *dockerNodeRunner) pullIfNecessary(image string) error {
	runner.pullMtx.Lock()
	defer runner.pullMtx.Unlock()

	images, err := runner.client.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		return err
	}
	for _, img := range images {
		for _, v := range img.RepoTags {
			if v == image {
				return nil
			}
		}
	}
	r, err := runner.client.ImagePull(context.Background(), image, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	ioutil.ReadAll(r)
	return nil
}
