package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"topview-ttk/internal/app/ttk-api/internal/config"
	"topview-ttk/internal/app/ttk-user/rpc/client/ssoservice"
	"topview-ttk/internal/app/ttk-user/rpc/client/userservice"
)

type ServiceContext struct {
	Config     config.Config
	UserClient userservice.UserService
	SsoClient  ssoservice.SsoService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UserClient: userservice.NewUserService(zrpc.MustNewClient(c.User)),
		SsoClient:  ssoservice.NewSsoService(zrpc.MustNewClient(c.User)),
	}
}
