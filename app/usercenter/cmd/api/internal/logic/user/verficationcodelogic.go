package user

import (
	"context"

	"usercenter-api/internal/svc"
	"usercenter-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerficationCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerficationCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerficationCodeLogic {
	return &VerficationCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerficationCodeLogic) VerficationCode(req *types.VerificationCodeRequest) (resp *types.VerificationCodeResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
