package ssoservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"topview-ttk/internal/pkg/common/ttkerr"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"topview-ttk/internal/app/ttk-user/rpc/internal/logic/ssoservice/login"
	"topview-ttk/internal/app/ttk-user/rpc/internal/util"
	"topview-ttk/internal/app/ttk-user/rpc/model"

	"topview-ttk/internal/app/ttk-user/rpc/internal/svc"
	"topview-ttk/internal/app/ttk-user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailPassLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEmailPassLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailPassLoginLogic {
	return &EmailPassLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EmailPassLoginLogic) EmailPassLogin(in *user.EmailPassLoginRequest) (*user.LoginResponse, error) {
	email := in.GetEmail()
	var userCredentials *model.TtkUserCredentials
	var err error

	if !util.ValidateEmail(email) {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.EmailValidError), "邮箱格式错误")

	}
	userCredentials, err = l.svcCtx.TtkUserInfoModel.FindUserCredentialsByEmail(l.ctx, email)

	if err != nil && errors.Is(err, sqlc.ErrNotFound) {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.PassportError), "邮箱或者密码错误")
	}

	if err != nil {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.DbError), "获取用户凭证数据失败，原因 : %v, 参数: %+v ", err, in)
	}
	if in.GetPass() == "" {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.PassportError), "邮箱或者密码错误")
	}

	inputPass := login.EncryptPasswordWithSalt(in.GetPass(), userCredentials.Salt.String)
	// test
	dbPass := login.EncryptPasswordWithSalt(userCredentials.Password, userCredentials.Salt.String)
	if dbPass != inputPass {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.PassportError), "邮箱或者密码错误")
	}

	userInfo, err := l.svcCtx.TtkUserInfoModel.FindOne(l.ctx, userCredentials.Id)

	if err != nil {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.DbError), "获取用户数据失败，原因: %v", err)
	}
	return &user.LoginResponse{
		UserInfo: &user.UserInfo{
			Id:       userInfo.Id,
			UserName: userInfo.TtkId,
			NickName: userInfo.NickName.String,
		},
	}, nil
}
