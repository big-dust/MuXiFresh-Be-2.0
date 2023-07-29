package logic

import (
	"MuXiFresh-Be-2.0/app/user/cmd/rpc/user/userclient"
	"MuXiFresh-Be-2.0/common/ctxData"
	"context"

	"MuXiFresh-Be-2.0/app/user/cmd/api/internal/svc"
	"MuXiFresh-Be-2.0/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserInfoLogic {
	return &SetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetUserInfoLogic) SetUserInfo(req *types.SetUserInfoReq) (resp *types.SetUserInfoResp, err error) {

	setUserInfoResp, err := l.svcCtx.UserClient.SetUserInfo(l.ctx, &userclient.SetUserInfoReq{
		UserId:   ctxData.GetUserIdFromCtx(l.ctx),
		Avatar:   req.Avatar,
		NickName: req.NickName,
	})
	if err != nil {
		return nil, err
	}
	return &types.SetUserInfoResp{
		Flag: setUserInfoResp.Flag,
	}, nil
}
