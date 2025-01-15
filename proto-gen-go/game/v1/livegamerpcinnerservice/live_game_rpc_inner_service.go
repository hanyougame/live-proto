// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: game.proto

package livegamerpcinnerservice

import (
	"context"

	"github.com/hanyougame/live-proto/proto-gen-go/game/v1"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddTripartiteTransferRecordReq           = v1.AddTripartiteTransferRecordReq
	AddTripartiteTransferRecordStatusReq     = v1.AddTripartiteTransferRecordStatusReq
	CategoryNameBase                         = v1.CategoryNameBase
	CreateCompensationRecordReq              = v1.CreateCompensationRecordReq
	CreateCompensationRecordResp             = v1.CreateCompensationRecordResp
	GameCategoryDetail                       = v1.GameCategoryDetail
	GameDetails                              = v1.GameDetails
	GameDetailsReq                           = v1.GameDetailsReq
	GameHandelFavoriteReq                    = v1.GameHandelFavoriteReq
	GamePlatformDetail                       = v1.GamePlatformDetail
	GameReply                                = v1.GameReply
	GameReq                                  = v1.GameReq
	GameUserAdjustmentReply                  = v1.GameUserAdjustmentReply
	GameUserAdjustmentReq                    = v1.GameUserAdjustmentReq
	GameUserBetCancelReply                   = v1.GameUserBetCancelReply
	GameUserBetCancelReq                     = v1.GameUserBetCancelReq
	GameUserBetReply                         = v1.GameUserBetReply
	GameUserBetReq                           = v1.GameUserBetReq
	GameUserRewardReply                      = v1.GameUserRewardReply
	GameUserRewardReq                        = v1.GameUserRewardReq
	GetCategoryListByCurrReply               = v1.GetCategoryListByCurrReply
	GetCategoryListByCurrReq                 = v1.GetCategoryListByCurrReq
	GetGameDetailsListReply                  = v1.GetGameDetailsListReply
	GetGameFavoriteListReq                   = v1.GetGameFavoriteListReq
	GetGameListByCategoryReq                 = v1.GetGameListByCategoryReq
	GetGameListByPlatformReq                 = v1.GetGameListByPlatformReq
	GetGameListBySearchReq                   = v1.GetGameListBySearchReq
	GetGameTransferBetOrderListReply         = v1.GetGameTransferBetOrderListReply
	GetGameTransferBetOrderListReplyBetInfo  = v1.GetGameTransferBetOrderListReplyBetInfo
	GetGameTransferBetOrderListReq           = v1.GetGameTransferBetOrderListReq
	GetGameTransferOrderStatusReply          = v1.GetGameTransferOrderStatusReply
	GetGameTransferOrderStatusReplyOrderInfo = v1.GetGameTransferOrderStatusReplyOrderInfo
	GetGameTransferOrderStatusReq            = v1.GetGameTransferOrderStatusReq
	GetHotGameListReq                        = v1.GetHotGameListReq
	GetHotPlatformListReply                  = v1.GetHotPlatformListReply
	GetHotPlatformListReq                    = v1.GetHotPlatformListReq
	GetK9GameAccessKeyReply                  = v1.GetK9GameAccessKeyReply
	GetK9GameAccessKeyReq                    = v1.GetK9GameAccessKeyReq
	GetPlatformListByCurrReply               = v1.GetPlatformListByCurrReply
	GetPlatformListByCurrReq                 = v1.GetPlatformListByCurrReq
	GetWalletTransferBalanceReply            = v1.GetWalletTransferBalanceReply
	GetWalletTransferBalanceReq              = v1.GetWalletTransferBalanceReq
	PlatformRedirectionBase                  = v1.PlatformRedirectionBase
	SingleEnterGameReply                     = v1.SingleEnterGameReply
	SingleEnterGameReq                       = v1.SingleEnterGameReq
	SingleEnterGameTryReply                  = v1.SingleEnterGameTryReply
	SingleEnterGameTryReq                    = v1.SingleEnterGameTryReq
	TransferCallbackReply                    = v1.TransferCallbackReply
	TransferCallbackReq                      = v1.TransferCallbackReq
	TransferEnterGameReply                   = v1.TransferEnterGameReply
	TransferEnterGameReq                     = v1.TransferEnterGameReq
	TripartiteTransferRecord                 = v1.TripartiteTransferRecord
	TripartiteTransferRecordStatusReq        = v1.TripartiteTransferRecordStatusReq
	WalletTransferInGameReply                = v1.WalletTransferInGameReply
	WalletTransferInGameReq                  = v1.WalletTransferInGameReq
	WalletTransferOutGameReply               = v1.WalletTransferOutGameReply
	WalletTransferOutGameReq                 = v1.WalletTransferOutGameReq

	LiveGameRpcInnerService interface {
		// 添加三方转账记录
		AddTripartiteTransferRecord(ctx context.Context, in *AddTripartiteTransferRecordReq, opts ...grpc.CallOption) (*GameReply, error)
		// 变更三方转账记录状态
		AddTripartiteTransferRecordStatus(ctx context.Context, in *AddTripartiteTransferRecordStatusReq, opts ...grpc.CallOption) (*GameReply, error)
		// 查询某一条的状态数据
		TripartiteTransferRecordStatus(ctx context.Context, in *TripartiteTransferRecordStatusReq, opts ...grpc.CallOption) (*TripartiteTransferRecord, error)
		// 创建补偿失败记录
		CreateCompensationFailedRecord(ctx context.Context, in *CreateCompensationRecordReq, opts ...grpc.CallOption) (*CreateCompensationRecordResp, error)
	}

	defaultLiveGameRpcInnerService struct {
		cli zrpc.Client
	}
)

func NewLiveGameRpcInnerService(cli zrpc.Client) LiveGameRpcInnerService {
	return &defaultLiveGameRpcInnerService{
		cli: cli,
	}
}

// 添加三方转账记录
func (m *defaultLiveGameRpcInnerService) AddTripartiteTransferRecord(ctx context.Context, in *AddTripartiteTransferRecordReq, opts ...grpc.CallOption) (*GameReply, error) {
	client := v1.NewLiveGameRpcInnerServiceClient(m.cli.Conn())
	return client.AddTripartiteTransferRecord(ctx, in, opts...)
}

// 变更三方转账记录状态
func (m *defaultLiveGameRpcInnerService) AddTripartiteTransferRecordStatus(ctx context.Context, in *AddTripartiteTransferRecordStatusReq, opts ...grpc.CallOption) (*GameReply, error) {
	client := v1.NewLiveGameRpcInnerServiceClient(m.cli.Conn())
	return client.AddTripartiteTransferRecordStatus(ctx, in, opts...)
}

// 查询某一条的状态数据
func (m *defaultLiveGameRpcInnerService) TripartiteTransferRecordStatus(ctx context.Context, in *TripartiteTransferRecordStatusReq, opts ...grpc.CallOption) (*TripartiteTransferRecord, error) {
	client := v1.NewLiveGameRpcInnerServiceClient(m.cli.Conn())
	return client.TripartiteTransferRecordStatus(ctx, in, opts...)
}

// 创建补偿失败记录
func (m *defaultLiveGameRpcInnerService) CreateCompensationFailedRecord(ctx context.Context, in *CreateCompensationRecordReq, opts ...grpc.CallOption) (*CreateCompensationRecordResp, error) {
	client := v1.NewLiveGameRpcInnerServiceClient(m.cli.Conn())
	return client.CreateCompensationFailedRecord(ctx, in, opts...)
}
