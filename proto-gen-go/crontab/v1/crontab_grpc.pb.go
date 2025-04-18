// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.29.3
// source: crontab/v1/crontab.proto

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
	LiveCrontabRpcService_Add_FullMethodName = "/user.v1.LiveCrontabRpcService/Add"
	LiveCrontabRpcService_Del_FullMethodName = "/user.v1.LiveCrontabRpcService/Del"
)

// LiveCrontabRpcServiceClient is the client API for LiveCrontabRpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LiveCrontabRpcServiceClient interface {
	// 添加定时任务
	Add(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*AddRes, error)
	// 删除定时任务
	Del(ctx context.Context, in *DelReq, opts ...grpc.CallOption) (*EmptyRes, error)
}

type liveCrontabRpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLiveCrontabRpcServiceClient(cc grpc.ClientConnInterface) LiveCrontabRpcServiceClient {
	return &liveCrontabRpcServiceClient{cc}
}

func (c *liveCrontabRpcServiceClient) Add(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*AddRes, error) {
	out := new(AddRes)
	err := c.cc.Invoke(ctx, LiveCrontabRpcService_Add_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveCrontabRpcServiceClient) Del(ctx context.Context, in *DelReq, opts ...grpc.CallOption) (*EmptyRes, error) {
	out := new(EmptyRes)
	err := c.cc.Invoke(ctx, LiveCrontabRpcService_Del_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LiveCrontabRpcServiceServer is the server API for LiveCrontabRpcService service.
// All implementations must embed UnimplementedLiveCrontabRpcServiceServer
// for forward compatibility
type LiveCrontabRpcServiceServer interface {
	// 添加定时任务
	Add(context.Context, *AddReq) (*AddRes, error)
	// 删除定时任务
	Del(context.Context, *DelReq) (*EmptyRes, error)
	mustEmbedUnimplementedLiveCrontabRpcServiceServer()
}

// UnimplementedLiveCrontabRpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLiveCrontabRpcServiceServer struct {
}

func (UnimplementedLiveCrontabRpcServiceServer) Add(context.Context, *AddReq) (*AddRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedLiveCrontabRpcServiceServer) Del(context.Context, *DelReq) (*EmptyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Del not implemented")
}
func (UnimplementedLiveCrontabRpcServiceServer) mustEmbedUnimplementedLiveCrontabRpcServiceServer() {}

// UnsafeLiveCrontabRpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LiveCrontabRpcServiceServer will
// result in compilation errors.
type UnsafeLiveCrontabRpcServiceServer interface {
	mustEmbedUnimplementedLiveCrontabRpcServiceServer()
}

func RegisterLiveCrontabRpcServiceServer(s grpc.ServiceRegistrar, srv LiveCrontabRpcServiceServer) {
	s.RegisterService(&LiveCrontabRpcService_ServiceDesc, srv)
}

func _LiveCrontabRpcService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveCrontabRpcServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveCrontabRpcService_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveCrontabRpcServiceServer).Add(ctx, req.(*AddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveCrontabRpcService_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveCrontabRpcServiceServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveCrontabRpcService_Del_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveCrontabRpcServiceServer).Del(ctx, req.(*DelReq))
	}
	return interceptor(ctx, in, info, handler)
}

// LiveCrontabRpcService_ServiceDesc is the grpc.ServiceDesc for LiveCrontabRpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LiveCrontabRpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.v1.LiveCrontabRpcService",
	HandlerType: (*LiveCrontabRpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _LiveCrontabRpcService_Add_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _LiveCrontabRpcService_Del_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crontab/v1/crontab.proto",
}
