package svc

import (
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/form/model"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/schedule/api/internal/config"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/schedule/rpc/scheduleclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	ScheduleClient scheduleclient.ScheduleClient
	UserInfoClient model.UserInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		ScheduleClient: scheduleclient.NewScheduleClient(zrpc.MustNewClient(c.ScheduleConf)),
		UserInfoClient: model.NewUserInfoModel(c.MongoDBConf.URL, c.MongoDBConf.DB, "user"),
	}
}
