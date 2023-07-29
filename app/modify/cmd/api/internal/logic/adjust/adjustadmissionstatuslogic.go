package adjust

import (
	"context"

	"MuXiFresh-Be-2.0/app/modify/cmd/api/internal/svc"
	"MuXiFresh-Be-2.0/app/modify/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdjustAdmissionstatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdjustAdmissionstatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdjustAdmissionstatusLogic {
	return &AdjustAdmissionstatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdjustAdmissionstatusLogic) AdjustAdmissionstatus(req *types.AdjustAdmissionstatusReq) (resp *types.AdjustAdmissionstatusResp, err error) {
	AdjustAdmissionstatusResp, err := l.svcCtx.SubmissionClient.AdjustAdmissionstatus(l.ctx, &pb.AdjustAdmissionstatusReq{
		AdmissionStatus:   req.AdmissionStatus,
	})
	if err != nil {
		return nil, err
	}
	return &types.AdjustAdmissionstatusResp{Flag: ModifytheUseravatarResp.Flag}, nil
}
