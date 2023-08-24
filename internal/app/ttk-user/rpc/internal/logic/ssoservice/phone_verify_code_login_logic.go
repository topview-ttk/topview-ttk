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

type PhoneVerifyCodeLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPhoneVerifyCodeLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhoneVerifyCodeLoginLogic {
	return &PhoneVerifyCodeLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PhoneVerifyCodeLoginLogic) PhoneVerifyCodeLogin(in *user.PhoneVerifyCodeLoginRequest) (*user.LoginResponse, error) {
	isValid := util.ValidatePhoneNumber(in.GetPhone())

	if !isValid {
		return handlePhoneLoginError("请输入正确的手机号码，当前手机号码不合法"), nil
	}

	code, err := l.svcCtx.Rdb.GetDel(l.ctx, "verification:"+in.GetPhone()).Result()
	if err != nil {
		return handlePhoneLoginError("系统繁忙或者验证码不存在，请重新发送验证码"), err
	}

	if code != in.GetVerificationCode() {
		return handlePhoneLoginError("验证码错误，请核实短信验证码是否正确"), nil
	}
	userInfo, err := l.svcCtx.TtkUserInfoModel.FindOneByPhone(l.ctx, in.GetPhone())
	if err != nil && errors.Is(err, sqlc.ErrNotFound) {
		userInfo = login.CreateDefaultUserInfo()
		userInfo.Phone = sql.NullString{String: in.GetPhone(), Valid: true}
		_, err := l.svcCtx.TtkUserInfoModel.Insert(l.ctx, userInfo)
		if err != nil {
			logx.Error(err)
			return handleEmailLoginError("网络繁忙，请重试尝试登录"), err
		}
		// 为了进行缓存用户数据，需要查询数据库，如果err大概率插入失败，让用户重新登录
		userInfo, err = l.svcCtx.TtkUserInfoModel.FindOneByPhone(l.ctx, in.GetPhone())
		if err != nil {
			logx.Error(err)
			return handleEmailLoginError("网络繁忙，请重试尝试登录"), err
		}
	} else if err != nil {
		logx.Error(err)
		return handlePhoneLoginError("网络繁忙，请重试尝试登录"), err
	}

	token, err := login.GenerateVfToken(in.DeviceInfo, in.ClientInfo, userInfo.Id)
	if err != nil {
		logx.Error(err)
		return &user.LoginResponse{
			StatusCode: user.StatusCode_INVALID_ARGUMENT,
			Message:    "系统繁忙，请重试！",
		}, nil
	}
	return &user.LoginResponse{
		StatusCode: user.StatusCode_OK,
		Message:    "登录成功，正在加载",
		Token:      token,
		UserInfo: &user.UserInfo{
			Id:       userInfo.Id,
			UserName: userInfo.TtkId,
			NickName: userInfo.NickName.String,
		},
	}, nil
}

func handlePhoneLoginError(message string) *user.LoginResponse {
	return &user.LoginResponse{
		StatusCode: user.StatusCode_INVALID_ARGUMENT,
		Message:    message,
	}
}
