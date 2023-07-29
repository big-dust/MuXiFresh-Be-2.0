package logic

import (
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/form/rpc/entryformclient"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/common/ctxData"
	"context"

	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/form/api/internal/svc"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/form/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFormLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateFormLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFormLogic {
	return &UpdateFormLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateFormLogic) UpdateForm(req *types.CreateReq) (resp *types.CreateResp, err error) {
	userId := ctxData.GetUserIdFromCtx(l.ctx)
	_, err = l.svcCtx.FormClient.UpdateForm(l.ctx, &entryformclient.CreateReq{
		FormId:        req.FormId,
		UserId:        userId,
		Avatar:        req.Avatar,
		Name:          req.Name,
		StuNumber:     req.StuNumber,
		School:        req.School,
		Major:         req.Major,
		Grade:         req.Grade,
		Gender:        req.Gender,
		Email:         req.Email,
		QQ:            req.QQ,
		Phone:         req.Phone,
		Group:         req.Group,
		Reason:        req.Reason,
		Knowledge:     req.Knowledge,
		SelfIntro:     req.SelfIntro,
		ExtraQuestion: req.ExtraQuestion,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateResp{
		Flag: true,
	}, nil
}
