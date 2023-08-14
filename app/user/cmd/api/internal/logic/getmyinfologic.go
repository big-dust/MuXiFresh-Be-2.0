package logic

import (
	"MuXiFresh-Be-2.0/app/user/cmd/rpc/user/pb"
	"MuXiFresh-Be-2.0/common/ctxData"
	"context"

	"MuXiFresh-Be-2.0/app/user/cmd/api/internal/svc"
	"MuXiFresh-Be-2.0/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMyInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyInfoLogic {
	return &GetMyInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMyInfoLogic) GetMyInfo(req *types.GetMyInfoReq) (resp *types.GetMyInfoResp, err error) {

	userId := ctxData.GetUserIdFromCtx(l.ctx)

	userInfo, err := l.svcCtx.UserClient.GetUserInfo(l.ctx, &pb.GetUserInfoReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	return &types.GetMyInfoResp{
		Avatar:    userInfo.Avatar,
		NickName:  userInfo.NickName,
		Name:      userInfo.Name,
		School:    userInfo.School,
		Group:     userInfo.Group,
		Email:     userInfo.Email,
		StudentID: userInfo.StudentID,
		QQ:        userInfo.QQ,
	}, nil
}
