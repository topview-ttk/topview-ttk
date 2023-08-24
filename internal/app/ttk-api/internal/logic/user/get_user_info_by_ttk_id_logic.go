package user

import (
	"context"
	"github.com/pkg/errors"
	"topview-ttk/internal/app/ttk-user/rpc/user"

	"topview-ttk/internal/app/ttk-api/internal/svc"
	"topview-ttk/internal/app/ttk-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByTTKIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoByTTKIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByTTKIdLogic {
	return &GetUserInfoByTTKIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoByTTKIdLogic) GetUserInfoByTTKId(req *types.GetUserInfoByTTKIdRequest) (resp *types.GetUserInfoResponse, err error) {
	rpcResp, err := l.svcCtx.UserClient.GetUserInfoByUid(l.ctx, &user.GetUserInfoByUidRequest{
		Uid: req.TTKId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.GetUserInfoResponse{UserInfo: types.UserInfo{TTkId: rpcResp.UserInfo.GetUserName()}}, nil
}
