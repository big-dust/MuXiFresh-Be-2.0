package logic

import (
	"context"

	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/schedule/rpc/internal/svc"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/schedule/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLogic) Check(in *pb.CheckReq) (*pb.CheckResp, error) {
	s, err := l.svcCtx.ScheduleClient.FindOne(l.ctx, in.ScheduleID)
	if err != nil {
		return nil, err
	}
	return &pb.CheckResp{
		Name:            s.Name,
		School:          s.School,
		Major:           s.Major,
		Group:           s.Group,
		EntryFormStatus: s.EntryFormStatus,
		AdmissionStatus: s.AdmissionStatus,
	}, nil
}
