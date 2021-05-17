package svc

import (
	"github.com/tal-tech/go-zero/rest"
	"go-zero-demo/book/service/search/cmd/api/internal/config"
	"go-zero-demo/book/service/search/cmd/api/internal/middleware"
)

type ServiceContext struct {
	Config  config.Config
	Example rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Example: middleware.NewExampleMiddleware().Handle,
	}
}
