package logic

import (
	"MuXiFresh-Be-2.0/app/task/model"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"MuXiFresh-Be-2.0/app/task/cmd/rpc/submission/internal/svc"
	"MuXiFresh-Be-2.0/app/task/cmd/rpc/submission/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetSubmissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetSubmissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetSubmissionLogic {
	return &SetSubmissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetSubmissionLogic) SetSubmission(in *pb.SetSubmissionReq) (*pb.SetSubmissionResp, error) {

	userId, err := primitive.ObjectIDFromHex(in.UserId)
	if err != nil {
		return nil, err
	}
	
	assignmentID, err := primitive.ObjectIDFromHex(in.AssignmentID)
	if err != nil {
		return nil, err
	}

	toSetSubmission := &model.Submission{
		UserId:       userId,
		AssignmentID: assignmentID,
		Urls:         in.Urls,
		UpdateAt:     time.Now(),
		CreateAt:     time.Now(),
	}

	submission, err := l.svcCtx.SubmissionModel.FindByUserIdAndAssignmentID(l.ctx, in.UserId, in.AssignmentID)

	if err != nil {
		if err = l.svcCtx.SubmissionModel.Insert(l.ctx, toSetSubmission); err != nil {
			return nil, err
		}
	} else {
		toSetSubmission.ID = submission.ID
		if _, err = l.svcCtx.SubmissionModel.Update(l.ctx, toSetSubmission); err != nil {
			return nil, err
		}
	}
	return &pb.SetSubmissionResp{
		Flag: true,
	}, nil
}
