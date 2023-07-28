package logic

import (
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/schedule/model"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/common/ctxData"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	uid, err := primitive.ObjectIDFromHex(userid)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.ScheduleClient.Insert(l.ctx, &model.Schedule{
		UserID:          uid,
		EntryFormStatus: "未提交",
		AdmissionStatus: "未报名",
		UpdateAt:        time.Now(),
		CreateAt:        time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateResp{
		Flag: true,
	}, nil
}
