package user

import (
	"context"
	"topview-ttk/internal/app/ttk-user/rpc/user"

	"topview-ttk/internal/app/ttk-api/internal/svc"
	"topview-ttk/internal/app/ttk-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GithubLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGithubLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GithubLoginLogic {
	return &GithubLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GithubLoginLogic) GithubLogin(req *types.GithubLoginRequest) (resp *types.LoginResponse, err error) {
	rpcResp, err := l.svcCtx.SsoClient.GithubLogin(l.ctx, &user.GitHubLoginRequest{
		Token		:    req.Token,
		DeviceInfo	: 	 req.DeviceInfo,
		ClientInfo	: 	 req.ClientInfo,
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
