package logic

import (
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/schedule/rpc/scheduleclient"
	"context"

	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/schedule/api/internal/svc"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/schedule/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateScheduleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateScheduleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateScheduleLogic {
	return &CreateScheduleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateScheduleLogic) CreateSchedule(req *types.CreateReq) (resp *types.CreateResp, err error) {
	_, err = l.svcCtx.ScheduleClient.Create(l.ctx, &scheduleclient.CreateReq{
		Token: req.Authorization,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateResp{
		Flag: true,
	}, nil
}
