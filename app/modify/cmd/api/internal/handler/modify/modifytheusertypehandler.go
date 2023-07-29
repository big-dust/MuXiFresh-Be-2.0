package modify

import (
	"net/http"

	"MuXiFresh-Be-2.0/app/modify/cmd/api/internal/logic/modify"
	"MuXiFresh-Be-2.0/app/modify/cmd/api/internal/svc"
	"MuXiFresh-Be-2.0/app/modify/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ModifytheUsertypeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ModifytheUsertypeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := modify.NewModifytheUsertypeLogic(r.Context(), svcCtx)
		resp, err := l.ModifytheUsertype(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
