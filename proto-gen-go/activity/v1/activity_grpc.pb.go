// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: activity/v1/activity.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	LiveActivityInnerService_RedPacketCampaignEvent_FullMethodName      = "/activity.v1.LiveActivityInnerService/RedPacketCampaignEvent"
	LiveActivityInnerService_CheckUserRedPacketCondition_FullMethodName = "/activity.v1.LiveActivityInnerService/CheckUserRedPacketCondition"
	LiveActivityInnerService_IncreaseUserRedPacketCount_FullMethodName  = "/activity.v1.LiveActivityInnerService/IncreaseUserRedPacketCount"
	LiveActivityInnerService_LuckySpinEvent_FullMethodName              = "/activity.v1.LiveActivityInnerService/LuckySpinEvent"
	LiveActivityInnerService_AddLuckyPoint_FullMethodName               = "/activity.v1.LiveActivityInnerService/AddLuckyPoint"
	LiveActivityInnerService_UseLuckyPoint_FullMethodName               = "/activity.v1.LiveActivityInnerService/UseLuckyPoint"
	LiveActivityInnerService_GetUserLuckyPoint_FullMethodName           = "/activity.v1.LiveActivityInnerService/GetUserLuckyPoint"
	LiveActivityInnerService_LuckyPointsAddList_FullMethodName          = "/activity.v1.LiveActivityInnerService/LuckyPointsAddList"
	LiveActivityInnerService_LuckyPointsUsedList_FullMethodName         = "/activity.v1.LiveActivityInnerService/LuckyPointsUsedList"
	LiveActivityInnerService_RewardList_FullMethodName                  = "/activity.v1.LiveActivityInnerService/RewardList"
	LiveActivityInnerService_GetUserInviteCount_FullMethodName          = "/activity.v1.LiveActivityInnerService/GetUserInviteCount"
)

// LiveActivityInnerServiceClient is the client API for LiveActivityInnerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LiveActivityInnerServiceClient interface {
	// 抢红包活动事件
	RedPacketCampaignEvent(ctx context.Context, in *RedPacketCampaignEventReq, opts ...grpc.CallOption) (*ActivityReply, error)
	// 判断用户抢红包条件
	CheckUserRedPacketCondition(ctx context.Context, in *CheckUserRedPacketConditionReq, opts ...grpc.CallOption) (*CheckUserRedPacketConditionReply, error)
	// 增加用户领取红包次数缓存
	IncreaseUserRedPacketCount(ctx context.Context, in *IncreaseUserRedPacketCountReq, opts ...grpc.CallOption) (*ActivityReply, error)
	// 幸运转盘活动事件
	LuckySpinEvent(ctx context.Context, in *LuckySpinEventReq, opts ...grpc.CallOption) (*ActivityReply, error)
	// 增加幸运值
	AddLuckyPoint(ctx context.Context, in *AddLuckyValReq, opts ...grpc.CallOption) (*ActivityReply, error)
	// 使用幸运值
	UseLuckyPoint(ctx context.Context, in *UseLuckyPointReq, opts ...grpc.CallOption) (*ActivityReply, error)
	// 我的幸运值
	GetUserLuckyPoint(ctx context.Context, in *GetLuckyPointReq, opts ...grpc.CallOption) (*GetLuckyPointReply, error)
	// 幸运值获取记录
	LuckyPointsAddList(ctx context.Context, in *GetLuckyPointListReq, opts ...grpc.CallOption) (*LuckyPointsAddListReply, error)
	// 幸运值消费记录
	LuckyPointsUsedList(ctx context.Context, in *GetLuckyPointListReq, opts ...grpc.CallOption) (*LuckyPointsUsedListReply, error)
	// 获奖记录 公告展示
	RewardList(ctx context.Context, in *RewardListReq, opts ...grpc.CallOption) (*RewardListReply, error)
	// 获取用户当前周期内有效邀请人数
	GetUserInviteCount(ctx context.Context, in *GetUserInviteCountReq, opts ...grpc.CallOption) (*GetUserInviteCountReply, error)
}

type liveActivityInnerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLiveActivityInnerServiceClient(cc grpc.ClientConnInterface) LiveActivityInnerServiceClient {
	return &liveActivityInnerServiceClient{cc}
}

func (c *liveActivityInnerServiceClient) RedPacketCampaignEvent(ctx context.Context, in *RedPacketCampaignEventReq, opts ...grpc.CallOption) (*ActivityReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ActivityReply)
	err := c.cc.Invoke(ctx, LiveActivityInnerService_RedPacketCampaignEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveActivityInnerServiceClient) CheckUserRedPacketCondition(ctx context.Context, in *CheckUserRedPacketConditionReq, opts ...grpc.CallOption) (*CheckUserRedPacketConditionReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckUserRedPacketConditionReply)
	err := c.cc.Invoke(ctx, LiveActivityInnerService_CheckUserRedPacketCondition_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveActivityInnerServiceClient) IncreaseUserRedPacketCount(ctx context.Context, in *IncreaseUserRedPacketCountReq, opts ...grpc.CallOption) (*ActivityReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ActivityReply)
	err := c.cc.Invoke(ctx, LiveActivityInnerService_IncreaseUserRedPacketCount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveActivityInnerServiceClient) LuckySpinEvent(ctx context.Context, in *LuckySpinEventReq, opts ...grpc.CallOption) (*ActivityReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ActivityReply)
	err := c.cc.Invoke(ctx, LiveActivityInnerService_LuckySpinEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveActivityInnerServiceClient) AddLuckyPoint(ctx context.Context, in *AddLuckyValReq, opts ...grpc.CallOption) (*ActivityReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ActivityReply)
	err := c.cc.Invoke(ctx, LiveActivityInnerService_AddLuckyPoint_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveActivityInnerServiceClient) UseLuckyPoint(ctx context.Context, in *UseLuckyPointReq, opts ...grpc.CallOption) (*ActivityReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ActivityReply)
	err := c.cc.Invoke(ctx, LiveActivityInnerService_UseLuckyPoint_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveActivityInnerServiceClient) GetUserLuckyPoint(ctx context.Context, in *GetLuckyPointReq, opts ...grpc.CallOption) (*GetLuckyPointReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLuckyPointReply)
	err := c.cc.Invoke(ctx, LiveActivityInnerService_GetUserLuckyPoint_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveActivityInnerServiceClient) LuckyPointsAddList(ctx context.Context, in *GetLuckyPointListReq, opts ...grpc.CallOption) (*LuckyPointsAddListReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LuckyPointsAddListReply)
	err := c.cc.Invoke(ctx, LiveActivityInnerService_LuckyPointsAddList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveActivityInnerServiceClient) LuckyPointsUsedList(ctx context.Context, in *GetLuckyPointListReq, opts ...grpc.CallOption) (*LuckyPointsUsedListReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LuckyPointsUsedListReply)
	err := c.cc.Invoke(ctx, LiveActivityInnerService_LuckyPointsUsedList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveActivityInnerServiceClient) RewardList(ctx context.Context, in *RewardListReq, opts ...grpc.CallOption) (*RewardListReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RewardListReply)
	err := c.cc.Invoke(ctx, LiveActivityInnerService_RewardList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveActivityInnerServiceClient) GetUserInviteCount(ctx context.Context, in *GetUserInviteCountReq, opts ...grpc.CallOption) (*GetUserInviteCountReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserInviteCountReply)
	err := c.cc.Invoke(ctx, LiveActivityInnerService_GetUserInviteCount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LiveActivityInnerServiceServer is the server API for LiveActivityInnerService service.
// All implementations must embed UnimplementedLiveActivityInnerServiceServer
// for forward compatibility.
type LiveActivityInnerServiceServer interface {
	// 抢红包活动事件
	RedPacketCampaignEvent(context.Context, *RedPacketCampaignEventReq) (*ActivityReply, error)
	// 判断用户抢红包条件
	CheckUserRedPacketCondition(context.Context, *CheckUserRedPacketConditionReq) (*CheckUserRedPacketConditionReply, error)
	// 增加用户领取红包次数缓存
	IncreaseUserRedPacketCount(context.Context, *IncreaseUserRedPacketCountReq) (*ActivityReply, error)
	// 幸运转盘活动事件
	LuckySpinEvent(context.Context, *LuckySpinEventReq) (*ActivityReply, error)
	// 增加幸运值
	AddLuckyPoint(context.Context, *AddLuckyValReq) (*ActivityReply, error)
	// 使用幸运值
	UseLuckyPoint(context.Context, *UseLuckyPointReq) (*ActivityReply, error)
	// 我的幸运值
	GetUserLuckyPoint(context.Context, *GetLuckyPointReq) (*GetLuckyPointReply, error)
	// 幸运值获取记录
	LuckyPointsAddList(context.Context, *GetLuckyPointListReq) (*LuckyPointsAddListReply, error)
	// 幸运值消费记录
	LuckyPointsUsedList(context.Context, *GetLuckyPointListReq) (*LuckyPointsUsedListReply, error)
	// 获奖记录 公告展示
	RewardList(context.Context, *RewardListReq) (*RewardListReply, error)
	// 获取用户当前周期内有效邀请人数
	GetUserInviteCount(context.Context, *GetUserInviteCountReq) (*GetUserInviteCountReply, error)
	mustEmbedUnimplementedLiveActivityInnerServiceServer()
}

// UnimplementedLiveActivityInnerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLiveActivityInnerServiceServer struct{}

func (UnimplementedLiveActivityInnerServiceServer) RedPacketCampaignEvent(context.Context, *RedPacketCampaignEventReq) (*ActivityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RedPacketCampaignEvent not implemented")
}
func (UnimplementedLiveActivityInnerServiceServer) CheckUserRedPacketCondition(context.Context, *CheckUserRedPacketConditionReq) (*CheckUserRedPacketConditionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserRedPacketCondition not implemented")
}
func (UnimplementedLiveActivityInnerServiceServer) IncreaseUserRedPacketCount(context.Context, *IncreaseUserRedPacketCountReq) (*ActivityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncreaseUserRedPacketCount not implemented")
}
func (UnimplementedLiveActivityInnerServiceServer) LuckySpinEvent(context.Context, *LuckySpinEventReq) (*ActivityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LuckySpinEvent not implemented")
}
func (UnimplementedLiveActivityInnerServiceServer) AddLuckyPoint(context.Context, *AddLuckyValReq) (*ActivityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLuckyPoint not implemented")
}
func (UnimplementedLiveActivityInnerServiceServer) UseLuckyPoint(context.Context, *UseLuckyPointReq) (*ActivityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UseLuckyPoint not implemented")
}
func (UnimplementedLiveActivityInnerServiceServer) GetUserLuckyPoint(context.Context, *GetLuckyPointReq) (*GetLuckyPointReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserLuckyPoint not implemented")
}
func (UnimplementedLiveActivityInnerServiceServer) LuckyPointsAddList(context.Context, *GetLuckyPointListReq) (*LuckyPointsAddListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LuckyPointsAddList not implemented")
}
func (UnimplementedLiveActivityInnerServiceServer) LuckyPointsUsedList(context.Context, *GetLuckyPointListReq) (*LuckyPointsUsedListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LuckyPointsUsedList not implemented")
}
func (UnimplementedLiveActivityInnerServiceServer) RewardList(context.Context, *RewardListReq) (*RewardListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RewardList not implemented")
}
func (UnimplementedLiveActivityInnerServiceServer) GetUserInviteCount(context.Context, *GetUserInviteCountReq) (*GetUserInviteCountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInviteCount not implemented")
}
func (UnimplementedLiveActivityInnerServiceServer) mustEmbedUnimplementedLiveActivityInnerServiceServer() {
}
func (UnimplementedLiveActivityInnerServiceServer) testEmbeddedByValue() {}

// UnsafeLiveActivityInnerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LiveActivityInnerServiceServer will
// result in compilation errors.
type UnsafeLiveActivityInnerServiceServer interface {
	mustEmbedUnimplementedLiveActivityInnerServiceServer()
}

func RegisterLiveActivityInnerServiceServer(s grpc.ServiceRegistrar, srv LiveActivityInnerServiceServer) {
	// If the following call pancis, it indicates UnimplementedLiveActivityInnerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LiveActivityInnerService_ServiceDesc, srv)
}

func _LiveActivityInnerService_RedPacketCampaignEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RedPacketCampaignEventReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveActivityInnerServiceServer).RedPacketCampaignEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveActivityInnerService_RedPacketCampaignEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveActivityInnerServiceServer).RedPacketCampaignEvent(ctx, req.(*RedPacketCampaignEventReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveActivityInnerService_CheckUserRedPacketCondition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserRedPacketConditionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveActivityInnerServiceServer).CheckUserRedPacketCondition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveActivityInnerService_CheckUserRedPacketCondition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveActivityInnerServiceServer).CheckUserRedPacketCondition(ctx, req.(*CheckUserRedPacketConditionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveActivityInnerService_IncreaseUserRedPacketCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncreaseUserRedPacketCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveActivityInnerServiceServer).IncreaseUserRedPacketCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveActivityInnerService_IncreaseUserRedPacketCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveActivityInnerServiceServer).IncreaseUserRedPacketCount(ctx, req.(*IncreaseUserRedPacketCountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveActivityInnerService_LuckySpinEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LuckySpinEventReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveActivityInnerServiceServer).LuckySpinEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveActivityInnerService_LuckySpinEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveActivityInnerServiceServer).LuckySpinEvent(ctx, req.(*LuckySpinEventReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveActivityInnerService_AddLuckyPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddLuckyValReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveActivityInnerServiceServer).AddLuckyPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveActivityInnerService_AddLuckyPoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveActivityInnerServiceServer).AddLuckyPoint(ctx, req.(*AddLuckyValReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveActivityInnerService_UseLuckyPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UseLuckyPointReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveActivityInnerServiceServer).UseLuckyPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveActivityInnerService_UseLuckyPoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveActivityInnerServiceServer).UseLuckyPoint(ctx, req.(*UseLuckyPointReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveActivityInnerService_GetUserLuckyPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLuckyPointReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveActivityInnerServiceServer).GetUserLuckyPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveActivityInnerService_GetUserLuckyPoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveActivityInnerServiceServer).GetUserLuckyPoint(ctx, req.(*GetLuckyPointReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveActivityInnerService_LuckyPointsAddList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLuckyPointListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveActivityInnerServiceServer).LuckyPointsAddList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveActivityInnerService_LuckyPointsAddList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveActivityInnerServiceServer).LuckyPointsAddList(ctx, req.(*GetLuckyPointListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveActivityInnerService_LuckyPointsUsedList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLuckyPointListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveActivityInnerServiceServer).LuckyPointsUsedList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveActivityInnerService_LuckyPointsUsedList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveActivityInnerServiceServer).LuckyPointsUsedList(ctx, req.(*GetLuckyPointListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveActivityInnerService_RewardList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RewardListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveActivityInnerServiceServer).RewardList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveActivityInnerService_RewardList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveActivityInnerServiceServer).RewardList(ctx, req.(*RewardListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveActivityInnerService_GetUserInviteCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInviteCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveActivityInnerServiceServer).GetUserInviteCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveActivityInnerService_GetUserInviteCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveActivityInnerServiceServer).GetUserInviteCount(ctx, req.(*GetUserInviteCountReq))
	}
	return interceptor(ctx, in, info, handler)
}

// LiveActivityInnerService_ServiceDesc is the grpc.ServiceDesc for LiveActivityInnerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LiveActivityInnerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "activity.v1.LiveActivityInnerService",
	HandlerType: (*LiveActivityInnerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RedPacketCampaignEvent",
			Handler:    _LiveActivityInnerService_RedPacketCampaignEvent_Handler,
		},
		{
			MethodName: "CheckUserRedPacketCondition",
			Handler:    _LiveActivityInnerService_CheckUserRedPacketCondition_Handler,
		},
		{
			MethodName: "IncreaseUserRedPacketCount",
			Handler:    _LiveActivityInnerService_IncreaseUserRedPacketCount_Handler,
		},
		{
			MethodName: "LuckySpinEvent",
			Handler:    _LiveActivityInnerService_LuckySpinEvent_Handler,
		},
		{
			MethodName: "AddLuckyPoint",
			Handler:    _LiveActivityInnerService_AddLuckyPoint_Handler,
		},
		{
			MethodName: "UseLuckyPoint",
			Handler:    _LiveActivityInnerService_UseLuckyPoint_Handler,
		},
		{
			MethodName: "GetUserLuckyPoint",
			Handler:    _LiveActivityInnerService_GetUserLuckyPoint_Handler,
		},
		{
			MethodName: "LuckyPointsAddList",
			Handler:    _LiveActivityInnerService_LuckyPointsAddList_Handler,
		},
		{
			MethodName: "LuckyPointsUsedList",
			Handler:    _LiveActivityInnerService_LuckyPointsUsedList_Handler,
		},
		{
			MethodName: "RewardList",
			Handler:    _LiveActivityInnerService_RewardList_Handler,
		},
		{
			MethodName: "GetUserInviteCount",
			Handler:    _LiveActivityInnerService_GetUserInviteCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "activity/v1/activity.proto",
}
