package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"go-zero-demo/datacenter/service/common/cmd/rpc/internal/config"
	"go-zero-demo/datacenter/service/common/model"
)

type ServiceContext struct {
	Config         config.Config
	AppConfigModel model.AppConfigModel
	BaseAppModel   model.BaseAppModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	apm := model.NewAppConfigModel(conn, c.CacheRedis)
	bam := model.NewBaseAppModel(conn, c.CacheRedis)
	return &ServiceContext{
		Config:         c,
		AppConfigModel: apm,
		BaseAppModel:   bam,
	}
}
