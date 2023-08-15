package user

import (
	"context"
	"topview-ttk/internal/app/ttk-user/rpc/user"

	"topview-ttk/internal/app/ttk-api/internal/svc"
	"topview-ttk/internal/app/ttk-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailOrTtkPassLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailOrTtkPassLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailOrTtkPassLoginLogic {
	return &EmailOrTtkPassLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailOrTtkPassLoginLogic) EmailOrTtkPassLogin(req *types.EmailOrTtkPassLoginRequest) (resp *types.LoginResponse, err error) {
	rpcResp, err := l.svcCtx.SsoClient.EmailOrTtkPassLogin(l.ctx, &user.EmailOrTTKPassLoginRequest{
		Email:      req.Email,
		TtkId:      req.TtkId,
		Pass:       req.Password,
		DeviceInfo: req.DeviceInfo,
		ClientInfo: req.ClientInfo,
	})

	if err != nil {
		logx.Error(err)
		return &types.LoginResponse{}, err
	}

	return &types.LoginResponse{
		StatusCode: int32(rpcResp.GetStatusCode().Number()),
		Message:    rpcResp.Message,
		// todo User_info
	}, err
}
