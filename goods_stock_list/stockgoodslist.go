package main

import (
	"flag"
	"fmt"

	"github.com/reation/micro_service_stock_service/goods_stock_list/internal/config"
	"github.com/reation/micro_service_stock_service/goods_stock_list/internal/server"
	"github.com/reation/micro_service_stock_service/goods_stock_list/internal/svc"
	"github.com/reation/micro_service_stock_service/protoc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/stockgoodslist.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		protoc.RegisterGetGoodsStockListServer(grpcServer, server.NewGetGoodsStockListServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
