// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: manage/v1/manage.proto

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
	LiveManageRpcService_GetLegalTenderInfo_FullMethodName    = "/manage.v1.LiveManageRpcService/GetLegalTenderInfo"
	LiveManageRpcService_GetCurrInfoById_FullMethodName       = "/manage.v1.LiveManageRpcService/GetCurrInfoById"
	LiveManageRpcService_GetIPGeolocation_FullMethodName      = "/manage.v1.LiveManageRpcService/GetIPGeolocation"
	LiveManageRpcService_BatchGetIPGeolocation_FullMethodName = "/manage.v1.LiveManageRpcService/BatchGetIPGeolocation"
	LiveManageRpcService_SendSms_FullMethodName               = "/manage.v1.LiveManageRpcService/SendSms"
	LiveManageRpcService_SmsBalance_FullMethodName            = "/manage.v1.LiveManageRpcService/SmsBalance"
	LiveManageRpcService_SendEmail_FullMethodName             = "/manage.v1.LiveManageRpcService/SendEmail"
)

// LiveManageRpcServiceClient is the client API for LiveManageRpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LiveManageRpcServiceClient interface {
	// 获取法定货币信息
	GetLegalTenderInfo(ctx context.Context, in *ManageReq, opts ...grpc.CallOption) (*CurrencyDetailReplyList, error)
	// 通过货币ID获取货币信息
	GetCurrInfoById(ctx context.Context, in *GetCurrInfoByIDReq, opts ...grpc.CallOption) (*CurrencyDetailReply, error)
	// =============== IP相关接口 ===============
	// 获取单个IP地理位置信息
	GetIPGeolocation(ctx context.Context, in *GetIPGeolocationReq, opts ...grpc.CallOption) (*GetIPGeolocationReply, error)
	// 批量获取IP地理位置信息
	BatchGetIPGeolocation(ctx context.Context, in *BatchGetIPGeolocationReq, opts ...grpc.CallOption) (*BatchGetIPGeolocationReply, error)
	SendSms(ctx context.Context, in *SendSmsReq, opts ...grpc.CallOption) (*SendSmsReply, error)
	SmsBalance(ctx context.Context, in *SmsBalanceReq, opts ...grpc.CallOption) (*SmsBalanceReply, error)
	SendEmail(ctx context.Context, in *SendEmailReq, opts ...grpc.CallOption) (*SendEmailReply, error)
}

type liveManageRpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLiveManageRpcServiceClient(cc grpc.ClientConnInterface) LiveManageRpcServiceClient {
	return &liveManageRpcServiceClient{cc}
}

func (c *liveManageRpcServiceClient) GetLegalTenderInfo(ctx context.Context, in *ManageReq, opts ...grpc.CallOption) (*CurrencyDetailReplyList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CurrencyDetailReplyList)
	err := c.cc.Invoke(ctx, LiveManageRpcService_GetLegalTenderInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveManageRpcServiceClient) GetCurrInfoById(ctx context.Context, in *GetCurrInfoByIDReq, opts ...grpc.CallOption) (*CurrencyDetailReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CurrencyDetailReply)
	err := c.cc.Invoke(ctx, LiveManageRpcService_GetCurrInfoById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveManageRpcServiceClient) GetIPGeolocation(ctx context.Context, in *GetIPGeolocationReq, opts ...grpc.CallOption) (*GetIPGeolocationReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetIPGeolocationReply)
	err := c.cc.Invoke(ctx, LiveManageRpcService_GetIPGeolocation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveManageRpcServiceClient) BatchGetIPGeolocation(ctx context.Context, in *BatchGetIPGeolocationReq, opts ...grpc.CallOption) (*BatchGetIPGeolocationReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BatchGetIPGeolocationReply)
	err := c.cc.Invoke(ctx, LiveManageRpcService_BatchGetIPGeolocation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveManageRpcServiceClient) SendSms(ctx context.Context, in *SendSmsReq, opts ...grpc.CallOption) (*SendSmsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendSmsReply)
	err := c.cc.Invoke(ctx, LiveManageRpcService_SendSms_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveManageRpcServiceClient) SmsBalance(ctx context.Context, in *SmsBalanceReq, opts ...grpc.CallOption) (*SmsBalanceReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SmsBalanceReply)
	err := c.cc.Invoke(ctx, LiveManageRpcService_SmsBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveManageRpcServiceClient) SendEmail(ctx context.Context, in *SendEmailReq, opts ...grpc.CallOption) (*SendEmailReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendEmailReply)
	err := c.cc.Invoke(ctx, LiveManageRpcService_SendEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LiveManageRpcServiceServer is the server API for LiveManageRpcService service.
// All implementations must embed UnimplementedLiveManageRpcServiceServer
// for forward compatibility.
type LiveManageRpcServiceServer interface {
	// 获取法定货币信息
	GetLegalTenderInfo(context.Context, *ManageReq) (*CurrencyDetailReplyList, error)
	// 通过货币ID获取货币信息
	GetCurrInfoById(context.Context, *GetCurrInfoByIDReq) (*CurrencyDetailReply, error)
	// =============== IP相关接口 ===============
	// 获取单个IP地理位置信息
	GetIPGeolocation(context.Context, *GetIPGeolocationReq) (*GetIPGeolocationReply, error)
	// 批量获取IP地理位置信息
	BatchGetIPGeolocation(context.Context, *BatchGetIPGeolocationReq) (*BatchGetIPGeolocationReply, error)
	SendSms(context.Context, *SendSmsReq) (*SendSmsReply, error)
	SmsBalance(context.Context, *SmsBalanceReq) (*SmsBalanceReply, error)
	SendEmail(context.Context, *SendEmailReq) (*SendEmailReply, error)
	mustEmbedUnimplementedLiveManageRpcServiceServer()
}

// UnimplementedLiveManageRpcServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLiveManageRpcServiceServer struct{}

func (UnimplementedLiveManageRpcServiceServer) GetLegalTenderInfo(context.Context, *ManageReq) (*CurrencyDetailReplyList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLegalTenderInfo not implemented")
}
func (UnimplementedLiveManageRpcServiceServer) GetCurrInfoById(context.Context, *GetCurrInfoByIDReq) (*CurrencyDetailReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrInfoById not implemented")
}
func (UnimplementedLiveManageRpcServiceServer) GetIPGeolocation(context.Context, *GetIPGeolocationReq) (*GetIPGeolocationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIPGeolocation not implemented")
}
func (UnimplementedLiveManageRpcServiceServer) BatchGetIPGeolocation(context.Context, *BatchGetIPGeolocationReq) (*BatchGetIPGeolocationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGetIPGeolocation not implemented")
}
func (UnimplementedLiveManageRpcServiceServer) SendSms(context.Context, *SendSmsReq) (*SendSmsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSms not implemented")
}
func (UnimplementedLiveManageRpcServiceServer) SmsBalance(context.Context, *SmsBalanceReq) (*SmsBalanceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SmsBalance not implemented")
}
func (UnimplementedLiveManageRpcServiceServer) SendEmail(context.Context, *SendEmailReq) (*SendEmailReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmail not implemented")
}
func (UnimplementedLiveManageRpcServiceServer) mustEmbedUnimplementedLiveManageRpcServiceServer() {}
func (UnimplementedLiveManageRpcServiceServer) testEmbeddedByValue()                              {}

// UnsafeLiveManageRpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LiveManageRpcServiceServer will
// result in compilation errors.
type UnsafeLiveManageRpcServiceServer interface {
	mustEmbedUnimplementedLiveManageRpcServiceServer()
}

func RegisterLiveManageRpcServiceServer(s grpc.ServiceRegistrar, srv LiveManageRpcServiceServer) {
	// If the following call pancis, it indicates UnimplementedLiveManageRpcServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LiveManageRpcService_ServiceDesc, srv)
}

func _LiveManageRpcService_GetLegalTenderInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveManageRpcServiceServer).GetLegalTenderInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveManageRpcService_GetLegalTenderInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveManageRpcServiceServer).GetLegalTenderInfo(ctx, req.(*ManageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveManageRpcService_GetCurrInfoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrInfoByIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveManageRpcServiceServer).GetCurrInfoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveManageRpcService_GetCurrInfoById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveManageRpcServiceServer).GetCurrInfoById(ctx, req.(*GetCurrInfoByIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveManageRpcService_GetIPGeolocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIPGeolocationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveManageRpcServiceServer).GetIPGeolocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveManageRpcService_GetIPGeolocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveManageRpcServiceServer).GetIPGeolocation(ctx, req.(*GetIPGeolocationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveManageRpcService_BatchGetIPGeolocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetIPGeolocationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveManageRpcServiceServer).BatchGetIPGeolocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveManageRpcService_BatchGetIPGeolocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveManageRpcServiceServer).BatchGetIPGeolocation(ctx, req.(*BatchGetIPGeolocationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveManageRpcService_SendSms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSmsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveManageRpcServiceServer).SendSms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveManageRpcService_SendSms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveManageRpcServiceServer).SendSms(ctx, req.(*SendSmsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveManageRpcService_SmsBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmsBalanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveManageRpcServiceServer).SmsBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveManageRpcService_SmsBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveManageRpcServiceServer).SmsBalance(ctx, req.(*SmsBalanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveManageRpcService_SendEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveManageRpcServiceServer).SendEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveManageRpcService_SendEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveManageRpcServiceServer).SendEmail(ctx, req.(*SendEmailReq))
	}
	return interceptor(ctx, in, info, handler)
}

// LiveManageRpcService_ServiceDesc is the grpc.ServiceDesc for LiveManageRpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LiveManageRpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "manage.v1.LiveManageRpcService",
	HandlerType: (*LiveManageRpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLegalTenderInfo",
			Handler:    _LiveManageRpcService_GetLegalTenderInfo_Handler,
		},
		{
			MethodName: "GetCurrInfoById",
			Handler:    _LiveManageRpcService_GetCurrInfoById_Handler,
		},
		{
			MethodName: "GetIPGeolocation",
			Handler:    _LiveManageRpcService_GetIPGeolocation_Handler,
		},
		{
			MethodName: "BatchGetIPGeolocation",
			Handler:    _LiveManageRpcService_BatchGetIPGeolocation_Handler,
		},
		{
			MethodName: "SendSms",
			Handler:    _LiveManageRpcService_SendSms_Handler,
		},
		{
			MethodName: "SmsBalance",
			Handler:    _LiveManageRpcService_SmsBalance_Handler,
		},
		{
			MethodName: "SendEmail",
			Handler:    _LiveManageRpcService_SendEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "manage/v1/manage.proto",
}
