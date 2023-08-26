package services

import (
	"shopingCar_go/constants"
	"shopingCar_go/dao"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Service struct {
	dao *dao.Dao
}

func NewService(dao *dao.Dao) *Service {
	return &Service{dao}

}

func (s *Service) dbConnection(ctx *gin.Context) *gorm.DB {
	value := ctx.Request.Context().Value(constants.ContextDbConnection)
	if value != nil {
		if db, ok := value.(*gorm.DB); ok {
			return db
		}
	}
	return nil

}
