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

type FacebookLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFacebookLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FacebookLoginLogic {
	return &FacebookLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FacebookLoginLogic) FacebookLogin(req *types.ThirdPartyLoginRequest) (resp *types.LoginResponse, err error) {
	var rpcLoginCommon = &user.LoginCommon{}
	err = copier.Copy(rpcLoginCommon, &req.LoginCommon)

	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	rpcResp, err := l.svcCtx.SsoClient.FacebookLogin(l.ctx, &user.ThirdPartyLoginRequest{
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
