package sso

import (
	"context"
	"topview-ttk/internal/app/ttk-user/rpc/user"

	"topview-ttk/internal/app/ttk-api/internal/svc"
	"topview-ttk/internal/app/ttk-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendPhoneVerificationCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendPhoneVerificationCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendPhoneVerificationCodeLogic {
	return &SendPhoneVerificationCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendPhoneVerificationCodeLogic) SendPhoneVerificationCode(req *types.SendPhoneVerificationCodeRequest) (resp *types.SendVerificationCodeResponse, err error) {
	rpcResp, err := l.svcCtx.SsoClient.SendPhoneVerificationCode(l.ctx, &user.SendPhoneVerificationCodeRequest{
		Phone:      req.Phone,
		DeviceInfo: req.DeviceInfo,
		ClientInfo: req.ClientInfo,
	})

	if err != nil {
		logx.Error(err)
		return &types.SendVerificationCodeResponse{}, err
	}

	return &types.SendVerificationCodeResponse{
		StatusCode: int32(rpcResp.GetStatusCode().Number()),
		Message:    rpcResp.Message,
	}, err
}
