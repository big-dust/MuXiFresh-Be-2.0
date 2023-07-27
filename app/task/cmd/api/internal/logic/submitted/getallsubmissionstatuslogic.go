package submitted

import (
	"MuXiFresh-Be-2.0/app/task/cmd/rpc/submission/pb"
	"context"
	"github.com/jinzhu/copier"

	"MuXiFresh-Be-2.0/app/task/cmd/api/internal/svc"
	"MuXiFresh-Be-2.0/app/task/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllSubmissionStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllSubmissionStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllSubmissionStatusLogic {
	return &GetAllSubmissionStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllSubmissionStatusLogic) GetAllSubmissionStatus(req *types.GetAllSubmissionStatusReq) (resp *types.GetAllSubmissionStatusResp, err error) {
	//管理员身份认证

	//get
	getAllStatusResp, err := l.svcCtx.SubmissionClient.GetAllSubmissionStatus(l.ctx, &pb.GetAllSubmissionStatusReq{
		AssignmentID: req.AssignmentID,
		Page:         int64(req.Page),
	})
	if err != nil {
		return nil, err
	}
	var completions []types.Completion
	copier.Copy(&completions, &getAllStatusResp.Completions)
	return &types.GetAllSubmissionStatusResp{
		Completions: completions,
	}, nil
}
