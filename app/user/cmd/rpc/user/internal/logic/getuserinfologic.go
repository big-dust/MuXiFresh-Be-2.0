package logic

import (
	"MuXiFresh-Be-2.0/common/convert"
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

	userInfo, err := l.svcCtx.UserInfoModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}

	group := "尚未报名"

	if !userInfo.EntryFormID.IsZero() {
		group = "简历审阅中"
	}

	schedule, err := l.svcCtx.ScheduleModel.FindOne(l.ctx, userInfo.ScheduleID.String()[10:34])
	if err != nil {
		return nil, err
	}

	if schedule.AdmissionStatus == globalKey.Internship || in.UserType != globalKey.Freshman {
		entryForm, err := l.svcCtx.EntryFormModel.FindOneByUserId(l.ctx, in.UserId)
		if err != nil {
			return nil, err
		}
		identity := "实习生"
		if in.UserType != globalKey.Freshman {
			identity = "成员"
		}
		group = entryForm.Grade + "级" + convert.GroupCvtChinese(entryForm.Group) + identity
	}

	return &pb.GetUserInfoResp{
		Avatar:    userInfo.Avatar,
		NickName:  userInfo.NickName,
		Email:     userInfo.Email,
		Group:     group,
		StudentID: userInfo.StudentID,
	}, nil
}
