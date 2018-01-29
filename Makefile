# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: kusd android ios kusd-cross swarm evm all test clean
.PHONY: kusd-linux kusd-linux-386 kusd-linux-amd64 kusd-linux-mips64 kusd-linux-mips64le
.PHONY: kusd-linux-arm kusd-linux-arm-5 kusd-linux-arm-6 kusd-linux-arm-7 kusd-linux-arm64
.PHONY: kusd-darwin kusd-darwin-386 kusd-darwin-amd64
.PHONY: kusd-windows kusd-windows-386 kusd-windows-amd64

GOBIN = $(pwd)/build/bin
GO ?= latest

kusd:
	build/env.sh go run build/ci.go install ./cmd/kusd
	@echo "Done building."
	@echo "Run \"$(GOBIN)/kusd\" to launch kusd."

bootnode:
	build/env.sh go run build/ci.go install ./cmd/bootnode
	@echo "Done building."
	@echo "Run \"$(GOBIN)/bootnode\" to launch bootnode."

swarm:
	build/env.sh go run build/ci.go install ./cmd/swarm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/swarm\" to launch swarm."

evm:
	build/env.sh go run build/ci.go install ./cmd/evm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/evm\" to start the evm."

all:
	build/env.sh go run build/ci.go install

android:
	build/env.sh go run build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/kusd.aar\" to use the library."

ios:
	build/env.sh go run build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/Kusd.framework\" to use the library."

test: all
	build/env.sh go run build/ci.go test

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

kusd-cross: kusd-linux kusd-darwin kusd-windows kusd-android kusd-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/kusd-*

kusd-linux: kusd-linux-386 kusd-linux-amd64 kusd-linux-arm kusd-linux-mips64 kusd-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/kusd-linux-*

kusd-linux-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/kusd
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/kusd-linux-* | grep 386

kusd-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/kusd
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/kusd-linux-* | grep amd64

kusd-linux-arm: kusd-linux-arm-5 kusd-linux-arm-6 kusd-linux-arm-7 kusd-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/kusd-linux-* | grep arm

kusd-linux-arm-5:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/kusd
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/kusd-linux-* | grep arm-5

kusd-linux-arm-6:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/kusd
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/kusd-linux-* | grep arm-6

kusd-linux-arm-7:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/kusd
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/kusd-linux-* | grep arm-7

kusd-linux-arm64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/kusd
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/kusd-linux-* | grep arm64

kusd-linux-mips:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/kusd
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/kusd-linux-* | grep mips

kusd-linux-mipsle:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/kusd
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/kusd-linux-* | grep mipsle

kusd-linux-mips64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/kusd
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/kusd-linux-* | grep mips64

kusd-linux-mips64le:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/kusd
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/kusd-linux-* | grep mips64le

kusd-darwin: kusd-darwin-386 kusd-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/kusd-darwin-*

kusd-darwin-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/kusd
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/kusd-darwin-* | grep 386

kusd-darwin-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/kusd
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/kusd-darwin-* | grep amd64

kusd-windows: kusd-windows-386 kusd-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/kusd-windows-*

kusd-windows-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/kusd
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/kusd-windows-* | grep 386

kusd-windows-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/kusd
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/kusd-windows-* | grep amd64


## Docker

docker-build-bootnode:
	docker build -t kowala-tech/bootnode -f bootnode.Dockerfile .

docker-build-kusd:
	docker build -t kowala-tech/kusd -f kusd.Dockerfile .
