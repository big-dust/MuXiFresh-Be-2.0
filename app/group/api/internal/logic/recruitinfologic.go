package logic

import (
	"context"

	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/group/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type RecruitInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecruitInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecruitInfoLogic {
	return &RecruitInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecruitInfoLogic) RecruitInfo() error {
	// todo: add your logic here and delete this line

	return nil
}
