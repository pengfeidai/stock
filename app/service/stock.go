package service

import (
	"go-stock/app/common"
	"go-stock/app/database/mysql"
	"go-stock/app/model"
)

var logger = common.Logger

/**
 * 获取stock列表
 */
func GetStocks(currentPage, pageSize int, stock model.Stock) (list []model.Stock, count int, err error) {
	db := mysql.DB
	if stock.Code != "" {
		db = db.Where("code = ?", stock.Code)
	}

	if stock.Name != "" {
		db = db.Where("name like ?", "%"+stock.Name+"%")
	}

	if err = db.Where("trigger_date IS NOT NULL").Offset(currentPage - 1).Limit(pageSize).Find(&list).Offset(-1).Limit(-1).Count(&count).Error; err != nil {
		logger.Error("service.GetStocks DB Find error:", err)
		return
	}
	return
}

/**
 * 更新stock
 */
func UpdateStock(code, performance, stockType string) error {
	var stock model.Stock
	err := mysql.DB.Model(&stock).Where("code = ?", code).Updates(model.Stock{Performance: performance, Type: stockType}).Error
	if err != nil {
		logger.Error("service.UpdateStockerror:", err)
		return err
	}
	return nil
}
