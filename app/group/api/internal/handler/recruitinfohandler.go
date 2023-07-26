package handler

import (
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/group/api/internal/logic"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/group/api/internal/svc"
	"greet/response"
	"net/http"
)

func RecruitInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewRecruitInfoLogic(r.Context(), svcCtx)
		err := l.RecruitInfo()
		response.Response(w, nil, err)

	}
}
