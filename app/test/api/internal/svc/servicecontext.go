package svc

import (
	"MuXiFresh-Be-2.0/app/test/api/internal/config"
	"MuXiFresh-Be-2.0/app/test/rpc/testclient"
	"MuXiFresh-Be-2.0/app/user/cmd/rpc/user/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	TestClient testclient.TestClient
	UserClient userclient.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		TestClient: testclient.NewTestClient(zrpc.MustNewClient(c.TestConf)),
		UserClient: userclient.NewUserClient(zrpc.MustNewClient(c.UserConf)),
	}
}
