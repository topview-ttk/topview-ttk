package userA

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"topview-ttk/internal/app/ttk-api/internal/svc"
	"topview-ttk/internal/app/ttk-api/internal/types"
	"topview-ttk/internal/app/ttk-user/rpc/user"

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
	uid, _ := l.ctx.Value("Uid").(json.Number).Int64()
	rpcResp, err := l.svcCtx.UserClient.GetUserSelfInfo(l.ctx, &user.GetUserInfoByUidRequest{Uid: uid})
	if err != nil {
		return nil, err
	}
	var u = types.UserInfo{}
	err = copier.Copy(&u, rpcResp.UserInfo)
	if err != nil {
		return nil, err
	}
	return &types.GetUserSelfResponse{UserInfo: u}, nil
}
