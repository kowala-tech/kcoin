package cluster

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Backend models a kubernetes cluster
type Backend interface {
	// Exists checks if the cluster exists or not.
	Exists() bool

	// Create Creates the cluster. See specific implementations of this interface for more details.
	Create() error

	// Delete deletes the cluster. See specific implementations of this interface for more details.
	Delete() error

	// Clientset returns the Clientset to the kubernetes cluster.
	Clientset() (*kubernetes.Clientset, error)

	// RestConfig returns the REST config to the cluster.
	RestConfig() (*rest.Config, error)
}

type Cluster interface {
	// Connect connects to the backend and initializes its network ID
	Connect() error

	// Initialize prepares a new cluster to be ready to start. It saves the networkID for future pods
	// to use it, generates a genesis and stores initial keys in the cluster.
	Initialize(networkID string) error

	// Cleanup deletes all pods, leaving the kluster in a fresh state
	Cleanup() error

	// Exec runs arbitrary console commands on a specific pod running kusd.
	Exec(podName, command string) (*ExecResponse, error)

	// GetBalance returns the balance of the coinbase of the specified node
	GetBalance(podName string) (float64, error)

	// RunArchiveNode Runs an archive nodes.
	RunArchiveNode() (string, error)

	// RunBootnode Runs the bootnode
	RunBootnode() error

	// RunGenesisValidator Runs a genesis validator
	RunGenesisValidator() (string, error)

	// RunValidator Runs a standard validator
	RunValidator() (string, error)

	// DeletePod removes a running pod
	DeletePod(podName string) error

	// TriggerGenesisValidation sends a transaction for the genesis validator to start.
	TriggerGenesisValidation() error

	// StoreString stores a configuration value in the cluster
	StoreString(key, value string) error

	// GetString retrieves a configuration value from the cluster
	GetString(key string) (string, error)
}
