package models

import (
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string `gorm:"not null"`
}

func (u *User) CreateUser() (*User, error) {
	err := db.Create(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) GetUser() (*User, error) {
	err := db.Where("username = ?", u.Username).First(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) GetUserByID(id uint) (*User, error) {
	err := db.Where("id = ?", id).Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}