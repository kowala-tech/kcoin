# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: kcoin android ios kcoin-cross swarm evm all test clean
.PHONY: kcoin-linux kcoin-linux-386 kcoin-linux-amd64 kcoin-linux-mips64 kcoin-linux-mips64le
.PHONY: kcoin-linux-arm kcoin-linux-arm-5 kcoin-linux-arm-6 kcoin-linux-arm-7 kcoin-linux-arm64
.PHONY: kcoin-darwin kcoin-darwin-386 kcoin-darwin-amd64
.PHONY: kcoin-windows kcoin-windows-386 kcoin-windows-amd64

GOBIN = $(pwd)/build/bin
GO ?= latest

kcoin:
	build/env.sh go run build/ci.go install ./cmd/kcoin
	@echo "Done building."
	@echo "Run \"$(GOBIN)/kcoin\" to launch kcoin."

bootnode:
	build/env.sh go run build/ci.go install ./cmd/bootnode
	@echo "Done building."
	@echo "Run \"$(GOBIN)/bootnode\" to launch bootnode."

faucet:
	build/env.sh go run build/ci.go install ./cmd/faucet
	@echo "Done building."
	@echo "Run \"$(GOBIN)/faucet\" to launch faucet."

evm:
	build/env.sh go run build/ci.go install ./cmd/evm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/evm\" to start the evm."

all:
	build/env.sh go run build/ci.go install

android:
	build/env.sh go run build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/kcoin.aar\" to use the library."

ios:
	build/env.sh go run build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/Kusd.framework\" to use the library."

test: all
	build/env.sh go run build/ci.go test

lint: all
	build/env.sh go run build/ci.go lint

clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go get -u golang.org/x/tools/cmd/stringer
	env GOBIN= go get -u github.com/jteeuwen/go-bindata/go-bindata
	env GOBIN= go get -u github.com/fjl/gencodec
	env GOBIN= go install ./cmd/abigen

# Cross Compilation Targets (xgo)

kcoin-cross: kcoin-linux kcoin-darwin kcoin-windows kcoin-android kcoin-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/kcoin-*

kcoin-linux: kcoin-linux-386 kcoin-linux-amd64 kcoin-linux-arm kcoin-linux-mips64 kcoin-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/kcoin-linux-*

kcoin-linux-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/kcoin
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/kcoin-linux-* | grep 386

kcoin-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/kcoin
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/kcoin-linux-* | grep amd64

kcoin-linux-arm: kcoin-linux-arm-5 kcoin-linux-arm-6 kcoin-linux-arm-7 kcoin-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/kcoin-linux-* | grep arm

kcoin-linux-arm-5:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/kcoin
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/kcoin-linux-* | grep arm-5

kcoin-linux-arm-6:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/kcoin
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/kcoin-linux-* | grep arm-6

kcoin-linux-arm-7:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/kcoin
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/kcoin-linux-* | grep arm-7

kcoin-linux-arm64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/kcoin
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/kcoin-linux-* | grep arm64

kcoin-linux-mips:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/kcoin
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/kcoin-linux-* | grep mips

kcoin-linux-mipsle:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/kcoin
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/kcoin-linux-* | grep mipsle

kcoin-linux-mips64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/kcoin
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/kcoin-linux-* | grep mips64

kcoin-linux-mips64le:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/kcoin
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/kcoin-linux-* | grep mips64le

kcoin-darwin: kcoin-darwin-386 kcoin-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/kcoin-darwin-*

kcoin-darwin-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/kcoin
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/kcoin-darwin-* | grep 386

kcoin-darwin-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/kcoin
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/kcoin-darwin-* | grep amd64

kcoin-windows: kcoin-windows-386 kcoin-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/kcoin-windows-*

kcoin-windows-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/kcoin
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/kcoin-windows-* | grep 386

kcoin-windows-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/kcoin
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/kcoin-windows-* | grep amd64


## Docker

docker-build-bootnode:
	docker build -t kowalatech/bootnode -f bootnode.Dockerfile .

docker-build-kusd:
	docker build -t kowalatech/kusd -f kcoin.Dockerfile .

docker-build-faucet:
	docker build -t kowalatech/faucet -f faucet.Dockerfile .


docker-publish-bootnode:
	docker push kowalatech/bootnode

docker-publish-kusd:
	docker push kowalatech/kusd

docker-publish-faucet:
	docker push kowalatech/faucet

## E2E tests

MINIKUBE_BIN := $(shell command -v minikube 2> /dev/null)
start_local_k8s:
ifndef MINIKUBE_BIN
	@echo "You must install minikube first..."
	@exit 1
endif
	@minikube start -p testing --kubernetes-version v1.9.0

GODOG_BIN := $(shell command -v godog 2> /dev/null)
e2e:
ifndef GODOG_BIN
	@echo "Installing godog..."
	@go get github.com/DATA-DOG/godog/cmd/godog
endif
	@build/k8s_env.sh sh -c "cd tests && godog ../features"
