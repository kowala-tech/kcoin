package cluster

// RunNode Runs the a node
func (client *cluster) RunNode(name string) error {
	pod, err := NewPodBuilder().
		Network("testnet").
		Name("validator").
		Bootnode(client.NetworkID).
		Build()
	if err != nil {
		return nil
	}

	_, err = client.Clientset.CoreV1().Pods(Namespace).Create(pod)
	if err != nil {
		return err
	}

	return client.waitForKusdPod(name)
}
