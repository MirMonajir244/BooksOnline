package routes

import (
	"github.com/MirMonajir244/BooksOnline/db"
	"github.com/MirMonajir244/BooksOnline/models"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
)

func getBooks(ctx *gin.Context) {
	books, err := models.GetAll(db.DB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "count not fetch data try again"})
	}
	var bookResponses []models.Book
	for _, book := range books {
		bookResponses = append(bookResponses, models.Book{
			Name:        book.Name,
			Description: book.Description,
			Author:      book.Author,
			Price:       book.Price,
			UserID:      book.UserID,
		})
	}

	ctx.JSON(http.StatusOK, bookResponses)
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
	// Get the uploaded file
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "file field is required"})
		return
	}

	// Create a new Book instance
	var books models.Book

	// Extract other fields from form data
	books.Name = ctx.PostForm("name")
	books.Author = ctx.PostForm("author")
	books.Description = ctx.PostForm("description")
	books.UserID = ctx.PostForm("userID")

	// Parse the price field
	priceStr := ctx.PostForm("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid price"})
		return
	}
	books.Price = price

	// Open the uploaded file
	openedFile, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not open the file"})
		return
	}
	defer openedFile.Close()

	// Read the file into a byte slice
	fileData, err := io.ReadAll(openedFile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not read file"})
		return
	}

	// Set the File field in the book struct (if applicable)
	books.Filename = fileData // Assuming FilePath is the field for storing file data

	// Save the book to the database
	if err := books.Save(db.DB); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Could Not create the Books": err.Error()})
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
