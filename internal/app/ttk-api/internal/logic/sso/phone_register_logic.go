package sso

import (
	"context"
	"topview-ttk/internal/app/ttk-user/rpc/user"

	"topview-ttk/internal/app/ttk-api/internal/svc"
	"topview-ttk/internal/app/ttk-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhoneRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPhoneRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhoneRegisterLogic {
	return &PhoneRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PhoneRegisterLogic) PhoneRegister(req *types.PhoneRegisterRequest) (resp *types.RegisterResponse, err error) {
	_, err = l.svcCtx.SsoClient.PhoneRegister(l.ctx, &user.PhoneRegisterRequest{
		NickName: req.NickName,
		Phone:    req.Phone,
		Password: req.Password,
	})

	if err != nil {
		logx.Error(err)
		return &types.RegisterResponse{}, err
	}

	return &types.RegisterResponse{}, err
}
