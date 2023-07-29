package handler

import (
	"MuXiFresh-Be-2.0/app/form/api/internal/logic"
	"MuXiFresh-Be-2.0/app/form/api/internal/svc"
	"MuXiFresh-Be-2.0/app/form/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"greet/response"
	"net/http"
)

func CreateFormHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCreateFormLogic(r.Context(), svcCtx)
		resp, err := l.CreateForm(&req)
		response.Response(w, resp, err)

	}
}
