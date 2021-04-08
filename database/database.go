package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseService struct {
	*gorm.DB
}

func InitDatabase() (*DatabaseService, error) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DatabaseService{
		db,
	}, nil
}

func Connect() (*DatabaseService, error) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DatabaseService{
		db,
	}, nil
}
