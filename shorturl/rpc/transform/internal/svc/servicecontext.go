package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"go-zero-demo/shorturl/rpc/transform/internal/config"
	"go-zero-demo/shorturl/rpc/transform/model"
)

type ServiceContext struct {
	Config config.Config
	Model model.ShorturlModel   // 手动代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model: model.NewShorturlModel(sqlx.NewMysql(c.DataSource), c.Cache), // 手动代码
	}
}
