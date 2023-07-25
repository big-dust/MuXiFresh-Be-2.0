package logic

import (
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/schedule/model"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/common/ctxData"
	"context"
	"time"

	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/schedule/rpc/internal/svc"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/schedule/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *pb.CreateReq) (*pb.CreateResp, error) {
	userid := ctxData.GetUserIdFromCtx(l.ctx)
	u, err := l.svcCtx.UserInfoClient.FindOne(l.ctx, userid)
	if err != nil {
		return nil, err
	}
	f, err := l.svcCtx.EntryFormClient.FindOne(l.ctx, u.EntryFormID.String()[10:34])
	if err != nil {
		return nil, err
	}
	l.svcCtx.ScheduleClient.Insert(l.ctx, &model.Schedule{
		UserID:          u.ID,
		Name:            f.Name,
		School:          f.School,
		Major:           f.Major,
		Group:           f.Group,
		EntryFormStatus: "已提交",
		AdmissionStatus: "已报名",
		UpdateAt:        time.Now(),
		CreateAt:        time.Now(),
	})
	return &pb.CreateResp{
		Flag: true,
	}, nil
}
