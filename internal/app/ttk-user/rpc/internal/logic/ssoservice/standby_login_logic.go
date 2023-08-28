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

type StandbyUserInfo struct {
	Nickname  string `json:"nickname"`
	ID        int64  `json:"id"`
	AvatarURL string `json:"avatar_url"`
	LoginType int64  `json:"loginType"`
}

type StandbyLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStandbyLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StandbyLoginLogic {
	return &StandbyLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StandbyLoginLogic) StandbyLogin(in *user.StandbyLoginRequest) (*user.LoginResponse, error) {
	num_ID, _ := strconv.ParseInt(in.GetThirdPartyId(), 10, 64)

	standbyUserInfo := StandbyUserInfo{
		Nickname:  in.GetNickname(),
		ID:        num_ID,
		AvatarURL: in.GetAvatarUrl(),
		LoginType: in.LoginType,
	}

	id, err := l.svcCtx.TtkThirdPartyBindingModel.FindUserIdByThirdPartyIdAndType(l.ctx, standbyUserInfo.ID, "github")

	if err != nil && errors.Is(err, sqlc.ErrNotFound) {
		if err := database.TransCtx(l.ctx, l.svcCtx.SqlConn, func(ctx context.Context, session sqlx.Session) error {
			userInfo := login.CreateDefaultUserInfo()
			userInfo.AvatarPath = standbyUserInfo.AvatarURL
			userInfo.NickName = standbyUserInfo.Nickname

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
				ThirdPartyBindingType: standbyUserInfo.LoginType,
				ThirdPartyId:          strconv.FormatInt(standbyUserInfo.ID, 10),
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

	return &user.LoginResponse{}, nil
}
