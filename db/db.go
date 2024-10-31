package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	conn_str := "user=mit password=ssl12345 dbname=BooksOnline sslmode=disable"
	DB, err = sql.Open("postgres", conn_str)
	if err != nil {
		panic("Could not connect to database")
	}
	CreateTables()
}

func CreateTables() {
	createTable := `CREATE TABLE IF NOT EXISTS Books (
    Name TEXT NOT NULL UNIQUE,
    Description TEXT,
    Author TEXT NOT NULL,
    Price NUMERIC NOT NULL,
    UserID TEXT UNIQUE NOT NULL,
    PRIMARY KEY (UserID)
)`
	_, err := DB.Exec(createTable)
	if err != nil {
		panic(err)
	}
}
