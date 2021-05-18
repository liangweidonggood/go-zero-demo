package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-demo/book/common/errorx"
	"net/http"

	"go-zero-demo/book/service/user/cmd/api/internal/config"
	"go-zero-demo/book/service/user/cmd/api/internal/handler"
	"go-zero-demo/book/service/user/cmd/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// logx 根据配置初始化
	logx.MustSetup(c.Logx)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	// 开启自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

/**
启动
go run user.go -f etc/user-api.yaml
测试jwt
curl -i -X POST \
    http://127.0.0.1:8888/user/login \
    -H 'content-type: application/json' \
    -d '{
      "username":"zs666",
      "password":"123456"
     }'

HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 18 May 2021 03:32:20 GMT
Content-Length: 277

{"id":1,"name":"张三","gender":"男","number":"","username":"","accessToken":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjEzOTUxNDAsImlhdCI6MTYyMTMwODc0MCwidXNlcklkIjoxfQ.ndKFofQbVYMyEB1WmKCMnXEDmDJqAChfK5FiHmym30c","accessExpire":1621395140,"refreshAfter":1621351940}

*/
