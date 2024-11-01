package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	group := server.Group("/")

	group.GET("/books", getBooks)
	group.GET("/books/:name", getBookByName)
	group.POST("/books", AddNewBook)
	group.PUT("/books/:name", UpdateBook)
}
