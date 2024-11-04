package routes

import (
	"github.com/MirMonajir244/BooksOnline/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/books", getBooks)
	server.GET("/books/:name", getBookByName)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/books", AddNewBook)
	authenticated.PUT("/books/:name", UpdateBookByName)
	authenticated.DELETE("/books/:name", DeleteBookByName)

	auth := server.Group("/auth")
	auth.POST("SignUp", SignUPUser)
	auth.POST("/Login", Login)
}
