package model

import (
	"time"

	"github.com/google/uuid"
)

// base sql constraints like the definition of ID algo...UUID in this case.
// you could also use gorm.Model if you ok with the embedded types that comes with it
type Base struct {
	ID        uuid.UUID `gorm:"primary_key; unique; type:uuid; column:id; default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type User struct {
	Base
	Username string `gorm:"unique_index;not_null" json:"username"`
	Email    string `gorm:"unique_index;not_null" json:"email"`
	Password string `gorm:"not_null" json:"password"`
}
