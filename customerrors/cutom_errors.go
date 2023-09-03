package customerrors

import (
	"errors"
	"fmt"
	"net/http"

	"shopingCar_go/constants"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func ErrorResponse(errorCode string, ctx *gin.Context, err error) {
	var status string
	var message string

	switch {
	case errorCode == string(constants.TypeError):
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
	case errorCode != string(constants.Unexpected):
		status = errorCode
		message = err.Error()
	}

	if status == "" && message == "" || errorCode == string(constants.Unexpected) {
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
