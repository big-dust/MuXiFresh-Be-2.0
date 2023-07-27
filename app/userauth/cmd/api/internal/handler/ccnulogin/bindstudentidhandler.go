package handler

import (
	"MuXiFresh-Be-2.0/app/userauth/cmd/api/internal/svc"
	"MuXiFresh-Be-2.0/app/userauth/cmd/api/internal/types"
	"MuXiFresh-Be-2.0/common/greet/response"
	"net/http"
)

func BindStudentIDHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BindStudentIDReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewBindStudentIDLogic(r.Context(), svcCtx)
		resp, err := l.BindStudentID(&req)
		response.Response(w, resp, err)

	}
}
