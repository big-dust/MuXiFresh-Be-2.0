package svc

import (
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/group/api/internal/config"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/group/rpc/getclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	IntroClient getclient.GetClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		IntroClient: getclient.NewGetClient(zrpc.MustNewClient(c.IntroConf)),
	}
}
