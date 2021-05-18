package svc

import (
	"github.com/tal-tech/go-zero/rest"
	"go-zero-demo/datacenter/service/user/cmd/api/internal/config"
	"go-zero-demo/datacenter/service/user/cmd/api/internal/middleware"
)

type ServiceContext struct {
	Config    config.Config
	Usercheck rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		Usercheck: middleware.NewUsercheckMiddleware().Handle,
	}
}
