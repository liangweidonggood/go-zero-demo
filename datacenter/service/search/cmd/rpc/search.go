package main

import (
	"flag"
	"fmt"

	"go-zero-demo/datacenter/service/search/cmd/rpc/internal/config"
	"go-zero-demo/datacenter/service/search/cmd/rpc/internal/server"
	"go-zero-demo/datacenter/service/search/cmd/rpc/internal/svc"
	"go-zero-demo/datacenter/service/search/cmd/rpc/search"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/search.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewSearchServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		search.RegisterSearchServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
