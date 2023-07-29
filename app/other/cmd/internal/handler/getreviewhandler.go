package handler

import (
	"MuXiFresh-Be-2.0/app/other/cmd/internal/logic"
	"MuXiFresh-Be-2.0/app/other/cmd/internal/svc"
	"MuXiFresh-Be-2.0/app/other/cmd/internal/types"
	"greet/response"
	"net/http"
)

func GetReviewHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetReviewReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetReviewLogic(r.Context(), svcCtx)
		resp, err := l.GetReview(&req)
		response.Response(w, resp, err)

	}
}
