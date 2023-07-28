package logic

import (
	"context"

	"MuXiFresh-Be-2.0/app/modify/cmd/rpc/adjust/adjust"
	"MuXiFresh-Be-2.0/app/modify/cmd/rpc/adjust/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdjustAdmissionstatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdjustAdmissionstatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdjustAdmissionstatusLogic {
	return &AdjustAdmissionstatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdjustAdmissionstatusLogic) AdjustAdmissionstatus(in *adjust.AdjustAdmissionstatusReq) (*adjust.AdjustAdmissionstatusResp, error) {
	_, err := l.svcCtx.Model.Insert(l.ctx, &model.schedule{
		AdmissionStatus: in.AdmissionStatus,
	})
	if err != nil {
		return nil, err
	}

	return &modifytheusertype.AdjustAdmissionstatusResp{
		Ok: true,
	}, nil
}
