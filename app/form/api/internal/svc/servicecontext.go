package svc

import (
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/form/api/internal/config"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/form/model"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/form/rpc/entryformclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config              config.Config
	FormClient          entryformclient.EntryFormClient
	UserInfoModelClient model.UserInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		FormClient:          entryformclient.NewEntryFormClient(zrpc.MustNewClient(c.FormConf)),
		UserInfoModelClient: model.NewUserInfoModel(c.MongoDBConf.URL, c.MongoDBConf.DB, "userinfo"),
	}
}
