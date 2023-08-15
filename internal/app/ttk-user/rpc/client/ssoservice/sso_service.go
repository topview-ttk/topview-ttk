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
	EmailOrTTKPassLoginRequest       = user.EmailOrTTKPassLoginRequest
	EmailVerifyCodeLoginRequest      = user.EmailVerifyCodeLoginRequest
	LoginResponse                    = user.LoginResponse
	PhonePassLoginRequest            = user.PhonePassLoginRequest
	PhoneVerifyCodeLoginRequest      = user.PhoneVerifyCodeLoginRequest
	Request                          = user.Request
	Response                         = user.Response
	SendEmailVerificationCodeRequest = user.SendEmailVerificationCodeRequest
	SendPhoneVerificationCodeRequest = user.SendPhoneVerificationCodeRequest
	SendVerificationCodeResponse     = user.SendVerificationCodeResponse
	UserInfo                         = user.UserInfo

	SsoService interface {
		// 发送验证码
		SendPhoneVerificationCode(ctx context.Context, in *SendPhoneVerificationCodeRequest, opts ...grpc.CallOption) (*SendVerificationCodeResponse, error)
		SendEmailVerificationCode(ctx context.Context, in *SendEmailVerificationCodeRequest, opts ...grpc.CallOption) (*SendVerificationCodeResponse, error)
		// 登录
		PhoneVerifyCodeLogin(ctx context.Context, in *PhoneVerifyCodeLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		EmailVerifyCodeLogin(ctx context.Context, in *EmailVerifyCodeLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		PhonePassLogin(ctx context.Context, in *PhonePassLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		EmailOrTtkPassLogin(ctx context.Context, in *EmailOrTTKPassLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
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

func (m *defaultSsoService) EmailOrTtkPassLogin(ctx context.Context, in *EmailOrTTKPassLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := user.NewSsoServiceClient(m.cli.Conn())
	return client.EmailOrTtkPassLogin(ctx, in, opts...)
}
