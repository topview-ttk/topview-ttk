package ssoservicelogic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

func (l *EmailVerifyCodeLoginLogic) EmailVerifyCodeLogin(in *user.EmailVerifyCodeLoginRequest) (*user.LoginResponse, error) {
	isValid := util.ValidateEmail(in.GetEmail())
	if !isValid {
		return handleEmailLoginError("请输入正确的邮箱，当前邮箱不合法"), nil
	}
	code, err := l.svcCtx.Rdb.GetDel(l.ctx, "verification:"+in.GetEmail()).Result()
	if err != nil {
		return handleEmailLoginError("系统繁忙或者验证码不存在，请重新发送验证码"), nil
	}

	if code != in.GetVerificationCode() {
		return handleEmailLoginError("验证码错误，请核实短信验证码是否正确"), nil
	}
	userInfo, err := l.svcCtx.TtkUserInfoModel.FindOneByEmail(l.ctx, in.GetEmail())
	fmt.Println(err)
	if err != nil && errors.Is(err, sqlc.ErrNotFound) {
		userInfo = login.CreateDefaultUserInfo()
		userInfo.Email = sql.NullString{String: in.GetEmail(), Valid: true}
		_, err := l.svcCtx.TtkUserInfoModel.Insert(l.ctx, userInfo)
		if err != nil {
			logx.Error(err)
			return handleEmailLoginError("网络繁忙，请重试尝试登录"), err
		}
		// 为了进行缓存用户数据，需要查询数据库，如果err大概率插入失败，让用户重新登录
		userInfo, err = l.svcCtx.TtkUserInfoModel.FindOneByEmail(l.ctx, in.GetEmail())
		if err != nil {
			logx.Error(err)
			return handleEmailLoginError("网络繁忙，请重试尝试登录"), err
		}
	} else if err != nil {
		logx.Error(err)
		return handleEmailLoginError("网络繁忙，请重试尝试登录"), err
	}

	return &user.LoginResponse{
		StatusCode: user.StatusCode_OK,
		Message:    "登录成功，正在加载",
		UserInfo: &user.UserInfo{
			Id:       userInfo.Id,
			UserName: userInfo.TtkId,
			NickName: userInfo.NickName.String,
		},
	}, nil
}

func handleEmailLoginError(message string) *user.LoginResponse {
	return &user.LoginResponse{
		StatusCode: 1,
		Message:    message,
	}
}
