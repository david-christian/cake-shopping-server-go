package dao

import (
	"fmt"
	"shopingCar_go/models"

	"gorm.io/gorm"
)

// func (d *Dao) IsEmailExists(body *models.UserRegister) (bool, error) {
// 	var user models.UserAccount
// 	if err := d.db.Where("email = ?", body.Email).First(&user).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return false, nil
// 		} else {
// 			return false, err
// 		}
// 	}
// 	return true, nil

// }

func (d *Dao) CreateUserAccount(db *gorm.DB, body *models.UserRegister, userId string) error {
	// 要創建 user 的資料
	user := models.UserAccount{
		UserId:   userId,
		Password: body.Password,
		Name:     body.Name,
		Email:    body.Email,
		Phone:    body.Phone,
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil

}

func (d *Dao) CreateUser(db *gorm.DB, userId string) error {
	user := models.User{
		UserId: userId,
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil

}

func (d *Dao) SelectUser(email string) (*models.UserAccount, error) {
	var user models.UserAccount
	err := d.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil

}

func (d *Dao) GetUser(userId string) (*models.UserAccount, error) {
	var user models.UserAccount
	var user2 *models.UserAccount
	fmt.Printf(">>>>>>>>>>>>%p", user2)
	err := d.db.
		Select("name", "email", "phone").
		Where("user_id = ?", userId).
		First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil

}
