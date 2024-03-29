OS   := $(shell go env GOOS)
ARCH := $(shell go env GOARCH)

BUF_VERSION                := 0.44.0
PROTOC_GEN_GO_VERSION      := 1.27.1
PROTOC_GEN_GO_GRPC_VERSION := 1.1.0

BIN_DIR := $(shell pwd)/bin

BUF                     := $(abspath $(BIN_DIR)/buf)
PROTOC_GEN_GO           := $(abspath $(BIN_DIR)/protoc-gen-go)
PROTOC_GEN_GO_GRPC      := $(abspath $(BIN_DIR)/protoc-gen-go-grpc)
PROTOC_GEN_GRPC_GATEWAY := $(abspath $(BIN_DIR)/protoc-gen-grpc-gateway)

buf: $(BUF)
$(BUF):
	curl -sSL "https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/buf-$(shell uname -s)-$(shell uname -m)" -o $(BUF) && chmod +x $(BUF)

protoc-gen-go: $(PROTOC_GEN_GO)
$(PROTOC_GEN_GO):
	curl -sSL https://github.com/protocolbuffers/protobuf-go/releases/download/v$(PROTOC_GEN_GO_VERSION)/protoc-gen-go.v$(PROTOC_GEN_GO_VERSION).$(OS).$(ARCH).tar.gz | tar -C $(BIN_DIR) -xzv protoc-gen-go

protoc-gen-go-grpc: $(PROTOC_GEN_GO_GRPC)
$(PROTOC_GEN_GO_GRPC):
	curl -sSL https://github.com/grpc/grpc-go/releases/download/cmd%2Fprotoc-gen-go-grpc%2Fv$(PROTOC_GEN_GO_GRPC_VERSION)/protoc-gen-go-grpc.v$(PROTOC_GEN_GO_GRPC_VERSION).$(OS).$(ARCH).tar.gz | tar -C $(BIN_DIR) -xzv ./protoc-gen-go-grpc

protoc-gen-grpc-gateway: $(PROTOC_GEN_GRPC_GATEWAY)
$(PROTOC_GEN_GRPC_GATEWAY):
	cd ./tools && go build -o ../bin/protoc-gen-grpc-gateway github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway

gen-proto: $(BUF) $(PROTOC_GEN_GO) $(PROTOC_GEN_GO_GRPC) $(PROTOC_GEN_GRPC_GATEWAY)
	$(BUF) generate \
		--path ./services/

clean:
	rm -f $(BIN_DIR)/*

gen-deploy-workflow-template:
	chmod +x build/workflows/deploy-workflow.sh && build/workflows/deploy-workflow.sh

gen-pull-request-workflow-template:
	chmod +x build/workflows/pull-request-workflow.sh && build/workflows/pull-request-workflow.sh
