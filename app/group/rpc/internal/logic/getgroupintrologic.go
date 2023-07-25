package logic

import (
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/group/rpc/internal/svc"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/group/rpc/pb"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupIntroLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupIntroLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupIntroLogic {
	return &GetGroupIntroLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupIntroLogic) GetGroupIntro(in *pb.GetReq) (*pb.GetResp, error) {
	group, err := l.svcCtx.GroupClient.FindByName(l.ctx, in.Name)

	if err != nil {
		return nil, err
	}
	return &pb.GetResp{
		Intro: group.Intro,
	}, nil
}
