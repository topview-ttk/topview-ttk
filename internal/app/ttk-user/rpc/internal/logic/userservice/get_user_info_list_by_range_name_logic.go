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

type GetUserInfoListByRangeNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoListByRangeNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoListByRangeNameLogic {
	return &GetUserInfoListByRangeNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoListByRangeNameLogic) GetUserInfoListByRangeName(in *user.GetUserInfoListByRangeNameRequest) (*user.GetUserInfoListByRangeNameResponse, error) {
	var data = &model.FindUserList{}
	data.RangeName = in.RangeName

	data.Offset = int((in.GetPage().GetCurrentPage() - 1) * in.GetPage().GetPageSize())
	data.Limit = int(in.GetPage().GetPageSize())
	modelUserList, err := l.svcCtx.TtkUserInfoModel.FindListByRangeName(l.ctx, data)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.DbError), "获取用户数据失败，原因: %v ,参数 :%+v", err, in)
	}
	var resp []*user.UserInfo
	for _, u := range modelUserList {
		var ur = &user.UserInfo{}
		err := copier.Copy(ur, u)
		if err != nil {
			return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.ServerCommonError), "这个问题几乎不可能，原因：%v,参数：%+v", err, in)
		}
		resp = append(resp, ur)
	}
	return &user.GetUserInfoListByRangeNameResponse{
		UserInfoList: resp,
	}, nil
}
