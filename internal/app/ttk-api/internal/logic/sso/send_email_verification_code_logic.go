package sso

import (
	"context"
	"topview-ttk/internal/app/ttk-user/rpc/user"

	"topview-ttk/internal/app/ttk-api/internal/svc"
	"topview-ttk/internal/app/ttk-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailVerificationCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendEmailVerificationCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailVerificationCodeLogic {
	return &SendEmailVerificationCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendEmailVerificationCodeLogic) SendEmailVerificationCode(req *types.SendEmailVerificationCodeRequest) (resp *types.SendVerificationCodeResponse, err error) {
	_, err = l.svcCtx.SsoClient.SendEmailVerificationCode(l.ctx, &user.SendEmailVerificationCodeRequest{
		Email:      req.Email,
		DeviceInfo: req.DeviceInfo,
		ClientInfo: req.ClientInfo,
	})

	if err != nil {
		logx.Error(err)
		return &types.SendVerificationCodeResponse{}, err
	}

	return &types.SendVerificationCodeResponse{}, err
}
