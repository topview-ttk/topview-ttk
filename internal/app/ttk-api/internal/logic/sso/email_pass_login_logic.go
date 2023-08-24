package sso

import (
	"context"
	"github.com/pkg/errors"
	"topview-ttk/internal/app/ttk-user/rpc/user"

	"topview-ttk/internal/app/ttk-api/internal/svc"
	"topview-ttk/internal/app/ttk-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailPassLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailPassLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailPassLoginLogic {
	return &EmailPassLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailPassLoginLogic) EmailPassLogin(req *types.EmailLoginRequest) (resp *types.LoginResponse, err error) {

	rpcResp, err := l.svcCtx.SsoClient.EmailPassLogin(l.ctx, &user.EmailPassLoginRequest{
		Email:      req.Email,
		Pass:       req.Password,
		DeviceInfo: req.DeviceInfo,
		ClientInfo: req.ClientInfo,
	})

	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	return &types.LoginResponse{
		StatusCode: int32(rpcResp.GetStatusCode().Number()),
		Message:    rpcResp.Message,
		Token:      resp.Token,
	}, err
}
