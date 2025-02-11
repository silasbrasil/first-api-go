package main

import (
	"github.com/gin-gonic/gin"

	"github.com/silasbrasil/first-api-go/controllers"
	"github.com/silasbrasil/first-api-go/models"
)

func main() {
	r := gin.Default()

	r.GET("/books", controllers.ListBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.GetBookById)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	models.ConnectDatabase()

	r.Run()
}
