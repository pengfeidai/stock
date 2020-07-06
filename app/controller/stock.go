package controller

import (
	"go-stock/app/common"
	"go-stock/app/middleware"
	"go-stock/app/model"
	"go-stock/app/schema"
	"go-stock/app/service"

	"github.com/gin-gonic/gin"
)

var logger = common.Logger

/**
 * 获取stock列表
 */
func GetStocks(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}

	var p schema.GetStocks
	if err := ctx.ValidateQuery(&p); err != nil {
		return
	}

	var stock = model.Stock{}
	if code, ok := c.GetQuery("code"); ok {
		stock.Code = code
	}

	if name, ok := c.GetQuery("name"); ok {
		stock.Name = name
	}

	list, count, err := service.GetStocks(p.CurrentPage, p.PageSize, stock)
	if err != nil {
		ctx.Response(common.ERROR, common.SERVER_ERROR, nil)
		return
	}
	data := map[string]interface{}{
		"count":       count,
		"currentPage": p.CurrentPage,
		"pageSize":    p.PageSize,
		"list":        list,
	}
	ctx.Response(common.SUCCESS, nil, data)
}

/**
 * stock 数据更新
 */
func UpdateStock(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}

	var p schema.UpdateStock
	if err := ctx.ValidateJSON(&p); err != nil {
		return
	}

	if err := service.UpdateStock(p.Code, p.Performance, p.StockType); err != nil {
		ctx.Response(common.ERROR, common.UPDATE_STOCK_FAIL, nil)
		return
	}
	data := map[string]string{
		"code": p.Code,
	}
	ctx.Response(common.SUCCESS, nil, data)
}
