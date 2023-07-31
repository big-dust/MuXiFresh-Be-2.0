package svc

import (
	externalModel2 "MuXiFresh-Be-2.0/app/form/model"
	"MuXiFresh-Be-2.0/app/user/cmd/rpc/user/internal/config"
	externalModel "MuXiFresh-Be-2.0/app/userauth/model"
)

type ServiceContext struct {
	Config         config.Config
	UserInfoModel  externalModel.UserInfoModel
	EntryFormModel externalModel2.EntryFormModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		UserInfoModel:  externalModel.NewUserInfoModel(c.MongoConf.URL, c.MongoConf.DB, "userinfo"),
		EntryFormModel: externalModel2.NewEntryFormModel(c.MongoConf.URL, c.MongoConf.DB, "entry_form"),
	}
}
