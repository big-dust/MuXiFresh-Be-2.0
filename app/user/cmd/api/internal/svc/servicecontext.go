package svc

import (
	"MuXiFresh-Be-2.0/app/user/cmd/api/internal/config"
	"MuXiFresh-Be-2.0/app/user/cmd/rpc/user/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	UserClient userclient.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UserClient: userclient.NewUserClient(zrpc.MustNewClient(c.UserConf)),
	}
}
