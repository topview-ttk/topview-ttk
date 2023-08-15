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

func (l *PhoneVerifyCodeLoginLogic) PhoneVerifyCodeLogin(req *types.PhoneVerifyCodeLoginRequest) (resp *types.SendVerificationCodeResponse, err error) {
	rpcResp, err := l.svcCtx.SsoClient.PhoneVerifyCodeLogin(l.ctx, &user.PhoneVerifyCodeLoginRequest{
		Phone:            req.Phone,
		VerificationCode: req.VerificationCode,
		DeviceInfo:       req.DeviceInfo,
		ClientInfo:       req.ClientInfo,
	})

	if err != nil {
		logx.Error(err)
		return &types.SendVerificationCodeResponse{}, err
	}

	return &types.SendVerificationCodeResponse{
		StatusCode: int32(rpcResp.GetStatusCode().Number()),
		Message:    rpcResp.Message,
		// todo User_info
	}, err
}
