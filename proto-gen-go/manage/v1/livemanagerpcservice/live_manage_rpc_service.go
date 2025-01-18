// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: manage.proto

package livemanagerpcservice

import (
	"context"

	"github.com/hanyougame/live-proto/proto-gen-go/manage/v1"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CurrencyDetailReply     = v1.CurrencyDetailReply
	CurrencyDetailReplyList = v1.CurrencyDetailReplyList
	GetCurrInfoByIDReq      = v1.GetCurrInfoByIDReq
	ManageReply             = v1.ManageReply
	ManageReq               = v1.ManageReq

	LiveManageRpcService interface {
		// 获取法定货币信息
		GetLegalTenderInfo(ctx context.Context, in *ManageReq, opts ...grpc.CallOption) (*CurrencyDetailReplyList, error)
		// 通过货币ID获取货币信息
		GetCurrInfoById(ctx context.Context, in *GetCurrInfoByIDReq, opts ...grpc.CallOption) (*CurrencyDetailReply, error)
	}

	defaultLiveManageRpcService struct {
		cli zrpc.Client
	}
)

func NewLiveManageRpcService(cli zrpc.Client) LiveManageRpcService {
	return &defaultLiveManageRpcService{
		cli: cli,
	}
}

// 获取法定货币信息
func (m *defaultLiveManageRpcService) GetLegalTenderInfo(ctx context.Context, in *ManageReq, opts ...grpc.CallOption) (*CurrencyDetailReplyList, error) {
	client := v1.NewLiveManageRpcServiceClient(m.cli.Conn())
	return client.GetLegalTenderInfo(ctx, in, opts...)
}

// 通过货币ID获取货币信息
func (m *defaultLiveManageRpcService) GetCurrInfoById(ctx context.Context, in *GetCurrInfoByIDReq, opts ...grpc.CallOption) (*CurrencyDetailReply, error) {
	client := v1.NewLiveManageRpcServiceClient(m.cli.Conn())
	return client.GetCurrInfoById(ctx, in, opts...)
}
