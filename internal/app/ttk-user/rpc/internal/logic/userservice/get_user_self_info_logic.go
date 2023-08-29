package userservicelogic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"topview-ttk/internal/app/ttk-user/model"
	"topview-ttk/internal/pkg/ttkerr"

	"topview-ttk/internal/app/ttk-user/rpc/internal/svc"
	"topview-ttk/internal/app/ttk-user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserSelfInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserSelfInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserSelfInfoLogic {
	return &GetUserSelfInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserSelfInfoLogic) GetUserSelfInfo(in *user.GetUserInfoByUidRequest) (*user.GetUserInfoResponse, error) {
	u, err := l.svcCtx.TtkUserInfoModel.FindOne(l.ctx, in.Uid)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.DbError), "获取用户数据失败，原因: %v ,参数 :%+v", err, in)
	}
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.UserNotFountError), "用户不存在 ,参数 :%+v", in)
	}
	var ur = &user.UserInfo{}
	err = copier.Copy(ur, u)

	if err != nil {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.ServerCommonError), "用户数据映射失败 ,参数 :%+v", u)
	}
	return &user.GetUserInfoResponse{UserInfo: ur}, nil
}
