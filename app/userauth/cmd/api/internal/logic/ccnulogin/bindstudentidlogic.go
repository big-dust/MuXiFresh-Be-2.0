package ccnulogin

import (
	"MuXiFresh-Be-2.0/app/userauth/cmd/rpc/accountCenter/pb"
	"MuXiFresh-Be-2.0/common/ctxData"
	"context"

	"MuXiFresh-Be-2.0/app/userauth/cmd/api/internal/svc"
	"MuXiFresh-Be-2.0/app/userauth/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindStudentIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBindStudentIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindStudentIDLogic {
	return &BindStudentIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BindStudentIDLogic) BindStudentID(req *types.BindStudentIDReq) (resp *types.BindStudentIDResp, err error) {

	bindResp, err := l.svcCtx.ActCenterClient.BindStudentID(l.ctx, &pb.BindingStudentIDReq{
		UserId:    ctxData.GetUserIdFromCtx(l.ctx),
		StudentID: req.StdudentID,
		Password:  req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &types.BindStudentIDResp{
		Flag: bindResp.Flag,
	}, nil
}
