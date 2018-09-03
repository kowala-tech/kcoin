package testnet

//Node represents an interface common to all nodes running on a docker container.
type Node interface {
	Start() error
	Stop() error
	ID() string
}

type node struct {
	nodeSpec     *NodeSpec
	dockerEngine DockerEngine
}

func (n *node) Start() error {
	err := n.dockerEngine.PullImage(n.nodeSpec.Image)
	if err != nil {
		return err
	}

	err = n.dockerEngine.CreateContainer(n.nodeSpec.Image, n.nodeSpec.ID, n.nodeSpec.NetworkID, n.nodeSpec.Cmd, nil, n.nodeSpec.PortMapping)
	if err != nil {
		return err
	}

	err = n.copyFiles()
	if err != nil {
		return err
	}

	err = n.dockerEngine.StartContainer(n.nodeSpec.ID)
	if err != nil {
		return err
	}

	return nil
}

func (n *node) copyFiles() error {
	for path, content := range n.nodeSpec.Files {
		err := n.dockerEngine.CopyToContainer(n.nodeSpec.ID, path, content)
		if err != nil {
			return err
		}
	}

	return nil
}

func (n *node) Stop() error {
	err := n.dockerEngine.StopAndRemoveContainer(n.nodeSpec.ID)
	if err != nil {
		return err
	}

	return nil
}

func (n *node) ID() string {
	return n.nodeSpec.ID
}
