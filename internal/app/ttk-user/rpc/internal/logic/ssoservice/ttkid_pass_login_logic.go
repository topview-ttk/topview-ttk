package ssoservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"topview-ttk/internal/app/ttk-user/model"
	"topview-ttk/internal/app/ttk-user/rpc/internal/logic/ssoservice/login"
	"topview-ttk/internal/app/ttk-user/rpc/internal/svc"
	"topview-ttk/internal/app/ttk-user/rpc/internal/util"
	"topview-ttk/internal/app/ttk-user/rpc/user"
	"topview-ttk/internal/pkg/ttkerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type TtkIdPassLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTtkIdPassLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TtkIdPassLoginLogic {
	return &TtkIdPassLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TtkIdPassLoginLogic) TtkIdPassLogin(in *user.TTKIdPassLoginRequest) (*user.LoginResponse, error) {
	Ttkid := in.GetTtkId()
	var userCredentials *model.TtkUserCredentials
	var err error

	if !util.ValidateTtkId(Ttkid) {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.TTKIdValidError), "TTKId格式错误")
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

	inputPass := login.EncryptPasswordWithSalt(in.GetPass(), userCredentials.Salt)
	// test
	dbPass := login.EncryptPasswordWithSalt(userCredentials.Password, userCredentials.Salt)
	if dbPass != inputPass {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.PassportError), "邮箱或者密码错误")
	}

	return &user.LoginResponse{
		Uid: userCredentials.Id,
	}, nil
}
