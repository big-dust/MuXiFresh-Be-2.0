package logic

import (
	"context"

	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/form/rpc/internal/svc"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/form/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type JudgeUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJudgeUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JudgeUserLogic {
	return &JudgeUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JudgeUserLogic) JudgeUser(in *pb.ClickReq) (*pb.ClickResp, error) {
	// todo: add your logic here and delete this line

	return &pb.ClickResp{}, nil
}
