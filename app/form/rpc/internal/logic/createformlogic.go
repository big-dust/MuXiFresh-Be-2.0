package logic

import (
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/form/model"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/form/rpc/internal/svc"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/form/rpc/pb"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CreateFormLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateFormLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateFormLogic {
	return &CreateFormLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateFormLogic) CreateForm(in *pb.CreateReq) (*pb.CreateResp, error) {
	userId, err := primitive.ObjectIDFromHex(in.UserId)
	if err != nil {
		return nil, err
	}
	formID, err := l.svcCtx.FormClient.InsertReturnID(l.ctx, &model.EntryForm{
		UserId:        userId,
		Avatar:        in.Avatar,
		Name:          in.Name,
		StuNumber:     in.StuNumber,
		School:        in.School,
		Major:         in.Major,
		Grade:         in.Grade,
		Gender:        in.Gender,
		Email:         in.Email,
		QQ:            in.ExtraQuestion,
		Phone:         in.Phone,
		Group:         in.Group,
		Reason:        in.Reason,
		Knowledge:     in.Knowledge,
		SelfIntro:     in.SelfIntro,
		ExtraQuestion: in.ExtraQuestion,
		UpdateAt:      time.Now(),
		CreateAt:      time.Now(),
	})

	if err != nil {
		return nil, err
	}
	return &pb.CreateResp{
		FormID: fmt.Sprint(formID)[10:34],
	}, nil
}
