package svc

import (
	formmodel "MuXiFresh-Be-2.0/app/form/model"
	"MuXiFresh-Be-2.0/app/schedule/model"
	"MuXiFresh-Be-2.0/app/schedule/rpc/internal/config"
	userauthModel "MuXiFresh-Be-2.0/app/userauth/model"
)

type ServiceContext struct {
	Config          config.Config
	ScheduleClient  model.ScheduleModel
	EntryFormClient formmodel.EntryFormModel
	UserInfoClient  userauthModel.UserInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		ScheduleClient:  model.NewScheduleModel(c.MongoDBConf.URL, c.MongoDBConf.DB, "schedule"),
		EntryFormClient: formmodel.NewEntryFormModel(c.MongoDBConf.URL, c.MongoDBConf.DB, "entry_form"),
		UserInfoClient:  userauthModel.NewUserInfoModel(c.MongoDBConf.URL, c.MongoDBConf.DB, "userinfo"),
	}
}
