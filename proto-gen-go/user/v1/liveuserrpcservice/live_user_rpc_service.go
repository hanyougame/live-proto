// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: user.proto

package liveuserrpcservice

import (
	"context"

	"github.com/hanyougame/live-proto/proto-gen-go/user/v1"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetInfoByUserTokenReq = v1.GetInfoByUserTokenReq
	GetUserBalanceReply   = v1.GetUserBalanceReply
	GetUserBalanceReq     = v1.GetUserBalanceReq
	GetUserInfoByIdReq    = v1.GetUserInfoByIdReq
	UpdateUserBalanceReq  = v1.UpdateUserBalanceReq
	UpdateUserBalanceResp = v1.UpdateUserBalanceResp
	UserDetailsInfoReply  = v1.UserDetailsInfoReply
	UserWalletInfo        = v1.UserWalletInfo

	LiveUserRpcService interface {
		// 根据token获取用户信息
		GetInfoByUserToken(ctx context.Context, in *GetInfoByUserTokenReq, opts ...grpc.CallOption) (*UserDetailsInfoReply, error)
		// 获取用户余额
		GetUserBalance(ctx context.Context, in *GetUserBalanceReq, opts ...grpc.CallOption) (*GetUserBalanceReply, error)
		// 修改用户余额
		UpdateUserBalance(ctx context.Context, in *UpdateUserBalanceReq, opts ...grpc.CallOption) (*UpdateUserBalanceResp, error)
		GetUserInfoById(ctx context.Context, in *GetUserInfoByIdReq, opts ...grpc.CallOption) (*UserDetailsInfoReply, error)
	}

	defaultLiveUserRpcService struct {
		cli zrpc.Client
	}
)

func NewLiveUserRpcService(cli zrpc.Client) LiveUserRpcService {
	return &defaultLiveUserRpcService{
		cli: cli,
	}
}

// 根据token获取用户信息
func (m *defaultLiveUserRpcService) GetInfoByUserToken(ctx context.Context, in *GetInfoByUserTokenReq, opts ...grpc.CallOption) (*UserDetailsInfoReply, error) {
	client := v1.NewLiveUserRpcServiceClient(m.cli.Conn())
	return client.GetInfoByUserToken(ctx, in, opts...)
}

// 获取用户余额
func (m *defaultLiveUserRpcService) GetUserBalance(ctx context.Context, in *GetUserBalanceReq, opts ...grpc.CallOption) (*GetUserBalanceReply, error) {
	client := v1.NewLiveUserRpcServiceClient(m.cli.Conn())
	return client.GetUserBalance(ctx, in, opts...)
}

// 修改用户余额
func (m *defaultLiveUserRpcService) UpdateUserBalance(ctx context.Context, in *UpdateUserBalanceReq, opts ...grpc.CallOption) (*UpdateUserBalanceResp, error) {
	client := v1.NewLiveUserRpcServiceClient(m.cli.Conn())
	return client.UpdateUserBalance(ctx, in, opts...)
}

func (m *defaultLiveUserRpcService) GetUserInfoById(ctx context.Context, in *GetUserInfoByIdReq, opts ...grpc.CallOption) (*UserDetailsInfoReply, error) {
	client := v1.NewLiveUserRpcServiceClient(m.cli.Conn())
	return client.GetUserInfoById(ctx, in, opts...)
}
