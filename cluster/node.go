package cluster

// RunNode Runs the a node
func (client *cluster) RunNode(name string) error {
	bootnode, err := client.GetString(bootnodeEnodeKey)
	if err != nil {
		return err
	}

	pod, err := NewPodBuilder().
		WithNetworkId(client.NetworkID).
		WithName(name).
		WithBootnode(bootnode).
		WithSyncMode("full").
		WithLogLevel(3).
		Build()
	if err != nil {
		return nil
	}

	_, err = client.Clientset.CoreV1().Pods(Namespace).Create(pod)
	if err != nil {
		return err
	}

	err = client.waitForKusdPod(name)
	if err != nil {
		return err
	}

	return client.waitForInitialSync(name)
}
