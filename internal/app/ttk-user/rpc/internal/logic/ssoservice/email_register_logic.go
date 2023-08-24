package ssoservicelogic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"topview-ttk/internal/app/ttk-user/rpc/internal/svc"
	"topview-ttk/internal/app/ttk-user/rpc/internal/util"
	"topview-ttk/internal/app/ttk-user/rpc/model"
	"topview-ttk/internal/app/ttk-user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEmailRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailRegisterLogic {
	return &EmailRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 注册
func (l *EmailRegisterLogic) EmailRegister(in *user.EmailRegisterRequest) (*user.RegisterResponse, error) {
	email := in.GetEmail()
	var userCredentials *model.TtkUserCredentials
	var err error

	if util.ValidateEmail(email) {
		userCredentials, err = l.svcCtx.TtkUserInfoModel.FindUserCredentialsByEmail(l.ctx, email)
	} else {
		return &user.RegisterResponse{
			StatusCode: user.StatusCode_INVALID_ARGUMENT,
			Message:    "该邮箱不合法",
		}, nil
	}
	if err != nil && !errors.Is(err, sqlc.ErrNotFound) {
		return &user.RegisterResponse{
			StatusCode: user.StatusCode_INVALID_ARGUMENT,
			Message:    "该邮箱已被注册",
		}, nil
	}

	if err != nil {
		return &user.RegisterResponse{
			StatusCode: user.StatusCode_INVALID_ARGUMENT,
			Message:    "网络繁忙，请重新登录",
		}, nil
	}

	//inputPass := login.EncryptPasswordWithSalt(in.GetPassword(), userCredentials.Salt.String)
	//inputNickname := in.GetNickname()

	_, err = l.svcCtx.TtkUserInfoModel.FindOne(l.ctx, userCredentials.Id)
	return &user.RegisterResponse{
		StatusCode: user.StatusCode_OK,
		Message:    "注册成功，正在加载",
	}, nil

}
