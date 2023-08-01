package logic

import (
	"MuXiFresh-Be-2.0/app/user/cmd/api/internal/svc"
	"MuXiFresh-Be-2.0/app/user/cmd/api/internal/types"
	"MuXiFresh-Be-2.0/app/user/cmd/rpc/user/pb"
	"MuXiFresh-Be-2.0/common/ctxData"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoReq) (resp *types.GetUserInfoResp, err error) {

	userId := ctxData.GetUserIdFromCtx(l.ctx)

	getUserTypeResp, err := l.svcCtx.UserClient.GetUserType(l.ctx, &pb.GetUserTypeReq{
		UserId: userId,
	})

	userInfo, err := l.svcCtx.UserClient.GetUserInfo(l.ctx, &pb.GetUserInfoReq{
		UserId:   userId,
		UserType: getUserTypeResp.UserType,
	})
	if err != nil {
		return nil, err
	}
	return &types.GetUserInfoResp{
		Avatar:    userInfo.Avatar,
		NickName:  userInfo.NickName,
		Email:     userInfo.Email,
		Group:     userInfo.Group,
		StudentID: userInfo.StudentID,
	}, nil
}
