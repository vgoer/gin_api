package main

import (
	"net/http"
	"vgoer/gin_api/bootstrap"
	"vgoer/gin_api/global"

	"github.com/gin-gonic/gin"
)

func main() {

	// 初始化配置
	bootstrap.InitializeConfig()

	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success!!!")

	server := gin.Default()

	// 测试路由
	server.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "safdsafdsaf")
	})

	server.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "hello pagedsdsafdsafsdfs")
	})

	// 启动服务器

	server.Run(":" + global.App.Config.App.Port)

}
