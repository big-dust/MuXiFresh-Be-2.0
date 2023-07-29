package logic

import (
	"MuXiFresh-Be-2.0/common/ctxData"
	"context"

	"MuXiFresh-Be-2.0/app/test/rpc/internal/svc"
	"MuXiFresh-Be-2.0/app/test/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type JudgeUserTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJudgeUserTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JudgeUserTypeLogic {
	return &JudgeUserTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JudgeUserTypeLogic) JudgeUserType(in *pb.UserTypeReq) (*pb.UserTypeResp, error) {
	uid := ctxData.GetUserIdFromCtx(l.ctx)
	u, err := l.svcCtx.UserInfoClient.FindOne(l.ctx, uid)
	if err != nil {
		return nil, err
	}

	return &pb.UserTypeResp{
		Type: u.UserType,
	}, nil
}
