package ssoservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"topview-ttk/internal/app/ttk-user/rpc/internal/logic/ssoservice/send"
	"topview-ttk/internal/app/ttk-user/rpc/internal/svc"
	"topview-ttk/internal/app/ttk-user/rpc/internal/util"
	"topview-ttk/internal/app/ttk-user/rpc/user"
	"topview-ttk/internal/pkg/ttkerr"

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
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.EmailValidError), "邮箱格式错误")
	}

	if !send.CanSendVerificationCode(l.ctx, l.svcCtx.Rdb, in.GetEmail()) {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.SendVerifyCodeFrequentError), "发送验证码频繁，参数：%+v", in)
	}

	code := send.GenerateVerificationCode()
	send.StoreVerificationCode(l.ctx, l.svcCtx.Rdb, in.GetEmail(), code)
	err := send.SMTPEmail(in.GetEmail(), code)
	if err != nil {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.ServerCommonError), "发送验证码失败，参数：%+v", in)
	}
	return &user.SendVerificationCodeResponse{}, nil
}
