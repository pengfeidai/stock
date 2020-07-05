package router

import (
	"go-stock/app/controller"

	"github.com/gin-gonic/gin"
)

func InitStockRouter(group *gin.RouterGroup) {
	router := group.Group("")
	{
		router.GET("/stocks", controller.GetStocks)
		router.POST("/stock", controller.UpdateStock)
	}
}
