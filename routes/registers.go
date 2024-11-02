package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	group := server.Group("/")
	auth := server.Group("/auth")

	group.GET("/books", getBooks)
	group.GET("/books/:name", getBookByName)
	group.POST("/books", AddNewBook)
	group.PUT("/books/:name", UpdateBookByName)
	group.DELETE("/books/:name", DeleteBookByName)
	auth.POST("SignUp", SignUPUser)
	auth.POST("/Login", Login)
}
