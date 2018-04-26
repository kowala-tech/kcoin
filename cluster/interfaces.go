package cluster

import (
	"math/big"

	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/kcoinclient"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Backend models a kubernetes cluster
type Backend interface {
	// Clientset returns the Clientset to the kubernetes cluster.
	Clientset() (*kubernetes.Clientset, error)

	// RestConfig returns the REST config to the cluster.
	RestConfig() (*rest.Config, error)

	// DockerEnv returns the environment variables necessary to connect to the private docker repository in the kubernetes cluster
	DockerEnv() ([]string, error)

	// IP returns a IP where the cluster can be accessed
	IP() (string, error)
}

type Cluster interface {
	// Connect connects to the backend and initializes its network ID
	Connect() error

	// RpcClient gets a client connected to a node in the cluster
	RpcClient() (*kcoinclient.Client, error)

	// Initialize prepares a new cluster to be ready to start. It saves the networkID for future pods
	// to use it, generates a genesis and stores initial keys in the cluster.
	Initialize(networkID string, seedAccount common.Address) error

	// Cleanup deletes all pods, leaving the kluster in a fresh state
	Cleanup() error

	// Exec runs arbitrary console commands on a specific pod running kcoin.
	Exec(podName, command string) (*ExecResponse, error)

	// GetBalance returns the balance of the coinbase of the specified node
	GetBalance(podName string) (*big.Int, error)

	// RunArchiveNode Runs an archive nodes.
	RunArchiveNode() (string, error)

	// RunBootnode Runs the bootnode
	RunBootnode() error

	// RunRpcNode Runs the rpc node
	RunRpcNode() (string, error)

	// RunNode runs a normal node
	RunNode(string) error

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
