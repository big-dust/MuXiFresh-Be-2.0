package logic

import (
	"context"

	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/group/api/internal/svc"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/group/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupIntroLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupIntroLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupIntroLogic {
	return &GroupIntroLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupIntroLogic) GroupIntro(req *types.GroupIntroReq) (resp *types.GroupIntroResp, err error) {
	// todo: add your logic here and delete this line

	return
}
