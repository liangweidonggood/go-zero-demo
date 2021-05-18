package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"net/http"

	"go-zero-demo/book/service/search/cmd/api/internal/config"
	"go-zero-demo/book/service/search/cmd/api/internal/handler"
	"go-zero-demo/book/service/search/cmd/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/search-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 全局中间件
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			logx.Info("global middleware")
			next(w, r)
		}
	})

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

/**
前提
user api 登录获取token
user rpc 供search调用
启动
go run search.go -f etc/search-api.yaml
调用/search/do验证jwt鉴权是否通过

curl -i -X GET \
    'http://127.0.0.1:8889/search/do?name=%E8%A5%BF%E6%B8%B8%E8%AE%B0'

TP/1.1 401 Unauthorized
Date: Tue, 18 May 2021 06:02:21 GMT
Content-Length: 0

带上token
curl -i -X GET \
    'http://127.0.0.1:8889/search/do?name=%E8%A5%BF%E6%B8%B8%E8%AE%B0' \
    -H 'authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjE0MDQyMzUsImlhdCI6MTYyMTMxNzgzNSwidXNlcklkIjoxfQ.w5NMs020yb21AQnovlWKrKAGdcHGqk0U2AKRQvIWQjo'

TP/1.1 200 OK
Content-Type: application/json
Date: Tue, 18 May 2021 06:04:49 GMT
Content-Length: 21

{"name":"","count":0}

最终结果
TP/1.1 200 OK
Content-Type: application/json
Date: Tue, 18 May 2021 06:51:54 GMT
Content-Length: 32

{"name":"西游记","count":100}

*/
