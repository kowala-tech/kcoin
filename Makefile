# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

PWD   := $(shell pwd)
GOBIN = $(PWD)/client/build/bin

.PHONY: kcoin
kcoin:
	cd client; build/env.sh go run build/ci.go install ./cmd/kcoin
	@echo "Done building."
	@echo "Run \"$(GOBIN)/kcoin\" to launch kcoin."

.PHONY: control
control:
	cd client; build/env.sh go run build/ci.go install ./cmd/control
	@echo "Done building."
	@echo "Run \"$(GOBIN)/control\" to launch control."

.PHONY: bootnode
bootnode:
	cd client; build/env.sh go run build/ci.go install ./cmd/bootnode
	@echo "Done building."
	@echo "Run \"$(GOBIN)/bootnode\" to launch bootnode."

.PHONY: faucet
faucet:
	cd client; build/env.sh go run build/ci.go install ./cmd/faucet
	@echo "Done building."
	@echo "Run \"$(GOBIN)/faucet\" to launch faucet."

.PHONY: evm
evm:
	cd client; build/env.sh go run build/ci.go install ./cmd/evm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/evm\" to start the evm."

.PHONY: genesis
genesis:
	cd client; build/env.sh go run build/ci.go install ./cmd/genesis
	@echo "Done building."
	@echo "Run \"$(GOBIN)/genesis\" to generate genesis files."

.PHONY: all
all:
	cd client; build/env.sh go run build/ci.go install

.PHONY: android
android:
	cd client; build/env.sh go run build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/kcoin.aar\" to use the library."

.PHONY: ios
ios:
	cd client; build/env.sh go run build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/Kusd.framework\" to use the library."

.PHONY: abigen
abigen:
	cd client; build/env.sh go run build/ci.go install ./cmd/abigen

.PHONY: test
test: all
	cd client; build/env.sh go run build/ci.go test

.PHONY: test_genesis
test_genesis: all
	cd client/knode/genesis && go test . || curl -X POST -H 'Content-type: application/json' --data '{"attachments":[{"actions":[{"type":"button","text":"Build link","url":"${DRONE_BUILD_LINK}"}],"title":"Build failure","pretext":"Some days, you just cant get rid of a bomb!","text":"*Network config has changed!* Regenerate golden files, this will also require a network restart.","mrkdwn_in":["text","pretext","title"],"color":"#ff0000"}]}' ${SLACK_APP_WEBHOOK}

.PHONY: test_notifications
test_notifications: dep
	cd notifications && \
	$(GOPATH)/bin/dep ensure --vendor-only && \
	go test ./... -tags=integration

.PHONY: test_truffle
test_truffle: client/contracts/truffle/node_modules
	cd client/contracts/truffle; npm run test

.PHONY: lint
lint: all
	cd client; build/env.sh go run build/ci.go lint

.PHONY: clean
clean:
	rm -fr client/build/_workspace/pkg/ $(GOBIN)/* client/build/bin/abigen client/contracts/truffle/node_modules

# Bindings tools

.PHONY: bindings
bindings:
	$(MAKE) -j 5 stringer go-bindata gencodec abigen bindings_node_modules
	go generate ./client/contracts/bindings/...

.PHONY: install_tools
install_tools: notifications_dep wallet_backend_dep abigen moq go-bindata stringer gencodec mockery protoc-gen-go stringer go-bindata gencodec

client/contracts/truffle/node_modules:
	cd client/contracts/truffle; npm ci

.PHONY: go_generate
go_generate: client/contracts/truffle/node_modules
	# force namehash first because the other contracts depend on this libraries.
	go generate ./client/contracts/bindings/utils/namehash.go
	go generate ./...

.PHONY: docker_go_generate
docker_go_generate:
	docker run --rm -v $(PWD):/go/src/github.com/kowala-tech/kcoin -w /go/src/github.com/kowala-tech/kcoin kowalatech/go:1.0.12 make go_generate

.PHONY: assert_no_changes
assert_no_changes:
	git status
	@if ! git diff-index --quiet HEAD; then echo 'There are uncommited go generate files.\nRun `make docker_go_generate` to regenerate all of them.'; exit 1; fi

.PHONY: notifications_dep
notifications_dep: dep
	cd notifications && \
	$(GOPATH)/bin/dep ensure --vendor-only

.PHONY: wallet_backend_dep
wallet_backend_dep: dep
	cd wallet-backend && \
	$(GOPATH)/bin/dep ensure --vendor-only

# Cross Compilation Targets (xgo)

.PHONY: go_repository_index_update
go_repository_index_update:
	@cd client; go run cmd/repository/main.go

.PHONY: repository_index
repository_index:
	@echo "generating index files"
	@aws s3 ls releases.kowala.io | cut -b32- - > index.txt

.PHONY: repository_index_update
repository_index_update: repository_index
	@echo "uploading index files"
	@aws s3 cp index.txt s3://releases.kowala.io --acl public-read

.PHONY: kcoin_cross
kcoin_cross: kcoin_cross_build kcoin_cross_compress kcoin_cross_rename
	@echo "Full cross compilation done."

.PHONY: kcoin_cross_build
kcoin_cross_build:
	cd client; build/env.sh go run build/ci.go xgo -- --go=latest --targets=linux/amd64,linux/arm64,darwin/amd64,windows/amd64 -v ./cmd/kcoin
	mv client/build/bin/kcoin-darwin-10.6-amd64 client/build/bin/kcoin-darwin-amd64
	mv client/build/bin/kcoin-windows-4.0-amd64.exe client/build/bin/kcoin-windows-amd64.exe

.PHONY: kcoin_cross_compress
kcoin_cross_compress:
	cd client/build/bin; for f in kcoin*; do zip $$f.zip $$f; rm $$f; done; cd -

.PHONY: kcoin_cross_rename
kcoin_cross_rename:
ifdef DRONE_TAG
	mkdir -p client/build/bin/tags/$(DRONE_TAG)
	cd client/build/bin && for f in kcoin-*; do \
		release=$$(echo $$f | awk '{ gsub("kcoin", "kcoin-stable"); print }');\
		version=$$(echo $$f | awk '{ gsub("kcoin", "tags/$(DRONE_TAG)/kcoin"); print }');\
		cp $$f $$release;\
		mv $$f $$version;\
	done;
else
	mkdir -p client/build/bin/commits/$(DRONE_COMMIT_SHA)
	cd client/build/bin && for f in kcoin-*; do \
		release=$$(echo $$f | awk '{ gsub("kcoin", "kcoin-unstable"); print }');\
		version=$$(echo $$f | awk '{ gsub("kcoin", "commits/$(DRONE_COMMIT_SHA)/kcoin"); print }');\
		cp $$f $$release;\
		mv $$f $$version;\
	done;
endif

## E2E tests

.PHONY: e2e
e2e: dep
	cd e2e && \
	$(GOPATH)/bin/dep ensure --vendor-only && \
	go build -a && \
	./e2e --features ./features

## Wallet app

.PHONY: wallet_app_tests
wallet_app_tests:
	@cd wallet-app; \
	yarn install --network-concurrency 1 && \
	yarn run test

## Docs
BUILD_DOCS := mkdocs build --clean --strict -d site
.PHONY: build_docs
build_docs:
	@cd docs; $(BUILD_DOCS)

.PHONY: build_docs_with_docker
build_docs_with_docker:
	@docker run --rm -v $(PWD)/docs:/documents kowalatech/mkdocs $(BUILD_DOCS)

## Dev docker images

.PHONY: dev_docker_images
dev_docker_images: dev_explorer_docker_image dev_explorer_sync_docker_image dev_kusd_docker_image dev_bootnode_docker_image dev_faucet_docker_image dev_wallet_backend_docker_image dev_transactions_persistance_docker_image dev_transactions_publisher_docker_image dev_backend_api_docker_image

.PHONY: dev_kusd_docker_image
dev_kusd_docker_image:
	docker build --build-arg CI --build-arg DRONE --build-arg DRONE_REPO --build-arg DRONE_COMMIT_SHA --build-arg DRONE_COMMIT_BRANCH --build-arg DRONE_TAG --build-arg DRONE_BUILD_NUMBER --build-arg DRONE_BUILD_EVENT -t kowalatech/kusd:dev -f client/release/kcoin.Dockerfile .

.PHONY: dev_bootnode_docker_image
dev_bootnode_docker_image:
	docker build --build-arg CI --build-arg DRONE --build-arg DRONE_REPO --build-arg DRONE_COMMIT_SHA --build-arg DRONE_COMMIT_BRANCH --build-arg DRONE_TAG --build-arg DRONE_BUILD_NUMBER --build-arg DRONE_BUILD_EVENT -t kowalatech/bootnode:dev -f client/release/bootnode.Dockerfile .

.PHONY: dev_faucet_docker_image
dev_faucet_docker_image:
	docker build --build-arg CI --build-arg DRONE --build-arg DRONE_REPO --build-arg DRONE_COMMIT_SHA --build-arg DRONE_COMMIT_BRANCH --build-arg DRONE_TAG --build-arg DRONE_BUILD_NUMBER --build-arg DRONE_BUILD_EVENT -t kowalatech/faucet:dev -f client/release/faucet.Dockerfile .

.PHONY: dev_wallet_backend_docker_image
dev_wallet_backend_docker_image:
	docker build --build-arg CI --build-arg DRONE --build-arg DRONE_REPO --build-arg DRONE_COMMIT_SHA --build-arg DRONE_COMMIT_BRANCH --build-arg DRONE_TAG --build-arg DRONE_BUILD_NUMBER --build-arg DRONE_BUILD_EVENT -t kowalatech/wallet_backend:dev -f wallet-backend/Dockerfile .

.PHONY: dev_transactions_persistance_docker_image
dev_transactions_persistance_docker_image:
	docker build --build-arg CI --build-arg DRONE --build-arg DRONE_REPO --build-arg DRONE_COMMIT_SHA --build-arg DRONE_COMMIT_BRANCH --build-arg DRONE_TAG --build-arg DRONE_BUILD_NUMBER --build-arg DRONE_BUILD_EVENT -t kowalatech/transactions_persistance:dev -f notifications/transactions_db_synchronize.Dockerfile .

.PHONY: dev_transactions_publisher_docker_image
dev_transactions_publisher_docker_image:
	docker build --build-arg CI --build-arg DRONE --build-arg DRONE_REPO --build-arg DRONE_COMMIT_SHA --build-arg DRONE_COMMIT_BRANCH --build-arg DRONE_TAG --build-arg DRONE_BUILD_NUMBER --build-arg DRONE_BUILD_EVENT -t kowalatech/transactions_publisher:dev -f notifications/transactions_publisher.Dockerfile .

.PHONY: dev_backend_api_docker_image
dev_backend_api_docker_image:
	docker build --build-arg CI --build-arg DRONE --build-arg DRONE_REPO --build-arg DRONE_COMMIT_SHA --build-arg DRONE_COMMIT_BRANCH --build-arg DRONE_TAG --build-arg DRONE_BUILD_NUMBER --build-arg DRONE_BUILD_EVENT -t kowalatech/backend_api:dev -f notifications/api.Dockerfile .

.PHONY: dev_explorer_docker_image
dev_explorer_docker_image:
	docker build --build-arg CI --build-arg DRONE --build-arg DRONE_REPO --build-arg DRONE_COMMIT_SHA --build-arg DRONE_COMMIT_BRANCH --build-arg DRONE_TAG --build-arg DRONE_BUILD_NUMBER --build-arg DRONE_BUILD_EVENT -t kowalatech/kexplorer -f explorer/web.Dockerfile .

.PHONY: dev_explorer_sync_docker_image
dev_explorer_sync_docker_image:
	docker build -t kowalatech/kexplorersync -f explorer/sync.Dockerfile .

# Tools

.PHONY: dep
DEP_BIN := $(shell command -v dep 2> /dev/null)
dep:
ifndef DEP_BIN
	@echo "Installing dep..."
	@go get github.com/golang/dep/cmd/dep
endif

.PHONY: stringer
STRINGER_BIN := $(shell command -v stringer 2> /dev/null)
stringer:
ifndef STRINGER_BIN
	@echo "Installing stringer..."
	@go get golang.org/x/tools/cmd/stringer
endif

.PHONY: go-bindata
GO_BINDATA_BIN := $(shell command -v go-bindata 2> /dev/null)
go-bindata:
ifndef GO_BINDATA_BIN
	@echo "Installing go-bindata..."
	@go get -u github.com/kevinburke/go-bindata/...
endif

.PHONY: gencodec
GENCODEC_BIN := $(shell command -v gencodec 2> /dev/null)
gencodec:
ifndef GENCODEC_BIN
	@echo "Installing gencodec..."
	@go get github.com/fjl/gencodec
endif

MOQ_BIN := $(shell command -v moq 2> /dev/null)
moq:
ifndef MOQ_BIN
	@echo "Installing moq..."
	@go get github.com/matryer/moq
endif

MOCKERY_BIN := $(shell command -v mockery 2> /dev/null)
mockery:
ifndef MOCKERY_BIN
	@echo "Installing mockery..."
	@go get github.com/vektra/mockery/.../
endif

PROTOC_GEN_BIN := $(shell command -v protoc-gen-go 2> /dev/null)
protoc-gen-go:
ifndef PROTOC_GEN_BIN
	@echo "Installing protoc-gen-go..."
	@go get github.com/golang/protobuf/protoc-gen-go
endif
