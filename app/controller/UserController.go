package controller

import (
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (user UserController) User(ctx *gin.Context) {
	ctx.String(200, "userpages api")
}
