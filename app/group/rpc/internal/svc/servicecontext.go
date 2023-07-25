package svc

import (
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/group/model"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/group/rpc/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	GroupClient model.GroupModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		GroupClient: model.NewGroupModel(c.MongoDBConf.URL, c.MongoDBConf.DB, "group"),
	}
}
