package logic

import (
	"MuXiFresh-Be-2.0/app/userauth/cmd/rpc/accountCenter/pb"
	"MuXiFresh-Be-2.0/common/ctxData"
	"context"

	"MuXiFresh-Be-2.0/app/userauth/cmd/api/internal/svc"
	"MuXiFresh-Be-2.0/app/userauth/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PassChangeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPassChangeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PassChangeLogic {
	return &PassChangeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PassChangeLogic) PassChange(req *types.PassChangeReq) (resp *types.PassChangeResp, err error) {
	passChangeResp, err := l.svcCtx.ActCenterClient.PassChange(l.ctx, &pb.PassChangeReq{
		Email:    ctxData.GetEmailFromCtx(l.ctx),
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &types.PassChangeResp{Flag: passChangeResp.Flag}, nil
}
