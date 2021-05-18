package svc

import (
	"github.com/tal-tech/go-zero/rest"
	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/config"
	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/middleware"
)

type ServiceContext struct {
	Config     config.Config
	Usercheck  rest.Middleware
	Admincheck rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		Usercheck:  middleware.NewUsercheckMiddleware().Handle,
		Admincheck: middleware.NewAdmincheckMiddleware().Handle,
	}
}
