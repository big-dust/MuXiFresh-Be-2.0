package submitted

import (
	"MuXiFresh-Be-2.0/app/task/cmd/rpc/submission/submissionclient"
	"MuXiFresh-Be-2.0/common/ctxData"
	"context"

	"MuXiFresh-Be-2.0/app/task/cmd/api/internal/svc"
	"MuXiFresh-Be-2.0/app/task/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMySubmissionStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMySubmissionStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMySubmissionStatusLogic {
	return &GetMySubmissionStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMySubmissionStatusLogic) GetMySubmissionStatus(req *types.GetMySubmissionStatusReq) (resp *types.GetMySubmissionStatusResp, err error) {
	getMyStatusResp, err := l.svcCtx.SubmissionClient.GetMySubmissionStatus(l.ctx, &submissionclient.GetMySubmissionStatusReq{
		AssignmentID: req.AssignmentID,
		UserId:       ctxData.GetUserIdFromCtx(l.ctx),
	})
	if err != nil {
		return nil, err
	}
	return &types.GetMySubmissionStatusResp{MySubmissionStatus: getMyStatusResp.Status}, nil
}
