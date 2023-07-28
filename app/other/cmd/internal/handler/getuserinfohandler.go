package handler

import (
	"MuXiFresh-Be-2.0/app/other/cmd/internal/logic"
	"MuXiFresh-Be-2.0/app/other/cmd/internal/svc"
	"MuXiFresh-Be-2.0/app/other/cmd/internal/types"
	"greet/response"
	"net/http"
)

func GetUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetUserInfo(&req)
		response.Response(w, resp, err)

	}
}
