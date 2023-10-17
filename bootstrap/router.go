package bootstrap

import (
	"vgoer/gin_api/global"
	"vgoer/gin_api/routes"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	server := gin.Default()

	// 前端项目静态资源
	server.StaticFile("/", "./static/dist/index.html")
	server.Static("/assets", "./static/dist/assets")
	server.StaticFile("/favicon.ico", "./static/dist/favicon.ico")
	// 其他静态资源
	server.Static("/public", "./static")
	server.Static("/storage", "./storage/app/public")

	// 注册 api 分组路由
	// 修改的
	routes.InitRouteApi(server)

	return server
}

// RunServer 启动服务器
func RunServer() {
	r := setupRouter()

	r.Run(":" + global.App.Config.App.Port)
}
