package svc

import (
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
	"go-zero-demo/book/service/search/cmd/api/internal/config"
	"go-zero-demo/book/service/search/cmd/api/internal/middleware"
	"go-zero-demo/book/service/user/cmd/rpc/userclient"
)

type ServiceContext struct {
	Config  config.Config
	Example rest.Middleware
	UserRpc userclient.User //不同项目组如何引用？
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Example: middleware.NewExampleMiddleware().Handle,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
