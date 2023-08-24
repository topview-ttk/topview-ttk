package userservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"topview-ttk/internal/app/ttk-user/rpc/model"
	"topview-ttk/internal/pkg/common/ttkerr"

	"topview-ttk/internal/app/ttk-user/rpc/internal/svc"
	"topview-ttk/internal/app/ttk-user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByTTKIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoByTTKIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByTTKIdLogic {
	return &GetUserInfoByTTKIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoByTTKIdLogic) GetUserInfoByTTKId(in *user.GetUserInfoByTTKIdRequest) (*user.GetUserInfoResponse, error) {

	u, err := l.svcCtx.TtkUserInfoModel.FindOneByTtkId(l.ctx, in.TtkId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.DbError), "获取用户数据失败，原因: %v ,参数 :%+v", err, in)
	}
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.UserNotFountError), "用户不存在 ,参数 :%+v", in)
	}
	return &user.GetUserInfoResponse{UserInfo: &user.UserInfo{
		Id:       u.Id,
		UserName: u.TtkId,
		NickName: u.NickName.String,
	}}, nil
}
