package ssoservicelogic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"topview-ttk/internal/app/ttk-user/rpc/internal/logic/ssoservice/login"
	"topview-ttk/internal/app/ttk-user/rpc/internal/util"

	"topview-ttk/internal/app/ttk-user/rpc/internal/svc"
	"topview-ttk/internal/app/ttk-user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhonePassLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPhonePassLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhonePassLoginLogic {
	return &PhonePassLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PhonePassLoginLogic) PhonePassLogin(in *user.PhonePassLoginRequest) (*user.LoginResponse, error) {
	phone := in.GetPhone()
	if !util.ValidatePhoneNumber(phone) {
		return &user.LoginResponse{
			StatusCode: user.StatusCode_INVALID_ARGUMENT,
			Message:    "手机号码或者密码错误",
		}, nil
	}
	userCredentials, err := l.svcCtx.TtkUserInfoModel.FindUserCredentialsByPhone(l.ctx, phone)
	if err != nil && errors.Is(err, sqlc.ErrNotFound) {
		return &user.LoginResponse{
			StatusCode: user.StatusCode_INVALID_ARGUMENT,
			Message:    "手机号码或者密码错误",
		}, nil
	}

	if err != nil {
		return &user.LoginResponse{
			StatusCode: user.StatusCode_INVALID_ARGUMENT,
			Message:    "网络繁忙，请重新登录",
		}, nil
	}

	if in.GetPass() == "" {
		return &user.LoginResponse{
			StatusCode: user.StatusCode_INVALID_ARGUMENT,
			Message:    "手机号码或者密码错误",
		}, nil
	}

	inputPass := login.EncryptPasswordWithSalt(in.GetPass(), userCredentials.Salt.String)
	// test
	dbPass := login.EncryptPasswordWithSalt(userCredentials.Password, userCredentials.Salt.String)

	if dbPass != inputPass {
		return &user.LoginResponse{
			StatusCode: user.StatusCode_INVALID_ARGUMENT,
			Message:    "手机号码或者密码错误",
		}, nil
	}
	userInfo, err := l.svcCtx.TtkUserInfoModel.FindOne(l.ctx, userCredentials.Id)
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
