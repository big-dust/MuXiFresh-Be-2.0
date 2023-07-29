package handler

import (
	"MuXiFresh-Be-2.0/app/user/cmd/api/internal/logic"
	"MuXiFresh-Be-2.0/app/user/cmd/api/internal/svc"
	"MuXiFresh-Be-2.0/app/user/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"greet/response"
	"net/http"
)

func GetAdminListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAdminListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetAdminListLogic(r.Context(), svcCtx)
		resp, err := l.GetAdminList(&req)
		response.Response(w, resp, err)

	}
}
