# Makefile
PROJECT_ROOT ?= ../live-server-rpc/app
SERVICES ?= game user
# 配置项目根目录和第三方依赖目录
PROTO_ROOT ?= ./proto
PROTO_THIRD_PARTY ?= $(PROTO_ROOT)/third_party
GENERATED_DIR ?= ./proto-gen-go
# 使用 wildcard 和 filter 组合查找所有 .proto 文件
PROTO_FILES := $(shell find $(PROTO_ROOT) -name "*.proto")

# 从 .proto 文件路径中提取服务目录
PROTO_SERVICES := $(sort $(dir $(PROTO_FILES)))
# RPC 服务生成配置
define generate_rpc
echo "Generating RPC service: $(1)"
bash ./scripts/genrpc.sh $(1)
endef

# 生成所有 RPC 服务
.PHONY: genrpc
genrpc: check-tools
	@$(foreach service,$(SERVICES),$(call generate_rpc,$(service));)

# 检查必要工具是否安装
.PHONY: check-tools
check-tools:
	@which protoc >/dev/null || (echo "Error: protoc not found" && exit 1)
	@which protoc-gen-go >/dev/null || (echo "Error: protoc-gen-go not found" && exit 1)
	@which protoc-gen-go-grpc >/dev/null || (echo "Error: protoc-gen-go-grpc not found" && exit 1)
	@which goctl >/dev/null || (echo "Error: goctl not found" && exit 1)


# 动态创建目录结构（仅为有 .proto 文件的目录创建）
prepare:
	@$(foreach service,$(PROTO_SERVICES), \
		mkdir -p $(GENERATED_DIR)/$(subst $(PROTO_ROOT)/,,$(patsubst %/,%,$(service))); \
	)

# 生成所有服务的 Go 代码
gengo: prepare
	@if [ -z "$(PROTO_FILES)" ]; then \
		echo "没有找到 .proto 文件"; \
	else \
		for service in $(PROTO_SERVICES); do \
			protoc \
			--proto_path=$(PROTO_ROOT) \
			--proto_path=$(PROTO_THIRD_PARTY) \
			--go_out=$(GENERATED_DIR) \
			--go_opt=paths=source_relative \
			--go-grpc_out=$(GENERATED_DIR) \
			--go-grpc_opt=paths=source_relative \
			$$service/*.proto; \
		done \
	fi


# 帮助信息
.PHONY: help
help:
	@echo "使用说明:"
	@echo "  make genrpc     - 生成所有配置的 RPC 服务"
	@echo "  make install-tools - 安装所需工具"
	@echo "  make clean      - 清理生成的文件"
	@echo ""
	@echo "可以通过环境变量自定义:"
	@echo "  SERVICES=\"service1 service2 ...\"  - 指定要生成的服务列表 (默认: game user)"
	@echo "  PROJECT_ROOT    - RPC 服务的基础目录 (默认: ../live-server-rpc/app)"

# 清理生成的文件
.PHONY: clean
clean:
	@echo "清理生成的文件..."
	@find $(PROJECT_ROOT) -type d -name "pb" -exec rm -rf {} +

# 安装依赖工具
.PHONY: install-tools
install-tools:
	@echo "Installing required tools..."
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/zeromicro/go-zero/tools/goctl@latest