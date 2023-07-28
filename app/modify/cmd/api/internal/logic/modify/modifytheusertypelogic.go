package modify

import (
	"context"

	"MuXiFresh-Be-2.0/app/modify/cmd/api/internal/svc"
	"MuXiFresh-Be-2.0/app/modify/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifytheUsertypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifytheUsertypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifytheUsertypeLogic {
	return &ModifytheUsertypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifytheUsertypeLogic) ModifytheUsertype(req *types.ModifytheUsertypeReq) (resp *types.ModifytheUsertypeResp, err error) {
	ModifytheUsertypeResp, err:= l.svcCtx.SubmissionClient.ModifytheUsertype(l.ctx, &pb.ModifytheUsertypeReq{
		Email:    req.Email,
		UserType: req.UserType,
		})
		if err != nil {
			return nil, err
		}
		return &types.ModifytheUsertypeResp{Flag: ModifytheUseravatarResp.Flag}, nil
}

