package main

import (
	"github.com/gin-gonic/gin"

	"github.com/silasbrasil/first-api-go/controllers"
	"github.com/silasbrasil/first-api-go/models"
)

func setupRouter() *gin.Engine {
	return gin.Default()
}

func listBooksRoute(r *gin.Engine) *gin.Engine {
	r.GET("/books", controllers.ListBooks)

	return r
}

func createBookRoute(r *gin.Engine) *gin.Engine {
	r.POST("/books", controllers.CreateBook)

	return r
}

func getBookByIdRoute(r *gin.Engine) *gin.Engine {
	r.GET("/books/:id", controllers.GetBookById)

	return r
}

func updateBookRoute(r *gin.Engine) *gin.Engine {
	r.PATCH("/books/:id", controllers.UpdateBook)

	return r
}

func deleteBookRoute(r *gin.Engine) *gin.Engine {
	r.DELETE("/books/:id", controllers.DeleteBook)

	return r
}

func main() {
	models.ConnectDatabase()

	r := setupRouter()
	listBooksRoute(r)
	createBookRoute(r)
	getBookByIdRoute(r)
	updateBookRoute(r)
	deleteBookRoute(r)

	r.Run()
}
