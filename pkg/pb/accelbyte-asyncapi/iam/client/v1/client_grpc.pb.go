// Copyright (c) 2024 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// source: IAM/AsyncAPI/client.yaml (0.1.0)

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: accelbyte-asyncapi/iam/client/v1/client.proto

package client

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ClientClientCreatedServiceClient is the client API for ClientClientCreatedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientClientCreatedServiceClient interface {
	OnMessage(ctx context.Context, in *ClientCreated, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type clientClientCreatedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientClientCreatedServiceClient(cc grpc.ClientConnInterface) ClientClientCreatedServiceClient {
	return &clientClientCreatedServiceClient{cc}
}

func (c *clientClientCreatedServiceClient) OnMessage(ctx context.Context, in *ClientCreated, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/accelbyte.iam.client.v1.ClientClientCreatedService/OnMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientClientCreatedServiceServer is the server API for ClientClientCreatedService service.
// All implementations must embed UnimplementedClientClientCreatedServiceServer
// for forward compatibility
type ClientClientCreatedServiceServer interface {
	OnMessage(context.Context, *ClientCreated) (*emptypb.Empty, error)
	mustEmbedUnimplementedClientClientCreatedServiceServer()
}

// UnimplementedClientClientCreatedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientClientCreatedServiceServer struct {
}

func (UnimplementedClientClientCreatedServiceServer) OnMessage(context.Context, *ClientCreated) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnMessage not implemented")
}
func (UnimplementedClientClientCreatedServiceServer) mustEmbedUnimplementedClientClientCreatedServiceServer() {
}

// UnsafeClientClientCreatedServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientClientCreatedServiceServer will
// result in compilation errors.
type UnsafeClientClientCreatedServiceServer interface {
	mustEmbedUnimplementedClientClientCreatedServiceServer()
}

func RegisterClientClientCreatedServiceServer(s grpc.ServiceRegistrar, srv ClientClientCreatedServiceServer) {
	s.RegisterService(&ClientClientCreatedService_ServiceDesc, srv)
}

func _ClientClientCreatedService_OnMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientCreated)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientClientCreatedServiceServer).OnMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accelbyte.iam.client.v1.ClientClientCreatedService/OnMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientClientCreatedServiceServer).OnMessage(ctx, req.(*ClientCreated))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientClientCreatedService_ServiceDesc is the grpc.ServiceDesc for ClientClientCreatedService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientClientCreatedService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accelbyte.iam.client.v1.ClientClientCreatedService",
	HandlerType: (*ClientClientCreatedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnMessage",
			Handler:    _ClientClientCreatedService_OnMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accelbyte-asyncapi/iam/client/v1/client.proto",
}

// ClientClientDeletedServiceClient is the client API for ClientClientDeletedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientClientDeletedServiceClient interface {
	OnMessage(ctx context.Context, in *ClientDeleted, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type clientClientDeletedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientClientDeletedServiceClient(cc grpc.ClientConnInterface) ClientClientDeletedServiceClient {
	return &clientClientDeletedServiceClient{cc}
}

func (c *clientClientDeletedServiceClient) OnMessage(ctx context.Context, in *ClientDeleted, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/accelbyte.iam.client.v1.ClientClientDeletedService/OnMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientClientDeletedServiceServer is the server API for ClientClientDeletedService service.
// All implementations must embed UnimplementedClientClientDeletedServiceServer
// for forward compatibility
type ClientClientDeletedServiceServer interface {
	OnMessage(context.Context, *ClientDeleted) (*emptypb.Empty, error)
	mustEmbedUnimplementedClientClientDeletedServiceServer()
}

// UnimplementedClientClientDeletedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientClientDeletedServiceServer struct {
}

func (UnimplementedClientClientDeletedServiceServer) OnMessage(context.Context, *ClientDeleted) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnMessage not implemented")
}
func (UnimplementedClientClientDeletedServiceServer) mustEmbedUnimplementedClientClientDeletedServiceServer() {
}

// UnsafeClientClientDeletedServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientClientDeletedServiceServer will
// result in compilation errors.
type UnsafeClientClientDeletedServiceServer interface {
	mustEmbedUnimplementedClientClientDeletedServiceServer()
}

func RegisterClientClientDeletedServiceServer(s grpc.ServiceRegistrar, srv ClientClientDeletedServiceServer) {
	s.RegisterService(&ClientClientDeletedService_ServiceDesc, srv)
}

func _ClientClientDeletedService_OnMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientDeleted)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientClientDeletedServiceServer).OnMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accelbyte.iam.client.v1.ClientClientDeletedService/OnMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientClientDeletedServiceServer).OnMessage(ctx, req.(*ClientDeleted))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientClientDeletedService_ServiceDesc is the grpc.ServiceDesc for ClientClientDeletedService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientClientDeletedService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accelbyte.iam.client.v1.ClientClientDeletedService",
	HandlerType: (*ClientClientDeletedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnMessage",
			Handler:    _ClientClientDeletedService_OnMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accelbyte-asyncapi/iam/client/v1/client.proto",
}

// ClientClientUpdatedServiceClient is the client API for ClientClientUpdatedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientClientUpdatedServiceClient interface {
	OnMessage(ctx context.Context, in *ClientUpdated, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type clientClientUpdatedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientClientUpdatedServiceClient(cc grpc.ClientConnInterface) ClientClientUpdatedServiceClient {
	return &clientClientUpdatedServiceClient{cc}
}

func (c *clientClientUpdatedServiceClient) OnMessage(ctx context.Context, in *ClientUpdated, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/accelbyte.iam.client.v1.ClientClientUpdatedService/OnMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientClientUpdatedServiceServer is the server API for ClientClientUpdatedService service.
// All implementations must embed UnimplementedClientClientUpdatedServiceServer
// for forward compatibility
type ClientClientUpdatedServiceServer interface {
	OnMessage(context.Context, *ClientUpdated) (*emptypb.Empty, error)
	mustEmbedUnimplementedClientClientUpdatedServiceServer()
}

// UnimplementedClientClientUpdatedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientClientUpdatedServiceServer struct {
}

func (UnimplementedClientClientUpdatedServiceServer) OnMessage(context.Context, *ClientUpdated) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnMessage not implemented")
}
func (UnimplementedClientClientUpdatedServiceServer) mustEmbedUnimplementedClientClientUpdatedServiceServer() {
}

// UnsafeClientClientUpdatedServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientClientUpdatedServiceServer will
// result in compilation errors.
type UnsafeClientClientUpdatedServiceServer interface {
	mustEmbedUnimplementedClientClientUpdatedServiceServer()
}

func RegisterClientClientUpdatedServiceServer(s grpc.ServiceRegistrar, srv ClientClientUpdatedServiceServer) {
	s.RegisterService(&ClientClientUpdatedService_ServiceDesc, srv)
}

func _ClientClientUpdatedService_OnMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientUpdated)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientClientUpdatedServiceServer).OnMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accelbyte.iam.client.v1.ClientClientUpdatedService/OnMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientClientUpdatedServiceServer).OnMessage(ctx, req.(*ClientUpdated))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientClientUpdatedService_ServiceDesc is the grpc.ServiceDesc for ClientClientUpdatedService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientClientUpdatedService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accelbyte.iam.client.v1.ClientClientUpdatedService",
	HandlerType: (*ClientClientUpdatedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnMessage",
			Handler:    _ClientClientUpdatedService_OnMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accelbyte-asyncapi/iam/client/v1/client.proto",
}

// ClientThirdPartyClientThirdPartyCreatedServiceClient is the client API for ClientThirdPartyClientThirdPartyCreatedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientThirdPartyClientThirdPartyCreatedServiceClient interface {
	OnMessage(ctx context.Context, in *ClientThirdPartyCreated, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type clientThirdPartyClientThirdPartyCreatedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientThirdPartyClientThirdPartyCreatedServiceClient(cc grpc.ClientConnInterface) ClientThirdPartyClientThirdPartyCreatedServiceClient {
	return &clientThirdPartyClientThirdPartyCreatedServiceClient{cc}
}

func (c *clientThirdPartyClientThirdPartyCreatedServiceClient) OnMessage(ctx context.Context, in *ClientThirdPartyCreated, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/accelbyte.iam.client.v1.ClientThirdPartyClientThirdPartyCreatedService/OnMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientThirdPartyClientThirdPartyCreatedServiceServer is the server API for ClientThirdPartyClientThirdPartyCreatedService service.
// All implementations must embed UnimplementedClientThirdPartyClientThirdPartyCreatedServiceServer
// for forward compatibility
type ClientThirdPartyClientThirdPartyCreatedServiceServer interface {
	OnMessage(context.Context, *ClientThirdPartyCreated) (*emptypb.Empty, error)
	mustEmbedUnimplementedClientThirdPartyClientThirdPartyCreatedServiceServer()
}

// UnimplementedClientThirdPartyClientThirdPartyCreatedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientThirdPartyClientThirdPartyCreatedServiceServer struct {
}

func (UnimplementedClientThirdPartyClientThirdPartyCreatedServiceServer) OnMessage(context.Context, *ClientThirdPartyCreated) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnMessage not implemented")
}
func (UnimplementedClientThirdPartyClientThirdPartyCreatedServiceServer) mustEmbedUnimplementedClientThirdPartyClientThirdPartyCreatedServiceServer() {
}

// UnsafeClientThirdPartyClientThirdPartyCreatedServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientThirdPartyClientThirdPartyCreatedServiceServer will
// result in compilation errors.
type UnsafeClientThirdPartyClientThirdPartyCreatedServiceServer interface {
	mustEmbedUnimplementedClientThirdPartyClientThirdPartyCreatedServiceServer()
}

func RegisterClientThirdPartyClientThirdPartyCreatedServiceServer(s grpc.ServiceRegistrar, srv ClientThirdPartyClientThirdPartyCreatedServiceServer) {
	s.RegisterService(&ClientThirdPartyClientThirdPartyCreatedService_ServiceDesc, srv)
}

func _ClientThirdPartyClientThirdPartyCreatedService_OnMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientThirdPartyCreated)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientThirdPartyClientThirdPartyCreatedServiceServer).OnMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accelbyte.iam.client.v1.ClientThirdPartyClientThirdPartyCreatedService/OnMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientThirdPartyClientThirdPartyCreatedServiceServer).OnMessage(ctx, req.(*ClientThirdPartyCreated))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientThirdPartyClientThirdPartyCreatedService_ServiceDesc is the grpc.ServiceDesc for ClientThirdPartyClientThirdPartyCreatedService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientThirdPartyClientThirdPartyCreatedService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accelbyte.iam.client.v1.ClientThirdPartyClientThirdPartyCreatedService",
	HandlerType: (*ClientThirdPartyClientThirdPartyCreatedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnMessage",
			Handler:    _ClientThirdPartyClientThirdPartyCreatedService_OnMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accelbyte-asyncapi/iam/client/v1/client.proto",
}

// ClientThirdPartyClientThirdPartyDeletedServiceClient is the client API for ClientThirdPartyClientThirdPartyDeletedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientThirdPartyClientThirdPartyDeletedServiceClient interface {
	OnMessage(ctx context.Context, in *ClientThirdPartyDeleted, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type clientThirdPartyClientThirdPartyDeletedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientThirdPartyClientThirdPartyDeletedServiceClient(cc grpc.ClientConnInterface) ClientThirdPartyClientThirdPartyDeletedServiceClient {
	return &clientThirdPartyClientThirdPartyDeletedServiceClient{cc}
}

func (c *clientThirdPartyClientThirdPartyDeletedServiceClient) OnMessage(ctx context.Context, in *ClientThirdPartyDeleted, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/accelbyte.iam.client.v1.ClientThirdPartyClientThirdPartyDeletedService/OnMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientThirdPartyClientThirdPartyDeletedServiceServer is the server API for ClientThirdPartyClientThirdPartyDeletedService service.
// All implementations must embed UnimplementedClientThirdPartyClientThirdPartyDeletedServiceServer
// for forward compatibility
type ClientThirdPartyClientThirdPartyDeletedServiceServer interface {
	OnMessage(context.Context, *ClientThirdPartyDeleted) (*emptypb.Empty, error)
	mustEmbedUnimplementedClientThirdPartyClientThirdPartyDeletedServiceServer()
}

// UnimplementedClientThirdPartyClientThirdPartyDeletedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientThirdPartyClientThirdPartyDeletedServiceServer struct {
}

func (UnimplementedClientThirdPartyClientThirdPartyDeletedServiceServer) OnMessage(context.Context, *ClientThirdPartyDeleted) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnMessage not implemented")
}
func (UnimplementedClientThirdPartyClientThirdPartyDeletedServiceServer) mustEmbedUnimplementedClientThirdPartyClientThirdPartyDeletedServiceServer() {
}

// UnsafeClientThirdPartyClientThirdPartyDeletedServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientThirdPartyClientThirdPartyDeletedServiceServer will
// result in compilation errors.
type UnsafeClientThirdPartyClientThirdPartyDeletedServiceServer interface {
	mustEmbedUnimplementedClientThirdPartyClientThirdPartyDeletedServiceServer()
}

func RegisterClientThirdPartyClientThirdPartyDeletedServiceServer(s grpc.ServiceRegistrar, srv ClientThirdPartyClientThirdPartyDeletedServiceServer) {
	s.RegisterService(&ClientThirdPartyClientThirdPartyDeletedService_ServiceDesc, srv)
}

func _ClientThirdPartyClientThirdPartyDeletedService_OnMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientThirdPartyDeleted)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientThirdPartyClientThirdPartyDeletedServiceServer).OnMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accelbyte.iam.client.v1.ClientThirdPartyClientThirdPartyDeletedService/OnMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientThirdPartyClientThirdPartyDeletedServiceServer).OnMessage(ctx, req.(*ClientThirdPartyDeleted))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientThirdPartyClientThirdPartyDeletedService_ServiceDesc is the grpc.ServiceDesc for ClientThirdPartyClientThirdPartyDeletedService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientThirdPartyClientThirdPartyDeletedService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accelbyte.iam.client.v1.ClientThirdPartyClientThirdPartyDeletedService",
	HandlerType: (*ClientThirdPartyClientThirdPartyDeletedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnMessage",
			Handler:    _ClientThirdPartyClientThirdPartyDeletedService_OnMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accelbyte-asyncapi/iam/client/v1/client.proto",
}

// ClientThirdPartyClientThirdPartyUpdatedServiceClient is the client API for ClientThirdPartyClientThirdPartyUpdatedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientThirdPartyClientThirdPartyUpdatedServiceClient interface {
	OnMessage(ctx context.Context, in *ClientThirdPartyUpdated, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type clientThirdPartyClientThirdPartyUpdatedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientThirdPartyClientThirdPartyUpdatedServiceClient(cc grpc.ClientConnInterface) ClientThirdPartyClientThirdPartyUpdatedServiceClient {
	return &clientThirdPartyClientThirdPartyUpdatedServiceClient{cc}
}

func (c *clientThirdPartyClientThirdPartyUpdatedServiceClient) OnMessage(ctx context.Context, in *ClientThirdPartyUpdated, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/accelbyte.iam.client.v1.ClientThirdPartyClientThirdPartyUpdatedService/OnMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientThirdPartyClientThirdPartyUpdatedServiceServer is the server API for ClientThirdPartyClientThirdPartyUpdatedService service.
// All implementations must embed UnimplementedClientThirdPartyClientThirdPartyUpdatedServiceServer
// for forward compatibility
type ClientThirdPartyClientThirdPartyUpdatedServiceServer interface {
	OnMessage(context.Context, *ClientThirdPartyUpdated) (*emptypb.Empty, error)
	mustEmbedUnimplementedClientThirdPartyClientThirdPartyUpdatedServiceServer()
}

// UnimplementedClientThirdPartyClientThirdPartyUpdatedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientThirdPartyClientThirdPartyUpdatedServiceServer struct {
}

func (UnimplementedClientThirdPartyClientThirdPartyUpdatedServiceServer) OnMessage(context.Context, *ClientThirdPartyUpdated) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnMessage not implemented")
}
func (UnimplementedClientThirdPartyClientThirdPartyUpdatedServiceServer) mustEmbedUnimplementedClientThirdPartyClientThirdPartyUpdatedServiceServer() {
}

// UnsafeClientThirdPartyClientThirdPartyUpdatedServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientThirdPartyClientThirdPartyUpdatedServiceServer will
// result in compilation errors.
type UnsafeClientThirdPartyClientThirdPartyUpdatedServiceServer interface {
	mustEmbedUnimplementedClientThirdPartyClientThirdPartyUpdatedServiceServer()
}

func RegisterClientThirdPartyClientThirdPartyUpdatedServiceServer(s grpc.ServiceRegistrar, srv ClientThirdPartyClientThirdPartyUpdatedServiceServer) {
	s.RegisterService(&ClientThirdPartyClientThirdPartyUpdatedService_ServiceDesc, srv)
}

func _ClientThirdPartyClientThirdPartyUpdatedService_OnMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientThirdPartyUpdated)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientThirdPartyClientThirdPartyUpdatedServiceServer).OnMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accelbyte.iam.client.v1.ClientThirdPartyClientThirdPartyUpdatedService/OnMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientThirdPartyClientThirdPartyUpdatedServiceServer).OnMessage(ctx, req.(*ClientThirdPartyUpdated))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientThirdPartyClientThirdPartyUpdatedService_ServiceDesc is the grpc.ServiceDesc for ClientThirdPartyClientThirdPartyUpdatedService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientThirdPartyClientThirdPartyUpdatedService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accelbyte.iam.client.v1.ClientThirdPartyClientThirdPartyUpdatedService",
	HandlerType: (*ClientThirdPartyClientThirdPartyUpdatedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnMessage",
			Handler:    _ClientThirdPartyClientThirdPartyUpdatedService_OnMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accelbyte-asyncapi/iam/client/v1/client.proto",
}

// ClientPermissionsClientPermissionsCreatedServiceClient is the client API for ClientPermissionsClientPermissionsCreatedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientPermissionsClientPermissionsCreatedServiceClient interface {
	OnMessage(ctx context.Context, in *ClientPermissionCreated, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type clientPermissionsClientPermissionsCreatedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientPermissionsClientPermissionsCreatedServiceClient(cc grpc.ClientConnInterface) ClientPermissionsClientPermissionsCreatedServiceClient {
	return &clientPermissionsClientPermissionsCreatedServiceClient{cc}
}

func (c *clientPermissionsClientPermissionsCreatedServiceClient) OnMessage(ctx context.Context, in *ClientPermissionCreated, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/accelbyte.iam.client.v1.ClientPermissionsClientPermissionsCreatedService/OnMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientPermissionsClientPermissionsCreatedServiceServer is the server API for ClientPermissionsClientPermissionsCreatedService service.
// All implementations must embed UnimplementedClientPermissionsClientPermissionsCreatedServiceServer
// for forward compatibility
type ClientPermissionsClientPermissionsCreatedServiceServer interface {
	OnMessage(context.Context, *ClientPermissionCreated) (*emptypb.Empty, error)
	mustEmbedUnimplementedClientPermissionsClientPermissionsCreatedServiceServer()
}

// UnimplementedClientPermissionsClientPermissionsCreatedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientPermissionsClientPermissionsCreatedServiceServer struct {
}

func (UnimplementedClientPermissionsClientPermissionsCreatedServiceServer) OnMessage(context.Context, *ClientPermissionCreated) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnMessage not implemented")
}
func (UnimplementedClientPermissionsClientPermissionsCreatedServiceServer) mustEmbedUnimplementedClientPermissionsClientPermissionsCreatedServiceServer() {
}

// UnsafeClientPermissionsClientPermissionsCreatedServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientPermissionsClientPermissionsCreatedServiceServer will
// result in compilation errors.
type UnsafeClientPermissionsClientPermissionsCreatedServiceServer interface {
	mustEmbedUnimplementedClientPermissionsClientPermissionsCreatedServiceServer()
}

func RegisterClientPermissionsClientPermissionsCreatedServiceServer(s grpc.ServiceRegistrar, srv ClientPermissionsClientPermissionsCreatedServiceServer) {
	s.RegisterService(&ClientPermissionsClientPermissionsCreatedService_ServiceDesc, srv)
}

func _ClientPermissionsClientPermissionsCreatedService_OnMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientPermissionCreated)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientPermissionsClientPermissionsCreatedServiceServer).OnMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accelbyte.iam.client.v1.ClientPermissionsClientPermissionsCreatedService/OnMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientPermissionsClientPermissionsCreatedServiceServer).OnMessage(ctx, req.(*ClientPermissionCreated))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientPermissionsClientPermissionsCreatedService_ServiceDesc is the grpc.ServiceDesc for ClientPermissionsClientPermissionsCreatedService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientPermissionsClientPermissionsCreatedService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accelbyte.iam.client.v1.ClientPermissionsClientPermissionsCreatedService",
	HandlerType: (*ClientPermissionsClientPermissionsCreatedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnMessage",
			Handler:    _ClientPermissionsClientPermissionsCreatedService_OnMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accelbyte-asyncapi/iam/client/v1/client.proto",
}

// ClientPermissionsClientPermissionsDeletedServiceClient is the client API for ClientPermissionsClientPermissionsDeletedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientPermissionsClientPermissionsDeletedServiceClient interface {
	OnMessage(ctx context.Context, in *ClientPermissionDeleted, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type clientPermissionsClientPermissionsDeletedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientPermissionsClientPermissionsDeletedServiceClient(cc grpc.ClientConnInterface) ClientPermissionsClientPermissionsDeletedServiceClient {
	return &clientPermissionsClientPermissionsDeletedServiceClient{cc}
}

func (c *clientPermissionsClientPermissionsDeletedServiceClient) OnMessage(ctx context.Context, in *ClientPermissionDeleted, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/accelbyte.iam.client.v1.ClientPermissionsClientPermissionsDeletedService/OnMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientPermissionsClientPermissionsDeletedServiceServer is the server API for ClientPermissionsClientPermissionsDeletedService service.
// All implementations must embed UnimplementedClientPermissionsClientPermissionsDeletedServiceServer
// for forward compatibility
type ClientPermissionsClientPermissionsDeletedServiceServer interface {
	OnMessage(context.Context, *ClientPermissionDeleted) (*emptypb.Empty, error)
	mustEmbedUnimplementedClientPermissionsClientPermissionsDeletedServiceServer()
}

// UnimplementedClientPermissionsClientPermissionsDeletedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientPermissionsClientPermissionsDeletedServiceServer struct {
}

func (UnimplementedClientPermissionsClientPermissionsDeletedServiceServer) OnMessage(context.Context, *ClientPermissionDeleted) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnMessage not implemented")
}
func (UnimplementedClientPermissionsClientPermissionsDeletedServiceServer) mustEmbedUnimplementedClientPermissionsClientPermissionsDeletedServiceServer() {
}

// UnsafeClientPermissionsClientPermissionsDeletedServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientPermissionsClientPermissionsDeletedServiceServer will
// result in compilation errors.
type UnsafeClientPermissionsClientPermissionsDeletedServiceServer interface {
	mustEmbedUnimplementedClientPermissionsClientPermissionsDeletedServiceServer()
}

func RegisterClientPermissionsClientPermissionsDeletedServiceServer(s grpc.ServiceRegistrar, srv ClientPermissionsClientPermissionsDeletedServiceServer) {
	s.RegisterService(&ClientPermissionsClientPermissionsDeletedService_ServiceDesc, srv)
}

func _ClientPermissionsClientPermissionsDeletedService_OnMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientPermissionDeleted)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientPermissionsClientPermissionsDeletedServiceServer).OnMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accelbyte.iam.client.v1.ClientPermissionsClientPermissionsDeletedService/OnMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientPermissionsClientPermissionsDeletedServiceServer).OnMessage(ctx, req.(*ClientPermissionDeleted))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientPermissionsClientPermissionsDeletedService_ServiceDesc is the grpc.ServiceDesc for ClientPermissionsClientPermissionsDeletedService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientPermissionsClientPermissionsDeletedService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accelbyte.iam.client.v1.ClientPermissionsClientPermissionsDeletedService",
	HandlerType: (*ClientPermissionsClientPermissionsDeletedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnMessage",
			Handler:    _ClientPermissionsClientPermissionsDeletedService_OnMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accelbyte-asyncapi/iam/client/v1/client.proto",
}

// ClientPermissionsClientPermissionsUpdatedServiceClient is the client API for ClientPermissionsClientPermissionsUpdatedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientPermissionsClientPermissionsUpdatedServiceClient interface {
	OnMessage(ctx context.Context, in *ClientPermissionUpdated, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type clientPermissionsClientPermissionsUpdatedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientPermissionsClientPermissionsUpdatedServiceClient(cc grpc.ClientConnInterface) ClientPermissionsClientPermissionsUpdatedServiceClient {
	return &clientPermissionsClientPermissionsUpdatedServiceClient{cc}
}

func (c *clientPermissionsClientPermissionsUpdatedServiceClient) OnMessage(ctx context.Context, in *ClientPermissionUpdated, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/accelbyte.iam.client.v1.ClientPermissionsClientPermissionsUpdatedService/OnMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientPermissionsClientPermissionsUpdatedServiceServer is the server API for ClientPermissionsClientPermissionsUpdatedService service.
// All implementations must embed UnimplementedClientPermissionsClientPermissionsUpdatedServiceServer
// for forward compatibility
type ClientPermissionsClientPermissionsUpdatedServiceServer interface {
	OnMessage(context.Context, *ClientPermissionUpdated) (*emptypb.Empty, error)
	mustEmbedUnimplementedClientPermissionsClientPermissionsUpdatedServiceServer()
}

// UnimplementedClientPermissionsClientPermissionsUpdatedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientPermissionsClientPermissionsUpdatedServiceServer struct {
}

func (UnimplementedClientPermissionsClientPermissionsUpdatedServiceServer) OnMessage(context.Context, *ClientPermissionUpdated) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnMessage not implemented")
}
func (UnimplementedClientPermissionsClientPermissionsUpdatedServiceServer) mustEmbedUnimplementedClientPermissionsClientPermissionsUpdatedServiceServer() {
}

// UnsafeClientPermissionsClientPermissionsUpdatedServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientPermissionsClientPermissionsUpdatedServiceServer will
// result in compilation errors.
type UnsafeClientPermissionsClientPermissionsUpdatedServiceServer interface {
	mustEmbedUnimplementedClientPermissionsClientPermissionsUpdatedServiceServer()
}

func RegisterClientPermissionsClientPermissionsUpdatedServiceServer(s grpc.ServiceRegistrar, srv ClientPermissionsClientPermissionsUpdatedServiceServer) {
	s.RegisterService(&ClientPermissionsClientPermissionsUpdatedService_ServiceDesc, srv)
}

func _ClientPermissionsClientPermissionsUpdatedService_OnMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientPermissionUpdated)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientPermissionsClientPermissionsUpdatedServiceServer).OnMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accelbyte.iam.client.v1.ClientPermissionsClientPermissionsUpdatedService/OnMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientPermissionsClientPermissionsUpdatedServiceServer).OnMessage(ctx, req.(*ClientPermissionUpdated))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientPermissionsClientPermissionsUpdatedService_ServiceDesc is the grpc.ServiceDesc for ClientPermissionsClientPermissionsUpdatedService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientPermissionsClientPermissionsUpdatedService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accelbyte.iam.client.v1.ClientPermissionsClientPermissionsUpdatedService",
	HandlerType: (*ClientPermissionsClientPermissionsUpdatedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnMessage",
			Handler:    _ClientPermissionsClientPermissionsUpdatedService_OnMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accelbyte-asyncapi/iam/client/v1/client.proto",
}
