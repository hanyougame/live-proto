// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: finance.proto

package liveauditrpcservice

import (
	"context"

	"github.com/hanyougame/live-proto/proto-gen-go/finance/v1"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddAuditReq                   = v1.AddAuditReq
	AddAuditResp                  = v1.AddAuditResp
	AddWithdrawLimitReq           = v1.AddWithdrawLimitReq
	AddWithdrawLimitResp          = v1.AddWithdrawLimitResp
	BalanceReq                    = v1.BalanceReq
	BalanceResp                   = v1.BalanceResp
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

	LiveAuditRpcService interface {
		// 添加稽核 传递用户币种和金额
		AddAudit(ctx context.Context, in *AddAuditReq, opts ...grpc.CallOption) (*AddAuditResp, error)
		// 更新稽核金额（打稽核） 传递用户币种和金额
		UpdateAuditAmount(ctx context.Context, in *UpdateAuditAmountReq, opts ...grpc.CallOption) (*UpdateAuditAmountResp, error)
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
