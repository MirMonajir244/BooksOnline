package main

import (
	"github.com/MirMonajir244/BooksOnline/db"
	"github.com/MirMonajir244/BooksOnline/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		log.Fatal("Unable to run the servers", err)
	}
}
