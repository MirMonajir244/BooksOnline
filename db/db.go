package db

import (
	"github.com/MirMonajir244/BooksOnline/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	conn_str := "host=localhost user=mit password=ssl12345 dbname=BooksOnline port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(conn_str), &gorm.Config{})
	if err != nil {
		panic("Could not connect to database")
	}

	err = DB.AutoMigrate(&models.Book{})
	if err != nil {
		return
	}
	err = DB.AutoMigrate(&models.Users{})
	if err != nil {
		return
	}
}
