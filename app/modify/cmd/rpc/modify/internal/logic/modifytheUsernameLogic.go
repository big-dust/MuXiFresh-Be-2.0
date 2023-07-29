package logic

import (
	"context"

	"MuXiFresh-Be-2.0/app/modify/cmd/rpc/modify/internal/svc"
	"MuXiFresh-Be-2.0/app/modify/cmd/rpc/modify/modify"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifytheUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifytheUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifytheUsernameLogic {
	return &ModifytheUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModifytheUsernameLogic) ModifytheUsername(in *modify.ModifytheUsernameReq) (*modify.ModifytheUsernameResp, error) {
	_, err := l.svcCtx.Model.Insert(l.ctx,&model.userinfo{
		NickName:   in.NickName,
    })
    if err != nil {
        return nil, err
    }

    return &modifytheusername.ModifytheUsernameResp{
        Ok: true,
    }, nil
}
