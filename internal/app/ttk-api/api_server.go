package main

import (
	"flag"
	"fmt"

	"topview-ttk/internal/app/ttk-api/internal/config"
	"topview-ttk/internal/app/ttk-api/internal/handler"
	"topview-ttk/internal/app/ttk-api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "internal/app/ttk-api/etc/api_server.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
