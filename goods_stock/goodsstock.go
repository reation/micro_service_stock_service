package main

import (
	"flag"
	"fmt"

	"github.com/reation/micro_service_stock_service/goods_stock/internal/config"
	"github.com/reation/micro_service_stock_service/goods_stock/internal/server"
	"github.com/reation/micro_service_stock_service/goods_stock/internal/svc"
	"github.com/reation/micro_service_stock_service/protoc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/goodsstock.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		protoc.RegisterStockOpServer(grpcServer, server.NewStockOpServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
