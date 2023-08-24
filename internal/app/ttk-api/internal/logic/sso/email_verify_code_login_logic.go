package sso

import (
	"context"
	"github.com/pkg/errors"
	"topview-ttk/internal/app/ttk-user/rpc/user"

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

	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	return &types.LoginResponse{
		StatusCode: int32(rpcResp.GetStatusCode().Number()),
		Message:    rpcResp.Message,
		Token:      resp.Token,
		// todo User_info
	}, err

}
