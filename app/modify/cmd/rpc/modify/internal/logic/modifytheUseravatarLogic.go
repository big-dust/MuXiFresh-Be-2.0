package logic

import (
	"context"

	"MuXiFresh-Be-2.0/app/modify/cmd/rpc/modify/internal/svc"
	"MuXiFresh-Be-2.0/app/modify/cmd/rpc/modify/modify"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifytheUseravatarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifytheUseravatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifytheUseravatarLogic {
	return &ModifytheUseravatarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModifytheUseravatarLogic) ModifytheUseravatar(in *modify.ModifytheUseravatarReq) (*modify.ModifytheUseravatarResp, error) {
	_, err := l.svcCtx.Model.Insert(l.ctx,&model.userinfo{
		Avatar:   in.Avatar,
    })
    if err != nil {
        return nil, err
    }

    return &modifytheuseravatar.ModifytheUseravatarResp{
        Ok: true,
    }, nil
}
