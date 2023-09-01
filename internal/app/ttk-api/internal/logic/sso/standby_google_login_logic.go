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

type StandbyGoogleLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStandbyGoogleLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StandbyGoogleLoginLogic {
	return &StandbyGoogleLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StandbyGoogleLoginLogic) StandbyGoogleLogin(req *types.StandbyLoginRequest) (resp *types.LoginResponse, err error) {
	var rpcClientInfo = &user.ClientInfo{}
	var rpcStandbyUserInfo = &user.StandbyUserInfo{}

	err = copier.Copy(rpcClientInfo, &req.ClientInfo)
	err = copier.Copy(rpcStandbyUserInfo, &req.StandbyUserInfo)

	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	rpcResp, err := l.svcCtx.SsoClient.StandbyFacebookLogin(l.ctx, &user.StandbyLoginRequest{
		StandbyInfo: rpcStandbyUserInfo,
		ClientInfo:  rpcClientInfo,
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
