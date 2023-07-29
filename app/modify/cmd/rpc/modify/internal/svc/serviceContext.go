package svc

import "MuXiFresh-Be-2.0/app/modify/cmd/rpc/modify/internal/config"

type ServiceContext struct {
	Config config.Config
	UserInfoClient model.UserInfo
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UserInfoClient:model.NewsUserInfoModel(c.MongoConf.URL,c.MongoConf.DB,collection:"UserInfo")

	}
}
