package controller

import (
	"vgoer/gin_api/app/common/request"
	"vgoer/gin_api/app/common/response"
	"vgoer/gin_api/app/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (user UserController) User(ctx *gin.Context) {
	ctx.String(200, "userpages api")
}

func (user UserController) Register(ctx *gin.Context) {
	var form request.Register
	if err := ctx.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(ctx, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Register(form); err != nil {
		response.BusinessFail(ctx, err.Error())
	} else {
		response.Success(ctx, user)
	}

}
