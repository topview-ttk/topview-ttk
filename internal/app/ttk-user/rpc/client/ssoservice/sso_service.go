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
	EmailVerifyCodeLoginRequest       = user.EmailVerifyCodeLoginRequest
	EmailVerifyCodeLoginResponse      = user.EmailVerifyCodeLoginResponse
	PhoneVerifyCodeLoginRequest       = user.PhoneVerifyCodeLoginRequest
	PhoneVerifyCodeLoginResponse      = user.PhoneVerifyCodeLoginResponse
	Request                           = user.Request
	Response                          = user.Response
	SendEmailVerificationCodeRequest  = user.SendEmailVerificationCodeRequest
	SendEmailVerificationCodeResponse = user.SendEmailVerificationCodeResponse
	SendPhoneVerificationCodeRequest  = user.SendPhoneVerificationCodeRequest
	SendPhoneVerificationCodeResponse = user.SendPhoneVerificationCodeResponse
	UserInfo                          = user.UserInfo

	SsoService interface {
		SendPhoneVerificationCode(ctx context.Context, in *SendPhoneVerificationCodeRequest, opts ...grpc.CallOption) (*SendPhoneVerificationCodeResponse, error)
		PhoneVerifyCodeLogin(ctx context.Context, in *PhoneVerifyCodeLoginRequest, opts ...grpc.CallOption) (*PhoneVerifyCodeLoginResponse, error)
		SendEmailVerificationCode(ctx context.Context, in *SendEmailVerificationCodeRequest, opts ...grpc.CallOption) (*SendEmailVerificationCodeResponse, error)
		EmailVerifyCodeLogin(ctx context.Context, in *EmailVerifyCodeLoginRequest, opts ...grpc.CallOption) (*EmailVerifyCodeLoginResponse, error)
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

func (m *defaultSsoService) SendPhoneVerificationCode(ctx context.Context, in *SendPhoneVerificationCodeRequest, opts ...grpc.CallOption) (*SendPhoneVerificationCodeResponse, error) {
	client := user.NewSsoServiceClient(m.cli.Conn())
	return client.SendPhoneVerificationCode(ctx, in, opts...)
}

func (m *defaultSsoService) PhoneVerifyCodeLogin(ctx context.Context, in *PhoneVerifyCodeLoginRequest, opts ...grpc.CallOption) (*PhoneVerifyCodeLoginResponse, error) {
	client := user.NewSsoServiceClient(m.cli.Conn())
	return client.PhoneVerifyCodeLogin(ctx, in, opts...)
}

func (m *defaultSsoService) SendEmailVerificationCode(ctx context.Context, in *SendEmailVerificationCodeRequest, opts ...grpc.CallOption) (*SendEmailVerificationCodeResponse, error) {
	client := user.NewSsoServiceClient(m.cli.Conn())
	return client.SendEmailVerificationCode(ctx, in, opts...)
}

func (m *defaultSsoService) EmailVerifyCodeLogin(ctx context.Context, in *EmailVerifyCodeLoginRequest, opts ...grpc.CallOption) (*EmailVerifyCodeLoginResponse, error) {
	client := user.NewSsoServiceClient(m.cli.Conn())
	return client.EmailVerifyCodeLogin(ctx, in, opts...)
}
