package ssoservicelogic

import (
	"context"
	"fmt"
	"topview-ttk/internal/app/ttk-user/rpc/internal/logic/ssoservice/send"
	"topview-ttk/internal/app/ttk-user/rpc/internal/util"

	"topview-ttk/internal/app/ttk-user/rpc/internal/svc"
	"topview-ttk/internal/app/ttk-user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailVerificationCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendEmailVerificationCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailVerificationCodeLogic {
	return &SendEmailVerificationCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendEmailVerificationCodeLogic) SendEmailVerificationCode(in *user.SendEmailVerificationCodeRequest) (*user.SendVerificationCodeResponse, error) {
	isValid := util.ValidateEmail(in.GetEmail())

	if !isValid {
		return &user.SendVerificationCodeResponse{
			StatusCode: 1,
			Message:    "请输入正确的邮箱，当前邮箱不合法",
		}, nil
	}

	if !send.CanSendVerificationCode(l.ctx, l.svcCtx.Rdb, in.GetEmail()) {
		return &user.SendVerificationCodeResponse{
			StatusCode: 1,
			Message:    "请求过多或网络繁忙，请重试尝试",
		}, nil
	}

	code := send.GenerateVerificationCode()
	send.StoreVerificationCode(l.ctx, l.svcCtx.Rdb, in.GetEmail(), code)
	err := send.SMTPEmail(in.GetEmail(), code)
	if err != nil {
		fmt.Println(err)
		logx.Error(err)
	}
	return &user.SendVerificationCodeResponse{
		StatusCode: 0,
		Message:    "验证码已发送，请注意查收",
	}, nil
}
