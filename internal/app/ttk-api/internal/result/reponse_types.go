package result

import (
	"topview-ttk/internal/pkg/ttkerr"
)

type ResponseSuccess struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type NullJson struct{}

type ResponseError struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

func Success(data interface{}) *ResponseSuccess {
	return &ResponseSuccess{ttkerr.OK, "OK", data}
}

func UnAuthorized(errMsg string) *ResponseError {
	return &ResponseError{ttkerr.Unauthorized, "用户鉴权失败:" + errMsg}
}

func Error(errCode uint32, errMsg string) *ResponseError {
	return &ResponseError{errCode, errMsg}
}
