package ssoservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"topview-ttk/internal/app/ttk-user/rpc/internal/logic/ssoservice/login"
	"topview-ttk/internal/app/ttk-user/rpc/internal/svc"
	"topview-ttk/internal/app/ttk-user/rpc/internal/util"
	"topview-ttk/internal/app/ttk-user/rpc/user"
	"topview-ttk/internal/pkg/common/ttkerr"

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
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.PhoneValidError), "手机格式错误")
	}
	userCredentials, err := l.svcCtx.TtkUserInfoModel.FindUserCredentialsByPhone(l.ctx, phone)
	if err != nil && errors.Is(err, sqlc.ErrNotFound) {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.PassportError), "手机或者密码错误")
	}

	if err != nil {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.DbError), "获取用户凭证数据失败，原因 : %v, 参数: %+v ", err, in)
	}
	if in.GetPass() == "" {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.PassportError), "手机或者密码错误")
	}

	inputPass := login.EncryptPasswordWithSalt(in.GetPass(), userCredentials.Salt.String)
	// test
	dbPass := login.EncryptPasswordWithSalt(userCredentials.Password, userCredentials.Salt.String)

	if dbPass != inputPass {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.PassportError), "手机或者密码错误")
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
