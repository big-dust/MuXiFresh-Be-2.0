package logic

import (
	"MuXiFresh-Be-2.0/app/userauth/model"
	"context"

	"MuXiFresh-Be-2.0/app/userauth/cmd/rpc/accountCenter/internal/svc"
	"MuXiFresh-Be-2.0/app/userauth/cmd/rpc/accountCenter/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PassChangeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPassChangeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PassChangeLogic {
	return &PassChangeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PassChangeLogic) PassChange(in *pb.PassChangeReq) (*pb.PassChangeResp, error) {
	_, err := l.svcCtx.UserAuthClient.UpdateByEm(l.ctx, &model.UserAuth{
		Email:    in.Email,
		Password: in.Password,
	})
	if err != nil {
		return nil, err
	}
	return &pb.PassChangeResp{
		Flag: true,
	}, nil
}
