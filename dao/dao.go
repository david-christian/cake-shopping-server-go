package dao

import (
	"shopingCar_go/constants"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Dao struct {
	db *gorm.DB
}

func NewDao(db *gorm.DB) *Dao {
	return &Dao{
		db: db,
	}

}

func (d *Dao) dbConnection(ctx *gin.Context) *gorm.DB {
	value := ctx.Request.Context().Value(constants.ContextDbConnection)
	if value != nil {
		if db, ok := value.(*gorm.DB); ok {
			return db
		}
	}
	return nil

}
