// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"topview-ttk/internal/app/ttk-user/rpc/internal/logic/userservice"
	"topview-ttk/internal/app/ttk-user/rpc/internal/svc"
	"topview-ttk/internal/app/ttk-user/rpc/user"
)

type UserServiceServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServiceServer
}

func NewUserServiceServer(svcCtx *svc.ServiceContext) *UserServiceServer {
	return &UserServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServiceServer) GetUserInfoByUid(ctx context.Context, in *user.GetUserInfoByUidRequest) (*user.GetUserInfoResponse, error) {
	l := userservicelogic.NewGetUserInfoByUidLogic(ctx, s.svcCtx)
	return l.GetUserInfoByUid(in)
}

func (s *UserServiceServer) GetUserInfoByUserName(ctx context.Context, in *user.GetUserInfoByUserNameRequest) (*user.GetUserInfoResponse, error) {
	l := userservicelogic.NewGetUserInfoByUserNameLogic(ctx, s.svcCtx)
	return l.GetUserInfoByUserName(in)
}
