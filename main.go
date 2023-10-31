package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tmbrody/taskquire/config"
	"github.com/tmbrody/taskquire/internal/database"
	"github.com/tmbrody/taskquire/middleware"
	"github.com/tmbrody/taskquire/routes"

	_ "github.com/go-sql-driver/mysql"
)

// main function initializes the server by loading environment variables, setting up the database connection,
// configuring trusted proxies, setting up middleware, and defining routes.
// It then starts the server on the specified port.
func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get port from environment variable
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	// Get trusted proxies from environment variable
	proxies := os.Getenv("TRUSTED_PROXIES")
	if proxies == "" {
		log.Fatal("TRUSTED_PROXIES environment variable is not set")
	}

	// Get JWT secret from environment variable
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is not set")
	}

	// Get database URL from environment variable and open database connection
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	// Close database connection when main function exits
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}()

	// Create database queries object and set it in the API config
	dbQueries := database.New(db)
	config.ApiCfg.DB = dbQueries
	config.ApiCfg.JwtSecret = jwtSecret

	// Create gin router with default middleware
	r := gin.Default()

	// Set trusted proxies for gin router
	trustedProxies := []string{proxies}
	err = r.SetTrustedProxies(trustedProxies)
	if err != nil {
		panic(err)
	}

	// Use middleware to set up database connection for each request
	r.Use(middleware.DatabaseSetupMiddleware())

	// Create API router group and set up routes
	apiRouter := r.Group("/api")
	routes.SetupRoutes(r)
	routes.SetupApiRoutes(apiRouter)

	// Start server on specified port
	err = r.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
