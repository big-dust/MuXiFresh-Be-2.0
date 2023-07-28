package logic

import (
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/test/model"
	"context"
	"strings"
	"time"

	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/test/rpc/internal/svc"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/test/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserTestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserTestLogic {
	return &UserTestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserTestLogic) UserTest(in *pb.TestReq) (*pb.TestResp, error) {
	cr := make([]int, 85)

	type Result struct {
		A int64
		B int64
		C int64
		F int64
		G int64
		H int64
		L int64
	}
	result := Result{
		A: 0,
		B: 0,
		C: 0,
		F: 0,
		G: 0,
		H: 0,
		L: 0,
	}
	for index, value := range in.Choice {
		if (index == 2) || (index == 3) || (index == 7) || (index == 8) || (index == 13) || (index == 16) || (index == 17) || (index == 19) || (index == 22) || (index == 23) || (index == 39) || (index == 41) || (index == 44) || (index == 45) || (index >= 48 && index <= 53) || (index >= 56 && index <= 60) || (index == 62) || (index == 68) || (index == 72) || (index == 74) || (index == 75) || (index >= 78 && index <= 82) {
			if strings.Compare(value, "A") == 0 {
				cr[index] = 2
			}
		}
		if (index == 76) || (index == 77) {
			if strings.Compare(value, "A") == 0 {
				cr[index] = 1
			}
		}
		if (index >= 2 && index <= 29) || (index >= 31 && index <= 41) || (index >= 43 && index <= 53) || (index >= 55 && index <= 65) || (index >= 67 && index <= 75) || (index >= 78 && index <= 83) {
			if strings.Compare(value, "B") == 0 {
				cr[index] = 1
			}
		}
		if (index >= 4 && index <= 6) || (index == 9) || (index == 11) || (index == 12) || (index == 14) || (index == 15) || (index == 18) || (index >= 24 && index <= 29) || (index >= 32 && index <= 38) || (index == 40) || (index == 46) || (index == 47) || (index == 61) || (index == 63) || (index == 64) || (index == 67) || (index >= 69 && index <= 71) || (index == 73) || (index == 83) {
			if strings.Compare(value, "C") == 0 {
				cr[index] = 2
			}
		}
		if (index == 30) || (index == 42) || (index == 54) || (index == 66) {
			if strings.Compare(value, "C") == 0 {
				cr[index] = 1
			}
		}
	}
	a := []int{3, 10, 19, 20, 30, 42, 54, 65, 76}
	b := []int{11, 21, 22, 31, 32, 43, 44, 55, 56, 66, 67, 77, 78}
	c := []int{4, 5, 12, 13, 23, 33, 34, 45, 46, 57, 58, 68, 79}
	f := []int{6, 14, 24, 35, 36, 47, 48, 59, 60, 70, 71, 80, 81}
	g := []int{7, 15, 25, 37, 49, 61, 72, 73, 82, 83}
	h := []int{8, 16, 17, 26, 27, 38, 39, 50, 51, 62, 63, 74, 84}
	ll := []int{9, 18, 28, 29, 40, 41, 52, 53, 64, 75}

	for i1, v1 := range cr {
		flag := false

		for _, v2 := range a {
			if i1 == (v2 - 1) {
				result.A += int64(v1)
				flag = true
				continue
			}
		}
		if flag == true {
			continue
		}

		for _, v3 := range b {
			if i1 == (v3 - 1) {
				result.B += int64(v1)
				flag = true
				continue
			}
		}
		if flag == true {
			continue
		}

		for _, v4 := range c {
			if i1 == (v4 - 1) {
				result.C += int64(v1)
				flag = true
				continue
			}
		}
		if flag == true {
			continue
		}

		for _, v5 := range f {
			if i1 == (v5 - 1) {
				result.F += int64(v1)
				flag = true
				continue
			}
		}
		if flag == true {
			continue
		}

		for _, v6 := range g {
			if i1 == (v6 - 1) {
				result.G += int64(v1)
				flag = true
				continue
			}
		}
		if flag == true {
			continue
		}

		for _, v7 := range h {
			if i1 == (v7 - 1) {
				result.H += int64(v1)
				flag = true
				continue
			}
		}
		if flag == true {
			continue
		}

		for _, v8 := range ll {
			if i1 == (v8 - 1) {
				result.L += int64(v1)
				flag = true
				continue
			}
		}
		if flag == true {
			continue
		}

	}

	_, err := l.svcCtx.UserInfoClient.Update(l.ctx, &model.UserInfo{
		TestChoice: in.Choice,
		TestResult: struct {
			LeQunXing   int64
			YouHengXing int64
			XingFenXing int64
			CongHuiXing int64
			JiaoJiXing  int64
			HuaiYiXing  int64
			WenDingXing int64
		}{
			LeQunXing:   result.A,
			YouHengXing: result.G,
			XingFenXing: result.F,
			CongHuiXing: result.B,
			JiaoJiXing:  result.H,
			HuaiYiXing:  result.L,
			WenDingXing: result.C,
		},
		UpdateAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.TestResp{
		Flag: true,
	}, nil
}
