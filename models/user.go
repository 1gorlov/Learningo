package models

import (
	"learningo/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

func All() ([]User, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	var users []User

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func Find(id string) (*User, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	var user User

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func Create(user User) (*User, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
