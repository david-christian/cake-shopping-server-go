package controllers

import (
	"net/http"
	"shopingCar_go/constants"
	"shopingCar_go/models"

	"github.com/gin-gonic/gin"
)

// func (c *UserController) Login(ctx *gin.Context) {
// 	var body models.UserAccount
// 	if err := ctx.ShouldBindJSON(&body); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"ok":    0,
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	token, err := c.service.Login(&body)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"ok":    0,
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{"ok": 1, "token": token})
// }

func (c *Controller) Register(ctx *gin.Context) {
	var body models.UserRegister
	if err := ctx.ShouldBind(&body); err != nil {
		ErrorResponse(string(constants.TypeError), ctx, err)
		return
	}
	token, err := c.service.Register(ctx, &body)
	if err != nil {
		ErrorResponse(string(constants.TypeError), ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"ok":    true,
		"token": token,
	})

}

func (c *Controller) Login(ctx *gin.Context) {
	var body models.UserLogin
	if err := ctx.ShouldBind(&body); err != nil {
		ErrorResponse(string(constants.TypeError), ctx, err)
		return
	}
	token, err := c.service.Login(&body)
	if err != nil {
		ErrorResponse(string(constants.TypeError), ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"ok":    true,
		"token": token,
	})

}

func (c *Controller) Get(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(string)
	data, err := c.service.Get(userId)
	if err != nil {
		ErrorResponse(string(constants.TypeError), ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"data": data,
	})

}
