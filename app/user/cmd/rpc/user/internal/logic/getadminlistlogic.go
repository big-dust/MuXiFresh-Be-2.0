package logic

import (
	"MuXiFresh-Be-2.0/app/user/cmd/rpc/user/internal/svc"
	"MuXiFresh-Be-2.0/app/user/cmd/rpc/user/pb"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAdminListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminListLogic {
	return &GetAdminListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAdminListLogic) GetAdminList(in *pb.GetAdminListReq) (*pb.GetAdminListResp, error) {
	userInfos, err := l.svcCtx.UserInfoModel.FindByUserType(l.ctx, in.UserType, l.svcCtx.Config.Limit, (in.Page-1)*l.svcCtx.Config.Limit)
	if err != nil {
		return nil, err
	}
	var list []*pb.One
	for _, userInfo := range userInfos {
		list = append(list, &pb.One{
			UserId:   userInfo.ID.String()[10:34],
			Nickname: userInfo.NickName,
			Avatar:   userInfo.Avatar,
		})
	}
	return &pb.GetAdminListResp{
		List: list,
	}, nil
}
