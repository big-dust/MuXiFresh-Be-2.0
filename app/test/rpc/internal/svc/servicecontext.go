package svc

import (
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/form/model"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/test/rpc/internal/config"
)

type ServiceContext struct {
	Config          config.Config
	UserInfoClient  model.UserInfoModel
	EntryFormClient model.EntryFormModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		UserInfoClient:  model.NewUserInfoModel(c.MongoDBConf.URL, c.MongoDBConf.DB, "userinfo"),
		EntryFormClient: model.NewEntryFormModel(c.MongoDBConf.URL, c.MongoDBConf.DB, "form"),
	}
}
