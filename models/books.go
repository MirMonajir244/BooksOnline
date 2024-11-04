package models

import (
	"errors"
	"gorm.io/gorm"
)

type Book struct {
	Name        string  `json:"name" gorm:"not null;unique;primaryKey"`
	Description string  `json:"description"`
	Author      string  `json:"author" gorm:"not null"`
	Price       float64 `json:"price"`
	UserID      int64   `json:"userID" gorm:"not null;unique;"`
	Filename    []byte  `json:"file" gorm:"type:bytea"`
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

func UpdateBook(db *gorm.DB, name string, updatedBook Book) error {
	var existingBook Book

	// Check if the book exists
	if err := db.Where("name = ?", name).First(&existingBook).Error; err != nil {
		return errors.New("book not found")
	}

	// Perform the update on the Book model
	result := db.Model(&existingBook).Updates(Book{
		Name:        updatedBook.Name,
		Author:      updatedBook.Author,
		Price:       updatedBook.Price,
		Description: updatedBook.Description,
	})

	// Check for errors
	if result.Error != nil {
		return errors.New("could not update book: " + result.Error.Error())
	}

	// Check if any rows were affected
	if result.RowsAffected == 0 {
		return errors.New("no changes made; values may be the same")
	}

	return nil // Return nil if everything went fine
}

func DeleteBook(db *gorm.DB, name string) error {
	var existingBook Book
	if err := db.Where("name = ?", name).First(&existingBook).Error; err != nil {
		return errors.New("book not found")
	}
	if err := db.Where("name = ?", name).Delete(&existingBook).Error; err != nil {
		return errors.New("book not found")
	}
	return nil
}
