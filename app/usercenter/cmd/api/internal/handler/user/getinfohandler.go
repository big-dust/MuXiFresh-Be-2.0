package handler

import (
	"api/internal/logic/user"
	"api/internal/svc"
	"api/internal/types"
	"github.com/Wishforpeace/zero-tools/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
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
