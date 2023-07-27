package logic

import (
	"MuXiFresh-Be-2.0/app/task/cmd/rpc/submission/internal/svc"
	"MuXiFresh-Be-2.0/app/task/cmd/rpc/submission/pb"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllSubmissionStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllSubmissionStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllSubmissionStatusLogic {
	return &GetAllSubmissionStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllSubmissionStatusLogic) GetAllSubmissionStatus(in *pb.GetAllSubmissionStatusReq) (*pb.GetAllSubmissionStatusResp, error) {

	submissions, err := l.svcCtx.SubmissionModel.FindByAssignmentID(l.ctx, in.AssignmentID, l.svcCtx.Config.Limit, (in.Page-1)*l.svcCtx.Config.Limit)
	if err != nil {
		return nil, err
	}
	var completions []*pb.Completion
	for _, submission := range submissions {
		form, err := l.svcCtx.FormModel.FindOneByUserId(l.ctx, submission.UserId.String()[10:34])
		if err != nil {
			return nil, err
		}
		completions = append(completions, &pb.Completion{
			Name:    form.Name,
			Grade:   form.Grade,
			College: form.College,
			Status:  submission.Status,
		})
	}
	return &pb.GetAllSubmissionStatusResp{
		Completions: completions,
	}, nil
}
