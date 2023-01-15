package logic

import (
	"context"
	"github.com/reation/micro_service_stock_service/config"
	"github.com/reation/micro_service_stock_service/model"

	"github.com/reation/micro_service_stock_service/goods_stock/internal/svc"
	"github.com/reation/micro_service_stock_service/protoc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsStockByGoodsIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGoodsStockByGoodsIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsStockByGoodsIDLogic {
	return &GetGoodsStockByGoodsIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGoodsStockByGoodsIDLogic) GetGoodsStockByGoodsID(in *protoc.GetGoodsStockRequest) (*protoc.GetGoodsStockResponse, error) {
	resp, err := l.svcCtx.StockModel.GetGoodsStockInfoByGoodsID(l.ctx, in.GetGoodsID())
	if err == model.ErrNotFound {
		return &protoc.GetGoodsStockResponse{States: config.GOODS_STOCK_STATE_EMPTY, GoodsNums: 0}, nil
	}
	if err != nil {
		return &protoc.GetGoodsStockResponse{States: config.GOODS_STOCK_STATE_ERROR, GoodsNums: 0}, nil
	}

	return &protoc.GetGoodsStockResponse{States: config.GOODS_STOCK_STATE_NORMAL, GoodsNums: resp.GoodsNum}, nil
}
