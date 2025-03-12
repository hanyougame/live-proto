// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.6
// Source: finance.proto

package livewithdrawlimitrpcservice

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

	LiveWithdrawLimitRpcService interface {
		// 添加提现限制 传递用户币种和金额
		AddWithdrawLimit(ctx context.Context, in *AddWithdrawLimitReq, opts ...grpc.CallOption) (*AddWithdrawLimitResp, error)
		// 更新提现限制金额 传递用户币种和金额
		UpdateWithdrawLimitAmount(ctx context.Context, in *UpdateWithdrawLimitAmountReq, opts ...grpc.CallOption) (*UpdateWithdrawLimitAmountResp, error)
	}

	defaultLiveWithdrawLimitRpcService struct {
		cli zrpc.Client
	}
)

func NewLiveWithdrawLimitRpcService(cli zrpc.Client) LiveWithdrawLimitRpcService {
	return &defaultLiveWithdrawLimitRpcService{
		cli: cli,
	}
}

// 添加提现限制 传递用户币种和金额
func (m *defaultLiveWithdrawLimitRpcService) AddWithdrawLimit(ctx context.Context, in *AddWithdrawLimitReq, opts ...grpc.CallOption) (*AddWithdrawLimitResp, error) {
	client := v1.NewLiveWithdrawLimitRpcServiceClient(m.cli.Conn())
	return client.AddWithdrawLimit(ctx, in, opts...)
}

// 更新提现限制金额 传递用户币种和金额
func (m *defaultLiveWithdrawLimitRpcService) UpdateWithdrawLimitAmount(ctx context.Context, in *UpdateWithdrawLimitAmountReq, opts ...grpc.CallOption) (*UpdateWithdrawLimitAmountResp, error) {
	client := v1.NewLiveWithdrawLimitRpcServiceClient(m.cli.Conn())
	return client.UpdateWithdrawLimitAmount(ctx, in, opts...)
}
