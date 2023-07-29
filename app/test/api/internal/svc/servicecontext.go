package svc

import (
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/test/api/internal/config"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/test/rpc/testclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	TestClient testclient.TestClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		TestClient: testclient.NewTestClient(zrpc.MustNewClient(c.TestConf)),
	}
}
