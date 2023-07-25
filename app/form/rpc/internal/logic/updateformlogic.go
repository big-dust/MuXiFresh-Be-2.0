package logic

import (
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/form/model"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/form/rpc/internal/svc"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/form/rpc/pb"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFormLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFormLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFormLogic {
	return &UpdateFormLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateFormLogic) UpdateForm(in *pb.CreateReq) (*pb.CreateResp, error) {
	u, err := primitive.ObjectIDFromHex(in.UserId)
	if err != nil {
		return nil, err
	}
	f, err := primitive.ObjectIDFromHex(in.FormId)
	if err != nil {
		return nil, err
	}
	//form := model.EntryForm{}
	//copier.Copy(&form,in)
	updateRet, err := l.svcCtx.FormClient.Update(l.ctx, &model.EntryForm{
		UserId:        u,
		ID:            f,
		Avatar:        in.Avatar,
		Name:          in.Name,
		StuNumber:     in.StuNumber,
		School:        in.School,
		Major:         in.Major,
		Grade:         in.Grade,
		Gender:        in.Gender,
		Email:         in.Email,
		QQ:            in.QQ,
		Phone:         in.Phone,
		Group:         in.Group,
		Reason:        in.Reason,
		Knowledge:     in.Knowledge,
		SelfIntro:     in.SelfIntro,
		ExtraQuestion: in.ExtraQuestion,
		UpdateAt:      time.Now(),
	})
	return &pb.CreateResp{
		FormID: fmt.Sprint(updateRet.UpsertedID),
	}, nil
}
