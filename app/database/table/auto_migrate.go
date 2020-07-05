package table

import (
	"go-stock/app/common"
	"go-stock/app/database/mysql"
	"go-stock/app/model"
)

var logger = common.Logger

func Init() {
	db := mysql.DB
	db.AutoMigrate(&model.User{})
	logger.Info("AutoMigrate tables success.")
}
