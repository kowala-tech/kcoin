package cluster

import (
	"fmt"
)

type staticCluster struct {
	k8sCluster
	ip         string
	dockerPort int
}

// NewStaticCluster returns a new Backend using an existing k8s cluster
func NewStaticCluster(ip string, dockerPort int) Backend {
	return &staticCluster{
		ip:         ip,
		dockerPort: dockerPort,
	}
}

// Exists Checks if the cluster exists
func (cluster *staticCluster) Exists() bool {
	return true
}

// Create Does nothing because the cluster already exists
func (cluster *staticCluster) Create() error {
	return nil
}

// Delete Does nothing because the cluster can't be deleted
func (cluster *staticCluster) Delete() error {
	return nil
}

// DockerEnv returns the environment variables necessary to connect to the private docker repository in the kubernetes cluster
func (cluster *staticCluster) DockerEnv() ([]string, error) {
	return []string{
		fmt.Sprintf("DOCKER_HOST=tcp://%v:%v", cluster.ip, cluster.dockerPort),
	}, nil
}

// ServiceAddr returns the ip of the cluster
func (cluster *staticCluster) IP() (string, error) {
	return cluster.ip, nil
}
