// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.6
// Source: game.proto

package livegameexternaltransferservice

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

	LiveGameExternalTransferService interface {
		// 进入游戏
		EnterGame(ctx context.Context, in *TransferEnterGameReq, opts ...grpc.CallOption) (*TransferEnterGameReply, error)
		// 转账钱包转入游戏
		WalletTransferInGame(ctx context.Context, in *WalletTransferInGameReq, opts ...grpc.CallOption) (*WalletTransferInGameReply, error)
		// 转账钱包转入
		WalletTransferOutGame(ctx context.Context, in *WalletTransferOutGameReq, opts ...grpc.CallOption) (*WalletTransferOutGameReply, error)
		// 转账钱包余额查询
		GetWalletTransferBalance(ctx context.Context, in *GetWalletTransferBalanceReq, opts ...grpc.CallOption) (*GetWalletTransferBalanceReply, error)
		// 查询用户游戏转账（转入、转出）订单状态
		GetGameTransferOrderStatus(ctx context.Context, in *GetGameTransferOrderStatusReq, opts ...grpc.CallOption) (*GetGameTransferOrderStatusReply, error)
		// 查询用户游戏投注订单列表
		GetGameTransferBetOrderList(ctx context.Context, in *GetGameTransferBetOrderListReq, opts ...grpc.CallOption) (*GetGameTransferBetOrderListReply, error)
	}

	defaultLiveGameExternalTransferService struct {
		cli zrpc.Client
	}
)

func NewLiveGameExternalTransferService(cli zrpc.Client) LiveGameExternalTransferService {
	return &defaultLiveGameExternalTransferService{
		cli: cli,
	}
}

// 进入游戏
func (m *defaultLiveGameExternalTransferService) EnterGame(ctx context.Context, in *TransferEnterGameReq, opts ...grpc.CallOption) (*TransferEnterGameReply, error) {
	client := v1.NewLiveGameExternalTransferServiceClient(m.cli.Conn())
	return client.EnterGame(ctx, in, opts...)
}

// 转账钱包转入游戏
func (m *defaultLiveGameExternalTransferService) WalletTransferInGame(ctx context.Context, in *WalletTransferInGameReq, opts ...grpc.CallOption) (*WalletTransferInGameReply, error) {
	client := v1.NewLiveGameExternalTransferServiceClient(m.cli.Conn())
	return client.WalletTransferInGame(ctx, in, opts...)
}

// 转账钱包转入
func (m *defaultLiveGameExternalTransferService) WalletTransferOutGame(ctx context.Context, in *WalletTransferOutGameReq, opts ...grpc.CallOption) (*WalletTransferOutGameReply, error) {
	client := v1.NewLiveGameExternalTransferServiceClient(m.cli.Conn())
	return client.WalletTransferOutGame(ctx, in, opts...)
}

// 转账钱包余额查询
func (m *defaultLiveGameExternalTransferService) GetWalletTransferBalance(ctx context.Context, in *GetWalletTransferBalanceReq, opts ...grpc.CallOption) (*GetWalletTransferBalanceReply, error) {
	client := v1.NewLiveGameExternalTransferServiceClient(m.cli.Conn())
	return client.GetWalletTransferBalance(ctx, in, opts...)
}

// 查询用户游戏转账（转入、转出）订单状态
func (m *defaultLiveGameExternalTransferService) GetGameTransferOrderStatus(ctx context.Context, in *GetGameTransferOrderStatusReq, opts ...grpc.CallOption) (*GetGameTransferOrderStatusReply, error) {
	client := v1.NewLiveGameExternalTransferServiceClient(m.cli.Conn())
	return client.GetGameTransferOrderStatus(ctx, in, opts...)
}

// 查询用户游戏投注订单列表
func (m *defaultLiveGameExternalTransferService) GetGameTransferBetOrderList(ctx context.Context, in *GetGameTransferBetOrderListReq, opts ...grpc.CallOption) (*GetGameTransferBetOrderListReply, error) {
	client := v1.NewLiveGameExternalTransferServiceClient(m.cli.Conn())
	return client.GetGameTransferBetOrderList(ctx, in, opts...)
}
