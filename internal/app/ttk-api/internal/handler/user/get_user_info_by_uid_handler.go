package user

import (
	"net/http"
	"topview-ttk/internal/app/ttk-api/internal/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"topview-ttk/internal/app/ttk-api/internal/logic/user"
	"topview-ttk/internal/app/ttk-api/internal/svc"
	"topview-ttk/internal/app/ttk-api/internal/types"
)

func GetUserInfoByUidHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserInfoByUidRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewGetUserInfoByUidLogic(r.Context(), svcCtx)
		resp, err := l.GetUserInfoByUid(&req)
		result.HttpResult(r, w, resp, err)
	}
}
