package model

import (
	"time"

	"github.com/google/uuid"
)

// custom initialization from gorm.Model, we are calling it Base
// can be used as standard for all structs that need the same default consts like ID as uuid
type Base struct {
	ID        uuid.UUID `gorm:"primary_key; unique; type:uuid; default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// User object with some rules for SQL schema generated from the "gorm:" tags
type User struct {
	Base
	Username string `gorm:"unique_index;not null" json:"username"`
	Email    string `gorm:"unique_index;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}
