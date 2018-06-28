SOURCE_DIR         := .
VENDOR_DIR         := $(SOURCE_DIR)/vendor
BIN_DIR            := bin
DEP_BIN            := $(BIN_DIR)/dep

.PHONY: all ci setup test integration

all: setup test
ci: setup integration

setup: $(VENDOR_DIR)

build_transactions_publisher:
	docker build -t kowalatech/transactions-publisher:latest -f transactions_publisher.Dockerfile .


test:
	@echo "Running unit tests..."
	@go test ./...

integration:
	@echo "Running integraton tests..."
	@go test ./... -tags=integration

$(DEP_BIN):
	@echo "Installing dep..."
	@mkdir -p bin
	@go get -u github.com/golang/dep/cmd/dep

$(VENDOR_DIR): $(DEP_BIN)
	@echo "Installing Go dependencies..."
	@dep ensure -vendor-only

generate_proto:
	protoc core/api/*.proto --go_out=plugins=grpc:.
	protoc blockchain/*.proto --go_out=plugins=grpc:.
