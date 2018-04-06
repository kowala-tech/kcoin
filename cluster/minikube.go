package cluster

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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
		return fmt.Errorf("Error creating cluster")
	}
	return nil
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
