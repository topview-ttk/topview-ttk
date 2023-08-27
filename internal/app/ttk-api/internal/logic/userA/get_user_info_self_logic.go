package userA

import (
	"context"
	"topview-ttk/internal/app/ttk-api/internal/svc"
	"topview-ttk/internal/app/ttk-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoSelfLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoSelfLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoSelfLogic {
	return &GetUserInfoSelfLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoSelfLogic) GetUserInfoSelf() (resp *types.GetUserSelfResponse, err error) {
	uid := l.ctx.Value("Uid")
	if _, ok := uid.(int64); ok {
	}
	// TODO 获取用户个人信息
	return
}
