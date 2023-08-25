package ssoservicelogic

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"topview-ttk/internal/app/ttk-user/rpc/internal/logic/ssoservice/login"
	"topview-ttk/internal/app/ttk-user/rpc/internal/svc"
	"topview-ttk/internal/app/ttk-user/rpc/internal/util"
	"topview-ttk/internal/app/ttk-user/rpc/user"
	"topview-ttk/internal/pkg/common/ttkerr"

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
	userInfo, err := l.svcCtx.TtkUserInfoModel.FindOneByEmail(l.ctx, in.GetEmail())
	if err != nil && errors.Is(err, sqlc.ErrNotFound) {
		userInfo = login.CreateDefaultUserInfo()
		userInfo.Email = sql.NullString{String: in.GetEmail(), Valid: true}
		_, err := l.svcCtx.TtkUserInfoModel.Insert(l.ctx, userInfo)
		if err != nil {
			return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.DbError), "邮箱注册失败, 原因: %v, 参数: %+v", err, in)
		}
		// 为了进行缓存用户数据，需要查询数据库，如果err大概率插入失败，让用户重新登录
		userInfo, err = l.svcCtx.TtkUserInfoModel.FindOneByEmail(l.ctx, in.GetEmail())
		if err != nil {
			return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.DbError), "获取用户信息失败, 原因: %v, 参数: %+v", err, in)
		}
	} else if err != nil {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.DbError), "获取用户信息失败, 原因: %v, 参数: %+v", err, in)
	}
	return &user.LoginResponse{
		UserInfo: &user.UserInfo{
			Id:       userInfo.Id,
			UserName: userInfo.TtkId,
			NickName: userInfo.NickName.String,
		},
	}, nil
}
