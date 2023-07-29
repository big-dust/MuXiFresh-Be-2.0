package submitted

import (
	"MuXiFresh-Be-2.0/app/task/cmd/rpc/submission/submissionclient"
	"MuXiFresh-Be-2.0/common/ctxData"
	"context"

	"MuXiFresh-Be-2.0/app/task/cmd/api/internal/svc"
	"MuXiFresh-Be-2.0/app/task/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetSubmissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetSubmissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetSubmissionLogic {
	return &SetSubmissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetSubmissionLogic) SetSubmission(req *types.SetSubmissionReq) (resp *types.SetSubmissionResp, err error) {
	setSubmissionResp, err := l.svcCtx.SubmissionClient.SetSubmission(l.ctx, &submissionclient.SetSubmissionReq{
		AssignmentID: req.AssignmentID,
		UserId:       ctxData.GetUserIdFromCtx(l.ctx),
		Urls:         req.Urls,
	})
	if err != nil {
		return nil, err
	}
	return &types.SetSubmissionResp{Flag: setSubmissionResp.Flag}, nil
}
