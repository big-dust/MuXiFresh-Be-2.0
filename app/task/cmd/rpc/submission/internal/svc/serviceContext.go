package svc

import (
	"MuXiFresh-Be-2.0/app/task/cmd/rpc/submission/internal/config"
	"MuXiFresh-Be-2.0/app/task/model"
)

type ServiceContext struct {
	Config          config.Config
	SubmissionModel model.SubmissionModel
	EntryFormModel  model.EntryFormModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		SubmissionModel: model.NewSubmissionModel(c.MongoConf.URL, c.MongoConf.DB, "submission"),
		EntryFormModel:  model.NewEntryFormModel(c.MongoConf.URL, c.MongoConf.DB, "entry_form"),
	}
}
