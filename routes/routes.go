package routes

import (
	"github.com/MirMonajir244/BooksOnline/db"
	"github.com/MirMonajir244/BooksOnline/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getBooks(ctx *gin.Context) {
	books, err := models.GetAll(db.DB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "count not fetch data try again"})
	}
	ctx.JSON(http.StatusOK, books)
}

func getBookByName(ctx *gin.Context) {
	name := ctx.Param("name")
	books, err := models.GetAll(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "No book available"})
		return
	}
	for _, book := range books {
		if book.Name == name {
			ctx.JSON(http.StatusOK, book)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "No book available"})
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

func UpdateBook(ctx *gin.Context) {
	var books models.Book
	name := ctx.Param("name")
	ctx.ShouldBindJSON(&books)
	err := models.UpdateBook(db.DB, name, books)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update the book", "Error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Updated the Book", "Books": books})
}
