// Code generated by goctl. DO NOT EDIT!
// Source: stock_goods.proto

package stockop

import (
	"context"

	"github.com/reation/micro_service_stock_service/protoc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetGoodsStockRequest  = protoc.GetGoodsStockRequest
	GetGoodsStockResponse = protoc.GetGoodsStockResponse
	StockGoodsData        = protoc.StockGoodsData
	StockOpGoodsID        = protoc.StockOpGoodsID

	StockOp interface {
		GetGoodsStockByGoodsIDList(ctx context.Context, in *GetGoodsStockRequest, opts ...grpc.CallOption) (*GetGoodsStockResponse, error)
	}

	defaultStockOp struct {
		cli zrpc.Client
	}
)

func NewStockOp(cli zrpc.Client) StockOp {
	return &defaultStockOp{
		cli: cli,
	}
}

func (m *defaultStockOp) GetGoodsStockByGoodsIDList(ctx context.Context, in *GetGoodsStockRequest, opts ...grpc.CallOption) (*GetGoodsStockResponse, error) {
	client := protoc.NewStockOpClient(m.cli.Conn())
	return client.GetGoodsStockByGoodsIDList(ctx, in, opts...)
}
