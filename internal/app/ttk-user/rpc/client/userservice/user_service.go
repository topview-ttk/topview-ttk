// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userservice

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

	UserService interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultUserService struct {
		cli zrpc.Client
	}
)

func NewUserService(cli zrpc.Client) UserService {
	return &defaultUserService{
		cli: cli,
	}
}

func (m *defaultUserService) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
