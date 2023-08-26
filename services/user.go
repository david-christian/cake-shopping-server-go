package services

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"os"
	"shopingCar_go/constants"
	"shopingCar_go/models"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	secretKey = []byte(os.Getenv(constants.EnvSecretkey))
)

type CustomClaims struct {
	UserId string
	jwt.RegisteredClaims
}

// func (s *UserService) Login(userAccount *models.UserAccount) (int, error) {
// 	// 這是 return 範例
// 	return 1, nil
// 	// 繼續寫這層的邏輯 + dao
// }

func (s *Service) Register(ctx *gin.Context, body *models.UserRegister) (string, error) {
	hash := sha256.Sum256([]byte(body.Password))
	// The type of hash is byte, so it needs to be converted to a string.
	body.Password = hex.EncodeToString(hash[:])
	userId := uuid.New().String()
	conn := s.dbConnection(ctx)
	// Control the transaction manually
	// transaction start
	tx := conn.Begin()
	if err := tx.Error; err != nil {
		return "", err
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := s.dao.CreateUser(tx, userId); err != nil {
		tx.Rollback()
		return "", err
	}
	if err := s.dao.CreateUserAccount(tx, body, userId); err != nil {
		tx.Rollback()
		return "", err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return "", err
	}
	// transaction end

	// created jwt token
	claims := CustomClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	t, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return t, nil

}

func (s *Service) Login(body *models.UserLogin) (string, error) {
	var user *models.UserAccount
	user, err := s.dao.SelectUser(body.Email)
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256([]byte(body.Password))
	body.Password = hex.EncodeToString(hash[:])
	if user.Password != body.Password {
		return "", errors.New("password do not match")
	}
	claims := CustomClaims{
		user.UserId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	t, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return t, nil

}

func (s *Service) Get(userId string) (*models.UserAccount, error) {
	var user *models.UserAccount
	user, err := s.dao.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
