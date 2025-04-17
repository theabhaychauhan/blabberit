package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"uniqueIndex;not null"`
	PublicKey string `gorm:"type:text;not null"`
}
