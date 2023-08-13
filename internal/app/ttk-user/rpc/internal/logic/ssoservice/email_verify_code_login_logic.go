package ssoservicelogic

import (
	"context"
	"database/sql"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"topview-ttk/internal/app/ttk-user/rpc/internal/logic/ssoservice/login"
	"topview-ttk/internal/app/ttk-user/rpc/internal/svc"
	"topview-ttk/internal/app/ttk-user/rpc/internal/util"
	"topview-ttk/internal/app/ttk-user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailVerifyCodeLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEmailVerifyCodeLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailVerifyCodeLoginLogic {
	return &EmailVerifyCodeLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EmailVerifyCodeLoginLogic) EmailVerifyCodeLogin(in *user.EmailVerifyCodeLoginRequest) (*user.EmailVerifyCodeLoginResponse, error) {
	isValid := util.ValidateEmail(in.GetEmail())
	if !isValid {
		return handleEmailLoginError("请输入正确的邮箱，当前邮箱不合法"), nil
	}
	code, err := l.svcCtx.Rdb.Get(l.ctx, "verification:"+in.GetEmail()).Result()
	if err != nil {
		return handleEmailLoginError("系统繁忙或者验证码不存在，请重新发送验证码"), err
	}

	if code != in.GetVerificationCode() {
		return handleEmailLoginError("验证码错误，请核实短信验证码是否正确"), nil
	}
	userInfo, err := l.svcCtx.TtkUserInfoModel.FindOneByEmail(l.ctx, in.GetEmail())
	if err != nil && errors.Is(err, sqlc.ErrNotFound) {
		userInfo = login.CreateDefaultUserInfo()
		userInfo.Email = sql.NullString{String: in.GetEmail(), Valid: true}
		r, err := l.svcCtx.TtkUserInfoModel.Insert(l.ctx, userInfo)
		if err != nil {
			return handleEmailLoginError("网络繁忙，请重试尝试登录"), err
		} else if a, _ := r.RowsAffected(); a != 1 {
			return handleEmailLoginError("网络繁忙，请重试尝试登录"), err
		}
	} else if err != nil {
		logx.Error(err)
		return handleEmailLoginError("网络繁忙，请重试尝试登录"), err
	}

	return &user.EmailVerifyCodeLoginResponse{}, nil
}

func handleEmailLoginError(message string) *user.EmailVerifyCodeLoginResponse {
	return &user.EmailVerifyCodeLoginResponse{
		StatusCode: 1,
		Message:    message,
	}
}
