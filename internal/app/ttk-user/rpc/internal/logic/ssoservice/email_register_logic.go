package ssoservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"topview-ttk/internal/app/ttk-user/rpc/internal/svc"
	"topview-ttk/internal/app/ttk-user/rpc/internal/util"
	"topview-ttk/internal/app/ttk-user/rpc/model"
	"topview-ttk/internal/app/ttk-user/rpc/user"
	"topview-ttk/internal/pkg/common/ttkerr"

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

func (l *EmailRegisterLogic) EmailRegister(in *user.EmailRegisterRequest) (*user.RegisterResponse, error) {
	email := in.GetEmail()
	var userCredentials *model.TtkUserCredentials
	var err error

	if util.ValidateEmail(email) {
		userCredentials, err = l.svcCtx.TtkUserInfoModel.FindUserCredentialsByEmail(l.ctx, email)
	} else {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.EmailValidError), "邮箱格式错误")
	}
	if err != nil && !errors.Is(err, sqlc.ErrNotFound) {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.EmailRegisteredError), "邮箱已被注册")
	}

	if err != nil {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.DbError), "邮箱已被注册")
	}

	//inputPass := login.EncryptPasswordWithSalt(in.GetPassword(), userCredentials.Salt.String)
	//inputNickname := in.GetNickname()

	_, err = l.svcCtx.TtkUserInfoModel.FindOne(l.ctx, userCredentials.Id)
	if err != nil {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.DbError), "邮箱注册失败，原因： %v , 参数: %+v", err, in)
	}
	return &user.RegisterResponse{}, nil

}
