package main

import (
	"MuXiFresh-Be-2.0/common/email"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/threading"

	"MuXiFresh-Be-2.0/app/form/api/internal/config"
	"MuXiFresh-Be-2.0/app/form/api/internal/handler"
	"MuXiFresh-Be-2.0/app/form/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/form-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	//加载邮箱配置
	email.Load(&c.EmailConf)
	//启动邮件Sender
	threading.GoSafe(email.Sender)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
