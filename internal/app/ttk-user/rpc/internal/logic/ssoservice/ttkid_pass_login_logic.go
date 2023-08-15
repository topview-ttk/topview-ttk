package ssoservicelogic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"topview-ttk/internal/app/ttk-user/rpc/internal/logic/ssoservice/login"
	"topview-ttk/internal/app/ttk-user/rpc/internal/util"
	"topview-ttk/internal/app/ttk-user/rpc/model"

	"topview-ttk/internal/app/ttk-user/rpc/internal/svc"
	"topview-ttk/internal/app/ttk-user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type TtkidPassLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTtkidPassLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TtkidPassLoginLogic {
	return &TtkidPassLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TtkidPassLoginLogic) TtkidPassLogin(in *user.TTKPassLoginRequest) (*user.LoginResponse, error) {
	Ttkid := in.GetTtkId()
	var userCredentials *model.TtkUserCredentials
	var err error

	if util.ValidateTtkId(Ttkid) {
		userCredentials, err = l.svcCtx.TtkUserInfoModel.FindUserCredentialsByTtkId(l.ctx, Ttkid)
	} else {
		return &user.LoginResponse{
			StatusCode: user.StatusCode_INVALID_ARGUMENT,
			Message:    "ttk_id或者密码错误",
		}, nil
	}
	if err != nil && errors.Is(err, sqlc.ErrNotFound) {
		return &user.LoginResponse{
			StatusCode: user.StatusCode_INVALID_ARGUMENT,
			Message:    "ttk_id或者密码错误",
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
			Message:    "ttk_id或者密码错误",
		}, nil
	}

	inputPass := login.EncryptPasswordWithSalt(in.GetPass(), userCredentials.Salt.String)
	// test
	dbPass := login.EncryptPasswordWithSalt(userCredentials.Password, userCredentials.Salt.String)
	if dbPass != inputPass {
		return &user.LoginResponse{
			StatusCode: user.StatusCode_INVALID_ARGUMENT,
			Message:    "ttk_id或者密码错误",
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
