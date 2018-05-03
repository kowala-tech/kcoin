package cluster

type Node struct {
	Image string
	Name  string
	Files map[string][]byte
	Cmd   []string
}

func BootnodeNode() (*Node, error) {
	return &Node{
		Image: "kowalatech/bootnode:dev",
		Name:  "bootnode",
		Cmd: []string{
			"--nodekeyhex", randStringBytes(64),
		},
		Files: make(map[string][]byte),
	}, nil
}
