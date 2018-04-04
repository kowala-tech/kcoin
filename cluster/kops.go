package cluster

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

type kopsCluster struct {
	k8sCluster
	Name     string
	s3Bucket string
}

// NewMinikubeCluster returns a new Backend using kops on aws
func NewKopsCluster(name string) Backend {
	return &kopsCluster{
		Name:     name + ".k8s.local",
		s3Bucket: "cluster." + name,
	}
}

// Exists Checks if the cluster exists on aws
func (cluster *kopsCluster) Exists() bool {
	log.Print("Checking presence of kops cluster...")

	cmd := exec.Command("kops", "validate", "cluster", "--name", cluster.Name, "--state", "s3://"+cluster.s3Bucket)
	cmd.Run()
	return cmd.ProcessState.Success()
}

// Create Creates the cluster. It requires `kops`, `kubectl` and `aws` binaries to be available, and aws must be configured.
func (cluster *kopsCluster) Create() error {
	if err := cluster.assertReady(); err != nil {
		return err
	}
	if err := cluster.createStateStore(); err != nil {
		return err
	}
	if err := cluster.createCluster(); err != nil {
		return err
	}
	if err := cluster.waitForCluster(); err != nil {
		return err
	}
	return nil
}

// Delete deletes the cluster and its state store.
func (cluster *kopsCluster) Delete() error {
	if err := cluster.assertReady(); err != nil {
		return err
	}
	if err := cluster.deleteCluster(); err != nil {
		return err
	}
	if err := cluster.deleteStateStore(); err != nil {
		return err
	}
	return nil
}

func (cluster *kopsCluster) createCluster() error {
	log.Println("Creating k8s cluster using kops")
	cmd := exec.Command("kops", "create", "cluster", "--name", cluster.Name, "--yes", "--state", "s3://"+cluster.s3Bucket, "--zones", "us-east-2b")
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()
	if !cmd.ProcessState.Success() {
		return fmt.Errorf("Error creating cluster")
	}
	return nil
}

func (cluster *kopsCluster) deleteCluster() error {
	log.Println("Deleting k8s cluster using kops")
	cmd := exec.Command("kops", "delete", "cluster", "--name", cluster.Name, "--yes", "--state", "s3://"+cluster.s3Bucket)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()
	if !cmd.ProcessState.Success() {
		return fmt.Errorf("Error creating cluster")
	}
	return nil
}

func (cluster *kopsCluster) waitForCluster() error {
	log.Println("Waiting for cluster to be up and running...")
	return WaitFor(10*time.Second, 5*time.Minute, func() bool {
		cmd := exec.Command("kops", "validate", "cluster", "--name", cluster.Name, "--state", "s3://"+cluster.s3Bucket)
		cmd.Run()
		return cmd.ProcessState.Success()
	})
}

func (cluster *kopsCluster) assertReady() error {
	log.Print("Checking kops dependencies...")
	// Make sure we have all required binaries installed
	binaries := []string{"kops", "kubectl", "aws"}
	for _, binary := range binaries {
		binCheckCmd := exec.Command("which", binary)
		binCheckCmd.Run()
		if !binCheckCmd.ProcessState.Success() {
			return fmt.Errorf("Missing dependency: %v", binary)
		}
	}

	// Check for aws configuration
	listUsersCmd := exec.Command("aws", "configure", "get", "aws_access_key_id")
	listUsersCmd.Run()
	if !listUsersCmd.ProcessState.Success() {
		return fmt.Errorf("aws is not configured. Please run `aws configure` first")
	}
	return nil
}

func (cluster *kopsCluster) createStateStore() error {
	log.Print("Creating state store on S3...")
	cmd := exec.Command("aws", "s3api", "create-bucket", "--bucket", cluster.s3Bucket, "--region", "us-east-2", "--create-bucket-configuration", "LocationConstraint=us-east-2")
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()
	if !cmd.ProcessState.Success() {
		return fmt.Errorf("Error creating state store S3 bucket")
	}
	return nil
}

func (cluster *kopsCluster) deleteStateStore() error {
	log.Print("Deleting state store from S3...")
	cmd := exec.Command("aws", "s3api", "delete-bucket", "--bucket", cluster.s3Bucket, "--region", "us-east-2")
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()
	if !cmd.ProcessState.Success() {
		return fmt.Errorf("Error deleting state store S3 bucket")
	}
	return nil
}
