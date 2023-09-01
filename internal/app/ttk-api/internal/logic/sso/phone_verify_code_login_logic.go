package sso

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"topview-ttk/internal/app/ttk-user/rpc/user"
	"topview-ttk/internal/pkg/token"

	"topview-ttk/internal/app/ttk-api/internal/svc"
	"topview-ttk/internal/app/ttk-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhoneVerifyCodeLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPhoneVerifyCodeLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhoneVerifyCodeLoginLogic {
	return &PhoneVerifyCodeLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PhoneVerifyCodeLoginLogic) PhoneVerifyCodeLogin(req *types.PhoneVerifyCodeLoginRequest) (resp *types.LoginResponse, err error) {
	var rpcClientInfo = &user.ClientInfo{}
	err = copier.Copy(rpcClientInfo, &req.ClientInfo)

	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	rpcResp, err := l.svcCtx.SsoClient.PhoneVerifyCodeLogin(l.ctx, &user.PhoneVerifyCodeLoginRequest{
		Phone:            req.Phone,
		VerificationCode: req.VerificationCode,
		ClientInfo:       rpcClientInfo,
	})

	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	t, err := token.GenerateVfToken(req.ClientInfo.DeviceInfo, req.ClientInfo.OSVersion, rpcResp.Uid)
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.LoginResponse{
		Token:        t,
		TokenExpires: token.TokenExpires.Nanoseconds(),
	}, err
}
