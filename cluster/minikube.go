package cluster

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var exportRegexp = regexp.MustCompile(`^export\ (.*)=\"(.*)\"$`)

type minikubeCluster struct {
	k8sCluster
	Name string
}

// NewMinikubeCluster returns a new Backend using minikube
func NewMinikubeCluster(name string) Backend {
	return &minikubeCluster{
		Name: name,
	}
}

// Exists Checks if the cluster exists on minikube
func (cluster *minikubeCluster) Exists() bool {
	log.Print("Checking presence of minikube cluster...")
	// Check for VirtualBox service
	statusCmd := exec.Command("minikube", "status", "-p", cluster.Name)
	statusCmd.Run()
	return statusCmd.ProcessState.Success()
}

// Create Creates the cluster using minikube. It requires `minikube` and `virtualbox`.
func (cluster *minikubeCluster) Create() error {
	if err := cluster.assertReady(); err != nil {
		return err
	}
	if cluster.Exists() {
		return fmt.Errorf("The cluster with this name already exists. Delete it first.")
	}
	log.Println("Creating k8s cluster using minikube")
	cmd := exec.Command("minikube", "start", "-p", cluster.Name, "--kubernetes-version", "v1.9.0")
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	return cluster.waitForCluster()
}

// Delete Deletes the cluster.
func (cluster *minikubeCluster) Delete() error {
	if err := cluster.assertReady(); err != nil {
		return err
	}
	log.Println("Deleting k8s cluster using minikube")
	cmd := exec.Command("minikube", "delete", "-p", cluster.Name)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	if !cmd.ProcessState.Success() {
		return fmt.Errorf("error creating cluster")
	}
	return nil
}

// DockerEnv returns the environment variables necessary to connect to the private docker repository in the kubernetes cluster
func (cluster *minikubeCluster) DockerEnv() ([]string, error) {
	statusCmd := exec.Command("minikube", "docker-env", "-p", cluster.Name, "--shell", "bash")
	stdout := &bytes.Buffer{}
	statusCmd.Stdout = stdout
	if err := statusCmd.Run(); err != nil {
		return nil, err
	}
	if !statusCmd.ProcessState.Success() {
		return nil, fmt.Errorf("error getting docker environment variables")
	}
	lines := strings.Split(stdout.String(), "\n")
	goodLines := make([]string, 0)
	for _, line := range lines {
		values := exportRegexp.FindAllStringSubmatch(line, 1)
		if values != nil {
			goodLines = append(goodLines, fmt.Sprintf("%v=%v", values[0][1], values[0][2]))
		}
	}

	return goodLines, nil
}

// ServiceAddr returns the ip:port pair for a specific service running in the cluster
func (cluster *minikubeCluster) ServiceAddr(serviceName string) (string, error) {
	statusCmd := exec.Command("minikube", "service", serviceName, "-p", cluster.Name, "-n", Namespace, "--url", "--format", "http://{{.IP}}:{{.Port}}")
	stdout := &bytes.Buffer{}
	statusCmd.Stdout = stdout
	statusCmd.Stderr = os.Stderr
	if err := statusCmd.Run(); err != nil {
		return "", err
	}
	if !statusCmd.ProcessState.Success() {
		return "", fmt.Errorf("error getting cluster IP")
	}
	url := strings.TrimSpace(stdout.String())
	return url[7:], nil
}

func (cluster *minikubeCluster) assertReady() error {
	log.Print("Checking minikube dependencies...")
	// Make sure we have all required binaries installed
	binaries := []string{"minikube", "VBoxManage"}
	for _, binary := range binaries {
		binCheckCmd := exec.Command("which", binary)
		binCheckCmd.Run()
		if !binCheckCmd.ProcessState.Success() {
			return fmt.Errorf("Missing dependency: %v", binary)
		}
	}

	// Check for VirtualBox service
	vmsListCmd := exec.Command("VBoxManage", "list", "vms")
	vmsListCmd.Run()
	if !vmsListCmd.ProcessState.Success() {
		return fmt.Errorf("VirtualBox doesn't seem to be running.")
	}
	return nil
}

func (cluster *minikubeCluster) waitForCluster() error {
	log.Println("Waiting for cluster to be up and running...")
	return WaitFor(10*time.Second, 5*time.Minute, func() bool {
		client, err := cluster.Clientset()
		if err != nil {
			return false
		}
		_, err = client.CoreV1().ServiceAccounts(apiv1.NamespaceDefault).Get("default", metav1.GetOptions{})
		return err == nil
	})
}
