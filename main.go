package main

import (
	"github.com/MirMonajir244/BooksOnline/db"
	"github.com/MirMonajir244/BooksOnline/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	db.InitDB()
	server := gin.Default()
	group := server.Group("/")

	group.GET("/books", getBooks)
	group.POST("/books", AddNewBook)
	err := server.Run(":8080")
	if err != nil {
		log.Fatal("Unable to run the servers", err)
	}
}

func getBooks(ctx *gin.Context) {
	books, err := models.GetAll(db.DB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "count not fetch data try again"})
	}
	ctx.JSON(http.StatusOK, books)
}

func AddNewBook(ctx *gin.Context) {
	var books models.Book
	// Bind JSON input directly to the book struct
	if err := ctx.ShouldBindJSON(&books); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err2 := books.Save(db.DB)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Could Not create the Books": err2.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Added New Book", "Books": books})
}
