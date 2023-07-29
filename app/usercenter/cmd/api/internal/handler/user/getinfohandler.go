package handler

import (
	"MuXiFresh-Be-2.0/common/greet/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"usercenter-api/internal/logic/user"
	"usercenter-api/internal/svc"
	"usercenter-api/internal/types"
)

func GetInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewGetInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetInfo(&req)
		response.Response(w, resp, err)

	}
}
