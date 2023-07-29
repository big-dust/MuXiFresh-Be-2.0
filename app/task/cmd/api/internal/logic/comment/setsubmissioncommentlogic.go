package comment

import (
	"MuXiFresh-Be-2.0/app/task/cmd/rpc/comment/commentclient"
	"MuXiFresh-Be-2.0/common/ctxData"
	"context"

	"MuXiFresh-Be-2.0/app/task/cmd/api/internal/svc"
	"MuXiFresh-Be-2.0/app/task/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetSubmissionCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetSubmissionCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetSubmissionCommentLogic {
	return &SetSubmissionCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetSubmissionCommentLogic) SetSubmissionComment(req *types.SetSubmissionCommentReq) (resp *types.SetSubmissionCommentResp, err error) {
	setCommentResp, err := l.svcCtx.CommentClient.SetSubmissionComment(l.ctx, &commentclient.SetSubmissionCommentReq{
		UserId:       ctxData.GetUserIdFromCtx(l.ctx),
		SubmissionID: req.SubmissionID,
		Content:      req.Content,
	})
	if err != nil {
		return nil, err
	}
	return &types.SetSubmissionCommentResp{
		Flag: setCommentResp.Flag,
	}, nil
}
