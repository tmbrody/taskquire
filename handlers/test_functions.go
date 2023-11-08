package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tmbrody/taskquire/config"
	"github.com/tmbrody/taskquire/internal/database"
)

func SetupTestDatabaseMiddleware(mockDBTX *MockDBTX) gin.HandlerFunc {
	queries := database.New(mockDBTX)
	return func(c *gin.Context) {
		c.Set(string(config.DbContextKey), queries)
		c.Next()
	}
}
