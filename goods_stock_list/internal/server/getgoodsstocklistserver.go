// Code generated by goctl. DO NOT EDIT!
// Source: stock_goods_list.proto

package server

import (
	"context"

	"github.com/reation/micro_service_stock_service/goods_stock_list/internal/logic"
	"github.com/reation/micro_service_stock_service/goods_stock_list/internal/svc"
	"github.com/reation/micro_service_stock_service/protoc"
)

type GetGoodsStockListServer struct {
	svcCtx *svc.ServiceContext
	protoc.UnimplementedGetGoodsStockListServer
}

func NewGetGoodsStockListServer(svcCtx *svc.ServiceContext) *GetGoodsStockListServer {
	return &GetGoodsStockListServer{
		svcCtx: svcCtx,
	}
}

func (s *GetGoodsStockListServer) GetGoodsStockByGoodsIDList(ctx context.Context, in *protoc.GetGoodsStockListRequest) (*protoc.GetGoodsStockListResponse, error) {
	l := logic.NewGetGoodsStockByGoodsIDListLogic(ctx, s.svcCtx)
	return l.GetGoodsStockByGoodsIDList(in)
}
