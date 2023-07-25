package logic

import (
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/schedule/rpc/scheduleclient"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/common/ctxData"
	"context"

	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/schedule/api/internal/svc"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/schedule/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckLogic) Check(req *types.CheckReq) (resp *types.CheckResp, err error) {
	userid := ctxData.GetUserIdFromCtx(l.ctx)
	u, err := l.svcCtx.UserInfoClient.FindOne(l.ctx, userid)
	if err != nil {
		return nil, err
	}
	req.ScheduleID = u.ScheduleID.String()[10:34]

	c, err := l.svcCtx.ScheduleClient.Check(l.ctx, &scheduleclient.CheckReq{
		ScheduleID: req.ScheduleID,
	})

	return &types.CheckResp{
		Name:            c.Name,
		School:          c.School,
		Major:           c.Major,
		Group:           c.Group,
		EntryFormStatus: c.EntryFormStatus,
		AdmissionStatus: c.AdmissionStatus,
	}, nil
}
