package logic

import (
	"MuXiFresh-Be-2.0/common/ctxData"
	"context"

	"MuXiFresh-Be-2.0/app/schedule/rpc/internal/svc"
	"MuXiFresh-Be-2.0/app/schedule/rpc/pb"

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
	userid := ctxData.GetUserIdFromCtx(l.ctx)
	u, err2 := l.svcCtx.UserInfoClient.FindOne(l.ctx, userid)
	if err2 != nil {
		return nil, err2
	}
	f, err2 := l.svcCtx.EntryFormClient.FindOne(l.ctx, u.EntryFormID.String()[10:34])
	if err2 != nil {
		return nil, err2
	}
	s, err := l.svcCtx.ScheduleClient.FindOne(l.ctx, in.ScheduleID)
	if err != nil {
		return nil, err
	}
	return &pb.CheckResp{
		Name:            f.Name,
		School:          f.School,
		Major:           f.Major,
		Group:           f.Group,
		EntryFormStatus: s.EntryFormStatus,
		AdmissionStatus: s.AdmissionStatus,
	}, nil
}
