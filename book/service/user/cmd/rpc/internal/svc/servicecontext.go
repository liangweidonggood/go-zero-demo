package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"go-zero-demo/book/service/user/cmd/rpc/internal/config"
	"go-zero-demo/book/service/user/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
	}
}
