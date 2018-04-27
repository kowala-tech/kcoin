package cluster

import (
	"fmt"
	"log"
	"math/big"
	"time"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RunValidator Runs a genesis validator
func (client *cluster) RunValidator() (string, error) {
	log.Println("Running a validator")
	idx, err := client.getNextValidatorIndex()
	if err != nil {
		return "", err
	}

	podName := fmt.Sprintf("validator-%d", idx)
	log.Printf("New validator name is `%v`", podName)

	if err := client.runValidatorPod(podName, int32(31910+idx)); err != nil {
		return "", err
	}

	if err := client.fundValidator(podName); err != nil {
		return "", err
	}

	if err := client.startValidation(podName); err != nil {
		return "", err
	}

	return podName, nil
}

func (client *cluster) runValidatorPod(podName string, port int32) error {
	bootnode, err := client.GetString(bootnodeEnodeKey)
	if err != nil {
		return err
	}

	pod := validatorPod(podName, client.NetworkID, bootnode, port)
	useGenesisFromConfigmap(&pod.Spec)
	useWALFromConfigmap(&pod.Spec)

	if _, err = client.Clientset.CoreV1().Pods(client.Namespace).Create(pod); err != nil {
		return err
	}
	if err = client.waitForKusdPod(podName); err != nil {
		return err
	}
	if err = client.waitForInitialSync(podName); err != nil {
		return err
	}
	return nil
}

func (client *cluster) fundValidator(podName string) error {
	resp, err := client.Exec(podName, `eth.coinbase`)
	if err != nil {
		return err
	}
	coinbaseQuotes := resp.StdOut

	log.Println("Transferring 50x min deposit to the new validator")
	_, err = client.Exec(
		GenesisValidatorPodName,
		fmt.Sprintf(`eth.sendTransaction({from:eth.coinbase, to: %v, value: 50*validator.getMinimumDeposit()})`, coinbaseQuotes))
	if err != nil {
		return err
	}

	log.Println("Waiting for funds to be available")
	err = WaitFor(2*time.Second, 1*time.Minute, func() bool {
		balance, err := client.GetBalance(podName)
		if err != nil {
			return false
		}
		return balance.Cmp(big.NewInt(0)) > 0
	})
	return err
}
func (client *cluster) startValidation(podName string) error {
	log.Println("Checking initial balance")
	initialBalance, err := client.GetBalance(podName)
	if err != nil {
		return err
	}

	log.Println("Setting deposit and starting validation")
	command := `
		personal.unlockAccount(eth.coinbase, "test");
		validator.setDeposit(10*validator.getMinimumDeposit());
		validator.start();
	`
	_, err = client.Exec(podName, command)
	if err != nil {
		return err
	}

	log.Println("Waiting for deposit to be accepted")
	err = WaitFor(2*time.Second, 1*time.Minute, func() bool {
		balance, err := client.GetBalance(podName)
		if err != nil {
			return false
		}
		return balance.Cmp(initialBalance) < 0
	})

	return err
}

func (client *cluster) getNextValidatorIndex() (int, error) {
	list, err := client.Clientset.CoreV1().Pods(client.Namespace).List(metav1.ListOptions{
		LabelSelector: "app=standard_validator",
	})
	if err != nil {
		return 0, err
	}
	max := 0
	var index int
	for _, pod := range list.Items {
		_, err := fmt.Sscanf(pod.Name, "validator-%d", &index)
		if err != nil {
			return 0, err
		}
		if index > max {
			max = index
		}
	}
	return max + 1, nil
}

func validatorPod(podName, networkID, bootnode string, port int32) *apiv1.Pod {
	return &apiv1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: podName,
			Labels: map[string]string{
				"network": "testnet",
				"app":     "standard_validator",
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
						"--syncmode", "full",
						"--deposit", "100001",
						"--gasprice", "20",
						"--bootnodes", bootnode,
						"--networkid", networkID,
						"--verbosity", "6",
						"--port", fmt.Sprintf("%v", port),
					},
				},
			},
		},
	}
}
