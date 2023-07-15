package handler

import (
	"github.com/Wishforpeace/zero-tools/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"usercenter-api/internal/logic/user"
	"usercenter-api/internal/svc"
	"usercenter-api/internal/types"
)

func VerficationCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VerificationCodeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewVerficationCodeLogic(r.Context(), svcCtx)
		resp, err := l.VerficationCode(&req)
		response.Response(w, resp, err)

	}
}
