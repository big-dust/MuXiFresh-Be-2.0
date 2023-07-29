package svc

import (
	"MuXiFresh-Be-2.0/app/form/api/internal/config"
	"MuXiFresh-Be-2.0/app/form/rpc/entryformclient"
	externalModel "MuXiFresh-Be-2.0/app/userauth/model"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config              config.Config
	FormClient          entryformclient.EntryFormClient
	UserInfoModelClient externalModel.UserInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		FormClient:          entryformclient.NewEntryFormClient(zrpc.MustNewClient(c.FormConf)),
		UserInfoModelClient: externalModel.NewUserInfoModel(c.MongoDBConf.URL, c.MongoDBConf.DB, "userinfo"),
	}
}
