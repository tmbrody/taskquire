package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/tmbrody/taskquire/config"
)

// DatabaseSetupMiddleware is a middleware function that sets the database connection
func DatabaseSetupMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(string(config.DbContextKey), config.ApiCfg.DB)
		c.Next()
	}
}
