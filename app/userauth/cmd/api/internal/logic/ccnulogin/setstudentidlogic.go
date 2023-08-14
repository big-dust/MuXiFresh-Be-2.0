package ccnulogin

import (
	"MuXiFresh-Be-2.0/app/userauth/cmd/rpc/accountCenter/accountcenterclient"
	"MuXiFresh-Be-2.0/common/ctxData"
	"context"

	"MuXiFresh-Be-2.0/app/userauth/cmd/api/internal/svc"
	"MuXiFresh-Be-2.0/app/userauth/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetStudentIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetStudentIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetStudentIDLogic {
	return &SetStudentIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetStudentIDLogic) SetStudentID(req *types.SetStudentIDReq) (resp *types.SetStudentIDResp, err error) {

	SetResp, err := l.svcCtx.AccountCenterClient.SetStudentID(l.ctx, &accountcenterclient.SetStudentIDReq{
		UserId:    ctxData.GetUserIdFromCtx(l.ctx),
		StudentID: req.StdudentID,
		Password:  req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &types.SetStudentIDResp{
		Flag: SetResp.Flag,
	}, nil
}
