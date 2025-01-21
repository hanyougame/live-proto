// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
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
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	LiveUserRpcService_GetInfoByUserToken_FullMethodName = "/user.v1.LiveUserRpcService/GetInfoByUserToken"
	LiveUserRpcService_GetUserBalance_FullMethodName     = "/user.v1.LiveUserRpcService/GetUserBalance"
	LiveUserRpcService_UpdateUserBalance_FullMethodName  = "/user.v1.LiveUserRpcService/UpdateUserBalance"
	LiveUserRpcService_GetUserInfoById_FullMethodName    = "/user.v1.LiveUserRpcService/GetUserInfoById"
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
	GetUserInfoById(ctx context.Context, in *GetUserInfoByIdReq, opts ...grpc.CallOption) (*UserDetailsInfoReply, error)
}

type liveUserRpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLiveUserRpcServiceClient(cc grpc.ClientConnInterface) LiveUserRpcServiceClient {
	return &liveUserRpcServiceClient{cc}
}

func (c *liveUserRpcServiceClient) GetInfoByUserToken(ctx context.Context, in *GetInfoByUserTokenReq, opts ...grpc.CallOption) (*UserDetailsInfoReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserDetailsInfoReply)
	err := c.cc.Invoke(ctx, LiveUserRpcService_GetInfoByUserToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveUserRpcServiceClient) GetUserBalance(ctx context.Context, in *GetUserBalanceReq, opts ...grpc.CallOption) (*GetUserBalanceReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserBalanceReply)
	err := c.cc.Invoke(ctx, LiveUserRpcService_GetUserBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveUserRpcServiceClient) UpdateUserBalance(ctx context.Context, in *UpdateUserBalanceReq, opts ...grpc.CallOption) (*UpdateUserBalanceResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUserBalanceResp)
	err := c.cc.Invoke(ctx, LiveUserRpcService_UpdateUserBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveUserRpcServiceClient) GetUserInfoById(ctx context.Context, in *GetUserInfoByIdReq, opts ...grpc.CallOption) (*UserDetailsInfoReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserDetailsInfoReply)
	err := c.cc.Invoke(ctx, LiveUserRpcService_GetUserInfoById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LiveUserRpcServiceServer is the server API for LiveUserRpcService service.
// All implementations must embed UnimplementedLiveUserRpcServiceServer
// for forward compatibility.
type LiveUserRpcServiceServer interface {
	// 根据token获取用户信息
	GetInfoByUserToken(context.Context, *GetInfoByUserTokenReq) (*UserDetailsInfoReply, error)
	// 获取用户余额
	GetUserBalance(context.Context, *GetUserBalanceReq) (*GetUserBalanceReply, error)
	// 修改用户余额
	UpdateUserBalance(context.Context, *UpdateUserBalanceReq) (*UpdateUserBalanceResp, error)
	GetUserInfoById(context.Context, *GetUserInfoByIdReq) (*UserDetailsInfoReply, error)
	mustEmbedUnimplementedLiveUserRpcServiceServer()
}

// UnimplementedLiveUserRpcServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLiveUserRpcServiceServer struct{}

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
func (UnimplementedLiveUserRpcServiceServer) mustEmbedUnimplementedLiveUserRpcServiceServer() {}
func (UnimplementedLiveUserRpcServiceServer) testEmbeddedByValue()                            {}

// UnsafeLiveUserRpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LiveUserRpcServiceServer will
// result in compilation errors.
type UnsafeLiveUserRpcServiceServer interface {
	mustEmbedUnimplementedLiveUserRpcServiceServer()
}

func RegisterLiveUserRpcServiceServer(s grpc.ServiceRegistrar, srv LiveUserRpcServiceServer) {
	// If the following call pancis, it indicates UnimplementedLiveUserRpcServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/v1/user.proto",
}
