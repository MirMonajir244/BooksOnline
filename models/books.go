package models

import (
	"errors"
	"github.com/MirMonajir244/BooksOnline/db"
)

type Book struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Author      string  `json:"author" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	UserID      string  `json:"userID" binding:"required"`
}

func (b *Book) Save() error {
	if db.DB == nil {
		return errors.New("db is not initialized")
	}
	query := `INSERT INTO Books(name, description, author, price, userID) VALUES($1, $2, $3, $4, $5)`
	re, err := db.DB.Prepare(query)
	if err != nil {
		return errors.New(err.Error())
	}
	_, err2 := re.Exec(b.Name, b.Description, b.Author, b.Price, b.UserID)
	if err2 != nil {
		return err
	}
	return nil
}

func GetAll() ([]Book, error) {
	query := `SELECT * FROM Books`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.Name, &book.Description, &book.Author, &book.Price, &book.UserID)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
