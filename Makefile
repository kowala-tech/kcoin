# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: kcoin android ios kcoin-cross swarm evm genesis all test clean
.PHONY: kcoin-cross kcoin-cross-compress kcoin-cross-build  kcoin-cross-rename
.PHONY: dep e2e
.PHONY: dev_docker_images dev_kusd_docker_image dev_bootnode_docker_image

GOBIN = $(pwd)/client/build/bin
GO ?= latest

NPROCS := 1
OS := $(shell uname)
ifeq ($(OS),Linux)
	NPROCS := $(shell grep -c ^processor /proc/cpuinfo)
else ifeq ($(OS),Darwin)
	NPROCS := $(shell sysctl -n hw.ncpu)
endif # $(OS)

kcoin: generate-contract-bindings
	cd client; build/env.sh go run build/ci.go install ./cmd/kcoin
	@echo "Done building."
	@echo "Run \"$(GOBIN)/kcoin\" to launch kcoin."

control:
	cd client; build/env.sh go run build/ci.go install ./cmd/control
	@echo "Done building."
	@echo "Run \"$(GOBIN)/control\" to launch control."

bootnode:
	cd client; build/env.sh go run build/ci.go install ./cmd/bootnode
	@echo "Done building."
	@echo "Run \"$(GOBIN)/bootnode\" to launch bootnode."

faucet:
	cd client; build/env.sh go run build/ci.go install ./cmd/faucet
	@echo "Done building."
	@echo "Run \"$(GOBIN)/faucet\" to launch faucet."

evm:
	cd client; build/env.sh go run build/ci.go install ./cmd/evm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/evm\" to start the evm."

genesis:
	cd client; build/env.sh go run build/ci.go install ./cmd/genesis
	@echo "Done building."
	@echo "Run \"$(GOBIN)/genesis\" to generate genesis files."

all:
	cd client; build/env.sh go run build/ci.go install

android:
	cd client; build/env.sh go run build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/kcoin.aar\" to use the library."

ios:
	cd client; build/env.sh go run build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/Kusd.framework\" to use the library."

test: all
	cd client; build/env.sh go run build/ci.go test

test_notifications: dep
	cd notifications && \
	$(GOPATH)/bin/dep ensure --vendor-only && \
	go test ./... -tags=integration

lint: all
	cd client; build/env.sh go run build/ci.go lint

clean:
	rm -fr client/build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

generate-contract-bindings: devtools
	cd client/bindings/consensus; go generate
	cd client/bindings/oracle; go generate
	cd client/bindings/ownership; go generate

devtools:
	env GOBIN= go get -u golang.org/x/tools/cmd/stringer
	env GOBIN= go get -u github.com/jteeuwen/go-bindata/go-bindata
	env GOBIN= go get -u github.com/fjl/gencodec
	env GOBIN= go install ./client/cmd/abigen

# Cross Compilation Targets (xgo)

kcoin-cross: kcoin-cross-build kcoin-cross-compress kcoin-cross-rename
	@echo "Full cross compilation done."

kcoin-cross-build:
	cd client; build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64,linux/arm64,darwin/amd64,windows/amd64 -v ./cmd/kcoin
	mv client/build/bin/kcoin-darwin-10.6-amd64 client/build/bin/kcoin-osx-10.6-amd64

kcoin-cross-compress:
	cd client/build/bin; for f in kcoin*; do zip $$f.zip $$f; rm $$f; done; cd -

kcoin-cross-rename:
ifdef DRONE_TAG
	cd client/build/bin && for f in kcoin-*; do \
		release=$$(echo $$f | awk '{ gsub("kcoin", "kcoin-stable"); print }');\
		version=$$(echo $$f | awk '{ gsub("kcoin", "kcoin-$(DRONE_TAG)"); print }');\
		cp $$f $$release;\
		mv $$f $$version;\
	done;
else
	cd client/build/bin && for f in kcoin-*; do \
		release=$$(echo $$f | awk '{ gsub("kcoin", "kcoin-unstable"); print }');\
		version=$$(echo $$f | awk '{ gsub("kcoin", "kcoin-$(DRONE_COMMIT_SHA)"); print }');\
		cp $$f $$release;\
		mv $$f $$version;\
	done;
endif

## E2E tests

e2e: dep
	cd e2e && \
	$(GOPATH)/bin/dep ensure --vendor-only && \
	go build -a && \
	./e2e --features ./features


## Dev docker images

dev_docker_images: dev_kusd_docker_image dev_bootnode_docker_image

dev_kusd_docker_image:
	docker build -t kowalatech/kusd:dev -f client/release/kcoin.Dockerfile .

dev_bootnode_docker_image:
	docker build -t kowalatech/bootnode:dev -f client/release/bootnode.Dockerfile .

# Tools

DEP_BIN := $(shell command -v dep 2> /dev/null)
dep:
ifndef DEP_BIN
	@echo "Installing dep..."
	@go get github.com/golang/dep/cmd/dep
endif
