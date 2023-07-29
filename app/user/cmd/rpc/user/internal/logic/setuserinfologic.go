package logic

import (
	"MuXiFresh-Be-2.0/app/user/cmd/rpc/user/internal/svc"
	"MuXiFresh-Be-2.0/app/user/cmd/rpc/user/pb"
	"MuXiFresh-Be-2.0/app/userauth/model"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserInfoLogic {
	return &SetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetUserInfoLogic) SetUserInfo(in *pb.SetUserInfoReq) (*pb.SetUserInfoResp, error) {

	uid, err := primitive.ObjectIDFromHex(in.UserId)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.UserInfoModel.Update(l.ctx, &model.UserInfo{
		ID:       uid,
		Avatar:   in.Avatar,
		NickName: in.NickName,
	})
	if err != nil {
		return nil, err
	}
	return &pb.SetUserInfoResp{
		Flag: true,
	}, nil
}
