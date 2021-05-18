package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-demo/datacenter/common/shared"
	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/middleware"
	"net/http"
	"strings"

	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/config"
	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/handler"
	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/datacenter-api.yaml", "the config file")

func dirhandler(patern, filedir string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		handler := http.StripPrefix(patern, http.FileServer(http.Dir(filedir)))
		handler.ServeHTTP(w, req)

	}
}

func staticFileHandler(engine *rest.Server) {
	//这里注册
	dirlevel := []string{":1", ":2", ":3", ":4", ":5", ":6", ":7", ":8"}
	patern := "/static/"
	dirpath := "./assets/"
	for i := 1; i < len(dirlevel); i++ {
		path := "/" + strings.Join(dirlevel[:i], "/")
		//最后生成 /asset
		engine.AddRoute(
			rest.Route{
				Method:  http.MethodGet,
				Path:    path,
				Handler: dirhandler(patern, dirpath),
			})
	}

}

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	//server := rest.MustNewServer(c.RestConf)
	server := rest.MustNewServer(c.RestConf, rest.WithNotAllowedHandler(middleware.NewCorsMiddleware().Handler()))
	defer server.Stop()

	server.Use(middleware.NewCorsMiddleware().Handle)
	//静太文件处理
	staticFileHandler(server)
	// 设置错误处理函数
	httpx.SetErrorHandler(shared.ErrorHandler)

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

/**
启动
go run datacenter.go -f etc/datacenter-api.yaml
测试获取站点信息
curl -i -X GET http://127.0.0.1:8888/common/appinfo -H 'content-type: application/json' -d '{"beid":1}'
测试七牛上传凭证
curl -i -X POST http://127.0.0.1:8888/common/qiuniu/token -H 'content-type: application/json' -d '{"beid":1}'

curl -i -X POST http://127.0.0.1:8888/common/snsinfo -H 'content-type: application/json' -d '{"beid":1,"ptyid":1,"back_url":"123"}'
*/
