package sso

import (
	"net/http"
	"topview-ttk/internal/app/ttk-api/internal/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"topview-ttk/internal/app/ttk-api/internal/logic/sso"
	"topview-ttk/internal/app/ttk-api/internal/svc"
	"topview-ttk/internal/app/ttk-api/internal/types"
)

func PhoneRegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PhoneRegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := sso.NewPhoneRegisterLogic(r.Context(), svcCtx)
		resp, err := l.PhoneRegister(&req)
		result.HttpResult(r, w, resp, err)
	}
}
