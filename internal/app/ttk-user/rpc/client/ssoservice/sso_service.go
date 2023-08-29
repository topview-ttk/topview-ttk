// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package ssoservice

import (
	"context"

	"topview-ttk/internal/app/ttk-user/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	EmailPassLoginRequest              = user.EmailPassLoginRequest
	EmailRegisterRequest               = user.EmailRegisterRequest
	EmailVerifyCodeLoginRequest        = user.EmailVerifyCodeLoginRequest
	GetUserInfoByTTKIdRequest          = user.GetUserInfoByTTKIdRequest
	GetUserInfoByUidRequest            = user.GetUserInfoByUidRequest
	GetUserInfoListByRangeNameRequest  = user.GetUserInfoListByRangeNameRequest
	GetUserInfoListByRangeNameResponse = user.GetUserInfoListByRangeNameResponse
	GetUserInfoResponse                = user.GetUserInfoResponse
	LoginCommon                        = user.LoginCommon
	LoginResponse                      = user.LoginResponse
	Page                               = user.Page
	PhonePassLoginRequest              = user.PhonePassLoginRequest
	PhoneRegisterRequest               = user.PhoneRegisterRequest
	PhoneVerifyCodeLoginRequest        = user.PhoneVerifyCodeLoginRequest
	RefreshTokenRequest                = user.RefreshTokenRequest
	RefreshTokenResponse               = user.RefreshTokenResponse
	RegisterResponse                   = user.RegisterResponse
	SendEmailVerificationCodeRequest   = user.SendEmailVerificationCodeRequest
	SendPhoneVerificationCodeRequest   = user.SendPhoneVerificationCodeRequest
	SendVerificationCodeResponse       = user.SendVerificationCodeResponse
	StandbyLoginRequest                = user.StandbyLoginRequest
	TTKIdPassLoginRequest              = user.TTKIdPassLoginRequest
	ThirdPartyLoginRequest             = user.ThirdPartyLoginRequest
	UserInfo                           = user.UserInfo

	SsoService interface {
		// 发送验证码
		SendPhoneVerificationCode(ctx context.Context, in *SendPhoneVerificationCodeRequest, opts ...grpc.CallOption) (*SendVerificationCodeResponse, error)
		SendEmailVerificationCode(ctx context.Context, in *SendEmailVerificationCodeRequest, opts ...grpc.CallOption) (*SendVerificationCodeResponse, error)
		// 登录
		PhoneVerifyCodeLogin(ctx context.Context, in *PhoneVerifyCodeLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		EmailVerifyCodeLogin(ctx context.Context, in *EmailVerifyCodeLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		PhonePassLogin(ctx context.Context, in *PhonePassLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		TTKIdPassLogin(ctx context.Context, in *TTKIdPassLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		EmailPassLogin(ctx context.Context, in *EmailPassLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		GoogleLogin(ctx context.Context, in *ThirdPartyLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		FacebookLogin(ctx context.Context, in *ThirdPartyLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		GithubLogin(ctx context.Context, in *ThirdPartyLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		StandbyLogin(ctx context.Context, in *StandbyLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		// 注册
		EmailRegister(ctx context.Context, in *EmailRegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
		PhoneRegister(ctx context.Context, in *PhoneRegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
		RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error)
	}

	defaultSsoService struct {
		cli zrpc.Client
	}
)

func NewSsoService(cli zrpc.Client) SsoService {
	return &defaultSsoService{
		cli: cli,
	}
}

// 发送验证码
func (m *defaultSsoService) SendPhoneVerificationCode(ctx context.Context, in *SendPhoneVerificationCodeRequest, opts ...grpc.CallOption) (*SendVerificationCodeResponse, error) {
	client := user.NewSsoServiceClient(m.cli.Conn())
	return client.SendPhoneVerificationCode(ctx, in, opts...)
}

func (m *defaultSsoService) SendEmailVerificationCode(ctx context.Context, in *SendEmailVerificationCodeRequest, opts ...grpc.CallOption) (*SendVerificationCodeResponse, error) {
	client := user.NewSsoServiceClient(m.cli.Conn())
	return client.SendEmailVerificationCode(ctx, in, opts...)
}

// 登录
func (m *defaultSsoService) PhoneVerifyCodeLogin(ctx context.Context, in *PhoneVerifyCodeLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := user.NewSsoServiceClient(m.cli.Conn())
	return client.PhoneVerifyCodeLogin(ctx, in, opts...)
}

func (m *defaultSsoService) EmailVerifyCodeLogin(ctx context.Context, in *EmailVerifyCodeLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := user.NewSsoServiceClient(m.cli.Conn())
	return client.EmailVerifyCodeLogin(ctx, in, opts...)
}

func (m *defaultSsoService) PhonePassLogin(ctx context.Context, in *PhonePassLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := user.NewSsoServiceClient(m.cli.Conn())
	return client.PhonePassLogin(ctx, in, opts...)
}

func (m *defaultSsoService) TTKIdPassLogin(ctx context.Context, in *TTKIdPassLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := user.NewSsoServiceClient(m.cli.Conn())
	return client.TTKIdPassLogin(ctx, in, opts...)
}

func (m *defaultSsoService) EmailPassLogin(ctx context.Context, in *EmailPassLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := user.NewSsoServiceClient(m.cli.Conn())
	return client.EmailPassLogin(ctx, in, opts...)
}

func (m *defaultSsoService) GoogleLogin(ctx context.Context, in *ThirdPartyLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := user.NewSsoServiceClient(m.cli.Conn())
	return client.GoogleLogin(ctx, in, opts...)
}

func (m *defaultSsoService) FacebookLogin(ctx context.Context, in *ThirdPartyLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := user.NewSsoServiceClient(m.cli.Conn())
	return client.FacebookLogin(ctx, in, opts...)
}

func (m *defaultSsoService) GithubLogin(ctx context.Context, in *ThirdPartyLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := user.NewSsoServiceClient(m.cli.Conn())
	return client.GithubLogin(ctx, in, opts...)
}

func (m *defaultSsoService) StandbyLogin(ctx context.Context, in *StandbyLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := user.NewSsoServiceClient(m.cli.Conn())
	return client.StandbyLogin(ctx, in, opts...)
}

// 注册
func (m *defaultSsoService) EmailRegister(ctx context.Context, in *EmailRegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	client := user.NewSsoServiceClient(m.cli.Conn())
	return client.EmailRegister(ctx, in, opts...)
}

func (m *defaultSsoService) PhoneRegister(ctx context.Context, in *PhoneRegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	client := user.NewSsoServiceClient(m.cli.Conn())
	return client.PhoneRegister(ctx, in, opts...)
}

func (m *defaultSsoService) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error) {
	client := user.NewSsoServiceClient(m.cli.Conn())
	return client.RefreshToken(ctx, in, opts...)
}
