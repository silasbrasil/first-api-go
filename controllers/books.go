package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silasbrasil/first-api-go/models"
)

func ListBooks(c *gin.Context) {
	var books []models.Book

	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"book": books})
}

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

func CreateBook(c *gin.Context) {
	var input CreateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Title}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}
