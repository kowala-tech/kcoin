package cluster

import (
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type k8sBackend struct {
	ip        string
	masterUrl string
	token     string
	env       []string
}

// Newk8sBackend returns a new Backend using an existing k8s cluster
func NewK8SBackend(ip, masterUrl, token string, env []string) Backend {
	return &k8sBackend{
		masterUrl: masterUrl,
		ip:        ip,
		token:     token,
		env:       env,
	}
}

func (backend *k8sBackend) RestConfig() (*rest.Config, error) {
	return &rest.Config{
		Host:        backend.masterUrl,
		BearerToken: backend.token,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}, nil
}

func (backend *k8sBackend) Clientset() (*kubernetes.Clientset, error) {
	log.Println("Connecting to the k8s cluster")
	config, err := backend.RestConfig()
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}

// DockerEnv returns the environment variables necessary to connect to the private docker repository in the kubernetes cluster
func (backend *k8sBackend) DockerEnv() ([]string, error) {
	return backend.env, nil
}

// IP returns the IP of the box to access open services
func (backend *k8sBackend) IP() (string, error) {
	return backend.ip, nil
}
