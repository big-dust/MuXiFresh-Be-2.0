package handler

import (
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/group/api/internal/logic"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/group/api/internal/svc"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/group/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"greet/response"
	"net/http"
)

func GetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetLogic(r.Context(), svcCtx)
		resp, err := l.Get(&req)
		response.Response(w, resp, err)

	}
}
