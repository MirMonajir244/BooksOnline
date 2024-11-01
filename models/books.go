package models

import (
	"errors"
	_ "github.com/MirMonajir244/BooksOnline/shared"
	"gorm.io/gorm"
)

type Book struct {
	Name        string  `json:"name" gorm:"not null;unique"`
	Description string  `json:"description"`
	Author      string  `json:"author" gorm:"not null"`
	Price       float64 `json:"price"`
	UserID      string  `json:"userID" gorm:"not null;unique"`
}

func (b *Book) Save(db *gorm.DB) error {
	err := db.Create(&b).Error
	if err != nil && !errors.Is(err, gorm.ErrDuplicatedKey) {
		return errors.New("could not save book" + err.Error())
	}
	return nil
}

func GetAll(db *gorm.DB) ([]Book, error) {
	var books []Book
	err := db.Find(&books).Error
	if err != nil {
		return books, errors.New("books not found")
	}
	return books, nil
}
