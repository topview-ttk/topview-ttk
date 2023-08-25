package ssoservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"topview-ttk/internal/app/ttk-user/rpc/internal/logic/ssoservice/login"
	"topview-ttk/internal/app/ttk-user/rpc/internal/svc"
	"topview-ttk/internal/app/ttk-user/rpc/internal/util"
	"topview-ttk/internal/app/ttk-user/rpc/model"
	"topview-ttk/internal/app/ttk-user/rpc/user"
	"topview-ttk/internal/pkg/common/ttkerr"

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

	if !util.ValidateTtkId(Ttkid) {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.TTIdValidError), "TTKId格式错误")
	}
	userCredentials, err = l.svcCtx.TtkUserInfoModel.FindUserCredentialsByTtkId(l.ctx, Ttkid)

	if err != nil && errors.Is(err, sqlc.ErrNotFound) {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.PassportError), "TTKId或者密码错误")
	}

	if err != nil {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.DbError), "获取用户凭证数据失败，原因 : %v, 参数: %+v ", err, in)
	}
	if in.GetPass() == "" {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.PassportError), "TTKId或者密码错误")
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
