package cluster

import (
	"fmt"
	"log"
	"math/rand"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	namespacePrefix = "kcoin-"
)

func (client *cluster) createNamespace() error {
	if err := client.findNamespace(); err != nil {
		return err
	}

	log.Printf("Creating `%v` namespace...\n", client.Namespace)
	ns := &apiv1.Namespace{}
	ns.Name = client.Namespace
	_, err := client.Clientset.CoreV1().Namespaces().Create(ns)

	return err
}

// findNamespace finds and configures an unused namespace in the cluster
func (client *cluster) findNamespace() error {
	attempts := 10

	for attempts > 0 {
		attempts -= 1
		randomNs := fmt.Sprintf("%v%v", namespacePrefix, rand.Intn(1000000))
		ns, _ := client.Clientset.CoreV1().Namespaces().Get(randomNs, metav1.GetOptions{})
		if ns == nil || ns.Name != randomNs { //Easiest way to check if it's found
			client.Namespace = randomNs
			return nil
		}
	}

	return fmt.Errorf("Tried to find an available namespace but couldn't find any")
}
