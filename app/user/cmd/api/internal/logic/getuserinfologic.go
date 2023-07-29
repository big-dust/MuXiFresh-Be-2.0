package logic

import (
	"MuXiFresh-Be-2.0/app/user/cmd/rpc/user/pb"
	"MuXiFresh-Be-2.0/common/ctxData"
	"MuXiFresh-Be-2.0/common/globalKey"
	"context"
	"errors"

	"MuXiFresh-Be-2.0/app/user/cmd/api/internal/svc"
	"MuXiFresh-Be-2.0/app/user/cmd/api/internal/types"

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

	getUserTypeResp, err := l.svcCtx.UserClient.GetUserType(l.ctx, &pb.GetUserTypeReq{
		UserId: ctxData.GetUserIdFromCtx(l.ctx),
	})

	userId := ctxData.GetUserIdFromCtx(l.ctx)
	if userId != globalKey.Empty {
		//管理员认证
		if err != nil {
			return nil, err
		}
		if getUserTypeResp.UserType != globalKey.Admin && getUserTypeResp.UserType != globalKey.SuperAdmin {
			return nil, errors.New("permission denied")
		}
		userId = req.UserId
	}

	userInfo, err := l.svcCtx.UserClient.GetUserInfo(l.ctx, &pb.GetUserInfoReq{
		UserId:   userId,
		UserType: getUserTypeResp.UserType,
	})

	return &types.GetUserInfoResp{
		Avatar:    userInfo.Avatar,
		NickName:  userInfo.NickName,
		Email:     userInfo.Email,
		Group:     userInfo.Group,
		StudentID: userInfo.StudentID,
	}, nil
}
