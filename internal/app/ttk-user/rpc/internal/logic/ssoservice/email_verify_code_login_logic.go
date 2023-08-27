package ssoservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"topview-ttk/internal/app/ttk-user/rpc/internal/logic/ssoservice/login"
	"topview-ttk/internal/app/ttk-user/rpc/internal/svc"
	"topview-ttk/internal/app/ttk-user/rpc/internal/util"
	"topview-ttk/internal/app/ttk-user/rpc/user"
	"topview-ttk/internal/pkg/ttkerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailVerifyCodeLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEmailVerifyCodeLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailVerifyCodeLoginLogic {
	return &EmailVerifyCodeLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EmailVerifyCodeLoginLogic) EmailVerifyCodeLogin(in *user.EmailVerifyCodeLoginRequest) (*user.LoginResponse, error) {
	isValid := util.ValidateEmail(in.GetEmail())
	if !isValid {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.EmailValidError), "邮箱格式错误")
	}
	code, err := l.svcCtx.Rdb.GetDel(l.ctx, "verification:"+in.GetEmail()).Result()
	if err != nil {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.VerifyCodeNotFoundError), "获取验证码失败, 原因: %v, 参数: %+v", err, in)
	}

	if code != in.GetVerificationCode() {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.VerifyCodeJudgeError), "验证码错误")
	}
	uc, err := l.svcCtx.TtkUserInfoModel.FindUserCredentialsByEmail(l.ctx, in.GetEmail())
	if err != nil && errors.Is(err, sqlc.ErrNotFound) {
		u := login.CreateDefaultUserInfo()
		u.Email = in.GetEmail()
		_, err := l.svcCtx.TtkUserInfoModel.Insert(l.ctx, u)
		if err != nil {
			return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.DbError), "邮箱注册失败, 原因: %v, 参数: %+v", err, in)
		}
		return &user.LoginResponse{
			Uid: u.Id,
		}, nil
	} else if err != nil {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.DbError), "获取用户信息失败, 原因: %v, 参数: %+v", err, in)
	}
	return &user.LoginResponse{
		Uid: uc.Id,
	}, nil
}
