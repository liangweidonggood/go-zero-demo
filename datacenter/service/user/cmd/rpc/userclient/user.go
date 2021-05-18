// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

//go:generate mockgen -destination ./user_mock.go -package userclient -source $GOFILE

package userclient

import (
	"context"

	"go-zero-demo/datacenter/service/user/cmd/rpc/user"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	Request          = user.Request
	RegisterReq      = user.RegisterReq
	LoginReq         = user.LoginReq
	JwtToken         = user.JwtToken
	AppUserLoginResp = user.AppUserLoginResp
	AppConfigResp    = user.AppConfigResp
	Response         = user.Response
	UserReply        = user.UserReply
	UserReq          = user.UserReq
	AppConfigReq     = user.AppConfigReq
	AppUserResp      = user.AppUserResp

	User interface {
		Ping(ctx context.Context, in *Request) (*Response, error)
		Register(ctx context.Context, in *RegisterReq) (*UserReply, error)
		Login(ctx context.Context, in *LoginReq) (*UserReply, error)
		UserInfo(ctx context.Context, in *UserReq) (*UserReply, error)
		SnsLogin(ctx context.Context, in *AppConfigReq) (*AppUserResp, error)
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

func (m *defaultUser) Ping(ctx context.Context, in *Request) (*Response, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Ping(ctx, in)
}

func (m *defaultUser) Register(ctx context.Context, in *RegisterReq) (*UserReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Register(ctx, in)
}

func (m *defaultUser) Login(ctx context.Context, in *LoginReq) (*UserReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in)
}

func (m *defaultUser) UserInfo(ctx context.Context, in *UserReq) (*UserReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserInfo(ctx, in)
}

func (m *defaultUser) SnsLogin(ctx context.Context, in *AppConfigReq) (*AppUserResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.SnsLogin(ctx, in)
}
