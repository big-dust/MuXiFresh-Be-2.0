package logic

import (
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/group/api/internal/svc"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/group/api/internal/types"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/group/rpc/getclient"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get(req *types.GetReq) (resp *types.GetResp, err error) {
	g, err := l.svcCtx.IntroClient.GetGroupIntro(l.ctx, &getclient.GetReq{
		Name: req.Name,
	})

	if err != nil {
		return nil, err
	}

	return &types.GetResp{
		Intro: g.Intro,
	}, nil
}
