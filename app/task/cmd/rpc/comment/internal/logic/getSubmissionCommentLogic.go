package logic

import (
	formmodel "MuXiFresh-Be-2.0/app/form/model"
	"MuXiFresh-Be-2.0/app/task/cmd/rpc/comment/internal/svc"
	"MuXiFresh-Be-2.0/app/task/cmd/rpc/comment/pb"
	"MuXiFresh-Be-2.0/app/task/model"
	"MuXiFresh-Be-2.0/common/convert"
	"MuXiFresh-Be-2.0/common/xerr"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSubmissionCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSubmissionCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubmissionCommentLogic {
	return &GetSubmissionCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSubmissionCommentLogic) GetSubmissionComment(in *pb.GetSubmissionCommentReq) (*pb.GetSubmissionCommentResp, error) {

	comments, err := l.svcCtx.CommentModel.FindBySubmissionID(l.ctx, in.SubmissionID)
	if err != nil && err != model.ErrNotFound {
		switch err {
		case model.ErrInvalidObjectId:
			return nil, xerr.ErrExistInvalidId.Status()
		default:
			return nil, xerr.NewErrCode(xerr.DB_ERROR).Status()
		}
	}
	var cmtsWithUserInfo []*pb.Comment
	for _, comment := range comments {
		userInfo, err := l.svcCtx.UserInfoModel.FindOne(l.ctx, comment.UserId.String()[10:34])
		if err != nil {
			return nil, xerr.NewErrCode(xerr.DB_ERROR).Status()
		}
		entryForm, err := l.svcCtx.EntryFormModel.FindOne(l.ctx, userInfo.EntryFormID.String()[10:34])
		if err != nil && err != model.ErrNotFound {
			return nil, xerr.NewErrCode(xerr.DB_ERROR).Status()
		}
		if entryForm == nil {
			entryForm = new(formmodel.EntryForm)
		}
		group := entryForm.Grade + "级" + convert.GroupCvtChinese(entryForm.Group) + "成员"
		cmtsWithUserInfo = append(cmtsWithUserInfo, &pb.Comment{
			CommentID: comment.ID.String()[10:34],
			Avatar:    userInfo.Avatar,
			NickName:  userInfo.NickName,
			Group:     group,
			Content:   comment.Content,
		})
	}
	return &pb.GetSubmissionCommentResp{
		Comments: cmtsWithUserInfo,
	}, nil
}
