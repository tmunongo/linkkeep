package models

import (
	config "github.com/tmunongo/linkkeep/api/src/db"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Link     string `gorm:"unique"`
	UserID   uint   `gorm:"not null"`
	User     User   `gorm:"foreignkey:UserID"`
}

func init() {
	config.Init()
	db = config.GetDB()
	db.AutoMigrate(&Link{})
}