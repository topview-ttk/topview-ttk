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

type PhonePassLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPhonePassLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhonePassLoginLogic {
	return &PhonePassLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PhonePassLoginLogic) PhonePassLogin(req *types.PhonePassLoginRequest) (resp *types.LoginResponse, err error) {
	var rpcClientInfo = &user.ClientInfo{}
	err = copier.Copy(rpcClientInfo, &req.ClientInfo)

	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	rpcResp, err := l.svcCtx.SsoClient.PhonePassLogin(l.ctx, &user.PhonePassLoginRequest{
		Phone:      req.Phone,
		Pass:       req.Password,
		ClientInfo: rpcClientInfo,
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
