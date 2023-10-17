package controller

import (
	"github.com/gin-gonic/gin"
)

type AdminController struct {
}

func (admin AdminController) Admin(ctx *gin.Context) {
	ctx.String(200, "admin pages")
}
