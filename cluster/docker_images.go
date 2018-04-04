package cluster

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"
)

// DockerBuilder builds images docker hosts
type DockerBuilder struct {
	env         []string
	buildErrors []chan error
	mtx         sync.Mutex
}

// NewDockerBuilder returns an instance of a docker image builder.
func NewDockerBuilder(env []string) *DockerBuilder {
	return &DockerBuilder{
		env:         env,
		buildErrors: make([]chan error, 0),
	}
}

func (builder *DockerBuilder) Build(tag, dockerfile string) chan error {
	builder.mtx.Lock()
	defer builder.mtx.Unlock()

	errChn := make(chan error)
	builder.buildErrors = append(builder.buildErrors, errChn)
	go func() {
		errChn <- builder.build(tag, dockerfile)
	}()
	return errChn
}

func (builder *DockerBuilder) Wait() []error {
	log.Println("Waiting for docker images to be built")
	builder.mtx.Lock()
	defer builder.mtx.Unlock()

	timeout := time.After(5 * time.Minute)

	errors := make([]error, 0)
	for _, buildError := range builder.buildErrors {
		select {
		case <-timeout:
			return []error{fmt.Errorf("Timeout building docker images")}
		case err := <-buildError:
			if err != nil {
				errors = append(errors, err)
			}
		}
	}
	builder.buildErrors = make([]chan error, 0)
	return errors
}

func (builder *DockerBuilder) build(tag, dockerfile string) error {
	cmd := exec.Command("docker", "build", "-t", tag, "-f", dockerfile, ".")

	env := os.Environ()
	for _, e := range builder.env {
		env = append(env, e)
	}
	cmd.Env = env
	if err := cmd.Run(); err != nil {
		return err
	}
	if !cmd.ProcessState.Success() {
		return fmt.Errorf("Error building docker image")
	}
	return nil
}
