// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: manage.proto

package livemanagerpcservice

import (
	"context"

	"github.com/hanyougame/live-proto/proto-gen-go/manage/v1"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BatchGetIPGeolocationReply = v1.BatchGetIPGeolocationReply
	BatchGetIPGeolocationReq   = v1.BatchGetIPGeolocationReq
	CurrencyDetailReply        = v1.CurrencyDetailReply
	CurrencyDetailReplyList    = v1.CurrencyDetailReplyList
	GetCurrInfoByIDReq         = v1.GetCurrInfoByIDReq
	GetIPGeolocationReply      = v1.GetIPGeolocationReply
	GetIPGeolocationReq        = v1.GetIPGeolocationReq
	IPGeolocationInfo          = v1.IPGeolocationInfo
	ManageReply                = v1.ManageReply
	ManageReq                  = v1.ManageReq
	SendEmailReply             = v1.SendEmailReply
	SendEmailReq               = v1.SendEmailReq
	SendSmsReply               = v1.SendSmsReply
	SendSmsReq                 = v1.SendSmsReq
	SmsBalanceReply            = v1.SmsBalanceReply
	SmsBalanceReq              = v1.SmsBalanceReq

	LiveManageRpcService interface {
		// 获取法定货币信息
		GetLegalTenderInfo(ctx context.Context, in *ManageReq, opts ...grpc.CallOption) (*CurrencyDetailReplyList, error)
		// 通过货币ID获取货币信息
		GetCurrInfoById(ctx context.Context, in *GetCurrInfoByIDReq, opts ...grpc.CallOption) (*CurrencyDetailReply, error)
		// =============== IP相关接口 ===============
		GetIPGeolocation(ctx context.Context, in *GetIPGeolocationReq, opts ...grpc.CallOption) (*GetIPGeolocationReply, error)
		// 批量获取IP地理位置信息
		BatchGetIPGeolocation(ctx context.Context, in *BatchGetIPGeolocationReq, opts ...grpc.CallOption) (*BatchGetIPGeolocationReply, error)
		SendSms(ctx context.Context, in *SendSmsReq, opts ...grpc.CallOption) (*SendSmsReply, error)
		SmsBalance(ctx context.Context, in *SmsBalanceReq, opts ...grpc.CallOption) (*SmsBalanceReply, error)
		SendEmail(ctx context.Context, in *SendEmailReq, opts ...grpc.CallOption) (*SendEmailReply, error)
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

// =============== IP相关接口 ===============
func (m *defaultLiveManageRpcService) GetIPGeolocation(ctx context.Context, in *GetIPGeolocationReq, opts ...grpc.CallOption) (*GetIPGeolocationReply, error) {
	client := v1.NewLiveManageRpcServiceClient(m.cli.Conn())
	return client.GetIPGeolocation(ctx, in, opts...)
}

// 批量获取IP地理位置信息
func (m *defaultLiveManageRpcService) BatchGetIPGeolocation(ctx context.Context, in *BatchGetIPGeolocationReq, opts ...grpc.CallOption) (*BatchGetIPGeolocationReply, error) {
	client := v1.NewLiveManageRpcServiceClient(m.cli.Conn())
	return client.BatchGetIPGeolocation(ctx, in, opts...)
}

func (m *defaultLiveManageRpcService) SendSms(ctx context.Context, in *SendSmsReq, opts ...grpc.CallOption) (*SendSmsReply, error) {
	client := v1.NewLiveManageRpcServiceClient(m.cli.Conn())
	return client.SendSms(ctx, in, opts...)
}

func (m *defaultLiveManageRpcService) SmsBalance(ctx context.Context, in *SmsBalanceReq, opts ...grpc.CallOption) (*SmsBalanceReply, error) {
	client := v1.NewLiveManageRpcServiceClient(m.cli.Conn())
	return client.SmsBalance(ctx, in, opts...)
}

func (m *defaultLiveManageRpcService) SendEmail(ctx context.Context, in *SendEmailReq, opts ...grpc.CallOption) (*SendEmailReply, error) {
	client := v1.NewLiveManageRpcServiceClient(m.cli.Conn())
	return client.SendEmail(ctx, in, opts...)
}
