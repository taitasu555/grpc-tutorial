# Makefile

# デフォルト設定
PROTO_PATH ?= go/services/command/proto
OUT_DIR    ?= $(PROTO_PATH)/pb

.PHONY: protoCommand
protoCommand:
	@echo "Using PROTO_PATH=$(PROTO_PATH)"
	@echo "Using OUT_DIR=$(OUT_DIR)"
	@echo "Generating Go code from proto files..."
	protoc \
	  --proto_path=$(PROTO_PATH) \
	  --go_out=paths=source_relative:$(OUT_DIR) \
	  --go-grpc_out=paths=source_relative:$(OUT_DIR) \
	  $(PROTO_PATH)/*.proto
	@echo "Generation complete."
