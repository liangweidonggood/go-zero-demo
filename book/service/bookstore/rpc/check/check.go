package main

import (
	"flag"
	"fmt"

	"go-zero-demo/book/service/bookstore/rpc/check/check"
	"go-zero-demo/book/service/bookstore/rpc/check/internal/config"
	"go-zero-demo/book/service/bookstore/rpc/check/internal/server"
	"go-zero-demo/book/service/bookstore/rpc/check/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/check.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewCheckerServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		check.RegisterCheckerServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
