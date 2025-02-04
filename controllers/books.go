package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"book": "My tile here"})
}
