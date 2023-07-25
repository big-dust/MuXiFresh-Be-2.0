package handler

import (
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/form/api/internal/logic"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/form/api/internal/svc"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/form/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"greet/response"
	"net/http"
)

func CheckFormHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CheckReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCheckFormLogic(r.Context(), svcCtx)
		resp, err := l.CheckForm(&req)
		response.Response(w, resp, err)

	}
}
