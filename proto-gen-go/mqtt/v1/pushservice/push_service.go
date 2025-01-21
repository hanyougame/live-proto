// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: mqtt.proto

package pushservice

import (
	"context"

	"github.com/hanyougame/live-proto/proto-gen-go/mqtt/v1"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	EnterQuitGameDelWdlReply = v1.EnterQuitGameDelWdlReply
	PushRequest              = v1.PushRequest
	PushResponse             = v1.PushResponse
	PushToUserRequest        = v1.PushToUserRequest
	PushToUsersRequest       = v1.PushToUsersRequest

	PushService interface {
		PushToAll(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushResponse, error)
		PushToUser(ctx context.Context, in *PushToUserRequest, opts ...grpc.CallOption) (*PushResponse, error)
		PushToUsers(ctx context.Context, in *PushToUsersRequest, opts ...grpc.CallOption) (*PushResponse, error)
	}

	defaultPushService struct {
		cli zrpc.Client
	}
)

func NewPushService(cli zrpc.Client) PushService {
	return &defaultPushService{
		cli: cli,
	}
}

func (m *defaultPushService) PushToAll(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushResponse, error) {
	client := v1.NewPushServiceClient(m.cli.Conn())
	return client.PushToAll(ctx, in, opts...)
}

func (m *defaultPushService) PushToUser(ctx context.Context, in *PushToUserRequest, opts ...grpc.CallOption) (*PushResponse, error) {
	client := v1.NewPushServiceClient(m.cli.Conn())
	return client.PushToUser(ctx, in, opts...)
}

func (m *defaultPushService) PushToUsers(ctx context.Context, in *PushToUsersRequest, opts ...grpc.CallOption) (*PushResponse, error) {
	client := v1.NewPushServiceClient(m.cli.Conn())
	return client.PushToUsers(ctx, in, opts...)
}
