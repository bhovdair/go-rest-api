package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductCategory struct {
	gorm.Model
	Code      string         `gorm:"type:varchar(255)" json:"code"`
	Name      string         `gorm:"type:varchar(255)" json:"name"`
	CreatedAt *time.Time     `gorm:"type:timestamp" json:"createdAt"`
	UpdatedAt *time.Time     `gorm:"type:timestamp" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"omitempty"`
}
