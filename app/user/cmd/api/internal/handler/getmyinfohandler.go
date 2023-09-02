package handler

import (
	"net/http"

	"MuXiFresh-Be-2.0/common/result"

	"MuXiFresh-Be-2.0/app/user/cmd/api/internal/logic"
	"MuXiFresh-Be-2.0/app/user/cmd/api/internal/svc"
	"MuXiFresh-Be-2.0/app/user/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetMyInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetMyInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewGetMyInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetMyInfo(&req)
		result.HttpResult(r, w, resp, err)
	}
}
