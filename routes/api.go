package routes

import (
	"vgoer/gin_api/app/controller"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func InitRouteApi(r *gin.Engine) {
	AdminRoutesInit(r)
	UserRoutesInit(r)
}

// admin 路由
func AdminRoutesInit(r *gin.Engine) {

	adminRoutes := r.Group("/api/admin")

	adminRoutes.GET("/", controller.AdminController{}.Admin)
}

// user 路由
func UserRoutesInit(r *gin.Engine) {

	userRoutes := r.Group("/api/user")

	userRoutes.GET("/", controller.UserController{}.User)
}
