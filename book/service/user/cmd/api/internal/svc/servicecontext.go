package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"go-zero-demo/book/service/user/cmd/api/internal/config"
	"go-zero-demo/book/service/user/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlconn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlconn, c.CacheRedis),
	}
}
