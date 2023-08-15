package ssoservicelogic

import (
	"context"
	"topview-ttk/internal/app/ttk-user/rpc/internal/logic/ssoservice/send"
	"topview-ttk/internal/app/ttk-user/rpc/internal/util"

	"topview-ttk/internal/app/ttk-user/rpc/internal/svc"
	"topview-ttk/internal/app/ttk-user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendPhoneVerificationCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendPhoneVerificationCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendPhoneVerificationCodeLogic {
	return &SendPhoneVerificationCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendPhoneVerificationCodeLogic) SendPhoneVerificationCode(in *user.SendPhoneVerificationCodeRequest) (*user.SendVerificationCodeResponse, error) {
	isValid := util.ValidatePhoneNumber(in.GetPhone())

	if !isValid {
		return &user.SendVerificationCodeResponse{
			StatusCode: 1,
			Message:    "请输入正确的手机号码，当前手机号码不合法",
		}, nil
	}

	if !send.CanSendVerificationCode(l.ctx, l.svcCtx.Rdb, in.GetPhone()) {
		return &user.SendVerificationCodeResponse{
			StatusCode: 1,
			Message:    "请求过多或网络繁忙，请重试尝试",
		}, nil
	}

	code := send.GenerateVerificationCode()
	send.StoreVerificationCode(l.ctx, l.svcCtx.Rdb, in.GetPhone(), code)
	err := send.AliyunSMS(in.GetPhone(), code)
	if err != nil {
		logx.Error(err)
	}
	return &user.SendVerificationCodeResponse{
		StatusCode: 0,
		Message:    "验证码已发送，请注意查收",
	}, nil
}
