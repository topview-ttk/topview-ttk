// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"topview-ttk/internal/app/ttk-user/rpc/internal/logic/ssoservice"
	"topview-ttk/internal/app/ttk-user/rpc/internal/svc"
	"topview-ttk/internal/app/ttk-user/rpc/user"
)

type SsoServiceServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedSsoServiceServer
}

func NewSsoServiceServer(svcCtx *svc.ServiceContext) *SsoServiceServer {
	return &SsoServiceServer{
		svcCtx: svcCtx,
	}
}

// 发送验证码
func (s *SsoServiceServer) SendPhoneVerificationCode(ctx context.Context, in *user.SendPhoneVerificationCodeRequest) (*user.SendVerificationCodeResponse, error) {
	l := ssoservicelogic.NewSendPhoneVerificationCodeLogic(ctx, s.svcCtx)
	return l.SendPhoneVerificationCode(in)
}

func (s *SsoServiceServer) SendEmailVerificationCode(ctx context.Context, in *user.SendEmailVerificationCodeRequest) (*user.SendVerificationCodeResponse, error) {
	l := ssoservicelogic.NewSendEmailVerificationCodeLogic(ctx, s.svcCtx)
	return l.SendEmailVerificationCode(in)
}

// 登录
func (s *SsoServiceServer) PhoneVerifyCodeLogin(ctx context.Context, in *user.PhoneVerifyCodeLoginRequest) (*user.LoginResponse, error) {
	l := ssoservicelogic.NewPhoneVerifyCodeLoginLogic(ctx, s.svcCtx)
	return l.PhoneVerifyCodeLogin(in)
}

func (s *SsoServiceServer) EmailVerifyCodeLogin(ctx context.Context, in *user.EmailVerifyCodeLoginRequest) (*user.LoginResponse, error) {
	l := ssoservicelogic.NewEmailVerifyCodeLoginLogic(ctx, s.svcCtx)
	return l.EmailVerifyCodeLogin(in)
}

func (s *SsoServiceServer) PhonePassLogin(ctx context.Context, in *user.PhonePassLoginRequest) (*user.LoginResponse, error) {
	l := ssoservicelogic.NewPhonePassLoginLogic(ctx, s.svcCtx)
	return l.PhonePassLogin(in)
}

func (s *SsoServiceServer) TTKIdPassLogin(ctx context.Context, in *user.TTKIdPassLoginRequest) (*user.LoginResponse, error) {
	l := ssoservicelogic.NewTTKIdPassLoginLogic(ctx, s.svcCtx)
	return l.TTKIdPassLogin(in)
}

func (s *SsoServiceServer) EmailPassLogin(ctx context.Context, in *user.EmailPassLoginRequest) (*user.LoginResponse, error) {
	l := ssoservicelogic.NewEmailPassLoginLogic(ctx, s.svcCtx)
	return l.EmailPassLogin(in)
}

func (s *SsoServiceServer) GoogleLogin(ctx context.Context, in *user.ThirdPartyLoginRequest) (*user.LoginResponse, error) {
	l := ssoservicelogic.NewGoogleLoginLogic(ctx, s.svcCtx)
	return l.GoogleLogin(in)
}

func (s *SsoServiceServer) FacebookLogin(ctx context.Context, in *user.ThirdPartyLoginRequest) (*user.LoginResponse, error) {
	l := ssoservicelogic.NewFacebookLoginLogic(ctx, s.svcCtx)
	return l.FacebookLogin(in)
}

func (s *SsoServiceServer) GithubLogin(ctx context.Context, in *user.ThirdPartyLoginRequest) (*user.LoginResponse, error) {
	l := ssoservicelogic.NewGithubLoginLogic(ctx, s.svcCtx)
	return l.GithubLogin(in)
}

func (s *SsoServiceServer) StandbyGoogleLogin(ctx context.Context, in *user.StandbyLoginRequest) (*user.LoginResponse, error) {
	l := ssoservicelogic.NewStandbyGoogleLoginLogic(ctx, s.svcCtx)
	return l.StandbyGoogleLogin(in)
}

func (s *SsoServiceServer) StandbyFacebookLogin(ctx context.Context, in *user.StandbyLoginRequest) (*user.LoginResponse, error) {
	l := ssoservicelogic.NewStandbyFacebookLoginLogic(ctx, s.svcCtx)
	return l.StandbyFacebookLogin(in)
}

// 注册
func (s *SsoServiceServer) EmailRegister(ctx context.Context, in *user.EmailRegisterRequest) (*user.RegisterResponse, error) {
	l := ssoservicelogic.NewEmailRegisterLogic(ctx, s.svcCtx)
	return l.EmailRegister(in)
}

func (s *SsoServiceServer) PhoneRegister(ctx context.Context, in *user.PhoneRegisterRequest) (*user.RegisterResponse, error) {
	l := ssoservicelogic.NewPhoneRegisterLogic(ctx, s.svcCtx)
	return l.PhoneRegister(in)
}

func (s *SsoServiceServer) RefreshToken(ctx context.Context, in *user.RefreshTokenRequest) (*user.RefreshTokenResponse, error) {
	l := ssoservicelogic.NewRefreshTokenLogic(ctx, s.svcCtx)
	return l.RefreshToken(in)
}
