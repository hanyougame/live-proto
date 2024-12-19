#!/bin/bash
# genrpc.sh

# 参数1: 服务名称
if [ $# -lt 1 ]; then
    echo "Usage: $0 <service-name>"
    exit 1
fi

SERVICE=$1
TARGET_DIR="../live-server-rpc/app"

# 检查目标目录是否存在，不存在则创建
mkdir -p ${TARGET_DIR}/${SERVICE}

# 检查proto文件是否存在
if [ ! -f "./proto/${SERVICE}/v1/${SERVICE}.proto" ]; then
    echo "Error: Proto file ./proto/${SERVICE}/v1/${SERVICE}.proto not found"
    exit 1
fi

# 生成rpc服务代码
goctl rpc protoc ./proto/${SERVICE}/v1/${SERVICE}.proto \
    --proto_path=. \
    --go_out=${TARGET_DIR}/${SERVICE} \
    --go-grpc_out=${TARGET_DIR}/${SERVICE} \
    --zrpc_out=${TARGET_DIR}/${SERVICE} \
    --client=true \
    --style=go_zero -m