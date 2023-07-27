package assigned

import (
	"MuXiFresh-Be-2.0/app/task/cmd/rpc/assignment/pb"
	"context"

	"MuXiFresh-Be-2.0/app/task/cmd/api/internal/svc"
	"MuXiFresh-Be-2.0/app/task/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetAssignmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetAssignmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetAssignmentLogic {
	return &SetAssignmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetAssignmentLogic) SetAssignment(req *types.SetAssignmentReq) (resp *types.SetAssignmentResp, err error) {

	//管理员身份认证

	//布置
	setAssignmentResp, err := l.svcCtx.AssignmentClient.SetAssignment(l.ctx, &pb.SetAssignmentReq{
		AssignmentID: req.AssignmentID,
		Group:        req.Group,
		TitleText:    req.TitleText,
		Content:      req.Content,
		Urls:         req.Urls,
	})
	if err != nil {
		return nil, err
	}
	return &types.SetAssignmentResp{Flag: setAssignmentResp.Flag}, nil
}
