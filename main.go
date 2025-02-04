package main

import (
	"github.com/gin-gonic/gin"

	"github.com/silasbrasil/first-api-go/controllers"
)

func main() {
	r := gin.Default()

	r.GET("/books", controllers.ListBooks)

	r.Run()
}
