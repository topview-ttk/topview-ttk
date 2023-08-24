package ssoservicelogic

import (
	"context"
	"errors"
	"time"
	"topview-ttk/internal/app/ttk-user/rpc/internal/logic/ssoservice/login"
	"topview-ttk/internal/pkg/common"

	"topview-ttk/internal/app/ttk-user/rpc/internal/svc"
	"topview-ttk/internal/app/ttk-user/rpc/user"

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
	token := in.Token
	uc, err := common.ParseToken(token)
	if err != nil {
		logx.Error("Token解析出错", err)
		return handleRefreshTokenError("会话过期，请重新登录"), err
	}
	expiresTime := uc.ExpiresAt.UnixNano()
	nowUnixMs := time.Now().UnixNano()
	halfExpiresMs := (uc.ExpiresAt.UnixNano() - uc.IssuedAt.UnixNano()) / 2
	if nowUnixMs < expiresTime && nowUnixMs > expiresTime-halfExpiresMs {
		// 这里是为了防止网络异常，确保对原有Token的容忍性，保证刷新的成功
		oriRst, _ := l.svcCtx.Rdb.Get(l.ctx, cacheTokenAndRefTokenKey+token).Result()
		if oriRst != "" {
			return &user.RefreshTokenResponse{StatusCode: user.StatusCode_OK, RefToken: oriRst}, nil
		}
		vfToken, err := login.GenerateVfToken(uc.DeviceInfo, uc.ClientInfo, uc.Uid)
		// 设置映射
		l.svcCtx.Rdb.Set(l.ctx, cacheTokenAndRefTokenKey+token, vfToken, refMapExpires)
		if err != nil {
			logx.Error("刷新Token失败", err)
			return handleRefreshTokenError("刷新Token失败"), err
		}
		return &user.RefreshTokenResponse{StatusCode: user.StatusCode_OK, RefToken: vfToken}, nil

	}
	err = errors.New("没过期，不要刷新")
	return handleRefreshTokenError("没过期，不要刷新"), err
}

func handleRefreshTokenError(message string) *user.RefreshTokenResponse {
	return &user.RefreshTokenResponse{
		StatusCode: user.StatusCode_INVALID_ARGUMENT,
		Message:    message,
	}
}
