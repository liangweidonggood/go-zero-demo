package main

import (
	"flag"
	"fmt"

	"go-zero-demo/datacenter/service/votes/cmd/rpc/internal/config"
	"go-zero-demo/datacenter/service/votes/cmd/rpc/internal/server"
	"go-zero-demo/datacenter/service/votes/cmd/rpc/internal/svc"
	"go-zero-demo/datacenter/service/votes/cmd/rpc/votes"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/votes.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewVotesServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		votes.RegisterVotesServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
