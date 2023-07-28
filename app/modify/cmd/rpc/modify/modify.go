package main

import (
	"flag"
	"fmt"

	"MuXiFresh-Be-2.0/app/modify/cmd/rpc/modify/internal/config"
	"MuXiFresh-Be-2.0/app/modify/cmd/rpc/modify/internal/server"
	"MuXiFresh-Be-2.0/app/modify/cmd/rpc/modify/internal/svc"
	"MuXiFresh-Be-2.0/app/modify/cmd/rpc/modify/modify"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/modify.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		modify.RegisterModifyServer(grpcServer, server.NewModifyServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
