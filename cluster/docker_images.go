package cluster

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"
)

// builds docker images in parallel
func (client *cluster) buildLocalDockerImages() error {
	failed := false
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := client.buildKusdLocalDockerImage("kowalatech/bootnode:dev", "bootnode.Dockerfile"); err != nil {
			failed = true
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := client.buildKusdLocalDockerImage("kowalatech/kusd:dev", "kcoin.Dockerfile"); err != nil {
			failed = true
		}
	}()

	done := make(chan struct{})
	go func() {
		wg.Wait()
		done <- struct{}{}
	}()

	select {
	case <-time.After(5 * time.Minute):
		return fmt.Errorf("Timeout building docker images")
	case <-done:
		if failed {
			return fmt.Errorf("Error building docker images")
		}
	}
	return nil
}

func (client *cluster) buildKusdLocalDockerImage(tag, dockerfile string) error {
	cmd := exec.Command("docker", "build", "-t", tag, "-f", dockerfile, ".")
	dockerEnv, err := client.Backend.DockerEnv()
	if err != nil {
		return err
	}
	env := os.Environ()
	for _, e := range dockerEnv {
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
