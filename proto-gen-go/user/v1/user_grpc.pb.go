// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.3
// source: user/v1/user.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	LiveUserRpcService_GetInfoByUserToken_FullMethodName    = "/user.v1.LiveUserRpcService/GetInfoByUserToken"
	LiveUserRpcService_GetUserBalance_FullMethodName        = "/user.v1.LiveUserRpcService/GetUserBalance"
	LiveUserRpcService_UpdateUserBalance_FullMethodName     = "/user.v1.LiveUserRpcService/UpdateUserBalance"
	LiveUserRpcService_GetUserInfoById_FullMethodName       = "/user.v1.LiveUserRpcService/GetUserInfoById"
	LiveUserRpcService_GetUserFullInfoById_FullMethodName   = "/user.v1.LiveUserRpcService/GetUserFullInfoById"
	LiveUserRpcService_GetUserFullMapInfo_FullMethodName    = "/user.v1.LiveUserRpcService/GetUserFullMapInfo"
	LiveUserRpcService_BatchGetUserBalance_FullMethodName   = "/user.v1.LiveUserRpcService/BatchGetUserBalance"
	LiveUserRpcService_BatchGetUserBalanceV2_FullMethodName = "/user.v1.LiveUserRpcService/BatchGetUserBalanceV2"
)

// LiveUserRpcServiceClient is the client API for LiveUserRpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LiveUserRpcServiceClient interface {
	// 根据token获取用户信息
	GetInfoByUserToken(ctx context.Context, in *GetInfoByUserTokenReq, opts ...grpc.CallOption) (*UserDetailsInfoReply, error)
	// 获取用户余额
	GetUserBalance(ctx context.Context, in *GetUserBalanceReq, opts ...grpc.CallOption) (*GetUserBalanceReply, error)
	// 修改用户余额
	UpdateUserBalance(ctx context.Context, in *UpdateUserBalanceReq, opts ...grpc.CallOption) (*UpdateUserBalanceResp, error)
	// 根据用户id获取部分用户信息
	GetUserInfoById(ctx context.Context, in *GetUserInfoByIdReq, opts ...grpc.CallOption) (*UserDetailsInfoReply, error)
	// 获取用户详细信息（全部-需要其他用户字段就在reply里面加）
	GetUserFullInfoById(ctx context.Context, in *GetUserFullInfoByIdReq, opts ...grpc.CallOption) (*GetUserFullInfoByIdReply, error)
	// 获取用户信息map
	GetUserFullMapInfo(ctx context.Context, in *GetUserFullMapInfoReq, opts ...grpc.CallOption) (*GetUserFullMapInfoReply, error)
	// 批量查询用户余额
	BatchGetUserBalance(ctx context.Context, in *BatchGetUserBalanceReq, opts ...grpc.CallOption) (*BatchGetUserBalanceResp, error)
	// 获取多用户多钱包余额接口
	BatchGetUserBalanceV2(ctx context.Context, in *BatchGetUserBalanceV2Req, opts ...grpc.CallOption) (*BatchGetUserBalanceV2Resp, error)
}

type liveUserRpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLiveUserRpcServiceClient(cc grpc.ClientConnInterface) LiveUserRpcServiceClient {
	return &liveUserRpcServiceClient{cc}
}

func (c *liveUserRpcServiceClient) GetInfoByUserToken(ctx context.Context, in *GetInfoByUserTokenReq, opts ...grpc.CallOption) (*UserDetailsInfoReply, error) {
	out := new(UserDetailsInfoReply)
	err := c.cc.Invoke(ctx, LiveUserRpcService_GetInfoByUserToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveUserRpcServiceClient) GetUserBalance(ctx context.Context, in *GetUserBalanceReq, opts ...grpc.CallOption) (*GetUserBalanceReply, error) {
	out := new(GetUserBalanceReply)
	err := c.cc.Invoke(ctx, LiveUserRpcService_GetUserBalance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveUserRpcServiceClient) UpdateUserBalance(ctx context.Context, in *UpdateUserBalanceReq, opts ...grpc.CallOption) (*UpdateUserBalanceResp, error) {
	out := new(UpdateUserBalanceResp)
	err := c.cc.Invoke(ctx, LiveUserRpcService_UpdateUserBalance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveUserRpcServiceClient) GetUserInfoById(ctx context.Context, in *GetUserInfoByIdReq, opts ...grpc.CallOption) (*UserDetailsInfoReply, error) {
	out := new(UserDetailsInfoReply)
	err := c.cc.Invoke(ctx, LiveUserRpcService_GetUserInfoById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveUserRpcServiceClient) GetUserFullInfoById(ctx context.Context, in *GetUserFullInfoByIdReq, opts ...grpc.CallOption) (*GetUserFullInfoByIdReply, error) {
	out := new(GetUserFullInfoByIdReply)
	err := c.cc.Invoke(ctx, LiveUserRpcService_GetUserFullInfoById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveUserRpcServiceClient) GetUserFullMapInfo(ctx context.Context, in *GetUserFullMapInfoReq, opts ...grpc.CallOption) (*GetUserFullMapInfoReply, error) {
	out := new(GetUserFullMapInfoReply)
	err := c.cc.Invoke(ctx, LiveUserRpcService_GetUserFullMapInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveUserRpcServiceClient) BatchGetUserBalance(ctx context.Context, in *BatchGetUserBalanceReq, opts ...grpc.CallOption) (*BatchGetUserBalanceResp, error) {
	out := new(BatchGetUserBalanceResp)
	err := c.cc.Invoke(ctx, LiveUserRpcService_BatchGetUserBalance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveUserRpcServiceClient) BatchGetUserBalanceV2(ctx context.Context, in *BatchGetUserBalanceV2Req, opts ...grpc.CallOption) (*BatchGetUserBalanceV2Resp, error) {
	out := new(BatchGetUserBalanceV2Resp)
	err := c.cc.Invoke(ctx, LiveUserRpcService_BatchGetUserBalanceV2_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LiveUserRpcServiceServer is the server API for LiveUserRpcService service.
// All implementations must embed UnimplementedLiveUserRpcServiceServer
// for forward compatibility
type LiveUserRpcServiceServer interface {
	// 根据token获取用户信息
	GetInfoByUserToken(context.Context, *GetInfoByUserTokenReq) (*UserDetailsInfoReply, error)
	// 获取用户余额
	GetUserBalance(context.Context, *GetUserBalanceReq) (*GetUserBalanceReply, error)
	// 修改用户余额
	UpdateUserBalance(context.Context, *UpdateUserBalanceReq) (*UpdateUserBalanceResp, error)
	// 根据用户id获取部分用户信息
	GetUserInfoById(context.Context, *GetUserInfoByIdReq) (*UserDetailsInfoReply, error)
	// 获取用户详细信息（全部-需要其他用户字段就在reply里面加）
	GetUserFullInfoById(context.Context, *GetUserFullInfoByIdReq) (*GetUserFullInfoByIdReply, error)
	// 获取用户信息map
	GetUserFullMapInfo(context.Context, *GetUserFullMapInfoReq) (*GetUserFullMapInfoReply, error)
	// 批量查询用户余额
	BatchGetUserBalance(context.Context, *BatchGetUserBalanceReq) (*BatchGetUserBalanceResp, error)
	// 获取多用户多钱包余额接口
	BatchGetUserBalanceV2(context.Context, *BatchGetUserBalanceV2Req) (*BatchGetUserBalanceV2Resp, error)
	mustEmbedUnimplementedLiveUserRpcServiceServer()
}

// UnimplementedLiveUserRpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLiveUserRpcServiceServer struct {
}

func (UnimplementedLiveUserRpcServiceServer) GetInfoByUserToken(context.Context, *GetInfoByUserTokenReq) (*UserDetailsInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfoByUserToken not implemented")
}
func (UnimplementedLiveUserRpcServiceServer) GetUserBalance(context.Context, *GetUserBalanceReq) (*GetUserBalanceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserBalance not implemented")
}
func (UnimplementedLiveUserRpcServiceServer) UpdateUserBalance(context.Context, *UpdateUserBalanceReq) (*UpdateUserBalanceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserBalance not implemented")
}
func (UnimplementedLiveUserRpcServiceServer) GetUserInfoById(context.Context, *GetUserInfoByIdReq) (*UserDetailsInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfoById not implemented")
}
func (UnimplementedLiveUserRpcServiceServer) GetUserFullInfoById(context.Context, *GetUserFullInfoByIdReq) (*GetUserFullInfoByIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFullInfoById not implemented")
}
func (UnimplementedLiveUserRpcServiceServer) GetUserFullMapInfo(context.Context, *GetUserFullMapInfoReq) (*GetUserFullMapInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFullMapInfo not implemented")
}
func (UnimplementedLiveUserRpcServiceServer) BatchGetUserBalance(context.Context, *BatchGetUserBalanceReq) (*BatchGetUserBalanceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGetUserBalance not implemented")
}
func (UnimplementedLiveUserRpcServiceServer) BatchGetUserBalanceV2(context.Context, *BatchGetUserBalanceV2Req) (*BatchGetUserBalanceV2Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGetUserBalanceV2 not implemented")
}
func (UnimplementedLiveUserRpcServiceServer) mustEmbedUnimplementedLiveUserRpcServiceServer() {}

// UnsafeLiveUserRpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LiveUserRpcServiceServer will
// result in compilation errors.
type UnsafeLiveUserRpcServiceServer interface {
	mustEmbedUnimplementedLiveUserRpcServiceServer()
}

func RegisterLiveUserRpcServiceServer(s grpc.ServiceRegistrar, srv LiveUserRpcServiceServer) {
	s.RegisterService(&LiveUserRpcService_ServiceDesc, srv)
}

func _LiveUserRpcService_GetInfoByUserToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoByUserTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveUserRpcServiceServer).GetInfoByUserToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveUserRpcService_GetInfoByUserToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveUserRpcServiceServer).GetInfoByUserToken(ctx, req.(*GetInfoByUserTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveUserRpcService_GetUserBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserBalanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveUserRpcServiceServer).GetUserBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveUserRpcService_GetUserBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveUserRpcServiceServer).GetUserBalance(ctx, req.(*GetUserBalanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveUserRpcService_UpdateUserBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserBalanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveUserRpcServiceServer).UpdateUserBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveUserRpcService_UpdateUserBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveUserRpcServiceServer).UpdateUserBalance(ctx, req.(*UpdateUserBalanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveUserRpcService_GetUserInfoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveUserRpcServiceServer).GetUserInfoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveUserRpcService_GetUserInfoById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveUserRpcServiceServer).GetUserInfoById(ctx, req.(*GetUserInfoByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveUserRpcService_GetUserFullInfoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserFullInfoByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveUserRpcServiceServer).GetUserFullInfoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveUserRpcService_GetUserFullInfoById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveUserRpcServiceServer).GetUserFullInfoById(ctx, req.(*GetUserFullInfoByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveUserRpcService_GetUserFullMapInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserFullMapInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveUserRpcServiceServer).GetUserFullMapInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveUserRpcService_GetUserFullMapInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveUserRpcServiceServer).GetUserFullMapInfo(ctx, req.(*GetUserFullMapInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveUserRpcService_BatchGetUserBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetUserBalanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveUserRpcServiceServer).BatchGetUserBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveUserRpcService_BatchGetUserBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveUserRpcServiceServer).BatchGetUserBalance(ctx, req.(*BatchGetUserBalanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveUserRpcService_BatchGetUserBalanceV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetUserBalanceV2Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveUserRpcServiceServer).BatchGetUserBalanceV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveUserRpcService_BatchGetUserBalanceV2_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveUserRpcServiceServer).BatchGetUserBalanceV2(ctx, req.(*BatchGetUserBalanceV2Req))
	}
	return interceptor(ctx, in, info, handler)
}

// LiveUserRpcService_ServiceDesc is the grpc.ServiceDesc for LiveUserRpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LiveUserRpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.v1.LiveUserRpcService",
	HandlerType: (*LiveUserRpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInfoByUserToken",
			Handler:    _LiveUserRpcService_GetInfoByUserToken_Handler,
		},
		{
			MethodName: "GetUserBalance",
			Handler:    _LiveUserRpcService_GetUserBalance_Handler,
		},
		{
			MethodName: "UpdateUserBalance",
			Handler:    _LiveUserRpcService_UpdateUserBalance_Handler,
		},
		{
			MethodName: "GetUserInfoById",
			Handler:    _LiveUserRpcService_GetUserInfoById_Handler,
		},
		{
			MethodName: "GetUserFullInfoById",
			Handler:    _LiveUserRpcService_GetUserFullInfoById_Handler,
		},
		{
			MethodName: "GetUserFullMapInfo",
			Handler:    _LiveUserRpcService_GetUserFullMapInfo_Handler,
		},
		{
			MethodName: "BatchGetUserBalance",
			Handler:    _LiveUserRpcService_BatchGetUserBalance_Handler,
		},
		{
			MethodName: "BatchGetUserBalanceV2",
			Handler:    _LiveUserRpcService_BatchGetUserBalanceV2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/v1/user.proto",
}
