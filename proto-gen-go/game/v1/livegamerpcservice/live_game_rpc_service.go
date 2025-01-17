// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: game.proto

package livegamerpcservice

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
	GetUserFavoriteIdsReply                  = v1.GetUserFavoriteIdsReply
	GetUserFavoriteIdsReq                    = v1.GetUserFavoriteIdsReq
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

	LiveGameRpcService interface {
		// 通过货币获取游戏类型列表
		GetGameCategoryListByCurr(ctx context.Context, in *GetCategoryListByCurrReq, opts ...grpc.CallOption) (*GetCategoryListByCurrReply, error)
		// 通过游戏类型获取游戏列表
		GetGameListByCategory(ctx context.Context, in *GetGameListByCategoryReq, opts ...grpc.CallOption) (*GetGameDetailsListReply, error)
		// 通过货币获取平台列表
		GetPlatformListByCurr(ctx context.Context, in *GetPlatformListByCurrReq, opts ...grpc.CallOption) (*GetPlatformListByCurrReply, error)
		// 通过平台获取游戏列表
		GetGameListByPlatform(ctx context.Context, in *GetGameListByPlatformReq, opts ...grpc.CallOption) (*GetGameDetailsListReply, error)
		// 通过搜索获取游戏列表
		GetGameListBySearch(ctx context.Context, in *GetGameListBySearchReq, opts ...grpc.CallOption) (*GetGameDetailsListReply, error)
		// 添加收藏
		GameAddFavorite(ctx context.Context, in *GameHandelFavoriteReq, opts ...grpc.CallOption) (*GameReply, error)
		// 移除收藏
		GameRemoveFavorite(ctx context.Context, in *GameHandelFavoriteReq, opts ...grpc.CallOption) (*GameReply, error)
		// 收藏列表
		GameFavoriteList(ctx context.Context, in *GetGameFavoriteListReq, opts ...grpc.CallOption) (*GetGameDetailsListReply, error)
		// 热门游戏列表
		GetHotGameList(ctx context.Context, in *GetHotGameListReq, opts ...grpc.CallOption) (*GetGameDetailsListReply, error)
		// 热门平台列表
		GetHotPlatformList(ctx context.Context, in *GetHotPlatformListReq, opts ...grpc.CallOption) (*GetHotPlatformListReply, error)
		// 根据游戏ID获取游戏详情
		GetGameDetails(ctx context.Context, in *GameDetailsReq, opts ...grpc.CallOption) (*GameDetails, error)
		// 获取用户收藏ID
		GetUserFavoriteIds(ctx context.Context, in *GetUserFavoriteIdsReq, opts ...grpc.CallOption) (*GetUserFavoriteIdsReply, error)
	}

	defaultLiveGameRpcService struct {
		cli zrpc.Client
	}
)

func NewLiveGameRpcService(cli zrpc.Client) LiveGameRpcService {
	return &defaultLiveGameRpcService{
		cli: cli,
	}
}

// 通过货币获取游戏类型列表
func (m *defaultLiveGameRpcService) GetGameCategoryListByCurr(ctx context.Context, in *GetCategoryListByCurrReq, opts ...grpc.CallOption) (*GetCategoryListByCurrReply, error) {
	client := v1.NewLiveGameRpcServiceClient(m.cli.Conn())
	return client.GetGameCategoryListByCurr(ctx, in, opts...)
}

// 通过游戏类型获取游戏列表
func (m *defaultLiveGameRpcService) GetGameListByCategory(ctx context.Context, in *GetGameListByCategoryReq, opts ...grpc.CallOption) (*GetGameDetailsListReply, error) {
	client := v1.NewLiveGameRpcServiceClient(m.cli.Conn())
	return client.GetGameListByCategory(ctx, in, opts...)
}

// 通过货币获取平台列表
func (m *defaultLiveGameRpcService) GetPlatformListByCurr(ctx context.Context, in *GetPlatformListByCurrReq, opts ...grpc.CallOption) (*GetPlatformListByCurrReply, error) {
	client := v1.NewLiveGameRpcServiceClient(m.cli.Conn())
	return client.GetPlatformListByCurr(ctx, in, opts...)
}

// 通过平台获取游戏列表
func (m *defaultLiveGameRpcService) GetGameListByPlatform(ctx context.Context, in *GetGameListByPlatformReq, opts ...grpc.CallOption) (*GetGameDetailsListReply, error) {
	client := v1.NewLiveGameRpcServiceClient(m.cli.Conn())
	return client.GetGameListByPlatform(ctx, in, opts...)
}

// 通过搜索获取游戏列表
func (m *defaultLiveGameRpcService) GetGameListBySearch(ctx context.Context, in *GetGameListBySearchReq, opts ...grpc.CallOption) (*GetGameDetailsListReply, error) {
	client := v1.NewLiveGameRpcServiceClient(m.cli.Conn())
	return client.GetGameListBySearch(ctx, in, opts...)
}

// 添加收藏
func (m *defaultLiveGameRpcService) GameAddFavorite(ctx context.Context, in *GameHandelFavoriteReq, opts ...grpc.CallOption) (*GameReply, error) {
	client := v1.NewLiveGameRpcServiceClient(m.cli.Conn())
	return client.GameAddFavorite(ctx, in, opts...)
}

// 移除收藏
func (m *defaultLiveGameRpcService) GameRemoveFavorite(ctx context.Context, in *GameHandelFavoriteReq, opts ...grpc.CallOption) (*GameReply, error) {
	client := v1.NewLiveGameRpcServiceClient(m.cli.Conn())
	return client.GameRemoveFavorite(ctx, in, opts...)
}

// 收藏列表
func (m *defaultLiveGameRpcService) GameFavoriteList(ctx context.Context, in *GetGameFavoriteListReq, opts ...grpc.CallOption) (*GetGameDetailsListReply, error) {
	client := v1.NewLiveGameRpcServiceClient(m.cli.Conn())
	return client.GameFavoriteList(ctx, in, opts...)
}

// 热门游戏列表
func (m *defaultLiveGameRpcService) GetHotGameList(ctx context.Context, in *GetHotGameListReq, opts ...grpc.CallOption) (*GetGameDetailsListReply, error) {
	client := v1.NewLiveGameRpcServiceClient(m.cli.Conn())
	return client.GetHotGameList(ctx, in, opts...)
}

// 热门平台列表
func (m *defaultLiveGameRpcService) GetHotPlatformList(ctx context.Context, in *GetHotPlatformListReq, opts ...grpc.CallOption) (*GetHotPlatformListReply, error) {
	client := v1.NewLiveGameRpcServiceClient(m.cli.Conn())
	return client.GetHotPlatformList(ctx, in, opts...)
}

// 根据游戏ID获取游戏详情
func (m *defaultLiveGameRpcService) GetGameDetails(ctx context.Context, in *GameDetailsReq, opts ...grpc.CallOption) (*GameDetails, error) {
	client := v1.NewLiveGameRpcServiceClient(m.cli.Conn())
	return client.GetGameDetails(ctx, in, opts...)
}

// 获取用户收藏ID
func (m *defaultLiveGameRpcService) GetUserFavoriteIds(ctx context.Context, in *GetUserFavoriteIdsReq, opts ...grpc.CallOption) (*GetUserFavoriteIdsReply, error) {
	client := v1.NewLiveGameRpcServiceClient(m.cli.Conn())
	return client.GetUserFavoriteIds(ctx, in, opts...)
}
