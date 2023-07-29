package adjust

import (
	"net/http"

	"MuXiFresh-Be-2.0/app/modify/cmd/api/internal/logic/adjust"
	"MuXiFresh-Be-2.0/app/modify/cmd/api/internal/svc"
	"MuXiFresh-Be-2.0/app/modify/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AdjustAdmissionstatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AdjustAdmissionstatusReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := adjust.NewAdjustAdmissionstatusLogic(r.Context(), svcCtx)
		resp, err := l.AdjustAdmissionstatus(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
