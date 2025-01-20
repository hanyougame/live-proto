// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.8
// source: finance/v1/finance.proto

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
	LivePaymentRpcService_PayIn_FullMethodName        = "/finance.v1.LivePaymentRpcService/PayIn"
	LivePaymentRpcService_PayOut_FullMethodName       = "/finance.v1.LivePaymentRpcService/PayOut"
	LivePaymentRpcService_PayInStatus_FullMethodName  = "/finance.v1.LivePaymentRpcService/PayInStatus"
	LivePaymentRpcService_PayOutStatus_FullMethodName = "/finance.v1.LivePaymentRpcService/PayOutStatus"
	LivePaymentRpcService_Balance_FullMethodName      = "/finance.v1.LivePaymentRpcService/Balance"
	LivePaymentRpcService_Recharge_FullMethodName     = "/finance.v1.LivePaymentRpcService/Recharge"
)

// LivePaymentRpcServiceClient is the client API for LivePaymentRpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LivePaymentRpcServiceClient interface {
	PayIn(ctx context.Context, in *PayInReq, opts ...grpc.CallOption) (*PayInResp, error)
	PayOut(ctx context.Context, in *PayOutReq, opts ...grpc.CallOption) (*PayOutResp, error)
	PayInStatus(ctx context.Context, in *PayInStatusReq, opts ...grpc.CallOption) (*PayInStatusResp, error)
	PayOutStatus(ctx context.Context, in *PayOutStatusReq, opts ...grpc.CallOption) (*PayOutStatusResp, error)
	Balance(ctx context.Context, in *BalanceReq, opts ...grpc.CallOption) (*BalanceResp, error)
	Recharge(ctx context.Context, in *RechargeReq, opts ...grpc.CallOption) (*RechargeResp, error)
}

type livePaymentRpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLivePaymentRpcServiceClient(cc grpc.ClientConnInterface) LivePaymentRpcServiceClient {
	return &livePaymentRpcServiceClient{cc}
}

func (c *livePaymentRpcServiceClient) PayIn(ctx context.Context, in *PayInReq, opts ...grpc.CallOption) (*PayInResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayInResp)
	err := c.cc.Invoke(ctx, LivePaymentRpcService_PayIn_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *livePaymentRpcServiceClient) PayOut(ctx context.Context, in *PayOutReq, opts ...grpc.CallOption) (*PayOutResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayOutResp)
	err := c.cc.Invoke(ctx, LivePaymentRpcService_PayOut_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *livePaymentRpcServiceClient) PayInStatus(ctx context.Context, in *PayInStatusReq, opts ...grpc.CallOption) (*PayInStatusResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayInStatusResp)
	err := c.cc.Invoke(ctx, LivePaymentRpcService_PayInStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *livePaymentRpcServiceClient) PayOutStatus(ctx context.Context, in *PayOutStatusReq, opts ...grpc.CallOption) (*PayOutStatusResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayOutStatusResp)
	err := c.cc.Invoke(ctx, LivePaymentRpcService_PayOutStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *livePaymentRpcServiceClient) Balance(ctx context.Context, in *BalanceReq, opts ...grpc.CallOption) (*BalanceResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BalanceResp)
	err := c.cc.Invoke(ctx, LivePaymentRpcService_Balance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *livePaymentRpcServiceClient) Recharge(ctx context.Context, in *RechargeReq, opts ...grpc.CallOption) (*RechargeResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RechargeResp)
	err := c.cc.Invoke(ctx, LivePaymentRpcService_Recharge_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LivePaymentRpcServiceServer is the server API for LivePaymentRpcService service.
// All implementations must embed UnimplementedLivePaymentRpcServiceServer
// for forward compatibility.
type LivePaymentRpcServiceServer interface {
	PayIn(context.Context, *PayInReq) (*PayInResp, error)
	PayOut(context.Context, *PayOutReq) (*PayOutResp, error)
	PayInStatus(context.Context, *PayInStatusReq) (*PayInStatusResp, error)
	PayOutStatus(context.Context, *PayOutStatusReq) (*PayOutStatusResp, error)
	Balance(context.Context, *BalanceReq) (*BalanceResp, error)
	Recharge(context.Context, *RechargeReq) (*RechargeResp, error)
	mustEmbedUnimplementedLivePaymentRpcServiceServer()
}

// UnimplementedLivePaymentRpcServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLivePaymentRpcServiceServer struct{}

func (UnimplementedLivePaymentRpcServiceServer) PayIn(context.Context, *PayInReq) (*PayInResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayIn not implemented")
}
func (UnimplementedLivePaymentRpcServiceServer) PayOut(context.Context, *PayOutReq) (*PayOutResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayOut not implemented")
}
func (UnimplementedLivePaymentRpcServiceServer) PayInStatus(context.Context, *PayInStatusReq) (*PayInStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayInStatus not implemented")
}
func (UnimplementedLivePaymentRpcServiceServer) PayOutStatus(context.Context, *PayOutStatusReq) (*PayOutStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayOutStatus not implemented")
}
func (UnimplementedLivePaymentRpcServiceServer) Balance(context.Context, *BalanceReq) (*BalanceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balance not implemented")
}
func (UnimplementedLivePaymentRpcServiceServer) Recharge(context.Context, *RechargeReq) (*RechargeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recharge not implemented")
}
func (UnimplementedLivePaymentRpcServiceServer) mustEmbedUnimplementedLivePaymentRpcServiceServer() {}
func (UnimplementedLivePaymentRpcServiceServer) testEmbeddedByValue()                               {}

// UnsafeLivePaymentRpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LivePaymentRpcServiceServer will
// result in compilation errors.
type UnsafeLivePaymentRpcServiceServer interface {
	mustEmbedUnimplementedLivePaymentRpcServiceServer()
}

func RegisterLivePaymentRpcServiceServer(s grpc.ServiceRegistrar, srv LivePaymentRpcServiceServer) {
	// If the following call pancis, it indicates UnimplementedLivePaymentRpcServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LivePaymentRpcService_ServiceDesc, srv)
}

func _LivePaymentRpcService_PayIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayInReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LivePaymentRpcServiceServer).PayIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LivePaymentRpcService_PayIn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LivePaymentRpcServiceServer).PayIn(ctx, req.(*PayInReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LivePaymentRpcService_PayOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayOutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LivePaymentRpcServiceServer).PayOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LivePaymentRpcService_PayOut_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LivePaymentRpcServiceServer).PayOut(ctx, req.(*PayOutReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LivePaymentRpcService_PayInStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayInStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LivePaymentRpcServiceServer).PayInStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LivePaymentRpcService_PayInStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LivePaymentRpcServiceServer).PayInStatus(ctx, req.(*PayInStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LivePaymentRpcService_PayOutStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayOutStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LivePaymentRpcServiceServer).PayOutStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LivePaymentRpcService_PayOutStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LivePaymentRpcServiceServer).PayOutStatus(ctx, req.(*PayOutStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LivePaymentRpcService_Balance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LivePaymentRpcServiceServer).Balance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LivePaymentRpcService_Balance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LivePaymentRpcServiceServer).Balance(ctx, req.(*BalanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LivePaymentRpcService_Recharge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RechargeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LivePaymentRpcServiceServer).Recharge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LivePaymentRpcService_Recharge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LivePaymentRpcServiceServer).Recharge(ctx, req.(*RechargeReq))
	}
	return interceptor(ctx, in, info, handler)
}

// LivePaymentRpcService_ServiceDesc is the grpc.ServiceDesc for LivePaymentRpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LivePaymentRpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "finance.v1.LivePaymentRpcService",
	HandlerType: (*LivePaymentRpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PayIn",
			Handler:    _LivePaymentRpcService_PayIn_Handler,
		},
		{
			MethodName: "PayOut",
			Handler:    _LivePaymentRpcService_PayOut_Handler,
		},
		{
			MethodName: "PayInStatus",
			Handler:    _LivePaymentRpcService_PayInStatus_Handler,
		},
		{
			MethodName: "PayOutStatus",
			Handler:    _LivePaymentRpcService_PayOutStatus_Handler,
		},
		{
			MethodName: "Balance",
			Handler:    _LivePaymentRpcService_Balance_Handler,
		},
		{
			MethodName: "Recharge",
			Handler:    _LivePaymentRpcService_Recharge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "finance/v1/finance.proto",
}

const (
	LiveAuditRpcService_AddAudit_FullMethodName          = "/finance.v1.LiveAuditRpcService/AddAudit"
	LiveAuditRpcService_UpdateAuditAmount_FullMethodName = "/finance.v1.LiveAuditRpcService/UpdateAuditAmount"
)

// LiveAuditRpcServiceClient is the client API for LiveAuditRpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LiveAuditRpcServiceClient interface {
	// 添加稽核 传递用户币种和金额
	AddAudit(ctx context.Context, in *AddAuditReq, opts ...grpc.CallOption) (*AddAuditResp, error)
	// 更新稽核金额（打稽核） 传递用户币种和金额
	UpdateAuditAmount(ctx context.Context, in *UpdateAuditAmountReq, opts ...grpc.CallOption) (*UpdateAuditAmountResp, error)
}

type liveAuditRpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLiveAuditRpcServiceClient(cc grpc.ClientConnInterface) LiveAuditRpcServiceClient {
	return &liveAuditRpcServiceClient{cc}
}

func (c *liveAuditRpcServiceClient) AddAudit(ctx context.Context, in *AddAuditReq, opts ...grpc.CallOption) (*AddAuditResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddAuditResp)
	err := c.cc.Invoke(ctx, LiveAuditRpcService_AddAudit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveAuditRpcServiceClient) UpdateAuditAmount(ctx context.Context, in *UpdateAuditAmountReq, opts ...grpc.CallOption) (*UpdateAuditAmountResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateAuditAmountResp)
	err := c.cc.Invoke(ctx, LiveAuditRpcService_UpdateAuditAmount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LiveAuditRpcServiceServer is the server API for LiveAuditRpcService service.
// All implementations must embed UnimplementedLiveAuditRpcServiceServer
// for forward compatibility.
type LiveAuditRpcServiceServer interface {
	// 添加稽核 传递用户币种和金额
	AddAudit(context.Context, *AddAuditReq) (*AddAuditResp, error)
	// 更新稽核金额（打稽核） 传递用户币种和金额
	UpdateAuditAmount(context.Context, *UpdateAuditAmountReq) (*UpdateAuditAmountResp, error)
	mustEmbedUnimplementedLiveAuditRpcServiceServer()
}

// UnimplementedLiveAuditRpcServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLiveAuditRpcServiceServer struct{}

func (UnimplementedLiveAuditRpcServiceServer) AddAudit(context.Context, *AddAuditReq) (*AddAuditResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAudit not implemented")
}
func (UnimplementedLiveAuditRpcServiceServer) UpdateAuditAmount(context.Context, *UpdateAuditAmountReq) (*UpdateAuditAmountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAuditAmount not implemented")
}
func (UnimplementedLiveAuditRpcServiceServer) mustEmbedUnimplementedLiveAuditRpcServiceServer() {}
func (UnimplementedLiveAuditRpcServiceServer) testEmbeddedByValue()                             {}

// UnsafeLiveAuditRpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LiveAuditRpcServiceServer will
// result in compilation errors.
type UnsafeLiveAuditRpcServiceServer interface {
	mustEmbedUnimplementedLiveAuditRpcServiceServer()
}

func RegisterLiveAuditRpcServiceServer(s grpc.ServiceRegistrar, srv LiveAuditRpcServiceServer) {
	// If the following call pancis, it indicates UnimplementedLiveAuditRpcServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LiveAuditRpcService_ServiceDesc, srv)
}

func _LiveAuditRpcService_AddAudit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAuditReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveAuditRpcServiceServer).AddAudit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveAuditRpcService_AddAudit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveAuditRpcServiceServer).AddAudit(ctx, req.(*AddAuditReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveAuditRpcService_UpdateAuditAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAuditAmountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveAuditRpcServiceServer).UpdateAuditAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveAuditRpcService_UpdateAuditAmount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveAuditRpcServiceServer).UpdateAuditAmount(ctx, req.(*UpdateAuditAmountReq))
	}
	return interceptor(ctx, in, info, handler)
}

// LiveAuditRpcService_ServiceDesc is the grpc.ServiceDesc for LiveAuditRpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LiveAuditRpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "finance.v1.LiveAuditRpcService",
	HandlerType: (*LiveAuditRpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddAudit",
			Handler:    _LiveAuditRpcService_AddAudit_Handler,
		},
		{
			MethodName: "UpdateAuditAmount",
			Handler:    _LiveAuditRpcService_UpdateAuditAmount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "finance/v1/finance.proto",
}

const (
	LiveWithdrawLimitRpcService_AddWithdrawLimit_FullMethodName          = "/finance.v1.LiveWithdrawLimitRpcService/AddWithdrawLimit"
	LiveWithdrawLimitRpcService_UpdateWithdrawLimitAmount_FullMethodName = "/finance.v1.LiveWithdrawLimitRpcService/UpdateWithdrawLimitAmount"
)

// LiveWithdrawLimitRpcServiceClient is the client API for LiveWithdrawLimitRpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LiveWithdrawLimitRpcServiceClient interface {
	// 添加提现限制 传递用户币种和金额
	AddWithdrawLimit(ctx context.Context, in *AddWithdrawLimitReq, opts ...grpc.CallOption) (*AddWithdrawLimitResp, error)
	// 更新提现限制金额 传递用户币种和金额
	UpdateWithdrawLimitAmount(ctx context.Context, in *UpdateWithdrawLimitAmountReq, opts ...grpc.CallOption) (*UpdateWithdrawLimitAmountResp, error)
}

type liveWithdrawLimitRpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLiveWithdrawLimitRpcServiceClient(cc grpc.ClientConnInterface) LiveWithdrawLimitRpcServiceClient {
	return &liveWithdrawLimitRpcServiceClient{cc}
}

func (c *liveWithdrawLimitRpcServiceClient) AddWithdrawLimit(ctx context.Context, in *AddWithdrawLimitReq, opts ...grpc.CallOption) (*AddWithdrawLimitResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddWithdrawLimitResp)
	err := c.cc.Invoke(ctx, LiveWithdrawLimitRpcService_AddWithdrawLimit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveWithdrawLimitRpcServiceClient) UpdateWithdrawLimitAmount(ctx context.Context, in *UpdateWithdrawLimitAmountReq, opts ...grpc.CallOption) (*UpdateWithdrawLimitAmountResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateWithdrawLimitAmountResp)
	err := c.cc.Invoke(ctx, LiveWithdrawLimitRpcService_UpdateWithdrawLimitAmount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LiveWithdrawLimitRpcServiceServer is the server API for LiveWithdrawLimitRpcService service.
// All implementations must embed UnimplementedLiveWithdrawLimitRpcServiceServer
// for forward compatibility.
type LiveWithdrawLimitRpcServiceServer interface {
	// 添加提现限制 传递用户币种和金额
	AddWithdrawLimit(context.Context, *AddWithdrawLimitReq) (*AddWithdrawLimitResp, error)
	// 更新提现限制金额 传递用户币种和金额
	UpdateWithdrawLimitAmount(context.Context, *UpdateWithdrawLimitAmountReq) (*UpdateWithdrawLimitAmountResp, error)
	mustEmbedUnimplementedLiveWithdrawLimitRpcServiceServer()
}

// UnimplementedLiveWithdrawLimitRpcServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLiveWithdrawLimitRpcServiceServer struct{}

func (UnimplementedLiveWithdrawLimitRpcServiceServer) AddWithdrawLimit(context.Context, *AddWithdrawLimitReq) (*AddWithdrawLimitResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddWithdrawLimit not implemented")
}
func (UnimplementedLiveWithdrawLimitRpcServiceServer) UpdateWithdrawLimitAmount(context.Context, *UpdateWithdrawLimitAmountReq) (*UpdateWithdrawLimitAmountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWithdrawLimitAmount not implemented")
}
func (UnimplementedLiveWithdrawLimitRpcServiceServer) mustEmbedUnimplementedLiveWithdrawLimitRpcServiceServer() {
}
func (UnimplementedLiveWithdrawLimitRpcServiceServer) testEmbeddedByValue() {}

// UnsafeLiveWithdrawLimitRpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LiveWithdrawLimitRpcServiceServer will
// result in compilation errors.
type UnsafeLiveWithdrawLimitRpcServiceServer interface {
	mustEmbedUnimplementedLiveWithdrawLimitRpcServiceServer()
}

func RegisterLiveWithdrawLimitRpcServiceServer(s grpc.ServiceRegistrar, srv LiveWithdrawLimitRpcServiceServer) {
	// If the following call pancis, it indicates UnimplementedLiveWithdrawLimitRpcServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LiveWithdrawLimitRpcService_ServiceDesc, srv)
}

func _LiveWithdrawLimitRpcService_AddWithdrawLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddWithdrawLimitReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveWithdrawLimitRpcServiceServer).AddWithdrawLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveWithdrawLimitRpcService_AddWithdrawLimit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveWithdrawLimitRpcServiceServer).AddWithdrawLimit(ctx, req.(*AddWithdrawLimitReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveWithdrawLimitRpcService_UpdateWithdrawLimitAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWithdrawLimitAmountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveWithdrawLimitRpcServiceServer).UpdateWithdrawLimitAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveWithdrawLimitRpcService_UpdateWithdrawLimitAmount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveWithdrawLimitRpcServiceServer).UpdateWithdrawLimitAmount(ctx, req.(*UpdateWithdrawLimitAmountReq))
	}
	return interceptor(ctx, in, info, handler)
}

// LiveWithdrawLimitRpcService_ServiceDesc is the grpc.ServiceDesc for LiveWithdrawLimitRpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LiveWithdrawLimitRpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "finance.v1.LiveWithdrawLimitRpcService",
	HandlerType: (*LiveWithdrawLimitRpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddWithdrawLimit",
			Handler:    _LiveWithdrawLimitRpcService_AddWithdrawLimit_Handler,
		},
		{
			MethodName: "UpdateWithdrawLimitAmount",
			Handler:    _LiveWithdrawLimitRpcService_UpdateWithdrawLimitAmount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "finance/v1/finance.proto",
}
