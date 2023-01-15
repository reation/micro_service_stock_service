package main

import (
	"context"
	"github.com/reation/micro_service_stock_service/protoc"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	StockGoodsListAddress = "192.168.1.104:9092"
	StockGoodsAddress     = "192.168.1.104:9093"
)

func main() {
	//stockGoodsList()
	stockGoods()
}

func stockGoods() {

	conn, err := grpc.Dial(StockGoodsAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := protoc.NewGetGoodsStockClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetGoodsStockByGoodsID(ctx, &protoc.GetGoodsStockRequest{GoodsID: 1})
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	log.Printf("states: %d", r.GetStates())
	log.Printf("goods_nums: %d", r.GetGoodsNums())
}

func stockGoodsList() {

	conn, err := grpc.Dial(StockGoodsListAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := protoc.NewGetGoodsStockListClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var idList = make([]*protoc.GetGoodsStockIDList, 4)
	idList[0] = &protoc.GetGoodsStockIDList{GoodsId: 1}
	idList[1] = &protoc.GetGoodsStockIDList{GoodsId: 2}
	idList[2] = &protoc.GetGoodsStockIDList{GoodsId: 3}
	idList[3] = &protoc.GetGoodsStockIDList{GoodsId: 4}
	r, err := c.GetGoodsStockByGoodsIDList(ctx, &protoc.GetGoodsStockListRequest{GoodsIDList: idList})
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	log.Printf("states: %d", r.GetStates())
	for _, v := range r.GetGoodsStockList() {
		log.Println(v)
	}

}
