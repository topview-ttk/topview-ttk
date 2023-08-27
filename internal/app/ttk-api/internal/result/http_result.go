package result

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"
	ttkerr2 "topview-ttk/internal/pkg/ttkerr"
)

func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		errCode := ttkerr2.ServerCommonError
		errMsg := ttkerr2.MapErrMsg(errCode)
		causeErr := errors.Cause(err)

		if e, ok := causeErr.(*ttkerr2.CodeError); ok {
			errCode = e.GetErrCode()
			errMsg = e.GetErrMsg()
		} else {
			if gs, ok := status.FromError(causeErr); ok {
				gc := uint32(gs.Code())
				if ttkerr2.IsCodeErr(gc) {
					errCode = gc
					errMsg = gs.Message()
				}
			}
		}
		logx.WithContext(r.Context()).Errorf("[API-ERR]: %+v ", err)
		httpx.WriteJson(w, http.StatusBadRequest, Error(errCode, errMsg))
	}
}

// ParamErrorResult .4
func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error) {
	errMsg := fmt.Sprintf("%s ,%s", ttkerr2.MapErrMsg(ttkerr2.RequestParamError), err.Error())
	httpx.WriteJson(w, http.StatusBadRequest, Error(ttkerr2.RequestParamError, errMsg))
}
