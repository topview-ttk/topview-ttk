package result

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"
	"topview-ttk/internal/pkg/common/ttkerr"
)

func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		errCode := ttkerr.ServerCommonError
		errMsg := ttkerr.MapErrMsg(errCode)
		causeErr := errors.Cause(err)

		if e, ok := causeErr.(*ttkerr.CodeError); ok {
			errCode = e.GetErrCode()
			errMsg = e.GetErrMsg()
		} else {
			if gs, ok := status.FromError(causeErr); ok {
				gc := uint32(gs.Code())
				if ttkerr.IsCodeErr(gc) {
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
	errMsg := fmt.Sprintf("%s ,%s", ttkerr.MapErrMsg(ttkerr.RequestParamError), err.Error())
	httpx.WriteJson(w, http.StatusBadRequest, Error(ttkerr.RequestParamError, errMsg))
}
