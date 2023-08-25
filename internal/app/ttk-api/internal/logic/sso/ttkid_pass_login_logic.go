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

type TtkidPassLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTtkidPassLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TtkidPassLoginLogic {
	return &TtkidPassLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TtkidPassLoginLogic) TtkidPassLogin(req *types.TTkIDLoginRequest) (resp *types.LoginResponse, err error) {
	rpcResp, err := l.svcCtx.SsoClient.TtkidPassLogin(l.ctx, &user.TTKPassLoginRequest{
		TtkId:      req.TTkId,
		Pass:       req.Password,
		DeviceInfo: req.DeviceInfo,
		ClientInfo: req.ClientInfo,
	})

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
