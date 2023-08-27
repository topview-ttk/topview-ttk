package ssoservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"time"
	"topview-ttk/internal/app/ttk-user/rpc/internal/svc"
	"topview-ttk/internal/app/ttk-user/rpc/user"
	"topview-ttk/internal/pkg/token"
	"topview-ttk/internal/pkg/ttkerr"

	"github.com/zeromicro/go-zero/core/logx"
)

const cacheTokenAndRefTokenKey = "cache:ttkLoginRefreshToken:"
const refMapExpires = 1 * time.Hour

type RefreshTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RefreshTokenLogic) RefreshToken(in *user.RefreshTokenRequest) (*user.RefreshTokenResponse, error) {
	t := in.Token
	uc, err := token.ParseToken(t)
	if err != nil {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.TokenExpireError), "会话过期，需重新登陆")

	}
	expiresTime := uc.ExpiresAt.UnixNano()
	nowUnixMs := time.Now().UnixNano()
	halfExpiresMs := (uc.ExpiresAt.UnixNano() - uc.IssuedAt.UnixNano()) / 2
	var vfToken string
	if nowUnixMs < expiresTime && nowUnixMs > expiresTime-halfExpiresMs {
		// 这里是为了防止网络异常，确保对原有Token的容忍性，保证刷新的成功
		oriRst, _ := l.svcCtx.Rdb.Get(l.ctx, cacheTokenAndRefTokenKey+t).Result()
		if oriRst != "" {
			return &user.RefreshTokenResponse{RefToken: oriRst}, nil
		}
		vfToken, err = token.GenerateVfToken(uc.DeviceInfo, uc.ClientInfo, uc.Uid)
		// 设置映射
		l.svcCtx.Rdb.Set(l.ctx, cacheTokenAndRefTokenKey+t, vfToken, refMapExpires)
		if err != nil {
			return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.TokenGenerateError), "刷新token失败，原因：%v", err)
		}

	}
	return &user.RefreshTokenResponse{RefToken: vfToken}, nil
}
