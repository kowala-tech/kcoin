package cluster

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"regexp"
	"time"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	bootnodeEnodeKey = "bootnode-enode"
	bootnodePort     = 33445
)

var (
	src               = rand.New(rand.NewSource(time.Now().UnixNano()))
	enodeSecretRegexp = regexp.MustCompile(`enode://([a-f0-9]*)@`)
)

// RunBootnode Runs the bootnode
func (client *cluster) RunBootnode() error {
	log.Println("Running bootnode")
	pod := bootnodePod("bootnode", randStringBytes(64))

	_, err := client.Clientset.CoreV1().Pods(Namespace).Create(pod)
	if err != nil {
		return err
	}
	err = client.waitForPod("bootnode")
	if err != nil {
		return err
	}
	enode, err := client.bootnodeEnode()
	if err != nil {
		return err
	}
	return client.StoreString(bootnodeEnodeKey, enode)
}

func (client *cluster) bootnodeEnode() (string, error) {
	log.Println("Reading bootnode enode from its stdout")
	pods := client.Clientset.CoreV1().Pods(Namespace)
	pod, err := pods.Get("bootnode", metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	logStream, err := pods.GetLogs("bootnode", &apiv1.PodLogOptions{}).Stream()
	if err != nil {
		return "", err
	}

	defer logStream.Close()
	data, err := ioutil.ReadAll(logStream)
	if err != nil {
		return "", err
	}
	found := enodeSecretRegexp.FindSubmatch(data)
	if len(found) != 2 {
		return "", fmt.Errorf("Bootnode enode not found in the logs.")
	}

	enodeSecret := string(found[1])
	ip := pod.Status.PodIP

	return fmt.Sprintf(`enode://%v@%v:%v`, enodeSecret, ip, bootnodePort), nil
}

func randStringBytes(n int) string {
	b := make([]byte, n/2)

	if _, err := src.Read(b); err != nil {
		panic(err)
	}

	return hex.EncodeToString(b)[:n]
}

func bootnodePod(podName, hexkey string) *apiv1.Pod {
	return &apiv1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: podName,
			Labels: map[string]string{
				"network": "testnet",
				"app":     "bootnode",
				"name":    podName,
			},
		},
		Spec: apiv1.PodSpec{
			Containers: []apiv1.Container{
				{
					Name:            podName,
					Image:           "kowalatech/bootnode:dev",
					ImagePullPolicy: apiv1.PullAlways,
					Ports: []apiv1.ContainerPort{
						{
							ContainerPort: 32233,
							Protocol:      apiv1.ProtocolUDP,
						},
					},
					Args: []string{
						"--nodekeyhex", hexkey,
						"--addr", fmt.Sprintf(":%v", bootnodePort),
					},
				},
			},
		},
	}
}
