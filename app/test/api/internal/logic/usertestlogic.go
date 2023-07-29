package logic

import (
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/test/rpc/testclient"
	"context"

	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/test/api/internal/svc"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/test/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserTestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserTestLogic {
	return &UserTestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserTestLogic) UserTest(req *types.TestReq) (resp *types.TestResp, err error) {
	_, err = l.svcCtx.TestClient.UserTest(l.ctx, &testclient.TestReq{
		Choice: req.Choice,
		Token:  req.Authorization,
	})
	if err != nil {
		return nil, err
	}
	return &types.TestResp{
		Flag: true,
	}, nil
}
