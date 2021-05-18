package config

import (
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Logx  logx.LogConf
	Mysql struct {
		DataSource string
	}
	CacheRedis cache.CacheConf
	Auth       struct {
		AccessSecret string
		AccessExpire int64
	}
}
