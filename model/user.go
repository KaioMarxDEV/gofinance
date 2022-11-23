package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique_index;not_null" json:"username"`
	Email    string `gorm:"unique_index;not_null" json:"email"`
	Password string `gorm:"not_null" json:"password"`
}
