package ssoservicelogic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"topview-ttk/internal/app/ttk-user/rpc/internal/svc"
)

type EmailOrTtkPassLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEmailOrTtkPassLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailOrTtkPassLoginLogic {
	return &EmailOrTtkPassLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//func (l *EmailOrTtkPassLoginLogic) EmailOrTtkPassLogin(in *user.EmailOrTTKPassLoginRequest) (*user.LoginResponse, error) {
//	email := in.GetEmail()
//	ttkId := in.GetTtkId()
//	var userCredentials *model.TtkUserCredentials
//	var err error
//
//	if util.ValidateEmail(email) {
//		userCredentials, err = l.svcCtx.TtkUserInfoModel.FindUserCredentialsByEmail(l.ctx, email)
//	} else if util.ValidateTtkId(ttkId) {
//		userCredentials, err = l.svcCtx.TtkUserInfoModel.FindUserCredentialsByTtkId(l.ctx, ttkId)
//	} else {
//		return &user.LoginResponse{
//			StatusCode: user.StatusCode_INVALID_ARGUMENT,
//			Message:    "TTK ID/邮箱或者密码错误",
//		}, nil
//	}
//	if err != nil && errors.Is(err, sqlc.ErrNotFound) {
//		return &user.LoginResponse{
//			StatusCode: user.StatusCode_INVALID_ARGUMENT,
//			Message:    "TTK ID/邮箱或者密码错误",
//		}, nil
//	}
//
//	if err != nil {
//		return &user.LoginResponse{
//			StatusCode: user.StatusCode_INVALID_ARGUMENT,
//			Message:    "网络繁忙，请重新登录",
//		}, nil
//	}
//	if in.GetPass() == "" {
//		return &user.LoginResponse{
//			StatusCode: user.StatusCode_INVALID_ARGUMENT,
//			Message:    "TTK ID/邮箱或者密码错误",
//		}, nil
//	}
//	inputPass := login.EncryptPasswordWithSalt(in.GetPass(), userCredentials.Salt.String)
//	// test
//	dbPass := login.EncryptPasswordWithSalt(userCredentials.Password, userCredentials.Salt.String)
//	if dbPass != inputPass {
//		return &user.LoginResponse{
//			StatusCode: user.StatusCode_INVALID_ARGUMENT,
//			Message:    "TTK ID/邮箱或者密码错误",
//		}, nil
//	}
//
//	userInfo, err := l.svcCtx.TtkUserInfoModel.FindOne(l.ctx, userCredentials.Id)
//	return &user.LoginResponse{
//		StatusCode: user.StatusCode_OK,
//		Message:    "登录成功，正在加载",
//		UserInfo: &user.UserInfo{
//			Id:       userInfo.Id,
//			UserName: userInfo.TtkId,
//			NickName: userInfo.NickName.String,
//		},
//	}, nil
//}
