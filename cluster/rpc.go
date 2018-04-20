package cluster

import (
	"fmt"
	"log"

	"github.com/kowala-tech/kcoin/kcoinclient"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const rpcPodName = "rpc"
const rpcPort = 31301

// RunRpcNode Runs the rpc node
func (client *cluster) RunRpcNode() (string, error) {
	log.Println("Running rpc node")
	bootnode, err := client.GetString(bootnodeEnodeKey)
	if err != nil {
		return "", err
	}
	pod := RpcNodePod(rpcPodName, client.NetworkID, bootnode, rpcPort)
	useGenesisFromConfigmap(&pod.Spec)

	_, err = client.Clientset.CoreV1().Pods(client.Namespace).Create(pod)
	if err != nil {
		return "", err
	}
	err = client.waitForKusdPod(rpcPodName)
	if err != nil {
		return "", err
	}

	service := RpcNodeService(rpcPodName, rpcPort)
	_, err = client.Clientset.CoreV1().Services(client.Namespace).Create(service)
	if err != nil {
		return "", err
	}

	return rpcPodName, nil
}

// RpcClient gets a client connected to the rpc node in the cluster
func (client *cluster) RpcClient() (*kcoinclient.Client, error) {
	ip, err := client.Backend.IP()
	if err != nil {
		return nil, err
	}
	rpcAddr := fmt.Sprintf("http://%v:%v", ip, rpcPort)
	return kcoinclient.Dial(rpcAddr)
}

func RpcNodePod(podName, networkID, bootnode string, port int32) *apiv1.Pod {
	return &apiv1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: podName,
			Labels: map[string]string{
				"network": "testnet",
				"app":     "rpc_node",
				"name":    podName,
			},
		},
		Spec: apiv1.PodSpec{
			Containers: []apiv1.Container{
				{
					Name:            podName,
					Image:           "kowalatech/kusd:dev",
					ImagePullPolicy: apiv1.PullNever,
					Ports: []apiv1.ContainerPort{
						{
							ContainerPort: port,
							Protocol:      apiv1.ProtocolTCP,
						},
					},
					Args: []string{
						"--syncmode", "fast",
						"--bootnodes", bootnode,
						"--networkid", networkID,
						"--verbosity", "6",
						"--rpc",
						"--rpcaddr", "0.0.0.0",
						"--rpccorsdomain", "*",
						"--rpcport", fmt.Sprintf("%v", port),
					},
				},
			},
		},
	}
}

func RpcNodeService(serviceName string, port int32) *apiv1.Service {
	return &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: serviceName,
		},
		Spec: apiv1.ServiceSpec{
			Type: apiv1.ServiceTypeNodePort,
			Selector: map[string]string{
				"network": "testnet",
				"app":     "rpc_node",
			},
			Ports: []apiv1.ServicePort{
				{
					Port:     port,
					NodePort: port,
				},
			},
		},
	}
}
