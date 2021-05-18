package main

import (
	"flag"
	"fmt"

	"go-zero-demo/datacenter/service/common/cmd/rpc/common"
	"go-zero-demo/datacenter/service/common/cmd/rpc/internal/config"
	"go-zero-demo/datacenter/service/common/cmd/rpc/internal/server"
	"go-zero-demo/datacenter/service/common/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/common.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewCommonServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		common.RegisterCommonServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

/**
启动
go run common.go -f etc/common.yaml

etcdctl member list
etcdctl get datacenterCommon.rpc --prefix

datacenterCommon.rpc/112442081613213983
127.0.0.1:8080

*/
