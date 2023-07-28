package modify

import (
	"context"

	"MuXiFresh-Be-2.0/app/modify/cmd/api/internal/svc"
	"MuXiFresh-Be-2.0/app/modify/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifytheUseravatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifytheUseravatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifytheUseravatarLogic {
	return &ModifytheUseravatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifytheUseravatarLogic) ModifytheUseravatar(req *types.ModifytheUseravatarReq) (resp *types.ModifytheUseravatarResp, err error) {
		ModifytheUseravatarResp, err := l.svcCtx.SubmissionClient.ModifytheUseravatar(l.ctx, &pb.ModifytheUseravatarReq{
			Avatar:   req.Avatar,
		})
		if err != nil {
			return nil, err
		}
		return &types.ModifytheUseravatarResp{Flag: ModifytheUseravatarResp.Flag}, nil
}
