package user

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetDB(db *gorm.DB) {
	DB = db
}

func CreateUser(u *User) error {
	return DB.Create(u).Error
}

func FindUserByUsername(username string) (*User, error) {
	var u User
	error := DB.Where("username = ?", username).First(&u).Error
	if error != nil {
		return nil, error
	}
	return &u, nil
}
