package logic

import (
	"MuXiFresh-Be-2.0/app/other/cmd/internal/svc"
	"MuXiFresh-Be-2.0/app/other/cmd/internal/types"
	pb2 "MuXiFresh-Be-2.0/app/userauth/cmd/rpc/userinfo/pb"
	"MuXiFresh-Be-2.0/common/ctxData"
	"MuXiFresh-Be-2.0/common/globalKey"
	"context"
	"errors"

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

	getUserTypeResp, err := l.svcCtx.UserInfoClient.GetUserType(l.ctx, &pb2.GetUserTypeReq{
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

	group := "the resume is under review"

	if getUserTypeResp.UserType != globalKey.Freshman {
		entryForm, err := l.svcCtx.EntryFormModel.FindOneByUserId(l.ctx, userId)
		if err != nil {
			return nil, err
		}
		group = entryForm.Group
	}

	userInfo, err := l.svcCtx.UserInfoModel.FindOne(l.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &types.GetUserInfoResp{
		Avatar:    userInfo.Avatar,
		NickName:  userInfo.NickName,
		Email:     userInfo.Email,
		Group:     group,
		StudentID: userInfo.StudentID,
	}, nil
}
