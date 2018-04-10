package cluster

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const genesisValidatorPodName = "genesis-validator"

// RunGenesisValidator Runs a genesis validator
func (client *cluster) RunGenesisValidator() (string, error) {
	log.Println("Running genesis validator")
	bootnode, err := client.GetString(bootnodeEnodeKey)
	if err != nil {
		return "", err
	}
	pod := genesisValidatorPod(genesisValidatorPodName, client.NetworkID, "d6e579085c82329c89fca7a9f012be59028ed53f", bootnode, 31300)
	useGenesisFromConfigmap(&pod.Spec)
	usePasswordFromConfigmap(&pod.Spec)
	useKeyFromConfigmap(&pod.Spec, "d6e579085c82329c89fca7a9f012be59028ed53f", "UTC--2018-01-16T16-31-38.006625000Z--d6e579085c82329c89fca7a9f012be59028ed53f")

	_, err = client.Clientset.CoreV1().Pods(Namespace).Create(pod)
	if err != nil {
		return "", err
	}
	return genesisValidatorPodName, client.waitForKusdPod(genesisValidatorPodName)
}

// TriggerGenesisValidation sends a transaction for the genesis validator to start.
func (client *cluster) TriggerGenesisValidation() error {
	log.Println("Triggering genesis validation")
	_, err := client.Exec(
		genesisValidatorPodName,
		`
			personal.unlockAccount(eth.coinbase, "test");
			eth.sendTransaction({from:eth.coinbase,to: "0x259be75d96876f2ada3d202722523e9cd4dd917d",value: 1})
		`)
	if err != nil {
		return err
	}

	return WaitFor(2*time.Second, 20*time.Second, func() bool {
		res, err := client.Exec(genesisValidatorPodName, `eth.blockNumber`)
		if err != nil {
			return false
		}
		parsed, err := strconv.Atoi(strings.TrimSpace(res.StdOut))
		if err != nil {
			return false
		}
		return parsed > 0
	})
}

func genesisValidatorPod(podName, networkID, pub_key, bootnode string, port int32) *apiv1.Pod {
	return &apiv1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: podName,
			Labels: map[string]string{
				"network": "testnet",
				"app":     "genesis_validator",
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
							Protocol:      apiv1.ProtocolUDP,
						},
						{
							ContainerPort: port,
							Protocol:      apiv1.ProtocolTCP,
						},
					},
					Args: []string{
						"--validate",
						"--deposit", "100001",
						"--password", "/kcoin/password.txt",
						"--unlock", fmt.Sprintf("0x%v", pub_key),
						"--coinbase", fmt.Sprintf("0x%v", pub_key),
						"--syncmode", "full",
						"--bootnodes", bootnode,
						"--networkid", networkID,
						"--lightpeers", "20",
						"--cache", "128",
						"--gasprice", "1",
						"--txpool.journal", "transactions.rlp",
						"--txpool.rejournal", "1h",
						"--txpool.pricelimit", "1",
						"--txpool.pricebump", "10",
						"--txpool.accountslots", "16",
						"--txpool.globalslots", "4096",
						"--txpool.accountqueue", "64",
						"--txpool.globalqueue", "1024",
						"--txpool.lifetime", "3h",
						"--gpoblocks", "10",
						"--gpopercentile", "50",
						"--maxpeers", "25",
						"--verbosity", "6",
						"--port", fmt.Sprintf("%v", port),
					},
				},
			},
		},
	}
}
