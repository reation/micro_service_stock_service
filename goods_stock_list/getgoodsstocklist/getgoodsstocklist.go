// Code generated by goctl. DO NOT EDIT!
// Source: stock_goods_list.proto

package getgoodsstocklist

import (
	"context"

	"github.com/reation/micro_service_stock_service/protoc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetGoodsStockDataList     = protoc.GetGoodsStockDataList
	GetGoodsStockIDList       = protoc.GetGoodsStockIDList
	GetGoodsStockListRequest  = protoc.GetGoodsStockListRequest
	GetGoodsStockListResponse = protoc.GetGoodsStockListResponse

	GetGoodsStockList interface {
		GetGoodsStockByGoodsIDList(ctx context.Context, in *GetGoodsStockListRequest, opts ...grpc.CallOption) (*GetGoodsStockListResponse, error)
	}

	defaultGetGoodsStockList struct {
		cli zrpc.Client
	}
)

func NewGetGoodsStockList(cli zrpc.Client) GetGoodsStockList {
	return &defaultGetGoodsStockList{
		cli: cli,
	}
}

func (m *defaultGetGoodsStockList) GetGoodsStockByGoodsIDList(ctx context.Context, in *GetGoodsStockListRequest, opts ...grpc.CallOption) (*GetGoodsStockListResponse, error) {
	client := protoc.NewGetGoodsStockListClient(m.cli.Conn())
	return client.GetGoodsStockByGoodsIDList(ctx, in, opts...)
}
