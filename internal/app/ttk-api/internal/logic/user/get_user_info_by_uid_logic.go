package user

import (
	"context"
	"github.com/pkg/errors"
	"topview-ttk/internal/app/ttk-user/rpc/user"

	"topview-ttk/internal/app/ttk-api/internal/svc"
	"topview-ttk/internal/app/ttk-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByUidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoByUidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByUidLogic {
	return &GetUserInfoByUidLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoByUidLogic) GetUserInfoByUid(req *types.GetUserInfoByUidRequest) (resp *types.GetUserInfoResponse, err error) {
	rpcResp, err := l.svcCtx.UserClient.GetUserInfoByUid(l.ctx, &user.GetUserInfoByUidRequest{
		Uid: req.Uid,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.GetUserInfoResponse{UserInfo: types.UserInfo{TTkId: rpcResp.UserInfo.GetUserName()}}, nil
}
