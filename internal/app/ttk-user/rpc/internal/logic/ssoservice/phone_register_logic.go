package ssoservicelogic

import (
	"context"

	"topview-ttk/internal/app/ttk-user/rpc/internal/svc"
	"topview-ttk/internal/app/ttk-user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhoneRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPhoneRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhoneRegisterLogic {
	return &PhoneRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PhoneRegisterLogic) PhoneRegister(in *user.PhoneRegisterRequest) (*user.RegisterResponse, error) {
	// todo: add your logic here and delete this line

	return &user.RegisterResponse{}, nil
}
