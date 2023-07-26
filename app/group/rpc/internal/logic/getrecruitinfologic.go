package logic

import (
	"context"

	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/group/rpc/internal/svc"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/group/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRecruitInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRecruitInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRecruitInfoLogic {
	return &GetRecruitInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRecruitInfoLogic) GetRecruitInfo(in *pb.RecruitInfoReq) (*pb.RecruitInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.RecruitInfoResp{}, nil
}
