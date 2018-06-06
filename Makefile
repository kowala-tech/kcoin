# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: kcoin android ios kcoin-cross swarm evm genesis all test clean
.PHONY: kcoin-cross kcoin-cross-compress kcoin-cross-build  kcoin-cross-rename

GOBIN = $(pwd)/build/bin
GO ?= latest

NPROCS := 1
OS := $(shell uname)
ifeq ($(OS),Linux)
	NPROCS := $(shell grep -c ^processor /proc/cpuinfo)
else ifeq ($(OS),Darwin)
	NPROCS := $(shell sysctl -n hw.ncpu)
endif # $(OS)

kcoin:
	build/env.sh go run build/ci.go install ./cmd/kcoin
	@echo "Done building."
	@echo "Run \"$(GOBIN)/kcoin\" to launch kcoin."

control:
	build/env.sh go run build/ci.go install ./cmd/control
	@echo "Done building."
	@echo "Run \"$(GOBIN)/control\" to launch control."

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

genesis:
	build/env.sh go run build/ci.go install ./cmd/genesis
	@echo "Done building."
	@echo "Run \"$(GOBIN)/genesis\" to generate genesis files."

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

kcoin-cross: kcoin-cross-build kcoin-cross-compress kcoin-cross-rename
	@echo "Full cross compilation done."

kcoin-cross-build:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64,linux/arm64,darwin/amd64,windows/amd64 -v ./cmd/kcoin
	mv build/bin/kcoin-darwin-10.6-amd64 build/bin/kcoin-osx-10.6-amd64

kcoin-cross-compress:
	cd build/bin; for f in kcoin*; do zip $$f.zip $$f; rm $$f; done; cd -

kcoin-cross-rename:
ifdef DRONE_TAG
	cd build/bin && for f in kcoin*; do cp $$f $${f/kcoin-/kcoin-stable-}; mv $$f $${f/kcoin-/kcoin-$(DRONE_TAG)-}; done; cd -
else
	cd build/bin && for f in kcoin*; do cp $$f $${f/kcoin-/kcoin-unstable-}; mv $$f $${f/kcoin-/kcoin-$(DRONE_COMMIT_SHA)-}; done; cd -
endif

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

GODOG_BIN := $(shell command -v godog 2> /dev/null)

e2e:
ifndef GODOG_BIN
	@echo "Installing godog..."
	@go get github.com/DATA-DOG/godog/cmd/godog
endif
	@build/env.sh sh -c "cd tests && godog -c=$(NPROCS) -f=progress ../features"
