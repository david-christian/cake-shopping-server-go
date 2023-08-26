package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"shopingCar_go/constants"
	"shopingCar_go/controllers"
	"shopingCar_go/dao"
	"shopingCar_go/router"
	"shopingCar_go/services"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbName = os.Getenv(constants.DbName)
	dbUser = os.Getenv(constants.DbUser)
	dbPass = os.Getenv(constants.DbPass)
	dbhost = os.Getenv(constants.DbHost)
)

const (
	dbSlowThreshold   = 200 * time.Millisecond
	dbMaxIdleConns    = 2
	dbMaxOpenConns    = 100
	dbConnMaxIdleTime = 50 * time.Second
	dbConnMaxLifetime = 50 * time.Second
)

func initializeControllers(db *gorm.DB) *controllers.Controller {
	initializeDao := dao.NewDao(db)
	initializeService := services.NewService(initializeDao)
	initializeController := controllers.NewController(initializeService)

	return initializeController

}

func main() {
	dbConfig := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		dbUser, dbPass, dbhost, dbName)
	connection, err := gorm.Open(mysql.Open(dbConfig), &gorm.Config{
		SkipDefaultTransaction: false,
		Logger: logger.New(log.New(os.Stdout, "", log.LstdFlags),
			logger.Config{
				SlowThreshold:             dbSlowThreshold,
				Colorful:                  false,
				IgnoreRecordNotFoundError: true,
				ParameterizedQueries:      false,
				LogLevel:                  logger.Info,
			}),
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatalln(err)
	}
	db, err := connection.DB()
	if err != nil {
		log.Fatalln(err)
	}

	db.SetConnMaxIdleTime(dbConnMaxIdleTime)
	db.SetConnMaxLifetime(dbConnMaxLifetime)
	db.SetMaxIdleConns(dbMaxIdleConns)
	db.SetMaxOpenConns(dbMaxOpenConns)

	c := initializeControllers(connection)

	server := router.NewRouter(c, connection)
	server.Run(":8080")

}
