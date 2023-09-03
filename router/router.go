package router

import (
	"context"
	"net/http"
	"os"
	"shopingCar_go/constants"
	"shopingCar_go/controllers"
	"strings"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func NewRouter(c *controllers.Controller, connection *gorm.DB) *gin.Engine {
	router := gin.Default()
	// database connection set in context
	router.Use(databaseMiddleware(connection))

	// User routes
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/register", c.Register)
		userRoutes.POST("/login", c.Login)
		userRoutes.GET("/get", authenticate, c.Get)
	}

	photoRoutes := router.Group("/photo")
	{
		photoRoutes.POST("/upload", c.UploadImage)
	}

	return router

}

func databaseMiddleware(connection *gorm.DB) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		ctx := gctx.Request.Context()
		ctx = context.WithValue(ctx, constants.ContextDbConnection, connection.WithContext(ctx))
		gctx.Request = gctx.Request.WithContext(ctx)
		gctx.Next()
	}

}

func authenticate(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")
	parts := strings.Split(bearerToken, " ")
	var token string
	if len(parts) == 2 && parts[0] == "Bearer" {
		token = parts[1]
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"status":  constants.Invalid,
			"message": "Invalid Bearer token format",
		})
		return
	}

	var secretKey = []byte(os.Getenv(constants.EnvSecretkey))
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"status":  constants.Unauthorized,
			"message": "Token Unauthorized",
		})
		return
	}

	clamis, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"status":  constants.Unexpected,
			"message": "Unable to retrieve token information.",
		})
	}

	ctx.Set("userId", clamis["UserId"].(string))
	ctx.Next()

}
