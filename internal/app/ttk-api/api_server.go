package main

import (
	"flag"
	"fmt"
	"topview-ttk/internal/app/ttk-api/internal/config"
	"topview-ttk/internal/app/ttk-api/internal/handler"
	"topview-ttk/internal/app/ttk-api/internal/middleware"
	"topview-ttk/internal/app/ttk-api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "internal/app/ttk-api/etc/api_server.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf,
		// 跨域
		rest.WithCors(),
		// Jwt鉴权
		rest.WithUnauthorizedCallback(middleware.NewUnAuthorizedMiddleware().Callback()))
	defer server.Stop()
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
