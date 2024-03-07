package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string         `gorm:"type:varchar(255)" json:"name"`
	Username  string         `gorm:"type:varchar(255)" json:"username"`
	Email     string         `gorm:"type:varchar(255)" json:"email"`
	Password  string         `gorm:"type:varchar(255)" json:"-"`
	ExpiredAt time.Time      `gorm:"type:timestamp" json:"expiredAt"`
	CreatedAt *time.Time     `gorm:"type:timestamp" json:"createdAt"`
	UpdatedAt *time.Time     `gorm:"type:timestamp" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"omitempty"`
}
