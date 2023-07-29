package svc

import (
	"MuXiFresh-Be-2.0/app/userauth/cmd/rpc/userinfo/internal/config"
	"MuXiFresh-Be-2.0/app/userauth/model"
)

type ServiceContext struct {
	Config        config.Config
	UserInfoModel model.UserInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserInfoModel: model.NewUserInfoModel(c.MongoConf.URL, c.MongoConf.DB, "userinfo"),
	}
}
