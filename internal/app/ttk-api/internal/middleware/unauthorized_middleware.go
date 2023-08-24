package middleware

import (
	"github.com/zeromicro/go-zero/rest/handler"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"topview-ttk/internal/app/ttk-api/internal/result"
)

type UnAuthorizedMiddleware struct {
}

func NewUnAuthorizedMiddleware() *UnAuthorizedMiddleware {
	return &UnAuthorizedMiddleware{}
}

func (a *UnAuthorizedMiddleware) Callback() handler.UnauthorizedCallback {
	return func(w http.ResponseWriter, r *http.Request, err error) {
		resp := result.UnAuthorized(err.Error())
		httpx.WriteJson(w, http.StatusUnauthorized, resp)
	}
}
