package main

import (
	"flag"
	"fmt"
	ssoserver "topview-ttk/internal/app/ttk-user/rpc/internal/server/ssoservice"
	userserver "topview-ttk/internal/app/ttk-user/rpc/internal/server/userservice"
	"topview-ttk/internal/pkg/interceptor/rpc"

	"topview-ttk/internal/app/ttk-user/rpc/internal/config"
	"topview-ttk/internal/app/ttk-user/rpc/internal/svc"
	"topview-ttk/internal/app/ttk-user/rpc/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "internal/app/ttk-user/rpc/etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServiceServer(grpcServer, userserver.NewUserServiceServer(ctx))
		user.RegisterSsoServiceServer(grpcServer, ssoserver.NewSsoServiceServer(ctx))
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	// 添加rpc侧日志拦截器
	s.AddUnaryInterceptors(rpc.LoggerInterceptor)
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
