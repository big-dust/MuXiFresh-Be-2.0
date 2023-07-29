package logic

import (
	"MuXiFresh-Be-2.0/app/test/rpc/internal/svc"
	"MuXiFresh-Be-2.0/app/test/rpc/pb"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckTestResultLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckTestResultLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckTestResultLogic {
	return &CheckTestResultLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckTestResultLogic) CheckTestResult(in *pb.TestInfoReq) (*pb.TestInfoResp, error) {
	u, err := l.svcCtx.UserInfoClient.FindOne(l.ctx, in.UserID)
	if err != nil {
		return nil, err
	}

	o, err := l.svcCtx.EntryFormClient.FindOne(l.ctx, u.EntryFormID.String()[10:34])
	if err != nil {
		return nil, err
	}

	return &pb.TestInfoResp{
		Name:        o.Name,
		Gender:      o.Gender,
		Major:       o.School + o.Major,
		Grade:       o.Grade,
		LeQunXing:   u.TestResult.LeQunXing,
		YouHengXing: u.TestResult.YouHengXing,
		XingFenXing: u.TestResult.XingFenXing,
		CongHuiXing: u.TestResult.CongHuiXing,
		JiaoJiXing:  u.TestResult.JiaoJiXing,
		HuaiYiXing:  u.TestResult.HuaiYiXing,
		WenDingXing: u.TestResult.WenDingXing,
		Choice:      u.TestChoice,
	}, nil
}
