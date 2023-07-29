package comment

import (
	"MuXiFresh-Be-2.0/app/task/cmd/api/internal/svc"
	"MuXiFresh-Be-2.0/app/task/cmd/api/internal/types"
	"MuXiFresh-Be-2.0/app/task/cmd/rpc/comment/pb"
	"context"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSubmissionCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSubmissionCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubmissionCommentLogic {
	return &GetSubmissionCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSubmissionCommentLogic) GetSubmissionComment(req *types.GetSubmissionCommentReq) (resp *types.GetSubmissionCommentResp, err error) {
	getCommentResp, err := l.svcCtx.CommentClient.GetSubmissionComment(l.ctx, &pb.GetSubmissionCommentReq{
		SubmissionID: req.SubmissionID,
	})
	if err != nil {
		return nil, err
	}
	var comments []types.Comment
	copier.Copy(&comments, &getCommentResp.Comments)
	return &types.GetSubmissionCommentResp{
		Comments: comments,
	}, nil
}
