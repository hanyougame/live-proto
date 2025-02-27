// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: agent.proto

package liveagentrpcservice

import (
	"context"

	"github.com/hanyougame/live-proto/proto-gen-go/agent/v1"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AgentDataCountTaskExecuteReq  = v1.AgentDataCountTaskExecuteReq
	AgentDataCountTaskExecuteResp = v1.AgentDataCountTaskExecuteResp
	AgentSettleTaskExecuteReq     = v1.AgentSettleTaskExecuteReq
	AgentSettleTaskExecuteResp    = v1.AgentSettleTaskExecuteResp

	LiveAgentRpcService interface {
		// 执行代理结算任务
		AgentSettleTaskExecute(ctx context.Context, in *AgentSettleTaskExecuteReq, opts ...grpc.CallOption) (*AgentSettleTaskExecuteResp, error)
		// 执行代理数据统计任务
		AgentDataCountTaskExecute(ctx context.Context, in *AgentDataCountTaskExecuteReq, opts ...grpc.CallOption) (*AgentDataCountTaskExecuteResp, error)
	}

	defaultLiveAgentRpcService struct {
		cli zrpc.Client
	}
)

func NewLiveAgentRpcService(cli zrpc.Client) LiveAgentRpcService {
	return &defaultLiveAgentRpcService{
		cli: cli,
	}
}

// 执行代理结算任务
func (m *defaultLiveAgentRpcService) AgentSettleTaskExecute(ctx context.Context, in *AgentSettleTaskExecuteReq, opts ...grpc.CallOption) (*AgentSettleTaskExecuteResp, error) {
	client := v1.NewLiveAgentRpcServiceClient(m.cli.Conn())
	return client.AgentSettleTaskExecute(ctx, in, opts...)
}

// 执行代理数据统计任务
func (m *defaultLiveAgentRpcService) AgentDataCountTaskExecute(ctx context.Context, in *AgentDataCountTaskExecuteReq, opts ...grpc.CallOption) (*AgentDataCountTaskExecuteResp, error) {
	client := v1.NewLiveAgentRpcServiceClient(m.cli.Conn())
	return client.AgentDataCountTaskExecute(ctx, in, opts...)
}
