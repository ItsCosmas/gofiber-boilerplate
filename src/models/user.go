package models

import (
	"gorm.io/gorm"
)

// User domain model
type User struct {
	gorm.Model
	ID             int
	ExternalID     string `gorm:"unique;NOT NULL;"`
	FullName       string `gorm:"NOT NULL"`
	ProfilePicture string `gorm:"size:255"`
	Bio            string `gorm:"size:255"`
	Email          string `gorm:"NOT NULL;UNIQUE"`
	Password       string `gorm:"NOT NULL"`
	Role           string `gorm:"NOT NULL;size:255;DEFAULT:'standard'"`
	Active         bool   `gorm:"NOT NULL;DEFAULT:true"`
}
