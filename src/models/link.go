package models

import (
	config "github.com/tmunongo/linkkeep/api/src/db"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Title       string `gorm:"not null"`
	URL         string `gorm:"unique"`
	Description string `gorm:"string"`
	Cover       string `gorm:"string"`
	UserID      uint   `gorm:"not null"`
}

func init() {
	config.Init()
	db = config.GetDB()
	db.AutoMigrate(&Link{})
}

func (link *Link) Save() (*Link, error) {
	err := db.Create(&link).Error

	if err != nil {
		return nil, err
	}

	return link, nil
}
