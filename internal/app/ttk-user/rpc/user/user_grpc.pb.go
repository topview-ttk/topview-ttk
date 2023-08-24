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
	UserService_GetUserInfoByUid_FullMethodName   = "/user.UserService/GetUserInfoByUid"
	UserService_GetUserInfoByTTKId_FullMethodName = "/user.UserService/GetUserInfoByTTKId"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUserInfoByUid(ctx context.Context, in *GetUserInfoByUidRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error)
	GetUserInfoByTTKId(ctx context.Context, in *GetUserInfoByTTKIdRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUserInfoByUid(ctx context.Context, in *GetUserInfoByUidRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error) {
	out := new(GetUserInfoResponse)
	err := c.cc.Invoke(ctx, UserService_GetUserInfoByUid_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserInfoByTTKId(ctx context.Context, in *GetUserInfoByTTKIdRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error) {
	out := new(GetUserInfoResponse)
	err := c.cc.Invoke(ctx, UserService_GetUserInfoByTTKId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	GetUserInfoByUid(context.Context, *GetUserInfoByUidRequest) (*GetUserInfoResponse, error)
	GetUserInfoByTTKId(context.Context, *GetUserInfoByTTKIdRequest) (*GetUserInfoResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetUserInfoByUid(context.Context, *GetUserInfoByUidRequest) (*GetUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfoByUid not implemented")
}
func (UnimplementedUserServiceServer) GetUserInfoByTTKId(context.Context, *GetUserInfoByTTKIdRequest) (*GetUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfoByTTKId not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetUserInfoByUid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoByUidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserInfoByUid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserInfoByUid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserInfoByUid(ctx, req.(*GetUserInfoByUidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserInfoByTTKId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoByTTKIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserInfoByTTKId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserInfoByTTKId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserInfoByTTKId(ctx, req.(*GetUserInfoByTTKIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInfoByUid",
			Handler:    _UserService_GetUserInfoByUid_Handler,
		},
		{
			MethodName: "GetUserInfoByTTKId",
			Handler:    _UserService_GetUserInfoByTTKId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

const (
	SsoService_SendPhoneVerificationCode_FullMethodName = "/user.SsoService/SendPhoneVerificationCode"
	SsoService_SendEmailVerificationCode_FullMethodName = "/user.SsoService/SendEmailVerificationCode"
	SsoService_PhoneVerifyCodeLogin_FullMethodName      = "/user.SsoService/PhoneVerifyCodeLogin"
	SsoService_EmailVerifyCodeLogin_FullMethodName      = "/user.SsoService/EmailVerifyCodeLogin"
	SsoService_PhonePassLogin_FullMethodName            = "/user.SsoService/PhonePassLogin"
	SsoService_TtkidPassLogin_FullMethodName            = "/user.SsoService/TtkidPassLogin"
	SsoService_EmailPassLogin_FullMethodName            = "/user.SsoService/EmailPassLogin"
	SsoService_GithubLogin_FullMethodName               = "/user.SsoService/GithubLogin"
	SsoService_EmailRegister_FullMethodName             = "/user.SsoService/EmailRegister"
	SsoService_PhoneRegister_FullMethodName             = "/user.SsoService/PhoneRegister"
	SsoService_RefreshToken_FullMethodName              = "/user.SsoService/RefreshToken"
)

// SsoServiceClient is the client API for SsoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SsoServiceClient interface {
	// 发送验证码
	SendPhoneVerificationCode(ctx context.Context, in *SendPhoneVerificationCodeRequest, opts ...grpc.CallOption) (*SendVerificationCodeResponse, error)
	SendEmailVerificationCode(ctx context.Context, in *SendEmailVerificationCodeRequest, opts ...grpc.CallOption) (*SendVerificationCodeResponse, error)
	// 登录
	PhoneVerifyCodeLogin(ctx context.Context, in *PhoneVerifyCodeLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	EmailVerifyCodeLogin(ctx context.Context, in *EmailVerifyCodeLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	PhonePassLogin(ctx context.Context, in *PhonePassLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	TtkidPassLogin(ctx context.Context, in *TTKPassLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	EmailPassLogin(ctx context.Context, in *EmailPassLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	GithubLogin(ctx context.Context, in *GitHubLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// 注册
	EmailRegister(ctx context.Context, in *EmailRegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	PhoneRegister(ctx context.Context, in *PhoneRegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error)
}

type ssoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSsoServiceClient(cc grpc.ClientConnInterface) SsoServiceClient {
	return &ssoServiceClient{cc}
}

func (c *ssoServiceClient) SendPhoneVerificationCode(ctx context.Context, in *SendPhoneVerificationCodeRequest, opts ...grpc.CallOption) (*SendVerificationCodeResponse, error) {
	out := new(SendVerificationCodeResponse)
	err := c.cc.Invoke(ctx, SsoService_SendPhoneVerificationCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoServiceClient) SendEmailVerificationCode(ctx context.Context, in *SendEmailVerificationCodeRequest, opts ...grpc.CallOption) (*SendVerificationCodeResponse, error) {
	out := new(SendVerificationCodeResponse)
	err := c.cc.Invoke(ctx, SsoService_SendEmailVerificationCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoServiceClient) PhoneVerifyCodeLogin(ctx context.Context, in *PhoneVerifyCodeLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, SsoService_PhoneVerifyCodeLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoServiceClient) EmailVerifyCodeLogin(ctx context.Context, in *EmailVerifyCodeLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, SsoService_EmailVerifyCodeLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoServiceClient) PhonePassLogin(ctx context.Context, in *PhonePassLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, SsoService_PhonePassLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoServiceClient) TtkidPassLogin(ctx context.Context, in *TTKPassLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, SsoService_TtkidPassLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoServiceClient) EmailPassLogin(ctx context.Context, in *EmailPassLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, SsoService_EmailPassLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoServiceClient) GithubLogin(ctx context.Context, in *GitHubLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, SsoService_GithubLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoServiceClient) EmailRegister(ctx context.Context, in *EmailRegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, SsoService_EmailRegister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoServiceClient) PhoneRegister(ctx context.Context, in *PhoneRegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, SsoService_PhoneRegister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoServiceClient) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error) {
	out := new(RefreshTokenResponse)
	err := c.cc.Invoke(ctx, SsoService_RefreshToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SsoServiceServer is the server API for SsoService service.
// All implementations must embed UnimplementedSsoServiceServer
// for forward compatibility
type SsoServiceServer interface {
	// 发送验证码
	SendPhoneVerificationCode(context.Context, *SendPhoneVerificationCodeRequest) (*SendVerificationCodeResponse, error)
	SendEmailVerificationCode(context.Context, *SendEmailVerificationCodeRequest) (*SendVerificationCodeResponse, error)
	// 登录
	PhoneVerifyCodeLogin(context.Context, *PhoneVerifyCodeLoginRequest) (*LoginResponse, error)
	EmailVerifyCodeLogin(context.Context, *EmailVerifyCodeLoginRequest) (*LoginResponse, error)
	PhonePassLogin(context.Context, *PhonePassLoginRequest) (*LoginResponse, error)
	TtkidPassLogin(context.Context, *TTKPassLoginRequest) (*LoginResponse, error)
	EmailPassLogin(context.Context, *EmailPassLoginRequest) (*LoginResponse, error)
	GithubLogin(context.Context, *GitHubLoginRequest) (*LoginResponse, error)
	// 注册
	EmailRegister(context.Context, *EmailRegisterRequest) (*RegisterResponse, error)
	PhoneRegister(context.Context, *PhoneRegisterRequest) (*RegisterResponse, error)
	RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error)
	mustEmbedUnimplementedSsoServiceServer()
}

// UnimplementedSsoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSsoServiceServer struct {
}

func (UnimplementedSsoServiceServer) SendPhoneVerificationCode(context.Context, *SendPhoneVerificationCodeRequest) (*SendVerificationCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPhoneVerificationCode not implemented")
}
func (UnimplementedSsoServiceServer) SendEmailVerificationCode(context.Context, *SendEmailVerificationCodeRequest) (*SendVerificationCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmailVerificationCode not implemented")
}
func (UnimplementedSsoServiceServer) PhoneVerifyCodeLogin(context.Context, *PhoneVerifyCodeLoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PhoneVerifyCodeLogin not implemented")
}
func (UnimplementedSsoServiceServer) EmailVerifyCodeLogin(context.Context, *EmailVerifyCodeLoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmailVerifyCodeLogin not implemented")
}
func (UnimplementedSsoServiceServer) PhonePassLogin(context.Context, *PhonePassLoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PhonePassLogin not implemented")
}
func (UnimplementedSsoServiceServer) TtkidPassLogin(context.Context, *TTKPassLoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TtkidPassLogin not implemented")
}
func (UnimplementedSsoServiceServer) EmailPassLogin(context.Context, *EmailPassLoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmailPassLogin not implemented")
}
func (UnimplementedSsoServiceServer) GithubLogin(context.Context, *GitHubLoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GithubLogin not implemented")
}
func (UnimplementedSsoServiceServer) EmailRegister(context.Context, *EmailRegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmailRegister not implemented")
}
func (UnimplementedSsoServiceServer) PhoneRegister(context.Context, *PhoneRegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PhoneRegister not implemented")
}
func (UnimplementedSsoServiceServer) RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
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

func _SsoService_SendEmailVerificationCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailVerificationCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsoServiceServer).SendEmailVerificationCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SsoService_SendEmailVerificationCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsoServiceServer).SendEmailVerificationCode(ctx, req.(*SendEmailVerificationCodeRequest))
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

func _SsoService_EmailVerifyCodeLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailVerifyCodeLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsoServiceServer).EmailVerifyCodeLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SsoService_EmailVerifyCodeLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsoServiceServer).EmailVerifyCodeLogin(ctx, req.(*EmailVerifyCodeLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SsoService_PhonePassLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhonePassLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsoServiceServer).PhonePassLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SsoService_PhonePassLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsoServiceServer).PhonePassLogin(ctx, req.(*PhonePassLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SsoService_TtkidPassLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TTKPassLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsoServiceServer).TtkidPassLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SsoService_TtkidPassLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsoServiceServer).TtkidPassLogin(ctx, req.(*TTKPassLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SsoService_EmailPassLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailPassLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsoServiceServer).EmailPassLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SsoService_EmailPassLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsoServiceServer).EmailPassLogin(ctx, req.(*EmailPassLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SsoService_GithubLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GitHubLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsoServiceServer).GithubLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SsoService_GithubLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsoServiceServer).GithubLogin(ctx, req.(*GitHubLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SsoService_EmailRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsoServiceServer).EmailRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SsoService_EmailRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsoServiceServer).EmailRegister(ctx, req.(*EmailRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SsoService_PhoneRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhoneRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsoServiceServer).PhoneRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SsoService_PhoneRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsoServiceServer).PhoneRegister(ctx, req.(*PhoneRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SsoService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsoServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SsoService_RefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsoServiceServer).RefreshToken(ctx, req.(*RefreshTokenRequest))
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
			MethodName: "SendEmailVerificationCode",
			Handler:    _SsoService_SendEmailVerificationCode_Handler,
		},
		{
			MethodName: "PhoneVerifyCodeLogin",
			Handler:    _SsoService_PhoneVerifyCodeLogin_Handler,
		},
		{
			MethodName: "EmailVerifyCodeLogin",
			Handler:    _SsoService_EmailVerifyCodeLogin_Handler,
		},
		{
			MethodName: "PhonePassLogin",
			Handler:    _SsoService_PhonePassLogin_Handler,
		},
		{
			MethodName: "TtkidPassLogin",
			Handler:    _SsoService_TtkidPassLogin_Handler,
		},
		{
			MethodName: "EmailPassLogin",
			Handler:    _SsoService_EmailPassLogin_Handler,
		},
		{
			MethodName: "GithubLogin",
			Handler:    _SsoService_GithubLogin_Handler,
		},
		{
			MethodName: "EmailRegister",
			Handler:    _SsoService_EmailRegister_Handler,
		},
		{
			MethodName: "PhoneRegister",
			Handler:    _SsoService_PhoneRegister_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _SsoService_RefreshToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
