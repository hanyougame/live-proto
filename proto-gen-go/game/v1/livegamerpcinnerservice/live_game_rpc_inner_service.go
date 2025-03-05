// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.6
// Source: game.proto

package livegamerpcinnerservice

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

	LiveGameRpcInnerService interface {
		// 添加三方转账记录
		AddTripartiteTransferRecord(ctx context.Context, in *AddTripartiteTransferRecordReq, opts ...grpc.CallOption) (*GameReply, error)
		// 变更三方转账记录状态
		AddTripartiteTransferRecordStatus(ctx context.Context, in *AddTripartiteTransferRecordStatusReq, opts ...grpc.CallOption) (*GameReply, error)
		// 转账钱包处理
		ProcessMessageTransferData(ctx context.Context, in *ProcessMessageTransferDataReq, opts ...grpc.CallOption) (*ProcessMessageTransferDataReply, error)
		ProcessMessageTransferSend(ctx context.Context, in *ProcessMessageTransferSendReq, opts ...grpc.CallOption) (*ProcessMessageTransferSendReply, error)
		// 查询某一条的状态数据
		TripartiteTransferRecordStatus(ctx context.Context, in *TripartiteTransferRecordStatusReq, opts ...grpc.CallOption) (*TripartiteTransferRecord, error)
		// 创建补偿失败记录
		CreateCompensationFailedRecord(ctx context.Context, in *CreateCompensationRecordReq, opts ...grpc.CallOption) (*CreateCompensationRecordResp, error)
		// 添加游戏下注记录(单一钱包)
		AddGameBetRecord(ctx context.Context, in *AddGameBetRecordReq, opts ...grpc.CallOption) (*AddGameBetRecordReply, error)
		// 变更游戏下注记录结算状态
		AddGameSettledRecord(ctx context.Context, in *AddGameSettledRecordReq, opts ...grpc.CallOption) (*AddGameBetBaseReply, error)
		// 变更游戏取消记录状态
		AddGameCancelRecord(ctx context.Context, in *AddGameCancelRecordReq, opts ...grpc.CallOption) (*AddGameBetBaseReply, error)
		// 变更游戏调整记录状态
		AddGameAdjustmentRecord(ctx context.Context, in *AddGameAdjustmentRecordReq, opts ...grpc.CallOption) (*AddGameBetBaseReply, error)
		// 添加游戏下注记录(转账钱包)
		AddTransferGameBetRecord(ctx context.Context, in *AddTransferGameBetRecordReq, opts ...grpc.CallOption) (*GameReply, error)
		// 发送游戏下注MQ
		SendGameBetBetMQ(ctx context.Context, in *SendGameBetBetMQReq, opts ...grpc.CallOption) (*GameReply, error)
		// 发送游戏下注结算MQ
		SendGameBetBetSettlementMQ(ctx context.Context, in *SendGameBetBetSettlementMQReq, opts ...grpc.CallOption) (*GameReply, error)
		// 添加最近游玩的游戏
		AddRecentlyGamePlay(ctx context.Context, in *AddRecentlyGamePlayReq, opts ...grpc.CallOption) (*GameReply, error)
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

// 转账钱包处理
func (m *defaultLiveGameRpcInnerService) ProcessMessageTransferData(ctx context.Context, in *ProcessMessageTransferDataReq, opts ...grpc.CallOption) (*ProcessMessageTransferDataReply, error) {
	client := v1.NewLiveGameRpcInnerServiceClient(m.cli.Conn())
	return client.ProcessMessageTransferData(ctx, in, opts...)
}

func (m *defaultLiveGameRpcInnerService) ProcessMessageTransferSend(ctx context.Context, in *ProcessMessageTransferSendReq, opts ...grpc.CallOption) (*ProcessMessageTransferSendReply, error) {
	client := v1.NewLiveGameRpcInnerServiceClient(m.cli.Conn())
	return client.ProcessMessageTransferSend(ctx, in, opts...)
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

// 添加游戏下注记录(单一钱包)
func (m *defaultLiveGameRpcInnerService) AddGameBetRecord(ctx context.Context, in *AddGameBetRecordReq, opts ...grpc.CallOption) (*AddGameBetRecordReply, error) {
	client := v1.NewLiveGameRpcInnerServiceClient(m.cli.Conn())
	return client.AddGameBetRecord(ctx, in, opts...)
}

// 变更游戏下注记录结算状态
func (m *defaultLiveGameRpcInnerService) AddGameSettledRecord(ctx context.Context, in *AddGameSettledRecordReq, opts ...grpc.CallOption) (*AddGameBetBaseReply, error) {
	client := v1.NewLiveGameRpcInnerServiceClient(m.cli.Conn())
	return client.AddGameSettledRecord(ctx, in, opts...)
}

// 变更游戏取消记录状态
func (m *defaultLiveGameRpcInnerService) AddGameCancelRecord(ctx context.Context, in *AddGameCancelRecordReq, opts ...grpc.CallOption) (*AddGameBetBaseReply, error) {
	client := v1.NewLiveGameRpcInnerServiceClient(m.cli.Conn())
	return client.AddGameCancelRecord(ctx, in, opts...)
}

// 变更游戏调整记录状态
func (m *defaultLiveGameRpcInnerService) AddGameAdjustmentRecord(ctx context.Context, in *AddGameAdjustmentRecordReq, opts ...grpc.CallOption) (*AddGameBetBaseReply, error) {
	client := v1.NewLiveGameRpcInnerServiceClient(m.cli.Conn())
	return client.AddGameAdjustmentRecord(ctx, in, opts...)
}

// 添加游戏下注记录(转账钱包)
func (m *defaultLiveGameRpcInnerService) AddTransferGameBetRecord(ctx context.Context, in *AddTransferGameBetRecordReq, opts ...grpc.CallOption) (*GameReply, error) {
	client := v1.NewLiveGameRpcInnerServiceClient(m.cli.Conn())
	return client.AddTransferGameBetRecord(ctx, in, opts...)
}

// 发送游戏下注MQ
func (m *defaultLiveGameRpcInnerService) SendGameBetBetMQ(ctx context.Context, in *SendGameBetBetMQReq, opts ...grpc.CallOption) (*GameReply, error) {
	client := v1.NewLiveGameRpcInnerServiceClient(m.cli.Conn())
	return client.SendGameBetBetMQ(ctx, in, opts...)
}

// 发送游戏下注结算MQ
func (m *defaultLiveGameRpcInnerService) SendGameBetBetSettlementMQ(ctx context.Context, in *SendGameBetBetSettlementMQReq, opts ...grpc.CallOption) (*GameReply, error) {
	client := v1.NewLiveGameRpcInnerServiceClient(m.cli.Conn())
	return client.SendGameBetBetSettlementMQ(ctx, in, opts...)
}

// 添加最近游玩的游戏
func (m *defaultLiveGameRpcInnerService) AddRecentlyGamePlay(ctx context.Context, in *AddRecentlyGamePlayReq, opts ...grpc.CallOption) (*GameReply, error) {
	client := v1.NewLiveGameRpcInnerServiceClient(m.cli.Conn())
	return client.AddRecentlyGamePlay(ctx, in, opts...)
}
