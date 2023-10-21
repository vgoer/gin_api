package routes

import (
	"vgoer/gin_api/app/common"
	"vgoer/gin_api/app/controller"
	"vgoer/gin_api/app/middleware"
	"vgoer/gin_api/app/services"

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

	// 中间件
	userRoutes.Use(gin.Logger(), middleware.Cors())
	userRoutes.Use(gin.Logger(), middleware.CustomRecovery())

	userRoutes.GET("/", controller.UserController{}.User)
	userRoutes.POST("/auth/register", controller.UserController{}.Register)
	userRoutes.POST("/auth/login", controller.Login)

	authRouter := userRoutes.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/auth/info", controller.Info)
		authRouter.POST("/auth/logout", controller.Logout)
		authRouter.POST("/imgeUpload", common.ImageUpload)
	}
}
