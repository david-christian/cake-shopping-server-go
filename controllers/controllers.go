package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"shopingCar_go/constants"
	"shopingCar_go/services"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// type Controllers struct {
// 	UserController *UserController
// 	// Other controllers go here
// }

type Controller struct {
	service *services.Service
}

func NewController(service *services.Service) *Controller {
	return &Controller{service}

}

// func ErrorResponse(ctx *gin.Context, err error) {
// 	errorCode, message := customerrors.DefinitionError(err)
// 	if errorCode == string(constants.Duplicated) {
// 		ctx.JSON(http.StatusConflict, gin.H{
// 			"ok":      false,
// 			"status":  errorCode,
// 			"message": message,
// 		})
// 		return
// 	}
// 	ctx.JSON(http.StatusBadRequest, gin.H{
// 		"ok":      false,
// 		"status":  errorCode,
// 		"message": message,
// 	})

// }

func ErrorResponse(errorCode string, ctx *gin.Context, err error) {
	var status string
	var message string

	switch errorCode {
	case string(constants.TypeError):
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			mysqlErr = err.(*mysql.MySQLError)
			if mysqlErr.Number == 1062 {
				status = string(constants.Duplicated)
				message = "duplicate entry error"
			}
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = string(constants.Invalid)
			message = "no rows found"
		}
	case string(constants.MaxFileError):
		status = string(constants.Invalid)
		message = "invalid file extension"

	case string(constants.NotAllowedExtensionError):
		status = string(constants.Invalid)
		message = "extension not allowed"

	case string(constants.FileConversionError):
		status = string(constants.Unexpected)
		message = "file conversion error"
	}

	if status == "" && message == "" {
		status = string(constants.Unexpected)
		message = "unexpected error"
	}

	if constants.Environment == constants.Develop {
		message += fmt.Sprintf(" system error: %s", err.Error())
	}

	if status == string(constants.Duplicated) {
		ctx.JSON(http.StatusConflict, gin.H{
			"ok":      false,
			"status":  status,
			"message": message,
		})
		return
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"status":  status,
			"message": message,
		})
	}
}
