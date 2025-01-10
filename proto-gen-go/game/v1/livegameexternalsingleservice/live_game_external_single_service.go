// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: game.proto

package livegameexternalsingleservice

import (
	"context"

	"github.com/hanyougame/live-proto/proto-gen-go/game/v1"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CategoryNameBase                         = v1.CategoryNameBase
	GameCategoryDetail                       = v1.GameCategoryDetail
	GameDetails                              = v1.GameDetails
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
	WalletTransferInGameReply                = v1.WalletTransferInGameReply
	WalletTransferInGameReq                  = v1.WalletTransferInGameReq
	WalletTransferOutGameReply               = v1.WalletTransferOutGameReply
	WalletTransferOutGameReq                 = v1.WalletTransferOutGameReq

	LiveGameExternalSingleService interface {
		// 进入游戏
		EnterGame(ctx context.Context, in *SingleEnterGameReq, opts ...grpc.CallOption) (*SingleEnterGameReply, error)
		// 进入游戏试玩
		EnterGameTry(ctx context.Context, in *SingleEnterGameTryReq, opts ...grpc.CallOption) (*SingleEnterGameTryReply, error)
	}

	defaultLiveGameExternalSingleService struct {
		cli zrpc.Client
	}
)

func NewLiveGameExternalSingleService(cli zrpc.Client) LiveGameExternalSingleService {
	return &defaultLiveGameExternalSingleService{
		cli: cli,
	}
}

// 进入游戏
func (m *defaultLiveGameExternalSingleService) EnterGame(ctx context.Context, in *SingleEnterGameReq, opts ...grpc.CallOption) (*SingleEnterGameReply, error) {
	client := v1.NewLiveGameExternalSingleServiceClient(m.cli.Conn())
	return client.EnterGame(ctx, in, opts...)
}

// 进入游戏试玩
func (m *defaultLiveGameExternalSingleService) EnterGameTry(ctx context.Context, in *SingleEnterGameTryReq, opts ...grpc.CallOption) (*SingleEnterGameTryReply, error) {
	client := v1.NewLiveGameExternalSingleServiceClient(m.cli.Conn())
	return client.EnterGameTry(ctx, in, opts...)
}