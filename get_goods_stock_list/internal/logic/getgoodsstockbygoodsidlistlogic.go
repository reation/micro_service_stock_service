package logic

import (
	"context"
	"github.com/reation/micro_service_stock_service/config"
	"github.com/reation/micro_service_stock_service/model"

	"github.com/reation/micro_service_stock_service/get_goods_stock_list/internal/svc"
	"github.com/reation/micro_service_stock_service/protoc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsStockByGoodsIDListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGoodsStockByGoodsIDListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsStockByGoodsIDListLogic {
	return &GetGoodsStockByGoodsIDListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGoodsStockByGoodsIDListLogic) GetGoodsStockByGoodsIDList(in *protoc.GetGoodsStockRequest) (*protoc.GetGoodsStockResponse, error) {
	states, goodsStockList, _ := l.GetGoodsStockList(in.GetGoodsIDList())
	if states != config.GOODS_STOCK_STATE_NORMAL {
		return &protoc.GetGoodsStockResponse{States: states, GoodsStockList: nil}, nil
	}
	var resp = make([]*protoc.StockGoodsData, len(*goodsStockList))
	for k, v := range *goodsStockList {
		resp[k] = &protoc.StockGoodsData{
			GoodsId:  v.GoodsId,
			GoodsNum: v.GoodsNum,
		}
	}

	return &protoc.GetGoodsStockResponse{States: states, GoodsStockList: resp}, nil
}

func (l *GetGoodsStockByGoodsIDListLogic) GetGoodsStockList(goodsIDList []*protoc.StockOpGoodsID) (int64, *[]model.GoodsStock, error) {
	var goodsIDS = make([]int64, len(goodsIDList))
	for k, v := range goodsIDList {
		goodsIDS[k] = v.GoodsId
	}
	goodsStockList, err := l.svcCtx.StockModel.GetGoodsStockListByGoodsIDList(l.ctx, goodsIDS)
	if err == model.ErrNotFound {
		return config.GOODS_STOCK_STATE_EMPTY, nil, nil
	}
	if err != nil {
		return config.GOODS_STOCK_STATE_ERROR, nil, nil
	}

	return config.GOODS_STOCK_STATE_NORMAL, goodsStockList, nil
}
