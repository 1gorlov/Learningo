package models

import (
	"learningo/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"not null;unique_index"`
	Password     string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
}

var db = database.Connect()

func (user *User) All() (*[]User, error) {
	var users []User
	err := db.Find(&users).Error
	return &users, err
}

func (user *User) Find(id string) (*User, error) {
	err := db.First(&user, id).Error
	return user, err
}

func (user *User) Create() (*User, error) {
	err := db.Create(&user).Error
	return user, err
}

// Needs more modification
func (user *User) Update(id string, nu User) (*User, error) {
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}
	user.Name = nu.Name
	user.Email = nu.Email

	err := db.Save(&user).Error

	return user, err
}

func (user *User) Delete() error {
	err := db.Delete(&user).Error
	return err
}
