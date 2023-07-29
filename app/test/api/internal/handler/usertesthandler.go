package handler

import (
	"MuXiFresh-Be-2.0/app/test/api/internal/logic"
	"MuXiFresh-Be-2.0/app/test/api/internal/svc"
	"MuXiFresh-Be-2.0/app/test/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"greet/response"
	"net/http"
)

func UserTestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TestReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserTestLogic(r.Context(), svcCtx)
		resp, err := l.UserTest(&req)
		response.Response(w, resp, err)

	}
}
