package controllers

import (
	"net/http"

	"shopingCar_go/constants"
	"shopingCar_go/customerrors"
	"shopingCar_go/models"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Register(ctx *gin.Context) {
	var body models.UserRegister
	if err := ctx.ShouldBind(&body); err != nil {
		customerrors.ErrorResponse(string(constants.TypeError), ctx, err)
		return
	}
	token, err := c.service.Register(ctx, &body)
	if err != nil {
		customerrors.ErrorResponse(string(constants.TypeError), ctx, err)
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
		customerrors.ErrorResponse(string(constants.TypeError), ctx, err)
		return
	}
	token, err := c.service.Login(&body)
	if err != nil {
		customerrors.ErrorResponse(string(constants.TypeError), ctx, err)
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
		customerrors.ErrorResponse(string(constants.TypeError), ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"data": data,
	})

}
