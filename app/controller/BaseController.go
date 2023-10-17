package controller

import (
	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (base BaseController) success(c *gin.Context) {
	c.String(200, "成功")
}

func (base BaseController) error(c *gin.Context) {
	c.String(200, "失败")
}
