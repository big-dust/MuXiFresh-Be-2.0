package logic

import (
	formmodel "MuXiFresh-Be-2.0/app/form/model"
	"MuXiFresh-Be-2.0/app/task/cmd/rpc/submission/internal/svc"
	"MuXiFresh-Be-2.0/app/task/cmd/rpc/submission/pb"
	"MuXiFresh-Be-2.0/app/task/model"
	"MuXiFresh-Be-2.0/common/xerr"
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

	submissions, err := l.svcCtx.SubmissionModel.FindByAssignmentID(l.ctx, in.AssignmentID)

	if err != nil {
		switch err {
		case model.ErrInvalidObjectId:
			return nil, xerr.ErrExistInvalidId.Status()
		default:
			return nil, xerr.NewErrCode(xerr.DB_ERROR).Status()
		}
	}

	var completions []*pb.Completion
	for _, submission := range submissions {
		userId := submission.UserId.String()[10:34]

		entryForm, err := l.svcCtx.EntryFormModel.FindOneByUserId(l.ctx, userId)
		if err != nil && err != model.ErrNotFound {
			return nil, xerr.NewErrCode(xerr.DB_ERROR).Status()
		}
		if entryForm == nil {
			entryForm = new(formmodel.EntryForm)
		}
		userInfo, err := l.svcCtx.UserInfoModel.FindOne(l.ctx, userId)
		if err != nil {
			return nil, xerr.NewErrCode(xerr.DB_ERROR).Status()
		}

		completions = append(completions, &pb.Completion{
			UserId: entryForm.UserId.String()[10:34],
			Name:   userInfo.Name,
			Avatar: entryForm.Avatar,
			Email:  userInfo.Email,
			Grade:  entryForm.Grade,
			School: userInfo.School,
			Status: submission.Status,
		})
	}
	return &pb.GetAllSubmissionStatusResp{
		Completions: completions,
	}, nil
}
