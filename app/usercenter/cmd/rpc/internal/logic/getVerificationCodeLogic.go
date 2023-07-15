package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVerificationCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVerificationCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVerificationCodeLogic {
	return &GetVerificationCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVerificationCodeLogic) GetVerificationCode(in *pb.VerificationCodeRequest) (*pb.VerificationCodeResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.VerificationCodeResponse{}, nil
}
