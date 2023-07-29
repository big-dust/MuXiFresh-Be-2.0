package svc

import (
	"MuXiFresh-Be-2.0/app/task/cmd/api/internal/config"
	"MuXiFresh-Be-2.0/app/task/cmd/rpc/assignment/assignmentclient"
	"MuXiFresh-Be-2.0/app/task/cmd/rpc/comment/commentclient"
	"MuXiFresh-Be-2.0/app/task/cmd/rpc/submission/submissionclient"
	"MuXiFresh-Be-2.0/app/user/cmd/rpc/user/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	AssignmentClient assignmentclient.AssignmentClient
	SubmissionClient submissionclient.SubmissionClient
	CommentClient    commentclient.CommentClient
	UserClient       userclient.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		AssignmentClient: assignmentclient.NewAssignmentClient(zrpc.MustNewClient(c.AssignmentConf)),
		SubmissionClient: submissionclient.NewSubmissionClient(zrpc.MustNewClient(c.SubmissionConf)),
		CommentClient:    commentclient.NewCommentClient(zrpc.MustNewClient(c.CommentConf)),
		UserClient:       userclient.NewUserClient(zrpc.MustNewClient(c.UserConf)),
	}
}
