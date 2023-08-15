package user

import (
	"context"
	"topview-ttk/internal/app/ttk-user/rpc/user"

	"topview-ttk/internal/app/ttk-api/internal/svc"
	"topview-ttk/internal/app/ttk-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailRegisterLogic {
	return &EmailRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailRegisterLogic) EmailRegister(req *types.EmailRegisterRequest) (resp *types.RegisterResponse, err error) {
	rpcResp, err := l.svcCtx.SsoClient.EmailRegister(l.ctx, &user.EmailRegisterRequest{
		Email:    req.Email,
		Nickname: req.NickName,
		Password: req.Password,
	})

	if err != nil {
		logx.Error(err)
		return &types.RegisterResponse{}, err
	}

	return &types.RegisterResponse{
		StatusCode: int32(rpcResp.GetStatusCode().Number()),
		Message:    rpcResp.Message,
		// todo User_info
	}, err
}
