// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"ZeZeIM/apps/user/rpc/internal/logic"
	"ZeZeIM/apps/user/rpc/internal/svc"
	"ZeZeIM/apps/user/rpc/pb/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Ping(ctx context.Context, in *user.PingReq) (*user.PingResp, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}

func (s *UserServer) Login(ctx context.Context, in *user.LoginReq) (*user.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) Register(ctx context.Context, in *user.RegisterReq) (*user.RegisterResp, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserServer) FindUser(ctx context.Context, in *user.FindUserReq) (*user.FindUserResp, error) {
	l := logic.NewFindUserLogic(ctx, s.svcCtx)
	return l.FindUser(in)
}

func (s *UserServer) AllUser(ctx context.Context, in *user.AllUserReq) (*user.AllUserResp, error) {
	l := logic.NewAllUserLogic(ctx, s.svcCtx)
	return l.AllUser(in)
}
