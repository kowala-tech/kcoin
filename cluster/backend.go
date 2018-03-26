package cluster

import (
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var kubeConfigPath = filepath.Join(os.Getenv("HOME"), ".kube", "config")

type k8sCluster struct {
}

func (cluster *k8sCluster) Clientset() (*kubernetes.Clientset, error) {
	log.Println("Connecting to the k8s cluster")
	config, err := cluster.RestConfig()
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}

func (cluster *k8sCluster) RestConfig() (*rest.Config, error) {
	return clientcmd.BuildConfigFromFlags("", kubeConfigPath)
}
