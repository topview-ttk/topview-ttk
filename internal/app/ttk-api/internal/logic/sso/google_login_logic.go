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

type GoogleLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoogleLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoogleLoginLogic {
	return &GoogleLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoogleLoginLogic) GoogleLogin(req *types.ThirdPartyLoginRequest) (resp *types.LoginResponse, err error) {
	var rpcLoginCommon = &user.LoginCommon{}
	err = copier.Copy(rpcLoginCommon, &req.LoginCommon)

	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	rpcResp, err := l.svcCtx.SsoClient.GoogleLogin(l.ctx, &user.ThirdPartyLoginRequest{
		AccessToken: req.Token,
		LoginCommon: rpcLoginCommon,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	t, err := token.GenerateVfToken(req.LoginCommon.DeviceInfo, req.LoginCommon.ClientInfo, rpcResp.Uid)
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.LoginResponse{
		Token:        t,
		TokenExpires: token.TokenExpires.Nanoseconds(),
	}, err
}
