package user

import (
	"context"
	"topview-ttk/internal/app/ttk-user/rpc/user"

	"topview-ttk/internal/app/ttk-api/internal/svc"
	"topview-ttk/internal/app/ttk-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhoneVerifyCodeLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPhoneVerifyCodeLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhoneVerifyCodeLoginLogic {
	return &PhoneVerifyCodeLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PhoneVerifyCodeLoginLogic) PhoneVerifyCodeLogin(req *types.PhoneVerifyCodeLoginRequest) (resp *types.PhoneVerifyCodeLoginResponse, err error) {
	rpcResp, err := l.svcCtx.SsoClient.PhoneVerifyCodeLogin(l.ctx, &user.PhoneVerifyCodeLoginRequest{
		Phone:            req.Phone,
		VerificationCode: req.Verification_code,
		DeviceInfo:       req.Device_info,
		ClientInfo:       req.Client_info,
	})

	if err != nil {
		logx.Error(err)
		return &types.PhoneVerifyCodeLoginResponse{}, err
	}

	return &types.PhoneVerifyCodeLoginResponse{
		Status_code: rpcResp.StatusCode,
		Message:     rpcResp.Message,
		// todo User_info
	}, err
	return
}
