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
	defer db.DB.Close()
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
	books, err := models.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "count not fetch data try again"})
	}
	ctx.JSON(http.StatusOK, books)
}

func AddNewBook(ctx *gin.Context) {
	var Books models.Book
	err := ctx.ShouldBindJSON(&Books)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Could Not Parse the Books": err})
		return
	}
	err2 := Books.Save()
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Could Not create the Books": err2.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Added New Book", "Books": Books})
}
