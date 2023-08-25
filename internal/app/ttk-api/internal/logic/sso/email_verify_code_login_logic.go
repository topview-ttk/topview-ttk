package sso

import (
	"context"
	"github.com/pkg/errors"
	"topview-ttk/internal/app/ttk-user/rpc/user"
	"topview-ttk/internal/pkg/common/token"

	"topview-ttk/internal/app/ttk-api/internal/svc"
	"topview-ttk/internal/app/ttk-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailVerifyCodeLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailVerifyCodeLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailVerifyCodeLoginLogic {
	return &EmailVerifyCodeLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailVerifyCodeLoginLogic) EmailVerifyCodeLogin(req *types.EmailVerifyCodeLoginRequest) (resp *types.LoginResponse, err error) {
	rpcResp, err := l.svcCtx.SsoClient.EmailVerifyCodeLogin(l.ctx, &user.EmailVerifyCodeLoginRequest{
		Email:            req.Email,
		VerificationCode: req.VerificationCode,
		DeviceInfo:       req.DeviceInfo,
		ClientInfo:       req.ClientInfo,
	})
	_ = rpcResp.UserInfo.UserName

	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	t, err := token.GenerateVfToken(req.DeviceInfo, req.ClientInfo, rpcResp.UserInfo.Id)
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.LoginResponse{
		Token: t,
	}, err

}
