package main

import (
	"flag"
	"fmt"

	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/test/rpc/internal/config"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/test/rpc/internal/server"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/test/rpc/internal/svc"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/test/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/test.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterTestClientServer(grpcServer, server.NewTestClientServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
