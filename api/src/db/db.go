package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)
func Init() {
	// connect to db
	dsn := "host=localhost user=postgres password=postgres dbname=linkkeep port=5432 sslmode=disable"

	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}