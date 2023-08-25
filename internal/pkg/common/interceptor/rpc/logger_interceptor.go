package rpc

import (
	"context"
	"topview-ttk/internal/pkg/common/ttkerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req)
	if err != nil {
		causeErr := errors.Cause(err)
		if e, ok := causeErr.(*ttkerr.CodeError); ok {
			logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", err)
			//convert grpc err
			err = status.Error(codes.Code(e.GetErrCode()), e.GetErrMsg())
		} else {
			logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", err)
		}
	}
	return resp, err
}
