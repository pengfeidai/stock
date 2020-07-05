package router

import (
	"fmt"
	"go-stock/app/config"

	"go-stock/app/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	config := config.Conf
	gin.SetMode(config.Mode)
	router := gin.New()
	// 404处理
	router.NoRoute(func(c *gin.Context) {
		ctx := middleware.Context{Ctx: c}
		path := c.Request.URL.Path
		method := c.Request.Method
		ctx.Response(404, fmt.Sprintf("%s %s not found", method, path), nil)
	})

	// 中间件
	router.Use(
		cors.Default(),
		middleware.Recovery(),
		middleware.Logger(),
	)

	// 路由分组加载
	group := router.Group(config.Url.Prefix)
	InitStockRouter(group)

	return router
}
