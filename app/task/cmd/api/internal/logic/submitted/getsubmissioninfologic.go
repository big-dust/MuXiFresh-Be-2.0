package submitted

import (
	"MuXiFresh-Be-2.0/app/task/cmd/api/internal/svc"
	"MuXiFresh-Be-2.0/app/task/cmd/api/internal/types"
	"MuXiFresh-Be-2.0/app/task/cmd/rpc/submission/submissionclient"
	"MuXiFresh-Be-2.0/app/user/cmd/rpc/user/userclient"
	"MuXiFresh-Be-2.0/common/ctxData"
	"MuXiFresh-Be-2.0/common/globalKey"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSubmissionInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSubmissionInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubmissionInfoLogic {
	return &GetSubmissionInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSubmissionInfoLogic) GetSubmissionInfo(req *types.GetSubmissionInfoReq) (resp *types.GetSubmissionInfoResp, err error) {
	var userId string
	if req.UserID != globalKey.Myself {
		//管理员身份认证
		getUserTypeResp, err := l.svcCtx.UserClient.GetUserType(l.ctx, &userclient.GetUserTypeReq{
			UserId: ctxData.GetUserIdFromCtx(l.ctx),
		})
		if err != nil {
			return nil, err
		}
		if getUserTypeResp.UserType != globalKey.Admin && getUserTypeResp.UserType != globalKey.SuperAdmin {
			return nil, errors.New("permission denied")
		}
		userId = req.UserID
	} else {
		userId = ctxData.GetUserIdFromCtx(l.ctx)
	}

	getSubmissionInfoResp, err := l.svcCtx.SubmissionClient.GetSubmissionInfo(l.ctx, &submissionclient.GetSubmissionInfoReq{
		AssignmentID: req.AssignmentID,
		UserId:       userId,
	})
	if err != nil {
		return nil, err
	}
	return &types.GetSubmissionInfoResp{
		SubmissionID: getSubmissionInfoResp.SubmissionID,
		Urls:         getSubmissionInfoResp.Urls,
	}, nil
}
