package logic

import (
	"MuXiFresh-Be-2.0/app/userauth/cmd/rpc/accountCenter/internal/svc"
	"MuXiFresh-Be-2.0/app/userauth/cmd/rpc/accountCenter/pb"
	"MuXiFresh-Be-2.0/app/userauth/model"
	"MuXiFresh-Be-2.0/common/tool"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindStudentIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBindStudentIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindStudentIDLogic {
	return &BindStudentIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BindStudentIDLogic) BindStudentID(in *pb.BindingStudentIDReq) (*pb.BindingStudentIDResp, error) {
	//一站式登录
	if !tool.CCNULogin(in.StudentID, in.Password) {
		return nil, errors.New("student_id or password is wrong")
	}
	//存userinfo
	uid, err := primitive.ObjectIDFromHex(in.UserId)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.UserInfoClient.Update(l.ctx, &model.UserInfo{
		ID:        uid,
		StudentID: in.StudentID,
		UpdateAt:  time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.BindingStudentIDResp{
		Flag: true,
	}, nil
}
