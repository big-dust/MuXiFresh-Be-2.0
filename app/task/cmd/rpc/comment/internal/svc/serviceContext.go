package svc

import (
	"MuXiFresh-Be-2.0/app/task/cmd/rpc/comment/internal/config"
	"MuXiFresh-Be-2.0/app/task/model"
	externalModel "MuXiFresh-Be-2.0/app/userauth/model"
)

type ServiceContext struct {
	Config          config.Config
	CommentModel    model.CommentModel
	UserInfoModel   externalModel.UserInfoModel
	SubmissionModel model.SubmissionModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		CommentModel:    model.NewCommentModel(c.MongoConf.URL, c.MongoConf.DB, "comment"),
		UserInfoModel:   externalModel.NewUserInfoModel(c.MongoConf.URL, c.MongoConf.DB, "userinfo"),
		SubmissionModel: model.NewSubmissionModel(c.MongoConf.URL, c.MongoConf.DB, "submission"),
	}
}
