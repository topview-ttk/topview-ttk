package userA

import (
	"net/http"
	"topview-ttk/internal/app/ttk-api/internal/result"

	"topview-ttk/internal/app/ttk-api/internal/logic/userA"
	"topview-ttk/internal/app/ttk-api/internal/svc"
)

func GetUserInfoSelfHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := userA.NewGetUserInfoSelfLogic(r.Context(), svcCtx)
		resp, err := l.GetUserInfoSelf()
		result.HttpResult(r, w, resp, err)
	}
}
