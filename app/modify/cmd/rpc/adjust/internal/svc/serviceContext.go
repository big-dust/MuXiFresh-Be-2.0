package svc

import "MuXiFresh-Be-2.0/app/modify/cmd/rpc/adjust/internal/config"

type ServiceContext struct {
	Config config.Config
	scheduleClient model.schedule
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		scheduleClient:model.NewscheduleModel(c.MongoConf.URL,c.MongoConf.DB,collection:"schedule")
	}
}
