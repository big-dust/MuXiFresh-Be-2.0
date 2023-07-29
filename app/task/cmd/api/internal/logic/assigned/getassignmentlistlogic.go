package assigned

import (
	"MuXiFresh-Be-2.0/app/task/cmd/rpc/assignment/pb"
	"context"
	"github.com/jinzhu/copier"

	"MuXiFresh-Be-2.0/app/task/cmd/api/internal/svc"
	"MuXiFresh-Be-2.0/app/task/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAssignmentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAssignmentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAssignmentListLogic {
	return &GetAssignmentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAssignmentListLogic) GetAssignmentList(req *types.GetAssignmentListReq) (resp *types.GetAssignmentListResp, err error) {
	getListResp, err := l.svcCtx.AssignmentClient.GetAssignmentList(l.ctx, &pb.GetAssignmentListReq{
		Group: req.Group,
	})
	if err != nil {
		return nil, err
	}
	var titles []types.Title
	copier.Copy(&titles, &getListResp.Titles)
	return &types.GetAssignmentListResp{
		Titles: titles,
	}, nil
}
