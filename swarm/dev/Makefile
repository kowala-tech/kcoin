.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/kowala-tech/kUSD/cmd/swarm
	go build -o bin/kusd     github.com/kowala-tech/kUSD/cmd/kusd
	go build -o bin/bootnode github.com/kowala-tech/kUSD/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/kowala-tech/kUSD/swarm/...
