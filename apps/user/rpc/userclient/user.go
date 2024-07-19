// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userclient

import (
	"context"

	"ZeZeIM/apps/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AllUserReq   = user.AllUserReq
	AllUserResp  = user.AllUserResp
	FindUserReq  = user.FindUserReq
	FindUserResp = user.FindUserResp
	LoginReq     = user.LoginReq
	LoginResp    = user.LoginResp
	PingReq      = user.PingReq
	PingResp     = user.PingResp
	RegisterReq  = user.RegisterReq
	RegisterResp = user.RegisterResp
	UserEntity   = user.UserEntity

	User interface {
		Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingResp, error)
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
		FindUser(ctx context.Context, in *FindUserReq, opts ...grpc.CallOption) (*FindUserResp, error)
		AllUser(ctx context.Context, in *AllUserReq, opts ...grpc.CallOption) (*AllUserResp, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

func (m *defaultUser) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUser) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUser) FindUser(ctx context.Context, in *FindUserReq, opts ...grpc.CallOption) (*FindUserResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.FindUser(ctx, in, opts...)
}

func (m *defaultUser) AllUser(ctx context.Context, in *AllUserReq, opts ...grpc.CallOption) (*AllUserResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.AllUser(ctx, in, opts...)
}
