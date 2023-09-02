package logic

import (
	"MuXiFresh-Be-2.0/app/task/model"
	"MuXiFresh-Be-2.0/common/xerr"
	"context"

	"MuXiFresh-Be-2.0/app/task/cmd/rpc/submission/internal/svc"
	"MuXiFresh-Be-2.0/app/task/cmd/rpc/submission/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSubmissionInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSubmissionInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubmissionInfoLogic {
	return &GetSubmissionInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSubmissionInfoLogic) GetSubmissionInfo(in *pb.GetSubmissionInfoReq) (*pb.GetSubmissionInfoResp, error) {

	submission, err := l.svcCtx.SubmissionModel.FindByUserIdAndAssignmentID(l.ctx, in.UserId, in.AssignmentID)
	if err != nil {
		switch err {
		case model.ErrNotFound:
			return nil, xerr.ErrNotFind
		case model.ErrInvalidObjectId:
			return nil, xerr.ErrExistInvalidId.Status()
		default:
			return nil, xerr.NewErrCode(xerr.DB_ERROR).Status()
		}
	}
	return &pb.GetSubmissionInfoResp{
		SubmissionID: submission.ID.String()[10:34],
		Urls:         submission.Urls,
	}, nil
}
