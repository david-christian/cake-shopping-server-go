package models

import "time"

type UserLogin struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserRegister struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

// table

type UserAccount struct {
	UserId    string    `gorm:"column:user_id;primaryKey"`
	Password  string    `gorm:"column:password;not null"`
	Name      string    `gorm:"column:name;not null"`
	Email     string    `gorm:"column:email;not null;uniqueIndex"`
	Phone     string    `gorm:"column:phone;not null"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
}

func (u UserAccount) TableName() string {
	return "user_account"

}

type User struct {
	UserId    string    `gorm:"column:user_id;primaryKey"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
}

func (u User) TableName() string {
	return "user"

}
