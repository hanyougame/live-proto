// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: game.proto

package livegameexternalservice

import (
	"context"

	"github.com/hanyougame/live-proto/proto-gen-go/game/v1"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
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
	GetGameTransferBetOrderListReply         = v1.GetGameTransferBetOrderListReply
	GetGameTransferBetOrderListReplyBetInfo  = v1.GetGameTransferBetOrderListReplyBetInfo
	GetGameTransferBetOrderListReq           = v1.GetGameTransferBetOrderListReq
	GetGameTransferOrderStatusReply          = v1.GetGameTransferOrderStatusReply
	GetGameTransferOrderStatusReplyOrderInfo = v1.GetGameTransferOrderStatusReplyOrderInfo
	GetGameTransferOrderStatusReq            = v1.GetGameTransferOrderStatusReq
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
	WalletTransferInGameReply                = v1.WalletTransferInGameReply
	WalletTransferInGameReq                  = v1.WalletTransferInGameReq
	WalletTransferOutGameReply               = v1.WalletTransferOutGameReply
	WalletTransferOutGameReq                 = v1.WalletTransferOutGameReq

	LiveGameExternalService interface {
		// 获取游戏资源信息
		K9GameResourceListSync(ctx context.Context, in *GameReq, opts ...grpc.CallOption) (*GameReply, error)
	}

	defaultLiveGameExternalService struct {
		cli zrpc.Client
	}
)

func NewLiveGameExternalService(cli zrpc.Client) LiveGameExternalService {
	return &defaultLiveGameExternalService{
		cli: cli,
	}
}

// 获取游戏资源信息
func (m *defaultLiveGameExternalService) K9GameResourceListSync(ctx context.Context, in *GameReq, opts ...grpc.CallOption) (*GameReply, error) {
	client := v1.NewLiveGameExternalServiceClient(m.cli.Conn())
	return client.K9GameResourceListSync(ctx, in, opts...)
}
