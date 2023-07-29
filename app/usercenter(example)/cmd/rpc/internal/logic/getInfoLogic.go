package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfoLogic {
	return &GetInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetInfoLogic) GetInfo(in *pb.GetInfoRequest) (*pb.GetInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetInfoResponse{}, nil
}
