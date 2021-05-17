package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"go-zero-demo/mall/order/api/internal/config"
	"go-zero-demo/mall/user/rpc/userclient"
)

type ServiceContext struct {
	Config      config.Config
	MallUserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		MallUserRpc: userclient.NewUser(zrpc.MustNewClient(c.MallUserRpc)),
	}
}
