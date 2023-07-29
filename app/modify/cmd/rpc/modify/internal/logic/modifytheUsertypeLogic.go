package logic

import (
	"context"

	"MuXiFresh-Be-2.0/app/modify/cmd/rpc/modify/internal/svc"
	"MuXiFresh-Be-2.0/app/modify/cmd/rpc/modify/modify"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifytheUsertypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifytheUsertypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifytheUsertypeLogic {
	return &ModifytheUsertypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModifytheUsertypeLogic) ModifytheUsertype(in *modify.ModifytheUsertypeReq) (*modify.ModifytheUsertypeResp, error) {
	_, err := l.svcCtx.Model.Insert(l.ctx,&model.userinfo{
		Email:    in.Email,
		UserType: in.UserType,
    })
    if err != nil {
        return nil, err
    }

    return &modifytheusertype.ModifytheUsertypeResp{
        Ok: true,
    }, nil
}
