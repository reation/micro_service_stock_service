package svc

import (
	"fmt"
	"github.com/reation/micro_service_stock_service/get_goods_stock_list/internal/config"
	"github.com/reation/micro_service_stock_service/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	StockModel model.GoodsStockModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	dataSourceUrl := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		c.Mysql.StockTable.User,
		c.Mysql.StockTable.Passwd,
		c.Mysql.StockTable.Addr,
		c.Mysql.StockTable.Port,
		c.Mysql.StockTable.DBName,
	)
	return &ServiceContext{
		Config:     c,
		StockModel: model.NewGoodsStockModel(sqlx.NewMysql(dataSourceUrl)),
	}
}
