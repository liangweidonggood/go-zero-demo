package svc

import (
	"context"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/core/syncx"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
	"go-zero-demo/datacenter/common/shared"
	"go-zero-demo/datacenter/service/common/cmd/rpc/commonclient"
	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/config"
	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/middleware"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

type ServiceContext struct {
	Config           config.Config
	Usercheck        rest.Middleware
	Admincheck       rest.Middleware
	Cache            cache.Cache
	RedisConn        *redis.Redis
	GreetMiddleware1 rest.Middleware
	GreetMiddleware2 rest.Middleware
	CommonRpc        commonclient.Common //公共
}

func timeInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	stime := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		return err
	}

	fmt.Printf("调用 %s 方法 耗时: %v\n", method, time.Now().Sub(stime))
	return nil
}

func NewServiceContext(c config.Config) *ServiceContext {
	//缓存
	ca := cache.New(c.CacheRedis, syncx.NewSharedCalls(), cache.NewStat("dc"), shared.ErrNotFound)
	rcon := redis.NewRedis(c.CacheRedis[0].Host, c.CacheRedis[0].Type, c.CacheRedis[0].Pass)
	//rpc
	cr := commonclient.NewCommon(zrpc.MustNewClient(c.CommonRpc, zrpc.WithUnaryClientInterceptor(timeInterceptor)))

	return &ServiceContext{
		Config:           c,
		Usercheck:        middleware.NewUsercheckMiddleware().Handle,
		Admincheck:       middleware.NewAdmincheckMiddleware().Handle,
		Cache:            ca,
		RedisConn:        rcon,
		GreetMiddleware1: greetMiddleware1,
		GreetMiddleware2: greetMiddleware2,
		CommonRpc:        cr,
	}
}
func greetMiddleware1(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("greetMiddleware1 request ... ")
		next(w, r)
		logx.Info("greetMiddleware1 reponse ... ")
	}
}

func greetMiddleware2(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("greetMiddleware2 request ... ")
		next(w, r)
		logx.Info("greetMiddleware2 reponse ... ")
	}
}
