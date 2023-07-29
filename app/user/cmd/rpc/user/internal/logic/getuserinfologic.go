package logic

import (
	"MuXiFresh-Be-2.0/common/globalKey"
	"context"

	"MuXiFresh-Be-2.0/app/user/cmd/rpc/user/internal/svc"
	"MuXiFresh-Be-2.0/app/user/cmd/rpc/user/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {

	group := "the resume is under review"

	if in.UserType != globalKey.Freshman {
		entryForm, err := l.svcCtx.EntryFormModel.FindOneByUserId(l.ctx, in.UserId)
		if err != nil {
			return nil, err
		}
		group = entryForm.Group
	}

	userInfo, err := l.svcCtx.UserInfoModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserInfoResp{
		Avatar:    userInfo.Avatar,
		NickName:  userInfo.NickName,
		Email:     userInfo.Email,
		Group:     group,
		StudentID: userInfo.StudentID,
	}, nil
}
