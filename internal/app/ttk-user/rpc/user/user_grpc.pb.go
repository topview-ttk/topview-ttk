// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: user.proto

package user

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
	User_Ping_FullMethodName = "/user.User/Ping"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, User_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _User_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

const (
	SsoService_SendPhoneVerificationCode_FullMethodName = "/user.SsoService/SendPhoneVerificationCode"
	SsoService_PhoneVerifyCodeLogin_FullMethodName      = "/user.SsoService/PhoneVerifyCodeLogin"
)

// SsoServiceClient is the client API for SsoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SsoServiceClient interface {
	SendPhoneVerificationCode(ctx context.Context, in *SendPhoneVerificationCodeRequest, opts ...grpc.CallOption) (*SendPhoneVerificationCodeResponse, error)
	PhoneVerifyCodeLogin(ctx context.Context, in *PhoneVerifyCodeLoginRequest, opts ...grpc.CallOption) (*PhoneVerifyCodeLoginResponse, error)
}

type ssoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSsoServiceClient(cc grpc.ClientConnInterface) SsoServiceClient {
	return &ssoServiceClient{cc}
}

func (c *ssoServiceClient) SendPhoneVerificationCode(ctx context.Context, in *SendPhoneVerificationCodeRequest, opts ...grpc.CallOption) (*SendPhoneVerificationCodeResponse, error) {
	out := new(SendPhoneVerificationCodeResponse)
	err := c.cc.Invoke(ctx, SsoService_SendPhoneVerificationCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoServiceClient) PhoneVerifyCodeLogin(ctx context.Context, in *PhoneVerifyCodeLoginRequest, opts ...grpc.CallOption) (*PhoneVerifyCodeLoginResponse, error) {
	out := new(PhoneVerifyCodeLoginResponse)
	err := c.cc.Invoke(ctx, SsoService_PhoneVerifyCodeLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SsoServiceServer is the server API for SsoService service.
// All implementations must embed UnimplementedSsoServiceServer
// for forward compatibility
type SsoServiceServer interface {
	SendPhoneVerificationCode(context.Context, *SendPhoneVerificationCodeRequest) (*SendPhoneVerificationCodeResponse, error)
	PhoneVerifyCodeLogin(context.Context, *PhoneVerifyCodeLoginRequest) (*PhoneVerifyCodeLoginResponse, error)
	mustEmbedUnimplementedSsoServiceServer()
}

// UnimplementedSsoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSsoServiceServer struct {
}

func (UnimplementedSsoServiceServer) SendPhoneVerificationCode(context.Context, *SendPhoneVerificationCodeRequest) (*SendPhoneVerificationCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPhoneVerificationCode not implemented")
}
func (UnimplementedSsoServiceServer) PhoneVerifyCodeLogin(context.Context, *PhoneVerifyCodeLoginRequest) (*PhoneVerifyCodeLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PhoneVerifyCodeLogin not implemented")
}
func (UnimplementedSsoServiceServer) mustEmbedUnimplementedSsoServiceServer() {}

// UnsafeSsoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SsoServiceServer will
// result in compilation errors.
type UnsafeSsoServiceServer interface {
	mustEmbedUnimplementedSsoServiceServer()
}

func RegisterSsoServiceServer(s grpc.ServiceRegistrar, srv SsoServiceServer) {
	s.RegisterService(&SsoService_ServiceDesc, srv)
}

func _SsoService_SendPhoneVerificationCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendPhoneVerificationCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsoServiceServer).SendPhoneVerificationCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SsoService_SendPhoneVerificationCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsoServiceServer).SendPhoneVerificationCode(ctx, req.(*SendPhoneVerificationCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SsoService_PhoneVerifyCodeLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhoneVerifyCodeLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsoServiceServer).PhoneVerifyCodeLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SsoService_PhoneVerifyCodeLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsoServiceServer).PhoneVerifyCodeLogin(ctx, req.(*PhoneVerifyCodeLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SsoService_ServiceDesc is the grpc.ServiceDesc for SsoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SsoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.SsoService",
	HandlerType: (*SsoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendPhoneVerificationCode",
			Handler:    _SsoService_SendPhoneVerificationCode_Handler,
		},
		{
			MethodName: "PhoneVerifyCodeLogin",
			Handler:    _SsoService_PhoneVerifyCodeLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
