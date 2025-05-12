package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb() (*gorm.DB, error) {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
