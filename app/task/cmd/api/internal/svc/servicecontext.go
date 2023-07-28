package svc

import (
	"MuXiFresh-Be-2.0/app/task/cmd/api/internal/config"
	"MuXiFresh-Be-2.0/app/task/cmd/rpc/assignment/assignmentclient"
	"MuXiFresh-Be-2.0/app/task/cmd/rpc/comment/commentclient"
	"MuXiFresh-Be-2.0/app/task/cmd/rpc/submission/submissionclient"
	"MuXiFresh-Be-2.0/app/userauth/cmd/rpc/userinfo/userinfoclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	AssignmentClient assignmentclient.AssignmentClient
	SubmissionClient submissionclient.SubmissionClient
	CommentClient    commentclient.CommentClient
	UserInfoClient   userinfoclient.UserInfoClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		AssignmentClient: assignmentclient.NewAssignmentClient(zrpc.MustNewClient(c.AssignmentConf)),
		SubmissionClient: submissionclient.NewSubmissionClient(zrpc.MustNewClient(c.SubmissionConf)),
		CommentClient:    commentclient.NewCommentClient(zrpc.MustNewClient(c.CommentConf)),
		UserInfoClient:   userinfoclient.NewUserInfoClient(zrpc.MustNewClient(c.UserInfoConf)),
	}
}
