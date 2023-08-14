package handler

import (
	logic "MuXiFresh-Be-2.0/app/userauth/cmd/api/internal/logic/ccnulogin"
	"MuXiFresh-Be-2.0/app/userauth/cmd/api/internal/svc"
	"MuXiFresh-Be-2.0/app/userauth/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"greet/response"
	"net/http"
)

func SetStudentIDHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetStudentIDReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSetStudentIDLogic(r.Context(), svcCtx)
		resp, err := l.SetStudentID(&req)
		response.Response(w, resp, err)

	}
}
