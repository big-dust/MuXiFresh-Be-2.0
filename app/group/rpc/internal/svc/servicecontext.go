package svc

import (
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/group/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
