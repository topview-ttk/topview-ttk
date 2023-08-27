package sso

import (
	"context"
	"github.com/pkg/errors"
	"topview-ttk/internal/app/ttk-user/rpc/user"

	"topview-ttk/internal/app/ttk-api/internal/svc"
	"topview-ttk/internal/app/ttk-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailRegisterLogic {
	return &EmailRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailRegisterLogic) EmailRegister(req *types.EmailRegisterRequest) (resp *types.RegisterResponse, err error) {
	_, err = l.svcCtx.SsoClient.EmailRegister(l.ctx, &user.EmailRegisterRequest{
		Email:    req.Email,
		NickName: req.NickName,
		Password: req.Password,
	})

	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	return &types.RegisterResponse{}, err
}
