package utils

import (
	"shopingCar_go/constants"

	"github.com/gin-gonic/gin"
)

func IsDevelop(ctx *gin.Context) bool {
	return constants.Environment == constants.Develop

}
