package models

import (
	"errors"
	"github.com/MirMonajir244/BooksOnline/utils"
	"gorm.io/gorm"
)

type Users struct {
	UserName     string `json:"userName" gorm:"not null;unique"`
	UserID       int64  `json:"userID" gorm:"not null;unique;primaryKey"`
	UserEmail    string `json:"userEmail" gorm:"not null;unique"`
	UserPassword string `json:"userPassword" gorm:"not null"`
	Books        []Book `json:"books" gorm:"foreignKey:UserID"`
}

func (user *Users) Create(db *gorm.DB) error {
	hashPass, hashErr := utils.HashPassword(user.UserPassword)
	if hashErr != nil {
		return errors.New("Internal server Issue")
	}
	user.UserPassword = hashPass
	err := db.Create(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrDuplicatedKey) {
		return errors.New("User already exists" + err.Error())
	}
	return nil
}

func (user *Users) ValidateCredentials(db *gorm.DB, userIdentifier, password string) error {
	// Variable to hold the user
	var foundUser Users

	// Query to find the user by email or username
	err := db.Where("user_email = ? OR user_name = ?", userIdentifier, userIdentifier).First(&foundUser).Error
	if err != nil {
		return errors.New("user not found")
	}

	// Check if the provided password matches the hashed password
	if !utils.CheckPasswordHash(password, foundUser.UserPassword) {
		return errors.New("incorrect password")
	}

	// Update the user variable to return valid user information
	*user = foundUser

	return nil // Valid credentials
}
