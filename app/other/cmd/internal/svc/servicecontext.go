package svc

import (
	"MuXiFresh-Be-2.0/app/other/cmd/internal/config"
	"MuXiFresh-Be-2.0/app/other/model"
	externalModel2 "MuXiFresh-Be-2.0/app/task/model"
	"MuXiFresh-Be-2.0/app/userauth/cmd/rpc/userinfo/userinfoclient"
	externalModel "MuXiFresh-Be-2.0/app/userauth/model"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	EntryFormModel externalModel2.EntryFormModel
	UserInfoModel  externalModel.UserInfoModel
	UserInfoClient userinfoclient.UserInfoClient
	ScheduleClient model.ScheduleModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		EntryFormModel: externalModel2.NewEntryFormModel(c.MongoConf.URL, c.MongoConf.DB, "entry_form"),
		UserInfoModel:  externalModel.NewUserInfoModel(c.MongoConf.URL, c.MongoConf.DB, "userinfo"),
		UserInfoClient: userinfoclient.NewUserInfoClient(zrpc.MustNewClient(c.UserInfoConf)),
		ScheduleClient: model.NewScheduleModel(c.MongoConf.URL, c.MongoConf.DB, "schedule"),
	}
}
