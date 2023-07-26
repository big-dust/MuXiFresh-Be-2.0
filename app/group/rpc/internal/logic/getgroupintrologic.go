package logic

import (
	"context"
	"strings"

	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/group/rpc/internal/svc"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/group/rpc/pb"

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

func (l *GetGroupIntroLogic) GetGroupIntro(in *pb.GroupIntroReq) (*pb.GroupIntroResp, error) {
	if strings.Compare(in.GroupName, "产品组") == 0 {
		return &pb.GroupIntroResp{
			Intro: "这是一段产品组的介绍",
		}, nil
	} else if strings.Compare(in.GroupName, "安卓组") == 0 {
		return &pb.GroupIntroResp{
			Intro: "这是一段安卓组的介绍",
		}, nil
	} else if strings.Compare(in.GroupName, "前端组") == 0 {
		return &pb.GroupIntroResp{
			Intro: "这是一段前端组的介绍",
		}, nil
	} else if strings.Compare(in.GroupName, "后端组") == 0 {
		return &pb.GroupIntroResp{
			Intro: "这是一段后端组的介绍",
		}, nil
	} else if strings.Compare(in.GroupName, "设计组") == 0 {
		return &pb.GroupIntroResp{
			Intro: "这是一段设计组的介绍",
		}, nil
	}
	return nil, nil
}
