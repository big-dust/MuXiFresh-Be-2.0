package svc

import (
	"MuXiFresh-Be-2.0/app/task/cmd/rpc/submission/internal/config"
	"MuXiFresh-Be-2.0/app/task/model"
)

type ServiceContext struct {
	Config          config.Config
	SubmissionModel model.SubmissionModel
	FormModel       model.FormModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		SubmissionModel: model.NewSubmissionModel(c.MongoConf.URL, c.MongoConf.DB, "submission"),
		FormModel:       model.NewFormModel(c.MongoConf.URL, c.MongoConf.DB, "form"),
	}
}
