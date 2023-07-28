package svc

import (
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/schedule/model"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/schedule/rpc/internal/config"
)

type ServiceContext struct {
	Config          config.Config
	ScheduleClient  model.ScheduleModel
	EntryFormClient model.EntryFormModel
	UserInfoClient  model.UserInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		ScheduleClient:  model.NewScheduleModel(c.MongoDBConf.URL, c.MongoDBConf.DB, "schedule"),
		EntryFormClient: model.NewEntryFormModel(c.MongoDBConf.URL, c.MongoDBConf.DB, "form"),
		UserInfoClient:  model.NewUserInfoModel(c.MongoDBConf.URL, c.MongoDBConf.DB, "userinfo"),
	}
}
