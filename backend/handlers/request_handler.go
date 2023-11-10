package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadinessHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

func ErrorHandler(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "Something went wrong",
	})
}

func HomeHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to the home page!",
	})
}

func AboutHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "This is the about page.",
	})
}
