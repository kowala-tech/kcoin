package cluster

import (
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const Namespace = "kcoin"

type cluster struct {
	Backend   Backend
	Clientset *kubernetes.Clientset
	NetworkID string
}

// NewClient Returns a client interface to the specified backend k8s
func NewCluster(backend Backend) Cluster {
	return &cluster{
		Backend: backend,
	}
}

func (client *cluster) Connect() error {
	clientset, err := client.Backend.Clientset()
	if err != nil {
		return err
	}
	client.Clientset = clientset

	// Ignore error, the config storage might not be initialized yet for new clusters
	networkID, _ := client.GetString(NetworkIDKey)
	client.NetworkID = networkID
	return nil
}

func (client *cluster) Cleanup() error {
	var zero int64 = 0
	err := client.Clientset.CoreV1().Pods(Namespace).DeleteCollection(&metav1.DeleteOptions{
		GracePeriodSeconds: &zero,
	}, metav1.ListOptions{})
	if err != nil {
		return err
	}
	err = WaitFor(1*time.Second, 20*time.Second, func() bool {
		list, err := client.Clientset.CoreV1().Pods(Namespace).List(metav1.ListOptions{})
		if err != nil {
			return false
		}
		return len(list.Items) == 0
	})
	if err != nil {
		return err
	}

	// Services can't be deleted as a collection...
	list, err := client.Clientset.CoreV1().Services(Namespace).List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, service := range list.Items {
		err = client.Clientset.CoreV1().Services(Namespace).Delete(service.Name, &metav1.DeleteOptions{
			GracePeriodSeconds: &zero,
		})
		if err != nil {
			return err
		}
	}

	return WaitFor(1*time.Second, 20*time.Second, func() bool {
		list, err := client.Clientset.CoreV1().Services(Namespace).List(metav1.ListOptions{})
		if err != nil {
			return false
		}
		return len(list.Items) == 0
	})
}

func (client *cluster) Initialize(networkID string) error {
	log.Println("Initializing cluster")
	client.NetworkID = networkID

	env, err := client.Backend.DockerEnv()
	if err != nil {
		return err
	}
	builder := NewDockerBuilder(env)
	builder.Build("kowalatech/bootnode:dev", "bootnode.Dockerfile")
	builder.Build("kowalatech/kusd:dev", "kcoin.Dockerfile")

	if err := client.createNamespace(); err != nil {
		return err
	}
	if err := client.StoreString(NetworkIDKey, networkID); err != nil {
		return err
	}
	if err := client.addKeys(); err != nil {
		return err
	}
	if err := client.addKeysPassword(); err != nil {
		return err
	}
	if err := client.generateGenesis(); err != nil {
		return err
	}
	if errs := builder.Wait(); len(errs) > 0 {
		return errs[0] // any error will do
	}

	return nil
}

func (client *cluster) DeletePod(podName string) error {
	return client.Clientset.CoreV1().Pods(Namespace).Delete(podName, &metav1.DeleteOptions{})
}

func (client *cluster) createNamespace() error {
	ns, err := client.Clientset.CoreV1().Namespaces().Get(Namespace, metav1.GetOptions{})

	// `err`` will be a NotFound if the namespace doesn't exist, and `ns` will be
	//   a struct with an empty Name. Just checking for the name match will cover
	//   a not found too.
	if ns.Name == Namespace {
		return nil
	}

	log.Printf("Creating `%v` namespace...\n", Namespace)
	ns = &apiv1.Namespace{}
	ns.Name = Namespace
	_, err = client.Clientset.CoreV1().Namespaces().Create(ns)
	return err
}

func (client *cluster) waitForPod(podName string) error {
	log.Printf("Waiting for pod `%v`...\n", podName)
	return WaitFor(2*time.Second, 1*time.Minute, func() bool {
		return client.isPodRunning(podName)
	})
}

func (client *cluster) isPodRunning(podName string) bool {
	pod, err := client.Clientset.CoreV1().Pods(Namespace).Get(podName, metav1.GetOptions{})
	if err != nil {
		return false
	}
	return pod != nil && pod.Status.Phase == apiv1.PodRunning
}

func (client *cluster) waitForKusdPod(podName string) error {
	log.Printf("Waiting for pod `%v`...\n", podName)
	return WaitFor(2*time.Second, 2*time.Minute, func() bool {
		return client.isKusdPodRunning(podName)
	})
}

func (client *cluster) GetBalance(podName string) (*big.Int, error) {
	resp, err := client.Exec(podName, fmt.Sprintf(`eth.getBalance(eth.coinbase)`))
	if err != nil {
		return nil, err
	}
	balance := big.NewInt(0)
	balance.SetString(strings.TrimSpace(resp.StdOut), 10)
	return balance, nil
}

func (client *cluster) isKusdPodRunning(podName string) bool {
	pod, err := client.Clientset.CoreV1().Pods(Namespace).Get(podName, metav1.GetOptions{})
	if pod == nil || err != nil {
		return false
	}
	if pod.Status.Phase != apiv1.PodRunning {
		return false
	}
	// Run any command just to check if it works
	_, err = client.Exec(podName, `console.log("Hello world")`)
	return err == nil
}

func (client *cluster) waitForInitialSync(podName string) error {
	log.Printf("Waiting for pod `%v` to finish initial sync...\n", podName)
	return WaitFor(2*time.Second, 5*time.Minute, func() bool {
		resp, err := client.Exec(podName, `eth.syncing`)
		return err == nil && resp.StdOut == "false\n"
	})
}
