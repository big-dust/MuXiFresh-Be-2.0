package logic

import (
	"MuXiFresh-Be-2.0/common/globalKey"
	"context"

	"MuXiFresh-Be-2.0/app/task/cmd/rpc/submission/internal/svc"
	"MuXiFresh-Be-2.0/app/task/cmd/rpc/submission/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMySubmissionStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMySubmissionStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMySubmissionStatusLogic {
	return &GetMySubmissionStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMySubmissionStatusLogic) GetMySubmissionStatus(in *pb.GetMySubmissionStatusReq) (*pb.GetMySubmissionStatusResp, error) {

	submission, err := l.svcCtx.SubmissionModel.FindByUserIdAndAssignmentID(l.ctx, in.UserId, in.AssignmentID)
	status := globalKey.Submitted
	if err != nil {
		status = globalKey.NotSubmitted
	} else {
		if submission.Status == globalKey.Reviewed {
			status = globalKey.Reviewed
		}
	}
	return &pb.GetMySubmissionStatusResp{
		Status: status,
	}, nil
}
