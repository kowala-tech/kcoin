package cluster

import (
	"fmt"
	"log"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RunArchiveNodes Runs archive nodes numbered from `from` to `to`.
func (client *cluster) RunArchiveNode() (string, error) {
	log.Println("Running archive nodes")
	idx, err := client.getNextArchiveNodeIndex()
	if err != nil {
		return "", err
	}

	podName := fmt.Sprintf("archive-node-%d", idx)
	bootnode, err := client.GetString(bootnodeEnodeKey)
	if err != nil {
		return "", err
	}
	pod := archiveNodePod(podName, client.NetworkID, bootnode, int32(31310+idx))
	useGenesisFromConfigmap(&pod.Spec)

	_, err = client.Clientset.CoreV1().Pods(Namespace).Create(pod)
	if err != nil {
		return "", err
	}

	log.Println("Waiting for archive node to be running")
	if err := client.waitForKusdPod(podName); err != nil {
		return "", err
	}
	return podName, nil
}

func (client *cluster) getNextArchiveNodeIndex() (int, error) {
	list, err := client.Clientset.CoreV1().Pods(Namespace).List(metav1.ListOptions{
		LabelSelector: "app=archive_node",
	})
	if err != nil {
		return 0, err
	}
	max := 0
	var index int
	for _, pod := range list.Items {
		_, err := fmt.Sscanf(pod.Name, "archive-node-%d", &index)
		if err != nil {
			return 0, err
		}
		if index > max {
			max = index
		}
	}
	return max + 1, nil
}

func archiveNodePod(podName, networkID, bootnode string, port int32) *apiv1.Pod {
	return &apiv1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: podName,
			Labels: map[string]string{
				"network": "testnet",
				"app":     "archive_node",
				"name":    podName,
			},
		},
		Spec: apiv1.PodSpec{
			Containers: []apiv1.Container{
				{
					Name:            podName,
					Image:           "kowalatech/kusd:dev",
					ImagePullPolicy: apiv1.PullAlways,
					Ports: []apiv1.ContainerPort{
						{
							ContainerPort: port,
							Protocol:      apiv1.ProtocolUDP,
						},
						{
							ContainerPort: port,
							Protocol:      apiv1.ProtocolTCP,
						},
					},
					Args: []string{
						"--syncmode", "fast",
						"--bootnodes", bootnode,
						"--networkid", networkID,
						"--gasprice", "1",
						"--verbosity", "6",
						"--port", fmt.Sprintf("%v", port),
					},
				},
			},
		},
	}
}
