package svc

import (
	externalModel2 "MuXiFresh-Be-2.0/app/form/model"
	"MuXiFresh-Be-2.0/app/review/cmd/api/internal/config"
	externalModel3 "MuXiFresh-Be-2.0/app/schedule/model"
	"MuXiFresh-Be-2.0/app/user/cmd/rpc/user/userclient"
	externalModel1 "MuXiFresh-Be-2.0/app/userauth/model"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	EntryFormModel externalModel2.EntryFormModel
	UserClient     userclient.UserClient
	ScheduleClient externalModel3.ScheduleModel
	UserInfoModel  externalModel1.UserInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		EntryFormModel: externalModel2.NewEntryFormModel(c.MongoConf.URL, c.MongoConf.DB, "entry_form"),
		UserClient:     userclient.NewUserClient(zrpc.MustNewClient(c.UserConf)),
		ScheduleClient: externalModel3.NewScheduleModel(c.MongoConf.URL, c.MongoConf.DB, "schedule"),
		UserInfoModel:  externalModel1.NewUserInfoModel(c.MongoConf.URL, c.MongoConf.DB, "userinfo"),
	}
}
