// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: finance.proto

package liveauditrpcservice

import (
	"context"
	"github.com/hanyougame/live-proto/proto-gen-go/finance/v1"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddAudit                      = v1.AddAudit
	AddAuditReq                   = v1.AddAuditReq
	AddAuditResp                  = v1.AddAuditResp
	AddWithdrawLimitReq           = v1.AddWithdrawLimitReq
	AddWithdrawLimitResp          = v1.AddWithdrawLimitResp
	AuditInfo                     = v1.AuditInfo
	BalanceReq                    = v1.BalanceReq
	BalanceResp                   = v1.BalanceResp
	BatchAddAuditReq              = v1.BatchAddAuditReq
	BatchAddAuditResp             = v1.BatchAddAuditResp
	ExchangeRateReq               = v1.ExchangeRateReq
	ExchangeRateResp              = v1.ExchangeRateResp
	GetAuditInfoReq               = v1.GetAuditInfoReq
	GetAuditInfoResp              = v1.GetAuditInfoResp
	GetAuditListReq               = v1.GetAuditListReq
	GetAuditListResp              = v1.GetAuditListResp
	PayInReq                      = v1.PayInReq
	PayInResp                     = v1.PayInResp
	PayInStatusReq                = v1.PayInStatusReq
	PayInStatusResp               = v1.PayInStatusResp
	PayOutReq                     = v1.PayOutReq
	PayOutResp                    = v1.PayOutResp
	PayOutStatusReq               = v1.PayOutStatusReq
	PayOutStatusResp              = v1.PayOutStatusResp
	RechargeReq                   = v1.RechargeReq
	RechargeResp                  = v1.RechargeResp
	UpdateAuditAmountReq          = v1.UpdateAuditAmountReq
	UpdateAuditAmountResp         = v1.UpdateAuditAmountResp
	UpdateWithdrawLimitAmountReq  = v1.UpdateWithdrawLimitAmountReq
	UpdateWithdrawLimitAmountResp = v1.UpdateWithdrawLimitAmountResp
	WithdrawReq                   = v1.WithdrawReq
	WithdrawResp                  = v1.WithdrawResp

	LiveAuditRpcService interface {
		// 添加稽核 传递用户币种和金额
		AddAudit(ctx context.Context, in *AddAuditReq, opts ...grpc.CallOption) (*AddAuditResp, error)
		// 更新稽核金额（打稽核） 传递用户币种和金额
		UpdateAuditAmount(ctx context.Context, in *UpdateAuditAmountReq, opts ...grpc.CallOption) (*UpdateAuditAmountResp, error)
		// 获取用户稽核信息
		GetAuditInfo(ctx context.Context, in *GetAuditInfoReq, opts ...grpc.CallOption) (*GetAuditInfoResp, error)
		// 获取指定用户稽核列表
		GetAuditList(ctx context.Context, in *GetAuditListReq, opts ...grpc.CallOption) (*GetAuditListResp, error)
		// 批量添加稽核信息
		BatchAddAudit(ctx context.Context, in *BatchAddAuditReq, opts ...grpc.CallOption) (*BatchAddAuditResp, error)
	}

	defaultLiveAuditRpcService struct {
		cli zrpc.Client
	}
)

func NewLiveAuditRpcService(cli zrpc.Client) LiveAuditRpcService {
	return &defaultLiveAuditRpcService{
		cli: cli,
	}
}

// 添加稽核 传递用户币种和金额
func (m *defaultLiveAuditRpcService) AddAudit(ctx context.Context, in *AddAuditReq, opts ...grpc.CallOption) (*AddAuditResp, error) {
	client := v1.NewLiveAuditRpcServiceClient(m.cli.Conn())
	return client.AddAudit(ctx, in, opts...)
}

// 更新稽核金额（打稽核） 传递用户币种和金额
func (m *defaultLiveAuditRpcService) UpdateAuditAmount(ctx context.Context, in *UpdateAuditAmountReq, opts ...grpc.CallOption) (*UpdateAuditAmountResp, error) {
	client := v1.NewLiveAuditRpcServiceClient(m.cli.Conn())
	return client.UpdateAuditAmount(ctx, in, opts...)
}

// 获取用户稽核信息
func (m *defaultLiveAuditRpcService) GetAuditInfo(ctx context.Context, in *GetAuditInfoReq, opts ...grpc.CallOption) (*GetAuditInfoResp, error) {
	client := v1.NewLiveAuditRpcServiceClient(m.cli.Conn())
	return client.GetAuditInfo(ctx, in, opts...)
}

// 获取指定用户稽核列表
func (m *defaultLiveAuditRpcService) GetAuditList(ctx context.Context, in *GetAuditListReq, opts ...grpc.CallOption) (*GetAuditListResp, error) {
	client := v1.NewLiveAuditRpcServiceClient(m.cli.Conn())
	return client.GetAuditList(ctx, in, opts...)
}

// 批量添加稽核信息
func (m *defaultLiveAuditRpcService) BatchAddAudit(ctx context.Context, in *BatchAddAuditReq, opts ...grpc.CallOption) (*BatchAddAuditResp, error) {
	client := v1.NewLiveAuditRpcServiceClient(m.cli.Conn())
	return client.BatchAddAudit(ctx, in, opts...)
}
