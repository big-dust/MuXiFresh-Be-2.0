package logic

import (
	"MuXiFresh-Be-2.0/app/test/api/internal/svc"
	"MuXiFresh-Be-2.0/app/test/api/internal/types"
	"MuXiFresh-Be-2.0/app/test/rpc/testclient"
	"MuXiFresh-Be-2.0/common/ctxData"
	"context"

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
	var uid string
	if req.UserID == "myself" {
		uid = ctxData.GetUserIdFromCtx(l.ctx)
	} else {
		uid = req.UserID
	}
	_, err = l.svcCtx.TestClient.CheckTestResult(l.ctx, &testclient.TestInfoReq{
		Token:  req.Authorization,
		UserID: uid,
	})
	if err != nil {
		return nil, err
	}

	userInfo, err := l.svcCtx.UserInfoClient.FindOne(l.ctx, uid)
	if err != nil {
		return nil, err
	}

	formid := userInfo.EntryFormID

	form, err := l.svcCtx.FormClient.FindOne(l.ctx, formid.String()[10:34])

	typesChoice := make([]types.ChoiceItem, len(userInfo.TestChoice))
	for i, item := range userInfo.TestChoice {
		typesChoice[i] = types.ChoiceItem(item)
	}

	return &types.TestInfoResp{
		Name:        userInfo.Name,
		Gender:      form.Gender,
		Major:       form.Major,
		Grade:       form.Grade,
		LeQunXing:   userInfo.TestResult.LeQunXing,
		YouHengXing: userInfo.TestResult.YouHengXing,
		XingFenXing: userInfo.TestResult.XingFenXing,
		CongHuiXing: userInfo.TestResult.CongHuiXing,
		JiaoJiXing:  userInfo.TestResult.JiaoJiXing,
		HuaiYiXing:  userInfo.TestResult.HuaiYiXing,
		WenDingXing: userInfo.TestResult.WenDingXing,
		Choice:      typesChoice,
	}, nil
}
