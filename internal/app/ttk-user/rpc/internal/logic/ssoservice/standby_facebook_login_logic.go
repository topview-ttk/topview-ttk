package ssoservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"
	"time"
	"topview-ttk/internal/app/ttk-user/model"
	"topview-ttk/internal/app/ttk-user/rpc/internal/logic/ssoservice/login"
	"topview-ttk/internal/pkg/database"
	"topview-ttk/internal/pkg/ttkerr"

	"topview-ttk/internal/app/ttk-user/rpc/internal/svc"
	"topview-ttk/internal/app/ttk-user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type StandbyFacebookLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStandbyFacebookLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StandbyFacebookLoginLogic {
	return &StandbyFacebookLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StandbyFacebookLoginLogic) StandbyFacebookLogin(in *user.StandbyLoginRequest) (*user.LoginResponse, error) {
	standbyInfo := in.GetStandbyInfo()

	num_ID, _ := strconv.ParseInt(standbyInfo.GetThirdPartyId(), 10, 64)

	id, err := l.svcCtx.TtkThirdPartyBindingModel.FindUserIdByThirdPartyIdAndType(l.ctx, num_ID, "google")

	if err != nil && errors.Is(err, sqlc.ErrNotFound) {
		if err := database.TransCtx(l.ctx, l.svcCtx.SqlConn, func(ctx context.Context, session sqlx.Session) error {
			userInfo := login.CreateDefaultUserInfo()
			userInfo.AvatarPath = standbyInfo.GetAvatarUrl()
			userInfo.NickName = standbyInfo.GetNickname()
			userInfo.Email = standbyInfo.GetEmail()

			_, err := l.svcCtx.TtkUserInfoModel.TransSaveCtx(ctx, session, userInfo)
			if err != nil {
				return err
			}
			uid := userInfo.Id
			if err != nil {
				return err
			}
			githubBinding := &model.TtkThirdPartyBinding{
				UserId:                uid,
				ThirdPartyBindingType: 2,
				ThirdPartyId:          standbyInfo.GetThirdPartyId(),
				CreatedAt:             time.Time{},
				UpdatedAt:             time.Now(),
			}
			_, err = l.svcCtx.TtkThirdPartyBindingModel.TransSaveCtx(ctx, session, githubBinding)
			if err != nil {
				return err
			}
			id = uid
			return nil
		}); err != nil {
			return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.DbError), "第三方绑定失败, 原因: %v, 参数: %+v", err, in)
		}
	} else if err != nil {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.DbError), "获取用户信息失败, 原因: %v, 参数: %+v", err, in)
	}
	return &user.LoginResponse{
		Uid: id,
	}, nil
}
