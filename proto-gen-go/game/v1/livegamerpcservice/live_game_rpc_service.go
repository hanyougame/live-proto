// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.6
// Source: game.proto

package livegamerpcservice

import (
	"context"

	"github.com/hanyougame/live-proto/proto-gen-go/game/v1"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddGameAdjustmentRecordReq               = v1.AddGameAdjustmentRecordReq
	AddGameBetBaseReply                      = v1.AddGameBetBaseReply
	AddGameBetRecordReply                    = v1.AddGameBetRecordReply
	AddGameBetRecordReq                      = v1.AddGameBetRecordReq
	AddGameCancelRecordReq                   = v1.AddGameCancelRecordReq
	AddGameSettledRecordReq                  = v1.AddGameSettledRecordReq
	AddRecentlyGamePlayReq                   = v1.AddRecentlyGamePlayReq
	AddTransferGameBetRecordReq              = v1.AddTransferGameBetRecordReq
	AddTripartiteTransferRecordReq           = v1.AddTripartiteTransferRecordReq
	AddTripartiteTransferRecordStatusReq     = v1.AddTripartiteTransferRecordStatusReq
	BetRecordInfo                            = v1.BetRecordInfo
	BetSummaryInfo                           = v1.BetSummaryInfo
	CategoryNameBase                         = v1.CategoryNameBase
	CreateCompensationRecordReq              = v1.CreateCompensationRecordReq
	CreateCompensationRecordResp             = v1.CreateCompensationRecordResp
	GameCategoryDetail                       = v1.GameCategoryDetail
	GameCategorySimpleDetail                 = v1.GameCategorySimpleDetail
	GameDetails                              = v1.GameDetails
	GameDetailsList                          = v1.GameDetailsList
	GameDetailsReq                           = v1.GameDetailsReq
	GameHandelFavoriteReq                    = v1.GameHandelFavoriteReq
	GamePlatformDetail                       = v1.GamePlatformDetail
	GamePlatformDetailsReq                   = v1.GamePlatformDetailsReq
	GamePlatformSimpleDetail                 = v1.GamePlatformSimpleDetail
	GameReply                                = v1.GameReply
	GameReq                                  = v1.GameReq
	GameSimpleDetails                        = v1.GameSimpleDetails
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
	GetCategorySimpleListByCurrReply         = v1.GetCategorySimpleListByCurrReply
	GetGameDetailsListReply                  = v1.GetGameDetailsListReply
	GetGameFavoriteListReq                   = v1.GetGameFavoriteListReq
	GetGameListByCategoryReq                 = v1.GetGameListByCategoryReq
	GetGameListByPlatformReq                 = v1.GetGameListByPlatformReq
	GetGameListBySearchReq                   = v1.GetGameListBySearchReq
	GetGameRecentlyListReq                   = v1.GetGameRecentlyListReq
	GetGameSimpleListBySearchReply           = v1.GetGameSimpleListBySearchReply
	GetGameTransferBetOrderListReply         = v1.GetGameTransferBetOrderListReply
	GetGameTransferBetOrderListReq           = v1.GetGameTransferBetOrderListReq
	GetGameTransferOrderStatusReply          = v1.GetGameTransferOrderStatusReply
	GetGameTransferOrderStatusReplyOrderInfo = v1.GetGameTransferOrderStatusReplyOrderInfo
	GetGameTransferOrderStatusReq            = v1.GetGameTransferOrderStatusReq
	GetHomeGameItemReq                       = v1.GetHomeGameItemReq
	GetHomeGameItemsReply                    = v1.GetHomeGameItemsReply
	GetHomePlatformItemsReply                = v1.GetHomePlatformItemsReply
	GetHotGameListReq                        = v1.GetHotGameListReq
	GetHotPlatformListReply                  = v1.GetHotPlatformListReply
	GetHotPlatformListReq                    = v1.GetHotPlatformListReq
	GetK9GameAccessKeyReply                  = v1.GetK9GameAccessKeyReply
	GetK9GameAccessKeyReq                    = v1.GetK9GameAccessKeyReq
	GetPlatListSimpleByCurrReply             = v1.GetPlatListSimpleByCurrReply
	GetPlatformListByCurrReply               = v1.GetPlatformListByCurrReply
	GetPlatformListByCurrReq                 = v1.GetPlatformListByCurrReq
	GetUserBetRecordListReply                = v1.GetUserBetRecordListReply
	GetUserBetRecordListReq                  = v1.GetUserBetRecordListReq
	GetUserBetRecordSummaryReply             = v1.GetUserBetRecordSummaryReply
	GetUserBetRecordSummaryReq               = v1.GetUserBetRecordSummaryReq
	GetUserFavoriteIdsReply                  = v1.GetUserFavoriteIdsReply
	GetUserFavoriteIdsReq                    = v1.GetUserFavoriteIdsReq
	GetWalletTransferBalanceReply            = v1.GetWalletTransferBalanceReply
	GetWalletTransferBalanceReq              = v1.GetWalletTransferBalanceReq
	PlatformDetailsList                      = v1.PlatformDetailsList
	PlatformRedirectionBase                  = v1.PlatformRedirectionBase
	ProcessMessageTransferDataReply          = v1.ProcessMessageTransferDataReply
	ProcessMessageTransferDataReq            = v1.ProcessMessageTransferDataReq
	ProcessMessageTransferSendReply          = v1.ProcessMessageTransferSendReply
	ProcessMessageTransferSendReq            = v1.ProcessMessageTransferSendReq
	SendGameBetBetMQReq                      = v1.SendGameBetBetMQReq
	SendGameBetBetSettlementMQReq            = v1.SendGameBetBetSettlementMQReq
	SingleEnterGameReply                     = v1.SingleEnterGameReply
	SingleEnterGameReq                       = v1.SingleEnterGameReq
	SingleEnterGameTryReply                  = v1.SingleEnterGameTryReply
	SingleEnterGameTryReq                    = v1.SingleEnterGameTryReq
	TransferBetRecord                        = v1.TransferBetRecord
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
		// 通过货币获取游戏类型列表(简单信息)
		GetGameCategorySimpleListByCurr(ctx context.Context, in *GetCategoryListByCurrReq, opts ...grpc.CallOption) (*GetCategorySimpleListByCurrReply, error)
		// 通过游戏类型获取游戏列表
		GetGameListByCategory(ctx context.Context, in *GetGameListByCategoryReq, opts ...grpc.CallOption) (*GetGameDetailsListReply, error)
		// 通过货币获取平台列表
		GetPlatformListByCurr(ctx context.Context, in *GetPlatformListByCurrReq, opts ...grpc.CallOption) (*GetPlatformListByCurrReply, error)
		// 通过货币获取平台列表(简单信息)
		GetPlatListSimpleByCurr(ctx context.Context, in *GetPlatformListByCurrReq, opts ...grpc.CallOption) (*GetPlatListSimpleByCurrReply, error)
		// 获取平台详情
		GetPlatformDetails(ctx context.Context, in *GamePlatformDetailsReq, opts ...grpc.CallOption) (*GamePlatformDetail, error)
		// 通过平台获取游戏列表
		GetGameListByPlatform(ctx context.Context, in *GetGameListByPlatformReq, opts ...grpc.CallOption) (*GetGameDetailsListReply, error)
		// 通过搜索获取游戏列表
		GetGameListBySearch(ctx context.Context, in *GetGameListBySearchReq, opts ...grpc.CallOption) (*GetGameDetailsListReply, error)
		// 通过搜索获取游戏列表 (简单信息)
		GetGameSimpleListBySearch(ctx context.Context, in *GetGameListBySearchReq, opts ...grpc.CallOption) (*GetGameSimpleListBySearchReply, error)
		// 添加收藏
		GameAddFavorite(ctx context.Context, in *GameHandelFavoriteReq, opts ...grpc.CallOption) (*GameReply, error)
		// 移除收藏
		GameRemoveFavorite(ctx context.Context, in *GameHandelFavoriteReq, opts ...grpc.CallOption) (*GameReply, error)
		// 收藏列表
		GameFavoriteList(ctx context.Context, in *GetGameFavoriteListReq, opts ...grpc.CallOption) (*GetGameDetailsListReply, error)
		// 热门游戏列表
		GetHotGameList(ctx context.Context, in *GetHotGameListReq, opts ...grpc.CallOption) (*GetGameDetailsListReply, error)
		// 最近游戏列表
		GetRecentlyGameList(ctx context.Context, in *GetGameRecentlyListReq, opts ...grpc.CallOption) (*GetGameDetailsListReply, error)
		// 热门平台列表
		GetHotPlatformList(ctx context.Context, in *GetHotPlatformListReq, opts ...grpc.CallOption) (*GetHotPlatformListReply, error)
		// 根据游戏ID获取游戏详情
		GetGameDetails(ctx context.Context, in *GameDetailsReq, opts ...grpc.CallOption) (*GameDetails, error)
		// 获取用户收藏ID
		GetUserFavoriteIds(ctx context.Context, in *GetUserFavoriteIdsReq, opts ...grpc.CallOption) (*GetUserFavoriteIdsReply, error)
		// 获取用户投注记录
		GetUserBetRecordList(ctx context.Context, in *GetUserBetRecordListReq, opts ...grpc.CallOption) (*GetUserBetRecordListReply, error)
		// 获取用户投注报表
		GetUserBetRecordSummary(ctx context.Context, in *GetUserBetRecordSummaryReq, opts ...grpc.CallOption) (*GetUserBetRecordSummaryReply, error)
		// 获取首页平台详情
		GetHomePlatformItems(ctx context.Context, in *GetHomeGameItemReq, opts ...grpc.CallOption) (*GetHomePlatformItemsReply, error)
		// 获取首页游戏详情
		GetHomeGameItems(ctx context.Context, in *GetHomeGameItemReq, opts ...grpc.CallOption) (*GetHomeGameItemsReply, error)
		// 获取指定数量的游戏列表
		GetHomeGameList(ctx context.Context, in *GetHomeGameItemReq, opts ...grpc.CallOption) (*GameDetailsList, error)
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

// 通过货币获取游戏类型列表(简单信息)
func (m *defaultLiveGameRpcService) GetGameCategorySimpleListByCurr(ctx context.Context, in *GetCategoryListByCurrReq, opts ...grpc.CallOption) (*GetCategorySimpleListByCurrReply, error) {
	client := v1.NewLiveGameRpcServiceClient(m.cli.Conn())
	return client.GetGameCategorySimpleListByCurr(ctx, in, opts...)
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

// 通过货币获取平台列表(简单信息)
func (m *defaultLiveGameRpcService) GetPlatListSimpleByCurr(ctx context.Context, in *GetPlatformListByCurrReq, opts ...grpc.CallOption) (*GetPlatListSimpleByCurrReply, error) {
	client := v1.NewLiveGameRpcServiceClient(m.cli.Conn())
	return client.GetPlatListSimpleByCurr(ctx, in, opts...)
}

// 获取平台详情
func (m *defaultLiveGameRpcService) GetPlatformDetails(ctx context.Context, in *GamePlatformDetailsReq, opts ...grpc.CallOption) (*GamePlatformDetail, error) {
	client := v1.NewLiveGameRpcServiceClient(m.cli.Conn())
	return client.GetPlatformDetails(ctx, in, opts...)
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

// 通过搜索获取游戏列表 (简单信息)
func (m *defaultLiveGameRpcService) GetGameSimpleListBySearch(ctx context.Context, in *GetGameListBySearchReq, opts ...grpc.CallOption) (*GetGameSimpleListBySearchReply, error) {
	client := v1.NewLiveGameRpcServiceClient(m.cli.Conn())
	return client.GetGameSimpleListBySearch(ctx, in, opts...)
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

// 最近游戏列表
func (m *defaultLiveGameRpcService) GetRecentlyGameList(ctx context.Context, in *GetGameRecentlyListReq, opts ...grpc.CallOption) (*GetGameDetailsListReply, error) {
	client := v1.NewLiveGameRpcServiceClient(m.cli.Conn())
	return client.GetRecentlyGameList(ctx, in, opts...)
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

// 获取用户投注记录
func (m *defaultLiveGameRpcService) GetUserBetRecordList(ctx context.Context, in *GetUserBetRecordListReq, opts ...grpc.CallOption) (*GetUserBetRecordListReply, error) {
	client := v1.NewLiveGameRpcServiceClient(m.cli.Conn())
	return client.GetUserBetRecordList(ctx, in, opts...)
}

// 获取用户投注报表
func (m *defaultLiveGameRpcService) GetUserBetRecordSummary(ctx context.Context, in *GetUserBetRecordSummaryReq, opts ...grpc.CallOption) (*GetUserBetRecordSummaryReply, error) {
	client := v1.NewLiveGameRpcServiceClient(m.cli.Conn())
	return client.GetUserBetRecordSummary(ctx, in, opts...)
}

// 获取首页平台详情
func (m *defaultLiveGameRpcService) GetHomePlatformItems(ctx context.Context, in *GetHomeGameItemReq, opts ...grpc.CallOption) (*GetHomePlatformItemsReply, error) {
	client := v1.NewLiveGameRpcServiceClient(m.cli.Conn())
	return client.GetHomePlatformItems(ctx, in, opts...)
}

// 获取首页游戏详情
func (m *defaultLiveGameRpcService) GetHomeGameItems(ctx context.Context, in *GetHomeGameItemReq, opts ...grpc.CallOption) (*GetHomeGameItemsReply, error) {
	client := v1.NewLiveGameRpcServiceClient(m.cli.Conn())
	return client.GetHomeGameItems(ctx, in, opts...)
}

// 获取指定数量的游戏列表
func (m *defaultLiveGameRpcService) GetHomeGameList(ctx context.Context, in *GetHomeGameItemReq, opts ...grpc.CallOption) (*GameDetailsList, error) {
	client := v1.NewLiveGameRpcServiceClient(m.cli.Conn())
	return client.GetHomeGameList(ctx, in, opts...)
}
