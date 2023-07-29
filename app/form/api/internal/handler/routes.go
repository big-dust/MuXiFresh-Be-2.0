// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"MuXiFresh-Be-2.0/app/form/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/form/judge",
				Handler: JudgeUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/form",
				Handler: CreateFormHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/form",
				Handler: UpdateFormHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/form/view",
				Handler: CheckFormHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/api/v2"),
	)
}
