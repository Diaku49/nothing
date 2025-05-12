package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Email          string  `gorm:"uniqueIndex;not null;size:255"`
	FullName       string  `gorm:"size:255"`
	Phone          *string `gorm:"size:16"`
	Password       *string
	RefreshTokenoh *string `gorm:"type:text"`
}
