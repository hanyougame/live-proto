// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.6
// Source: game.proto

package livegamecallbackrpcservice

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
	GameValueItem                            = v1.GameValueItem
	GetCategoryListByCurrReply               = v1.GetCategoryListByCurrReply
	GetCategoryListByCurrReq                 = v1.GetCategoryListByCurrReq
	GetCategorySimpleListByCurrReply         = v1.GetCategorySimpleListByCurrReply
	GetGameConfInfoReply                     = v1.GetGameConfInfoReply
	GetGameConfInfoReq                       = v1.GetGameConfInfoReq
	GetGameDetailsByThirdReply               = v1.GetGameDetailsByThirdReply
	GetGameDetailsByThirdReq                 = v1.GetGameDetailsByThirdReq
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
	GetNewGameListReply                      = v1.GetNewGameListReply
	GetNewGameListReq                        = v1.GetNewGameListReq
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

	LiveGameCallbackRpcService interface {
		// 获取accessKey
		GetK9GameAccessKey(ctx context.Context, in *GetK9GameAccessKeyReq, opts ...grpc.CallOption) (*GetK9GameAccessKeyReply, error)
		// 游戏投注
		GameUserBet(ctx context.Context, in *GameUserBetReq, opts ...grpc.CallOption) (*GameUserBetReply, error)
		// 订单派奖
		GameUserReward(ctx context.Context, in *GameUserRewardReq, opts ...grpc.CallOption) (*GameUserRewardReply, error)
		// 取消游戏投注
		GameUserBetCancel(ctx context.Context, in *GameUserBetCancelReq, opts ...grpc.CallOption) (*GameUserBetCancelReply, error)
		// 调账
		GameUserAdjustment(ctx context.Context, in *GameUserAdjustmentReq, opts ...grpc.CallOption) (*GameUserAdjustmentReply, error)
		// 订单回调
		TransferCallback(ctx context.Context, in *TransferCallbackReq, opts ...grpc.CallOption) (*TransferCallbackReply, error)
	}

	defaultLiveGameCallbackRpcService struct {
		cli zrpc.Client
	}
)

func NewLiveGameCallbackRpcService(cli zrpc.Client) LiveGameCallbackRpcService {
	return &defaultLiveGameCallbackRpcService{
		cli: cli,
	}
}

// 获取accessKey
func (m *defaultLiveGameCallbackRpcService) GetK9GameAccessKey(ctx context.Context, in *GetK9GameAccessKeyReq, opts ...grpc.CallOption) (*GetK9GameAccessKeyReply, error) {
	client := v1.NewLiveGameCallbackRpcServiceClient(m.cli.Conn())
	return client.GetK9GameAccessKey(ctx, in, opts...)
}

// 游戏投注
func (m *defaultLiveGameCallbackRpcService) GameUserBet(ctx context.Context, in *GameUserBetReq, opts ...grpc.CallOption) (*GameUserBetReply, error) {
	client := v1.NewLiveGameCallbackRpcServiceClient(m.cli.Conn())
	return client.GameUserBet(ctx, in, opts...)
}

// 订单派奖
func (m *defaultLiveGameCallbackRpcService) GameUserReward(ctx context.Context, in *GameUserRewardReq, opts ...grpc.CallOption) (*GameUserRewardReply, error) {
	client := v1.NewLiveGameCallbackRpcServiceClient(m.cli.Conn())
	return client.GameUserReward(ctx, in, opts...)
}

// 取消游戏投注
func (m *defaultLiveGameCallbackRpcService) GameUserBetCancel(ctx context.Context, in *GameUserBetCancelReq, opts ...grpc.CallOption) (*GameUserBetCancelReply, error) {
	client := v1.NewLiveGameCallbackRpcServiceClient(m.cli.Conn())
	return client.GameUserBetCancel(ctx, in, opts...)
}

// 调账
func (m *defaultLiveGameCallbackRpcService) GameUserAdjustment(ctx context.Context, in *GameUserAdjustmentReq, opts ...grpc.CallOption) (*GameUserAdjustmentReply, error) {
	client := v1.NewLiveGameCallbackRpcServiceClient(m.cli.Conn())
	return client.GameUserAdjustment(ctx, in, opts...)
}

// 订单回调
func (m *defaultLiveGameCallbackRpcService) TransferCallback(ctx context.Context, in *TransferCallbackReq, opts ...grpc.CallOption) (*TransferCallbackReply, error) {
	client := v1.NewLiveGameCallbackRpcServiceClient(m.cli.Conn())
	return client.TransferCallback(ctx, in, opts...)
}
