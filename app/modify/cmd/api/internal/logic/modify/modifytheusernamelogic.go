package modify

import (
	"context"

	"MuXiFresh-Be-2.0/app/modify/cmd/api/internal/svc"
	"MuXiFresh-Be-2.0/app/modify/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifytheUsernameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifytheUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifytheUsernameLogic {
	return &ModifytheUsernameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifytheUsernameLogic) ModifytheUsername(req *types.ModifytheUsernameReq) (resp *types.ModifytheUsernameResp, err error) {
	ModifytheUsernameResp, err := l.svcCtx.SubmissionClient.ModifytheUsername(l.ctx, &pb.ModifytheUsernameReq{
	    NickName:   req.NickName,
		})
		if err != nil {
			return nil, err
		}
		return &types.ModifytheUsernameResp{Flag: ModifytheUsernameResp.Flag}, nil
}
