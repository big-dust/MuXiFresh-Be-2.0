package logic

import (
	"MuXiFresh-Be-2.0/app/test/rpc/testclient"
	"MuXiFresh-Be-2.0/app/user/cmd/rpc/user/userclient"
	"context"
	"strings"

	"MuXiFresh-Be-2.0/app/test/api/internal/svc"
	"MuXiFresh-Be-2.0/app/test/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckTestResultLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckTestResultLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckTestResultLogic {
	return &CheckTestResultLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckTestResultLogic) CheckTestResult(req *types.TestInfoReq) (resp *types.TestInfoResp, err error) {
	ut, err := l.svcCtx.UserClient.GetUserType(l.ctx, &userclient.GetUserTypeReq{
		UserId: req.UserID,
	})
	if err != nil {
		return nil, err
	}
	if (strings.Compare(ut.UserType, "admin") == 0) || (strings.Compare(ut.UserType, "super_admin") == 0) {
		r, err := l.svcCtx.TestClient.CheckTestResult(l.ctx, &testclient.TestInfoReq{
			UserID: req.UserID,
		})
		if err != nil {
			return nil, err
		}
		return &types.TestInfoResp{
			Name:        r.Name,
			Gender:      r.Gender,
			Major:       r.Major,
			Grade:       r.Grade,
			LeQunXing:   r.LeQunXing,
			YouHengXing: r.YouHengXing,
			XingFenXing: r.XingFenXing,
			CongHuiXing: r.CongHuiXing,
			JiaoJiXing:  r.JiaoJiXing,
			HuaiYiXing:  r.HuaiYiXing,
			WenDingXing: r.WenDingXing,
			Choice:      nil,
		}, nil
	}
	return nil, nil
}
