package sso

import (
	"context"
	"topview-ttk/internal/app/ttk-user/rpc/user"

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
		logx.Error(err)
		return &types.LoginResponse{}, err
	}

	return &types.LoginResponse{
		StatusCode: int32(rpcResp.GetStatusCode().Number()),
		Message:    rpcResp.Message,
		UserInfo:   types.UserInfo{TTkId: rpcResp.UserInfo.GetUserName()},
		Token:      rpcResp.Token,
	}, err
}
